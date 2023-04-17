create table if not exists books (
    id serial primary key,
    title varchar not null,
    author varchar not null,
    version int not null default 1
);