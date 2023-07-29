FROM mysql

COPY mysql.conf /etc/mysql/mysql.conf

ENV MYSQL_ROOT_PASSWORD Kvado2023
ENV MYSQL_DATABASE librarydb

COPY init.sql /docker-entrypoint-initdb.d/