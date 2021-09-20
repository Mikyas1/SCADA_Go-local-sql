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
    process_id int null,
    count int not null,
    cyl_type int null,
    constraint QWeekly_5_pk
        primary key (id)
);

-- =====================================================

select process_id, SUM(count) AS count, DATE_FORMAT(process_time, '%Y-%m-%d') AS period, namem from QWeekly_1 WHERE process_time >= '2021-09-16 00:05:00' AND process_time <= '2021-09-18 00:00:00' AND cyl_type IN (1,2) GROUP BY process_id, DATE_FORMAT(process_time, '%Y-%m-%d') ORDER BY DATE_FORMAT(process_time, '%Y-%m-%d'), process_id;

-- ====================================================
create table QSearch_1
(
    id int auto_increment,
    count int not null,
    sort_out int not null,
    plat_input int not null,
    process_id int not null,
    namem varchar(50) not null,
    count_by_machine_1 int not null,
    count_by_machine_2 int not null,
    count_by_machine_3 int not null,
    process_time datetime not null,
    cyl_type int null,
    constraint QSearch_1_pk
        primary key (id)
);

-- ====================================================

select process_id, SUM(count) as count, SUM(plat_input) as plat_input, SUM(sort_out) as sort_out, SUM(count_by_machine_1) as cbm1, SUM(count_by_machine_2) as cbm2, SUM(count_by_machine_3) as cbm3, namem from QSearch_1 WHERE process_time >= '2021-09-19 00:00:00' AND process_time <= '2021-09-20 00:00:00' AND cyl_type IN (1,2) GROUP BY process_id;

-- ====================================================