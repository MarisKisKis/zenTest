CREATE TABLE IF NOT EXISTS "users"
(
    "id"   INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "age"  int
);