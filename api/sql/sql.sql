CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int AUTO_INCREMENT primary key,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL unique,
    email varchar(50) NOT NULL unique,
    senha varchar(100) NOT NULL,
    criadoEm timestamp default current_timestamp()
) ENGINE=InnoDB;

CREATE TABLE seguidores(
    usuario_id int NOT NULL,
    FOREIGN KEY (usuario_id) 
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int NOT NULL,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY (usuario_id, seguidor_id)
)ENGINE=InnoDB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) NOT NULL,
    conteudo varchar(300) NOT NULL,
    autor_id int not null, 
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp()
)ENGINE=InnoDB;