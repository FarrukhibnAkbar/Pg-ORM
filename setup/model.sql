create table users (
	user_id   serial not null primary key,
	firstname varchar(20) not null,
	lastname  varchar(20) not null,
	create_at timestamptz default current_timestamp
);