package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CopyFile 复制文件
func CopyFile(sourceFile, targetFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
		}
	}(source)

	target, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer func(target *os.File) {
		err := target.Close()
		if err != nil {
		}
	}(target)

	_, err = io.Copy(target, source)
	if err != nil {
		return err
	}
	return nil
}

// FileExist 判断文件是否存在
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

// DirExist 判断目录是否存在
func DirExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else {
		// 检查是否是目录
		info, _ := os.Stat(dirPath)
		if info.IsDir() {
			return true
		} else {
			return false
		}
	}
}

// CreateFile 创建文件
func CreateFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	return nil
}

// MkdirAll 递归创建目录
func MkdirAll(dirPath string) bool {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return false
	} else {
		return true
	}
}

// FindByFileName 查找目录下匹配文件名的文件
func FindByFileName(dir, filename string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), filename) {
			files = append(files, info.Name())
		}
		return nil
	})
	return files
}

// MoveFile 移动文件
func MoveFile(src, dst string) error {
	// 先复制
	if err := CopyFile(src, dst); err != nil {
		return err
	}
	// 再删除
	if err := os.Remove(src); err != nil {
		return err
	}
	return nil
}

// CopyFileWithTempAndRename 复制文件，先去掉后缀名作为临时文件，复制完成后再重命名
// 此方法会先检查临时文件是否存在，如果存在则说明正在复制，返回错误
func CopyFileWithTempAndRename(sourceFile, targetFile string) error {
	// 获取目标文件的目录和文件名
	targetDir := filepath.Dir(targetFile)
	targetName := filepath.Base(targetFile)

	// 去掉后缀名作为临时文件名
	ext := filepath.Ext(targetName)
	nameWithoutExt := strings.TrimSuffix(targetName, ext)
	tempFile := filepath.Join(targetDir, nameWithoutExt)

	// 检查临时文件是否存在，如果存在说明正在复制
	if FileExist(tempFile) {
		return os.ErrExist // 返回文件已存在错误，表示正在复制中
	}

	// 确保目标目录存在
	if !DirExist(targetDir) {
		if !MkdirAll(targetDir) {
			return os.ErrPermission
		}
	}

	// 先复制到临时文件（无后缀名）
	if err := CopyFile(sourceFile, tempFile); err != nil {
		return err
	}

	// 复制成功后，将临时文件重命名为目标文件
	if err := os.Rename(tempFile, targetFile); err != nil {
		// 如果重命名失败，清理临时文件
		os.Remove(tempFile)
		return err
	}

	return nil
}
