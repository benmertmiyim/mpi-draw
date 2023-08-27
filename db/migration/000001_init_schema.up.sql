CREATE TABLE "codes" (
  "code" varchar PRIMARY KEY,
  "phone" varchar,
  "email" varchar,
  "birthday" date,
  "name_surname" varchar,
  "address" varchar,
  "city_code" int,
  "registered_date" timestamp,
  "ip" varchar,
  "active" boolean NOT NULL DEFAULT TRUE,
  "banned" boolean NOT NULL DEFAULT FALSE,
  "banned_reason" varchar,
  "tc_no" varchar,
  "used" boolean NOT NULL DEFAULT FALSE
);