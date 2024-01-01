package mysql

// authパッケージ用
var LoginMail = "SELECT id, password FROM users WHERE email = ? AND register_type = 1 "
var LoginGoogle = "SELECT id FROM users WHERE email = ? AND register_type = 2"
var RegisterUser = "INSERT INTO users (email, password, name) VALUES (?, ?, ?)"
