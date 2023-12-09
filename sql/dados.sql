insert into usuarios (nome, nick, email, senha)
values

("Usuário 1", "usuario_1", "usuario1@hotmail.com", "'$2a$10$khz6.9uT3lEQLowjGdmNmummo5oq5gMZz/vCEm9vcqWEsEr8qCitO'"), 
("Usuário 2", "usuario_2", "usuario2@hotmail.com", "'$2a$10$khz6.9uT3lEQLowjGdmNmummo5oq5gMZz/vCEm9vcqWEsEr8qCitO'"), 
("Usuário 3", "usuario_3", "usuario3@hotmail.com", "'$2a$10$khz6.9uT3lEQLowjGdmNmummo5oq5gMZz/vCEm9vcqWEsEr8qCitO'");


insert into seguidores (usuario_id, seguidor_id)
values 
(1, 2),
(3, 1),
(1, 3);