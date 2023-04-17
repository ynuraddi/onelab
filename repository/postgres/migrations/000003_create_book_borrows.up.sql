create table if not exists book_borrows (
    id serial primary key,
    book_id int not null,
    user_id int not null,
    borrow_date timestamp default now() not null,
    return_date timestamp default null,
    foreign key (user_id) references users(id),
    foreign key (book_id) references books(id),
    version int not null default 1
);