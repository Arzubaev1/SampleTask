CREATE TABLE phone_number(
    "id" VARCHAR PRIMARY KEY,
    "user_id" VARCHAR NOT NULL REFERENCES "users"("id"),
    "phone" VARCHAR NOT NULL,
    "is_fax" BOOLEAN DEFAULT TRUE,
    "description" VARCHAR NOT NULL
);