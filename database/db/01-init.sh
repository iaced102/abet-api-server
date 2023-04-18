set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS users (
	  id CHAR(80) NOT NULL PRIMARY KEY,
	  user_name CHAR(26) NOT NULL ,
	  user_type INT,
	  created_at CHAR(50),
      updated_at CHAR(50),
      password CHAR(200),
      crypt_password CHAR(400),
	  UNIQUE(id, user_name)
	);
  CREATE TABLE IF NOT EXISTS document (
	  id CHAR(80) NOT NULL PRIMARY KEY,
    name CHAR(80),
	  created_by CHAR(80) NOT NULL ,
	  created_at CHAR(50),
    updated_at CHAR(50),
    assessor_id CHAR(150),
    verifier_id CHAR(150),
    superviser_id CHAR(80),
	  UNIQUE(id)
	);
  CREATE TABLE IF NOT EXISTS report (
	  id CHAR(80) NOT NULL PRIMARY KEY,
	  document_id CHAR(80),
    field CHAR(20),
	  UNIQUE(id)
	);
  CREATE TABLE IF NOT EXISTS report (
	  id CHAR(80) NOT NULL PRIMARY KEY,
	  document_id CHAR(80),
    field CHAR(20),
	  UNIQUE(id)
	);
  CREATE TABLE IF NOT EXISTS detail_report (
	  id CHAR(80) NOT NULL PRIMARY KEY,
	  student_id CHAR(25),
    first_name CHAR(30),
    last_name CHAR(30),
    class_id CHAR(30),
    report_id CHAR(80),
    value text,
    UNIQUE(id)
	);
  COMMIT;
EOSQL