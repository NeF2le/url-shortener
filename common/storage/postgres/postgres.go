package postgres

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	Username string `yaml:"user" env:"USER" env-default:"postgres"`
	Password string `yaml:"password" env:"PASSWORD" env-default:"postgres"`
	Database string `yaml:"db" env:"DB" env-default:"postgres"`
	MaxConns int    `yaml:"max_conns" env:"MAX_CONNS" env-default:"10"`
	MinConns int    `yaml:"min_conns" env:"MIN_CONNS" env-default:"5"`
}
