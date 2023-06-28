package app

import (
	"flag"
	"log"
	"os"
	"reflect"

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
		log.SetPrefix("[INFO] ")
		log.Println("Config File NOT FOUND!!!Use system environment instead.")
		// listen_port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Println("Cannot convert variable PORT to int.")
		}
		env = &Env{
			Path:       os.Getenv("WB_PATH"),
			Token:      os.Getenv("TOKEN"),
			Secret:     os.Getenv("SECRET"),
			ListenPort: *port,
		}
		// 如果没有环境变量，或者环境变量为空，抛出错误日志并退出程序.
		if env.Path == "" || env.Token == "" || env.Secret == "" {
			log.SetPrefix("[ERROR] ")
			// 使用反射遍历env结构体，找出哪个字段为空
			envType := reflect.TypeOf(env).Elem()
			// 遍历结构体字段
			for i := 0; i < envType.NumField(); i++ {
				field := envType.Field(i)
				value := reflect.ValueOf(env).Elem().Field(i).Interface()
				// 判断字段值是否为空
				if value == "" {
					log.Println(field.Name, "field is empty")
				}
			}
			log.Fatal("Please checkout your system env and try again.")
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
