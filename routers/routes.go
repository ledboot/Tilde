package routers

import (
	"github.com/astaxie/beego"
	"github.com/ledboot/Tilde/api"
)

func init() {
	beego.Router("/api/jobs", &api.JobApi{}, "get:List;post:Post")
}
