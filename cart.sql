CREATE TABLE `items` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `status` varchar(191) DEFAULT 'effected',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `item_id` char(128) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `currency` char(8),
  `alias` char(128) DEFAULT NULL,
  `badge` longtext,
  `category` longtext,
  `pic` longtext,
  `urls` longtext,
  `desc` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci



CREATE TABLE `cart_items` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT,
 `status` varchar(16) DEFAULT 'effected',
 `created_at` datetime(3) DEFAULT NULL,
 `updated_at` datetime(3) DEFAULT NULL,
 `cart_id` char(128) DEFAULT NULL,
 `cart_item_id` char(128),
 `item_id` char(128) DEFAULT NULL,
 `count` bigint DEFAULT NULL,
 `amount` double DEFAULT NULL,
 `currency` char(8),
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `carts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` longtext,
  `merchant_id` longtext,
  `status` varchar(16) DEFAULT 'effected',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `store_id` longtext,
  `user_id` longtext,
  `cart_id` longtext,
  `total_amount` double DEFAULT NULL,
  `total_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
