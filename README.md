# Go_Lang_Course

## Run docker for sql:

``` bash
docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql

docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
````

``` bash
docker run -it --link some-mysql:mysql -rm mysql sh -c 'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -P "MYSQL_PORT_3306_TCP_PORT" -uroot -p  "$MYSQL_ENV_MYSQL_ROOT_PASSWORD"

docker run -it --rm --link some-postgres:postgres postrgres psql -h postgres -U postgres
```