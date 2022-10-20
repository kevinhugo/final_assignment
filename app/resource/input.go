package resource

type InputUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"required"`
}

type InputPhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type InputSocialMedia struct {
	ID              uint   `json:"id"`
	Name            string `json:"name" binding:"required"`
	SocialMedialUrl string `json:"social_media_url" binding:"required"`
}

type UpdateSocialMedia struct {
	Name            string `json:"name" binding:"required"`
	SocialMedialUrl string `json:"social_media_url" binding:"required"`
}

type InputComment struct {
	ID      uint   `json:"id"`
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UpdateComment struct {
	Message string `json:"message" binding:"required"`
}
