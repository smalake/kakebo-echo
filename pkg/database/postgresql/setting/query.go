package setting

var GetAdminByUID = `
	SELECT group_admin
	FROM users
	WHERE uid = $1
`

var GetName = `
	SELECT name
	FROM users
	WHERE uid = $1
`
var UpdateName = `
	UPDATE users
	SET
		name = $1,
		updated_at = $2
	WHERE uid = $3
`
