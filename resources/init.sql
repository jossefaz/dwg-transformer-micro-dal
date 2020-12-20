create table if not exists CAD_check_errors
(
    check_status_id int not null
        primary key,
    error_code      int not null
);

create table if not exists CAD_check_status
(
    ID          int auto_increment
        primary key,
    status_code int          null,
    last_update timestamp    null,
    path        varchar(265) null,
    ref_num     int          null,
    system_code int          null
);

