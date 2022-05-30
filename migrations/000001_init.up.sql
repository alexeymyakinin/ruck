create table account
(
    id       bigserial primary key,
    email    varchar(255) not null,
    username varchar(255) not null,
    password varchar      not null
);
create unique index account_email_uix on account (email);
create unique index account_username_uix on account (username);

create table chat
(
    id   bigserial primary key,
    name varchar(255) not null
);

create table chat_account
(
    id      bigserial primary key,
    chat    bigint not null references chat (id),
    account bigint not null references account (id)
);
create index chat_account_chat_ix on chat_account (chat);
create index chat_account_account_ix on chat_account (account);

create table chat_message
(
    id        bigserial primary key,
    text      varchar(255),
    chat      bigint    not null references chat (id),
    sender    bigint    not null references account (id),
    timestamp timestamp not null default current_timestamp
);
create index chat_message_chat_ix on chat_message (chat);
create index chat_message_sender_ix on chat_message (sender);

