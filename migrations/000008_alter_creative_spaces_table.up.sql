ALTER TABLE creative_spaces
ADD calendar_link varchar(1000);

ALTER TABLE creative_spaces
ADD calendar_work_day_indexes integer[];

UPDATE creative_spaces
SET calendar_work_day_indexes = '{}';

ALTER TABLE creative_spaces
ALTER COLUMN calendar_work_day_indexes SET NOT NULL;
