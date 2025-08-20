CREATE SCHEMA IF NOT EXISTS vopi_chat;

CREATE TABLE "vopi_chat"."bots" (
	"id" SERIAL NOT NULL UNIQUE,
	"name" VARCHAR(255) NOT NULL,
	"uri" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."bots" IS 'tabela de integracao com os bots';
COMMENT ON COLUMN "vopi_chat"."bots"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."person_types" (
	"id" SERIAL NOT NULL UNIQUE,
	"description" VARCHAR(255),
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."person_types" IS 'tabela de tipo de pessoa';
COMMENT ON COLUMN "vopi_chat"."person_types"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."messages" (
	"id" SERIAL NOT NULL UNIQUE,
	"bot_id" INTEGER,
	"chat_participant_id" INTEGER,
	"payload" TEXT NOT NULL,
	"payload_type" INTEGER NOT NULL,
	"header" JSONB,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."messages" IS 'tabela de mensagens trocadas através do chat';
COMMENT ON COLUMN "vopi_chat"."messages"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."bot_flows" (
	"id" SERIAL NOT NULL UNIQUE,
	"bot_id" INTEGER NOT NULL,
	"flow_id" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."bot_flows" IS 'tabela de fluxos que o bot possui';
COMMENT ON COLUMN "vopi_chat"."bot_flows"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."persons" (
	"id" SERIAL NOT NULL UNIQUE,
	"person_type_id" INTEGER,
	-- nome da pessoa
	"name" VARCHAR(255) NOT NULL,
	"document" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT current_timestamp,
	"updated_at" TIMESTAMP NOT NULL DEFAULT current_timestamp,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."persons" IS 'tabela de dados gerais de pessoas';
COMMENT ON COLUMN "vopi_chat"."persons"."name" IS 'nome da pessoa';
COMMENT ON COLUMN "vopi_chat"."persons"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."contacts" (
	"id" SERIAL NOT NULL UNIQUE,
	"person_id" INTEGER NOT NULL,
	-- nome da pessoa
	"contact" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT current_timestamp,
	"updated_at" TIMESTAMP NOT NULL DEFAULT current_timestamp,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."contacts" IS 'tabela de dados de contatos';
COMMENT ON COLUMN "vopi_chat"."contacts"."contact" IS 'nome da pessoa';
COMMENT ON COLUMN "vopi_chat"."contacts"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."chat_participants" (
	"id" SERIAL NOT NULL UNIQUE,
	"contact_id" INTEGER NOT NULL,
	"chat_id" INTEGER NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"deleted_at" TIMESTAMP,
	-- true - ativo; false - inativo
	"active" BOOLEAN NOT NULL DEFAULT true,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."chat_participants" IS 'tabela de participantes do chat(pessoa, atendente)';
COMMENT ON COLUMN "vopi_chat"."chat_participants"."active" IS 'true - ativo; false - inativo';


CREATE TABLE "vopi_chat"."chats" (
	"id" SERIAL NOT NULL UNIQUE,
	"bot_flow_id" INTEGER NOT NULL,
	"chat_status_id" INTEGER NOT NULL,
	"chat_identifier" VARCHAR(255) NOT NULL,
	PRIMARY KEY("id")
);

COMMENT ON TABLE "vopi_chat"."chats" IS 'tabela que representa os chats(canais)';

ALTER TABLE "vopi_chat"."chat_participants"
ADD FOREIGN KEY("chat_id") REFERENCES "vopi_chat"."chats"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."chat_participants"
ADD FOREIGN KEY("contact_id") REFERENCES "vopi_chat"."contacts"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."messages"
ADD FOREIGN KEY("chat_participant_id") REFERENCES "vopi_chat"."chat_participants"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."messages"
ADD FOREIGN KEY("bot_id") REFERENCES "vopi_chat"."bots"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."bot_flows"
ADD FOREIGN KEY("bot_id") REFERENCES "vopi_chat"."bots"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."chats"
ADD FOREIGN KEY("bot_flow_id") REFERENCES "vopi_chat"."bot_flows"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."persons"
ADD FOREIGN KEY("person_type_id") REFERENCES "vopi_chat"."person_types"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "vopi_chat"."contacts"
ADD FOREIGN KEY("person_id") REFERENCES "vopi_chat"."persons"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;