CREATE DATABASE IF NOT EXISTS `lemonilo`;

USE `lemonilo`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(27) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `email` varchar(50) NOT NULL,
  `address` text NOT NULL,
  `password` varchar(50) NOT NULL,
  `status` enum('on','off') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);
COMMIT;