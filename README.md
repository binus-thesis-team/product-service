# catalog-service

## Database with docker

1. run docker compose

```shell
docker compose up
```

2. run migration

```shell
sql-migrate up -config=dbconfig.yml -env="development"
go run database/migrate/up.go
```

3. connect from db client with credential

```
host=localhost
port=5432
db=product
user=user
password=password
```

### Sql-Migrate
- For complete documentation : https://github.com/rubenv/sql-migrate
- Create new file sql : sql-migrate new -config=dbconfig.yml -env="local" create_table_test
- Check status : sql-migrate status -config=dbconfig.yml -env="local"
- Dryrun query : sql-migrate up -config=dbconfig.yml -env="local" -dryrun
- Run query : sql-migrate up -config=dbconfig.yml -env="local"
