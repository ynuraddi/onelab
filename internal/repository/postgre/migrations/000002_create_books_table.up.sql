create table if not exists books (
    book_id serial primary key,
    name varchar (50) not null,
    author varchar (50) not null,
    unique(name, author)
);