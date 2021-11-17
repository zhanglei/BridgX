package constants

type Status string

const (
	Undefined Status = "UNDEFINED"
	Pending   Status = "PENDING"
	Timeout   Status = "TIMEOUT"
	Starting  Status = "STARTING"
	Running   Status = "RUNNING"
	Deleted   Status = "DELETED"
	Deleting  Status = "DELETING"
)
