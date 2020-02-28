-- MySQL dump 10.13  Distrib 5.7.29, for Linux (x86_64)
--
-- Host: localhost    Database: pseletivo
-- ------------------------------------------------------
-- Server version	5.7.29-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `candidatos`
--

DROP TABLE IF EXISTS `candidatos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `candidatos` (
  `nome` varchar(80) DEFAULT NULL,
  `dre` int(11) NOT NULL,
  `curso` varchar(80) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  `data_nasc` date DEFAULT NULL,
  `telegram` int(16) DEFAULT NULL,
  `tem_note` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`dre`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `candidatos`
--

LOCK TABLES `candidatos` WRITE;
/*!40000 ALTER TABLE `candidatos` DISABLE KEYS */;
/*!40000 ALTER TABLE `candidatos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ouvinte`
--

DROP TABLE IF EXISTS `ouvinte`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ouvinte` (
  `nome` varchar(80) DEFAULT NULL,
  `est_ufrj` tinyint(1) DEFAULT NULL,
  `dre` int(16) NOT NULL,
  `contato_1` varchar(16) DEFAULT NULL,
  `contato_2` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`dre`),
  CONSTRAINT `ouvinte_ibfk_1` FOREIGN KEY (`dre`) REFERENCES `candidatos` (`dre`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ouvinte`
--

LOCK TABLES `ouvinte` WRITE;
/*!40000 ALTER TABLE `ouvinte` DISABLE KEYS */;
/*!40000 ALTER TABLE `ouvinte` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `palestrantes`
--

DROP TABLE IF EXISTS `palestrantes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `palestrantes` (
  `nome` varchar(80) DEFAULT NULL,
  `formacao` varchar(255) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `ID_palestrante` varchar(20) NOT NULL,
  `tema` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID_palestrante`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `palestrantes`
--

LOCK TABLES `palestrantes` WRITE;
/*!40000 ALTER TABLE `palestrantes` DISABLE KEYS */;
/*!40000 ALTER TABLE `palestrantes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `tema` varchar(100) NOT NULL,
  `atividade` varchar(100) DEFAULT NULL,
  `ID_palestrante` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`tema`),
  KEY `ID_palestrante` (`ID_palestrante`),
  CONSTRAINT `tags_ibfk_1` FOREIGN KEY (`ID_palestrante`) REFERENCES `palestrantes` (`ID_palestrante`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-28 13:47:45
