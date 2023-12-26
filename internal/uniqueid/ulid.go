//nolint:gosec // math/randでも問題なしと判断したため
package uniqueid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() (ulid.ULID, error) {
	time := time.Now()
	return MakeULID(time)
}

func MakeULID(t time.Time) (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	ms := ulid.Timestamp(t)
	return ulid.New(ms, entropy)
}
