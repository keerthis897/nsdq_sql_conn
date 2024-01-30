package sqlserver

import (
	"context"
	"database/sql"
)

func CreateTableIfNotExists(ctx context.Context, conn *sql.DB) error {
	_, err := conn.ExecContext(ctx, `
	IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[students]') AND type in (N'U'))
	BEGIN
		CREATE TABLE students (
			ID INT PRIMARY KEY IDENTITY,
			Name NVARCHAR(50)
		)
	END`)
	return err
}

func InsertData(ctx context.Context, conn *sql.DB, name string) (int64, error) {
	result, err := conn.ExecContext(ctx, "INSERT INTO students (name) VALUES ('John6')")
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
