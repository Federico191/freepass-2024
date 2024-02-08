create table if not exists candidate_informations
(
    election_number int primary key,
    vision       varchar(255)                        null,
    mission      varchar(255)                        null,
    achievement  varchar(255)                        null,
    experience   text                                null,
    created_at   timestamp default CURRENT_TIMESTAMP null,
    updated_at   timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    deleted_at   timestamp                           null,
    constraint candidates_informations_election_number_fk
        foreign key (election_number) references candidates (election_number)
);