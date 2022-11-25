package account

import (
	"context"
	"testing"
	"time"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/mock"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

func TestMockExample(t *testing.T) {
	s := merchantService{storage: repository.New()}
	// mock the merchant get db call
	s.storage.Merchant = mock.Merchant{
		GetFunc: func(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
			return types.Merchant{
				ID:        "test",
				Name:      "test name",
				Email:     "test@test.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil
		},
	}

	res, err := s.Get(context.Background(), &filters.Merchant{})
	if err != nil {
		t.Fatal(err)
	}
	if res.ID == "" {
		t.Fatal("missing merchant ID")
	}
}
