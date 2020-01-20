package api

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/ledboot/tilde/logger"
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
