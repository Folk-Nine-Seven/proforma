package model

type (
	Task struct {
		Id          string `json:"id"`
		Summary     string `json:"summary" binding:"required"`
		Description string `json:"description,omitempty"`
		Status      Status `json:"status"`
		SubTasks    []Task `json:"subtasks,omitempty" binding:"dive"`
		Metadata
	}
)
