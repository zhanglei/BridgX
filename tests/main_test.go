package tests

import (
	"os"
	"testing"

	"github.com/galaxy-future/BridgX/config"
	"github.com/galaxy-future/BridgX/internal/clients"
	"github.com/galaxy-future/BridgX/internal/logs"
)

func TestMain(m *testing.M) {
	//因为是相对路径，需要把conf文件copy到tests目录下
	config.Init()
	logs.Init()
	clients.Init()
	exitCode := m.Run()
	os.Exit(exitCode)
}
