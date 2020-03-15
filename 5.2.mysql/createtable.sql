CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL DEFAULT '',
    `department` VARCHAR(64) NOT NULL DEFAULT '',
    `created` DATETIME NOT NULL DEFAULT '1970-01-01',
    PRIMARY KEY (`uid`)
);