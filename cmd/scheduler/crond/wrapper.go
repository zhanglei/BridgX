package crond

import (
	"sync"
	"time"

	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/tovenja/cron/v3"
)

var cronServer *cron.Cron
var jm sync.Map
var im sync.Map

type XJob interface {
	GetVersionNo() string
	SetVersionNo(v string)
	UniqueKey() string
	cron.Job
}

type cronLogger struct {
}

func (c cronLogger) Info(msg string, keysAndValues ...interface{}) {
	logs.Logger.Infof(msg, keysAndValues)
}

func (c cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logs.Logger.Errorf(msg, keysAndValues)
	logs.Logger.Error(err)
}

func Init() {
	cronServer = cron.New(cron.WithChain(cron.Recover(&cronLogger{})))
}

func AddFixedIntervalSecondsJob(interval int, job cron.Job) {
	cronServer.Schedule(cron.Every(time.Duration(interval)*time.Second), job)
}

func AddFixedIntervalSecondsXJob(interval int, job XJob) {
	id := cronServer.Schedule(cron.Every(time.Duration(interval)*time.Second), job)
	jm.Store(job.UniqueKey(), job)
	im.Store(job.UniqueKey(), id)
}

func RemoveXJob(key string) {
	id, ok := im.Load(key)
	if ok {
		cronServer.Remove(id.(cron.EntryID))
	}
}

func GetXJob(key string) XJob {
	job, ok := jm.Load(key)
	if ok {
		return job.(XJob)
	}
	return nil
}

func Run() {
	cronServer.Run()
}

func Stop() {
	ctx := cronServer.Stop()
	<-ctx.Done()
}
