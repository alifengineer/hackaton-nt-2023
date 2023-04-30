
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
    "guid" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "first_name" varchar(255) NOT NULL,
    "last_name" varchar(255) NOT NULL,
    "email" varchar NOT NULL UNIQUE,
    "password" varchar(255) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" INTEGER DEFAULT 0
);

CREATE UNIQUE INDEX "users_phone_deleted_at_unique" ON "users" ("email", "deleted_at");

INSERT INTO "users"("first_name", "last_name", "email","password")
VALUES ('Tolib', 'Dilmurodov', 'tolibdilmurodov98@gmail.com', '$argon2id$v=19$models=65536,t=3,p=4$335KxmjmC5iB3YrgWBJ4Gw$dNTUoBZezp9zTTF29jGQ4gDMgEx5k+klCaiVkmTEoEo');
-- tolib2003