package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func formatLine(line string) string {
	// Remove trailing punctuation
	line = strings.TrimRightFunc(line, func(r rune) bool {
		return unicode.IsPunct(r)
	})

	line = strings.ToLower(line)

	return line
}

func login(browser *rod.Browser, email *string, password *string) *rod.Page {
	page := browser.MustPage("https://app.kanpla.dk/login").MustWaitStable()
	page.MustElement(`input[type="email"]`).MustInput(*email).MustType(input.Enter).MustWaitStable()
	page.MustElement(`input[type="password"]`).MustInput(*password).MustType(input.Enter).MustWaitStable()
	page.MustElementX(`//button[div/span[text() = "Menu"]]`).MustClick().MustWaitStable()
	return page
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
	loginEmail := os.Getenv("SUPEAT_KANPLA_EMAIL")
	loginPassword := os.Getenv("SUPEAT_KANPLA_PASSWORD")
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	log.Println("Logging in...")
	page := login(browser, &loginEmail, &loginPassword)

	log.Println("Scraping main dishes...")
	weeklyMainDishes := scrapeWeeklyDishes(page, `//p[text()[contains(., "VARM RET")]]/following-sibling::p`)

	log.Println("Scraping cold cuts...")
	weeklyColdCuts := scrapeWeeklyDishes(page, `//p[text()[contains(., "PÅLÆG 1")]]/following-sibling::p`)

	log.Println("Scraping salads...")
	weeklySalads := scrapeWeeklyDishes(page, `//p[text()[contains(., "SALAT 1")]]/following-sibling::p`)

	log.Println("Structuring scraped dishes...")
	var dailyMenusThisWeek WeeklyMenu
	for i := 0; i < 5; i++ {
		dailyMenusThisWeek[i] = DailyMenu{
			MainDish: weeklyMainDishes[i],
			ColdCuts: weeklyColdCuts[i],
			Salads:   weeklySalads[i],
		}
	}

	log.Println("Monday:")
	log.Println("Main dish:", dailyMenusThisWeek[0].MainDish.Items[0])
	log.Println("Cold cut:", dailyMenusThisWeek[0].ColdCuts.Items[0])
	log.Println("Salads:", dailyMenusThisWeek[0].Salads.Items[0])
}
