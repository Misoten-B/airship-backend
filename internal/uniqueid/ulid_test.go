package uniqueid_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Misoten-B/airship-backend/internal/uniqueid"
)

func TestULID(t *testing.T) {
	now := time.Now()
	before := now.Add(-1 * time.Hour)
	after := now.Add(1 * time.Hour)

	nowULID, err := uniqueid.MakeULID(now)
	if err != nil {
		t.Fatal(err)
	}

	beforeULID, err := uniqueid.MakeULID(before)
	if err != nil {
		t.Fatal(err)
	}

	afterULID, err := uniqueid.MakeULID(after)
	if err != nil {
		t.Fatal(err)
	}

	result := strings.Compare(nowULID.String(), beforeULID.String())
	if result != 1 {
		t.Errorf("MakeULID(now) = %s, MakeULID(before) = %s", nowULID.String(), beforeULID.String())
	}

	result = strings.Compare(nowULID.String(), afterULID.String())
	if result != -1 {
		t.Errorf("MakeULID(now) = %s, MakeULID(after) = %s", nowULID.String(), afterULID.String())
	}
}
