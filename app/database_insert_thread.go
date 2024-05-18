package app

/* insert in openai_threads new record */
func (app *App) insertUsernameAndThreadId(username, chatId, threadId string) (int, error) {

	insertQuery := `INSERT INTO openai_threads (username, chat_id, thread_id) VALUES (?, ?, ?)`
	inserResult, insertError := app.MySQL.Exec(insertQuery, username, chatId, threadId)
	if insertError != nil {
		return 1, insertError
	}
	lastId, lastIdError := inserResult.LastInsertId()
	if lastIdError != nil {
		return int(lastId), lastIdError
	}
	return int(lastId), nil

}
