package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)
type Config struct {
	Server							string `mapstructure:"ZABBIX_API_ENDPOINT"`
	User								string `mapstructure:"ZABBIX_USER"`
	Password						string `mapstructure:"ZABBIX_PASSWORD"`
	SslSkip							bool `mapstructure:"ZABBIX_SKIP_SSL"`
	StrictRegister 			bool `mapstructure:"ZE3000_STRICT_METRIC_REG"`
	SingleMetric 				bool `mapstructure:"ZE3000_SINGLE_METRIC"`
	SingleMetricHelp		string `mapstructure:"ZE3000_SINGLE_METRIC_HELP"`
	MainHostPort				string `mapstructure:"ZE3000_HOST_PORT"`
	MetricNamespace  		string `mapstructure:"ZE3000_METRIC_NAMESPACE"`
	MetricSubsystem  		string `mapstructure:"ZE3000_METRIC_SUBSYSTEM"`
	MetricNamePrefix 		string `mapstructure:"ZE3000_METRIC_NAME_PREFIX"`
	MetricNameField  		string `mapstructure:"ZE3000_METRIC_NAME_FIELD"`
	MetricValue      		string `mapstructure:"ZE3000_METRIC_VALUE"`
	MetricHelpField  		string `mapstructure:"ZE3000_METRIC_HELP"`
	MetricUriPath    		string `mapstructure:"ZE3000_METRIC_URI_PATH"`
	SourceRefresh    		string `mapstructure:"ZE3000_ZABBIX_REFRESH_DELAY_SEC"`
	MetricLabels     		string `mapstructure:"ZE3000_ZABBIX_METRIC_LABELS"`
	Query            		string `mapstructure:"ZE3000_ZABBIX_QUERY"`

}
var Cnf Config

func loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("zabbix-exporter-config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
			return
	}

	err = viper.Unmarshal(&config)
	return
}

func init(){
	fmt.Println(Cnf)
	var err error
	Cnf, err = loadConfig("./config/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Println(Cnf)
}

// var (
// 	Server     = getEnv("ZABBIX_API_ENDPOINT", "http://zabbix/api_jsonrpc.php")
// 	User       = getEnv("ZABBIX_USER", "admin")
// 	Password   = getEnv("ZABBIX_PASSWORD", "admin")
// 	SslSkip, _ = strconv.ParseBool(getEnv("ZABBIX_SKIP_SSL", "true"))
// 	//if true = panic on duplicate metric, if false = results may vary, better check query or choose unique ZE3000_METRIC_NAME_FIELD
// 	// or use ZE3000_STRICT_METRIC_WKAROUND = true it adds _%num% at the end of metric name
// 	StrictRegister, _ = strconv.ParseBool(getEnv("ZE3000_STRICT_METRIC_REG", "true"))
// 	SingleMetric, _   = strconv.ParseBool(getEnv("ZE3000_SINGLE_METRIC", "false"))
// 	SingleMetricHelp  = getEnv("ZE3000_SINGLE_METRIC_HELP", "single description")

// 	MainHostPort     = getEnv("ZE3000_HOST_PORT", "localhost:8080")
// 	MetricNamespace  = getEnv("ZE3000_METRIC_NAMESPACE", "zbx")
// 	MetricSubsystem  = getEnv("ZE3000_METRIC_SUBSYSTEM", "subsystem")
// 	MetricNamePrefix = getEnv("ZE3000_METRIC_NAME_PREFIX", "prefix")
// 	MetricNameField  = getEnv("ZE3000_METRIC_NAME_FIELD", "key_")
// 	MetricValue      = getEnv("ZE3000_METRIC_VALUE", "lastvalue")
// 	MetricHelpField  = getEnv("ZE3000_METRIC_HELP", "description")
// 	MetricUriPath    = getEnv("ZE3000_METRIC_URI_PATH", "/metrics")
// 	SourceRefresh    = getEnv("ZE3000_ZABBIX_REFRESH_DELAY_SEC", "10")
// 	MetricLabels     = strings.TrimSpace(getEnv("ZE3000_ZABBIX_METRIC_LABELS", "name,itemid,key_,hosts>host,hosts>name,interfaces>ip,interface>dns"))
// 	Query            = getEnv("ZE3000_ZABBIX_QUERY", `{     "jsonrpc": "2.0",     "method": "item.get",     "params": {     	"application":"My Super Application",         "output": ["itemid","key_","description","lastvalue"],         "selectDependencies": "extend",         "selectHosts": ["name","status","host"],         "selectInterfaces": ["ip","dns"],         "sortfield":"key_"     },     "auth": "%auth-token%",     "id": 1 }`)
// )

// func getEnv(key, fallback string) string {
// 	value, exist := os.LookupEnv(key)

// 	if !exist {
// 		log.Print("Config ", key, ": ", fallback)
// 		return fallback
// 	}
// 	if key != "ZABBIX_PASSWORD" {
// 		log.Print("Config ", key, ": ", value)
// 	}
// 	return value
// }
