## Postgress
To run the postgress image in docker:
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

First example of how to insert info into the table
```sql
INSERT INTO users VALUE (1, 22, 'John', 'Smith', 'john@mail.com');
```

Second example of how to insert info into the table allows to skip some fields
```sql
INSERT INTO users (age, email, first_name, last_name)
```
```sql
VALUES (30, 'bobby@mail.com', 'Bobby', 'Singer');
```

To see users in the table
```sql
SELECT * FROM users;
```

```sql
SELECT id, email FROM users;
```

Filter results
```sql
SELECT * FROM users WHERE email='bobby@mail.com';
```

```sql
SELECT * FROM users WHERE email='bobby@mail.com';
```

```sql
SELECT * FROM users WHERE age > 22 OR last_name='Singer';
```

```sql
SELECT * FROM users WHERE age > 22 AND last_name='Singer';
```


Update records
```sql
UPDATE users SET first_name='Johny' WHERE id=1;
```

Delete records
```sql
DELETE FROM users WHERE id=1;
```


###Bcrypt

To hash the password
```bash
go run bcrypt/bcrypt.go hash "secret-password"
```

To compare the password (not matching)
```bash
go run bcrypt/bcrypt.go compare "secret-password" {put generated hash from previous command here in single quotes}
```
