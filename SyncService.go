package goflow_client

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type SyncService interface {
	/**
	 * 全量同步机构数据
	 */
	SyncOrgAll(orgs []Org) (err error)

	/**
	 * 增量同步机构数据
	 */
	SyncOrgIncr(orgs []Org) (err error)

	/**
	 * 全量同步用户数据
	 */
	SyncUserAll(users []User) (err error)

	/**
	 * 增量同步用户数据
	 */
	SyncUserIncr(users []User) (err error)
}

type SyncServiceImpl struct {
	client *Client
}

func getSyncService(client *Client) SyncService {
	return SyncServiceImpl{
		client: client,
	}
}

func (s SyncServiceImpl) SyncOrgAll(orgs []Org) (err error) {
	var param []map[string]any
	err = mapstructure.Decode(orgs, &param)
	if err != nil {
		return
	}
	result, err := httpPost[any](s.client, "client/org/syncAll", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (s SyncServiceImpl) SyncOrgIncr(orgs []Org) (err error) {
	var param []map[string]any
	err = mapstructure.Decode(orgs, &param)
	if err != nil {
		return
	}
	result, err := httpPost[any](s.client, "client/org/syncIncr", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (s SyncServiceImpl) SyncUserAll(users []User) (err error) {
	var param []map[string]any
	err = mapstructure.Decode(users, &param)
	if err != nil {
		return
	}
	result, err := httpPost[any](s.client, "client/user/syncAll", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (s SyncServiceImpl) SyncUserIncr(users []User) (err error) {
	var param []map[string]any
	err = mapstructure.Decode(users, &param)
	if err != nil {
		return
	}
	result, err := httpPost[any](s.client, "client/user/syncIncr", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

type Org struct {
	Id       string  `json:"id"`       // 机构Id
	Name     string  `json:"name"`     // 机构名称
	Deleted  bool    `json:"deleted"`  // 是否被删除（true-已删除）
	Type     OrgType `json:"type"`     // 机构分类
	Children []Org   `json:"children"` // 下级机构
}

type OrgType struct {
	Id   string `json:"id"`   // 机构分类Id
	Name string `json:"name"` // 机构分类名称
}

type User struct {
	Id      string   `json:"id"`      // 用户Id
	Name    string   `json:"name"`    // 用户名称
	Deleted bool     `json:"deleted"` // 是否被删除（true-已删除）
	OrgIds  []string `json:"orgIds"`  // 所属机构Id
}
