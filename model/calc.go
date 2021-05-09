package model

type CalcAddPaylod struct{
	X *int `uri:"x" binding:"required"`
	Y *int `uri:"y" binding:"required"`
}

type CalcDivPayload struct {
	X *int `uri:"x" binding:"required"`
	Y *int `uri:"y" binding:"required"`
}
