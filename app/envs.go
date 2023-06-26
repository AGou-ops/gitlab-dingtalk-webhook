package app

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Env struct {
	Path       string
	Token      string
	Secret     string
	ListenPort int
}

func GetEnv() *Env {
	// 设置服务默认监听端口, 可以从命令行参数传入
	port := flag.Int("p", 8787, "service listen port")
	flag.Parse()
	viper.SetDefault("ListenPort", *port)

	// 设置配置文件路径及文件名
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")

	var env *Env
	if err := viper.ReadInConfig(); err != nil {
		log.SetPrefix("[INFO]")
		log.Println("Config File NOT FOUND!!!Use system environment instead.")
		listen_port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Println("Cannot convert variable PORT to int.")
		}
		env = &Env{
			Path:       os.Getenv("WB_PATH"),
			Token:      os.Getenv("TOKEN"),
			Secret:     os.Getenv("SECRET"),
			ListenPort: listen_port,
		}
		// 如果没有环境变量，或者环境变量为空，抛出错误日志并退出程序.
		if env.Path == "" || env.Token == "" || env.Secret == "" {
			log.SetPrefix("[ERROR]")
			log.Fatal("Please checkout your system env.")
		}
	} else {
		env = &Env{
			Path:       viper.GetString("Path"),
			Token:      viper.GetString("Token"),
			Secret:     viper.GetString("Secret"),
			ListenPort: viper.GetInt("ListenPort"),
		}
	}

	return env
}
