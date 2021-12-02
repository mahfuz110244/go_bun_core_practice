CREATE TYPE address_type_enum AS ENUM ('default', 'shipping','invoice');

CREATE TABLE "contact_type" (
  "contact_type_name" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "is_active" bool NOT NULL DEFAULT true,
"is_location_group" bool NOT NULL DEFAULT true,
  "contact_type_id" int2 NOT NULL DEFAULT nextval('contact_type_contact_type_id_seq'::regclass),
  CONSTRAINT "contact_type_pkey" PRIMARY KEY ("contact_type_id"),
  CONSTRAINT "contact_type_contact_type_name_un" UNIQUE ("contact_type_name")
);

CREATE TABLE "contact_group" (
  "contact_group_name" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "is_active" bool NOT NULL DEFAULT true,
  "contact_group_id" int2 NOT NULL DEFAULT nextval('contact_group_contact_group_id_seq'::regclass),
  "contact_type_id" int2 NOT NULL,
  CONSTRAINT "contact_group_pkey" PRIMARY KEY ("contact_group_id"),
  CONSTRAINT "contact_type_id_fk" FOREIGN KEY ("contact_type_id") REFERENCES "contact_type" ("contact_type_id") ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT "contact_type_sub_group_contact_sub_group_name_un" UNIQUE ("contact_group_name")
);


CREATE TABLE "contacts" (
  "contact_id" uuid NOT NULL DEFAULT uuid_generate_v1(),
  "first_name" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "last_name" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "business_name" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "job_title" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "location_group_id" uuid,
  "image" varchar COLLATE "pg_catalog"."default",
  "description" varchar COLLATE "pg_catalog"."default",
  "contact_group_id" int2 NOT NULL DEFAULT 0,
  CONSTRAINT "contacts_contact_id_pk" PRIMARY KEY ("contact_id"),
  CONSTRAINT "contact_group_id_fk" FOREIGN KEY ("contact_group_id") REFERENCES "contact_group" ("contact_group_id") ON DELETE NO ACTION ON UPDATE NO ACTION
);


CREATE TABLE "contacts_address" (
  "contacts_address_id" uuid NOT NULL DEFAULT uuid_generate_v1(),
  "contact_id" uuid NOT NULL,
  "address" varchar COLLATE "pg_catalog"."default",
  "zip" varchar COLLATE "pg_catalog"."default",
  "city" varchar COLLATE "pg_catalog"."default",
  "country" varchar COLLATE "pg_catalog"."default",
  "mobile" varchar COLLATE "pg_catalog"."default",
  "telephone" varchar COLLATE "pg_catalog"."default",
  "email" varchar COLLATE "pg_catalog"."default",
  "website" varchar COLLATE "pg_catalog"."default",
  "is_active" bool NOT NULL DEFAULT true,
  "address_type" "address_type_enum" NOT NULL DEFAULT 'default'::address_type_enum,
  CONSTRAINT "contacts_address_contacts_address_id_pk" PRIMARY KEY ("contacts_address_id"),
  CONSTRAINT "contact_id_fk" FOREIGN KEY ("contact_id") REFERENCES "contacts" ("contact_id") ON DELETE NO ACTION ON UPDATE NO ACTION
);
