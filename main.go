package main

import (
	"fmt"
	"time"
	"translateTool/models"

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
	u2 := models.GetEntryById(1)
	fmt.Printf("user2: %v\n", *u2)

	u3 := models.GetEntry("en", "zh_cn", "Apple")
	fmt.Printf("user3: %v\n", u3)

	beego.Run()
}
