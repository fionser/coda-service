package controllers

import (
	"coda-service/models"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/astaxie/beego"
)

var metaPath string

func init() {
	workdir, err := os.Getwd()
	if err != nil {
		beego.BeeLogger.Warn("Can not get the working directory")
		return
	}
	metaPath = filepath.Join(workdir, "meta")
	if _, err = os.Stat(metaPath); os.IsNotExist(err) {
		err = os.Mkdir(metaPath, os.ModePerm)
		if err != nil {
			beego.BeeLogger.Warn("Can not create directory: %v", metaPath)
			metaPath = "."
		}
	}
}

// Operations about Protocol
type KeyGenController struct {
	beego.Controller
}

func (u *KeyGenController) Put() {
	u.Post()
}

func checkProtocol(protocol string) bool {
	possibleProtocols := models.GetAllProtocols()
	for _, possible := range possibleProtocols {
		if possible.Name == protocol {
			return true
		}
	}
	return false
}

func (u *KeyGenController) Post() {
	protocol := u.Ctx.Input.Param(":protocol")
	if !checkProtocol(protocol) {
		beego.BeeLogger.Warn("Invalid protocol %v", protocol)
		u.Abort("500")
		return
	}

	output, err := exec.Command("./core", "keygen", protocol, metaPath).Output()
	if err != nil {
		beego.BeeLogger.Error("Key generation failed. %v\n%v", err.Error(), string(output))
		return
	}
}
