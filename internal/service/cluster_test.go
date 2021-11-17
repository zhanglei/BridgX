package service

import (
	"testing"

	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/pkg/cloud"
)

func TestCalcUnusedInstancesId(t *testing.T) {
	cloudInstances := []cloud.Instance{
		{
			Id: "1",
		},
		{
			Id: "2",
		},
	}
	bridgeXInstances := []model.Instance{
		{
			InstanceId: "1",
		},
	}
	unusedInstanceIds := calcUnusedInstancesId(cloudInstances, bridgeXInstances)
	if len(unusedInstanceIds) != 1 || unusedInstanceIds[0] != "2" {
		t.Errorf("failed in calc ununsed instance want [1] , got %v", unusedInstanceIds)
	}
}
