package app

import (
	"flag"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	Path       string
	Token      string
	Secret     string
	ListenPort int
}

func GetEnv() Env {
	viper.SetDefault("Path", os.Getenv("path"))
	viper.SetDefault("Token", os.Getenv("dingtalk_token"))
	viper.SetDefault("Secret", os.Getenv("dingtalk_secret"))

	// 设置服务默认监听端口, 可以从命令行参数传入
	port := flag.Int("p", 8787, "service listen port")
	flag.Parse()
	viper.SetDefault("ListenPort", *port)

	// 设置配置文件路径及文件名
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic("Config File NOT FOUND!!!")
	}

	env := Env{
		Path:       viper.GetString("Path"),
		Token:      viper.GetString("Token"),
		Secret:     viper.GetString("Secret"),
		ListenPort: viper.GetInt("ListenPort"),
	}
	return env
}
