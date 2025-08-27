package migrate

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Apply(db *sql.DB, dir string, direction string) error {
	// Placeholder for migration logic
	if direction != "up" && direction != "down" {
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	// Simulate migration application
	entries, err := os.ReadDir(dir)

	if err != nil {
		return err

	}

	var files []string

	for _, e := range entries {
		name := e.Name()
		if strings.HasSuffix(name, "."+direction+".sql") {
			files = append(files, filepath.Join(dir, name))
		}

		sort.Strings(files)

		for _, f := range files {
			b, err := os.ReadFile(f)
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %v", f, err)
			}
			_, err = db.Exec(string(b))
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", f, err)
			}
			fmt.Printf("✔ Applied migration: %s\n", f)
		}

	}

	return nil
}
