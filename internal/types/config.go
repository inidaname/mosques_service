package types

type Config struct {
	Server struct {
		GRPCPort string `yaml:"grpc_port"`
		HTTPPort string `yaml:"http_port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Schema   string `yaml:"schema"`
		SSLMode  string `yaml:"sslmode"`
		Url      string `yaml:"url"`
	} `yaml:"database"`
	DBurl  string
	APIurl string
}
