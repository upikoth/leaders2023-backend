CREATE TABLE scores (
	id serial not null primary key,
	user_id integer not null,
	creative_space_id integer not null,
	booking_id integer not null,
	comment varchar(1000),
	rating integer not null,

	constraint fk_user_id
		foreign key(user_id) 
			references users(id),

	constraint fk_creative_space_id
		foreign key(creative_space_id) 
			references creative_spaces(id),

	constraint fk_booking_id
		foreign key(booking_id) 
			references bookings(id)
);
