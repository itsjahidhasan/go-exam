CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_name VARCHAR(100) NOT NULL, 
        cash_balance FLOAT NOT NULL
	);

