package goflow_client

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ProcessInstanceService interface {
	/**
	 * 启动流程
	 */
	Start(request StartRequest) (instanceResult InstanceResult, err error)

	/**
	 * 实例结束，短信通知哪些人
	 */
	UpdateSmsReceiver(instanceId string, userIds []string) (err error)

	/**
	 * 回滚事务，针对启动流程时，放弃启动
	 */
	TxRollback(instanceId, userId string) (err error)

	/**
	 * 提交事务，针对启动流程时，确认提交
	 */
	TxCommit(instanceId, userId string) (err error)

	/**
	 * 撤销实例
	 */
	Cancel(instanceId, userId string) (err error)

	/**
	 * 终止实例
	 */
	Stop(instanceId, reason string) (err error)

	/**
	 * 删除实例
	 */
	Delete(instanceId string) (err error)

	/**
	 * 获取流程类型集合
	 */
	GetTypeList() (list []ProcessType, err error)

	/**
	 * 获取不同 "列表" 下的 "状态"下拉框
	 * listType：all-全部列表，create-我发起列表，doing-待办列表，done-已办列表，cc-抄送列表
	 * processKey： 流程标识
	 * includeCancel：是否包含撤销、草稿状态
	 */
	GetStateList(listType, processKey string, includeCancel bool) (list []State, err error)

	/**
	 * 获取实例详情
	 */
	GetDetail(instanceId string) (instanceDetail InstanceDetail, err error)

	/**
	 * 获取实例流水
	 */
	GetFlows(instanceId string) (flows []InstanceFlow, err error)

	/**
	 * 获取实例节点集合
	 */
	GetNodes(instanceId string) (nodes []InstanceNode, err error)

	/**
	 * 根据实例id集合获取实例
	 */
	GetListByIds(instanceIds []string) (list []InstanceList, err error)

	/**
	 * 全部
	 */
	GetAllList(query InstanceAllQuery) (totalElements, totalPage int64, list []InstanceList, err error)
	GetAllCount(query InstanceAllQuery) (count int, err error)

	/**
	 * 我发起
	 */
	GetCreateList(query InstanceCreateQuery) (totalElements, totalPage int64, list []InstanceList, err error)
	GetCreateCount(query InstanceCreateQuery) (count int, err error)

	/**
	 * 待我审批
	 */
	GetTodoList(query InstanceTodoQuery) (totalElements, totalPage int64, list []InstanceDealList, err error)
	GetTodoCount(query InstanceTodoQuery) (count int, err error)

	/**
	 * 我已审批
	 */
	GetDoneList(query InstanceDoneQuery) (totalElements, totalPage int64, list []InstanceDealList, err error)
	GetDoneCount(query InstanceDoneQuery) (count int, err error)

	/**
	 * 抄送我的
	 */
	GetCcList(query InstanceCcQuery) (totalElements, totalPage int64, list []InstanceDealList, err error)
	GetCcCount(query InstanceCcQuery) (count int, err error)

	/**
	 * 更新实例标题
	 */
	UpdateTitle(instanceId, userId, title string) (err error)

	/**
	 * 删除实例列表展示字段
	 */
	DeleteFields(instanceId, userId string) (err error)

	/**
	 * 新增实例列表展示字段
	 */
	AddFields(instanceId, userId string, fields []Field) (err error)
}

type ProcessInstanceServiceImpl struct {
	client *Client
}

func getProcessInstanceService(client *Client) ProcessInstanceService {
	return ProcessInstanceServiceImpl{
		client: client,
	}
}

