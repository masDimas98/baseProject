package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"goTest/initialize"
	"goTest/internal/model"
	"goTest/internal/model/baseModel"
	"goTest/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var rClient1 = initialize.RClient1
var secretKey = []byte(os.Getenv("SECRET_KEY"))

func LoginHandler(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {

	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, utils.BadRequestMapper(err.Error())
	}

	findUser, err := GetUserByEmail(c, user.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.UserNotFoundMapper("invalid username or password")
		} else {
			return nil, utils.InternalServerErrorMapper(err.Error())
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password))
	if err != nil {
		return nil, utils.UnauthorizedMapper("invalid username or password")
	}
	tokenString, err := createToken(c, baseModel.Session{
		Name:      findUser.Name,
		Email:     findUser.Email,
		LoginTime: time.Now().Unix(),
		ExpiresIn: time.Now().Add(time.Minute * 5).Unix(),
		IPAddress: c.GetHeader("X-Forwarded-For"),
	})
	if err != nil {
		return nil, utils.InternalServerErrorMapper(err.Error())
	}
	return utils.SuccessMapper(model.TokenResponse{AccessToken: tokenString}), nil
}

func VerifyRequest(c *gin.Context) (eResponse *baseModel.ErrorResponse) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return utils.UnauthorizedMapper("invalid token")
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(c, tokenString)
	if err != nil {
		return utils.UnauthorizedMapper(err.Error())
	}

	return nil

}

func verifyToken(c *gin.Context, tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	session := token.Claims.(jwt.MapClaims)
	email, ok := session["email"].(string)
	if !ok {
		return errors.New("invalid token")
	}

	exp, ok := session["expiresIn"]
	if !ok {
		return errors.New("invalid token")
	}
	if int64(exp.(float64)) < time.Now().Unix() {
		return errors.New("token expired")
	}

	_, err = rClient1.Get(c, email).Result()
	if err != nil {
		return errors.New("logged out")
	}

	return nil
}

func createToken(c *gin.Context, session baseModel.Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":      session.Name,
			"email":     session.Email,
			"loginTime": session.LoginTime,
			"expiresIn": session.ExpiresIn,
			"ipAddress": session.IPAddress,
		})

	data, err := json.Marshal(session)

	_, err = rClient1.Set(c, session.Email, data, time.Minute*5).Result()
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Logout(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, utils.UnauthorizedMapper("invalid token")
	}
	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, utils.UnauthorizedMapper("invalid token")
	}

	session := token.Claims.(jwt.MapClaims)
	email, ok := session["email"].(string)
	if !ok {
		return nil, utils.UnauthorizedMapper("invalid token")
	}
	_, err = rClient1.Del(c, email).Result()
	if err != nil {
		return nil, utils.InternalServerErrorMapper("failed to delete session")
	}
	return utils.SuccessMapper("Log out successfully"), nil
}
