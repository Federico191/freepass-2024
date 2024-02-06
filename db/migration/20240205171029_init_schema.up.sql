create table if not exists posts
(
    ID          int auto_increment
        primary key,
    description text         null,
    pic_url     varchar(255) null
);

create table if not exists users
(
    username   varchar(20)                          not null
        primary key,
    email      varchar(50)                          not null,
    full_name  varchar(255)                         not null,
    password   varchar(60)                          not null,
    is_admin   tinyint(1) default 0                 null,
    created_at timestamp  default CURRENT_TIMESTAMP null,
    updated_at timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint user_uk
        unique (email)
);

create table if not exists accounts
(
    ID                    int auto_increment
        primary key,
    avatar                varchar(255)                         null,
    username              varchar(20)                          not null,
    birth_date            date                                 not null,
    is_voted              tinyint(1) default 0                 null,
    voted_election_number int                                  null,
    created_at            timestamp  default CURRENT_TIMESTAMP null,
    updated_at            timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    deleted_at            timestamp                            null,
    constraint accounts_users_username_fk
        foreign key (username) references users (username)
);

create table if not exists candidates
(
    ID              int auto_increment
        primary key,
    election_number int                                 null,
    account_id      int                                 not null,
    leader_id       int                                 null,
    created_at      timestamp default CURRENT_TIMESTAMP null,
    updated_at      timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    deleted_at      timestamp                           null,
    constraint candidates_uk
        unique (account_id, election_number),
    constraint candidates_accounts_ID_fk
        foreign key (account_id) references accounts (ID)
);

create table if not exists account_comments
(
    post_id      int                                 not null,
    account_id   int                                 not null,
    candidate_id int                                 not null,
    comment      varchar(255)                        not null,
    created_at   timestamp default CURRENT_TIMESTAMP null,
    updated_at   timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    deleted_at   int                                 null,
    primary key (post_id, account_id, candidate_id),
    constraint account_comments_accounts_ID_fk
        foreign key (account_id) references accounts (ID),
    constraint account_comments_candidates_ID_fk
        foreign key (candidate_id) references candidates (ID),
    constraint account_comments_posts_ID_fk
        foreign key (post_id) references posts (ID)
);

create table if not exists candidate_posts
(
    post_id      int                                 not null,
    candidate_id int                                 not null,
    created_at   timestamp default CURRENT_TIMESTAMP null,
    updated_at   timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    deleted_at   timestamp                           null,
    primary key (post_id, candidate_id),
    constraint candidate_posts_candidates_ID_fk
        foreign key (candidate_id) references candidates (ID),
    constraint candidate_posts_posts_ID_fk
        foreign key (post_id) references posts (ID)
);

create table if not exists election_periods
(
    ID             int         not null
        primary key,
    start_time     timestamp   not null,
    end_time       timestamp   not null,
    admin_username varchar(20) null,
    constraint election_period_users_username_fk
        foreign key (admin_username) references users (username)
);

alter table candidates
    add constraint candidate_vice_fk
        foreign key (leader_id) references candidates (ID);

ALTER TABLE candidates
    ADD INDEX idx_election_number (election_number);

alter table accounts
    add constraint accounts_candidates_fk
        foreign key (voted_election_number) references candidates (election_number);