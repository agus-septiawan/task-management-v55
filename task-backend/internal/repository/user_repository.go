package repository

import (
	"database/sql"
	"fmt"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetByID(id int) (*model.User, error)
	GetByOAuth(provider, oauthID string) (*model.User, error)
	GetAll(page, limit int) ([]model.User, int, error)
	Update(user *model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository membuat instance UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Create membuat user baru
func (r *userRepository) Create(user *model.User) error {
	var hashedPassword *string

	// Hash password only if provided (not for OAuth users)
	if user.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		hashedStr := string(hashed)
		hashedPassword = &hashedStr
	}

	query := `
		INSERT INTO users (email, name, password, role, oauth_provider, oauth_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, user.Email, user.Name, hashedPassword, user.Role, user.OauthProvider, user.OauthID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	user.ID = int(id)
	return nil
}

// GetByEmail mengambil user berdasarkan email
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, name, password, role, oauth_provider, oauth_id, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	var user model.User
	var password sql.NullString
	var oauthProvider sql.NullString
	var oauthID sql.NullString

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&password,
		&user.Role,
		&oauthProvider,
		&oauthID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	if password.Valid {
		user.Password = password.String
	}

	if oauthProvider.Valid {
		user.OauthProvider = &oauthProvider.String
	}

	if oauthID.Valid {
		user.OauthID = &oauthID.String
	}
	return &user, nil
}

// GetByID mengambil user berdasarkan ID
func (r *userRepository) GetByID(id int) (*model.User, error) {
	query := `
		SELECT id, email, name, password, role, oauth_provider, oauth_id, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	var user model.User
	var password sql.NullString
	var oauthProvider sql.NullString
	var oauthID sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&password,
		&user.Role,
		&oauthProvider,
		&oauthID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	if password.Valid {
		user.Password = password.String
	}

	if oauthProvider.Valid {
		user.OauthProvider = &oauthProvider.String
	}

	if oauthID.Valid {
		user.OauthID = &oauthID.String
	}
	return &user, nil
}

// GetByOAuth mengambil user berdasarkan OAuth provider dan ID
func (r *userRepository) GetByOAuth(provider, oauthID string) (*model.User, error) {
	query := `
		SELECT id, email, name, password, role, oauth_provider, oauth_id, created_at, updated_at
		FROM users
		WHERE oauth_provider = ? AND oauth_id = ?
	`

	var user model.User
	var password sql.NullString
	var oauthProvider sql.NullString
	var oauthIDField sql.NullString

	err := r.db.QueryRow(query, provider, oauthID).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&password,
		&user.Role,
		&oauthProvider,
		&oauthIDField,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by OAuth: %w", err)
	}

	if password.Valid {
		user.Password = password.String
	}

	if oauthProvider.Valid {
		user.OauthProvider = &oauthProvider.String
	}

	if oauthIDField.Valid {
		user.OauthID = &oauthIDField.String
	}
	return &user, nil
}

// GetAll mengambil semua user dengan pagination
func (r *userRepository) GetAll(page, limit int) ([]model.User, int, error) {
	// Hitung offset
	offset := (page - 1) * limit

	query := `
		SELECT id, email, name, role, oauth_provider, oauth_id, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		var oauthProvider sql.NullString
		var oauthID sql.NullString

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Role,
			&oauthProvider,
			&oauthID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan user: %w", err)
		}

		if oauthProvider.Valid {
			user.OauthProvider = &oauthProvider.String
		}

		if oauthID.Valid {
			user.OauthID = &oauthID.String
		}

		users = append(users, user)
	}

	// Hitung total user
	var total int
	countQuery := "SELECT COUNT(*) FROM users"
	err = r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	return users, total, nil
}

// Update mengupdate user
func (r *userRepository) Update(user *model.User) error {
	query := `
		UPDATE users
		SET name = ?, role = ?, oauth_provider = ?, oauth_id = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err := r.db.Exec(query, user.Name, user.Role, user.OauthProvider, user.OauthID, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// Delete menghapus user
func (r *userRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
