package main

import (
	"fmt"
	"github.com/zhengweiye/goflow_client"
)

func main() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "prod",
	})
	userId := "16"
	instanceResult, err := client.GetProcessInstanceService().Start(goflow_client.StartRequest{
		StartUserId:   userId,
		ProcessKey:    "score_explain",
		BusinessId:    "1",
		BusinessTitle: "积分申述申请",
		NextUserIds:   []string{"10", "100"},
	})
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", instanceResult)
	client.GetProcessInstanceService().TxCommit(instanceResult.Id, userId)
}

func main2() {
	client := goflow_client.Create(goflow_client.Option{
		//Host:      "http://127.0.0.1:8888",
		Host:      "https://zhifa.cftzqinzhou.com",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	err := client.GetProcessInstanceService().TxCommit("8794c7c4-be04-4d3d-b7ec-c28e51a9947e", "16")
	fmt.Println("异常：", err)
}

func main3() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	err := client.GetProcessInstanceService().TxRollback("7d50e076-e41d-4fb9-b688-7b70feef868b", "3")
	fmt.Println("异常：", err)
}
