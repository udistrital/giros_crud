package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroContabilizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroEstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroOrdenPagoDetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:GiroReversarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:OrdenDevolucionDetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RelacionAutorizacionDetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/giros_crud/controllers:RetiroBeneficiarioGiroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
