package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 配置
type Server struct {
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	Redis Redis `json:"redis" yaml:"redis"`
	Logger Logger `json:"logger" yaml:"logger"`
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

var CONFIG *Server
var VP *viper.Viper

var FetcherList map[string]struct{}


const defaultConfigFile = "config/config.yaml"

func init()  {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: #{err}\n"))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}

	VP = v
}