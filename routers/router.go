// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/giros_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/cuenta_bancaria",
			beego.NSInclude(
				&controllers.CuentaBancariaController{},
			),
		),

		beego.NSNamespace("/giro",
			beego.NSInclude(
				&controllers.GiroController{},
			),
		),

		beego.NSNamespace("/forma_pago",
			beego.NSInclude(
				&controllers.FormaPagoController{},
			),
		),

		beego.NSNamespace("/orden_devolucion_detalle",
			beego.NSInclude(
				&controllers.OrdenDevolucionDetalleController{},
			),
		),

		beego.NSNamespace("/giro_orden_pago_detalle",
			beego.NSInclude(
				&controllers.GiroOrdenPagoDetalleController{},
			),
		),

		beego.NSNamespace("/giro_reversar",
			beego.NSInclude(
				&controllers.GiroReversarController{},
			),
		),

		beego.NSNamespace("/relacion_autorizacion_detalle",
			beego.NSInclude(
				&controllers.RelacionAutorizacionDetalleController{},
			),
		),

		beego.NSNamespace("/giro_contabilizacion",
			beego.NSInclude(
				&controllers.GiroContabilizacionController{},
			),
		),

		beego.NSNamespace("/giro_estado",
			beego.NSInclude(
				&controllers.GiroEstadoController{},
			),
		),

		beego.NSNamespace("/retiro_beneficiario_giro",
			beego.NSInclude(
				&controllers.RetiroBeneficiarioGiroController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
