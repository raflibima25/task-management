package request

import (
	"time"

	"github.com/raflibima25/task-management/internal/model"
)

type CreateTaskRequest struct {
	Title        string            `json:"title" binding:"required"`
	Description  *string           `json:"description,omitempty"`
	Status       *model.TaskStatus `json:"status,omitempty"`
	Priority     *model.Priority   `json:"priority,omitempty"`
	DueDate      *time.Time        `json:"due_date,omitempty"`
	CategoroyIDs []string          `json:"category_ids,omitempty"`
}

type UpdateTaskRequest struct {
	Title       *string           `json:"title,omitempty"`
	Description *string           `json:"description,omitempty"`
	Status      *model.TaskStatus `json:"status,omitempty"`
	Priority    *model.Priority   `json:"priority,omitempty"`
	DueDate     *time.Time        `json:"due_date,omitempty"`
	CategoryIDs []string          `json:"category_ids,omitempty"`
}

type TaskFilterRequest struct {
	Status     *model.TaskStatus `form:"status,omitempty"`
	Priority   *model.Priority   `form:"priority,omitempty"`
	CategoryID *string           `form:"category_id,omitempty"`
	SearchTerm *string           `form:"search_term,omitempty"`
	StartDate  *time.Time        `form:"start_date,omitempty" time_format:"2006-01-02"`
	EndDate    *time.Time        `form:"end_date,omitempty" time_format:"2006-01-02"`
	Page       int               `form:"page,default=1" binding:"min=1"`
	Limit      int               `form:"limit,default=10" binding:"min=1,max=100"`
}
