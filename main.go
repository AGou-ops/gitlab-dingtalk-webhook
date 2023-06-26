package main

import (
	"fmt"

	"github.com/AGou-ops/gitlab-dingtalk-webhook/app"
	"github.com/AGou-ops/gitlab-dingtalk-webhook/dingtalk"
	"github.com/AGou-ops/gitlab-dingtalk-webhook/gitlab"
	"github.com/gin-gonic/gin"
)

func main() {
	// 从环境变量或者配置文件中获取配置信息
	env := app.GetEnv()

	webhook, _ := gitlab.New(
		gitlab.Options.Secret("bjPyrYvx-hwwd1LSw8TS"),
	)
	dingtalk_robot := dingtalk.NewRobot(env.Token, env.Secret)

	r := gin.Default()
	err := app.StartService(*env, r, webhook, dingtalk_robot)
	if err != nil {
		fmt.Println(err)
	}
	r.Run(fmt.Sprintf(":%d", env.ListenPort))
}
