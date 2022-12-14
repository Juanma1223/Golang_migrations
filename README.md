# Golang_migrations

# Usage

### You can use either single dash (-dir) or double dash (--dir) for any flag listed, in case these flags are not passed then default values will be used instead

- -dir: Directory where migrations are located (Default value is /doc/db/migrations)
- -h: Database host url (Default value is localhost)
- -u: Database user (Default value is root)
- -p: Database password (Default value is root)
- -P: Database port (Default value is 3306)
- -d: Database we are connecting to (Default value is test)


## Apply all migrations
gomigrate -dir doc/migrations -h <host> -u <dbUser> -p <dbPassword> -P <dbPort>
