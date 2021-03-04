-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2021-03-01 14:48:44.252

-- foreign keys
ALTER TABLE giros.giro_contabilizacion
    DROP CONSTRAINT contabilizacion_giro_giro;

ALTER TABLE giros.giro
    DROP CONSTRAINT giro_cuenta_bancaria;

ALTER TABLE giros.giro_estado
    DROP CONSTRAINT giro_estado_giro_giro;

ALTER TABLE giros.giro
    DROP CONSTRAINT giro_forma_pago;

ALTER TABLE giros.giro_orden_pago_detalle
    DROP CONSTRAINT giro_oden_pago_detalle_giro;

ALTER TABLE giros.giro_reversar
    DROP CONSTRAINT giro_reversar_giro;

ALTER TABLE giros.orden_devolucion_detalle
    DROP CONSTRAINT orden_devolucion_detalle_giro;

ALTER TABLE giros.relacion_autorizacion_detalle
    DROP CONSTRAINT realcion_autorizacion_detalle_giro;

ALTER TABLE giros.retiro_beneficiario_giro
    DROP CONSTRAINT retiro_beneficiario_giro_giro;

-- tables
DROP TABLE giros.cuenta_bancaria;

DROP TABLE giros.forma_pago;

DROP TABLE giros.giro;

DROP TABLE giros.giro_contabilizacion;

DROP TABLE giros.giro_estado;

DROP TABLE giros.giro_orden_pago_detalle;

DROP TABLE giros.giro_reversar;

DROP TABLE giros.orden_devolucion_detalle;

DROP TABLE giros.relacion_autorizacion_detalle;

DROP TABLE giros.retiro_beneficiario_giro;

-- End of file.

