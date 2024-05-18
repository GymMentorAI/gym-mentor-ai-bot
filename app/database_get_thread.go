package app

import "log"

/* Query to openai_threads for previous thread*/
func (app *App) getThreadInfo(username string) ThreadInfo {
	var data ThreadInfo
	selectThreadQuery := `SELECT id, thread_id FROM openai_threads WHERE username = ? LIMIT 1`

	rows, rowsError := app.MySQL.Query(selectThreadQuery, username)
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
