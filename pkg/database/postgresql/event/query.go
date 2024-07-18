package event

var GetID = "SELECT id, group_id FROM users WHERE uid = $1"
var EventCreate = "INSERT INTO events (amount, category, store_name, memo, date, group_id, revision, create_user, update_user, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
var EventGetAll = "SELECT id, amount, category, memo, store_name, date FROM events WHERE group_id IN (SELECT group_id FROM users WHERE uid = $1)"
var EventGetOne = "SELECT * FROM events LEFT JOIN users ON users.uid = events.group_id WHERE users.uid = $1 AND events.id = $2"
var GetRevision = "SELECT revision FROM groups WHERE id = $1"
var UpdateRevision = "UPDATE groups SET revision = revision + 1 WHERE id = $1 RETURNING revision"
