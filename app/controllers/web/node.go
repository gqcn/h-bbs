package web

import (
	"bbs/app/constants"
	response "bbs/app/funcs/response"
	"bbs/app/model/nodes"
	postsModel "bbs/app/model/posts"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	indexTpl string = "web/node/index.html"
)

// Controller Base
type NodeController struct{}

func (c *NodeController) Index(r *ghttp.Request) {
	id := r.GetRouterVar("nodeId").Int()
	pageNum := r.GetQueryInt("page", 1)
	if id == 0 {
		response.RedirectBackWithError(r, gerror.New("节点未找到"))
	}

	item, err := g.DB().Table(nodes.Table).Where("id = ? and is_delete = ?", id, 0).One()
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if item.IsEmpty() {
		response.RedirectBackWithError(r, gerror.New("节点未找到"))
	}

	items, _ := g.DB().Table(postsModel.Table+" p").
		Fields("p.id,p.title,p.uid,p.nid,p.view_num,p.comment_num,p.create_at,u.name,u.avatar,n.name as node_name").
		InnerJoin("users u", "u.id = p.uid").
		InnerJoin("nodes n", "n.id = p.nid").
		Where("p.nid = ?", id).
		Order("create_at DESC").
		Page(pageNum, 40).
		All()

	response.ViewExit(r, constants.WebLayoutTplPath, g.Map{"mainTpl": indexTpl, "node": item, "posts": items})
}
