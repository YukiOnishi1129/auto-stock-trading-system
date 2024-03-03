CREATE TABLE IF NOT EXISTS users(
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `name` VARCHAR(191) NULL,
   `email` VARCHAR(191) NULL,
   `password` VARCHAR(255) NOT NULL,
   `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
   `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
   `deleted_at` DATETIME(3) NULL,
   UNIQUE INDEX `users_email_key`(`email`),
   PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;