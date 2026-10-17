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

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

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
	if idx := strings.Index(base64Data, ","); idx != -1 {
		base64Data = base64Data[idx+1:]
	}

	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("base64 解码失败: %v", err)
	}

	fullPath := filepath.Join(directory, fileName)
	err = os.WriteFile(fullPath, decoded, 0644)
	if err != nil {
		return fmt.Errorf("文件写入失败: %v", err)
	}
	return nil
}

// GetEditedFileName 生成带 _edited 后缀的文件名，严格保留原始扩展名
func (a *App) GetEditedFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, ext)
	lowerExt := strings.ToLower(ext)

	if lowerExt == ".jpg" || lowerExt == ".jpeg" {
		return fmt.Sprintf("%s_edited%s", name, lowerExt)
	}
	return fmt.Sprintf("%s_edited.png", name)
}
