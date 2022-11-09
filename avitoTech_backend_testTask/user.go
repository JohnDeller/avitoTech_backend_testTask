package avitotech

type User struct {
	Id      int     `json:"id" db:"id" binding:"required"`
	Balance float32 `json:"balance" binding:"required"`
}
