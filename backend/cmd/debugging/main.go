package main

import (
	"backend/internals/database"
	"backend/internals/store"
	"context"
	"fmt"
)

func main() {
	db, _ := database.New(
		"postgresql://postgres:admin@localhost:5432/quikshop?sslmode=disable",
		30,
		30,
		"15m",
	)

	s := store.NewUserStore(db)

	// ? Create User
	// user := models.UserDetail{
	// 	Email:    "a@b.c",
	// 	Password: "pass",
	// 	Role:     "Customer",
	// }
	// fmt.Println("Before : ", user)
	// s.Create(context.Background(), &user)
	// fmt.Println("")
	// fmt.Println("After : ", user)

	// ? Get User by Email
	// user, err := s.GetByEmail(context.Background(), "a@b.cc")
	// fmt.Println(user, err)

	// ? Get User by ID
	user, err := s.GetByUserID(context.Background(), 1)
	fmt.Println(user, err)
}
