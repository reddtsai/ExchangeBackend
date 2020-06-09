CREATE TABLE `ORDER` (
  `ID` varchar(35) NOT NULL,
  `SYMBOL` varchar(20) NOT NULL,
  `TYPE` varchar(5) NOT NULL,
  `AMOUNT` varchar(45) NOT NULL,
  `PRICE` varchar(45) NOT NULL,
  `UNIX` int NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci