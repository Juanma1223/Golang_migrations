# Golang_migrations

This CLI and golang library is intended to manage MySQL migrations, you can either create empty migrations and fill them as needed or parse a create table migration from a golang struct file

# Usage

### You can use either single dash (-dir) or double dash (--dir) for any flag listed, in case these flags are not passed then default values will be used instead, you can use the enviroments listed in the migrationConf.json file and set a default database name to use.

- -dir: Directory where migrations are located (Default value is /doc/db/migrations)
- -h: Database host url (Default value is localhost)
- -u: Database user (Default value is root)
- -p: Database password (Default value is root)
- -P: Database port (Default value is 3306)
- -d: Database we are connecting to (Default value is test)
- -create: Create new migration file (If this flag is used, no migrations will be applied)
- -fix: Fixes migration files versions if they are repeated or not sequential
- -parse: Creates migration from golang struct, this flag receives the path to the file
- -version: Returns database migrations version
- -change: Flag to change the setted default database name in a specific enviroment

## Apply all migrations
migration -dir doc/migrations -h "host" -u "dbUser" -p "dbPassword" -P "dbPort"

## Create new migration file
migration --create create-users-table -dir migrations/users

## Change default database name in a selected enviroment
migration -change
