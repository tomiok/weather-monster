CREATE TABLE webhook
(
    `id`           INT(11)      NOT NULL AUTO_INCREMENT,
    `city_id`      INT(11)      NOT NULL,
    `callback_url` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`city_id`) REFERENCES city (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT charset = utf8;