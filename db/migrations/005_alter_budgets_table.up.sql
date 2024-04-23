ALTER TABLE `budgets` 
ADD COLUMN `currency_code` CHAR(3) NOT NULL DEFAULT 'BGN' AFTER `description`, 
RENAME TO  `accounts` ;
