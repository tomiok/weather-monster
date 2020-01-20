CREATE TABLE city
(
    `id`         INT(11)      NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(100) NOT NULL,
    `latitude`   FLOAT(11) NOT NULL,
    `longitude`  FLOAT(11) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (name),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT charset = utf8;