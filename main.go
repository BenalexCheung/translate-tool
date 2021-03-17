package main

import (
	"time"

	_ "translateTool/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

}

func main() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	// orm.RegisterModelWithPrefix("prefix_", new(User))

	beego.Run()
}
