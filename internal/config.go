package internal

type Server struct {
	BindHttpAddr string `toml:"bind_addr_http"`
	ReadTimeout  int    `toml:"read_timeout"`
	WriteTimeout int    `toml:"write_timeout"`
}

type Cors struct {
	Methods     []string `toml:"methods"`
	Origins     []string `toml:"urls"`
	Headers     []string `toml:"headers"`
	Credentials bool     `toml:"credentials"`
	Debug       bool     `toml:"debug"`
}

type Context struct {
	Timeout int `toml:"timeout"`
}

type Config struct {
	Server  Server  `toml:"server"`
	Cors    Cors    `toml:"cors"`
	Context Context `toml:"context"`
}

func NewConfig() *Config {
	return &Config{}
}
