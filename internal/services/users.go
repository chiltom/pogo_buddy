package services

import (
  "errors"
  "github.com/chiltom/pogo_buddy/internal/db"
  "github.com/chiltom/pogo_buddy/internal/models"

  "golang.org/x/crypto/bcrypt"
)

type UserService struct {
  store *db.UserStore
}

func NewUserService(store *db.UserStore) *UserService {
  return &UserService{store: store}
}

func (s *UserService) Create(user *models.User) (int, error) {
  if user.Email == "" || user.Password == "" {
    return 0, errors.New("email and password are required")
  }
  
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    return 0, err
  }
  user.Password = string(hashedPassword)

  return s.store.Create(*user)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
  if email == "" {
    return nil, errors.New("email is required")
  }
  return s.store.GetByEmail(email)
}

func (s *UserService) Update(user *models.User) error {
  if user.ID == 0 {
    return errors.New("user id is required")
  }

  if user.Password != "" {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
      return err
    }
    user.Password = string(hashedPassword)
  }
  return s.store.Update(*user)
}

func (s *UserService) Delete(id int) error {
  if id == 0 {
    return errors.New("user id is required")
  }
  return s.store.Delete(id)
}

func (s *UserService) CheckPassword(email, password string) (bool, error) {
  user, err := s.store.GetByEmail(email)
  if err != nil {
    return false, err
  }
  if user == nil {
    return false, nil
  }
  return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil, nil
}
