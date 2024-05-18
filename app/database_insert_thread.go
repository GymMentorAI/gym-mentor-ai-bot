package app

/* insert in openai_threads new record */
func (app *App) insertUsernameAndThreadId(username, threadId string) (int, error) {

	insertQuery := `ÌNSERT INTO openai_threads (username, thread_id) VALUES (?, ?)`
	inserResult, insertError := app.MySQL.Exec(insertQuery, username, threadId)
	if insertError != nil {
		return 1, insertError
	}
	lastId, lastIdError := inserResult.LastInsertId()
	if lastIdError != nil {
		return int(lastId), lastIdError
	}
	return int(lastId), nil

}
