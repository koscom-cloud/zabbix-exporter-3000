package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server           string `mapstructure:"ZABBIX_API_ENDPOINT"`
	User             string `mapstructure:"ZABBIX_USER"`
	Password         string `mapstructure:"ZABBIX_PASSWORD"`
	SslSkip          bool   `mapstructure:"ZABBIX_SKIP_SSL"`
	StrictRegister   bool   `mapstructure:"ZE3000_STRICT_METRIC_REG"`
	SingleMetric     bool   `mapstructure:"ZE3000_SINGLE_METRIC"`
	SingleMetricHelp string `mapstructure:"ZE3000_SINGLE_METRIC_HELP"`
	MainHostPort     string `mapstructure:"ZE3000_HOST_PORT"`
	MetricNamespace  string `mapstructure:"ZE3000_METRIC_NAMESPACE"`
	MetricSubsystem  string `mapstructure:"ZE3000_METRIC_SUBSYSTEM"`
	MetricNamePrefix string `mapstructure:"ZE3000_METRIC_NAME_PREFIX"`
	MetricNameField  string `mapstructure:"ZE3000_METRIC_NAME_FIELD"`
	MetricValue      string `mapstructure:"ZE3000_METRIC_VALUE"`
	MetricHelpField  string `mapstructure:"ZE3000_METRIC_HELP"`
	MetricUriPath    string `mapstructure:"ZE3000_METRIC_URI_PATH"`
	SourceRefresh    string `mapstructure:"ZE3000_ZABBIX_REFRESH_DELAY_SEC"`
	MetricLabels     string `mapstructure:"ZE3000_ZABBIX_METRIC_LABELS"`
	Query            string `mapstructure:"ZE3000_ZABBIX_QUERY"`
}

var Cnf Config

func loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("zabbix-exporter-config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func init() {
	fmt.Println(Cnf)
	var err error
	configPath := getEnv("ZABBIX_EXPORTER_CONFIG_PATH","./")
	Cnf, err = loadConfig(configPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Println(Cnf)
}

func getEnv(key, fallback string) string {
	value, exist := os.LookupEnv(key)

	if !exist {
		log.Print("Config ", key, ": ", fallback)
		return fallback
	}
	if key != "ZABBIX_PASSWORD" {
		log.Print("Config ", key, ": ", value)
	}
	return value
}
