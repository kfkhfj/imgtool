package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App 结构体
type App struct {
	ctx context.Context
}

// NewApp 创建一个新的 App 实例
func NewApp() *App {
	return &App{}
}

// startup 在应用启动时调用
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// SelectDirectory 弹出 Windows 原生文件夹选择对话框
func (a *App) SelectDirectory() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "请选择要保存图片的目录",
	})
	return dir, err
}

// SaveFileToDisk 接收 Base64 编码的图片数据，解码并写入本地磁盘
func (a *App) SaveFileToDisk(directory string, fileName string, base64Data string) error {
	// 去除 DataURL 前缀 (例如 "data:image/png;base64,")
	if idx := strings.Index(base64Data, ","); idx != -1 {
		base64Data = base64Data[idx+1:]
	}

	// Base64 解码为字节数组
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("base64 解码失败: %v", err)
	}

	// 拼接完整的文件保存路径
	fullPath := filepath.Join(directory, fileName)

	// 写入文件，权限设为 0644
	err = os.WriteFile(fullPath, decoded, 0644)
	if err != nil {
		return fmt.Errorf("文件写入失败: %v", err)
	}

	return nil
}

// GetEditedFileName 生成带 _edited 后缀的文件名
func (a *App) GetEditedFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, ext)
	return fmt.Sprintf("%s_edited.png", name)
}
