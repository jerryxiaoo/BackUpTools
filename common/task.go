package common

// Task 结构体
type Task struct {
	TaskName       string   `json:"taskName"`
	TaskStatus     string   `json:"taskStatus"`
	LastTimeBackup string   `json:"lastTimeBakup"`
	SourcePath     string   `json:"sourcePath"`
	TargetPath     []string `json:"targetPath"`
}
