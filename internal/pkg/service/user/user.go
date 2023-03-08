package user

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/ports"
	"github.com/OscarLlamas6/hex-arch-go/settings"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service struct {
		repo ports.UserRepository
	}
)

func NewService(repo ports.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(user *entity.User) error {

	// Is valid email address
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	hashedPass := hashAndSalt(user.Password)
	fmt.Printf("Hashed password: %s\n", hashedPass)
	user.Password = hashedPass

	return s.repo.Create(user)
}

func (s *Service) Login(credentials *entity.DefaultCredentials) (string, error) {
	user := &entity.User{}

	if err := s.repo.First(user, "email = ?", credentials.Email); err != nil {
		return "", err
	}

	if err := tryMatchPassword(user.Password, credentials.Password); err != nil {
		return "", err
	}

	return createToken(user)
}

func tryMatchPassword(userPassword, credentialPassword string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentialPassword)); err != nil {
		return err
	}

	return nil
}

func hashAndSalt(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func createToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = user.ID
	claims["userName"] = user.Name
	claims["userEmail"] = user.Email
	claims["role"] = "standar"
	claims["exp"] = time.Hour * 24

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := jwtToken.SignedString([]byte(settings.AppConfig.JWTKey))
	if err != nil {
		return "", err
	}

	return strToken, nil
}
