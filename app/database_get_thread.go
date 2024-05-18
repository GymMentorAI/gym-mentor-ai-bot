package app

import "log"

/* Query to openai_threads for previous thread*/
func (app *App) getThreadInfo(chatId string) ThreadInfo {
	var data ThreadInfo
	selectThreadQuery := `SELECT id, thread_id FROM openai_threads WHERE chat_id = ? LIMIT 1`

	rows, rowsError := app.MySQL.Query(selectThreadQuery, chatId)
	if rowsError != nil {
		log.Println("ERROR: failed select in openai_threads", rowsError)
		return data
	}
	defer rows.Close()
	for rows.Next() {
		scanError := rows.Scan(&data.MySQLId, &data.ThreadId)
		if scanError != nil {
			log.Println("ERROR: failed row scan in select for openai_threads", scanError)
		}
	}
	return data

}
