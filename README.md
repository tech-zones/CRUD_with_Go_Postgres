#CRUD app in Golanf and Postgres

#Table SQL
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT UNIQUE
);

```
