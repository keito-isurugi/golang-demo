CREATE TABLE `golang_demo`.`users` (
  `id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `age` INT NULL,
  PRIMARY KEY (`id`));

	insert into users (id, name, age) values(1, "yamada", 30);
	
	select * from users;
  