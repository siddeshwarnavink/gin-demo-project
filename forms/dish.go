package forms

type DishForm struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Thumbnail   string `form:"thumbnail" binding:"required,url"`
}
