package api

import (
	"github.com/ledboot/Tilde/models"
	"github.com/ledboot/Tilde/spiders/sy9d"
)

type JobApi struct {
	BaseController
}

func (j *JobApi) List() {
	job := &models.Job{}
	//j.Ctx.Output.JSON(job, true, true)
	sy9d.Run()
	j.Success(job)
}

func (j *JobApi) Post() {

}

func (j *JobApi) Get() {

}
