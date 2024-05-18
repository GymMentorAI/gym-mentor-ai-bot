package app

import "log"

/* create MySQL necessary tables: openai_threads or die*/
func (app *App) createMysqlTables() {
	createOpenAiThreads := `

	CREATE TABLE IF NOT EXISTS openai_threads (
		id INTEGER auto_increment NOT NULL,
		username varchar(200) NOT NULL,
		chat_id varchar(200) NOT NULL,
		thread_id varchar(200) NOT NULL,
		PRIMARY KEY (ID)
	)
	ENGINE=InnoDB
	DEFAULT CHARSET=utf8mb4
	COLLATE=utf8mb4_0900_ai_ci;	
	`

	_, openAiThreadsError := app.MySQL.Exec(createOpenAiThreads)
	if openAiThreadsError != nil {
		log.Fatalln("FATAL ERROR: create openai_threads failed", openAiThreadsError)
	}

}
