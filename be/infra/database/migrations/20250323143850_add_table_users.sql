-- +goose Up
CREATE TABLE users (
                       id int AUTO_INCREMENT NOT NULL COMMENT 'ユーザーID'
    , login_id VARCHAR(30) NOT NULL COMMENT 'ログインID'
    , name VARCHAR(50) NOT NULL COMMENT '氏名'
    , password VARCHAR(300) NOT NULL COMMENT 'パスワード'
    , change_pass_flag BOOLEAN DEFAULT true NOT NULL COMMENT 'パスワード変更フラグ'
    , error_count INT DEFAULT 0 NOT NULL COMMENT 'エラー回数'
    , error_datetime DATETIME COMMENT 'エラー時刻'
    , lock_datetime DATETIME COMMENT 'ロック開始時刻'
    , created_at datetime DEFAULT now() NOT NULL COMMENT '作成日'
    , updated_at datetime DEFAULT now() NOT NULL COMMENT '更新日'
    , deleted_flg boolean DEFAULT false NOT NULL COMMENT '削除フラグ'
    , deleted_at datetime COMMENT '削除日'
    , CONSTRAINT users_PKC PRIMARY KEY (id)
) COMMENT 'ユーザー情報' ;

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd