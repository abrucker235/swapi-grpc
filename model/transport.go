package model

type Transport struct {
	ID                   int32  `json:"id"`
	Name                 string `json:"name"`
	Model                string `json:"model"`
	Manufacturer         string `json:"manufacturer"`
	Passengers           string `json:"passengers"`
	Consumables          string `json:"consumables"`
	Length               string `json:"length"`
	Crew                 string `json:"crew"`
	CargoCapacity        string `json:"cargo_capacity"`
	CostInCredits        string `json:"cost_in_credits"`
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
	Created              string `json:"created"`
	Edited               string `json:"edited"`
}
