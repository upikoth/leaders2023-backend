ALTER TABLE users
DROP COLUMN email;

ALTER TABLE users
ADD phone varchar(11);

UPDATE users
SET phone = id;

ALTER TABLE users
ALTER COLUMN phone SET NOT NULL;

ALTER TABLE users
ADD CONSTRAINT users_phone_unique UNIQUE (phone);

ALTER TABLE users
ADD role varchar(20);

UPDATE users
SET role = 'admin';

ALTER TABLE users
ALTER COLUMN role SET NOT NULL;
