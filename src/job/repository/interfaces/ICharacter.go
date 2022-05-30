package interfaces

type CharacterRepository interface {
	RestoreLife() (int64, error)
	RestoreStamina() (int64, error)
}
