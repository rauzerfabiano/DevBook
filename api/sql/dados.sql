INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$2Xx90olKhd4oiKgXAU1W2eG5c9OXGj2WQoCvAJpxwZryUyu586n72"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$2Xx90olKhd4oiKgXAU1W2eG5c9OXGj2WQoCvAJpxwZryUyu586n72"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$2Xx90olKhd4oiKgXAU1W2eG5c9OXGj2WQoCvAJpxwZryUyu586n72");

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);


