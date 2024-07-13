package event

import (
	"context"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/event"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/pkg/errors"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	formattedDate, _ := time.Parse("2006-01-02T15:04:05.000Z", "2024-07-07T00:00:00.000Z")
	success := map[string]struct {
		events model.EventCreate
		event1 model.Event
		event2 model.Event
		uid    string
		err    error
	}{
		"イベントが1件の場合": {
			events: model.EventCreate{
				Amount1:   100,
				Amount2:   0,
				Category1: 1,
				Category2: 1,
				Memo1:     "ファミマで買い物",
				Memo2:     "",
				StoreName: "ファミマ",
				Date:      "2024-07-07T00:00:00.000Z",
			},
			event1: model.Event{
				Amount:    100,
				Category:  1,
				Memo:      "ファミマで買い物",
				StoreName: "ファミマ",
				Date:      formattedDate,
			},
			uid: "testuid",
			err: nil,
		},
		"イベントが1件でメモと店名が空欄の場合": {
			events: model.EventCreate{
				Amount1:   100,
				Amount2:   0,
				Category1: 1,
				Category2: 1,
				Memo1:     "",
				Memo2:     "",
				StoreName: "",
				Date:      "2024-07-07T00:00:00.000Z",
			},
			event1: model.Event{
				Amount:    100,
				Category:  1,
				Memo:      "",
				StoreName: "",
				Date:      formattedDate,
			},
			uid: "testuid",
			err: nil,
		},
		"イベントが2件の場合": {
			events: model.EventCreate{
				Amount1:   100,
				Amount2:   200,
				Category1: 1,
				Category2: 2,
				Memo1:     "ファミマで買い物",
				Memo2:     "ファミマで他にも買ったよ",
				StoreName: "ファミマ",
				Date:      "2024-07-07T00:00:00.000Z",
			},
			event1: model.Event{
				Amount:    100,
				Category:  1,
				Memo:      "ファミマで買い物",
				StoreName: "ファミマ",
				Date:      formattedDate,
			},
			event2: model.Event{
				Amount:    200,
				Category:  2,
				Memo:      "ファミマで他にも買ったよ",
				StoreName: "ファミマ",
				Date:      formattedDate,
			},
			uid: "testuid",
			err: nil,
		},
	}

	fail := map[string]struct {
		events  model.EventCreate
		uid     string
		wantErr error
	}{
		"金額1が不正な値の場合": {
			events: model.EventCreate{
				Amount1:   -100,
				Amount2:   200,
				Category1: 1,
				Category2: 2,
				Memo1:     "ファミマで買い物",
				Memo2:     "ファミマで他にも買ったよ",
				StoreName: "ファミマ",
				Date:      "2024-07-07T00:00:00.000Z",
			},
			uid:     "testuid",
			wantErr: errors.BadRequest,
		},
		"カテゴリーの値が不正な値の場合": {
			events: model.EventCreate{
				Amount1:   100,
				Amount2:   200,
				Category1: 10,
				Category2: 2,
				Memo1:     "ファミマで買い物",
				Memo2:     "ファミマで他にも買ったよ",
				StoreName: "ファミマ",
				Date:      "2024-07-07T00:00:00.000Z",
			},
			uid:     "testuid",
			wantErr: errors.BadRequest,
		},
	}

	// Mock設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := event.NewMockEventRepository(ctrl)
	mockTransaction := transaction.NewMockTransactionRepository(ctrl)
	mock := New(mockRepo, mockTransaction)

	t.Run("成功", func(t *testing.T) {
		for tn, tt := range success {
			t.Run(tn, func(t *testing.T) {
				mockRepo.EXPECT().Create(tt.event1, tt.uid).Return(tt.err)
				if tt.events.Amount2 > 0 {
					mockRepo.EXPECT().Create(tt.event2, tt.uid).Return(tt.err)
				}
				// トランザクション内での期待動作を設定
				mockTransaction.EXPECT().Transaction(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(*sqlx.Tx) error) error {
						tx := &sqlx.Tx{} // 必要に応じて適切なモックトランザクションを作成
						return fn(tx)
					},
				)
				err := mock.Create(tt.events, tt.uid)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("失敗", func(t *testing.T) {
		for tn, tt := range fail {
			t.Run(tn, func(t *testing.T) {
				err := mock.Create(tt.events, tt.uid)
				assert.Equal(t, err, tt.wantErr)
			})
		}
	})
}
