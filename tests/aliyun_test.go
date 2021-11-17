package tests

import (
	"testing"

	"github.com/galaxy-future/BridgX/pkg/cloud/aliyun"
)

func TestGetAliyunClient(t *testing.T) {
	c, err := aliyun.New("a", "b", "cn-beijing")
	t.Logf("err:%v\n", err)
	region, err := c.GetRegions()

	t.Logf("err:%v\n", err)
	t.Logf("regions:%v\n", region)

}
