package config

type Config struct {
	Cutover struct {
		InstanceName string `yaml:"instancename" env:"" env-description:"Cutover Instance Name (Non_prod, UAT or PROD"`
		BaseURL      string `yaml:"baseurl"`
		AppToken     string `yaml:"apptoken"`
		WorkSpace    string `yaml:"workspace"`
	} `yaml:"cutover"`
}
