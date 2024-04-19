package main

import (
	"fmt"
	"github.com/zhengweiye/goflow_client"
)

func main() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "https://zhifa.cftzqinzhou.com",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	//nodes, err := client.GetProcessDefineService().GetNodes("punish_scene_decide", "16", nil)
	nodes, err := client.GetProcessDefineService().GetNodes("case_main", "16", nil)
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", nodes)
}

func main2() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "https://zhifa.cftzqinzhou.com",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	nodes, err := client.GetProcessInstanceService().GetNodes("a43f7908-5962-4f5a-90aa-40c1be17183d")
	fmt.Println("异常：", err)
	for _, n := range nodes {
		fmt.Println(n.NodeName, "==", n.NodeState)
	}
}
