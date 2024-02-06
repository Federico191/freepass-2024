create table if not exists candidate_informations
(
    candidate_id int primary key,
    vision       varchar(255)                        null,
    mission      varchar(255)                        null,
    achievement  varchar(255)                        null,
    experience   text                                null,
    created_at   timestamp default CURRENT_TIMESTAMP null,
    updated_at   timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    deleted_at   timestamp                           null,
    constraint candidates_informations_candidates_ID_fk
        foreign key (candidate_id) references candidates (ID)
);