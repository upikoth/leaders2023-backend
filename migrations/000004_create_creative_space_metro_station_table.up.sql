CREATE TABLE creative_space_metro_station (
	metro_station_id integer not null,
	creative_space_id integer not null,
	distance_in_minutes integer not null,

	constraint fk_metro_station
		foreign key(metro_station_id) 
			references metro_stations(id),

	constraint fk_creative_space
		foreign key(creative_space_id) 
			references creative_spaces(id)
);
