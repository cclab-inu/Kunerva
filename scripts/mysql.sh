#!/bin/bash

sudo apt install mysql-server -y

$ sudo mysql -u root

mysql> USE mysql;
mysql> SELECT User, Host, plugin FROM mysql.user;

mysql> UPDATE user SET plugin='mysql_native_password' WHERE User='root';
mysql> FLUSH PRIVILEGES;
mysql> SELECT User, Host, plugin FROM mysql.user;

mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '변경할-비밀번호';
mysql> FLUSTH PRIVILEGES;

