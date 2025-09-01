CREATE TABLE IF NOT EXISTS restaurant_menu (
		id SERIAL PRIMARY KEY,
		menu_name VARCHAR(150) NOT NULL, 
        price FLOAT NOT NULL,
        restaurant_id INT,
        FOREIGN KEY (restaurant_id) REFERENCES restaurant(id)
	);
