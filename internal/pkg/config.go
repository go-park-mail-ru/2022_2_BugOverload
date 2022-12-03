package pkg

type ServerHTTP struct {
	ServiceName        string `toml:"service_name"`
	BindAddr           string `toml:"bind_addr_http"`
	ReadTimeout        int    `toml:"read_timeout"`
	WriteTimeout       int    `toml:"write_timeout"`
	Protocol           string `toml:"protocol"`
	FileTLSCertificate string `toml:"tls_certificate_file"`
	FileTLSKey         string `toml:"tls_key_file"`
}

type Cors struct {
	Methods     []string `toml:"methods"`
	Origins     []string `toml:"urls"`
	Headers     []string `toml:"headers"`
	Credentials bool     `toml:"credentials"`
	Debug       bool     `toml:"api"`
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

type Metrics struct {
	BindHTTPAddr string `toml:"bind_addr_http"`
}

type ServerGRPC struct {
	ServiceName       string `toml:"service_name"`
	BindAddr          string `toml:"bind_addr_http"`
	ConnectionTimeout int    `toml:"connection_timeout"`
	WorkTimeout       int    `toml:"work_timeout"`
	URL               string `toml:"URL"`
}

type Config struct {
	ServerHTTPApi       ServerHTTP `toml:"server_http_api"`
	ServerGRPCImage     ServerGRPC `toml:"server_grpc_image"`
	ServerGRPCWarehouse ServerGRPC `toml:"server_grpc_warehouse"`
	ServerGRPCUser      ServerGRPC `toml:"server_grpc_user"`
	ServerGRPCAuth      ServerGRPC `toml:"server_grpc_auth"`
	Cors                Cors       `toml:"cors"`
	Metrics             Metrics    `toml:"metrics"`
	S3                  S3         `toml:"S3"`
	Context             Context    `toml:"context"`
	Logger              Logger     `toml:"logger"`
}

func NewConfig() *Config {
	return &Config{}
}
