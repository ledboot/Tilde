package main

import (
	"github.com/astaxie/beego"
	"github.com/ledboot/tilde/api"
)

func init() {
	beego.Router("/api/jobs", &api.JobApi{}, "get:List;post:Post")
}
