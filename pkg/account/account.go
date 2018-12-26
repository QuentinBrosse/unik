package account

import (
	"log"
	"regexp"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/quentinbrosse/scwgame/pkg/std/db"
	"github.com/quentinbrosse/scwgame/pkg/std/jwt"
	"github.com/quentinbrosse/scwgame/pkg/std/password"
	pb "github.com/quentinbrosse/scwgame/protobufs/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Token struct {
	jwtGo.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Username string `gorm:"type:varchar(100);unique;not null"`
	Password string
	Token    string `gorm:"-"`
}

var emailRegex *regexp.Regexp

func init() {
	db.GetDatabase().AutoMigrate(&Account{})
	emailRegex = regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}`)
}

func (a *Account) Create() error {
	database := db.GetDatabase()

	err := a.ValidateCreation()
	if err != nil {
		return err
	}

	hashedPass, err := password.HashPassword(a.Password)
	if err != nil {
		return UnknownError
	}

	a.Password = string(hashedPass)

	existingAccount := &Account{}
	err = database.
		Where("email = ?", a.Email).Or("username = ?", a.Username).
		First(existingAccount).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Printf("error: cannot check user existence (%s)", err)
		return UnknownError
	}
	if existingAccount.ID > 0 {
		log.Printf("error: %s", UserAlreadyExists.Error())
		return UserAlreadyExists
	}

	err = database.Create(a).Error
	if err != nil {
		log.Printf("error: cannot create account (%s)", err)
		return UnknownError
	}

	log.Printf("info: new user created: #%v %s", a.ID, a.Username)

	jwtToken, err := jwt.NewSigned(a.ID, []string{jwt.AudienceGame})
	if err != nil {
		return UnknownError
	}

	a.Token = jwtToken

	return nil
}

func (a *Account) SignIn() error {
	database := db.GetDatabase()

	err := a.ValidateSignIn()
	if err != nil {
		return err
	}

	candidate := &Account{}
	err = database.Debug().
		Where("email = ?", a.Email).
		First(candidate).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Printf("info: %s", InvalidCredentials.Error())
			return InvalidCredentials
		}
		log.Printf("error: cannot check user existence (%s)", err)
		return UnknownError
	}

	if !password.CompareHashAndPassword(candidate.Password, a.Password) {
		log.Printf("info: %s", InvalidCredentials.Error())
		return InvalidCredentials
	}

	*a = *candidate

	log.Printf("info: signed in user: #%v %s", a.ID, a.Username)

	jwtToken, err := jwt.NewSigned(a.ID, []string{jwt.AudienceGame})
	if err != nil {
		return UnknownError
	}

	a.Token = jwtToken

	return nil
}

func (a *Account) ValidateCreation() error {
	if a.Email == "" {
		return status.Errorf(codes.InvalidArgument, "missing email")
	}

	if !emailRegex.MatchString(a.Email) {
		return status.Errorf(codes.InvalidArgument, "invalid email")
	}

	if a.Username == "" {
		return status.Errorf(codes.InvalidArgument, "missing username")
	}

	if a.Password == "" {
		return status.Errorf(codes.InvalidArgument, "missing password")
	}

	if len(a.Password) < 6 {
		return status.Errorf(codes.InvalidArgument, "invalid password (len < 6)")
	}

	return nil
}

func (a *Account) ValidateSignIn() error {
	if a.Email == "" {
		return status.Errorf(codes.InvalidArgument, "missing email")
	}

	if !emailRegex.MatchString(a.Email) {
		return status.Errorf(codes.InvalidArgument, "invalid email")
	}

	if a.Password == "" {
		return status.Errorf(codes.InvalidArgument, "missing password")
	}

	if len(a.Password) < 6 {
		return InvalidCredentials
	}

	return nil
}

func (a *Account) ToPb() *pb.Account {
	return &pb.Account{
		Id:       uint32(a.ID),
		Email:    a.Email,
		Username: a.Username,
		Token:    a.Token,
	}
}
