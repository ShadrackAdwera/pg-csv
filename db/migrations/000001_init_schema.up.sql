CREATE TYPE "e_sales_channel" AS ENUM (
  'Offline',
  'Online'
);

CREATE TYPE "e_order_priority" AS ENUM (
  'C',
  'H',
  'M',
  'L'
);

CREATE TABLE "sales" (
  "id" bigserial PRIMARY KEY,
  "region" varchar(50) NOT NULL,
  "country" varchar(50) NOT NULL,
  "item_type" varchar(25) NOT NULL,
  "sales_channel" e_sales_channel NOT NULL DEFAULT 'Online',
  "order_priority" e_order_priority NOT NULL DEFAULT 'M',
  "order_date" timestamptz NOT NULL DEFAULT (now()),
  "order_id" bigint NOT NULL,
  "ship_date" timestamptz NOT NULL DEFAULT (now()),
  "units_sold" int NOT NULL,
  "unit_price" real NOT NULL,
  "unit_cost" real NOT NULL,
  "total_revenue" bigint NOT NULL,
  "total_cost" bigint NOT NULL,
  "total_profit" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);