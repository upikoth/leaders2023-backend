ALTER TABLE creative_spaces
ADD title varchar(100);

UPDATE creative_spaces
SET title = '';

ALTER TABLE creative_spaces
ALTER COLUMN title SET NOT NULL;

ALTER TABLE creative_spaces
ADD address varchar(100);

UPDATE creative_spaces
SET address = '';

ALTER TABLE creative_spaces
ALTER COLUMN address SET NOT NULL;

ALTER TABLE creative_spaces
ALTER COLUMN photos type text[];
