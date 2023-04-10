create table if not exists users(
   user_id serial primary key,
   user_name varchar (50) unique not null,
   login varchar (50) unique not null,
   password varchar (72) not null
);