package main

import (
	"fmt"
	"github.com/zhengweiye/goflow_client"
	"os"
)

func main() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})

	data, _ := os.ReadFile("test/task/测试.txt")

	instanceResult, err := client.GetTaskService().Execution(goflow_client.ExecutionRequest{
		TaskId:            "a5ea2529-1861-4cbd-9885-91387ec878dc",
		UserId:            "2",
		IsPass:            false,
		Opinion:           "",
		Variable:          nil,
		NextHandleUserIds: nil,
		Files: []goflow_client.File{
			{Name: "测试.txt", Data: data},
		},
	})
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", instanceResult)
}
