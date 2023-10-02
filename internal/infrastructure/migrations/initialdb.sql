-- +migrate Up
CREATE TYPE "ability_type" AS ENUM (
  'support',
  'action',
  'reaction'
);

CREATE TYPE "item_type" AS ENUM (
  'general',
  'weapon',
  'ammunition',
  'armor',
  'conjuration',
  'potion'
);

CREATE TYPE "weapon_type" AS ENUM (
  'bruise',
  'drilling',
  'cut',
  'drilling_and_cut'
);

CREATE TYPE "weapon_distance" AS ENUM (
  'melee',
  'rod',
  'short',
  'medium',
  'long'
);

CREATE TYPE "armor_type" AS ENUM (
  'light',
  'heavy',
  'shield'
);

CREATE TYPE "conjuration_item_buff" AS ENUM (
  'magic_asset',
  'arcane_core',
  'mystic_core',
  'focus_magic',
  'register'
);

CREATE TABLE "ability" (
  "id" uuid PRIMARY KEY NOT NULL,
  "prerequisite" uuid,
  "name" varchar NOT NULL,
  "type" ability_type NOT NULL,
  "mana" int,
  "description" varchar NOT NULL,
  "level" int DEFAULT 1,
  "event" varchar
);

CREATE TABLE "attribute" (
  "id" uuid PRIMARY KEY NOT NULL,
  "strenght" int NOT NULL,
  "agility" int NOT NULL,
  "inteligence" int NOT NULL,
  "willing" int NOT NULL
);

CREATE TABLE "race" (
  "id" uuid PRIMARY KEY,
  "attribute_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "biology" text NOT NULL,
  "culture" text NOT NULL
);

CREATE TABLE "race_ability" (
  "id" uuid PRIMARY KEY,
  "ability_id" uuid,
  "race_id" uuid,
  "is_default" boolean NOT NULL DEFAULT false
);

CREATE TABLE "kind" (
  "id" uuid PRIMARY KEY,
  "attribute_id" uuid,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "races" text NOT NULL
);

CREATE TABLE "kind_ability" (
  "id" uuid PRIMARY KEY,
  "ability_id" uuid,
  "kind_id" uuid
);

CREATE TABLE "item" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" int NOT NULL,
  "weight" float NOT NULL,
  "description" varchar NOT NULL,
  "type" item_type NOT NULL
);

CREATE TABLE "weapon" (
  "id" uuid PRIMARY KEY,
  "damage" int NOT NULL,
  "type" weapon_type NOT NULL,
  "one_hand_fn" int,
  "two_hand_fn" int,
  "distance" weapon_distance NOT NULL,
  "observation" varchar[] NOT NULL
);

CREATE TABLE "armor" (
  "id" uuid PRIMARY KEY,
  "defense" int NOT NULL,
  "fn" int NOT NULL,
  "type" armor_type NOT NULL
);

CREATE TABLE "conjuration_item" (
  "id" uuid PRIMARY KEY,
  "fn" int NOT NULL,
  "buff" conjuration_item_buff[] NOT NULL
);

CREATE TABLE "potion" (
  "id" uuid PRIMARY KEY,
  "event" varchar NOT NULL
);

ALTER TABLE "ability" ADD FOREIGN KEY ("prerequisite") REFERENCES "ability" ("id");

ALTER TABLE "race" ADD FOREIGN KEY ("attribute_id") REFERENCES "attribute" ("id");

ALTER TABLE "race_ability" ADD FOREIGN KEY ("ability_id") REFERENCES "ability" ("id");

ALTER TABLE "race_ability" ADD FOREIGN KEY ("race_id") REFERENCES "race" ("id");

ALTER TABLE "kind" ADD FOREIGN KEY ("attribute_id") REFERENCES "attribute" ("id");

ALTER TABLE "kind_ability" ADD FOREIGN KEY ("ability_id") REFERENCES "ability" ("id");

ALTER TABLE "kind_ability" ADD FOREIGN KEY ("kind_id") REFERENCES "kind" ("id");

ALTER TABLE "weapon" ADD FOREIGN KEY ("id") REFERENCES "item" ("id");

ALTER TABLE "armor" ADD FOREIGN KEY ("id") REFERENCES "item" ("id");

ALTER TABLE "conjuration_item" ADD FOREIGN KEY ("id") REFERENCES "item" ("id");

ALTER TABLE "potion" ADD FOREIGN KEY ("id") REFERENCES "item" ("id");

-- +migrate Down
DROP TABLE "potion";

DROP TABLE "conjuration_item";

DROP TABLE "armor";

DROP TABLE "weapon";

DROP TABLE "item";

DROP TABLE "kind_ability";

DROP TABLE "kind";

DROP TABLE "race_ability";

DROP TABLE "race";

DROP TABLE "ability";

DROP TABLE "attribute";

DROP TYPE "conjuration_item_buff";

DROP TYPE "armor_type";

DROP TYPE "weapon_distance";

DROP TYPE "weapon_type";

DROP TYPE "item_type";

DROP TYPE "ability_type";