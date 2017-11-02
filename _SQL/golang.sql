-- phpMyAdmin SQL Dump
-- version 4.7.3
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Ноя 02 2017 г., 13:17
-- Версия сервера: 5.7.19
-- Версия PHP: 7.1.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `golang`
--

-- --------------------------------------------------------

--
-- Структура таблицы `backend`
--

CREATE TABLE `backend` (
  `id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `url` varchar(50) DEFAULT NULL,
  `login` varchar(50) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `backend`
--

INSERT INTO `backend` (`id`, `domain_id`, `url`, `login`, `password`, `status`, `created_at`, `updated_at`) VALUES
(1, 3, 'http://infocorpus.wezom.net', 'kuchura.d.wezom@gmail.com', 'Swqa12345', 1, '2017-11-02 07:13:10', '2017-11-02 07:13:10');

-- --------------------------------------------------------

--
-- Структура таблицы `domains`
--

CREATE TABLE `domains` (
  `id` int(11) NOT NULL,
  `name` varchar(150) NOT NULL,
  `url` varchar(150) NOT NULL,
  `description` text,
  `status` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `domains`
--

INSERT INTO `domains` (`id`, `name`, `url`, `description`, `status`, `created_at`, `updated_at`) VALUES
(1, 'tirmarket.com.ua', 'http://tirmarket.com.ua', NULL, 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(3, 'infocorpus.wezom.net', 'http://infocorpus.wezom.net', NULL, 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(4, 'lester.com.ua', 'http://lester.com.ua', NULL, 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(6, 'gmail.com', 'https://gmail.com', 'Описание, краткая информация', 1, '2017-11-02 08:31:19', '2017-11-02 08:31:19');

-- --------------------------------------------------------

--
-- Структура таблицы `ftp`
--

CREATE TABLE `ftp` (
  `id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `hostname` varchar(150) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `ftp`
--

INSERT INTO `ftp` (`id`, `domain_id`, `hostname`, `username`, `password`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, '91.200.60.68', 'almetcentr_ftp', 'Y1u3W5x8', 1, '2017-10-30 07:07:38', '2017-10-30 07:07:38'),
(2, 1, '91.200.60.68', 'almetcentr_ftp_seo', 'F1a3R5x8', 1, '2017-10-31 10:07:38', '2017-10-31 10:07:38'),
(3, 3, '91.200.60.69', 'infocorpus_ftp', '7J7c9A7v', 1, '2017-11-01 10:07:38', '2017-11-01 10:07:38');

-- --------------------------------------------------------

--
-- Структура таблицы `hosting`
--

CREATE TABLE `hosting` (
  `id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `url` varchar(50) DEFAULT NULL,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `hosting`
--

INSERT INTO `hosting` (`id`, `domain_id`, `url`, `username`, `password`, `status`, `created_at`, `updated_at`) VALUES
(1, 3, '91.200.60.68:3000', 'root', 'GJoy4ewrieurf', 1, '2017-11-02 08:06:00', '2017-11-02 08:06:00');

-- --------------------------------------------------------

--
-- Структура таблицы `mysql`
--

CREATE TABLE `mysql` (
  `id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `phpmyadmin` varchar(50) DEFAULT NULL,
  `hostname` varchar(50) NOT NULL,
  `database_name` varchar(50) DEFAULT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `mysql`
--

INSERT INTO `mysql` (`id`, `domain_id`, `phpmyadmin`, `hostname`, `database_name`, `username`, `password`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 'http://91.200.60.69/phpmyadmin/', '91.200.60.69', 'test_db', 'test_db', 'I4c9R4z5', 1, '2017-11-01 14:46:18', '2017-11-01 14:46:18'),
(2, 3, 'http://91.200.60.69/phpmyadmin/', '91.200.60.69', 'infocorpus_db', 'infocorpus_db', '8L8p6I4v', 1, '2017-11-01 14:46:18', '2017-11-01 14:46:18');

-- --------------------------------------------------------

--
-- Структура таблицы `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(150) NOT NULL,
  `token` varchar(150) NOT NULL,
  `name` varchar(150) NOT NULL,
  `role` tinyint(1) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `token`, `name`, `role`, `status`) VALUES
(1, 'kuchura.d@gmail.com', '4ceb2016943ef079185ad57a1380113d', 'a1aa10018d6111e780bb000c2', 'Dmitry', 0, 1),
(2, 'ivan@gmail.com', '9273137617fc5e5e2f8749a12af17d14', 'ddkiEqaYR7bjl910ZLFFTkJ8a', 'Ivan', 0, 0);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `backend`
--
ALTER TABLE `backend`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Индексы таблицы `domains`
--
ALTER TABLE `domains`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `ftp`
--
ALTER TABLE `ftp`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `hosting`
--
ALTER TABLE `hosting`
  ADD UNIQUE KEY `id` (`id`);

--
-- Индексы таблицы `mysql`
--
ALTER TABLE `mysql`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `backend`
--
ALTER TABLE `backend`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT для таблицы `domains`
--
ALTER TABLE `domains`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
--
-- AUTO_INCREMENT для таблицы `ftp`
--
ALTER TABLE `ftp`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT для таблицы `hosting`
--
ALTER TABLE `hosting`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT для таблицы `mysql`
--
ALTER TABLE `mysql`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT для таблицы `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
