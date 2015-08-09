
# to be run after gorp creates all mysql tables.. Run `go run main.go` then run this
# ansible should do this hmm. perhaps.
# upgrade scripts may be necessary for the database as it can't be updated automagically

# table (type_id) -> config table type
alter table table__items add constraint type_id foreign key (`type_id`) references config__table_type(`id`);

# config primary key set to name
alter table config add primary key(name);

# link category_id to item
alter table item add constraint category_id foreign key (`category_id`) references category(`id`);

# add index to deleted
alter table config__modifier add index (`deleted`);
alter table config__modifier add index (`group_id`);
alter table config__modifier_group add index (`deleted`);

create table category_printer (
	`printer_id` varchar(255) default null,
	`category_id` int(11) default null,
	`item_id` int(11) default null,
	unique unique_index (`printer_id`, `category_id`),
	unique unique_index_2 (`printer_id`, `item_id`)
) engine=InnoDB default charset=utf8;
