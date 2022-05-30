package service

import (
	"fmt"

	r "github.com/rodrigoherera/game-jobs/src/job/repository/interfaces"
)

type CharacterHandler struct {
	characterRepo r.CharacterRepository
}

func NewCharacterHandler(characterRepo r.CharacterRepository) *CharacterHandler {
	return &CharacterHandler{
		characterRepo: characterRepo,
	}
}

func (ch *CharacterHandler) RestoreLife() {
	fmt.Println("RestoreLife Task initiate")
	result, err := ch.characterRepo.RestoreLife()

	fmt.Printf("RestoreLife rows affected: %d", result)
	fmt.Println()

	if err != nil {
		fmt.Printf("RestoreLife ERROR: %v", err.Error())
	}
	fmt.Println("RestoreLife Task finish")
}

func (ch *CharacterHandler) RestoreStamina() {
	fmt.Println("RestoreStamina Task initiate")
	result, err := ch.characterRepo.RestoreStamina()

	fmt.Printf("RestoreLife rows affected: %d", result)
	fmt.Println()

	if err != nil {
		fmt.Printf("RestoreLife ERROR: %v", err.Error())
	}
	fmt.Println("RestoreStamina Task finish")
}
