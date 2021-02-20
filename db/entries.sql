BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `entries` (
	`id`	INT AUTO_INCREMENT,
	`word_zh_cn`	TEXT,
	`word_zh_tw`	TEXT,
	`word_en`	TEXT,
	`word_de`	TEXT,
	`created`	DATETIME,
	PRIMARY KEY(`id`)
);
COMMIT;
