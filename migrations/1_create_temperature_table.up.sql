CREATE TABLE temperature
(
    `id`         INT(11) NOT NULL AUTO_INCREMENT,
    `city_id`    INT(11) NOT NULL,
    `min`        INT(8)  NOT NULL,
    `max`        INT(8)  NOT NULL,
    `timestamp`  INT(11),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`city_id`) REFERENCES city (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT charset = utf8;