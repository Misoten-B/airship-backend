package mapper_test

import (
	"testing"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/mapper"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
)

func TestStatus_ToGormStatus(t *testing.T) {
	cases := []struct {
		name   string
		status shared.Status
		want   int
	}{
		{
			name:   "status is in progress",
			status: shared.StatusInProgress{},
			want:   model.GormStatusInProgress,
		},
		{
			name:   "status is completed",
			status: shared.StatusCompleted{},
			want:   model.GormStatusCompleted,
		},
	}

	t.Parallel()

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := mapper.ToGormStatus(tc.status)

			if got != tc.want {
				t.Errorf("got: %d, want: %d", got, tc.want)
			}
		})
	}
}
