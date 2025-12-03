package task

// Task 结构体
type Task struct {
	TaskName   string   `json:"taskName"`
	SourcePath string   `json:"sourcePath"`
	TargetPath []string `json:"targetPath"`
}

func NewTask() Task {
	return Task{}
}

// TaskQueue结构体
type TaskQueue struct {
}

func NewTaskQueue() TaskQueue {

	return TaskQueue{}
}
