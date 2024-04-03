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

	/*err := client.GetSyncService().SyncUserIncr([]goflow_client.User{
		{
			Id:      "1",
			Name:    "张三",
			Deleted: false,
			OrgIds:  []string{"1"},
		}, {
			Id:      "2",
			Name:    "李四",
			Deleted: false,
			OrgIds:  []string{"1"},
		}, {
			Id:      "3",
			Name:    "王五",
			Deleted: false,
			OrgIds:  []string{"1"},
		},
	})*/

	err := client.GetSyncService().SyncOrgIncr([]goflow_client.Org{
		{
			Id:      "1",
			Name:    "钦保科技",
			Deleted: false,
			Type: goflow_client.OrgType{
				Id:   "government",
				Name: "政府单位",
			},
		},
	})
	fmt.Println("异常：", err)
}
