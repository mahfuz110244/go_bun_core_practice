DROP TABLE IF EXISTS status CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS public.status (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
	deleted_at TIMESTAMP WITHOUT TIME ZONE NULL DEFAULT NULL,
	created_by UUID,
	updated_by UUID,
	name VARCHAR(20) NOT NULL,
	description VARCHAR (255),
	is_active BOOLEAN NOT NULL DEFAULT TRUE,
	order_no smallint DEFAULT 0,
    total_bill numeric(12,2) DEFAULT 0,
	UNIQUE(name)
);

-- Step 1: Create trigger Function
CREATE OR REPLACE FUNCTION trigger_set_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW() AT TIME ZONE 'utc';
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Step 2: Then create a trigger for each table that has the column updated_at
DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_set_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_timestamp()', t,t);
    END loop;
END;
$$ language 'plpgsql';
