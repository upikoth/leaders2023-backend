ALTER TABLE users
DROP COLUMN role;

ALTER TABLE users
DROP COLUMN phone;

ALTER TABLE users
ADD email varchar(255);

UPDATE users
SET email = CONCAT(id, '@mail.ru');

ALTER TABLE users
ALTER COLUMN email SET NOT NULL;

ALTER TABLE users
ADD CONSTRAINT users_email_unique UNIQUE (email);
