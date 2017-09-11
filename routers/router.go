// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"coda-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/protocol", &controllers.ProtocolController{}),
		beego.NSRouter("/encrypt/:src:string&:dst:string", &controllers.EncryptionController{}),
		beego.NSRouter("/session/", &controllers.SessionController{}),
		beego.NSRouter("/keygen/:protocol", &controllers.KeyGenController{}),
	)
	beego.AddNamespace(ns)
}
