package migrate

import (
	"database/sql"
	"fmt"
)

func Apply(db *sql.DB, direction string) error {
	if direction != "up" && direction != "down" {
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	if direction == "up" {
		_, err := db.Exec(UpSQL())
		if err != nil {
			return fmt.Errorf("failed to execute up migration: %v", err)
		}

		fmt.Printf("✔ Applied migration")
	}

	if direction == "down" {
		_, err := db.Exec(DownSQL())
		if err != nil {
			return fmt.Errorf("failed to execute up migration: %v", err)
		}
		fmt.Printf("✔ Applied migration")
	}
	fmt.Println()

	return nil
}
