-- +migrate Up
create extension pgcrypto;

create table base (
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp
);

create table templates (
	id serial not null primary key,
	json_schema json not null
) inherits (base);

create table forms (
	id serial not null primary key,
	template_id int null references templates(id) on delete cascade,
	form_data json not null
) inherits (base);

-- +migrate Down
drop table forms;
drop table templates;
drop table base;
drop extension pgcrypto;