CREATE TABLE IF NOT EXISTS "inquiries" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"transaction_id" text NULL,
	"payment_code" text NULL,
	CONSTRAINT "inquiries_pkey" PRIMARY KEY (id),
	CONSTRAINT "inquiries_transaction_id_key" UNIQUE (transaction_id)
);