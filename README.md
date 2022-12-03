# Go tests and snippets

A lot of this code comes from books and is a playground to experiment and test, prior to implementation in other repos and production projects.
thus not all the code is working correctly at all times since these are meant to be failed attempts and rough work.

Topics are broadly focused on algos, networking, DB, distributed systems and pen testing. Go 1.18.3 is required for this repo.

## Run docker for postgres:

``` bash
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres

docker run -it --rm --link some-postgres:postgres postrgres psql -h postgres -U postgres
````

## Run docker for mysql:

``` bash
docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql

docker run -it --link some-mysql:mysql --rm mysql sh -c 'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -P "$MYSQL_PORT_3306_TCP_PORT" -uroot -p  "$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
```

## Run docker for mongoDB:

``` bash
docker run --name some-mongo -p 27017:27017 mongo

docker start some-mongo

docker run -it --link some-mongo:mongo --rm mongo sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/store"'

```