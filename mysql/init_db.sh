#!/bin/bash

sudo mysql -u root

mysql> USE mysql;

mysql> SELECT User, Host, plugin FROM mysql.user;

mysql> update user set plugin='mysql_native_password' where user='root';

mysql> ALTER user 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'cclab';

mysql> FLUSH PRIVILEGES;

mysql> quit
