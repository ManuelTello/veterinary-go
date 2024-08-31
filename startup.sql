create database [veterinary]
go

use [veterinary]
go

create table roles(
	id int not null identity,
	name varchar(250),
	description varchar(250),
	primary key(id)
);
go

if not exists(select * from [roles] where [roles].id = 0 or [roles].id = 1 or [roles].id = 2) 
insert into [roles](name,description) values
('admin','Veterinary sys admin, all permissions granted'),
('doctor','Veterinary doctor'),
('client','Veterinary client');
go

create table accounts_details(
	id int not null identity,
	first_name varchar(250) not null,
	last_name varchar(250) not null,
	email varchar(250) not null,
	date_created datetime not null,
	phone_number int null,
	primary key(id)
);
go

create table accounts(
	id int not null identity,
	password varchar(250) not null,
	username varchar(250) not null,
	details_id int not null,
	primary key(id),
	foreign key(details_id) references accounts_details(id)
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