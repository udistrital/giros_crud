DO $$
BEGIN
  IF EXISTS(SELECT *
    FROM information_schema.columns
    WHERE table_name='cuenta_bancaria' and column_name='sucursal_id')
  THEN
      ALTER TABLE giros.cuenta_bancaria RENAME COLUMN sucursal_id TO banco_id;
  END IF;
END $$;
