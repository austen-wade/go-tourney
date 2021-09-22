CREATE DATABASE tourneys;

CREATE TABLE tourney_events (
  event_id serial PRIMARY KEY,
  name VARCHAR (255) NOT NULL,
  date DATE NOT NULL,
  game_name VARCHAR (255) NOT NULL,
  number_entrants INT NOT NULL
);

CREATE TABLE tourney_entrants (
  entrant_id serial PRIMARY KEY,
  entrant_tag VARCHAR (255) NOT NULL,
  initial_seed INT NOT NULL,
  final_placement INT
);

CREATE TABLE tourney_sets (
  set_id serial PRIMARY KEY,
  entrant1_id INT NOT NULL,
  entrant2_id INT NOT NULL,
  entrant1_result VARCHAR (255) NOT NULL,
  entrant2_result VARCHAR (255) NOT NULL
);

insert into tourney_events (name, date, game_name, number_entrants) values ('Best in Class', CURRENT_DATE, 'Horse', 16);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Damaris', 1);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Gray', 2);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Manon', 3);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Juliana', 4);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Robinet', 5);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Beale', 6);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Ashlan', 7);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Uriah', 8);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Brynn', 9);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Willi', 10);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Meghan', 11);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Lilllie', 12);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Gaston', 13);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Payton', 14);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Tore', 15);
insert into tourney_entrants (entrant_tag, initial_seed) values ('Jessica', 16);
