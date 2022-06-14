## Postgress
To run the project locally:
```bash
go docker-compose up
```
In another terminal run queries for db
```bash
docker exec -it gallery-db-1 /usr/bin/psql -U baloo -d lenslocked
```
to quit `control d`


`DROP TABLE IF EXISTS users` <br>
`CREATE TABLE users (` <br>
`id SERIAL,` <br>
`age INT,` <br>
`first_name TEXT,` <br>
`last_name TEXT,` <br>
`email TEXT` <br>
`);` <br>


###Notes:
Creating SQL tables
```sql
CREATE TABLE table_name (
    field_name TYPE PRIMARY KEY,
    field_name TYPE(args) CONTRANTS,
);
```


