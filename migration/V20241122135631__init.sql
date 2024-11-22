create database "SweAuth";
\connect "SweAuth";
create schema sweauth;
alter database "SweAuth" set search_path to sweauth;
create extension if not exists pgcrypto;