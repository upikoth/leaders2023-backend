ALTER TABLE creative_spaces
DROP COLUMN title;

ALTER TABLE creative_spaces
DROP COLUMN address;

ALTER TABLE creative_spaces
ALTER COLUMN photos type varchar(255)[];
