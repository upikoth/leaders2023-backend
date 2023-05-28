package store

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type Score struct {
	tableName       struct{} `pg:"scores"` //nolint:unused // Имя таблицы
	Id              int      `pg:"id"`
	UserId          int      `pg:"user_id"`
	CreativeSpaceId int      `pg:"creative_space_id"`
	BookingId       int      `pg:"booking_id"`
	Comment         string   `pg:"comment"`
	Rating          int      `pg:"rating"`
}

func (s *Store) CreateScore(score Score) (int, error) {
	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// Записываем в таблицу score.
		result, scoreErr := tx.
			Model(&score).
			OnConflict("DO NOTHING").
			Insert()

		if scoreErr != nil {
			return scoreErr
		}

		if result.RowsAffected() == 0 {
			return constants.ErrScorePostDbError
		}

		return nil
	})

	if storeErr != nil {
		return 0, storeErr
	}

	return score.Id, nil
}
