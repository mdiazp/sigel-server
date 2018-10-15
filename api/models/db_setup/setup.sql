drop table notification cascade;
drop table reservation cascade;
drop table users cascade;
drop table local_admin;
drop table area_admin;
drop table local cascade;
drop table area cascade;

create table k_user(
    id                          serial primary key,
    username                    varchar(100) not null unique, 
    name                        varchar(100) not null,
    email                       varchar(100) not null,
    send_notifications_to_email boolean not null,
    rol                         varchar(100) not null,
    enable                      boolean not null
);

create table area(
    id                          serial primary key,
    name                        varchar(100) not null,
    description                 varchar(1024) not null,
    location                    varchar(1024),
);

create table local(
    id                          serial primary key,
    area_id                     integer references area(id) on delete cascade not null,
    name                        varchar(100) not null,
    description                 varchar(1024) not null,
    location                    varchar(1024) not null,
    working_moths               varchar(12) not null,
    working_week_days           varchar(7) not null,
    working_begin_time_hours    integer not null,
    working_begin_time_minutes  integer not null,
    working_end_time_hours      integer not null,
    working_end_time_minutes    integer not null,
    enable_to_reserve           boolean not null
);

create table reservation(
    id                      serial primary key,
    user_id                 integer references k_user(id) on delete cascade not null,
    local_id                integer references local(id) on delete cascade not null,
    activity_name           varchar(100) not null,
    activity_description    varchar(1024) not null,
    begin_time              timestamp not null,
    end_time                timestamp not null,
    confirmed               boolean not null,
    pending                 boolean not null
);

create table notification(
    id                      serial primary key,
    to_user                 integer references k_user(id) on delete cascade not null,
    message                 varchar(1024) not null,
    creation_time           timestamp not null,
    readed                  boolean not null
);

create table area_admin(
    id                      serial primary key,
    user_id                 integer references k_user(id) on delete cascade not null,
    area_id                 integer references area(id) on delete cascade not null,
);

create table local_admin(
    id                      serial primary key,
    user_id                 integer references k_user(id) on delete cascade not null,
    local_id                integer references local(id) on delete cascade not null,
);