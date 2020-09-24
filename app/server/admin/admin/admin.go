package admin

import (
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	layout    string = "admin/layout.html"
	listTpl   string = "admin/admin/list.html"
	createTpl string = "admin/admin/create.html"
	editTpl   string = "admin/admin/edit.html"
	errorTpl  string = "admin/error.html"
)

type Controller struct{}

func (c *Controller) List(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": listTpl})
	}
}



