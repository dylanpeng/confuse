package config

type Config struct {
	DBConfigs []*DBConfig
}

type DBConfig struct {
	Name         string
	UserName     string
	Password     string
	SourceUrl    string
	Port         string
	DataBaseName string
}
