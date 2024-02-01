# Hexagonal Architecture sample application

## Development

### Init local environment
1. Copy file `.env.example` and rename to `.env`

2. Update env vars to fit your local

3. Start local services
    ```shell
    make local-db
    ```

4. Migration

   ```shell
   make db/migrate
   ```

5. Run the server
    ```shell
    make run
    ```

### Unit test

```shell
make test
```

### Create new migration file

```shell
sql-migrate new -env="development" create-books-table
```

- Result: `Created migration migrations/20231219104808-create-books-table.sql`