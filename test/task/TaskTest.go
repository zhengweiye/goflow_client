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
		Env:       "prod",
	})
	_, err := client.GetTaskService().Execution(goflow_client.ExecutionRequest{
		TaskId:        "c1c00070-f904-4822-8f0e-b2161f66adcf",
		UserId:        "101",
		HandleResult:  "pass",
		HandleOpinion: "通过",
		Variable:      map[string]any{},
	})
	fmt.Println("异常：", err)
}
