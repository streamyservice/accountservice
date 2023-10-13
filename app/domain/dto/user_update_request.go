package dto


import(
	"time"
)

type UserUpdate struct{
	FullName string `json:"fullname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Image string `json:"image"`
	Gender string  `json:"gender"`
}