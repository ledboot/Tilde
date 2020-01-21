package api

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/ledboot/Tilde/logger"
	"github.com/ledboot/Tilde/models"
	"strings"
)

const (
	DefaultPageNo   = 1
	DefaultPageSize = 10
)

type BaseController struct {
	BaseApi
}

type BaseApi struct {
	beego.Controller
}

func (b *BaseApi) DecodeJsonReq(v interface{}) error {
	err := json.Unmarshal(b.Ctx.Input.CopyBody(1<<32), v)
	if err != nil {
		logger.Errorf("Error while decoding the json request, error: %v, %v",
			err, string(b.Ctx.Input.CopyBody(1 << 32)[:]))
		return errors.New("Invalid json request")
	}
	return nil
}

func (b *BaseController) GetQueryParams() models.QueryParams {
	pageNo, _ := b.GetInt("pageNo", DefaultPageNo)
	pageSize, _ := b.GetInt("pageSize", DefaultPageSize)
	sortby := b.Input().Get("sortby")
	filter := b.Input().Get("filter")
	qmap := map[string]interface{}{}
	if filter != "" {
		filters := strings.Split(filter, ",")
		for _, param := range filters {
			params := strings.Split(param, "=")
			if len(params) != 2 {
				continue
			}
			qmap[params[0]] = params[1]
		}
	}
	return models.QueryParams{PageSize: pageSize, PageNo: pageNo, SortBy: sortby, Query: qmap}
}
