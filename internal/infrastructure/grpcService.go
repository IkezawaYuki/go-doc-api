package infrastructure

import (
	"context"
	"fmt"
	"github.com/IkezawaYuki/go-dog-api/pkg/pb/dog"
)

type DogService struct {
}

func (s *DogService) FindSmartDog(ctx context.Context, message *dog.FindSmartDogMessage) (*dog.FindSmartDogResponse, error) {
	switch message.GetDogId() {
	case "one":
		return &dog.FindSmartDogResponse{
			Name: "Number 1",
			Kind: "ゴールデンレトリパー",
		}, nil
	case "two":
		return &dog.FindSmartDogResponse{
			Name: "Number 2",
			Kind: "ブルドッグ",
		}, nil
	case "three":
		return &dog.FindSmartDogResponse{
			Name: "Number 3",
			Kind: "ビーグル",
		}, nil
	}
	return nil, fmt.Errorf("not found")
}
