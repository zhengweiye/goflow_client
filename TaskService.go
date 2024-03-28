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
	Execution(req ExecutionRequest) (instanceResult InstanceResult, err error)

	/**
	 * 回退权限
	 */
	RollbackLimit(instanceId, userId string) (allow bool, err error)

	/**
	 * 回退
	 */
	Rollback(instanceId, userId string) (err error)

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

func (t TaskServiceImpl) Execution(req ExecutionRequest) (instanceResult InstanceResult, err error) {
	var param map[string]any
	err = mapstructure.Decode(req, &param)
	if err != nil {
		return
	}
	result, err := httpPost[InstanceResult](t.client, "client/task/execute", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	instanceResult = result.Data
	return
}

func (t TaskServiceImpl) RollbackLimit(instanceId, userId string) (allow bool, err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[bool](t.client, "client/task/rollbackLimit", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	allow = result.Data
	return
}

func (t TaskServiceImpl) Rollback(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](t.client, "client/task/rollback", param)
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
	Variable          map[string]any `json:"variable"`          // 变量（选填）
	NextHandleUserIds []string       `json:"nextHandleUserIds"` // 下一个节点审批人（选填）
	Files             []File         `json:"files"`             // 附件（选填）
}

type File struct {
	Name string `json:"name"` // 文件名称
	Data []byte `json:"data"` // 文件字节数组
}

type InstanceResult struct {
	InstanceId    string `json:"instanceId"`    // 实例Id
	InstanceState string `json:"instanceState"` // 实例状态（cancel-取消,draft-草稿,doing-正在处理,reject-驳回,pass-同意）
	BusinessId    string `json:"businessId"`    // 业务Id
	BusinessType  string `json:"businessType"`  // 业务类型
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
	Id   string `json:"id"`
	Name string `json:"name"`
}
