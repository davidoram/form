-- +migrate Up
create extension pgcrypto;

create table base (
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp
);

create table templates (
	id serial not null primary key,
	external_id uuid not null,
	version int not null,
	json_schema json not null,
	unique (external_id, version)
) inherits (base);

create table forms (
	id serial not null primary key,
	external_id uuid not null,
	template_id int null references templates(id) on delete cascade,
	form_data json not null,
	unique (external_id)
) inherits (base);

-- +migrate Down
drop table forms;
drop table templates;
drop table base;
drop extension pgcrypto;