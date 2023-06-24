package main

import (
	"fmt"
	"net/http"

	"github.com/AGou-ops/gitlab-dingtalk-webhook/gitlab"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := gitlab.New(
		gitlab.Options.Secret("bjPyrYvx-hwwd1LSw8TS"),
	)
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		payload, err := hook.Parse(
			r,
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

				fmt.Println("获取到push event")
				fmt.Printf("%v", push)
				w.Write([]byte("test push event"))
		}
	})
	http.ListenAndServe(":3000", nil)
}
