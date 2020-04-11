package types

type GifRequest struct {
	Bucket    string `yaml:"bucket" json:"bucket"`
	KeyPrefix string `yaml:"key_prefix" json:"key_prefix"`
	Year      int    `yaml:"year" json:"year"`
	Month     int    `yaml:"month" json:"month"`
	Day       int    `yaml:"day" json:"day"`
}
