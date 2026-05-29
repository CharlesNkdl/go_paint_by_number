package config

type Config struct {
	ImagePath string
}

func NewConfig() *Config {
	return &Config{
		ImagePath: "/Users/charlesnikdel/code/Perso/Go/go_paint_by_number/test_case/SCR-20260327-kvpq.png",
	}
}
