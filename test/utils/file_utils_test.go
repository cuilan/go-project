package utils_test

import (
	"go-project/internal/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFileWithTempAndRename(t *testing.T) {
	// 创建临时目录用于测试
	tempDir, err := os.MkdirTemp("", "test_copy_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// 创建源文件
	sourceFile := filepath.Join(tempDir, "source.txt")
	sourceContent := "这是测试内容"
	err = os.WriteFile(sourceFile, []byte(sourceContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 定义目标文件
	targetFile := filepath.Join(tempDir, "target.txt")
	tempFile := filepath.Join(tempDir, "target") // 无后缀名的临时文件

	t.Run("正常复制文件", func(t *testing.T) {
		// 执行复制操作
		err := utils.CopyFileWithTempAndRename(sourceFile, targetFile)
		if err != nil {
			t.Fatalf("复制文件失败: %v", err)
		}

		// 验证目标文件存在且内容正确
		if !utils.FileExist(targetFile) {
			t.Error("目标文件不存在")
		}

		content, err := os.ReadFile(targetFile)
		if err != nil {
			t.Fatalf("读取目标文件失败: %v", err)
		}

		if string(content) != sourceContent {
			t.Errorf("文件内容不匹配，期望: %s, 实际: %s", sourceContent, string(content))
		}

		// 验证临时文件已被清理
		if utils.FileExist(tempFile) {
			t.Error("临时文件没有被清理")
		}

		// 清理目标文件
		os.Remove(targetFile)
	})

	t.Run("临时文件存在时返回错误", func(t *testing.T) {
		// 先创建临时文件
		err := os.WriteFile(tempFile, []byte("临时文件"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 尝试复制，应该返回错误
		err = utils.CopyFileWithTempAndRename(sourceFile, targetFile)
		if err != os.ErrExist {
			t.Errorf("期望返回 os.ErrExist 错误，实际: %v", err)
		}

		// 目标文件不应该被创建
		if utils.FileExist(targetFile) {
			t.Error("目标文件不应该被创建")
		}

		// 清理临时文件
		os.Remove(tempFile)
	})

	t.Run("源文件不存在时返回错误", func(t *testing.T) {
		nonExistentFile := filepath.Join(tempDir, "non_existent.txt")
		err := utils.CopyFileWithTempAndRename(nonExistentFile, targetFile)
		if err == nil {
			t.Error("期望返回错误，但成功了")
		}

		// 目标文件和临时文件都不应该存在
		if utils.FileExist(targetFile) {
			t.Error("目标文件不应该被创建")
		}
		if utils.FileExist(tempFile) {
			t.Error("临时文件不应该被创建")
		}
	})

	t.Run("目标目录不存在时自动创建", func(t *testing.T) {
		subDir := filepath.Join(tempDir, "subdir")
		targetFileInSubDir := filepath.Join(subDir, "target.txt")

		err := utils.CopyFileWithTempAndRename(sourceFile, targetFileInSubDir)
		if err != nil {
			t.Fatalf("复制文件到子目录失败: %v", err)
		}

		// 验证子目录被创建
		if !utils.DirExist(subDir) {
			t.Error("子目录没有被创建")
		}

		// 验证文件存在且内容正确
		if !utils.FileExist(targetFileInSubDir) {
			t.Error("目标文件不存在")
		}

		content, err := os.ReadFile(targetFileInSubDir)
		if err != nil {
			t.Fatalf("读取目标文件失败: %v", err)
		}

		if string(content) != sourceContent {
			t.Errorf("文件内容不匹配，期望: %s, 实际: %s", sourceContent, string(content))
		}
	})
}
