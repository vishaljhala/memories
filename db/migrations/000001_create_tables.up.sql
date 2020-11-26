CREATE TABLE IF NOT EXISTS Users (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  username VARCHAR (50) UNIQUE NOT NULL,
  password VARCHAR (50) NOT NULL,
  email VARCHAR (300) UNIQUE NOT NULL
  );

CREATE TABLE IF NOT EXISTS Photos(
  ID serial PRIMARY KEY,
  CreatedAt time NOT NULL,
  UpdatedAt time NOT NULL,
  UserFk int NOT NULL,
  FilePath VARCHAR (512) NOT NULL
  );
