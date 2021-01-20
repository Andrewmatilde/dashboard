create database if not exists github_info;

create table comment
(
    id int not null,
    issue_id int not null,
    body longtext not null,
    primary key (id, issue_id)
);

create table pull_request_comment
(
    id int not null,
    pull_request_id int not null,
    body longtext not null,
    primary key (id, pull_request_id)
);

create table issue
(
    id int not null
        primary key,
    number int not null,
    repository_id int not null,
    closed tinyint(1) not null,
    closed_at datetime null,
    closed_week datetime null ,
    created_at datetime not null,
    created_week datetime not null,
    title varchar(1000) not null,
    url varchar(1000) not null,
    constraint issue_number_repository_id_uindex
        unique (number, repository_id)
);

create table issue_team
(
    issue_id int not null,
    team_id int not null,
    primary key (issue_id, team_id)
);

create table issue_version_fixed
(
    issue_id int not null,
    version_id int not null,
    primary key (issue_id,version_id)
);

create table issue_version_affected
(
    issue_id int not null,
    version_id int not null,
    primary key (issue_id,version_id)
);

create table pull_request
(
    id int not null
        primary key,
    number int not null,
    repository_id int not null,
    closed tinyint(1) not null,
    closed_at datetime null,
    closed_week datetime null ,
    created_at datetime not null,
    created_week datetime not null,
    merged tinyint(1) not null,
    merged_at datetime null,
    merged_week datetime null,
    title varchar(1000) not null,
    url varchar(1000) not null,
    constraint pull_request_number_repository_id_uindex
        unique (number, repository_id)
);

create table pull_request_team
(
    pull_request_id int not null,
    team_id int not null,
    primary key (pull_request_id, team_id)
);

create table label
(
    id int auto_increment
        primary key,
    name varchar(100) not null,
    repository_id int not null,
    constraint label_name_uindex
        unique (id, name)
);

create table issue_label
(
    issue_id int not null,
    label_id int not null,
    primary key (issue_id, label_id)
);

create table pull_request_label
(
    pull_request_id int not null,
    label_id int not null,
    primary key (pull_request_id, label_id)
);

create table label_severity_weight
(
    label_id int not null,
    label_name varchar(100) not null,
    weight float not null,
    primary key (label_id)
);

create table repository
(
    id int not null
        primary key,
    owner varchar(100) not null,
    repo_name varchar(100) not null,
    url varchar(1000) not null
);

create table timeline_repository
(
    datetime datetime not null,
    repository_id int not null,
    constraint datetime
        unique (datetime, repository_id)
);

create table label_sig
(
    label_id int not null,
    label_name varchar(100) not null,
    primary key (label_id)
);


create table if not exists team
(
    id int auto_increment
        primary key,
    name varchar(40) not null,
    size int null
);

create table if not exists team_bug_jail
(
    time datetime null,
    team_id int null,
    DI float null,
    in_jail tinyint(1) not null
);

create table if not exists team_issue
(
    issue_id int not null,
    team_id int not null,
    primary key (issue_id, team_id)
);

create table timeline
(
    datetime datetime not null,
    constraint datetime
        unique (datetime)
);

create table week_line
(
    week datetime not null unique
);

create table user
(
    id int not null
        primary key,
    login varchar(100) not null,
    email varchar(100) not null,
    constraint user_login_name_uindex
        unique (login)
);

create table user_issue
(
    user_id int not null,
    issue_id int not null,
    primary key (user_id, issue_id)
);

create table user_bug_jail
(
    time datetime not null,
    user_id int not null,
    in_jail tinyint(1) null,
    critical int null,
    di float null,
    primary key (user_id, time)
);

create table user_pull_request
(
    user_id int not null,
    pull_request_id int not null,
    primary key (user_id, pull_request_id)
);

create table version
(
    id int not null
        primary key,
    major int not null,
    minor int not null,
    patch int not null
);

create table tag
(
    id int auto_increment
        primary key,
    name varchar(100) not null,
    repository_id int not null,
    constraint name unique (name,repository_id)
);

create table coverage_timeline
(
    time datetime not null,
    repo_id int not null,
    coverage float not null
);
