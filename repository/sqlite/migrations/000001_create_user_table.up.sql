create table if not exists users(
   id serial primary key,
   login varchar unique not null,
   email varchar unique not null,
   hash_password varchar not null,
   is_active boolean not null default true,
   version int not null default 1
);