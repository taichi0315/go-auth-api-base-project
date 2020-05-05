CREATE SCHEMA IF NOT EXISTS `auth_api` DEFAULT CHARACTER SET utf8;
USE `auth_api`;

CREATE TABLE IF NOT EXISTS `auth_api`.`users` (
  `id`         BIGINT(20)   UNSIGNED NOT NULL AUTO_INCREMENT,
  `username`   VARCHAR(255) NOT NULL UNIQUE,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `auth_api`.`user_auths` (
  `user_id`         BIGINT(20)   UNSIGNED NOT NULL,
  `email`           VARCHAR(255) NOT NULL UNIQUE,
  `hash`            VARCHAR(255) NOT NULL,
  `updated_at`      DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_auths_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `auth_api`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `auth_api`.`auth_tokens` (
  `id`              BIGINT(20)   UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`         BIGINT(20)   UNSIGNED NOT NULL,
  `token`           VARCHAR(255) NOT NULL,
  `expiry`          DATETIME,
  `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_auth_tokens_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `auth_api`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
ENGINE = InnoDB;
