package monitors

import (
	"github.com/galaxy-future/BridgX/internal/constants"
	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/galaxy-future/BridgX/internal/model"
)

//TaskKiller 负责将执行时间超过最大执行时间的任务设置为失败
type TaskKiller struct {
}

func (m TaskKiller) Run() {
	tasks, err := model.GetExpireRunningTask(constants.DefaultTaskMaxRunningDuration)
	if err != nil {
		logs.Logger.Error("failed to get expire running task , err : %v", err)
		return
	}
	if len(tasks) == 0 {
		return
	}

	var taskIds []int64
	for _, task := range tasks {
		taskIds = append(taskIds, task.Id)
	}

	err = model.UpdateTaskStatus(taskIds, constants.TaskStatusFailed)
	if err != nil {
		logs.Logger.Error(err)
	}

}
