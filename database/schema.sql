
set storage_engine=INNODB;

drop database orderchef;
create database orderchef;
use orderchef;

-- Printer

create table printer (
	id int auto_increment not null,
	name varchar(255) not null,
	location varchar(255) null,

	primary key (id)
);

-- Config

create table config (
	version varchar(10) not null,
	setup tinyint(1) not null default 0,
	client_id varchar(255) null,
	api_key varchar(255) null
);

create table config__receipt (
	printer_id int null,
	receipt text not null,

	foreign key (printer_id) references printer (id)
);
create index config__receipt_printer_idx on config__receipt (printer_id) using hash;

create table config__modifier_group (
	id int auto_increment not null,
	name varchar(255) not null,
	number_required tinyint(10) not null default 0,

	primary key (id)
);

create table config__modifier (
	id int auto_increment not null,
	group_id int not null,
	name varchar(255) not null,
	price float(4, 2) not null default 0.0,

	primary key (id),
	foreign key (group_id) references config__modifier_group (id)
);
create index config__modifier_group_idx on config__modifier (group_id) using hash;

create table config__order_type (
	id int auto_increment not null,
	name varchar(255) not null,

	primary key (id)
);

create table config__table_type (
	id int auto_increment not null,
	name varchar(255) not null,

	primary key (id)
);

-- Table

create table table__items (
	id int auto_increment not null,
	name varchar(255) not null,
	type_id int not null,
	table_number varchar(255) null,
	location varchar(255) null,

	primary key (id),
	foreign key (type_id) references config__table_type (id)
);
create index table_type_idx on table__items (type_id) using hash;

create table category (
	id int auto_increment not null,
	name varchar(255) not null,
	description text null,

	primary key (id)
);

create table category_printer (
	printer_id int null,
	category_id int null,

	foreign key (printer_id) references printer (id) on delete set null,
	foreign key (category_id) references category (id) on delete set null
);
create index category_printer_printer_idx on category_printer (printer_id) using hash;
create index category_printer_category_idx on category_printer (category_id) using hash;

-- Users

create table customer (
	id int auto_increment not null,
	name varchar(255) not null,
	email varchar(255) null,
	telephone varchar(255) null,
	postcode varchar(255) null,

	primary key (id)
);
create index customer_name_idx on customer (name) using btree;

create table employee (
	id int auto_increment not null,
	name varchar(255) not null,
	manager tinyint(1) not null default 0,
	passkey varchar(4) default "0000",
	last_login int(11) null,

	primary key (id)
);
create index employee_passkey_idx on employee (passkey) using hash;

-- Item

create table item (
	id int auto_increment not null,
	name varchar(255) not null,
	description text null,
	price float(4, 2) not null default 0.0,
	category_id int not null,

	primary key (id),
	foreign key (category_id) references category (id)
);
create index item_name_idx on item (name) using btree;
create index item_category_idx on item (category_id) using hash;

-- Orders

create table order__group (
	id int auto_increment not null,
	table_id int not null,
	cleared tinyint(1) not null default 0,
	cleared_when int(11) null,

	primary key (id),
	foreign key (table_id) references table__items (id)
);
create index order__group_table_idx on order__group (table_id) using hash;

create table order__group_member (
	id int auto_increment not null,
	type_id int null,
	group_id int not null,

	primary key (id),
	foreign key (group_id) references order__group (id),
	foreign key (type_id) references config__order_type (id) on delete set null
);
create index order_group_member_idx on order__group_member (group_id) using hash;

create table order__item (
	id int auto_increment not null,
	item_id int not null,
	order_id int not null,

	quantity tinyint(10) not null default 1,
	notes text null,

	primary key (id),
	foreign key (item_id) references item (id),
	foreign key (order_id) references order__group_member (id)
);
create index order__item_item_idx on order__item (item_id) using hash;
create index order__item_order_idx on order__item (order_id) using hash;

create table order__item_modifier (
	order_item_id int not null,
	modifier_id int not null,

	foreign key (order_item_id) references order__item (id),
	foreign key (modifier_id) references config__modifier (id)
);
create index order__item_modifier_order_idx on order__item_modifier (order_item_id) using hash;
create index order__item_modifier_modifier_idx on order__item_modifier (modifier_id) using hash;

-- Receipt

create table order__receipt (
	id int auto_increment not null,
	order_group_id int not null,
	printer_id int not null,
	employee_id int not null,

	primary key (id),
	foreign key (order_group_id) references order__group (id),
	foreign key (printer_id) references printer (id),
	foreign key (employee_id) references employee (id)
);
