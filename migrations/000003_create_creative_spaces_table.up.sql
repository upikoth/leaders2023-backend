CREATE TABLE creative_spaces (
	id serial not null primary key,
	landlord_id integer not null,
	photos varchar(255)[] not null,
	price_per_hour integer not null,
	latitude real not null,
	longitude real not null,
	working_hours_start_at varchar(10) not null,
	working_hours_end_at varchar(10) not null,
	description varchar(1000) not null,

	constraint fk_landlord_id
		foreign key(landlord_id) 
			references users(id)
);
