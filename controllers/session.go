package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	cfgPath = ""
)

func init() {
	workdir, err := os.Getwd()
	if err != nil {
		beego.BeeLogger.Warn("Can not get the working directory")
		return
	}

	cfgPath = filepath.Join(workdir, "conf/_coda.cfg")
	if _, err = os.Stat(cfgPath); os.IsNotExist(err) {
		beego.BeeLogger.Warn("Can not find %v", cfgPath)
	}
}

// Operations about Protocol
type SessionController struct {
	beego.Controller
}

func (u *SessionController) Put() {
}

func getClients(clientList string) string {
	return strings.Replace(clientList, ",", " ", -1)
}

func copyCfgFile(srcFile, dstFile string) error {
	writer, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer writer.Close()
	reader, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(writer, reader)
	return err
}

func appendUserNameToCfg(userName string) error {
	workdir, err := os.Getwd()
	if err != nil {
		return errors.New("Can not get the working directory")
	}

	tmpFile := filepath.Join(workdir, "coda.cfg")
	_ = os.Remove(tmpFile)

	if err = copyCfgFile(cfgPath, tmpFile); err != nil {
		return err
	}

	file, err := os.OpenFile(tmpFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("Can not open file: " + tmpFile)
	}

	file.Write([]byte("user_name " + userName + "\n"))
	defer file.Close()
	return nil
}

func (u *SessionController) saveSchemaFile() (string, bool) {
	workdir, err := os.Getwd()
	if err != nil {
		workdir = "./"
	}
	_, h, err := u.GetFile("schemaFile")
	if err != nil {
		beego.BeeLogger.Error("%v", err.Error())
		return "", false
	}
	err = u.SaveToFile("schemaFile", filepath.Join(workdir, h.Filename))

	if err != nil {
		beego.BeeLogger.Error("%v", err.Error())
		return "", false
	}
	return h.Filename, true
}

func (u *SessionController) Post() {
	userName := u.GetString("userName")
	sessionName := u.GetString("sessionName")
	clientList := u.GetString("clientList")
	clients := getClients(clientList)
	protocol := u.GetString("protocol")

	if !checkProtocol(protocol) {
		beego.BeeLogger.Warn("Invalid protocol %v", protocol)
		u.Abort("500")
		return
	}

	if userName == "" || sessionName == "" || clients == "" {
		beego.BeeLogger.Warn("Invalid parameters")
		u.Abort("500")
		return
	}

	if err := appendUserNameToCfg(userName); err != nil {
		beego.BeeLogger.Warn("%v", err.Error())
		u.Abort("500")
		return
	}

	schemaFile, ok := u.saveSchemaFile()
	if !ok {
		u.Abort("500")
		return
	}

	cmdline := fmt.Sprintf("./client init %v %v %v %v", sessionName, protocol, schemaFile, clients)
	_, err := exec.Command("./client", "init", sessionName, protocol, schemaFile, clients).Output()
	if err != nil {
		u.Ctx.WriteString(cmdline + "\n")
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Ctx.WriteString("Session created!")
}
