package types

type Config struct {
	Polygon struct {
		Addr string `yaml:"addr"`
	} `yaml:"polygon"`

	Databases struct {
		Postgres string `yaml:"postgres"`
	} `yaml:"databases"`
}
