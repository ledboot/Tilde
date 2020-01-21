package api

import (
	"encoding/json"
	"fmt"
	"github.com/ledboot/Tilde/logger"
	"github.com/ledboot/Tilde/utils/hack"
	"net/http"
)

type Result struct {
	Msg     string      `json:"msg"`
	Code    int         `json:"code"`
	Content interface{} `json:"content"`
}

func (b *BaseController) Success(data interface{}) {
	b.Ctx.Output.SetStatus(http.StatusOK)
	b.Data["json"] = Result{Code: http.StatusOK, Content: data}
	b.ServeJSON()
}

func (b *BaseController) AbortBadRequest(msg string) {
	logger.Info("Abort BadRequest error. %s", msg)
	b.CustomAbort(http.StatusBadRequest, hack.String(b.errorResult(http.StatusBadRequest, msg)))
}

func (b *BaseController) AbortInternalServerError(msg string) {
	logger.Error("Abort BadRequest error. %s", msg)
	b.CustomAbort(http.StatusInternalServerError, hack.String(b.errorResult(http.StatusInternalServerError, msg)))
}

func (b *BaseController) errorResult(code int, msg string) []byte {
	r := Result{Msg: msg, Code: code}
	body, err := json.Marshal(r)
	if err != nil {
		logger.Errorf("Json Marshal error. %v", err)
		b.CustomAbort(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return body
}

func (b *BaseController) AbortBadRequestFormat(paramName string) {
	msg := fmt.Sprintf("Invalid param %s !", paramName)
	b.AbortBadRequest(msg)
}
