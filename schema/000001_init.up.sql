CREATE TABLE links
(
	link_id serial PRIMARY KEY,
	short_path text  NOT NULL,
	original_URL text NOT NULL,
	creation_date timestamp NOT NULL,
	update_date timestamp NOT NULL,
	visits_count bigint NOT NULL
);

CREATE TABLE users
(
	user_id serial PRIMARY KEY,
	name text NOT NULL,
	username text NOT NULL,
	password_hash text NOT NULL
);

CREATE TABLE users_links
(
	ul_id serial PRIMARY KEY,
	fk_user_id bigint REFERENCES users(user_id) ON DELETE CASCADE NOT NULL,
	fk_link_id bigint REFERENCES links(link_id) ON DELETE CASCADE NOT NULL
);

