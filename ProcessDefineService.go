package goflow_client

import (
	"fmt"
)

type ProcessDefineService interface {
	/**
	 * 启动实例之前，查询流程节点
	 * processKey 流程标识
	 * userId 用户Id
	 * processVar 变量（如果流程有分支,根据变量来决定走哪条分支）
	 */
	GetNodes(processKey, userId string, processVar map[string]any) (nodes []Node, err error)

	/**
	 * 启动实例之前，查询第一个审批节点
	 * processKey 流程标识
	 * userId 用户Id
	 * processVar 变量（如果流程有分支,根据变量来决定走哪条分支）
	 */
	GetFirstCheckNode(processKey, userId string, processVar map[string]any) (node Node, err error)

	/**
	 * 审批时，根据当前任务Id查询下一个节点
	 */
	GetNextNodes(taskId string, processVar map[string]any) (nodes []Node, err error)
}

type ProcessDefineServiceImpl struct {
	client *Client
}

func getProcessDefineService(client *Client) ProcessDefineService {
	return ProcessDefineServiceImpl{
		client: client,
	}
}

func (p ProcessDefineServiceImpl) GetNodes(processKey, userId string, processVar map[string]any) (nodes []Node, err error) {
	param := map[string]any{
		"processKey": processKey,
		"userId":     userId,
		"processVar": processVar,
	}
	result, err := httpPost[[]Node](p.client, "client/define/getNodes", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	nodes = result.Data
	return
}

func (p ProcessDefineServiceImpl) GetFirstCheckNode(processKey, userId string, processVar map[string]any) (node Node, err error) {
	param := map[string]any{
		"processKey": processKey,
		"userId":     userId,
		"processVar": processVar,
	}
	result, err := httpPost[Node](p.client, "client/define/getFirstCheckNode", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	node = result.Data
	return
}

func (p ProcessDefineServiceImpl) GetNextNodes(taskId string, processVar map[string]any) (nodes []Node, err error) {
	param := map[string]any{
		"taskId":     taskId,
		"processVar": processVar,
	}
	result, err := httpPost[[]Node](p.client, "client/define/getNextNodes", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	nodes = result.Data
	return
}

type Node struct {
	NodeId      string     `json:"nodeId"`      // 节点Id
	NodeName    string     `json:"nodeName"`    // 节点名称
	AllowSelect bool       `json:"allowSelect"` // 是否允许选择人员
	Users       []NodeUser `json:"users"`       // 节点负责人
}

type NodeUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
