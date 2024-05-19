package models

type Powerstats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

type Biography struct {
	FullName string `json:"fullName"`
}

type Images struct {
	XS string `json:"xs"`
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
}

type Superhero struct {
	Name       string     `json:"name"`
	Powerstats Powerstats `json:"powerstats"`
	Biography  Biography  `json:"biography"`
	Images     Images     `json:"images"`
}
