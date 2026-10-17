package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend
var assets embed.FS

func main() {
	// 初始化后端 App
	app := NewApp()

	// 配置窗口参数
	err := wails.Run(&options.App{
		Title:     "离线图片批量处理工具",
		Width:     1366,
		Height:    868,
		MinWidth:  1024,
		MinHeight: 700,
		AssetServer: &assetserver.Options{
			Assets: assets, // 将 frontend 文件夹嵌入到 exe 中
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app, // 将 app.go 中的方法暴露给前端 JS 调用
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			ZoomFactor:           1.0,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
