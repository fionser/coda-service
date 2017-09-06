package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var (
	// Protocols Array of protocols that run on CODA
	Protocols []*Protocol
)

// Protocol Defining the possible protocols run on CODA.
type Protocol struct {
	Name        string
	Description string
}

const (
	PROTOCOL_FILE = "/home/riku/go/src/coda-service/conf/service.json"
)

func init() {
	var jsonConfig config.JSONConfig
	configer, err := jsonConfig.Parse(PROTOCOL_FILE)
	if err == nil {
		protocols, err := configer.DIY("Protocols")
		if err == nil {
			protocolsCasted := protocols.([]interface{})
			if len(protocolsCasted) > 0 {
				Protocols = make([]*Protocol, len(protocolsCasted))
				for idx, protocol := range protocolsCasted {
					protocolCasted := protocol.(map[string]interface{})
					name := protocolCasted["Name"].(string)
					description := protocolCasted["Description"].(string)
					Protocols[idx] = &Protocol{name, description}
				}
			}
		}
	} else {
		beego.BeeLogger.Alert("Can not open JSON. %v", err.Error())
	}
}

// GetAllProtocols Return all the possible protocols
func GetAllProtocols() []*Protocol {
	return Protocols
}
