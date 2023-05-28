ALTER TABLE creative_spaces
ADD space_type varchar(50);

UPDATE creative_spaces
SET space_type = '';

ALTER TABLE creative_spaces
ALTER COLUMN space_type SET NOT NULL;

ALTER TABLE creative_spaces
ADD area integer;

UPDATE creative_spaces
SET area = 1;

ALTER TABLE creative_spaces
ALTER COLUMN area SET NOT NULL;

ALTER TABLE creative_spaces
ADD capacity integer;

UPDATE creative_spaces
SET capacity = 1;

ALTER TABLE creative_spaces
ALTER COLUMN capacity SET NOT NULL;
