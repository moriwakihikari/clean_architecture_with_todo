CREATE DATABASE IF NOT EXISTS todo_list;
USE todo_list;

CREATE TABLE IF NOT EXISTS `users`
(
 `id`               INT(20) AUTO_INCREMENT,
 `first_name`       VARCHAR(50) NOT NULL,
 `last_name`        VARCHAR(50) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL, 
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `todos`
(
 `id`               INT(20) AUTO_INCREMENT,
 `title`            VARCHAR(50) NOT NULL,
 `note`             VARCHAR(50) NOT NULL,
 `day_time`         Datetime DEFAULT NULL,
 `user_id`          Int,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL, 
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) 
        REFERENCES users(`id`)
        ON DELETE CASCADE
);
INSERT INTO users (first_name, last_name) VALUES ("Hikari", "Moriwaki");
INSERT INTO todos (title, note, user_id) VALUES ("title1", "note1", 1);