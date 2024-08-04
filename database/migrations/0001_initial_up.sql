CREATE TABLE `sa_admins` (
  `id` int(11) NOT NULL,
  `player_name` varchar(128) DEFAULT NULL,
  `player_steamid` varchar(64) NOT NULL,
  `flags` text DEFAULT NULL,
  `immunity` int(11) NOT NULL DEFAULT 0,
  `server_id` int(11) DEFAULT NULL,
  `ends` timestamp NULL DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `group_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `sa_servers` (
  `id` int(11) NOT NULL,
  `hostname` varchar(128) DEFAULT NULL,
  `address` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;