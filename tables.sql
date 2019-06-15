create table users (
  id bigserial primary key ,
  login text unique not null ,
  password text not null ,
  photo text
) ;

create table memes (
  id bigserial primary key ,
  image text unique not null ,
  tags text default '',
  owner bigint references users(id) ,
  likes bigint default 0,
  date timestamp default current_timestamp
) ;

create table saved_memes (
  users bigint references users(id) ,
  memes bigint references memes(id)
) ;