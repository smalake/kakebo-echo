package private

var PrivateCreate = "INSERT INTO privates (amount, category, store_name, memo, date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
var PrivateGetAll = "SELECT id, amount, category, memo, store_name, date FROM privates WHERE user_id IN (SELECT id FROM users WHERE uid = $1)"
var PrivateGetOne = "SELECT * FROM privates LEFT JOIN users ON users.uid = privates.group_id WHERE users.uid = $1 AND privates.id = $2"
