ALTER TABLE creative_spaces
ADD status varchar(20);

UPDATE creative_spaces
SET status = 'confirmation';

ALTER TABLE creative_spaces
ALTER COLUMN status SET NOT NULL;
