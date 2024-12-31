package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goTest/initialize"
	"goTest/internal/model"
	"goTest/internal/model/baseModel"
	"goTest/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var db = initialize.Db.Collection("Users")
var rClient0 = initialize.RClient0

func GetAllUsers(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {
	var userList []model.User
	findRedis, err := rClient0.Get(c, "getUsers").Result()
	if err != nil {

		find, _ := db.Find(c, bson.D{})
		err = find.All(c, &userList)
		if len(userList) == 0 {
			return nil, utils.NotFoundMapper("user not found")
		}

		data, err := json.Marshal(userList)
		if err != nil {
			return nil, utils.InternalServerErrorMapper(err.Error())
		}

		_, err = rClient0.Set(c, "getUsers", data, 30*time.Second).Result()
		if err != nil {
			return nil, utils.InternalServerErrorMapper(err.Error())
		}

	} else {

		err := json.Unmarshal([]byte(findRedis), &userList)
		if err != nil {
			return nil, utils.InternalServerErrorMapper(err.Error())
		}

	}
	return utils.SuccessMapper(userList), nil
}

func GetUser(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {

	var user model.User
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	err = db.FindOne(c, bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NotFoundMapper(err.Error())
		} else {
			return nil, utils.InternalServerErrorMapper(err.Error())
		}
	}

	return utils.SuccessMapper(user), nil
}

func GetUserByEmail(c *gin.Context, email string) (user *model.User, err error) {

	err = db.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		} else {
			return nil, err
		}
	}

	return user, nil

}

func UpdateUser(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, utils.BadRequestMapper(err.Error())
	}

	user.UpdatedAt = time.Now()

	update := bson.M{"$set": user}
	oId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	data, err := db.UpdateOne(c, bson.M{"_id": oId}, update)
	if data.ModifiedCount == 0 {
		return nil, utils.NotFoundMapper(err.Error())
	}

	_, err = rClient0.Del(c, "getUsers").Result()
	if err != nil {
		return nil, utils.InternalServerErrorMapper(err.Error())
	}

	user.ID = oId
	return utils.SuccessMapper(user), nil
}

func DeleteUser(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {
	oId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	data, err := db.DeleteOne(c, bson.M{"_id": oId})
	if err != nil {
		return nil, utils.InternalServerErrorMapper(err.Error())
	}

	return utils.SuccessMapper(data), nil
}

func CreateUser(c *gin.Context) (sResponse *baseModel.SuccessResponse, eResponse *baseModel.ErrorResponse) {
	var user model.User
	err := c.BindJSON(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.InternalServerErrorMapper(err.Error())
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	one, err := db.InsertOne(c, &user)
	if err != nil {
		return nil, utils.InternalServerErrorMapper(err.Error())
	}

	_, err = rClient0.Del(c, "getUsers").Result()
	if err != nil {
		return nil, utils.InternalServerErrorMapper(err.Error())
	}

	return utils.CreatedMapper(one), nil
}
