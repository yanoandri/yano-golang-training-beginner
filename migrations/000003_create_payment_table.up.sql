CREATE TABLE "payments" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"transaction_id" text NULL,
	"payment_code" text NULL,
	"name" text NULL,
	"amount" integer NULL,
	CONSTRAINT "payments_pkey" PRIMARY KEY (id),
	CONSTRAINT "payments_transaction_id_key" UNIQUE (transaction_id)
);