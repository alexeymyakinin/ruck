CREATE SCHEMA "chat";

CREATE TABLE "chat"."user"
(
    "id"       BIGSERIAL PRIMARY KEY,
    "email"    VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "password" VARCHAR      NOT NULL
);
CREATE UNIQUE INDEX "user_email_uix" ON "chat"."user" ("email");
CREATE UNIQUE INDEX "user_username_uix" ON "chat"."user" ("username");

CREATE TABLE "chat"."chat"
(
    "id"   BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "chat"."chat_user"
(
    "id"   BIGSERIAL PRIMARY KEY,
    "chat" BIGINT NOT NULL REFERENCES "chat"."chat" ("id"),
    "user" BIGINT NOT NULL REFERENCES "chat"."user" ("id")
);
CREATE INDEX "chat_user_chat_ix" ON "chat"."chat_user" ("chat");
CREATE INDEX "chat_user_user_ix" ON "chat"."chat_user" ("user");

CREATE TABLE "chat"."chat_message"
(
    "id"        BIGSERIAL PRIMARY KEY,
    "text"      VARCHAR(255),
    "chat"      BIGINT    NOT NULL REFERENCES "chat"."chat" ("id"),
    "sender"    BIGINT    NOT NULL REFERENCES "chat"."user" ("id"),
    "timestamp" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX "chat_message_chat_ix" ON "chat"."chat_message" ("chat");
CREATE INDEX "chat_message_sender_ix" ON "chat"."chat_message" ("sender");
