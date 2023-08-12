-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: localhost    Database: lms
-- ------------------------------------------------------
-- Server version	8.0.33-0ubuntu0.20.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (40,'wertyujk','xcvb','fghnm',1,0,0,0,0),(42,'asdfg','dfghj','sdfg',1,0,0,0,1),(43,'dfghj','cvbn','fvbn',2,2,0,0,0),(45,'werty','dfgh','qwert',1,0,0,0,1),(46,'dfghj','dfghj','xcvb',3,14,0,0,0),(47,'wedf','3','sdf',3,6,0,0,0),(48,'fhjfyh','vgh','hgjh',5,15,0,0,0),(49,'asdf','rtyhjk','dcvbn',4,4,0,0,0),(50,'qwertyui','cvbn','dfghj',4,0,0,0,2),(51,'somyaqwertyuj','cvb','dfgh',1,1,0,0,0),(52,'asd','asfd','asd',3,0,0,0,0),(53,'asdthth','asfdgfh','asdhg5',35,35,0,0,0),(54,'aer','dgf','df',4,3,0,0,0),(55,'sdf','fv','fsd',3,2,0,0,0),(56,'sdf','sdf','fsd',3,3,0,0,0),(57,'g','gf','gf',5,5,0,0,0),(58,'gqwerqwer','gfqwrqw','gfqwrqwr',5,5,0,0,0),(59,'asd','asd','asd',3,0,0,0,0),(60,'sdff','asdf','dsf',4,45,0,0,0),(61,'check','check','check',4,0,0,0,0),(62,'cheee','cheee','cheee',5,0,0,0,0),(63,'blackbird','bb','bbb',5,5,0,0,0),(64,'','','',0,0,0,0,0),(65,'demo','demo','978-0-12345-678-9',1,2,0,0,0);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `request`
--

DROP TABLE IF EXISTS `request`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `request` (
  `requestId` int NOT NULL AUTO_INCREMENT,
  `bookId` int NOT NULL,
  `userId` int NOT NULL,
  `status` varchar(255) NOT NULL,
  `temp_col` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`requestId`)
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `request`
--

LOCK TABLES `request` WRITE;
/*!40000 ALTER TABLE `request` DISABLE KEYS */;
INSERT INTO `request` VALUES (100,37,81,'owned',0),(106,58,73,'owned',0),(109,57,73,'owned',0),(110,54,73,'owned',0),(111,40,73,'owned',0),(112,53,73,'owned',0),(113,49,73,'owned',0),(114,46,73,'owned',0),(116,38,73,'owned',0),(117,37,73,'owned',0),(118,61,74,'owned',0),(119,62,74,'issue requested',0),(121,45,75,'return requested',0),(124,50,73,'owned',0),(125,50,87,'owned',0),(128,51,73,'issue requested',0),(129,40,92,'issue requested',0),(130,42,92,'return requested',0);
/*!40000 ALTER TABLE `request` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (1,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `userId` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `hash` char(60) NOT NULL,
  `adminId` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (73,'somya','somya@gmail.com','$2a$10$nZ3O3hmbvIHcxfBqsaF..ei6IjGDbtSoOoXEAFvF5QabEpxa3yBSS',0),(75,'arpita','arpita@gmail.com','$2a$10$83MlzWR8GqkDmwITl2I3her9awSJb9OOCgHKaU5lSPg9L3dJUd38W',0),(77,'ishan','ishan@gmail.com','$2a$10$7b7qS/4BBsO50Vd674zC1uerNw5iSSygn9dMRkQbXY8S6fbw4vyb6',0),(78,'altoids','al@gmail.com','$2a$10$BygcyrORKRtj9tfZKT8yL.t7a.IdShUGrc9soBofap4G0kzwqWj2W',1),(81,'ashish','ashish@gmail.com','$2a$10$DR71NgtWW9FcOi6e5yEpPOruSfSFxb/y2UHwAkq6/UkUmUAo9l4qS',0),(83,'qwer','qwer@gmail.com','$2a$10$M3rR0M0h2mAV2ZAUZFs6Q.Nd6qYgdvtjGrE.0CLcunUw4Tf5W40NS',1),(84,'ashpect','ash@gmail.com','$2a$10$uBcjS6z2AgCdo5MbfXOV5O6ypYegZs4B9T9toA6PY.HFQTb1GEnLW',0),(86,'','','$2a$10$KKTfT8Jl37nM6LaDl.jkpeR4O6/b/rUrAk/7HA9AbyvgcYZVnJtf2',0),(87,'bbaaa','bb@gmail.com','$2a$10$hn1JiL4olTyh.I21UJHQK.3eHROGF6gteM9ZcXdm30gTbEh.auWAm',0),(88,'admin','admin@gmail.com','$2a$10$S87eMhzmuo9poXIDRCw5JeYssj0JOzSh7Rtl/it4IMjKScWHTjRq.',1),(89,'Lakshay','lakshay@gmail.com','$2a$10$DDYUc3d.K9fpMm/uhSIRAOxDhX4sFS55MVFF4wh6eOE4FzZRSGnia',0),(90,'Manan','manan@gmail.com','$2a$10$anwQ5j53iZ98GyOTAl4UG.9QKFJdjfubSWUavWtvLge18/x1uZkt6',1),(91,'qwertyuio','sdfg@gmail.com','$2a$10$p0c8SpHWE81vkPyfeQCIs.YP1IGR7mvEpi/0SFyI813tO0waXAQUe',0),(92,'bhagat','bhagat@gmail.com','$2a$10$gnyqIS0mR3rNmYi6a66Pmudx36eWsDmHH3BGmZDRkYvsArZlhR1qq',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-12 22:09:40
