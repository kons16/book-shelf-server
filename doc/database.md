# データベース

使用 DB : MySQL  
ORM : gorp  
マイグレーションツール : sql-migrate

## テーブル情報

### users

| カラム名      | 型     | オプション  | 内容 |
| ------------- | ------ | ----------- | ---- |
| id            | BIGINT | PRIMARY KEY |      |
| name          | string |             |      |
| email         | string |             |      |
| password_hash | string |             |      |
| created_at    | time   |             |      |
| updated_at    | time   |             |      |

<br>

### user_session

| カラム名   | 型     | オプション                                      | 内容 |
| ---------- | ------ | ----------------------------------------------- | ---- |
| id         | BIGINT | PRIMARY KEY                                     |      |
| user_id    | BIGINT | FOREIGN KEY ON DELETE CASCADE ON UPDATE CASCADE |      |
| session    | string |                                                 |      |
| created_at | time   |                                                 |      |
| updated_at | time   |                                                 |      |

<br>

### books

| カラム名   | 型     | オプション  | 内容 |
| ---------- | ------ | ----------- | ---- |
| id         | BIGINT | PRIMARY KEY |      |
| title      | string |             |      |
| author     | string |             |      |
| created_at | time   |             |      |
| updated_at | time   |             |      |

<br>

### user_books

| カラム名   | 型     | オプション                                      | 内容 |
| ---------- | ------ | ----------------------------------------------- | ---- |
| id         | BIGINT | PRIMARY KEY                                     |      |
| user_id    | BIGINT | FOREIGN KEY ON DELETE CASCADE ON UPDATE CASCADE |      |
| book_id    | BIGINT | FOREIGN KEY ON DELETE CASCADE ON UPDATE CASCADE |      |
| created_at | time   |                                                 |      |
| updated_at | time   |                                                 |      |
