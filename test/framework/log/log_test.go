package log

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"testing"
)

//https://www.cnblogs.com/jiujuan/p/15542743.html

func TestLogrus(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")
}

func TestLogWithRequestId(t *testing.T) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return
	}
	log.Println("requestId={}, hello", newUUID)
}
