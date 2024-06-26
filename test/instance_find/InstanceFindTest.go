package main

import (
	"fmt"
	"github.com/zhengweiye/goflow_client"
)

func main10() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	totalElements, totalPage, list, err := client.GetProcessInstanceService().GetCcList(goflow_client.InstanceCcQuery{
		CurPage:   1,
		PageSize:  10,
		CurUserId: "1",
	})
	fmt.Println("异常：", err)
	fmt.Println("总记录数：", totalElements, ", 总页数：", totalPage)
	fmt.Printf("记录：%+v\n", list)
}

func main9() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	totalElements, totalPage, list, err := client.GetProcessInstanceService().GetDoneList(goflow_client.InstanceDoneQuery{
		CurPage:   1,
		PageSize:  10,
		CurUserId: "1",
	})
	fmt.Println("异常：", err)
	fmt.Println("总记录数：", totalElements, ", 总页数：", totalPage)
	fmt.Printf("记录：%+v\n", list)
}

func main8() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	totalElements, totalPage, list, err := client.GetProcessInstanceService().GetTodoList(goflow_client.InstanceTodoQuery{
		CurPage:   1,
		PageSize:  10,
		CurUserId: "2",
	})
	fmt.Println("异常：", err)
	fmt.Println("总记录数：", totalElements, ", 总页数：", totalPage)
	fmt.Printf("记录：%+v\n", list)
	for _, i := range list {
		fmt.Println(i.TaskId)
	}
}

func main7() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	totalElements, totalPage, list, err := client.GetProcessInstanceService().GetCreateList(goflow_client.InstanceCreateQuery{
		CurPage:   1,
		PageSize:  10,
		CurUserId: "3",
	})
	fmt.Println("异常：", err)
	fmt.Println("总记录数：", totalElements, ", 总页数：", totalPage)
	fmt.Printf("记录：%+v\n", list)
}

func main6() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	totalElements, totalPage, list, err := client.GetProcessInstanceService().GetAllList(goflow_client.InstanceAllQuery{
		CurPage:  1,
		PageSize: 10,
		UserIds:  []string{"1"},
	})
	fmt.Println("异常：", err)
	fmt.Println("总记录数：", totalElements, ", 总页数：", totalPage)
	fmt.Printf("记录：%+v\n", list)
}

func main3() {
	client := goflow_client.Create(goflow_client.Option{
		Host:      "http://127.0.0.1:8888",
		AppId:     "enforce",
		AppSecret: "enforce",
		Env:       "test",
	})
	list, err := client.GetProcessInstanceService().GetTypeList()
	fmt.Println("异常：", err)
	fmt.Printf("结果：%+v\n", list)
}
