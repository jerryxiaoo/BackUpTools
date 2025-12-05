package task

import (
	"backupTools/common"
	"backupTools/config"
	"backupTools/tool"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CopyToTargetPath(taskList []common.Task) error {

	//开启循环，拿到每一个对象
	for _, task := range taskList {
		//拿到每个对象的多个地址
		for _, singleTargetPath := range task.TargetPath {
			//filepath.WalkDir ：功能是遍历指定目录，并对每个文件和子目录执行给定的函数。(func)
			err := filepath.WalkDir(task.SourcePath, NewWalkDirFunc(task.SourcePath, singleTargetPath))
			if err != nil {
				fmt.Printf(err.Error())
				return err
			}
		}

	}
	return nil
}

// 利用闭包，因为上一层函数只能接受固定的参数，但我们还需要传其他参数进来
// 拷贝的逻辑处理
func NewWalkDirFunc(sourcePath, targetPath string) fs.WalkDirFunc {

	return func(path string, fileInfo fs.DirEntry, err error) error {

		//先拿记录日志的路径
		configPath := config.GetLogConfigPath()

		//三个参数都是上一层函数 filepath.WalkDir 这个标准库函数提供的
		//path 扫描到当前文件或文件夹的路径

		//拿到config

		// 1. 处理遍历错误（如权限不足）
		if err != nil {
			fmt.Printf("遍历路径 %s 出错：%v\n", path, err)
			return nil // 返回nil继续遍历，不终止
		}

		// 2.处理隐藏目录 以.开头
		if fileInfo.IsDir() && strings.HasPrefix(fileInfo.Name(), ".") {
			// 跳过该目录的所有子项
			return filepath.SkipDir
		}

		//计算当前文件的相对=对于参数的相对路径，将目标地址拼接在这个相对路径前面，就是目标路径的真实绝对路径
		/*
			假如 "sourcePath": "E:\\Document"  targetPath": “D:\\”
			这时候遍历到 E:\\Document\\1\\2\\3.txt
			就需要先计算出 1\\2\\3.txt
			再把 D:\\ 拼上去，得出最终的地址 D:\\1\\2\\3.txt
		*/
		//计算当前path相对于 sourceDir 的相对路径
		rel, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return err
		}
		//fmt.Println(rel, targetPath)
		//拼接
		finalTargetPath := filepath.Join(targetPath, rel)

		// 3. 处理文件夹
		if fileInfo.IsDir() {
			_, err := os.Stat(finalTargetPath)
			if err == nil {
				msg := fmt.Sprintf("文件夹%s已经存在，跳过当前文件夹", fileInfo.Name())
				tool.Info(msg)
				tool.WriteLogFile(msg, configPath)
				return nil
			}
			//设置文件夹的权限，从源文件夹去取
			dirInfo, _ := fileInfo.Info()
			dirMode := dirInfo.Mode().Perm()
			fmt.Println(dirMode)
			//创建文件夹
			err = os.MkdirAll(finalTargetPath, dirMode)
			if err != nil {
				return fmt.Errorf("创建目录 %s 失败：%w", finalTargetPath, err)
			}
			msg := fmt.Sprintf("已创建目录:%s", finalTargetPath)
			tool.Info(msg)
			tool.WriteLogFile(msg, configPath)
			//fmt.Printf("已创建目录：%s\n", finalTargetPath)
			// 目录处理完成，跳过后续文件逻辑
			return nil

		}

		// 4. 处理文件(不是文件夹的都是文件)
		if !fileInfo.IsDir() {

			//先判断是否存在同名文件
			TargetPathFileInfo, err := os.Stat(finalTargetPath)
			// 表示文件存在，
			if err == nil {
				//先判断是否同样的修改时间
				targetFileModeTime := TargetPathFileInfo.ModTime()
				if err != nil {
					return err
				}
				sourcePathFile, _ := fileInfo.Info()
				sourceFileModeTime := sourcePathFile.ModTime()
				if targetFileModeTime.Before(sourceFileModeTime) {
					err := CopyToTargetPathDetail(path, finalTargetPath, fileInfo)
					if err != nil {
						tool.Error(err.Error())
					}

				} else {
					msg := fmt.Sprintf("文件%s已经存在且修改时间一样，跳过当前文件", fileInfo.Name())
					tool.Info(msg)
					tool.WriteLogFile(msg, configPath)
					return nil
				}

			} else {

				err := CopyToTargetPathDetail(path, finalTargetPath, fileInfo)
				if err != nil {
					tool.Error(err.Error())
				}
				return nil
			}
		}

		if err != nil {
			msg := fmt.Sprintf("备份失败：%w", err)
			tool.Warning(msg)
			return fmt.Errorf(err.Error())
		}
		return nil
	}
}

// 拷贝的操作处理
func CopyToTargetPathDetail(sourcePath, finalTargetPath string, fileInfo fs.DirEntry) error {
	//先拿记录日志的路径
	configPath := config.GetLogConfigPath()

	//打开源文件
	readFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("打开源文件失败", sourcePath)
		return err
	}
	defer readFile.Close()
	//可以拷贝一个权限
	readFileInfo, err := fileInfo.Info()
	if err != nil {
		return err
	}
	writerFile, err := os.OpenFile(finalTargetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("创建写入文件失败", finalTargetPath)
		return err
	}
	defer writerFile.Close()

	_, err = io.Copy(writerFile, readFile)
	if err != nil {
		fmt.Println("拷贝读取流到写入流失败")
		return err
	}
	// 保留文件修改时间
	if err := os.Chtimes(finalTargetPath, readFileInfo.ModTime(), readFileInfo.ModTime()); err != nil {
		return fmt.Errorf("更新文件时间 %s 失败：%w", finalTargetPath, err)
	}

	msg := fmt.Sprintf("备份成功：%s ---> %s", sourcePath+""+fileInfo.Name(), finalTargetPath)
	tool.Info(msg)
	tool.WriteLogFile(msg, configPath)

	return nil
}
