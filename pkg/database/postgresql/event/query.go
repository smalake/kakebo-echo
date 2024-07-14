package event

var EventCreate = "INSERT INTO events (amount, category, store_name, memo, date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
var EventGetAll = "SELECT * FROM events LEFT JOIN users ON users.uid = events.group_id WHERE users.uid = $1"
var EventGetOne = "SELECT * FROM events LEFT JOIN users ON users.uid = events.group_id WHERE users.uid = $1 AND events.id = $2"
