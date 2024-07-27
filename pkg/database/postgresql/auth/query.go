package auth

// authパッケージ用
var CheckUserByUid = "SELECT group_admin FROM users WHERE uid = $1"
var CreateUser = "INSERT INTO users (uid, name, group_id, register_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
var CreateGroup = "INSERT INTO groups (revision, created_at, updated_at) VALUES (0, $1, $2) RETURNING id"
var CreateRevision = "INSERT INTO private_revision (revision, user_id, created_at, updated_at) VALUES (0, $1, $2, $3)"
var UpdateGroupID = "UPDATE users SET group_id = $1 WHERE id = $2"
