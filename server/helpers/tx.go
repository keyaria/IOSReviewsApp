package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		ThrowIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		ThrowIfError(errCommit)
	}
}
