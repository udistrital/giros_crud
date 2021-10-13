ALTER TABLE giros.cuenta_bancaria
RENAME COLUMN banco_id to sucursal_id;

COMMENT ON COLUMN giros.cuenta_bancaria.sucursal_id IS 'Traer de terceros_crud, tabla info_complementaria_tercero';
