package model

import "time"

// enum untuk status task
type TaskStatus string

// enum untuk prioritas task
type Priority string

// konstanta untuk task status
const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
)

// konstanta untuk prioritas task
const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

type Task struct {
	ID          string     `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Description *string    `json:"description,omitempty" gorm:"type:text"`
	Status      TaskStatus `json:"status" gorm:"type:task_status;default:'TODO';not null"`
	Priority    Priority   `json:"priority" gorm:"type:task_priority;default:'MEDIUM';not null"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	UserID      string     `json:"user_id" gorm:"type:uuid;not null"`
	User        User       `json:"-" gorm:"foreignKey:UserID"`
	Categories  []Category `json:"categories,omitempty" gorm:"many2many:task_categories"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

func (t *Task) BeforeCreate() (err error) {
	return nil
}

func (Task) TableName() string {
	return "tasks"
}
