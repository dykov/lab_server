create user meme_admin with
  login
  createdb
  inherit
  encrypted password '12345' ;

create database meme_database with
  owner = meme_admin ;