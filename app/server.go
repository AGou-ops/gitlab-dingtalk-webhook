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

		switch payload := payload.(type) {
		case gitlab.MergeRequestEventPayload:
			push := payload
			sendMsg2Dingtalk(robot, push)
		case gitlab.CommentEventPayload:
			push := payload
			sendComment2Dingtalk(robot, push)
		}
	})
	return nil
}

func sendMsg2Dingtalk(
	robot *dingtalk.Robot,
	plain gitlab.MergeRequestEventPayload,
) {
	title := "😀mergeRequest"
	var text string
	if plain.ObjectAttributes.State == "opened" {
		text = fmt.Sprintf(
			"### **%s** 发起了一个%s \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n @18557519596",
			plain.User.Name,
			plain.ObjectKind,
			plain.Repository.Name,
			plain.Repository.Homepage,
			plain.ObjectAttributes.LastCommit.Title,
			plain.ObjectAttributes.LastCommit.URL,
			plain.ObjectAttributes.SourceBranch,
			plain.ObjectAttributes.TargetBranch,
			plain.ObjectAttributes.Title,
			plain.ObjectAttributes.State,
		)
	} else {
		text = fmt.Sprintf(
			"### **%s** 关闭了一个%s \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n @18557519596",
			plain.Assignees[0].Name,
			plain.ObjectKind,
			plain.Repository.Name,
			plain.Repository.Homepage,
			plain.ObjectAttributes.LastCommit.Title,
			plain.ObjectAttributes.LastCommit.URL,
			plain.ObjectAttributes.SourceBranch,
			plain.ObjectAttributes.TargetBranch,
			plain.ObjectAttributes.Title,
			plain.ObjectAttributes.State,
		)
	}

	atMobiles := []string{
		"18557519596",
	}
	isAtAll := false
	robot.SendMarkdownMessage(
		title,
		text,
		atMobiles,
		isAtAll,
	)
}

func sendComment2Dingtalk(
	robot *dingtalk.Robot,
	plain gitlab.CommentEventPayload,
) {
	title := "😀Comment"
	text := fmt.Sprintf(
		"### **%s** 在[%s](%s) 发了一条评论 \n --- \n > #### %s",
		plain.User.Name,
		plain.Project.Name,
		plain.ObjectAttributes.URL,
		plain.ObjectAttributes.Note,
	)

	atMobiles := []string{
		"18557519596",
	}
	isAtAll := false
	robot.SendMarkdownMessage(
		title,
		text,
		atMobiles,
		isAtAll,
	)
}
