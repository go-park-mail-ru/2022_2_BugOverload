package pkg

type Server struct {
	BindHTTPAddr string `toml:"bind_addr_http"`
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

type Logger struct {
	LogLevel string `toml:"log_level"`
	LogAddr  string `toml:"log_path"`
}

type S3 struct {
	ID       string `toml:"aws_access_key_id"`
	Secret   string `toml:"aws_secret_access_key"`
	Region   string `toml:"region"`
	Endpoint string `toml:"endpoint"`
}

type Config struct {
	Server  Server  `toml:"server"`
	Cors    Cors    `toml:"cors"`
	S3      S3      `toml:"S3"`
	Context Context `toml:"context"`
	Logger  Logger  `toml:"logger"`
}

func NewConfig() *Config {
	return &Config{}
}
