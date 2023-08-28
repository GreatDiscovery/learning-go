package log

import (
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestLogWithRequestId(t *testing.T) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return
	}
	log.Println("requestId={}, hello", newUUID)
}
