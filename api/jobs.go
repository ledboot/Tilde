package api

import "github.com/ledboot/tilde/models"

type JobApi struct {
	BaseController
}

func (j *JobApi) List() {
	job := &models.Job{}
	j.Ctx.Output.JSON(job, true, true)
}

func (j *JobApi) Post() {

}

func (j *JobApi) Get() {

}
