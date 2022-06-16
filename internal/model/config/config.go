package config

type All struct {
	Mysql Mysql `yaml:"Mysql"`
}

type Mysql struct {
	DriverSourceName string `yaml:"DriverSourceName"`
	DataSourceName   string `yaml:"DataSourceName"`
}
