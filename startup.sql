create database [veterinary]
go

use [veterinary]
go

create table roles(
	id int not null identity,
	role_name varchar(250) not null,
	role_description varchar(250),
	primary key(id)
);
go

if not exists(select * from [roles] where [roles].id = 0 or [roles].id = 1 or [roles].id = 2) 
insert into [roles](role_name,role_description) values
('admin','System admin, all permissions granted'),
('doctor','Can view pets and owners information, and modify them'),
('client','Usual pet owner, can edit their own profile, add new pets and reserve turns');
go

create table accounts(
	id int not null identity,
	email varchar(250) not null,
	password varchar(250) not null,
	created_on datetime not null,
	first_name varchar(250) not null,
	last_name varchar(250) not null,
	phone_number int,
	alternative_number int,
	primary key(id)
);
go

create table rel_accounts_roles(
	id int not null identity,
	account_id int,
	role_id int,
	primary key(id),
	foreign key(account_id) references accounts(id),
	foreign key(role_id) references roles(id)
);
go

create table audit(
	id int not null identity,
	code varchar(250) not null,
	account_id int not null,
	date_started_on datetime not null,
	date_dies_on datetime not null,
	primary key(id),
	foreign key(account_id) references accounts(id)
);
go