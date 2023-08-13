CREATE TABLE `books` (
  `bookId` int NOT NULL AUTO_INCREMENT,
  `bookName` varchar(255) NOT NULL,
  `publisher` varchar(255) NOT NULL,
  `isbn` varchar(255) NOT NULL,
  `edition` int NOT NULL,
  `quantity` int NOT NULL,
  `requestId` int NOT NULL DEFAULT '0',
  `userId` int NOT NULL DEFAULT '0',
  `issuedQuantity` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`bookId`)
)

CREATE TABLE `request` (
  `requestId` int NOT NULL AUTO_INCREMENT,
  `bookId` int NOT NULL,
  `userId` int NOT NULL,
  `status` varchar(255) NOT NULL,
  `temp_col` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`requestId`)
)

CREATE TABLE `user` (
  `userId` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `hash` char(60) NOT NULL,
  `adminId` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`userId`)
)