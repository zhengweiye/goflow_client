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
		Env:       "test",
	})
	instanceResult, err := client.GetProcessInstanceService().Start(goflow_client.StartRequest{
		StartUserId:   "1",
		ProcessKey:    "leave",
		AutoSubmit:    true,
		BusinessId:    "1",
		BusinessTitle: "郑伟业请假申请",
		Fields: []goflow_client.Field{
			{Key: "days", Name: "请假天数", Value: "1天", SortNum: 1},
			{Key: "reason", Name: "请假理由", Value: "家里有事", SortNum: 2},
		},
	})
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", instanceResult)
}
