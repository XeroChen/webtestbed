package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	RspString := "hello testbed!\r\n\r\n"

	RspString += "\r\n\r\nHeaders:\r\n"
	for k, v := range c.Ctx.Request.Header {
		RspString = RspString + fmt.Sprintf("%s: %s\r\n", k, v[0])
		if k == "Waf-Test-Case" {
			c.Ctx.Output.Header("Waf-Test-Case", v[0])
		}
	}

	RspString += "\r\n\r\nURL Params:\r\n"
	for k, v := range c.Input() {
		RspString = RspString + fmt.Sprintf("%s: %s\r\n", k, v)
	}

	RspString += "\r\n\r\nParams:\r\n"
	for k, v := range c.Ctx.Input.Params() {
		RspString = RspString + fmt.Sprintf("%s: %s\r\n", k, v)
	}

	c.Ctx.Output.Body([]byte(RspString))
	// c.Ctx.WriteString(RspString)

	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}

func (c *MainController) Post() {

	RspString := "hello testbed!\r\n\r\n"

	RspString += "\r\n\r\nHeaders:\r\n"
	for k, v := range c.Ctx.Request.Header {
		RspString = RspString + fmt.Sprintf("%s: %s\r\n", k, v[0])
		if k == "Waf-Test-Case" {
			c.Ctx.Output.Header("Waf-Test-Case", v[0])
		}
	}

	RspString += "\r\n\r\nURL Params:\r\n"
	for k, v := range c.Input() {
		RspString = RspString + fmt.Sprintf("%s: %s\r\n", k, v)
	}

	RspString += "\r\n\r\nParams:\r\n"
	for k, v := range c.Ctx.Input.Params() {
		RspString = RspString + fmt.Sprintf("%s: %s\r\n", k, v)
	}

	RspString += "\r\n\r\nBody:\r\n"
	RspString = RspString + string(c.Ctx.Input.RequestBody)

	c.Ctx.Output.Body([]byte(RspString))

	// c.Ctx.WriteString(RspString)

	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}
