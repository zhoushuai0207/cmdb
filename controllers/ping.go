package controllers

type PingController struct {
	BaseController
}

func (c *PingController) URLMapping() {
	c.Mapping("Get", c.CheckHealth)
}

func (c *PingController) CheckHealth() {
	//c.Ctx.WriteString("OK")
	c.JsonResult(0, "ok")
}
