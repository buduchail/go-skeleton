CREATE TABLE `altares` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `id_difunto` bigint(11) NOT NULL,
  `dedicatoria` varchar(32) NOT NULL DEFAULT '',
  `niveles` int(11) NOT NULL,
  `repisas` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `difuntos` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `nombre` varchar(32) NOT NULL DEFAULT '',
  `nacido` varchar(32) NOT NULL DEFAULT '',
  `fallecido` varchar(32) NOT NULL DEFAULT '',
  `campo` varchar(16) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `difuntos` VALUES (1,'Frida Kahlo','6 de julio de 1907','13 de julio de 1954','Arts'),(2,'Cantinflas','12 de agosto de 1911','20 de abril de 1993','Entertainment'),(3,'Emiliano Zapata','8 de agosto de 1879','10 de abril de 1919','Politics');

CREATE TABLE `ofrendas` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `tipo` varchar(16) NOT NULL DEFAULT '',
  `nombre` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `ofrendas` VALUES (1,'food','Chocolate'),(2,'food','Calabaza'),(3,'food','Fruta'),(4,'food','Pan de muerto'),(5,'food','Tamal'),(6,'food','Mole'),(7,'drink','Agua'),(8,'drink','Tequila'),(9,'drink','Mezcal'),(10,'flower','Cempasúchil'),(11,'flower','Amaranto'),(12,'flower','Alheli'),(13,'flower','Nube');
