CREATE TABLE authors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

INSERT INTO authors (name) VALUES ('Александр Пушкин');
INSERT INTO authors (name) VALUES ('Лев Толстой');
INSERT INTO authors (name) VALUES ('Фёдор Достоевский');
INSERT INTO authors (name) VALUES ('Михаил Булгаков');
INSERT INTO authors (name) VALUES ('Антон Чехов');
INSERT INTO authors (name) VALUES ('Николай Гоголь');
INSERT INTO authors (name) VALUES ('Иван Тургенев');
INSERT INTO authors (name) VALUES ('Алексей Толстой');
INSERT INTO authors (name) VALUES ('Максим Горький');
INSERT INTO authors (name) VALUES ('Илья Ильф');
INSERT INTO authors (name) VALUES ('Евгений Петров');
INSERT INTO authors (name) VALUES ('Владимир Набоков');
INSERT INTO authors (name) VALUES ('Александр Солженицын');
INSERT INTO authors (name) VALUES ('Сергей Есенин');
INSERT INTO authors (name) VALUES ('Михаил Шолохов');
INSERT INTO authors (name) VALUES ('Борис Пастернак');
INSERT INTO authors (name) VALUES ('Анна Ахматова');
INSERT INTO authors (name) VALUES ('Марина Цветаева');
INSERT INTO authors (name) VALUES ('Сергей Довлатов');
INSERT INTO authors (name) VALUES ('Андрей Курков');

INSERT INTO books (title, author_id) VALUES ('Евгений Онегин', 1);
INSERT INTO books (title, author_id) VALUES ('Война и мир', 2);
INSERT INTO books (title, author_id) VALUES ('Преступление и наказание', 3);
INSERT INTO books (title, author_id) VALUES ('Мастер и Маргарита', 4);
INSERT INTO books (title, author_id) VALUES ('Чайка', 5);
INSERT INTO books (title, author_id) VALUES ('Мёртвые души', 6);
INSERT INTO books (title, author_id) VALUES ('Отцы и дети', 7);
INSERT INTO books (title, author_id) VALUES ('Гиперболоид инженера Гарина', 8);
INSERT INTO books (title, author_id) VALUES ('Детство', 9);
INSERT INTO books (title, author_id) VALUES ('Двенадцать стульев', 10);
INSERT INTO books (title, author_id) VALUES ('Собачье сердце', 10);
INSERT INTO books (title, author_id) VALUES ('Лолита', 12);
INSERT INTO books (title, author_id) VALUES ('Один день Ивана Денисовича', 13);
INSERT INTO books (title, author_id) VALUES ('Поэма о крыловоде', 14);
INSERT INTO books (title, author_id) VALUES ('Тихий Дон', 15);
INSERT INTO books (title, author_id) VALUES ('Доктор Живаго', 16);
INSERT INTO books (title, author_id) VALUES ('Реквием', 17);
INSERT INTO books (title, author_id) VALUES ('Анна Ахматова. Стихотворения', 18);
INSERT INTO books (title, author_id) VALUES ('Москва-Петушки', 19);
INSERT INTO books (title, author_id) VALUES ('Белая гвардия', 20);