package main

import (
	"fmt"
	"strings"

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
	fmt.Println(env.Secret)

	r := gin.Default()
	r.Use(app.MethodNotAllowed())
	r.POST(env.Path, func(ctx *gin.Context) {

		payload, err := webhook.Parse(
			ctx,
			gitlab.PushEvents,
			gitlab.MergeRequestEvents,
		)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				fmt.Println("err: event not found!")
			}
		}

		switch payload := payload.(type) {
		case gitlab.PushEventPayload:
			push := payload
			fmt.Println(push.UserName)
			fmt.Println(strings.Repeat("=", 100))
			fmt.Println("send message to dingtalk")
			dingtalk_robot.SendTextMessage(fmt.Sprintf("%+v", push), []string{"15628960878"}, false)
		}
	})
	r.Run(fmt.Sprintf(":%d", env.ListenPort))
}
