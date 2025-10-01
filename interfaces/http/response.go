package http

type UserPostRequest struct {
	UserName string `json:"username" validate:"required"`
}

type UserPutRequest struct {
	Name string `json:"username" validate:"required"`
}

type UserResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type UserIndexResponse struct {
	Users []UserResponse `json:"users"`
}

type UserGetResponse struct {
	User UserResponse `json:"user"`
}