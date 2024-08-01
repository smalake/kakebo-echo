package private

var PrivateCreate = "INSERT INTO privates (amount, category, store_name, memo, date, user_id, revision, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, (SELECT id FROM users WHERE uid = $6), $7, $8, $9) RETURNING id"
var PrivateGetAll = "SELECT id, amount, category, memo, store_name, date FROM privates WHERE user_id IN (SELECT id FROM users WHERE uid = $1)"
var PrivateGetOne = `
	SELECT
		id,
		amount,
		category,
		memo,
		store_name,
		date,
		created_at,
		updated_at
	FROM privates
	WHERE id = $1
	AND user_id
		IN (SELECT id FROM users WHERE uid = $2)
`
var PrivateUpdate = `
	UPDATE privates SET
		amount = $1,
		category = $2,
		memo = $3,
		store_name = $4,
		date = $5,
		updated_at = $6,
		revision = $7
	WHERE id = $8
	AND user_id
		IN (SELECT id FROM users WHERE uid = $9)
`
var PrivateDelete = `
	DELETE FROM privates
	WHERE id = $1
	AND user_id
		IN (SELECT id FROM users WHERE uid = $2)
`
var GetRevision = "SELECT revision FROM private_revision WHERE user_id IN (SELECT id FROM users WHERE uid = $1)"
var UpdateRevision = "UPDATE private_revision SET revision = revision + 1, updated_at = $1 WHERE user_id IN (SELECT id FROM users WHERE uid = $2) RETURNING revision"
