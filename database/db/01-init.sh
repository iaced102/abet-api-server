set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS users (
	  id varchar(80) NOT NULL PRIMARY KEY,
	  user_name varchar(26) NOT NULL ,
	  user_type INT,
	  created_at varchar(50),
      updated_at varchar(50),
      password varchar(200),
      crypt_password varchar(400),
      first_name varchar(20),
      last_name varchar(20),
      email varchar(80),
      custom_field text,
	  UNIQUE(id, user_name)
	);
  CREATE TABLE IF NOT EXISTS document (
	  id varchar(80) NOT NULL PRIMARY KEY,
    name varchar(80),
	  created_by varchar(80) NOT NULL ,
	  created_at varchar(50),
    updated_at varchar(50),
    assessor_id varchar(150),
    verifier_id varchar(150),
    superviser_id varchar(80),
    evaluate_field varchar(20),
    so_document_id varchar(80)
	  UNIQUE(id)
	);
  CREATE TABLE IF NOT EXISTS so_document (
    id varchar(80) NOT NULL PRIMARY KEY,
    name varchar(80),
    created_at varchar(80) NOT NULL ,
	  created_by varchar(80) NOT NULL ,
    identifier_id varchar(60),
    description text,
	  UNIQUE(id)
  );
  CREATE TABLE IF NOT EXISTS report (
	  id varchar(80) NOT NULL PRIMARY KEY,
	  document_id varchar(80),
    field varchar(20),
	  UNIQUE(id)
	);
  CREATE TABLE IF NOT EXISTS detail_report (
	  id varchar(80) NOT NULL PRIMARY KEY,
	  student_id varchar(25),
    first_name varchar(30),
    last_name varchar(30),
    class_id varchar(30),
    report_id varchar(80),
    value text,
    UNIQUE(id)
	);
  COMMIT;
EOSQL