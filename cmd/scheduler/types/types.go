package types

import (
	"github.com/tovenja/cron/v3"
)

type Scheduler struct {
	Interval int
	Monitor  cron.Job
}
