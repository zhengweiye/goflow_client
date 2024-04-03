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
	instanceResult, err := client.GetProcessInstanceService().Start(goflow_client.StartRequest{
		StartUserId:   "3",
		ProcessKey:    "leave2",
		AutoSubmit:    true,
		BusinessId:    "3",
		BusinessTitle: "郑伟业请假申请3",
		Fields: []goflow_client.Field{
			{Key: "days", Name: "请假天数", Value: "1天", SortNum: 1},
			{Key: "reason", Name: "请假理由", Value: "家里有事", SortNum: 2},
		},
		Variable: map[string]any{
			"days": 2,
		},
	})
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", instanceResult)
}

func main() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	err := client.GetProcessInstanceService().TxCommit("476c88d8-8413-4855-95b6-05eee245d3bb", "3")
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
