/**
 * @Author Kokutas
 * @Description 欢迎页面
 * @Date 2021/2/16 23:41
 **/
package view

import (
	"github.com/Kokutas/application/lib"
	"github.com/Kokutas/application/themes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func welcome(_ fyne.Window) fyne.CanvasObject {
	// 设置title
	title := canvas.NewText("GB/T 28181-2016客户端工具V2.0", color.RGBA{R: 60, G: 200, B: 0, A: 0})
	title.Alignment = fyne.TextAlignCenter
	//title.TextSize = 20
	title.Refresh()
	logo := canvas.NewImageFromResource(themes.AliCloudLogo)
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		title.TextSize = 16
		logo.SetMinSize(fyne.NewSize(20, 75))
	} else {
		title.TextSize = 22
		logo.SetMinSize(fyne.NewSize(100, 125))
	}
	title.Refresh()
	logo.Refresh()
	return container.NewCenter(
		container.NewVBox(
			title,
			logo,
			widget.NewLabel(""),
			container.NewGridWithRows(1,
				widget.NewHyperlinkWithStyle("官方网站", lib.ParseURL("https://www.baidu.com/"), fyne.TextAlignTrailing, fyne.TextStyle{}),
				widget.NewHyperlinkWithStyle("使用教程", lib.ParseURL("https://www.baidu.com/"), fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewHyperlinkWithStyle("联系我们", lib.ParseURL("https://www.baidu.com/"), fyne.TextAlignLeading, fyne.TextStyle{}),
			),
			// 版权信息
			widget.NewLabelWithStyle("© 2009-2021 xxxx 版权所有", fyne.TextAlignCenter, fyne.TextStyle{}),
			// 增值电信业务经营许可证
			widget.NewLabelWithStyle("增值电信业务经营许可证： xxxx", fyne.TextAlignCenter, fyne.TextStyle{}),
		),
	)

}
