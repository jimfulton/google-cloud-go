package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"cloud.google.com/go/internal/testutil"
	"github.com/google/go-cmp/cmp"
)

func captureStdout(f func()) string {
	saved := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = saved }()

	outC := make(chan string)
	// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	// back to normal state
	w.Close()
	return <-outC
}

func assertEqual(t *testing.T, got, want interface{}, opts ...cmp.Option) {
	if !testutil.Equal(got, want, opts...) {
		_, fpath, lno, ok := runtime.Caller(1)
		if ok {
			_, fname := filepath.Split(fpath)
			t.Errorf("%s:%d: Didn't match:\nGot: %s\nExpected: %s",
				fname, lno, got, want)
		} else {
			t.Errorf("Didn't match:\nGot: %s\nExpected: %s", got, want)
		}
	}
}

func assertNoError(t *testing.T, err error) bool {
	if err != nil {
		_, fpath, lno, ok := runtime.Caller(1)
		if ok {
			_, fname := filepath.Split(fpath)
			t.Errorf("%s:%d: %s", fname, lno, err)
		} else {
			t.Error(err)
		}
		return true
	}
	return false
}