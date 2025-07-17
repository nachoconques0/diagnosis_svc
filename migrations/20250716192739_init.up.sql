BEGIN;

CREATE SCHEMA IF NOT EXISTS top_doctor;

CREATE OR REPLACE FUNCTION top_doctor.set_updated_at ()
    RETURNS TRIGGER STABLE
    AS $plpgsql$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$plpgsql$
LANGUAGE plpgsql;

COMMIT;