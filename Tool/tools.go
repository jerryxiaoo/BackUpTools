package Tool

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// 利用闭包，因为上一层函数只能接受固定的参数，但我们还需要传其他参数进来
func NewWalkDirFunc(sourcePath string, targetPath string) fs.WalkDirFunc {

	return func(path string, fileInfo fs.DirEntry, err error) error {
		//三个参数都是上一层函数 filepath.WalkDir 这个标准库函数提供的
		//path 扫描到当前文件或文件夹的路径

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
			//设置默认权限
			dirMode := fs.FileMode(0755)
			//创建文件夹
			err := os.MkdirAll(finalTargetPath, dirMode)
			if err != nil {
				return fmt.Errorf("创建目录 %s 失败：%w", finalTargetPath, err)
			}
			fmt.Printf("已创建目录：%s\n", finalTargetPath)
			// 目录处理完成，跳过后续文件逻辑
			return nil

		}

		// 4. 处理文件(不是文件夹的都是文件)
		if !fileInfo.IsDir() {
			//打开源文件
			readFile, err := os.Open(path)
			if err != nil {
				fmt.Println("打开源文件失败", path)
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

			fmt.Printf("备份成功：%s ---> %s\n", sourcePath+""+fileInfo.Name(), finalTargetPath)
			return nil
		}

		if err != nil {
			return fmt.Errorf("备份失败：%w", err)
		}
		return nil
	}
}
