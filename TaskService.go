package goflow_client

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type TaskService interface {
	/**
	 * 获取当前任务对应的处理结果
	 * taskId 任务Id
	 */
	GetTaskResults(taskId string) (results []TaskResult, err error)

	/**
	 * 执行任务
	 */
	Execution(req ExecutionRequest) (nextTasks []Task, err error)

	/**
	 * 获取处理意见
	 */
	GetHandleByUserId(instanceId, userId, nodeKey string) (taskHandles []TaskHandle, err error)
	GetHandleByTaskId(taskId string) (taskHandles []TaskHandle, err error)
	GetHandleByInstanceId(instanceId string) (taskHandles []TaskHandle, err error)

	/**
	 * 抄送已读
	 */
	ReadCc(taskId, userId string) (err error)

	/**
	 * 强制回退
	 */
	ForceRollback(instanceId, userId string) (err error)

	/**
	 * 获取附件列表
	 * instanceId 实例Id
	 */
	GetFiles(instanceId string) (files []TaskFile, err error)

	/**
	 * 上传附件
	 * taskId 任务Id
	 * fileName 文件名称
	 * data 文件字节数组
	 */
	UploadFile(taskId, fileName string, data []byte) (err error)

	/**
	 * 删除附件
	 * fileId 文件Id
	 */
	DeleteFile(fileId string) (err error)

	/**
	 * 加签
	 */
	AddUsers(taskId string, userIds []string) (err error)

	/**
	 * 减签
	 */
	RemoveUsers(taskId string, userIds []string) (err error)
}

type TaskServiceImpl struct {
	client *Client
}

func getTaskService(client *Client) TaskService {
	return TaskServiceImpl{
		client: client,
	}
}

func (t TaskServiceImpl) GetTaskResults(taskId string) (results []TaskResult, err error) {
	param := map[string]any{
		"taskId": taskId,
	}
	result, err := httpPost[[]TaskResult](t.client, "client/task/getTaskResults", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	results = result.Data
	return
}

func (t TaskServiceImpl) Execution(req ExecutionRequest) (nextTasks []Task, err error) {
	var param map[string]any
	err = mapstructure.Decode(req, &param)
	if err != nil {
		return
	}
	result, err := httpPost[[]Task](t.client, "client/task/execute", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	nextTasks = result.Data
	return
}

func (t TaskServiceImpl) GetHandleByUserId(instanceId, userId, nodeKey string) (taskHandles []TaskHandle, err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
		"nodeKey":    nodeKey,
	}
	result, err := httpPost[[]TaskHandle](t.client, "client/task/getHandleByUserId", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	taskHandles = result.Data
	return
}

func (t TaskServiceImpl) GetHandleByTaskId(taskId string) (taskHandles []TaskHandle, err error) {
	param := map[string]any{
		"taskId": taskId,
	}
	result, err := httpPost[[]TaskHandle](t.client, "client/task/getHandleByTaskId", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	taskHandles = result.Data
	return
}

func (t TaskServiceImpl) GetHandleByInstanceId(instanceId string) (taskHandles []TaskHandle, err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[[]TaskHandle](t.client, "client/task/getHandleByInstanceId", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	taskHandles = result.Data
	return
}

func (t TaskServiceImpl) ReadCc(taskId, userId string) (err error) {
	param := map[string]any{
		"taskId": taskId,
		"userId": userId,
	}
	result, err := httpPost[any](t.client, "client/task/read", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) ForceRollback(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](t.client, "client/task/forceRollback", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) GetFiles(instanceId string) (files []TaskFile, err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[[]TaskFile](t.client, "client/task/getFiles", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	files = result.Data
	return
}

func (t TaskServiceImpl) UploadFile(taskId, fileName string, data []byte) (err error) {
	param := map[string]any{
		"taskId":   taskId,
		"fileName": fileName,
		"data":     data,
	}
	result, err := httpPost[any](t.client, "client/task/uploadFile", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) DeleteFile(fileId string) (err error) {
	param := map[string]any{
		"fileId": fileId,
	}
	result, err := httpPost[any](t.client, "client/task/deleteFile", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) AddUsers(taskId string, userIds []string) (err error) {
	param := map[string]any{
		"taskId":  taskId,
		"userIds": userIds,
	}
	result, err := httpPost[any](t.client, "client/task/addUsers", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) RemoveUsers(taskId string, userIds []string) (err error) {
	param := map[string]any{
		"taskId":  taskId,
		"userIds": userIds,
	}
	result, err := httpPost[any](t.client, "client/task/removeUsers", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

type ExecutionRequest struct {
	TaskId            string         `json:"taskId"`            // 任务Id（必填）
	UserId            string         `json:"userId"`            // 用户Id（必填）
	HandleResult      string         `json:"handleResult"`      // 是否通过,继续往下流转（必填）
	HandleOpinion     string         `json:"handleOpinion"`     // 审批意见（选填）
	HandleRemark      string         `json:"handleRemark"`      // 备注
	ForceEnd          bool           `json:"forceEnd"`          // 强制结束流程, 当前人审批通过时(此时流程尚未结束),其可以直接结束流程,而不往下继续
	Variable          map[string]any `json:"variable"`          // 变量（选填）
	NextHandleUserIds []string       `json:"nextHandleUserIds"` // 下一个节点审批人（选填）
	Files             []File         `json:"files"`             // 附件（选填）
	SignFile          File           `json:"signFile"`          // 签名文件（选填）signFile和signFileId二选一
	SignFileId        string         `json:"signFileId"`        // 签名文件Id（选填）signFile和signFileId二选一
}

type File struct {
	Name string `json:"name"` // 文件名称
	Data []byte `json:"data"` // 文件字节数组
}

type TaskFile struct {
	InstanceId string `json:"instanceId"` // 实例Id
	TaskId     string `json:"taskId"`     // 任务Id
	FileId     string `json:"fileId"`     // 文件Id
	FileName   string `json:"fileName"`   // 文件名称
	FilePath   string `json:"filePath"`   // 下载地址
	FileType   string `json:"fileType"`   // 文件后缀
}

type TaskResult struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FlowToNext bool   `json:"flowToNext"`
}

type TaskHandle struct {
	UserId           string `json:"userId"`
	UserName         string `json:"userName"`
	NodeKey          string `json:"nodeKey"`
	NodeName         string `json:"nodeName"`
	HandleResultCode string `json:"handleResultCode"`
	HandleResultName string `json:"handleResultName"`
	HandleOpinion    string `json:"handleOpinion"`
	HandleRemark     string `json:"handleRemark"`
	HandleTime       string `json:"handleTime"`
	SignFileId       string `json:"signFileId"`
	SignFilePath     string `json:"signFilePath"`
}

type Task struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}
