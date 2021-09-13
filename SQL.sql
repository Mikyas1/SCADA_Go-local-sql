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

-- ===================================================
create table BDashboard_2
(
    count int default 0 not null,
    final_datetime datetime not null,
    id int auto_increment,
    constraint BDashboard_2_pk
        primary key (id)
);

create unique index BDashboard_2_final_datetime_uindex
    on BDashboard_1 (final_datetime);

-- ===================================================
create table QWeekly_5
(
    id int auto_increment,
    process_time datetime not null,
    namem varchar(50) not null,
    process_id int,
    count int not null,
    cyl_type int,
    constraint QWeekly_5_pk
        primary key (id)
);

create unique index QWeekly_5_process_time_uindex
    on QWeekly_5 (process_time);