func (p ProcessInstanceServiceImpl) Start(request StartRequest) (instanceResult InstanceResult, err error) {
	var param map[string]any
	err = mapstructure.Decode(request, &param)
	if err != nil {
		return
	}
	result, err := httpPost[InstanceResult](p.client, "client/instance/start", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	instanceResult = result.Data
	return
}

/*func (p ProcessInstanceServiceImpl) StartChild(request StartChildRequest) (instanceResult InstanceResult, err error) {
	var param map[string]any
	err = mapstructure.Decode(request, &param)
	if err != nil {
		return
	}
	result, err := httpPost[InstanceResult](p.client, "client/instance/startChildren", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	instanceResult = result.Data
	return
}*/

func (p ProcessInstanceServiceImpl) UpdateSmsReceiver(instanceId string, userIds []string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userIds":    userIds,
	}
	result, err := httpPost[any](p.client, "client/instance/updateSmsReceiver", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) TxRollback(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](p.client, "client/instance/txRollback", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) TxCommit(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](p.client, "client/instance/txCommit", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) Cancel(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](p.client, "client/instance/cancel", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) Stop(instanceId, reason string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"reason":     reason,
	}
	result, err := httpPost[any](p.client, "client/instance/stop", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) Delete(instanceId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[any](p.client, "client/instance/delete", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) GetTypeList() (list []ProcessType, err error) {
	param := map[string]any{}
	result, err := httpPost[[]ProcessType](p.client, "client/instance/getTypeList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	list = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetStateList(listType, processKey string, includeCancel bool) (list []State, err error) {
	param := map[string]any{
		"listType":      listType,
		"processKey":    processKey,
		"includeCancel": includeCancel,
	}

	result, err := httpPost[[]State](p.client, "client/instance/getStateList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	list = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetDetail(instanceId string) (instanceDetail InstanceDetail, err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[InstanceDetail](p.client, "client/instance/getInstance", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	instanceDetail = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetFlows(instanceId string) (flows []InstanceFlow, err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[[]InstanceFlow](p.client, "client/instance/getLogs", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	flows = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetNodes(instanceId string) (nodes []InstanceNode, err error) {
	param := map[string]any{
		"instanceId": instanceId,
	}
	result, err := httpPost[[]InstanceNode](p.client, "client/instance/getNodes", param)
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

func (p ProcessInstanceServiceImpl) GetListByIds(instanceIds []string) (list []InstanceList, err error) {
	param := map[string]any{
		"ids": instanceIds,
	}
	result, err := httpPost[[]InstanceList](p.client, "client/instance/getListByIds", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	list = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetAllList(query InstanceAllQuery) (totalElements, totalPage int64, list []InstanceList, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[PageData[InstanceList]](p.client, "client/instance/getAllList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	totalElements = result.Data.TotalElements
	totalPage = result.Data.TotalPage
	list = result.Data.List
	return
}

func (p ProcessInstanceServiceImpl) GetAllCount(query InstanceAllQuery) (count int, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[int](p.client, "client/instance/getAllCount", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	count = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetCreateList(query InstanceCreateQuery) (totalElements, totalPage int64, list []InstanceList, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[PageData[InstanceList]](p.client, "client/instance/getCreateList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	totalElements = result.Data.TotalElements
	totalPage = result.Data.TotalPage
	list = result.Data.List
	return
}

func (p ProcessInstanceServiceImpl) GetCreateCount(query InstanceCreateQuery) (count int, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[int](p.client, "client/instance/getCreateCount", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	count = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetTodoList(query InstanceTodoQuery) (totalElements, totalPage int64, list []InstanceDealList, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[PageData[InstanceDealList]](p.client, "client/instance/getTodoList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	totalElements = result.Data.TotalElements
	totalPage = result.Data.TotalPage
	list = result.Data.List
	return
}

func (p ProcessInstanceServiceImpl) GetTodoCount(query InstanceTodoQuery) (count int, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[int](p.client, "client/instance/getTodoCount", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	count = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetDoneList(query InstanceDoneQuery) (totalElements, totalPage int64, list []InstanceDealList, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[PageData[InstanceDealList]](p.client, "client/instance/getDoneList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	totalElements = result.Data.TotalElements
	totalPage = result.Data.TotalPage
	list = result.Data.List
	return
}

func (p ProcessInstanceServiceImpl) GetDoneCount(query InstanceDoneQuery) (count int, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[int](p.client, "client/instance/getDoneCount", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	count = result.Data
	return
}

func (p ProcessInstanceServiceImpl) GetCcList(query InstanceCcQuery) (totalElements, totalPage int64, list []InstanceDealList, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[PageData[InstanceDealList]](p.client, "client/instance/getCcList", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	totalElements = result.Data.TotalElements
	totalPage = result.Data.TotalPage
	list = result.Data.List
	return
}

func (p ProcessInstanceServiceImpl) GetCcCount(query InstanceCcQuery) (count int, err error) {
	var param map[string]any
	err = mapstructure.Decode(query, &param)
	if err != nil {
		return
	}
	result, err := httpPost[int](p.client, "client/instance/getCcCount", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
	}
	count = result.Data
	return
}

func (p ProcessInstanceServiceImpl) UpdateTitle(instanceId, userId, title string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
		"title":      title,
	}
	result, err := httpPost[any](p.client, "client/instance/updateTitle", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) DeleteFields(instanceId, userId string) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
	}
	result, err := httpPost[any](p.client, "client/instance/deleteFields", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

func (p ProcessInstanceServiceImpl) AddFields(instanceId, userId string, fields []Field) (err error) {
	param := map[string]any{
		"instanceId": instanceId,
		"userId":     userId,
		"fields":     fields,
	}
	result, err := httpPost[any](p.client, "client/instance/addFields", param)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf(result.Msg)
		return
	}
	return
}

type StartRequest struct {
	StartUserId      string         `json:"startUserId"`      // 发起人Id（必填）
	ProcessKey       string         `json:"processKey"`       // 流程标识（必填）
	AutoSubmit       bool           `json:"autoSubmit"`       // 是否自动第一个发起节点（必填）
	BusinessId       string         `json:"businessId"`       // 业务Id（必填）
	BusinessTitle    string         `json:"businessTitle"`    // 业务标题（必填）
	Fields           []Field        `json:"fields"`           // 列表字段（选填）
	Users            []NodeUsers    `json:"users"`            // 启动时,各个节点对应的负责人（选填）
	NextUserIds      []string       `json:"nextUserIds"`      // 启动时,手工选择第一个节点的审批人
	EndNoticeUserIds []string       `json:"endNoticeUserIds"` // 审批结束时,通知人员Id
	Variable         map[string]any `json:"variable"`         // 变量（选填）
	Remark           string         `json:"remark"`           // 备注（选填）
}

type StartChildRequest struct {
	StartRequest
	TaskId string `json:"taskId"` // 父节点任务Id（必填）
}

type Field struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	SortNum int    `json:"sortNum"`
}

type NodeUsers struct {
	NodeId  string   `json:"nodeId"`
	UserIds []string `json:"userIds"`
}

type ProcessType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type InstanceResult struct {
	Id           string        `json:"id"`           // 实例Id
	State        InstanceState `json:"state"`        // 实例状态（cancel-取消,draft-草稿,doing-正在处理,实例结果-在工作流后台动态配置）
	BusinessId   string        `json:"businessId"`   // 业务Id
	BusinessType string        `json:"businessType"` // 业务类型
}

type InstanceState struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type InstanceFlow struct {
	OperUser    string `json:"operUser"`
	OperTime    string `json:"operTime"`
	OperContent string `json:"operContent"`
}

type InstanceNode struct {
	TaskId      string           `json:"taskId"`
	NodeKey     string           `json:"nodeKey"`
	NodeName    string           `json:"nodeName"`
	NodeState   InstanceState    `json:"nodeState"`
	FinishTime  string           `json:"finishTime"`
	ConsumeTime string           `json:"consumeTime"` // 耗时时长
	Users       []User           `json:"users"`
	Handles     []InstanceHandle `json:"handles"`
}

type User struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	AvatarFile UserAvatar `json:"avatarFile"`
}

type UserAvatar struct {
	FileId   string `json:"fileId"`
	FilePath string `json:"filePath"`
}

type InstanceHandle struct {
	User    User          `json:"user"`
	State   InstanceState `json:"state"`
	Time    string        `json:"time"`
	Opinion string        `json:"opinion"`
	Files   []TaskFile    `json:"files"`
}

type State struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type InstanceListField struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InstanceDetail struct {
	Id             string              `json:"id"`
	BusinessId     string              `json:"businessId"`
	BusinessTitle  string              `json:"businessTitle"`
	BusinessTable  string              `json:"businessTable"`
	TypeCode       string              `json:"typeCode"`
	TypeName       string              `json:"typeName"`
	Fields         []InstanceListField `json:"fields"`
	LimitTime      string              `json:"limitTime"`
	EndTime        string              `json:"endTime"`
	CreateUserId   string              `json:"createUserId"`
	CreateUserName string              `json:"createUserName"`
	CreateTime     string              `json:"createTime"`
	DoingTasks     []DoingTask         `json:"doingTasks"`
	State          InstanceState       `json:"state"`
	StateShow      InstanceState       `json:"stateShow"`
}

type DoingTask struct {
	TaskId   string   `json:"taskId"`
	NodeId   string   `json:"nodeId"`
	NodeName string   `json:"nodeName"`
	FormUrl  string   `json:"formUrl"`
	UserIds  []string `json:"userIds"`
}

type InstanceList struct {
	Id            string              `json:"id"`
	BusinessId    string              `json:"businessId"`
	BusinessTitle string              `json:"businessTitle"`
	TypeCode      string              `json:"typeCode"` // 流程类型
	TypeName      string              `json:"typeName"` // 流程类型
	Fields        []InstanceListField `json:"fields"`
	State         State               `json:"state"`
	CreateUser    string              `json:"createUser"`
	CreateTime    string              `json:"createTime"`
	FinishTime    string              `json:"finishTime"` // 实例完成时间
	LimitTime     string              `json:"limitTime"`
	CurNodeName   string              `json:"curNodeName"`
	ConsumeTime   string              `json:"consumeTime"` // 耗时
	IsOverDate    bool                `json:"isOverDate"`  // 是否逾期
	IsNearDate    bool                `json:"isNearDate"`  // 是否临期
	OverDate      string              `json:"overDate"`
	NearDate      string              `json:"nearDate"`
	UrgeCount     int                 `json:"urgeCount"` // 催办次数
}

type InstanceDealList struct {
	Id            string              `json:"id"`
	BusinessId    string              `json:"businessId"`
	BusinessTitle string              `json:"businessTitle"`
	TypeCode      string              `json:"typeCode"` // 流程类型
	TypeName      string              `json:"typeName"` // 流程类型
	Fields        []InstanceListField `json:"fields"`
	State         State               `json:"state"`
	CreateUser    string              `json:"createUser"`
	CreateTime    string              `json:"createTime"`
	FinishTime    string              `json:"finishTime"` // 实例完成时间
	LimitTime     string              `json:"limitTime"`
	TaskId        string              `json:"taskId"`
	NodeKey       string              `json:"nodeKey"`
	NodeName      string              `json:"nodeName"`
}

type InstanceAllQuery struct {
	CurPage         int      `json:"curPage"`         // 当前页码（必填）
	PageSize        int      `json:"pageSize"`        // 每页记录数（必填）
	ProcessCategory string   `json:"processCategory"` // 流程分类（选填）
	ProcessKey      string   `json:"processKey"`      // 流程类型（选填）
	Title           string   `json:"title"`           // 标题（选填）
	State           string   `json:"state"`           // 状态（选填）
	UserIds         []string `json:"userIds"`         // 查看的用户id数组（必填）
	OverDate        bool     `json:"overDate"`        // 查看逾期
	NearDate        bool     `json:"nearDate"`        // 查看临期
	IncludeCancel   bool     `json:"includeCancel"`   // 是否包含已撤销、草稿的记录
}

type InstanceCreateQuery struct {
	CurPage         int    `json:"curPage"`         // 当前页码（必填）
	PageSize        int    `json:"pageSize"`        // 每页记录数（必填）
	ProcessCategory string `json:"processCategory"` // 流程分类（选填）
	ProcessKey      string `json:"processKey"`      // 流程类型（选填）
	Title           string `json:"title"`           // 标题（选填）
	State           string `json:"state"`           // 状态（选填）
	CurUserId       string `json:"curUserId"`       // 用户id（必填）
	OverDate        bool   `json:"overDate"`        // 查看逾期
	NearDate        bool   `json:"nearDate"`        // 查看临期
}

type InstanceTodoQuery struct {
	CurPage         int      `json:"curPage"`         // 当前页码（必填）
	PageSize        int      `json:"pageSize"`        // 每页记录数（必填）
	ProcessCategory string   `json:"processCategory"` // 流程分类（选填）
	ProcessKey      string   `json:"processKey"`      // 流程类型（选填）
	Title           string   `json:"title"`           // 标题（选填）
	CurUserId       string   `json:"curUserId"`       // 用户id（必填）
	IncludeNodeKeys []string `json:"includeNodeKeys"` // 包含哪些节点,为空则包含所有节点（选填）
	ExcludeNodeKeys []string `json:"excludeNodeKeys"` // 排除哪些节点,为空则不排除任何节点（选填）
}

type InstanceDoneQuery struct {
	CurPage         int      `json:"curPage"`         // 当前页码（必填）
	PageSize        int      `json:"pageSize"`        // 每页记录数（必填）
	ProcessCategory string   `json:"processCategory"` // 流程分类（选填）
	ProcessKey      string   `json:"processKey"`      // 流程类型（选填）
	Title           string   `json:"title"`           // 标题（选填）
	CurUserId       string   `json:"curUserId"`       // 用户id（必填）
	IncludeNodeKeys []string `json:"includeNodeKeys"` // 包含哪些节点,为空则包含所有节点（选填）
	ExcludeNodeKeys []string `json:"excludeNodeKeys"` // 排除哪些节点,为空则不排除任何节点（选填）
}

type InstanceCcQuery struct {
	CurPage         int    `json:"curPage"`         // 当前页码（必填）
	PageSize        int    `json:"pageSize"`        // 每页记录数（必填）
	ProcessCategory string `json:"processCategory"` // 流程分类（选填）
	ProcessKey      string `json:"processKey"`      // 流程类型（选填）
	Title           string `json:"title"`           // 标题（选填）
	CurUserId       string `json:"curUserId"`       // 用户id（必填）
}
