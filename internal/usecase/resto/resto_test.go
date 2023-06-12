package resto

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/mocks"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/menu"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/order"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/user"
	"github.com/golang/mock/gomock"
)

func Test_restoUsecase_GetMenuList(t *testing.T) {
	type fields struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		{
			name: "success get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return([]model.MenuItem{
							{
								OrderCode: "a",
								Name:      "test product 1",
								Price:     10000,
								Type:      constant.MenuTypeFood,
							},
						}, nil)

					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want: []model.MenuItem{
				{
					OrderCode: "a",
					Name:      "test product 1",
					Price:     10000,
					Type:      constant.MenuTypeFood,
				},
			},
			wantErr: false,
		},
		{
			name: "fail get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return(nil, errors.New("mock error"))

					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &restoUsecase{
				menuRepo:  tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
				userRepo:  tt.fields.userRepo,
			}
			got, err := m.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.GetMenuList() = %v, want %v", got, tt.want)
			}
		})
	}
}
