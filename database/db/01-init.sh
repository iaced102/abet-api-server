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
  COMMIT;
EOSQL