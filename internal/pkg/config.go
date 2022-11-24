package pkg

type ServerHTTP struct {
	BindHTTPAddr       string `toml:"bind_addr_http"`
	ReadTimeout        int    `toml:"read_timeout"`
	WriteTimeout       int    `toml:"write_timeout"`
	Protocol           string `toml:"protocol"`
	FileTLSCertificate string `toml:"tls_certificate_file"`
	FileTLSKey         string `toml:"tls_key_file"`
}

type ServerGRPC struct {
	BindHTTPAddr      string `toml:"bind_addr_http"`
	ConnectionTimeout int    `toml:"connection_timeout"`
}

type URls struct {
	ImageServiceURL string `toml:"image_service"`
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
	Endpoint string `toml:"endpoint"`
}

type Config struct {
	ServerHTTP ServerHTTP `toml:"server_http"`
	ServerGRPC ServerGRPC `toml:"server_grpc"`
	Cors       Cors       `toml:"cors"`
	S3         S3         `toml:"S3"`
	Context    Context    `toml:"context"`
	Logger     Logger     `toml:"logger"`
	URls       URls       `toml:"URLs"`
}

func NewConfig() *Config {
	return &Config{}
}
