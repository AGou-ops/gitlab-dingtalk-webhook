package app

import (
	"fmt"

	"github.com/AGou-ops/gitlab-dingtalk-webhook/dingtalk"
	"github.com/AGou-ops/gitlab-dingtalk-webhook/gitlab"
	"github.com/gin-gonic/gin"
)

func StartService(
	env Env,
	r *gin.Engine,
	wb *gitlab.Webhook,
	robot *dingtalk.Robot,
) error {
	// 使用中间件
	r.Use(MethodNotAllowed())
	r.POST(env.Path, func(ctx *gin.Context) {
		payload, err := wb.Parse(
			ctx,
			gitlab.PushEvents,
			gitlab.MergeRequestEvents,
			gitlab.CommentEvents,
		)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				fmt.Println("err: event not found!")
			}
		}
		sendMsg(robot, payload)
	})
	return nil
}
