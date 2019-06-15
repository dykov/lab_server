create or replace procedure add_user( login text , password text , photo text )
  language sql
  as $$
    insert into users (login, password, photo)
    values (login , password , photo);
  $$ ;

create or replace procedure delete_user( id bigint )
  language sql
  as $$
    delete from users
    where users.id = id ;
  $$ ;

create or replace procedure add_meme( img text , tags text , owner bigint )
  language sql
  as $$
    insert into memes ( image, tags, owner)
    values ( img , tags , owner );
  $$ ;

create or replace procedure delete_meme( id bigint )
  language sql
  as $$
    delete from memes
    where memes.id = id ;
  $$ ;

create or replace procedure like_meme( user_id bigint , meme_id bigint )
  language sql
  as $$
    insert into saved_memes (users, memes)
    values (user_id , meme_id);
    
    update memes
    set likes = likes+1
    where id = meme_id;
  $$