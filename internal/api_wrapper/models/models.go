package models

import (
	"database/sql"
	"os"
	"path"

	_ "modernc.org/sqlite"
)

const CREATE_TABLE_CMD_EXEC = `CREATE TABLE IF NOT EXISTS cmd_execution (
	id 		string not null,
	status 		string not null,
	output 		string,
	exit_code 	int,
	error 		string,
	primary key (id)
)`

type Models struct {
	db           *sql.DB
	CMDExecution CMDExecutionModel
}

func (m *Models) Close() error {
	return m.db.Close()
}

func Initialize() (*Models, error) {
	models := new(Models)
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dataPath := path.Join(homePath, ".pve-api-wrapper")
	dbFile := path.Join(dataPath, "storage.db")
	if err := os.MkdirAll(dataPath, os.ModePerm); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}

	createQuerys := []string{
		CREATE_TABLE_CMD_EXEC,
	}

	for _, q := range createQuerys {
		_, err = db.Exec(q)
		if err != nil {
			return nil, err
		}
	}

	models.db = db
	models.CMDExecution = CMDExecutionModel{models: models}
	return models, nil
}
