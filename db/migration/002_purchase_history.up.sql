
CREATE TABLE if not EXISTS purchase_history(
    id SERIAL PRIMARY KEY,
    dish_id INT,
    restaurant_id INT,
    user_id INT,
    amount FLOAT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
    FOREIGN KEY (restaurant_id) REFERENCES restaurant(id)
    FOREIGN KEY (dish_id) REFERENCES restaurant_menu(id)
)