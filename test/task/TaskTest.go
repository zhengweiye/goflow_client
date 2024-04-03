package main

import (
	"fmt"
	"github.com/zhengweiye/goflow_client"
)

func main1() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	err := client.GetTaskService().ReadCc("0d04b107-ed97-4d9d-ba90-8f47ddc05e54", "6")
	fmt.Println("异常：", err)
}

func main() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	instanceResult, err := client.GetTaskService().Execution(goflow_client.ExecutionRequest{
		TaskId:        "b1ecb8a4-776d-4ecf-afda-50cc037a73c0",
		UserId:        "2",
		HandleResult:  "pass",
		HandleOpinion: "同意了",
		Variable: map[string]any{
			"days": 2,
		},
		NextHandleUserIds: nil,
	})
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", instanceResult)
}
