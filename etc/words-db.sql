CREATE SCHEMA `wordsdb`;

CREATE TABLE `wordsdb`.`words` (
  `word_id` integer PRIMARY KEY AUTO_INCREMENT,
  `word` varchar(255) NOT NULL,
  `part_of_speech` varchar(255) NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `wordsdb`.`definitions` (
  `definition_id` integer PRIMARY KEY AUTO_INCREMENT,
  `word_id` integer,
  `definition` varchar(255) NOT NULL,
  `example_sentence` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `wordsdb`.`synonyms` (
  `synonym_id` integer PRIMARY KEY AUTO_INCREMENT,
  `word_id` integer,
  `synonym` varchar(255) NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `wordsdb`.`definitions` ADD FOREIGN KEY (`word_id`) REFERENCES `wordsdb`.`words` (`word_id`);

ALTER TABLE `wordsdb`.`synonyms` ADD FOREIGN KEY (`word_id`) REFERENCES `wordsdb`.`words` (`word_id`);
