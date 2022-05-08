package model

type ServerConfig struct {
	Name     string     `mapstructure:"name"`
	Port     int        `mapstructure:"port"`
	GormInfo GormConfig `mapstructure:"gorm"`
}

type GormConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}
