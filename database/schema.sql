
# to be run after gorp creates all mysql tables.. Run `go run main.go` then run this
# ansible should do this hmm. perhaps.
# upgrade scripts may be necessary for the database as it can't be updated automagically

# table (type_id) -> config table type
alter table table__items add constraint type_id foreign key (`type_id`) references config__table_type(`id`);

# config primary key set to name
alter table config add primary key(name);
