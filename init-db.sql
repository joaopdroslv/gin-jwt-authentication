
CREATE DATABASE IF NOT EXISTS main;

USE main;

CREATE TABLE IF NOT EXISTS user_statuses (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name            VARCHAR(32) NOT NULL UNIQUE,
    description     TEXT DEFAULT NULL,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO main.user_statuses (
    name,
    description
) VALUES
    ("active", NULL),
    (
        "inactive",
        "A inactive user cannot perform any action."
    ),
    (
        "email_confirmation",
        "When a user's account is awaiting the email confirmation, it's not considered active yet."
    ),
    (
        "password_creation",
        "When a user's account is awaiting the creation of the password, it's not considered active yet."
    ),
    (
        "deleted_account",
        "When a user requests the deletion of their account, it's considered inactive."
    )
;

SET @active_user_status_id := (SELECT id FROM main.user_statuses WHERE name = "active" LIMIT 1);

CREATE TABLE IF NOT EXISTS user_credentials (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    email           VARCHAR(128) UNIQUE NOT NULL,
    password_hash   VARCHAR(255) DEFAULT NULL,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id                  BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_credential_id  BIGINT UNSIGNED NOT NULL,
    user_status_id      BIGINT UNSIGNED NOT NULL,
    name                VARCHAR(128) NOT NULL,
    birthdate           DATE NOT NULL,
    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_user_credential FOREIGN KEY (user_credential_id) REFERENCES main.user_credentials(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_user_status FOREIGN KEY (user_status_id) REFERENCES main.user_statuses(id),

    UNIQUE KEY uk_user_user_credential (user_credential_id),
    INDEX idx_user_with_status (id, user_status_id)
);

INSERT INTO main.user_credentials (
    email,
    password_hash
) VALUES
    (
        "system@system.com",
        "$2a$12$sZ.BjwbUgXAigyfepBLH7uUXijODjjRUMEGEKRKCitjAN8yciNjhe" -- bcrypt(12) == "systemsystem123"
    )
;

SET @system_user_credential_id := (SELECT id FROM main.user_credentials WHERE email = "system@system.com" LIMIT 1);

INSERT INTO main.users (
    name,
    birthdate,
    user_credential_id,
    user_status_id
) VALUES
    (
        "system",
        "2000-01-01",
        @system_user_credential_id,
        @active_user_status_id
    )
;

SET @system_user_id := (
    SELECT id
    FROM main.users
    WHERE user_credential_id = @system_user_credential_id
    LIMIT 1
);
