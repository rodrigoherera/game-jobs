package repository

import "gorm.io/gorm"

type CharacterRepo struct {
	db *gorm.DB
}

func NewCharacterRepo(db *gorm.DB) *CharacterRepo {
	return &CharacterRepo{
		db: db,
	}
}

func (cr *CharacterRepo) RestoreLife() (int64, error) {
	tx := cr.db.Begin()

	result := tx.Table("characters").
		Where("life < max_life").
		UpdateColumn("life", gorm.Expr("if (life + 5 > max_life, life + (max_life - life), life + 5)"))

	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	tx.Commit()

	return result.RowsAffected, result.Error
}

func (cr *CharacterRepo) RestoreStamina() (int64, error) {
	/*tx := cr.db.Begin()

	result := tx.Table("characters").
		Where("life < max_life").
		UpdateColumn("life", gorm.Expr("life + ?", 5))

	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	tx.Commit()

	return result.RowsAffected, result.Error*/
	return 0, nil
}
