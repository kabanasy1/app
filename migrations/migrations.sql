create table users(
    id serial not null
        constraint user_pk
        primary key,
    name text,
    surname text,
    room int64,
    phone text,
    position text
)