CREATE TABLE user (
    user_id VARCHAR(36) NOT NULL PRIMARY KEY,
    user_name VARCHAR(24) NOT NULL
);

CREATE TABLE question (
    question_id VARCHAR(36) NOT NULL PRIMARY KEY,
    question_name VARCHAR(50) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    question_type VARCHAR(10) NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    updated_at DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
    question_body VARCHAR(256) NOT NULL,
    FOREIGN KEY fk_user_id(user_id) REFERENCES user(user_id)
);

CREATE TABLE question_value_map (
    question_id VARCHAR(36) NOT NULL,
    value_order TINYINT UNSIGNED NOT NULL,
    question_value VARCHAR(256) NOT NULL,
    FOREIGN KEY fk_question_id(question_id) REFERENCES question(question_id)
);

CREATE TABLE question_answer_map (
    question_id VARCHAR(36) NOT NULL,
    answer_order TINYINT UNSIGNED NOT NULL,
    question_answer VARCHAR(256) NOT NULL,
    FOREIGN KEY fk_question_id(question_id) REFERENCES question(question_id)
);

CREATE TABLE question_tag(
    question_tag_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    question_tag_name VARCHAR(50) NOT NULL
);

CREATE TABLE question_tag_map(
    question_id VARCHAR(36) NOT NULL,
    question_tag_id    INT UNSIGNED,
    FOREIGN KEY fk_question_id(question_id) REFERENCES question(question_id),
    FOREIGN KEY fk_question_tag(question_tag_id) REFERENCES question_tag(question_tag_id)
);

CREATE TABLE question_collection(
    collection_id VARCHAR(36) NOT NULL PRIMARY KEY,
    collection_name VARCHAR(50) NOT NULL,
    collection_description VARCHAR(256),
    user_id VARCHAR(36) NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    updated_at DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
    FOREIGN KEY fk_user_id(user_id) REFERENCES user(user_id)
);

CREATE TABLE question_collection_map(
    collection_id VARCHAR(36) NOT NULL,
    question_order INT UNSIGNED NOT NULL,
    question_id VARCHAR(36) NOT NULL,
    FOREIGN KEY fk_collection_id(collection_id) REFERENCES question_collection(collection_id)
);

insert into user values
    ("c74b366b-5289-d4f5-1d80-66fedaafe97b", "hoge"),
    ("7125174d-6b38-d3ed-b808-ae0a4416f589", "huga"),
    ("116aa423-d551-2b81-2091-132e022d40c5", "piyo"),
    ("30df398b-63f3-351d-2cd2-78ce9df47faf", "foo");

insert into question values
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2","ほげくえすちょん", "c74b366b-5289-d4f5-1d80-66fedaafe97b", "4taku", default, default,"hogeってhoge?"),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5","ふがくえすちょん", "7125174d-6b38-d3ed-b808-ae0a4416f589", "anaume", default, default, "以下の空欄を埋めなさい");

insert into question_value_map values
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", 0, "hoge"),
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", 1, "huga"),
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", 2, "piyo"),
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", 3, "foo"),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", 0, "[]って[]?");

insert into question_answer_map values
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", 0, "hoge"),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", 0, "hoge"),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", 1, "hoge");

insert into question_tag values
    (default, "国語"),
    (default, "算数"),
    (default, "理科"),
    (default, "社会"),
    (default, "Go"),
    (default, "Dart"),
    (default, "C"),
    (default, "JavaScript");

insert into question_tag_map values
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "国語")),
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "算数")),
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "理科")),
    ("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "社会")),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "Go")),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "Dart")),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "C")),
    ("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", (SELECT question_tag_id FROM question_tag WHERE question_tag_name = "JavaScript"));

insert into question_collection values
    ("81b989be-5f8a-a979-ed6d-7a4613ff08e3", "piyo問題集", "テスト用のコレクション", "116aa423-d551-2b81-2091-132e022d40c5", default, default),
    ("2a2fd701-bb7e-6719-1f7b-ec67ffa5a269", "foo問題集", "ンョシクレコの用トステ", "30df398b-63f3-351d-2cd2-78ce9df47faf", default, default);

insert into question_collection_map values
    ("81b989be-5f8a-a979-ed6d-7a4613ff08e3", 0, "1a4ee1ec-1073-f9fb-8281-f3041d15a9d2"),
    ("81b989be-5f8a-a979-ed6d-7a4613ff08e3", 1, "c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5"),
    ("2a2fd701-bb7e-6719-1f7b-ec67ffa5a269", 0, "c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5"),
    ("2a2fd701-bb7e-6719-1f7b-ec67ffa5a269", 1, "1a4ee1ec-1073-f9fb-8281-f3041d15a9d2");
