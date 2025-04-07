package adapters

import (
	"UsersFree/src/users/domain/entities"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(user entities.User) error {
	existingDevice, err := m.GetByEsp32Serial(*user.Id_esp32)
	if err != nil && err.Error() != "device not found" {
		return fmt.Errorf("error checking ESP32 serial: %v", err)
	}
	if existingDevice != nil {
		return errors.New("ESP32 serial number already in use. Please enter a different one.")
	}

	err = m.InsertEsp32Serial(*user.Id_esp32)
	if err != nil {
		return fmt.Errorf("error inserting ESP32 serial into the devices table: %v", err)
	}

	query := `INSERT INTO users (name, lastName, email, backupEmail, age, password, id_esp32) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = m.conn.Exec(query, user.Name, user.LastName, user.Email, user.BackupEmail, user.Age, user.Password, user.Id_esp32)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	return nil
}

func (m *MySQL) InsertEsp32Serial(serial string) error {
	serialNumber, err := strconv.ParseInt(serial, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid serial number: %v", err)
	}

	query := `INSERT INTO esp32_devices (serial_number) VALUES (?)`
	_, err = m.conn.Exec(query, serialNumber)
	if err != nil {
		return fmt.Errorf("failed to insert ESP32 serial: %v", err)
	}
	return nil
}

func (m *MySQL) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	query := `SELECT id, name, lastName, email, backupEmail, age, password, id_esp32 FROM users WHERE email = ? LIMIT 1`
	fmt.Println("Executing query:", query, "with email:", email)

	err := m.conn.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No user found for email:", email)
			return entities.User{}, errors.New("user not found")
		}
		fmt.Println("Error executing query:", err)
		return entities.User{}, err
	}
	return user, nil
}

func (m *MySQL) GetByEsp32Serial(serial string) (*entities.User, error) {
	var user entities.User
	query := `SELECT id, name, lastName, email, backupEmail, age, password, id_esp32 FROM users WHERE id_esp32 = ? LIMIT 1`
	err := m.conn.QueryRow(query, serial).Scan(
		&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to check ESP32 serial: %v", err)
	}

	return &user, nil
}

func (m *MySQL) GetAll() ([]entities.User, error) {
	query := "SELECT id, name, lastName, email, backupEmail, age, password, id_esp32 FROM users"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

func (m *MySQL) GetById(id int) (entities.User, error) {
	query := "SELECT id, name, lastName, email, backupEmail, age, password, id_esp32 FROM users WHERE id = ?"
	row := m.conn.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32)
	if err == sql.ErrNoRows {
		return entities.User{}, errors.New("user not found")
	} else if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return user, nil
}

func (m *MySQL) Edit(user entities.User) error {
	query := "UPDATE users SET name = ?, lastName = ?, email = ?, backupEmail = ?, age = ?, password = ? WHERE id = ?"
	_, err := m.conn.Exec(query, user.Name, user.LastName, user.Email, user.BackupEmail, user.Age, user.Password, user.ID)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
