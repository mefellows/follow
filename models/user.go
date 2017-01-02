package models

// User is someone using the system
type User struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	DeviceID  string `json:"device_id"`
}
