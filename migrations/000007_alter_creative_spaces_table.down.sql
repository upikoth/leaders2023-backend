ALTER TABLE creative_spaces
ADD working_hours_start_at varchar(10);

UPDATE creative_spaces
SET working_hours_start_at = '';

ALTER TABLE creative_spaces
ALTER COLUMN working_hours_start_at SET NOT NULL;

ALTER TABLE creative_spaces
ADD working_hours_end_at varchar(10);

UPDATE creative_spaces
SET working_hours_end_at = '';

ALTER TABLE creative_spaces
ALTER COLUMN working_hours_end_at SET NOT NULL;
