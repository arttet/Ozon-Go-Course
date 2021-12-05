package async

import (
	"context"
	"testing"
	"time"
)

func TestAsyncSet(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, to)

	err := ac.Add(ctx, "k", "v")
	if err != ErrTimeout {
		t.Error("Expected timeout")
	}

	to = time.Millisecond * 3
	ctx, _ = context.WithTimeout(ctxBase, to)

	err = ac.Add(ctx, "k", "v")
	if err != nil {
		t.Errorf("Expected Set, %v", err)
	}
}

func TestAsyncGet(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond
	key := "k"
	value1 := "v"

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, to)

	_ = ac.Add(ctxBase, key, value1)
	v, err := ac.Get(ctx, key)
	if err != ErrTimeout {
		t.Error("Expected timeout")
	}

	ctx, _ = context.WithTimeout(ctxBase, to * 5)
	v, err = ac.Get(ctx, key)
	if err != nil {
		t.Error("Expected Get")
	}
	if v != value1 {
		t.Errorf("Expected <%v>, got <%v>", value1, v)
	}

	ctx, _ = context.WithTimeout(ctxBase, time.Microsecond)
	_ = ac.Add(ctx, "check", "ups")
	time.Sleep(time.Second)
	v, err = ac.Get(context.Background(), "check")
	if err == ErrTimeout {
		t.Errorf("Must Get, %v", err)
	}
	if v != "" {
		t.Errorf("Get should return nothing, got %v", v)
	}
}
