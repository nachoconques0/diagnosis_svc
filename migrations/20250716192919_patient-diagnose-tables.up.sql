BEGIN;

CREATE TABLE top_doctor.patient (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  dni TEXT NOT NULL,
  email TEXT NOT NULL,
  phone TEXT,
  address TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TRIGGER set_updated_at
  BEFORE INSERT OR UPDATE ON top_doctor.patient
  FOR EACH ROW
  EXECUTE PROCEDURE top_doctor.set_updated_at ();
	
CREATE TABLE top_doctor.diagnose (
  id UUID PRIMARY KEY,
  patient_id UUID NOT NULL REFERENCES top_doctor.patient(id),
  diagnosis TEXT NOT NULL,
  prescription TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE TRIGGER set_updated_at
  BEFORE INSERT OR UPDATE ON top_doctor.diagnose
  FOR EACH ROW
  EXECUTE PROCEDURE top_doctor.set_updated_at ();


CREATE TABLE top_doctor.user (
  id UUID PRIMARY KEY,
  nickname TEXT NOT NULL,
  password TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;