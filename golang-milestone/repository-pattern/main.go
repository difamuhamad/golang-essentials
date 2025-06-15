package main

import (
	"context"
	"fmt"
	"repository-pattern/database"
	"repository-pattern/model"
	"repository-pattern/repository"
)

func main() {
	ctx := context.Background()
	db := database.GetConnection()
	defer db.Close()

	// Dependency Injection Here
	commentRepo := repository.NewCommentRepository(db)

	// find all data comments
	comments, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, cm := range comments {
		fmt.Printf("data %d: %v\n", cm.Id, cm)
	}

	// find all data by id
	comment, err := commentRepo.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data : %v", comment)

	// insert data
	id, err := commentRepo.Insert(ctx, model.Comment{Email: "test@gmail.com", Comment: "komentar yuk"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("lastId: %v", id)
}
