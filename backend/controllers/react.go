package controllers

import "github.com/beego/beego/v2/server/web"

type ReactController struct {
    web.Controller
}

func (c *ReactController) ServeReact() {
    c.TplName = "react/index.html"
}
