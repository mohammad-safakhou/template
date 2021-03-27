create database backend;

\c backend;

create table if not exists accounts
(
    id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	phone_number VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
);

create table if not exists roles
(
    id          SERIAL primary key,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL ,
    name        varchar(300) NOT NULL,
    description varchar(1000)
);

create table if not exists permissions
(
    id      SERIAL primary key,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    route   varchar(300),
    method varchar(300),
    role_id int,
    foreign key (role_id) references roles (id) on delete CASCADE
);

create table if not exists role_user
(
    id      SERIAL primary key,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    user_id int,
    foreign key (user_id) references accounts (id),
    role_id int,
    foreign key (role_id) references roles (id)
);

insert into roles (created_at, is_active, name, description) VALUES (NOW(),true, 'test', '');
insert into role_user (created_at, user_id, role_id) values (NOW(), 1, 2);