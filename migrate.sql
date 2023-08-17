create table metadatas
(
    id           integer
        primary key autoincrement,
    name         text,
    content_type text,
    extension    text,
    size         integer,
    path         text     default null,
    url          text     default null,
    created_at   datetime default current_timestamp,
    updated_at   datetime default current_timestamp,
    deleted_at   datetime default null
);
