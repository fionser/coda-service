package controllers

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

var pkPath string

func checkPubKe() {
	workdir, err := os.Getwd()
	if err != nil {
		beego.BeeLogger.Warn("Can not get the working directory")
		return
	}
	pkPath = filepath.Join(workdir, "meta", "pk.pk")
	if _, err = os.Stat(pkPath); os.IsNotExist(err) {
		beego.BeeLogger.Warn("Can not find the public key %v", pkPath)
		pkPath = ""
		return
	}
}

// Operations about Protocol
type EncryptionController struct {
	beego.Controller
}

func (u *EncryptionController) Put() {
	u.Post()
}

func (u *EncryptionController) Post() {
	srcFile := u.Ctx.Input.Param(":src")
	dstDir := u.Ctx.Input.Param(":dst")
	beego.BeeLogger.Info("%v -> %v", srcFile, dstDir)
}
