package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CMDExecutionStatus string

const (
	CMD_EXEC_STATUS_RUNNING   CMDExecutionStatus = "RUNNING"
	CMD_EXEC_STATUS_SUCCEEDED CMDExecutionStatus = "SUCCEEDED"
	CMD_EXEC_STATUS_FAILED    CMDExecutionStatus = "FAILED"
)

type CMDExecution struct {
	ID       string             `json:"id"`
	Status   CMDExecutionStatus `json:"status"`
	ExitCode *int               `json:"exitCode"`
	Output   *string            `json:"output"`
	Error    *string            `json:"error"`
}

type CMDExecutionModel struct {
	models *Models
}

func (service *CMDExecutionModel) New() (id string, err error) {
	id = uuid.New().String()
	status := CMD_EXEC_STATUS_RUNNING
	stmt := `INSERT INTO cmd_execution(
			id,
			status
		) values(?, ?)`
	if _, err = service.models.db.Exec(
		stmt,
		id,
		status,
	); err != nil {
		return "", err
	}

	return id, nil
}

func (service *CMDExecutionModel) upsert(
	id string,
	status CMDExecutionStatus,
	exitCode *int,
	output *string,
	error error,
) (err error) {
	strerr := new(string)
	if error != nil {
		msg := error.Error()
		strerr = &msg
	}
	stmt := `INSERT OR REPLACE INTO cmd_execution(
			id,
			status,
			exit_code,
			output,
			error
		) values(?, ?, ?, ?, ?)`
	if _, err = service.models.db.Exec(
		stmt,
		id,
		status,
		exitCode,
		output,
		strerr,
	); err != nil {
		return err
	}

	return nil
}

func (s *CMDExecutionModel) SetFailed(
	id string,
	err error,
) error {
	return s.upsert(id, CMD_EXEC_STATUS_FAILED, nil, nil, err)
}

func (s *CMDExecutionModel) SetSucceeded(
	id,
	output string,
	exitCode int,
) error {
	return s.upsert(id, CMD_EXEC_STATUS_SUCCEEDED, &exitCode, &output, nil)
}

func (s *CMDExecutionModel) Get(id string) (result *CMDExecution, err error) {
	if id == "" {
		return result, errors.New(`please provide a valid id`)
	}

	stmt := fmt.Sprintf(`SELECT
		id,
		status,
		exit_code,
		output,
		error
	FROM cmd_execution
	WHERE id = "%s"
	`, id)
	rows, err := s.models.db.Query(stmt)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	result = new(CMDExecution)
	if err := rows.Scan(
		&result.ID,
		&result.Status,
		&result.ExitCode,
		&result.Output,
		&result.Error,
	); err != nil {
		return nil, err
	}

	return result, nil
}
