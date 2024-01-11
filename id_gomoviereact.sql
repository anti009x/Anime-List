-- MySQL database dump

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

-- MySQL specific settings
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- Create genres table
CREATE TABLE genres (
    id INT UNSIGNED NOT NULL,
    genre_name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

-- Create movies table
CREATE TABLE movies (
    id INT UNSIGNED NOT NULL,
    title VARCHAR(255),
    description TEXT,
    year INT,
    release_date DATE,
    runtime INT,
    rating INT,
    mpaa_rating VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

-- Create movies_genres table
CREATE TABLE movies_genres (
    id INT UNSIGNED NOT NULL,
    movie_id INT UNSIGNED,
    genre_id INT UNSIGNED,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (genre_id) REFERENCES genres(id),
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);

-- Insert data into genres table
INSERT INTO genres (id, genre_name, created_at, updated_at)
VALUES
    (1, 'Drama', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (2, 'Crime', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (3, 'Action', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (4, 'Comic Book', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (5, 'Sci-Fi', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (6, 'Mystery', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (7, 'Adventure', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (8, 'Comedy', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (9, 'Romance', '2021-05-17 00:00:00', '2021-05-17 00:00:00');

-- Insert data into movies table
INSERT INTO movies (id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at)
VALUES
    (1, 'Spy x Family', 'Corrupt politicians, frenzied nationalists, and other warmongering forces constantly jeopardize the thin veneer of peace between neighboring countries Ostania and Westalis. In spite of their plots, renowned spy and master of disguise Twilight fulfills dangerous missions one after another in the hope that no child will have to experience the horrors of war.', 2022, '2022-10-14', 142, 5, 'R', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (2, 'One Piece', 'Gol D. Roger was known as the Pirate King the strongest and most infamous being to have sailed the Grand Line. The capture and execution of Roger by the World Government brought a change throughout the world', 1999, '1999-03-24', 175, 5, 'R', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (4, 'Kaguya-sama wa Kokurasetai', 'At the renowned Shuchiin Academy, Miyuki Shirogane and Kaguya Shinomiya are the student bodys top representatives. Ranked the top student in the nation and respected by peers and mentors alike, Miyuki serves as the student council president.', 2019, '2019-04-14', 102, 4, 'R', '2021-05-17 00:00:00', '2021-05-17 00:00:00'),
    (3, 'Komi-san wa Comyushou desu', 'Hitohito Tadano is an ordinary boy who heads into his first day of high school with a clear plan: to avoid trouble and do his best to blend in with others. Unfortunately, he fails right away when he takes the seat beside the schools madonna—Shouko Komi.', 2021, '2021-07-18', 152, 5, 'PG13', '2021-05-17 00:00:00', '2021-05-17 00:00:00');

-- Set auto-increment values for sequences
ALTER TABLE genres AUTO_INCREMENT = 10;
ALTER TABLE movies_genres AUTO_INCREMENT = 1;
ALTER TABLE movies AUTO_INCREMENT = 5;

-- MySQL specific settings
SET FOREIGN_KEY_CHECKS = 1;
