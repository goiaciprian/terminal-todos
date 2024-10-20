create table todos(
	id integer primary key autoincrement not null,
	title text not null,
	description text,
	created_at timestamp default(datetime('now', 'localtime')) not null,
	completed integer default (0),
	completed_at timestamp
);