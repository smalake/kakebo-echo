package setting

var GetAdminByUID = `
	SELECT group_admin
	FROM users
	WHERE uid = $1
`
