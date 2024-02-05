drop table if exists account_comments;
drop table if exists account_likes;
drop table if exists candidate_posts;
alter table accounts
    drop constraint accounts_candidates_fk;
drop table if exists candidates;
drop table if exists accounts;
drop table if exists election_periods;
drop table if exists posts;
drop table if exists users;