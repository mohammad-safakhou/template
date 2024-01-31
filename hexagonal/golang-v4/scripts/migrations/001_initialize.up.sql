CREATE TABLE IF NOT EXISTS public.accounts
(
    "imsi"       TEXT      NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    CONSTRAINT "accounts_pk" PRIMARY KEY ("imsi")
) WITH (
      OIDS= FALSE
    );
