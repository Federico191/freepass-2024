create table account_votes
(
    account_id            int                   not null,
    is_voted              boolean default false null,
    voted_election_number int                   null,
    constraint account_votes_pk
        primary key (account_id),
    constraint account_votes_pk_2
        unique (voted_election_number),
    constraint account_votes_accounts_ID_fk
        foreign key (account_id) references accounts (ID),
    constraint account_votes_candidates_election_number_fk
        foreign key (voted_election_number) references candidates (election_number)
);

