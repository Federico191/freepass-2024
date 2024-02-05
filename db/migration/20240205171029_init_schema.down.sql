drop table if exists account_comments;

alter table accounts
    drop foreign key accounts_candidates_fk;

drop table if exists candidate_posts;

drop table if exists candidates;

drop table if exists accounts;

drop table if exists election_periods;

drop table if exists posts;

drop table if exists users;