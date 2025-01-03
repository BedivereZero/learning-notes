CREATE TABLE IF NOT EXISTS `examples` (
    `id`    INT AUTO_INCREMENT PRIMARY KEY,
    `name`  VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS `scenes` (
    `id`            INT AUTO_INCREMENT PRIMARY KEY,
    `name`          VARCHAR(255),
    `example_id`    INT NOT NULL
);

CREATE TABLE IF NOT EXISTS `scene_resource_bindings` (
    `scene_id`      INT NOT NULL,
    `resource_id`   INT NOT NULL,
    PRIMARY KEY (`scene_id`, `resource_id`)
);

CREATE TABLE IF NOT EXISTS `users` (
    `id`            INT AUTO_INCREMENT PRIMARY KEY,
    `name`          VARCHAR(255)
);
