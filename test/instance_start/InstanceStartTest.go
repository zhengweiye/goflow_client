package main

import (
	"fmt"
	"github.com/zhengweiye/goflow_client"
)

func main1() {
	client := goflow_client.Create(goflow_client.Option{
		//Host:      "http://127.0.0.1:8888",
		Host:      "https://zhifa.cftzqinzhou.com",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	instanceResult, err := client.GetProcessInstanceService().Start(goflow_client.StartRequest{
		StartUserId:   "16",
		ProcessKey:    "case_main",
		AutoSubmit:    true,
		BusinessId:    "111",
		BusinessTitle: "测试案件",
		Fields: []goflow_client.Field{
			{Key: "days", Name: "请假天数", Value: "1天", SortNum: 1},
			{Key: "reason", Name: "请假理由", Value: "家里有事", SortNum: 2},
		},
		Variable: map[string]any{
			"days": 2,
		},
		Users: []goflow_client.NodeUsers{
			{
				NodeId:  "Mjn0M-O5",
				UserIds: []string{"16"},
			},
		},
	})
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", instanceResult)
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
