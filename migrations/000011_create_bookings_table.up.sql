CREATE TABLE bookings (
	id serial not null primary key,
	tenant_id integer not null,
	landlord_id integer not null,
	creative_space_id integer not null,
	status varchar(20) not null,
	full_price integer not null,

	constraint fk_tenant_id
		foreign key(tenant_id) 
			references users(id),

	constraint fk_landlord_id
		foreign key(landlord_id) 
			references users(id),

	constraint fk_creative_space_id
		foreign key(creative_space_id) 
			references creative_spaces(id)
);
