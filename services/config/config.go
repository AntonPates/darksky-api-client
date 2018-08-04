package config

type Config interface {
	Load(fpath string) error
	Locations() []Location
	Output() (string, string)
	RawDir() string
}

var _ Config = New()

func New() Config {
	return &conf{}
}
