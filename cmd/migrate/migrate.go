package main

import (
	"fmt"
	"go-exam/db"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Println("Migration started ....")
	arg := os.Args
	var direction string
	if len(arg) < 2 {
		log.Println("No arg passed")
		os.Exit(1)
	}
	direction = arg[1]

	conn, err := db.Open()
	if err != nil {
		log.Println("DB connection error:", err)
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Println("Failed to start transaction")
	}
	defer tx.Rollback()

	dir := filepath.FromSlash("db/migration")
	info, err := os.Stat(dir)
	if os.IsNotExist(err) || !info.IsDir() {
		log.Println("db/migration is not exits")
		os.Exit(1)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println("Failed to read migration folder:", err)
		os.Exit(1)
	}
	if len(files) == 0 {
		log.Println("Migration folder is empty")
		os.Exit(1)
	}
	var mFiles []string
	for _, e := range files {
		name := e.Name()
		if strings.HasSuffix(name, "."+direction+".sql") {
			path := filepath.Join(dir, name)
			mFiles = append(mFiles, path)
		}
	}

	for _, f := range mFiles {
		content, err := os.ReadFile(f)
		if err != nil {
			log.Fatalln("Failed to read file: ", f)
		}
		sql := string(content)
		log.Println("Migration started for: -------", f)
		_, err = tx.Exec(sql)
		if err != nil {
			log.Println("Migration failed on this file:", f, "error: ", err)
		}
		fmt.Println("✅ Applied migration:", f)
	}

	if err := tx.Commit(); err != nil {
		log.Fatalln("Transaction Failed:", err)
	}

	log.Println("🎉 Migration completed successfully")
	log.Println("File count:", len(mFiles), " direction:", direction)

}
