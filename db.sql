CREATE TABLE IF NOT EXISTS `students` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(50) NOT NULL,
    `last_name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `academy` CHAR(4) NOT NULL,

    PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `registrations` (
    `student_id` INT NOT NULL,
    `player_uuid` VARCHAR(36) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(`student_id`) REFERENCES `students`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `verify_intents` (
    `id` VARCHAR(36) NOT NULL DEFAULT UUID(),
    `student_id` INT NOT NULL,
    `player_uuid` VARCHAR(36) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expires_at` TIMESTAMP NOT NULL DEFAULT ADDTIME(CURRENT_TIMESTAMP, "1:00:00"),

    PRIMARY KEY(`id`),
    FOREIGN KEY(`student_id`) REFERENCES `students`(`id`) ON DELETE CASCADE
);