-- +goose Up
CREATE TABLE chirps (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  body TEXT NOT NULL,
  user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE -- CHECK FOREIGN KEY CONSTRAINTS HERE: https://www.postgresql.org/docs/current/ddl-constraints.html#DDL-CONSTRAINTS-FK
);

-- +goose Down
DROP TABLE chirps;
