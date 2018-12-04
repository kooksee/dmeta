package cnst

type consumer struct {
	Task         string
	TaskResult   string
	TaskCallback string
	TaskIds      string
	TaskErrCount string
	Log          string
	OssSave      string
}

var Consumer = consumer{
	Task:         "task",
	TaskIds:      "task-ids",
	TaskResult:   "task-result",
	TaskCallback: "task-callback",
	TaskErrCount: "task-err-count",
	Log:          "log",
	OssSave:      "oss_save",
}
