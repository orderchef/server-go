-- MySQL dump 10.13  Distrib 5.5.44, for debian-linux-gnu (armv7l)
--
-- Host: localhost    Database: orderchef
-- ------------------------------------------------------
-- Server version	5.5.44-0+deb7u1

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
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'Sushi',''),(2,'Rice and Noodles',''),(3,'Classical Side Dishes',''),(4,'Drinks',''),(5,'Ice Cream',''),(6,'Bento Box',''),(7,'Ramen','');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_printer`
--

DROP TABLE IF EXISTS `category_printer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category_printer` (
  `printer_id` varchar(255) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `item_id` int(11) DEFAULT NULL,
  UNIQUE KEY `unique_index` (`printer_id`,`category_id`),
  UNIQUE KEY `unique_index_2` (`printer_id`,`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_printer`
--

LOCK TABLES `category_printer` WRITE;
/*!40000 ALTER TABLE `category_printer` DISABLE KEYS */;
INSERT INTO `category_printer` VALUES ('sushi',1,NULL),('kitchen',2,NULL),('kitchen',3,NULL),('receipt',4,NULL),('receipt',5,NULL),('kitchen',6,NULL),('kitchen',7,NULL);
/*!40000 ALTER TABLE `category_printer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config`
--

DROP TABLE IF EXISTS `config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config` (
  `name` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config`
--

