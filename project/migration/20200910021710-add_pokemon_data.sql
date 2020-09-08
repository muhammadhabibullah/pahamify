-- +migrate Up
START TRANSACTION;

INSERT INTO `types` (`name`)
VALUES ('Grass')
ON DUPLICATE KEY UPDATE `name`=`name`;

INSERT INTO `pokemons` (`id`, `number`, `name`)
VALUES ('6a876ba0-49d1-4fb6-8e0a-63ddd965d6a8', '001', 'Bulbasaur');

INSERT INTO `pokemon_types` (`pokemon_id`, `type_name`)
VALUES ('6a876ba0-49d1-4fb6-8e0a-63ddd965d6a8', 'Grass')
ON DUPLICATE KEY UPDATE `type_name`=`type_name`;

COMMIT;

-- +migrate Down
