package main

import (
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func attemptLogin(
	page *rod.Page,
	email *string,
	password *string,
	logger *log.Logger,
) {
	logger.Println("entering email...")
	emailInput := page.
		MustElement(`input[type="email"]`).
		MustInput(*email)
	page.MustWaitStable()
	emailInput.MustType(input.Enter)
	page.MustWaitStable()

	logger.Println("entering password...")
	page.
		MustElement(`input[type="password"]`).
		MustWaitStable().
		MustInput(*password).
		MustWaitStable().
		MustType(input.Enter).
		MustWaitStable()

	logger.Println("opening weekly menu page...")
	page.
		MustElementX(`//button[div/span[text() = "Menu"]]`).
		MustClick().
		MustWaitStable()
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
