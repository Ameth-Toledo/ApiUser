package entities

type User struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	LastName    string  `json:"lastName"`
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	Age         int32   `json:"age"`
	BackupEmail *string `json:"backupEmail"`
	Id_esp32    *string `json:"esp32Serial"`
}
