package repo

import (
	"context"
	"fmt"
	"mongodb-api/model"
	"mongodb-api/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserRepoImpl struct {
	db *mongo.Database
	ctx context.Context
}

func CreateUserRepoImpl(db *mongo.Database, ctx context.Context) user.UserRepo {
	return &UserRepoImpl{db, ctx}
}

func (e *UserRepoImpl) InsertUser(user *model.User) (*model.User, error) {
	_, err := e.db.Collection("user").InsertOne(e.ctx, user)
	if err != nil {
		fmt.Printf("error insert user %v \n", err)
		return nil, fmt.Errorf("Failed add data user")
	}
	return user, nil
}

func (e *UserRepoImpl) FindAllUser() (*[]model.User, error) {
	res, err := e.db.Collection("user").Find(e.ctx, bson.D{})
	if err != nil {
		fmt.Printf("error find data user %v \n", err)
		return nil, fmt.Errorf("failed find data")
	}
	defer res.Close(e.ctx)
	var user []model.User
	for res.Next(e.ctx) {
		var data model.User
		err := res.Decode(&data)
		if err != nil {
			fmt.Printf("error find data decode %v \n", err)
			return nil, fmt.Errorf("failed find data")
		}
		user = append(user, data) 
	}
	return &user, nil
}

func (e *UserRepoImpl) FindUser(id string) (*model.User, error) {
	var user model.User
	ids, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = e.db.Collection("user").FindOne(e.ctx, bson.M{"_id": ids}).Decode(&user)
	if err != nil {
		fmt.Printf("error find user id %v \n", err)
		return nil, fmt.Errorf("User not exist")
	}
	return &user, nil
}

func (e *UserRepoImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	ids, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = e.db.Collection("user").UpdateOne(e.ctx, bson.M{"_id": ids}, bson.M{"$set": user})
	if err != nil {
		fmt.Printf("error update %v \n", err)
		return nil, fmt.Errorf("Error update user")
	}
	user, err = e.FindUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *UserRepoImpl) DeleteUser(id string) error {
	ids, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = e.db.Collection("user").DeleteOne(e.ctx, bson.M{"_id": ids})
	if err != nil {
		fmt.Printf("error deleted data %v \n", err)
		return fmt.Errorf("Id user is not exist")
	}
	return nil
}