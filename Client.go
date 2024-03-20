package goflow_client

import "sync"

type Client struct {
	host      string
	appId     string
	appSecret string
	env       string
}

var clientObj *Client
var clientOnce sync.Once

type Option struct {
	Host      string
	AppId     string
	AppSecret string
	Env       string
}

func Create(opt Option) *Client {
	clientOnce.Do(func() {
		clientObj = &Client{
			host:      opt.Host,
			appId:     opt.AppId,
			appSecret: opt.AppSecret,
			env:       opt.Env,
		}
		clientObj.sync()
	})
	return clientObj
}

func (c *Client) GetSyncService() SyncService {
	return getSyncService(c)
}

func (c *Client) GetProcessDefineService() ProcessDefineService {
	return getProcessDefineService(c)
}

func (c *Client) GetProcessInstanceService() ProcessInstanceService {
	return getProcessInstanceService(c)
}

func (c *Client) GetTaskService() TaskService {
	return getTaskService(c)
}

func (c *Client) SetOrgAllHook() {

}

func (c *Client) SetOrgIncrHook() {

}

func (c *Client) SetUserAllHook() {

}

func (c *Client) SetUserIncrHook() {

}

func (c *Client) sync() {
	// 模式一：业务系统定时调用SyncService接口进行数据同步
	// 模式二：业务系统创建Client时，设置钩子函数，Client定时执行构造函数
}
