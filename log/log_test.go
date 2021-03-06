package log

import (
	"fmt"
	"testing"
)

func TestLogLevel(t *testing.T) {
	const (
		msg1 = "log line1"
		msg2 = "log line2"
		msg3 = "log line3"
	)
	setMockLogger(fmt.Sprintf("%shello: %s", msg2, msg3), false)
	level := "info"
	SetLevel(level)
	if GetLevel() != level {
		t.Fatalf("Expected %q, got %q", level, GetLevel())
	}
	Debug.Println(msg1)             //not logged
	Info.Print(msg2)                //logged
	Error.Printf("hello: %s", msg3) //logged
	globals().defaultLogger.(*mockLogger).Verify(t)
}

type mockLogger struct {
	fatal         bool
	logged        string
	expected      string
	fatalExpected bool
}

func setMockLogger(expected string, fatalExpected bool) {
	state.defaultLogger = &mockLogger{
		expected:      expected,
		fatalExpected: fatalExpected,
	}
}

func (ml *mockLogger) Printf(format string, v ...interface{}) {
	ml.logged += fmt.Sprintf(format, v...)
}
func (ml *mockLogger) Print(v ...interface{}) {
	ml.logged += fmt.Sprint(v...)
}
func (ml *mockLogger) Println(v ...interface{}) {
	ml.logged += fmt.Sprintln(v...)
}
func (ml *mockLogger) Fatal(v ...interface{}) {
	ml.fatal = true
	ml.Print(v...)
}
func (ml *mockLogger) Fatalf(format string, v ...interface{}) {
	ml.fatal = true
	ml.Printf(format, v...)
}
func (ml *mockLogger) Verify(t *testing.T) {
	if ml.logged != ml.expected {
		t.Errorf("Expected %q, got %q", ml.expected, ml.logged)
	}
	if ml.fatal != ml.fatalExpected {
		t.Errorf("Expected fatal %v, got %v", ml.fatalExpected, ml.fatal)
	}
}
func (ml *mockLogger) Flush() {

}

func (ml *mockLogger) Log(l Level, s string) {
	ml.Print(s)
}
