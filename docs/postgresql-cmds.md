## Index

1. [Postgresql commands](#postgresql-commands)
2. [Postgresql query commands](#postgresql-query-commands)
3. [Extra postgresql commands](#extra-postgresql-commands)
6. [Popular Dbs, and engines](#popular-databases)

## POSTGRESQL COMMANDS

- Start Postgres service in background

```bash
  sudo service postgresql start 
```

- Change Postgres password

```bash
  sudo passwd postgres
```

- Enter the `psql` shell:

```bash
  # Mac:
  psql postgres
  # Linux:
  sudo -u postgres psql
```

- Enter database straight from terminal:

```bash
  psql "<connection_string>"
  # connection_string is of format: "protocol://username:password@host:port/database"
  # You can then run: \dt // to list tables
```

- List all databases

```bash
  \l
  # or simply run `psql -l` from terminal
```

- Connect/Switch to a specific database:

```bash
  \c database_name
```

- List all tables in the current database:

```bash
  \dt
  # Use \dt+ to see table sizes and descriptions
```

- Describe a specific table's columns and data types:

```bash
  \d table_name
```

- Exit the `psql` shell:

```bash
  \q
```

- Create a new database

```bash
  CREATE DATABASE database_name
```

- Connect to the new database

```bash
  \c database_name

  # you should see database_name#
```

- Set the user password (Linux only)

```bash
  ALTER USER postgres WITH PASSWORD 'postgres';
```

## Postgresql Query Commands

From here you can run queries like

```bash
  SELECT version();
```

## Extra Postgresql Commands

- Fix database collation mismatch warnings (OS update fix):

```sql
  REINDEX DATABASE database_name;
  ALTER DATABASE database_name REFRESH COLLATION VERSION;
```

- Backup a database to a file (Run from terminal):

```bash
  pg_dump -U username -d database_name -F c -f backup_file.dump
```

- Restore a database from a file (Run from terminal):

```bash
  pg_restore -U username -d database_name -1 backup_file.dump
```

- View running queries and active connections:

```sql
  SELECT pid, user, pg_stat_activity.query, start_time 
  FROM pg_stat_activity 
  WHERE state = 'active';
```

- Forcefully terminate a stuck query or connection:

```sql
  SELECT pg_terminate_backend(pid);
```

## Popular Databases

[PostgreSQL](https://www.postgresql.org/): A fantastic open-source SQL database.
[MySQL](https://www.mysql.com/): Another open-source SQL database. Less fantastic IMO.
[MongoDB](https://www.mongodb.com/): A popular open-source NoSQL document database.
[Firebase](https://firebase.google.com/): A popular cloud-based NoSQL database service.
[SQLite](https://www.sqlite.org/index.html): A popular embedded SQL database.

Browse [DB Engines](https://db-engines.com/en/ranking) to dive deeper into the world of database technologies.
