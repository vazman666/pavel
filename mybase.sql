-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: mybase
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `mybase`
--

DROP TABLE IF EXISTS `mybase`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mybase` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `partnum` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` int(11) NOT NULL,
  `price` float NOT NULL,
  `provider` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `color` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=496 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mybase`
--

LOCK TABLES `mybase` WRITE;
/*!40000 ALTER TABLE `mybase` DISABLE KEYS */;
INSERT INTO `mybase` VALUES (388,'3013281',2,119.68,'froza','втулка стабилизатора','Vlad','22 November 2021','#ffffff'),(417,'4060A711',2,488,'','опорный подшипник','Vlad','02 December 2021','#ffffff'),(442,'107818755',1,260.98,'mikado','мотор омывателя','Pavel','17 December 2021','#ffffff'),(463,'BB5016',1,1401,'forum','колодки','Vlad','20 December 2021','#ffffff'),(464,'-pa268af',1,743.43,'mikado','Колодки','Vlad','20 December 2021','#ffffff'),(467,'243001',2,265.74,'mikado','ремкомплект','Pavel','21 December 2021','#ffffff'),(472,'gaskol-003-kit',1,312,'mikado','Уплотнительное кольцо','Vlad','23 December 2021','#ffffff'),(473,'cb-1802gp-std',1,578.46,'mikado','вкладыши','Vlad','27 December 2021','#ffffff'),(474,'ms-1845a-std',1,665.76,'mikado','вкладыши','Vlad','27 December 2021','#ffffff'),(475,'ah2554-j0',1,282.3,'mikado','сальник','Vlad','27 December 2021','#ffffff'),(476,'AH2057R0',1,144.9,'tiss','сальник','Vlad','23 December 2021','#ffffff'),(477,'1064A041',2,159,'Шате-М','проклакда','Vlad','23 December 2021','#ffffff'),(478,'tw-1845a-std',1,175.07,'mikado','полукольца','Vlad','23 December 2021','#ffffff'),(483,'-f4225',1,98.26,'mikado','сальник','Vlad','27 December 2021','#ffffff'),(484,'-f422',1,95.82,'mikado','сальник','Vlad','27 December 2021','#ffffff'),(485,'md149743',1,395.37,'mikado','втулка','Vlad','28 December 2021','#ffffff'),(486,'kk19017',1,2985.91,'mika','ступица','Anton','28 December 2021','#ffffff'),(488,'kk19017',1,2985.91,'mikado','ступица','Anton','29 December 2021','#ffffff'),(489,'340-91k',1,280.18,'mikado','термостат','Anton','29 December 2021','#ffffff'),(490,'MD377544',1,683.91,'froza','проставка','Vlad','29 December 2021','#ffffff'),(491,'M25901',2,378,'forum','подшипник опоры','Pavel','30 December 2021','#ffffff'),(492,'5033F7',2,446.3,'forum','чаша','Pavel','30 December 2021','#ffffff'),(493,'KSA580STD',1,2295,'forum','амортизатор','Pavel','30 December 2021','#ffffff'),(494,'KSA581STD',1,2295,'forum','амотризатор','Pavel','30 December 2021','#ffffff'),(495,'1775982',1,1222.92,'froza','провод','Vlad','30 December 2021','#ffffff');
/*!40000 ALTER TABLE `mybase` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-31  9:31:22
