package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/Zaanter/be-goserver/pkg/entities"
	"github.com/Zaanter/be-goserver/pkg/utils"
)

type UserRepository interface {
	FindAll() ([]entities.IUser, error)
	GetUser(uid string) (*entities.IUser, error)
	Create(user *entities.IUser) (*entities.IUser, error)
	Delete(uid string) error
	// Update(user *entities.IUser) (*entities.IUser, error)
}

type repo struct{}

func NewUserRepository() UserRepository {
	return &repo{}
}

const (
	projectId       string = "course-web-a2a46"
	USER_COLLECTION string = "users"
)

func (*repo) Create(user *entities.IUser) (*entities.IUser, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	userDate, _ := utils.ParseDate(user.Dob)
	argLocation, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	docId, _, err := client.Collection(USER_COLLECTION).Add(ctx, map[string]interface{}{
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
		"email":     user.Email,
		"dob": time.Date(userDate.Year, time.Month(userDate.Month),
			userDate.Day, 0, 0, 0, 0, argLocation),
		"deleted": false,
	})
	if err != nil {
		log.Fatalf("Failed to create a new user: %v", err)
		return nil, err
	}

	fmt.Printf("Creating new user with uid %v\n", docId.ID)
	user.Uid = docId.ID
	return user, nil
}

func (*repo) FindAll() ([]entities.IUser, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	iter := client.Collection(USER_COLLECTION).Documents(ctx)
	var users []entities.IUser
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of users: %v", err)
			return nil, err
		}
		user := entities.IUser{
			Uid:       doc.Ref.ID,
			Firstname: doc.Data()["firstname"].(string),
			Lastname:  doc.Data()["lastname"].(string),
			Email:     doc.Data()["email"].(string),
			Dob:       doc.Data()["dob"].(time.Time).Local().String(),
			Deleted:   doc.Data()["deleted"].(bool),
		}

		users = append(users, user)
	}
	return users, nil
}

func (*repo) GetUser(uid string) (*entities.IUser, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	userDoc, _ := client.Collection(USER_COLLECTION).Doc(uid).Get(ctx)
	user := entities.IUser{
		Uid:       userDoc.Ref.ID,
		Firstname: userDoc.Data()["firstname"].(string),
		Lastname:  userDoc.Data()["lastname"].(string),
		Email:     userDoc.Data()["email"].(string),
		Dob:       userDoc.Data()["dob"].(time.Time).Local().String(),
		Deleted:   userDoc.Data()["deleted"].(bool),
	}

	return &user, nil
}

func (*repo) Delete(uid string) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return err
	}
	defer client.Close()

	_, deletedError := client.Collection(USER_COLLECTION).Doc(uid).Delete(ctx)
	if deletedError != nil {
		log.Fatalf("failed deleting user: %v", err)
		return err
	}

	return nil
}
