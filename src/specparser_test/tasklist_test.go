package specparser_test

import (
	"specparser"
	"testing"
	"time"
)

func TestNewTaskList(t *testing.T) {
	spec, err := specparser.NewTaskSpec("* * * * * command")
	_, err = specparser.NewTaskList(spec, time.Now(), 10)

	if err != nil {
		t.Error("failed initialization")
	}
}
