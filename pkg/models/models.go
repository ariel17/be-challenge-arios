package models

type Person struct {
	Name        string
	Position    *string `json:"position"`
	DateOfBirth string  `json:"date_of_birth"`
	Nationality string  `json:"nationality"`
}

type Team struct {
	TLA       string   `json:"tla"`
	Name      string   `json:"name"`
	ShortName string   `json:"short_name"`
	AreaName  string   `json:"area_name"`
	Address   string   `json:"address"`
	Coach     *Person  `json:"coach,omitempty"`
	Players   []Person `json:"players,omitempty"`
}

type Competition struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	AreaName string `json:"area_name"`
	Teams    []Team `json:"teams"`
}