package main

type Dish struct {
	Title string   `json:"title"`
	Items []string `json:"items"`
}

type DailyMenu struct {
	MainDish Dish `json:"mainDish"`
	ColdCuts Dish `json:"coldCuts"`
	Salads   Dish `json:"salads"`
}

type WeeklyMenu [5]DailyMenu
