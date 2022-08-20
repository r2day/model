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
