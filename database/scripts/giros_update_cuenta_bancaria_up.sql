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