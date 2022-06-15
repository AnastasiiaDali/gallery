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


`DROP TABLE IF EXISTS users;` <br>
`CREATE TABLE users (` <br>
`id SERIAL,` <br>
`age INT,` <br>
`first_name TEXT,` <br>
`last_name TEXT,` <br>
`email TEXT` <br>
`);` <br>


###Notes:
Creating SQL table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
```


