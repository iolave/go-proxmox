package models

import "github.com/google/uuid"

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
