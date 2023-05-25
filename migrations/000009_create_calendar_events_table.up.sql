CREATE TABLE calendar_events (
	id serial not null primary key,
	date varchar(10) not null,
	creative_space_id integer not null
);
