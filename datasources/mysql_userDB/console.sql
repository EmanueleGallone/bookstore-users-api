CREATE DATABASE IF NOT EXISTS users_db
CHARACTER SET `utf8`;

use users_db;

create table if not exists users
(
    ID          int auto_increment,
    FirstName   varchar(255)                       not null,
    LastName    varchar(255)                       not null,
    Email       varchar(255)                       not null,
    DateCreated datetime default CURRENT_TIMESTAMP not null,
    Status      varchar(45)                        not null,
    Password    varchar(255)                       not null,
    constraint user_Email_uindex
        unique (Email),
    constraint user_ID_uindex
        unique (ID)
);

alter table user
    add primary key (ID);
