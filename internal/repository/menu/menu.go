package menu

import (
	"context"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	menuData := make([]model.MenuItem, 0)

	if err := m.db.WithContext(ctx).Where(model.MenuItem{Type: constant.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}

func (m *menuRepo) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenu")
	defer span.End()

	var menuData model.MenuItem

	if err := m.db.WithContext(ctx).Where(model.MenuItem{OrderCode: orderCode}).First(&menuData).Error; err != nil {
		return menuData, err
	}

	return menuData, nil
}
