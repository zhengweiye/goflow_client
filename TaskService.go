package goflow_client

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type TaskService interface {
	/**
	 * 查询任务详情
	 */
	GetTask(taskId string) (task *TaskVo, err error)

	/**
	 * 查询任务集合
	 */
	GetTasks(instanceId string) (tasks []TaskVo, err error)

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
	 * 强制回退到 "userId 对应的最近一次处理"
	 */
	RollbackLatestApproval(instanceId, userId string) (err error)

	/**
	 * 强制回退到 "指定任务"
	 */
	RollbackTask(taskId string) (err error)

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

	/**
	 * 替换某个流程下，某个任务的负责人-->只使用于或签模式
	 * nodeId和nodeName 二选一
	 */

	ReplaceUsers(processKey, nodeId, nodeName string, userIds []string, instanceIds []string) (err error)

	/**
	 * 根据任务Id获取对应 “流程节点” 的配置信息
	 */
	GetTaskNodeConf(taskId string) (taskNodeConf TaskNodeConfVo, err error)
}

type TaskServiceImpl struct {
	client *Client
}

func getTaskService(client *Client) TaskService {
	return TaskServiceImpl{
		client: client,
	}
}

func (t TaskServiceImpl) GetTask(taskId string) (task *TaskVo, err error) {
	param := map[string]any{
		"taskId": taskId,
	}
	result, err := httpPost[*TaskVo](t.client, "client/task/getTask", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	task = result.Data
	return
}

func (t TaskServiceImpl) GetTasks(instanceId string) (tasks []TaskVo, err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[[]TaskVo](t.client, "client/task/getTasks", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	tasks = result.Data
	return
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

func (t TaskServiceImpl) RollbackLatestApproval(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](t.client, "client/task/rollbackLatestApproval", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) RollbackTask(taskId string) (err error) {
	param := map[string]any{
		"taskId": taskId,
	}
	result, err := httpPost[any](t.client, "client/task/rollbackTask", param)
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

func (t TaskServiceImpl) ReplaceUsers(processKey, nodeId, nodeName string, userIds []string, instanceIds []string) (err error) {
	param := map[string]any{
		"processKey":  processKey,
		"nodeId":      nodeId,
		"nodeName":    nodeName,
		"userIds":     userIds,
		"instanceIds": instanceIds,
	}
	result, err := httpPost[any](t.client, "client/task/replaceUsers", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (t TaskServiceImpl) GetTaskNodeConf(taskId string) (taskNodeConf TaskNodeConfVo, err error) {
	param := map[string]any{
		"taskId": taskId,
	}
	result, err := httpPost[TaskNodeConfVo](t.client, "client/task/getTaskNodeConf", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	taskNodeConf = result.Data
	return
}

type ExecutionRequest struct {
	TaskId            string         `json:"taskId"`            // 任务Id（必填）
	UserId            string         `json:"userId"`            // 用户Id（必填）
	HandleResult      string         `json:"handleResult"`      // 是否通过,继续往下流转（必填）
	HandleOpinion     string         `json:"handleOpinion"`     // 审批意见（选填）
	HandleRemark      string         `json:"handleRemark"`      // 备注
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
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TaskHandle struct {
	UserId           string `json:"userId"`
	UserName         string `json:"userName"`
	TaskId           string `json:"taskId"`
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

type TaskNodeConfVo struct {
	PerformType         string  `json:"performType"`         // 审批模式（orSign-或签（只需一个人审批）；counterSign-会签（每个人都要审批））
	CounterSignSequence string  `json:"counterSignSequence"` // 会签顺序（order-顺序执行,parallel-同时执行）
	PassVoteSymbol      string  `json:"passVoteSymbol"`      // 会签通过符号（eq-等于,gt-大于,gte-大于等于,neq-不等于,lt-小于,lte-小于等于）
	PassVoteRate        float32 `json:"passVoteRate"`        // 会签通过比例（该节点通过人数的比例）
}

/*type TaskVo struct {
	Id           string     `json:"id"`
	TenantId     string     `json:"tenantId"`
	EnvId        string     `json:"envId"`
	InstanceId   string     `json:"instanceId"`
	PrevTaskId   *string    `json:"prevTaskId"`
	Key          string     `json:"key"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	EndTime      *time.Time `json:"endTime"`
	State        string     `json:"state"`
	HandleResult *string    `json:"handleResult"`
	CreateTime   time.Time  `json:"createTime"`
	SortNo       int64      `json:"sortNo"`
}*/

type TaskVo struct {
	Id           string     `json:"id"`
	Key          string     `json:"key"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	EndTime      string     `json:"endTime"`
	State        string     `json:"state"`
	HandleResult string     `json:"handleResult"`
	CreateTime   string     `json:"createTime"`
	Users        []TaskUser `json:"users"`
	NextNodes    []TaskNode `json:"nextNodes"`
}

type TaskNode struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TaskUser struct {
	UserId        string `json:"userId"`
	UserName      string `json:"userName"`
	State         string `json:"state"`
	HandleResult  string `json:"handleResult"`
	HandleOpinion string `json:"handleOpinion"`
	HandleRemark  string `json:"handleRemark"`
	HandleTime    string `json:"handleTime"`
}
