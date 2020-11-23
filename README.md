# bookstore_users-api
Users API

Database

```
$ docker-compose up -d
```

```mysql
create table users
(
    id           bigint auto_increment
        primary key,
    first_name   varchar(255) null,
    last_name    varchar(126) null,
    email        varchar(255) not null,
    date_created datetime     not null,
    status       varchar(255) not null,
    password     varchar(255) not null,
    constraint users_email_idx
        unique (email)
);
```