package dao

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"productAccounting-v1/cmd/admin-api/storage/driver"
	"productAccounting-v1/internal/domain/entity"
	"productAccounting-v1/internal/domain/enum"
)

type AuthStorage struct {
	driver *driver.Neo4jDriver
}

func NewAuthStorage(
	driver *driver.Neo4jDriver) *AuthStorage {
	return &AuthStorage{
		driver: driver,
	}
}

func (s *AuthStorage) CreateUser(user *entity.User) error {
	session := s.driver.GetSession()
	defer session.Close()

	newGuid := uuid.New()
	_, err := session.Run("CREATE (:User {guid: &guid, email: &email, password: &password, role: &role})", map[string]interface{}{
		"guid":     newGuid.String(),
		"email":    user.Email,
		"password": user.Password,
		"role":     user.Role,
	})

	if err != nil {
		return err
	}

	user.ID = newGuid
	return nil
}

func (s *AuthStorage) GetUser(email string, password string) (*entity.User, error) {
	session := s.driver.GetSession()
	defer session.Close()

	result, err := session.Run("MATCH (u:User) WHERE u.email = &email AND u.password = &password RETURN u", map[string]interface{}{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	var user *entity.User
	if result.Next() {
		prop := result.Record().Values[0].(neo4j.Node).Props
		guid, _ := uuid.Parse(prop["guid"].(string))

		user = &entity.User{
			ID:       guid,
			Email:    prop["email"].(string),
			Password: prop["password"].(string),
			Role:     enum.ParseRoles(prop["role"].(string)),
		}
	}

	return user, nil
}
