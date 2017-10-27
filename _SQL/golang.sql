-- phpMyAdmin SQL Dump
-- version 4.7.3
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Окт 27 2017 г., 15:47
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
-- Структура таблицы `domains`
--

CREATE TABLE `domains` (
  `id` int(11) NOT NULL,
  `name` varchar(150) NOT NULL,
  `url` varchar(150) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `domains`
--

INSERT INTO `domains` (`id`, `name`, `url`, `status`, `created_at`, `updated_at`) VALUES
(1, 'tirmarket.com.ua', 'http://tirmarket.com.ua', 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(2, 'wezom.com.ua', 'https://wezom.com.ua', 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(3, 'infocorpus.wezom.net', 'http://infocorpus.wezom.net', 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(4, 'lester.com.ua', 'http://lester.com.ua', 1, '2017-10-27 11:27:25', '2017-10-27 11:27:25'),
(5, 'gmail.com', 'https://gmail.com', 1, '2017-10-27 11:38:13', '2017-10-27 11:38:13');

-- --------------------------------------------------------

--
-- Структура таблицы `ftp`
--

CREATE TABLE `ftp` (
  `id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `host` int(100) NOT NULL,
  `login` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Структура таблицы `mysql`
--

CREATE TABLE `mysql` (
  `id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `hostname` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
-- AUTO_INCREMENT для таблицы `domains`
--
ALTER TABLE `domains`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT для таблицы `ftp`
--
ALTER TABLE `ftp`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT для таблицы `mysql`
--
ALTER TABLE `mysql`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT для таблицы `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
