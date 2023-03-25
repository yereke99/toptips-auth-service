package service

import (
	"TopTips/repository"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthService interface {
	Check(contact, role string) (bool, error)
	GiveTokenService(contact, role string) (string, error)
	Create(contact string, code int) error
	ValidateSMS(contact string, code_validate int) (bool, error)
	Login(contact, password string) (bool, error)
	//CheckTokenDriver(tokens string) (*models.DriverModel, error)
	//CheckTokenUser(token string) (*models.UserModel, error)
}

type authService struct {
	db repository.Database
}

func NewAuthService(ds repository.Database) *authService {
	return &authService{
		db: ds,
	}
}

func (s *authService) Check(contact, role string) (bool, error) {
	res, err := s.db.CheckFromRepo(contact, role)

	if err != nil {
		return false, err
	}

	if !res {
		return false, nil
	}

	return true, nil
}

func (s *authService) GiveTokenService(contact, role string) (string, error) {
	token, err := s.db.GiveToken(contact, role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Create(contact string, code int) error {
	e := s.db.Create(contact, code)
	if e != nil {
		return e
	}

	//resp, err := http.Get(config.ConfigSMS(contact, code))
	resp, err := http.Get("")
	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))

	return nil
}

func (s *authService) ValidateSMS(contact string, code_validate int) (bool, error) {
	code, err := s.db.ValidateSMS(contact)

	if err != nil {
		return false, err
	}

	if code == 0 && err != nil {
		return false, errors.New("The code id 0.")
	}

	if code_validate != code {
		return false, nil
	}

	return true, nil
}

func (s *authService) Login(contact, password string) (bool, error) {
	pass, err := s.db.Login(contact)
	if err != nil {
		return false, err
	}

	if pass != password {
		return false, nil
	}

	return true, nil
}
