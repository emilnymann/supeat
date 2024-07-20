package main

import (
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
)

func attemptLogin(
	browser *rod.Browser,
	email *string,
	password *string,
	logger *log.Logger,
) (*rod.Page, error) {
	var err error
	handleError := func(
		err error, message string,
	) bool {
		if err != nil {
			logger.Printf("%s: %v", message, err)
			return true
		}

		return false
	}

	page, err := func() (*rod.Page, error) {
		page, err := browser.Page(
			proto.TargetCreateTarget{URL: "https://app.kanpla.dk/login"},
		)
		if handleError(err, "Failed to open login page") {
			return nil, err
		}

		err = page.WaitStable(time.Second)
		if handleError(err, "Page did not stabilize") {
			return nil, err
		}

		emailInput, err := page.Element(
			`input[type="email"]`,
		)
		if handleError(err, "Failed to find email input element") {
			return nil, err
		}

		err = emailInput.Input(*email)
		if handleError(err, "Failed to input email") {
			return nil, err
		}

		err = emailInput.Type(input.Enter)
		if handleError(err, "Failed to press Enter after email input") {
			return nil, err
		}

		err = emailInput.WaitStable(time.Second)
		if handleError(err, "Email input did not stabilize") {
			return nil, err
		}

		passwordInput, err := page.Element(
			`input[type="password"]`,
		)
		if handleError(err, "Failed to find password input element") {
			return nil, err
		}

		err = passwordInput.Input(*password)
		if handleError(err, "Failed to input password") {
			return nil, err
		}

		err = passwordInput.Type(input.Enter)
		if handleError(err, "Failed to press Enter after password input") {
			return nil, err
		}

		err = page.WaitStable(time.Second)
		if handleError(err, "Page did not stabilize after password input") {
			return nil, err
		}

		menuButton, err := page.ElementX(
			`//button[div/span[text() = "Menu"]]`,
		)
		if handleError(err, "Failed to find menu button") {
			return nil, err
		}

		err = menuButton.Click(proto.InputMouseButtonLeft, 1)
		if handleError(err, "Failed to click menu button") {
			return nil, err
		}

		err = page.WaitStable(2 * time.Second)
		if handleError(err, "Page did not stabilize after clicking menu button") {
			return nil, err
		}

		return page, nil
	}()

	return page, err
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
