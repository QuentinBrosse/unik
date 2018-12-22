package account

import (
	"log"

	"github.com/quentinbrosse/scwgame/pkg/std/password"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/quentinbrosse/scwgame/pkg/std/db"
	"github.com/quentinbrosse/scwgame/pkg/std/jwt"
	pb "github.com/quentinbrosse/scwgame/protobufs/account"
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

func init() {
	db.GetDatabase().AutoMigrate(&Account{})
}

func (a *Account) Create() error {
	database := db.GetDatabase()

	// TODO: Validate the request

	hashedPass, err := password.HashPassword(a.Password)
	if err != nil {
		return UnknownError
	}

	a.Password = string(hashedPass)

	var count int
	err = database.
		Model(&Account{}).
		Where("email = ?", a.Email).Or("username = ?", a.Username).
		Count(&count).Error
	if err != nil {
		log.Printf("error: cannot check user existence (%s)", err)
		return UnknownError
	}
	if count > 0 {
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

func (a *Account) ToPb() *pb.Account {
	return &pb.Account{
		Id:    uint32(a.ID),
		Email: a.Email,
		Token: a.Token,
	}
}
