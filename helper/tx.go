package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	r := recover()
	if r != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(r)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
