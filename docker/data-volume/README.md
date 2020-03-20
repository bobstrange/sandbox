# Use data volume

Build a data volume container image.

```bash
docker image build -t bobstrange/data-volume:latest .
```

Create a data volume container.

```bash
docker container run --name data-volume bobstrange/data-volume
```

Even if container is stopped, there is a still container in the storage.

```bash
docker container run --rm --name mysql \
> -e "MYSQL_ALLOW_EMPTY_PASSWORD=yes" \
> -e "MYSQL_DATABASE=volume-test" \
> -e "MYSQL_USER=example" \
> -e "MYSQL_PASSWORD=example" \
> --volumes-from data-volume \
> mysql
```

Create a test table and insert some records.

```bash
docker container exec -it mysql mysql -u root -p volume-test
Enter password:
```

```mysql
mysql> CREATE TABLE user(
    -> id int PRIMARY KEY AUTO_INCREMENT,
    -> name VARCHAR(255)
    -> ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;
Query OK, 0 rows affected (0.01 sec)

mysql> INSERT INTO user (name) VALUES ('test'), ('test1'), ('test2');
Query OK, 3 rows affected (0.01 sec)
Records: 3  Duplicates: 0  Warnings: 0
```

Stop mysql container.(Remember you passed `--rm` options when you started the container)
Start mysql container.

```bash
docker container run --rm --name mysql \
> -e "MYSQL_ALLOW_EMPTY_PASSWORD=yes" \
> -e "MYSQL_DATABASE=volume-test" \
> -e "MYSQL_USER=example" \
> -e "MYSQL_PASSWORD=example" \
> --volumes-from data-volume \
> mysql
```

There is still data there.

```bash
docker container exec -it mysql mysql -u root -p volume-test
Enter password:
```

```mysql

mysql> SELECT * FROM user;
+----+-------+
| id | name  |
+----+-------+
|  1 | test  |
|  2 | test1 |
|  3 | test2 |
+----+-------+
3 rows in set (0.00 sec)

```


