package model

type Form struct {
	Id       string `json:"id"`
	Secret   string `json:"secret"`
	Script   string `json:"script"`
	Title    string `json:"title"`
	Delivery string `json:"delivery"`
}

type Config struct {
	Title  string `toml:"title"`
	Id     string `toml:"id"`
	Secret string `toml:"secret"`
}
