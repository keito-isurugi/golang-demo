CREATE TABLE `golang_demo`.`players` (
  `id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`));

	insert into players (id, name) values(1, "yamada");
	
	select * from players