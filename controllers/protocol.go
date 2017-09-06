package controllers

import (
	"coda-service/models"

	"github.com/astaxie/beego"
)

// Operations about Protocol
type ProtocolController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Protocols
// @Success 200 {object} models.Protocol
// @router / [get]
func (u *ProtocolController) GetAll() {
	protocols := models.GetAllProtocols()
	u.Data["json"] = protocols
	u.ServeJSON()
}
