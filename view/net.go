/**
 * @Author Kokutas
 * @Description //TODO
 * @Date 2021/2/16 23:56
 **/
package view

import (
	"application/handler"
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
)

func network(w fyne.Window)fyne.CanvasObject{
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()
	// 设置窗口内容
	return  container.NewAppTabs(
		container.NewTabItem("内网信息", local(ctx,w)),
		container.NewTabItem("公网信息", public(ctx,w)))
}
func local(ctx context.Context,w fyne.Window)fyne.CanvasObject{
	// 设置内容
	content:=widget.NewMultiLineEntry()
	content.SetPlaceHolder("内网信息查询失败")
	content.Refresh()
	result,err:= handler.LocalNetwork(ctx)
	if err != nil {
		log.Error(err)
		content.SetText("内网信息查询失败")
		return content
	}
	content.SetText(fmt.Sprintf("%v",result.String()))
	content.Refresh()
	return content
}
func public(ctx context.Context,w fyne.Window)fyne.CanvasObject{
	// 设置内容
	content:=widget.NewMultiLineEntry()
	content.SetPlaceHolder("公网信息查询失败")
	content.Refresh()
	result,err:= handler.PublicNetwork(ctx)
	if err != nil {
		log.Error(err)
		content.SetText("公网信息查询失败")
		return content
	}
	content.SetText(fmt.Sprintf("%v",result.String()))
	content.Refresh()
	return content
}