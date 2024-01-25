package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/AGou-ops/gitlab-dingtalk-webhook/dingtalk"
	"github.com/AGou-ops/gitlab-dingtalk-webhook/gitlab"
)

var (
	atMobiles = []string{
		"156xxxxxxxx",
	}
	isAtAll = false
)

func sendMsg(robot *dingtalk.Robot, payload interface{}) {
	switch payload := payload.(type) {
	case gitlab.MergeRequestEventPayload:
		sendMR2Dingtalk(robot, payload)
	case gitlab.CommentEventPayload:
		sendComment2Dingtalk(robot, payload)
	case gitlab.PushEventPayload:
		sendPush2Dingtalk(robot, payload)
	}
}

func sendPush2Dingtalk(
	robot *dingtalk.Robot, plain gitlab.PushEventPayload,
) {
	title := fmt.Sprintf("😀%s提交了push", plain.UserName)
	changedFiles := func(plain gitlab.PushEventPayload) string {
		var changedFileList string
		for _, commit := range plain.Commits {
			for _, item := range commit.Added {
				changedFileList += fmt.Sprintf("> ####  <font color='green'>- [A] %s</font> \n ", item)
			}
			for _, item := range commit.Modified {
				changedFileList += fmt.Sprintf("> ####  <font color='orange'>- [M] %s</font> \n ", item)
			}
			for _, item := range commit.Removed {
				changedFileList += fmt.Sprintf("> ####  <font color='red'>- [D] %s</font> \n ", item)
			}
		}
		return changedFileList
	}(plain)
	text := fmt.Sprintf(
		"### **%s** 提交了push \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交分支名称：[%s](%s) \n > #### 最新commitID为： [%s](%s) \n > #### 最新提交信息：%s \n > #### 最新变更文件： \n %s \n @%s",
		plain.UserName,
		plain.Project.Name,
		plain.Project.WebURL,
		plain.Ref,
		plain.Project.WebURL+"/-/tree/"+strings.TrimPrefix(plain.Ref, "refs/heads/"),
		plain.After[:8],
		plain.Project.WebURL+"/-/commit/"+plain.After,
		plain.Commits[len(plain.Commits)-1].Message,
		changedFiles,
		fmt.Sprint(strings.Join(atMobiles, "，@")),
	)
	if err := robot.SendMarkdownMessage(
		title,
		text,
		atMobiles,
		isAtAll,
	); err != nil {
		log.Println("Failed to send Markdown Message: ", err)
	}
}

func sendMR2Dingtalk(
	robot *dingtalk.Robot,
	plain gitlab.MergeRequestEventPayload,
) {
	title := fmt.Sprintf("😀%s发起了mergeRequest", plain.User.Name)
	var text string
	switch plain.ObjectAttributes.Action {
	case "open":
		text = fmt.Sprintf(
			"### **%s** 发起了一个%s \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n > #### MR链接地址：[点我直达](%s/-/merge_requests) \n @%s",
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
			plain.Project.WebURL,
			fmt.Sprint(strings.Join(atMobiles, "，@")),
		)
	case "update":
		text = fmt.Sprintf(
			"### **%s** 更新了%s \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 上次commitID：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n > #### MR链接地址：[点我直达](%s/-/merge_requests) \n @%s",
			plain.User.Name,
			plain.ObjectKind,
			plain.Repository.Name,
			plain.Repository.Homepage,
			plain.ObjectAttributes.LastCommit.Title,
			plain.ObjectAttributes.LastCommit.URL,
			plain.ObjectAttributes.Oldrev[:8],
			fmt.Sprintf(
				"http://git.nblh.local/nlp/Management/-/commit/%s",
				plain.ObjectAttributes.Oldrev,
			),
			plain.ObjectAttributes.SourceBranch,
			plain.ObjectAttributes.TargetBranch,
			plain.ObjectAttributes.Title,
			plain.ObjectAttributes.State,
			plain.Project.WebURL,
			fmt.Sprint(strings.Join(atMobiles, "，@")),
		)
	case "merge":
		text = fmt.Sprintf(
			"### **%s** 完成了%s合并 \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n @%s",
			plain.Reviewers[0].Name,
			plain.ObjectKind,
			plain.Repository.Name,
			plain.Repository.Homepage,
			plain.ObjectAttributes.LastCommit.Title,
			plain.ObjectAttributes.LastCommit.URL,
			plain.ObjectAttributes.SourceBranch,
			plain.ObjectAttributes.TargetBranch,
			plain.ObjectAttributes.Title,
			plain.ObjectAttributes.State,
			fmt.Sprint(strings.Join(atMobiles, "，@")),
		)
	case "close":
		text = fmt.Sprintf(
			"### **%s** 关闭了一个%s \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n @%s",
			plain.Reviewers[0].Name,
			plain.ObjectKind,
			plain.Repository.Name,
			plain.Repository.Homepage,
			plain.ObjectAttributes.LastCommit.Title,
			plain.ObjectAttributes.LastCommit.URL,
			plain.ObjectAttributes.SourceBranch,
			plain.ObjectAttributes.TargetBranch,
			plain.ObjectAttributes.Title,
			plain.ObjectAttributes.State,
			fmt.Sprint(strings.Join(atMobiles, "，@")),
		)
	case "reopen":
		text = fmt.Sprintf(
			"### **%s** 重新打开了一个%s \n --- \n > #### 项目名称：[%s](%s) \n > #### 提交信息：[%s](%s) \n > #### 合并分支：%s --> %s \n > #### MR标题名称：%s \n > #### MR当前状态：<font color='green'><b>%s</b></font> \n @%s",
			plain.Reviewers[0].Name,
			plain.ObjectKind,
			plain.Repository.Name,
			plain.Repository.Homepage,
			plain.ObjectAttributes.LastCommit.Title,
			plain.ObjectAttributes.LastCommit.URL,
			plain.ObjectAttributes.SourceBranch,
			plain.ObjectAttributes.TargetBranch,
			plain.ObjectAttributes.Title,
			plain.ObjectAttributes.State,
			fmt.Sprint(strings.Join(atMobiles, "，@")),
		)
	default:
		text = "未知action."
	}

	if err := robot.SendMarkdownMessage(
		title,
		text,
		atMobiles,
		isAtAll,
	); err != nil {
		log.Println("Failed to send Markdown Message: ", err)
	}
}

func sendComment2Dingtalk(
	robot *dingtalk.Robot,
	plain gitlab.CommentEventPayload,
) {
	title := fmt.Sprintf("😀%s发了一条Comment", plain.User.Name)
	text := fmt.Sprintf(
		"### **%s** 在[%s](%s) 发了一条评论 \n --- \n > #### %s",
		plain.User.Name,
		plain.Project.Name,
		plain.ObjectAttributes.URL,
		plain.ObjectAttributes.Note,
	)

	if err := robot.SendMarkdownMessage(
		title,
		text,
		atMobiles,
		isAtAll,
	); err != nil {
		log.Println("Failed to send Markdown Message: ", err)
	}
}
