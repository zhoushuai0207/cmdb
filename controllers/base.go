package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"io"
)

type BaseController struct {
	beego.Controller
}

// JsonResult 响应 json 结果
func (c *BaseController) JsonResult(code int, msg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
	c.StopRun()
}
