package main

import (
	_ "backend/routers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	logs.SetLevel(logs.LevelEmergency)

}


func main() {
	beego.Run()
}
