package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error400() {
	c.Ctx.ResponseWriter.WriteHeader(400)
	//c.CustomAbort(500, "")
	c.JsonResult(400, " Bad Request")
}

func (c *ErrorController) Error403() {
	c.Ctx.ResponseWriter.WriteHeader(403)
	c.JsonResult(403, " Forbidden")
}
func (c *ErrorController) Error404() {
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.JsonResult(404, "Page Not Find")
}

func (c *ErrorController) Error405() {
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.JsonResult(405, "Method Not Allowed")
}

func (c *ErrorController) Error500() {
	c.Ctx.ResponseWriter.WriteHeader(500)
	c.JsonResult(500, "Server Error")
}
