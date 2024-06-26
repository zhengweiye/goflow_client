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
		//Host:      "http://127.0.0.1:8888",
		Host:      "https://zhifa.cftzqinzhou.com",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	_, err := client.GetTaskService().Execution(goflow_client.ExecutionRequest{
		TaskId:            "27d24a50-3d13-4776-a66c-4c9aa1c8a2fc",
		UserId:            "23",
		HandleResult:      "pass",
		HandleOpinion:     "同意了",
		Variable:          map[string]any{},
		NextHandleUserIds: []string{},
	})
	fmt.Println("异常：", err)
}
