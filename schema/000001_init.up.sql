CREATE TABLE users
(
	user_id SERIAL PRIMARY KEY,
	name text NOT NULL,
	username text NOT NULL,
	password_hash text NOT NULL
);

CREATE TABLE links
(
	link_id SERIAL PRIMARY KEY,
	owner_id int REFERENCES users(user_id) ON DELETE CASCADE NOT NULL,
	original_URL text NOT NULL,
	creation_date timestamp NOT NULL,
	update_date timestamp NOT NULL,
	visits_count int NOT NULL
);

