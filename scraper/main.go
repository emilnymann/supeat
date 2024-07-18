package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-rod/rod"
)

func formatLine(line string) string {
	// Remove trailing punctuation
	line = strings.TrimRightFunc(line, func(r rune) bool {
		return unicode.IsPunct(r)
	})

	line = strings.ToLower(line)

	return line
}

func scrapeWeeklyDishes(page *rod.Page, xpath string) []Dish {
	dishElements := page.MustElementsX(xpath)[0:10]
	return parseDishes(dishElements)
}

func parseDishes(elements rod.Elements) []Dish {
	dishes := make([]Dish, len(elements)/2)
	re := regexp.MustCompile(`\s*\r?\n\s*`)

	for i := 0; i < len(elements); i += 2 {
		title := formatLine(elements[i].MustText())
		items := re.Split(elements[i+1].MustText(), -1)

		for j := range items {
			items[j] = formatLine(items[j])
		}

		dishes[i/2] = Dish{
			Title: title,
			Items: items,
		}
	}

	return dishes
}

func main() {
	logger := log.New(os.Stdout, "Supeat Scraper: ", log.LstdFlags)

	logger.Println("reading environment variables...")
	loginEmail := os.Getenv("SUPEAT_KANPLA_EMAIL")
	loginPassword := os.Getenv("SUPEAT_KANPLA_PASSWORD")

	if (len(loginEmail) <= 0) || (len(loginPassword) <= 0) {
		logger.Println("email or password environment variables not provided, exiting...")
		os.Exit(1)
	}

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	logger.Println("attempting login...")
	page, err := attemptLogin(
		browser,
		&loginEmail,
		&loginPassword,
		logger,
	)
	if err != nil {
		logger.Println("failed to login -", err)
		os.Exit(1)
	}

	logger.Println("scraping main dishes...")
	weeklyMainDishes := scrapeWeeklyDishes(page, `//p[text()[contains(., "VARM RET")]]/following-sibling::p`)

	logger.Println("scraping cold cuts...")
	weeklyColdCuts := scrapeWeeklyDishes(page, `//p[text()[contains(., "PÅLÆG 1")]]/following-sibling::p`)

	logger.Println("scraping salads...")
	weeklySalads := scrapeWeeklyDishes(page, `//p[text()[contains(., "SALAT 1")]]/following-sibling::p`)

	logger.Println("structuring scraped dishes...")
	var dailyMenusThisWeek WeeklyMenu
	for i := 0; i < 5; i++ {
		dailyMenusThisWeek[i] = DailyMenu{
			MainDish: weeklyMainDishes[i],
			ColdCuts: weeklyColdCuts[i],
			Salads:   weeklySalads[i],
		}
	}

	logger.Println("finished!")

	jsonResult, err := json.Marshal(dailyMenusThisWeek)
	if err != nil {
		fmt.Println("error marshaling dishes to JSON:", err)
		os.Exit(1)
	}

	os.Stdout.Write(jsonResult)
}
