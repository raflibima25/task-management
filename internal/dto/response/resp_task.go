package response

import (
	"time"

	"github.com/raflibima25/task-management/internal/model"
)

type TaskResponse struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	Description *string            `json:"description,omitempty"`
	Status      model.TaskStatus   `json:"status"`
	Priority    model.Priority     `json:"priority"`
	DueDate     *time.Time         `json:"due_date,omitempty"`
	UserID      string             `json:"user_id"`
	Categories  []CategoryResponse `json:"categories,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func FromTask(task *model.Task) TaskResponse {
	response := TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
		UserID:      task.UserID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	// convert categories jika ada
	if len(task.Categories) > 0 {
		response.Categories = make([]CategoryResponse, len(task.Categories))
		for i, category := range task.Categories {
			response.Categories[i] = FromCategory(&category)
		}
	}

	return response
}

type TasksResponse struct {
	Tasks      []TaskResponse `json:"tasks"`
	TotalCount int64          `json:"total_count"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPage  int            `json:"total_page"`
	HasMore    bool           `json:"has_more"`
}
