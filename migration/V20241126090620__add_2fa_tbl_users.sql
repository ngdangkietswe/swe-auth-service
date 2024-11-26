set search_path to sweauth;

alter table users
    add column enable_2fa boolean not null default false,
    add column secret_2fa varchar(255);