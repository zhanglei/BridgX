package service

import "time"

type BaseResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
}

func NewBaseResponse() BaseResponse {
	return BaseResponse{}
}

type ConnectRequest struct {
}

type ConnectResponse struct {
	BaseResponse BaseResponse
}

type ScaleUpData struct {
	TaskIds []int64 `json:"task_ids"`
}

type ScaleDownData struct {
	TaskIds []int64 `json:"task_ids"`
}

type ScaleUpRequest struct {
	Id  int64 `json:"id"`
	Num int64 `json:"num"`
}

type ScaleUpResponse struct {
	Data         *ScaleUpData `json:"data"`
	BaseResponse BaseResponse
}

type ListNodesRequest struct {
	Id       int64 `json:"id"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"` //默认10
}

type ListNodesResponse struct {
	Page         int64      `json:"page"`
	PageSize     int64      `json:"page_size"`
	QueryCount   int64      `json:"query_count"`
	Data         []NodeInfo `json:"data"`
	BaseResponse BaseResponse
}

type NodeInfo struct {
	Id       int64  `json:"id"`
	Ip       string `json:"ip"`        //节点 ip 地址
	VmId     string `json:"vm_id"`     //节点实例编号
	Status   int64  `json:"status"`    //节点状态（1 成功/3 失败）
	NodeType string `json:"node_type"` // 节点类型（手动 manual/定时 crontab）
}

type ScaleDownRequest struct {
	Id    int64    `json:"id"`
	Nodes []string `json:"nodes"`
}

type ScaleDownResponse struct {
	Data         *ScaleDownData `json:"data"`
	BaseResponse BaseResponse
}

type QueryTaskRequest struct {
	Id int64 `json:"id"`
}

type QueryTaskResponse struct {
	Data         *TaskInfo `json:"data"`
	BaseResponse BaseResponse
}

type TaskInfo struct {
	Id       int64     `json:"id"`
	PoolName string    `json:"pool_name"`
	Status   int64     `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Stat     []int64   `json:"Stat"`
}

type EcsProcess interface {
	Connect(req *ConnectRequest) *ConnectResponse
	// ScaledUp 扩容
	ScaledUp(req *ScaleUpRequest) *ScaleUpResponse
	// ScaleDown 缩容
	ScaleDown(req *ScaleDownRequest) *ScaleDownResponse
	// ListNodes 查询服务池 IPs
	ListNodes(req *ListNodesRequest) *ListNodesResponse

	// QueryTask 查任务状态
	QueryTask(req *QueryTaskRequest) *QueryTaskResponse
}
