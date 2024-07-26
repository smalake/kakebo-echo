package event

var GetID = "SELECT id, group_id FROM users WHERE uid = $1"
var EventCreate = "INSERT INTO events (amount, category, store_name, memo, date, group_id, revision, create_user, update_user, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
var EventGetAll = "SELECT id, amount, category, store_name, date FROM events WHERE group_id IN (SELECT group_id FROM users WHERE uid = $1)"
var EventGetOne = `
	SELECT
		events.id AS id,
		amount,
		category,
		memo,
		store_name,
		date,
		create_u.name AS create_user,
		update_u.name AS update_user,
		events.created_at AS created_at,
		events.updated_at AS updated_at
	FROM events
	JOIN users AS create_u ON create_u.id = events.create_user
	JOIN users AS update_u ON update_u.id = events.update_user
	WHERE events.id = $1
	AND events.group_id 
		IN (SELECT group_id FROM users WHERE uid = $2)
`

var EventUpdate = `
	UPDATE events SET
		amount = $1,
		category = $2,
		memo = $3,
		store_name = $4,
		date = $5,
		update_user = $6,
		updated_at = $7,
		revision = $8
	WHERE id = $9
	AND group_id
		IN (SELECT group_id FROM users WHERE uid = $10)
`
var GetRevision = "SELECT revision FROM groups WHERE id = $1"
var UpdateRevision = "UPDATE groups SET revision = revision + 1 WHERE id = $1 RETURNING revision"
