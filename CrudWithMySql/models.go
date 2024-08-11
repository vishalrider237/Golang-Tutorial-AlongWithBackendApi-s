package CrudWithMySql

type Employee struct {
	ID       uint   `json:"id" gorm:"primaryKey"` // Auto-generated ID
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Position string `json:"position" binding:"required"`
	Salary   int    `json:"salary" binding:"required"`
}
