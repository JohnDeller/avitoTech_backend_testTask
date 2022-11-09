package avitotech

type Order struct {
	Id      int     `json:"id" db:"id"`
	UserId  int     `json:"user_id" binding:"required"`
	Balance float32 `json:"balance" binding:"required"`
}
