package controllers

import (
	"strconv"

	"github.com/beego/beego/v2/server/web"

	"fmt"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Get() {
	userId := c.Ctx.Input.Param(":id")
	parsedId, err := strconv.ParseInt(userId, 10, 64)

	println(parsedId)

	if err != nil || parsedId < 0 {
        c.Ctx.Output.SetStatus(400)
        c.Ctx.Output.Body([]byte("Invalid user ID"))
        return
    }


	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body([]byte(fmt.Sprintf("Hello, Human, your id %s!", userId)))

}
