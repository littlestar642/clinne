package model

type Metadata struct {
	FilePath string `json:"filepath"`
	Answer   bool   `json:"answer"`
	Rules    string `json:"rules"`
}
