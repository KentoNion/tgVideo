package domain

import (
	"context"

	"tgVideoCall/models"
)

type Storage interface {
	GetAdmin(ctx context.Context,userID int) (models.Admin, error)
}
