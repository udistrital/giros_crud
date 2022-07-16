-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2021-03-01 14:48:44.252

CREATE SCHEMA IF NOT EXISTS giros;
-- ALTER SCHEMA giros OWNER TO test;
SET search_path TO pg_catalog,public,giros;

-- tables
-- Table: cuenta_bancaria
CREATE TABLE giros.cuenta_bancaria (
    id serial  NOT NULL,
    nombre_id int  NOT NULL,
    numero_cuenta int  NOT NULL,
    banco_id int  NOT NULL,
    tipo_cuenta int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT cuenta_bancaria_pk PRIMARY KEY (id)
);

-- Table: forma_pago
CREATE TABLE giros.forma_pago (
    id serial  NOT NULL,
    numero_orden decimal(5,2)  NULL,
    forma_pago_id int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT forma_pago_pk PRIMARY KEY (id)
);

-- Table: giro
CREATE TABLE giros.giro (
    id serial  NOT NULL,
    cuenta_bancaria_id int  NOT NULL,
    forma_pago_id int  NOT NULL,
    vigencia int  NOT NULL,
    area_funcional int  NOT NULL,
    tipo_documento_id int  NOT NULL,
    tipo_giro_id int  NOT NULL,
    numero_seleccion int  NOT NULL,
    valor_total decimal(10,2)  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT giro_pk PRIMARY KEY (id)
);

-- Table: giro_contabilizacion
CREATE TABLE giros.giro_contabilizacion (
    id serial  NOT NULL,
    giro_id int  NOT NULL,
    concepto varchar  NOT NULL,
    codigo int  NOT NULL,
    tipo_comprobante int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT giro_contabilizacion_pk PRIMARY KEY (id)
);

-- Table: giro_estado
CREATE TABLE giros.giro_estado (
    Id serial  NOT NULL,
    giro_id int  NOT NULL,
    estado_giro_id int  NOT NULL,
    responsable int  NOT NULL,
    observaciones varchar  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT giro_estado_pk PRIMARY KEY (Id)
);

-- Table: giro_orden_pago_detalle
CREATE TABLE giros.giro_orden_pago_detalle (
    id serial  NOT NULL,
    giro_id int  NOT NULL,
    vigencia int  NOT NULL,
    oden_pago_id int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT giro_orden_pago_detalle_pk PRIMARY KEY (id)
);

-- Table: giro_reversar
CREATE TABLE giros.giro_reversar (
    id serial  NOT NULL,
    giro_Id int  NOT NULL,
    responsable int  NOT NULL,
    observaciones varchar  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT giro_reversar_pk PRIMARY KEY (id)
);

-- Table: orden_devolucion_detalle
CREATE TABLE giros.orden_devolucion_detalle (
    id serial  NOT NULL,
    giro_id int  NOT NULL,
    vigencia int  NOT NULL,
    orden_devolucion_id int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT orden_devolucion_detalle_pk PRIMARY KEY (id)
);

-- Table: relacion_autorizacion_detalle
CREATE TABLE giros.relacion_autorizacion_detalle (
    id serial  NOT NULL,
    giro_id int  NOT NULL,
    vigencia int  NOT NULL,
    relacion_autorizacion_id int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT relacion_autorizacion_detalle_pk PRIMARY KEY (id)
);

-- Table: retiro_beneficiario_giro
CREATE TABLE giros.retiro_beneficiario_giro (
    id serial  NOT NULL,
    giro_id int  NOT NULL,
    consecutivo int  NOT NULL,
    tipo_documento_beneficiario int  NOT NULL,
    documento_beneficiario int  NOT NULL,
    observaciones int  NOT NULL,
    activo boolean  NOT NULL,
    fecha_creacion timestamp  NOT NULL,
    fecha_modificacion timestamp  NOT NULL,
    CONSTRAINT retiro_beneficiario_giro_pk PRIMARY KEY (id)
);

-- ALTER TABLE giros.cuenta_bancaria OWNER TO test;
-- ALTER TABLE giros.forma_pago OWNER TO test;
-- ALTER TABLE giros.giro_contabilizacion OWNER TO test;
-- ALTER TABLE giros.giro_estado OWNER TO test;
-- ALTER TABLE giros.giro_orden_pago_detalle OWNER TO test;
-- ALTER TABLE giros.giro_reversar OWNER TO test;
-- ALTER TABLE giros.giro OWNER TO test;
-- ALTER TABLE giros.orden_devolucion_detalle OWNER TO test;
-- ALTER TABLE giros.relacion_autorizacion_detalle OWNER TO test;
-- ALTER TABLE giros.retiro_beneficiario_giro OWNER TO test;



-- foreign keys
-- Reference: contabilizacion_giro_giro (table: giro_contabilizacion)
ALTER TABLE giros.giro_contabilizacion ADD CONSTRAINT contabilizacion_giro_giro
    FOREIGN KEY (giro_id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: giro_cuenta_bancaria (table: giro)
ALTER TABLE giros.giro ADD CONSTRAINT giro_cuenta_bancaria
    FOREIGN KEY (cuenta_bancaria_id)
    REFERENCES giros.cuenta_bancaria (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: giro_estado_giro_giro (table: giro_estado)
ALTER TABLE giros.giro_estado ADD CONSTRAINT giro_estado_giro_giro
    FOREIGN KEY (giro_id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: giro_forma_pago (table: giro)
ALTER TABLE giros.giro ADD CONSTRAINT giro_forma_pago
    FOREIGN KEY (forma_pago_id)
    REFERENCES giros.forma_pago (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: giro_oden_pago_detalle_giro (table: giro_orden_pago_detalle)
ALTER TABLE giros.giro_orden_pago_detalle ADD CONSTRAINT giro_oden_pago_detalle_giro
    FOREIGN KEY (giro_id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: giro_reversar_giro (table: giro_reversar)
ALTER TABLE giros.giro_reversar ADD CONSTRAINT giro_reversar_giro
    FOREIGN KEY (giro_Id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: orden_devolucion_detalle_giro (table: orden_devolucion_detalle)
ALTER TABLE giros.orden_devolucion_detalle ADD CONSTRAINT orden_devolucion_detalle_giro
    FOREIGN KEY (giro_id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: realcion_autorizacion_detalle_giro (table: relacion_autorizacion_detalle)
ALTER TABLE giros.relacion_autorizacion_detalle ADD CONSTRAINT realcion_autorizacion_detalle_giro
    FOREIGN KEY (giro_id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: retiro_beneficiario_giro_giro (table: retiro_beneficiario_giro)
ALTER TABLE giros.retiro_beneficiario_giro ADD CONSTRAINT retiro_beneficiario_giro_giro
    FOREIGN KEY (giro_id)
    REFERENCES giros.giro (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.

