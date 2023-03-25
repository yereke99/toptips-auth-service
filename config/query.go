package config

var (
	createTableSMS = `CREATE TABLE IF NOT EXISTS sms_cache(
		id SERIAL PRIMARY KEY,
		contact varchar(50),
		code INTEGER
	);`
	profile = `CREATE TABLE IF NOT EXISTS profile(
		id SERIAL PRIMARY KEY,
		firstName varchar(50),
		lastName varchar(50),
		phone varchar(50),
		status varchar(50),
		password varchar(50),
		role varchar(50),
		userId BIGINT,
		token varchar(500)
	);`
)
