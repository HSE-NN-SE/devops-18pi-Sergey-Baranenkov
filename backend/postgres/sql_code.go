package postgres

var UserTable = `create table if not exists users (
    /*secret_info*/
    user_id bigserial primary key,
	email text not null,
	token bytea not null,
	first_name text not null check ( length(first_name) > 0 ),
	last_name text not null check ( length(last_name) > 0),
    sex char(1) not null check ( sex in ('M', 'F')),
    /*basic_info*/
    avatar_ref text default 'hash_path/def_avatar.jpg', /*исправить путь*/
    bg_ref text default 'hash_path/def_bg',
	tel decimal(20),
	city text default '',
	birthday date, /*обработать*/
	status int2 check (0 <= status and status <= 5) default 0, /*0-не указано 1 1-женат 2-не женат 3-влюблен 4-все сложно 5-в акт.поиске */

    /*hobbies*/
	hobby text default '',
	fav_music text default '',
	fav_films text default '',
	fav_books text default '',
	fav_games text default '',
	other_interests text default '',

	/*privacy*/
	who_can_message text not null default 'all' check(who_can_message in ('all','fo')), /*fo - friends only*/
	who_can_see_info text not null default 'all' check(who_can_see_info in ('all','fo')),

    /*edu_and_emp*/
    edu_and_emp_info jsonb
); create index if not exists user_id_idx on users (user_id);`

var PostsTable = `create table if not exists users (
    /*secret_info*/
    user_id bigserial primary key,
	email text not null,
	token bytea not null,
	first_name text not null check ( length(first_name) > 0 ),
	last_name text not null check ( length(last_name) > 0),
    sex char(1) not null check ( sex in ('M', 'F')),
    /*basic_info*/
    avatar_ref text default 'hash_path/def_avatar.jpg', /*исправить путь*/
    bg_ref text default 'hash_path/def_bg',
	tel decimal(20),
	city text default '',
	birthday date, /*обработать*/
	status int2 check (0 <= status and status <= 5) default 0, /*0-не указано 1 1-женат 2-не женат 3-влюблен 4-все сложно 5-в акт.поиске */

    /*hobbies*/
	hobby text default '',
	fav_music text default '',
	fav_films text default '',
	fav_books text default '',
	fav_games text default '',
	other_interests text default '',

	/*privacy*/
	who_can_message text not null default 'all' check(who_can_message in ('all','fo')), /*fo - friends only*/
	who_can_see_info text not null default 'all' check(who_can_see_info in ('all','fo')),

    /*edu_and_emp*/
    edu_and_emp_info jsonb
); create index if not exists user_id_idx on users (user_id);`

var CommentsTable = `create table if not exists comments(
    obj_id bigserial primary key,
    auth_id bigserial references users (user_id) on delete cascade,
    text text not null,

    num_likes integer not null default 0,
    path ltree not null /*сюда записывается rel.path(.comment_id)*/
); create index if not exists tree_path_idx on comments using gist (path);

DROP TRIGGER IF EXISTS insert_comment_trigger on public.comments;
DROP function if exists insert_comment_process;
CREATE FUNCTION insert_comment_process() RETURNS trigger AS $insert_comment_process$
    BEGIN
       	new.path = new.path || new.obj_id::text;
       	Return new;
    END;
$insert_comment_process$ LANGUAGE plpgsql;

CREATE TRIGGER insert_comment_trigger
	BEFORE INSERT ON comments
FOR EACH ROW EXECUTE PROCEDURE insert_comment_process ();`

