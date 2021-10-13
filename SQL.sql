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
create table BDashboard_3
(
    count int default 0 not null,
    final_datetime datetime not null,
    id int auto_increment,
    constraint BDashboard_3_pk
        primary key (id)
);

create unique index BDashboard_3_final_datetime_uindex
    on BDashboard_3 (final_datetime);

-- ===================================================
create table QWeekly_3
(
    id int auto_increment,
    process_time datetime not null,
    namem varchar(50) not null,
    process_id int null,
    count int not null,
    cyl_type int null,
    constraint QWeekly_3_pk
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

create table QDashboard_1
(
    residual int not null,
    check_net int not null,
    count int not null,
    final_datetime datetime not null
);

create unique index QDashboard_1_final_datetime_uindex
    on QDashboard_1 (final_datetime);


-- =====================================================


create table QReport_3
(
    machine_id varchar(20) not null,
    process_date varchar(255) not null,
    gtem400 int null,
    gtem350 int null,
    gtem300 int null,
    gtem250 int null,
    gtem200 int null,
    gtem150 int null,
    gtem100 int null,
    gtem050 int null,
    value int null,
    gte050 int null,
    gte100 int null,
    gte150 int null,
    gte200 int null,
    gte250 int null,
    gte300 int null,
    gte350 int null,
    gte400 int null,
    sum int null,
    m200x200 int null,
    diff int null,
    start_point float null,
    accuracy float null,
    cylinder_type int null,
    final_datetime datetime null
);

-- ===================================================================================