package service

import (
	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/galaxy-future/BridgX/pkg/utils"
)

type DcpService struct {
}

func (s *DcpService) EntryLog(method string, req interface{}) {
	logs.Logger.Info("[service] entry log | func:%v | req:%v", method, utils.ObjToJson(req))
}

func (s *DcpService) ExitLog(method string, req interface{}, resp interface{}, err error) {
	logs.Logger.Info("[service] exit log | func:%v | req:%v | resp:%v | err:%v", method, utils.ObjToJson(req), utils.ObjToJson(resp), err)
}

func (s *DcpService) Connect(req *ConnectRequest) (resp *ConnectResponse) {
	var err error
	s.EntryLog("AccessService", req)
	defer func() {
		s.ExitLog("AccessService", req, resp, err)
	}()
	resp = &ConnectResponse{
		BaseResponse: NewBaseResponse(),
	}

	return resp
}

func (s *DcpService) ScaledUp(req *ScaleUpRequest) (resp *ScaleUpResponse) {
	var err error
	s.EntryLog("ScaledUp", req)
	defer func() {
		s.ExitLog("ScaledUp", req, resp, err)
	}()
	resp = &ScaleUpResponse{
		BaseResponse: NewBaseResponse(),
	}
	resp.Data = &ScaleUpData{
		TaskIds: []int64{},
	}
	return resp
}

func (s *DcpService) ScaleDown(req *ScaleDownRequest) (resp *ScaleDownResponse) {
	var err error
	s.EntryLog("ScaledDown", req)
	defer func() {
		s.ExitLog("ScaledDown", req, resp, err)
	}()
	resp = &ScaleDownResponse{
		BaseResponse: NewBaseResponse(),
	}

	resp.Data = &ScaleDownData{
		TaskIds: []int64{},
	}
	return resp
}

func (s *DcpService) ListNodes(req *ListNodesRequest) (resp *ListNodesResponse) {
	var err error
	s.EntryLog("ListNodes", req)
	defer func() {
		s.ExitLog("ListNodes", req, resp, err)
	}()
	resp = &ListNodesResponse{
		BaseResponse: NewBaseResponse(),
	}
	return resp
}

func (s *DcpService) QueryTask(req *QueryTaskRequest) (resp *QueryTaskResponse) {
	var err error
	s.EntryLog("QueryTask", req)
	defer func() {
		s.ExitLog("QueryTask", req, resp, err)
	}()
	resp = &QueryTaskResponse{
		BaseResponse: NewBaseResponse(),
	}
	return resp
}
