CREATE TYPE "kodetransaksi" AS ENUM (
  'C',
  'D'
);

CREATE TABLE "nasabah" (
  "id" SERIAL PRIMARY KEY,
  "nama" varchar,
  "nik" varchar(16) NOT NULL,
  "no_hp" varchar(13) NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "rekening" (
  "no_rekening" varchar(10) PRIMARY KEY,
  "nasabah_id" int,
  "saldo" float NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "transaksi" (
  "id" SERIAL PRIMARY KEY,
  "no_rekening" varchar(10),
  "nominal" float NOT NULL,
  "kode_transaksi" kodetransaksi,
  "created_at" timestamp DEFAULT (now())
);

CREATE UNIQUE INDEX ON "nasabah" ("nik");

CREATE UNIQUE INDEX ON "nasabah" ("no_hp");

CREATE UNIQUE INDEX ON "rekening" ("no_rekening");

CREATE INDEX "kode_transaksi_idx" ON "transaksi" ("kode_transaksi");

ALTER TABLE "rekening" ADD FOREIGN KEY ("nasabah_id") REFERENCES "nasabah" ("id");

ALTER TABLE "transaksi" ADD FOREIGN KEY ("no_rekening") REFERENCES "rekening" ("no_rekening");
