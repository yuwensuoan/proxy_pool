package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"proxy_pool/app/fetcher"
)

// 配置
type Server struct {
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	Redis Redis `json:"redis" yaml:"redis"`
	Logger Logger `json:"logger" yaml:"logger"`
	App App `json:"app" yaml:"app"`
	Sqlite3 `json:"sqlite3" yaml:"sqlite3"`
}

type App struct {
	Debug bool `json:"debug" yaml:"debug"`
	DbType string `json:"db_type" yaml:"dbtype"`
	ProxyVerify string `json:"proxy_verify" yaml:"proxyverify"`
	Addr string `json:"addr" yaml:"addr"`
}

// Mysql配置
type Mysql struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Host string `json:"host" yaml:"host"`
	Port int `json:"port" yaml:"port"`
	Dbname string `json:"dbname" yaml:"dbname"`
	Charset string `json:"charset" yaml:"charset"`
}

// Redis配置
type Redis struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
	DB int `json:"db" yaml:"db"`
}

// 日志配置
type Logger struct {
	Filepath string `json:"filepath" yaml:"filepath"`
	Filename string `json:"filename" yaml:"filename"`
}

type Sqlite3 struct {
	Database string `json:"database" yaml:"database"`
}

var CONFIG *Server
var VP *viper.Viper
var FetcherList []interface{}


const defaultConfigFile = "config/config.yml"

// 配置初始化
func init()  {
	// 需要执行的爬虫代理
	FetcherList = []interface{}{
		fetcher.CloudFetcher{},
		fetcher.SyrahFetcher{},
	}

	VP = viper.New()
	VP.SetConfigFile(defaultConfigFile)

	err := VP.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: #{err}\n"))
	}

	VP.WatchConfig()

	VP.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := VP.Unmarshal(&CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := VP.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}
}