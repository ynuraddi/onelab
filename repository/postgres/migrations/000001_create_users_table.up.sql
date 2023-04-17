create table if not exists users(
   id serial primary key,
   name varchar unique not null,
   login varchar unique not null,
   password varchar not null,
   is_active bool not null,
   version int not null default 1
);