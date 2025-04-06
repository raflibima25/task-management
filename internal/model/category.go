package model

type Category struct {
	ID        string  `json:"id" gorm:"type:uuid;primary_key;default;gen_random_uuid()"`
	Name      string  `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Color     *string `json:"color,omitempty" gorm:"type:varchar(7)"`
	Tasks     []Task  `json:"tasks,omitempty" gorm:"many2many:task_categories"`
	CreatedAt string  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt string  `json:"deleted_at" gorm:"index"`
}

func (c *Category) BeforeCreate() (err error) {
	return nil
}

func (Category) TableName() string {
	return "categories"
}