LOCK TABLES `config` WRITE;
/*!40000 ALTER TABLE `config` DISABLE KEYS */;
INSERT INTO `config` VALUES ('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef'),('is_setup','1'),('venue_name','orderchef');
/*!40000 ALTER TABLE `config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config__modifier`
--

DROP TABLE IF EXISTS `config__modifier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config__modifier` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `deleted` (`deleted`),
  KEY `group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config__modifier`
--

LOCK TABLES `config__modifier` WRITE;
/*!40000 ALTER TABLE `config__modifier` DISABLE KEYS */;
/*!40000 ALTER TABLE `config__modifier` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config__modifier_group`
--

DROP TABLE IF EXISTS `config__modifier_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config__modifier_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `choice_required` tinyint(1) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `deleted` (`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config__modifier_group`
--

LOCK TABLES `config__modifier_group` WRITE;
/*!40000 ALTER TABLE `config__modifier_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `config__modifier_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config__order_type`
--

DROP TABLE IF EXISTS `config__order_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config__order_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config__order_type`
--

LOCK TABLES `config__order_type` WRITE;
/*!40000 ALTER TABLE `config__order_type` DISABLE KEYS */;
INSERT INTO `config__order_type` VALUES (1,'Drinks',''),(2,'Starter',''),(3,'Main',''),(4,'Desert','');
/*!40000 ALTER TABLE `config__order_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config__payment_method`
--

DROP TABLE IF EXISTS `config__payment_method`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config__payment_method` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config__payment_method`
--

LOCK TABLES `config__payment_method` WRITE;
/*!40000 ALTER TABLE `config__payment_method` DISABLE KEYS */;
INSERT INTO `config__payment_method` VALUES (1,'Card'),(2,'Cash');
/*!40000 ALTER TABLE `config__payment_method` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config__receipt`
--

DROP TABLE IF EXISTS `config__receipt`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config__receipt` (
  `printer_id` int(11) DEFAULT NULL,
  `receipt` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config__receipt`
--

LOCK TABLES `config__receipt` WRITE;
/*!40000 ALTER TABLE `config__receipt` DISABLE KEYS */;
/*!40000 ALTER TABLE `config__receipt` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `config__table_type`
--

DROP TABLE IF EXISTS `config__table_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config__table_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config__table_type`
--

LOCK TABLES `config__table_type` WRITE;
/*!40000 ALTER TABLE `config__table_type` DISABLE KEYS */;
INSERT INTO `config__table_type` VALUES (1,'BAR'),(2,'Restaurant');
/*!40000 ALTER TABLE `config__table_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `telephone` varchar(255) DEFAULT NULL,
  `postcode` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employee`
--

DROP TABLE IF EXISTS `employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `manager` tinyint(1) DEFAULT NULL,
  `passkey` varchar(255) DEFAULT NULL,
  `last_login` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee`
--

LOCK TABLES `employee` WRITE;
/*!40000 ALTER TABLE `employee` DISABLE KEYS */;
/*!40000 ALTER TABLE `employee` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item`
--

DROP TABLE IF EXISTS `item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item`
--

LOCK TABLES `item` WRITE;
/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES (1,'Edemame','Soy beans in a pod',3,3),(2,'Miso Soup','',2.200000047683716,3),(3,'Goma Wakame','Cold japanese mixed seaweed salad',3.5,3),(4,'Okonomiyaki Skewer','Japanese savoury pancake filled with cabbage and octopus',4.800000190734863,3),(5,'Gyoza Vegetable','',3.799999952316284,3),(6,'Gyoza Chicken','',4.199999809265137,3),(7,'Prawn Tempura','',5.199999809265137,3),(8,'Soft Shell Crab Tempura','',6.949999809265137,3),(9,'Vegetable Tempura','',4.800000190734863,3),(10,'Chicken Katsu','',3.950000047683716,3),(11,'Prawn Katsu','',4.199999809265137,3),(12,'Pumpkin Koroke','',3.299999952316284,3),(13,'Takoyaki','',4.949999809265137,3),(14,'Chilli Squid','',5.800000190734863,3),(15,'Chicken Yakitori','',4.949999809265137,3),(16,'Chicken Karaage','',4.5,3),(17,'Sake','',6.5,3),(18,'Maguro','',7.900000095367432,3),(19,'Amaebi','',7.900000095367432,3),(20,'Assorted Sashimi','',9.899999618530273,3),(21,'Kappa','',2.799999952316284,1),(22,'Abokado','',2.799999952316284,1),(23,'Takuwan','',2.799999952316284,1),(24,'Sake','',3.5999999046325684,1),(25,'Maguro','',4,1),(26,'Ebi','',4,1),(27,'Taberu Roll','',9.949999809265137,1),(28,'Cowley Roll','',8.800000190734863,1),(29,'Eel Dragon Roll','',11.899999618530273,1),(30,'Spicy Tuna Roll','',10.5,1),(31,'California Roll','',6.5,1),(32,'Salmon Skin Roll','',6,1),(33,'Crispy Prawn Roll','',8.5,1),(34,'Spider Roll','',9.25,1),(35,'Chicken Katsu Roll','',7.5,1),(36,'Yasai Vegetarian Roll','',6,1),(37,'Nigiri Sake','',2.5,1),(38,'Nigiri Maguro','',3.5999999046325684,1),(39,'Nigiri Unagi','',4.199999809265137,1),(40,'Nigiri Temago','',2,1),(41,'Nigiri Tako','',2.5,1),(42,'Nigiri Ebi','',2.700000047683716,1),(43,'Nigiri Amaebi','',3.200000047683716,1),(44,'Nigiri Inari','',2,1),(45,'Gunkan Kaiso','',2.5,1),(46,'Gunkan Masago','',3.5,1),(47,'Gunkan Ikura','',4.800000190734863,1),(48,'Gunkan Tobikko','',4,1),(49,'Temaki Salmon & Avocado Hand Roll','',3.5,1),(50,'Temaki Eel & Cucumber','',4.300000190734863,1),(51,'Temaki Soft Shell Crab','',4.300000190734863,1),(52,'Temaki Spicy Tuna & Cucumber','',3.5,1),(53,'Temaki Crabstick & Avocado','',3.5,1),(54,'Temaki Prawn & Avocado','',3.5,1),(55,'Temaki Vegetarian','',3.5,1),(56,'Spicy Seafood Udon','',9.5,2),(57,'Chicken Curry Udon','',7.949999809265137,2),(58,'Chicken Yakisoba','',7.949999809265137,2),(59,'Seafood Yakisoba','',8.949999809265137,2),(60,'Vegetarian Yakisoba','',7.199999809265137,2),(61,'Chicken YakiUdon','',7.949999809265137,2),(62,'Seafood YakiUdon','',8.949999809265137,2),(63,'Vegetarian YakiUdon','',7.199999809265137,2),(64,'Chicken Katsudon','',7.949999809265137,2),(65,'Pork Katsudon','',7.949999809265137,2),(66,'Unagi Don','',14.949999809265137,2),(67,'Beef Guydon','',7.949999809265137,2),(68,'Extra Egg','',0.800000011920929,2),(69,'Chicken Katsu Curry','',7.949999809265137,2),(70,'Pork Katsu Curry','',7.949999809265137,2),(71,'Pumpkin Katsu Curry','',6.800000190734863,2),(72,'Sake Chahan','',7.5,2),(73,'Chicken Teriyaki','',11.5,6),(74,'Salmon Teriyaki','',13.5,6),(75,'Beef Teriyaki','',15.949999809265137,6),(76,'Chicken Katsu','',11.5,6),(77,'Pork Katsu','',11.5,6),(78,'Pumpkin Koroke','',9.5,6),(79,'Tonkotsu Pork','',8.949999809265137,7),(80,'Chicken Ramen','',7.949999809265137,7),(81,'Vegetarian Ramen','',7.5,7),(82,'Coca Cola','',2.3499999046325684,4),(83,'Orange Frobishers Juice','',2.5,4),(84,'Apple Frobishers Juice','',2.5,4),(85,'Mango Frobishers Juice','',2.5,4),(86,'Pineapple Frobishers Juice','',2.5,4),(87,'Lemon MangaJo','',2.5999999046325684,4),(88,'Pomegranate MangaJo','',2.5999999046325684,4),(89,'Original Ramune','',2.5999999046325684,4),(90,'Lychee Ramune','',2.5999999046325684,4),(91,'Melon Ramune','',2.5999999046325684,4),(92,'Strawberry Ramune','',2.5999999046325684,4),(93,'Calpis','',2.200000047683716,4),(94,'Strathmore Water','',1.7999999523162842,4),(95,'Green Tea','',2,4),(96,'Ice Oolong Tea','',2,4),(97,'Ice Green Tea','',2.200000047683716,4),(98,'Asahi','',3.5,4);
/*!40000 ALTER TABLE `item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item__modifier`
--

DROP TABLE IF EXISTS `item__modifier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item__modifier` (
  `item_id` int(11) DEFAULT NULL,
  `modifier_group_id` int(11) DEFAULT NULL,
  UNIQUE KEY `item_id` (`item_id`,`modifier_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item__modifier`
--

LOCK TABLES `item__modifier` WRITE;
/*!40000 ALTER TABLE `item__modifier` DISABLE KEYS */;
/*!40000 ALTER TABLE `item__modifier` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order__bill`
--

DROP TABLE IF EXISTS `order__bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order__bill` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) DEFAULT NULL,
  `paid` tinyint(1) DEFAULT NULL,
  `paid_amount` double DEFAULT NULL,
  `total` double DEFAULT NULL,
  `payment_method_id` int(11) DEFAULT NULL,
  `printed_at` datetime DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order__bill`
--

LOCK TABLES `order__bill` WRITE;
/*!40000 ALTER TABLE `order__bill` DISABLE KEYS */;
INSERT INTO `order__bill` VALUES (3,2,1,6.5,6.5,2,'2015-09-01 19:13:17','2015-09-01 18:37:09');
/*!40000 ALTER TABLE `order__bill` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order__bill_item`
--

DROP TABLE IF EXISTS `order__bill_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order__bill_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bill_id` int(11) DEFAULT NULL,
  `order_item_id` int(11) DEFAULT NULL,
  `item_name` varchar(255) DEFAULT NULL,
  `item_price` double DEFAULT NULL,
  `discount` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order__bill_item`
--

LOCK TABLES `order__bill_item` WRITE;
/*!40000 ALTER TABLE `order__bill_item` DISABLE KEYS */;
INSERT INTO `order__bill_item` VALUES (17,3,7,'California Roll',6.5,0);
/*!40000 ALTER TABLE `order__bill_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order__group`
--

DROP TABLE IF EXISTS `order__group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order__group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table_id` int(11) DEFAULT NULL,
  `cleared` tinyint(1) DEFAULT NULL,
  `cleared_when` datetime DEFAULT NULL,
  `covers` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order__group`
--

LOCK TABLES `order__group` WRITE;
/*!40000 ALTER TABLE `order__group` DISABLE KEYS */;
INSERT INTO `order__group` VALUES (1,3,0,NULL,0),(2,2,1,'2015-09-01 19:13:27',0),(3,1,0,NULL,0),(4,9,0,NULL,0);
/*!40000 ALTER TABLE `order__group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order__group_member`
--

DROP TABLE IF EXISTS `order__group_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order__group_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_id` int(11) DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
  `printed_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order__group_member`
--

LOCK TABLES `order__group_member` WRITE;
/*!40000 ALTER TABLE `order__group_member` DISABLE KEYS */;
INSERT INTO `order__group_member` VALUES (4,3,2,'2015-09-01 19:13:08'),(5,1,3,'2015-09-01 19:16:25'),(6,1,4,NULL);
/*!40000 ALTER TABLE `order__group_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order__item`
--

DROP TABLE IF EXISTS `order__item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order__item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_id` int(11) DEFAULT NULL,
  `order_id` int(11) DEFAULT NULL,
  `notes` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order__item`
--

LOCK TABLES `order__item` WRITE;
/*!40000 ALTER TABLE `order__item` DISABLE KEYS */;
INSERT INTO `order__item` VALUES (7,31,4,''),(12,98,6,''),(13,98,6,''),(14,98,6,'');
/*!40000 ALTER TABLE `order__item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order__item_modifier`
--

DROP TABLE IF EXISTS `order__item_modifier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order__item_modifier` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_item_id` int(11) DEFAULT NULL,
  `modifier_group_id` int(11) DEFAULT NULL,
  `modifier_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order__item_modifier`
--

LOCK TABLES `order__item_modifier` WRITE;
/*!40000 ALTER TABLE `order__item_modifier` DISABLE KEYS */;
/*!40000 ALTER TABLE `order__item_modifier` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `table__items`
--

DROP TABLE IF EXISTS `table__items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `table__items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `table_number` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type_id` (`type_id`),
  CONSTRAINT `type_id` FOREIGN KEY (`type_id`) REFERENCES `config__table_type` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `table__items`
--

LOCK TABLES `table__items` WRITE;
/*!40000 ALTER TABLE `table__items` DISABLE KEYS */;
INSERT INTO `table__items` VALUES (1,1,'Bar 2','2','Bar'),(2,1,'Bar 3','3','Bar'),(3,1,'Bar 4','4','Bar'),(4,2,'Table 1','1','Restaurant'),(5,2,'Table 2','2','Restaurant'),(6,2,'Table 6','6','Restaurant'),(7,2,'Table 5','5','Restaurant'),(8,2,'Table 4','4','Restaurant'),(9,2,'Table 3','3','Restaurant'),(10,2,'Table 7','7','Restaurant'),(11,2,'Table 5A','5A','Restaurant'),(12,2,'Table 5C','5C','Restaurant'),(13,1,'Bar 1','1','Bar'),(14,2,'Table 1A','1A','Restaurant'),(15,2,'Table 2A','2A','Restaurant');
/*!40000 ALTER TABLE `table__items` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-09-02  8:59:38