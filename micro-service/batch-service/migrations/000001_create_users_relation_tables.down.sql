ALTER TABLE accounts DROP FOREIGN KEY accounts_userId_fkey;

ALTER TABLE sessions DROP FOREIGN KEY sessions_userId_fkey;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS accounts;

DROP TABLE IF EXISTS sessions;

DROP TABLE IF EXISTS verification_tokens;

