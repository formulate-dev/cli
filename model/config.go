package model

type Config struct {
	Title    string         `toml:"title"`
	Internal ConfigInternal `toml:"internal"`
}

type ConfigInternal struct {
	CustomId string `toml:"custom_id"`
	Id       string `toml:"id"`
	Secret   string `toml:"secret"`
}
