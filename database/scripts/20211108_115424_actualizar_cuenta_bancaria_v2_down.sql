DO $$
BEGIN
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='nombre_cuenta')
  THEN
      ALTER TABLE giros.cuenta_bancaria ALTER COLUMN nombre_cuenta TYPE integer USING (nombre_cuenta::integer);
      ALTER TABLE giros.cuenta_bancaria RENAME COLUMN nombre_cuenta TO nombre_id;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='numero_cuenta')
  THEN
      ALTER TABLE giros.cuenta_bancaria ALTER COLUMN numero_cuenta TYPE integer;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='tipo_cuenta_id')
  THEN
      ALTER TABLE giros.cuenta_bancaria RENAME COLUMN tipo_cuenta_id TO tipo_cuenta;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='pagadora')
  THEN
      ALTER TABLE giros.cuenta_bancaria DROP COLUMN pagadora;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='recaudadora')
  THEN
      ALTER TABLE giros.cuenta_bancaria DROP COLUMN recaudadora;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='cuatro_por_mil')
  THEN
      ALTER TABLE giros.cuenta_bancaria DROP COLUMN cuatro_por_mil;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='recurso_id')
  THEN
      ALTER TABLE giros.cuenta_bancaria DROP COLUMN recurso_id;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='area_funcional')
  THEN
      ALTER TABLE giros.cuenta_bancaria DROP COLUMN area_funcional;
  END IF;
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='divisa_id')
  THEN
      ALTER TABLE giros.cuenta_bancaria DROP COLUMN divisa_id;
  END IF;
END $$;
