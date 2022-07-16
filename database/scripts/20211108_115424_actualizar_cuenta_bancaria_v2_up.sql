ALTER TABLE giros.cuenta_bancaria
ALTER COLUMN numero_cuenta TYPE numeric;

ALTER TABLE giros.cuenta_bancaria
RENAME COLUMN nombre_id TO nombre_cuenta;

ALTER TABLE giros.cuenta_bancaria
ALTER COLUMN nombre_cuenta TYPE varchar;

ALTER TABLE giros.cuenta_bancaria
RENAME COLUMN tipo_cuenta TO tipo_cuenta_id;

ALTER TABLE giros.cuenta_bancaria
ADD COLUMN pagadora boolean NOT NULL,
ADD COLUMN recaudadora boolean NOT NULL,
ADD COLUMN cuatro_por_mil boolean NOT NULL,
ADD COLUMN recurso_id varchar NOT NULL,
ADD COLUMN area_funcional integer NOT NULL,
ADD COLUMN divisa_id integer NOT NULL;

COMMENT ON COLUMN giros.cuenta_bancaria.numero_cuenta IS 'Numero de cuenta bancaria';
COMMENT ON COLUMN giros.cuenta_bancaria.nombre_cuenta IS 'Nombre de la cuenta bancaria';
COMMENT ON COLUMN giros.cuenta_bancaria.tipo_cuenta_id IS 'Traer de parametros_crud, tabla parametros';
COMMENT ON COLUMN giros.cuenta_bancaria.pagadora IS 'True si la cuenta bancaria es pagadora';
COMMENT ON COLUMN giros.cuenta_bancaria.recaudadora IS 'True si la cuenta bancaria es recaudadora';
COMMENT ON COLUMN giros.cuenta_bancaria.cuatro_por_mil IS 'True si la cuenta bancaria tiene 4x1000';
COMMENT ON COLUMN giros.cuenta_bancaria.recurso_id IS 'Traer de plan_cuentas_mongo_crud, tabla arbol_rubro';
COMMENT ON COLUMN giros.cuenta_bancaria.area_funcional IS 'Area funcional de la cuenta bancaria';
COMMENT ON COLUMN giros.cuenta_bancaria.divisa_id IS 'Traer de parametros_crud, tabla parametros';
