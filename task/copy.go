package task

import (
	"backupTools/common"
	"backupTools/tool"
	"fmt"
	"path/filepath"
)

func CopyToTargetPath(taskList []common.Task) error {
	//开启循环，拿到每一个对象
	for _, task := range taskList {
		//拿到每个对象的多个地址
		for _, singleTargetPath := range task.TargetPath {
			//filepath.WalkDir ：功能是遍历指定目录，并对每个文件和子目录执行给定的函数。(func)
			err := filepath.WalkDir(task.SourcePath, tool.NewWalkDirFunc(task.SourcePath, singleTargetPath))
			if err != nil {
				fmt.Printf(err.Error())
				return err
			}
		}

	}
	return nil
}
