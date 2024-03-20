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
	err := client.GetSyncService().SyncUserIncr([]goflow_client.User{
		{
			Id:      "5",
			Name:    "赵六2",
			Deleted: false,
			OrgIds:  []string{"1"},
		}, {
			Id:      "6",
			Name:    "王八2",
			Deleted: false,
			OrgIds:  []string{"1"},
		},
	})
	fmt.Println("异常：", err)
}
