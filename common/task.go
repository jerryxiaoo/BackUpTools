package common

// Task 结构体
type Task struct {
	TaskName   string   `json:"taskName"`
	SourcePath string   `json:"sourcePath"`
	TargetPath []string `json:"targetPath"`
}
