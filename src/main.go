package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"text/tabwriter"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "web", "pass", "db", 3306, "cnapp")

type task struct {
	ID     int
	Title  string
	Status string
}

func main() {
	slog.Info("Connecting to database", "connection string", dsn)
	db, err := openDB(dsn)
	if err != nil {
		slog.Error("Error connecting to database: ", err)
	}
	defer db.Close()

	tasks := []task{}

	stmt := `SELECT id, title, status FROM tasks`
	rows, err := db.Query(stmt)
	if err != nil {
		slog.Error("Error querying database: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		t := task{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			slog.Error("Error scanning row: ", err)
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		slog.Error("Error iterating rows: ", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "\nID\tTitle\tStatus")
	for _, t := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\n", t.ID, t.Title, t.Status)
	}
	w.Flush()
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
