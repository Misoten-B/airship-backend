//nolint:gosec // math/randでも問題なしと判断したため
package uniqueid

import (
	"io"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID(ms uint64, entropy io.Reader) (ulid.ULID, error) {
	return ulid.New(ms, entropy)
}

func MakeULID() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return NewULID(ms, entropy)
}
