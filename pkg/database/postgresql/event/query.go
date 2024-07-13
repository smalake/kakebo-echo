package event

var EventCreate = "INSERT INTO event (amount, category, store_name, memo, date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
var EventGetAll = "SELECT * FROM event WHERE uid = $1"
