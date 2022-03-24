package types

// This struct will be used in the configuration struct as a way to
// specify all the connection strings to databases.
type Databases struct {
	Redis    string `yaml:"redis"`
	Postgres string `yaml:"postgres"`
}

// This struct will be used in the configuration struct as a way to
// specify all the internal configurations related to Polygon and
// customize its behavior.
type Polygon struct {
	Addr string `yaml:"addr"`
}

// This is the struct which will be used to parse the configuration
// supplied from a YAML file. The configuration is going to help us
// configure the application as well as the behavior that we want to
// achieve on our Polygon instance.
type Config struct {
	Polygon   Polygon   `yaml:"polygon"`
	Databases Databases `yaml:"databases"`
}
