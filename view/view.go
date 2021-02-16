/**
 * @Author Kokutas
 * @Description //TODO
 * @Date 2021/2/16 23:31
 **/
package view

import (
	"application/lib"
	"application/themes"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
)
func Start(){
	// 1.创建APP
	application :=app.NewWithID("www.kokutas.gb28181")
	// 2.设置自定义主题
	application.Settings().SetTheme(themes.DarkTheme())
	// 3.根据应用创建一个标题为：GB/T 28181-2016 的GUI窗口对象
	w := application.NewWindow("GB/T 28181-2016")
	// 设置窗口的尺寸
	x, h := lib.ScreenSize()
	sx := float32(x)
	sh := float32(h)
	// 如果当前设备是移动端
	if fyne.CurrentDevice().IsMobile() {
		// 设置窗口的尺寸为最大
		w.Resize(fyne.NewSize(sx, sh))
	} else {
		// 设置窗口的尺寸为40% sy,(1-float32(sx)/float32(sh))*sh
		w.Resize(fyne.NewSize(0.5*sx, 0.3*sh))
	}
	// 设置导航内容
	navs := container.NewAppTabs(
		container.NewTabItemWithIcon("欢迎", theme.HomeIcon(), welcome(w)),
		container.NewTabItemWithIcon("欢迎1", theme.HomeIcon(), welcome(w)),
		container.NewTabItemWithIcon("欢迎2", theme.HomeIcon(), welcome(w)),
		container.NewTabItemWithIcon("欢迎3", theme.HomeIcon(), welcome(w)),
	)
	// 判断设备类型: 移动端
	if fyne.CurrentDevice().IsMobile(){
		// 导航居左
		navs.SetTabLocation(container.TabLocationLeading)
		// 判断设备类型：PC端
	}else{
		// 导航置底
		navs.SetTabLocation(container.TabLocationBottom)
	}
	// 设置窗口内容
	w.SetContent(navs)
	// 设置窗口居中
	w.CenterOnScreen()
	// 启动并展示窗口
	w.ShowAndRun()
}
