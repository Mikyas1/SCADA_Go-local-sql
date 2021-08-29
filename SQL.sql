create table table_name
(
    count int default 0 not null,
    final_datetime datetime not null,
    id int auto_increment,
    constraint table_name_pk
        primary key (id)
);

create unique index table_name_final_datetime_uindex
	on table_name (final_datetime);
