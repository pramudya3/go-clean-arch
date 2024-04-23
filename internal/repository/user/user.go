package user

import (
	"context"
	"fmt"

	"github.com/pramudya3/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	*mongo.Database
}

func (u *userRepository) InsertUser(ctx context.Context, user *domain.User) error {
	_, err := u.Database.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed insert one to database: %v", err.Error())
	}

	return nil
}

func (u *userRepository) FetchAllUsers(ctx context.Context) ([]*domain.User, error) {
	collection := u.Database.Collection("users")

	users := []*domain.User{}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := domain.User{}

	err := u.Database.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("no records found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to find email: %v, err: %v", email, err.Error())
	}

	return &user, nil
}

func (u *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user := domain.User{}

	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed get object id, %w", err)
	}

	err = u.Database.Collection("users").FindOne(ctx, bson.M{"_id": idObject}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("no records found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to find id: %v, err: %v", id, err.Error())
	}

	return &user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	coll := u.Database.Collection("users")

	idObject, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		return fmt.Errorf("failed get object id: %v", err)
	}

	updated := bson.M{"$set": user}

	_, err = coll.UpdateOne(ctx, bson.M{"_id": idObject}, updated)
	if err != nil {
		return fmt.Errorf("failed to update user, with id: %s, err: %v", idObject, err)
	}
	return nil
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		Database: db,
	}
}
