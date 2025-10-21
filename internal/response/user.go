
package response

import "ai_admin_project/internal/model"

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func ToUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
