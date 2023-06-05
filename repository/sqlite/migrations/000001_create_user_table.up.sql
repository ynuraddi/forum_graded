create table if not exists users (
   id integer primary key autoincrement,
   login varchar(50) unique not null,
   email varchar(50) unique not null,
   hash_password varchar(100) not null,
   is_active boolean default true,
   version integer not null default 1
);