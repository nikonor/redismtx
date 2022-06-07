package redismtx

import (
	"testing"
	"time"
	
	"github.com/go-redis/redis"
)

func TestName(t *testing.T) {
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	
	o := Init(rc, "redismtx", time.Second)
	
	if err := rc.Del("redismtx::test").Err(); err != nil && err != redis.Nil {
		t.Fatal("0::" + err.Error())
	}
	
	got, err := o.Lock("test")
	println("#1", got)
	if err != nil {
		t.Fatal("#1.1::" + err.Error())
	}
	if !got {
		t.Fatal("#1.2 want=true, got=false")
	}
	
	got, err = o.Lock("test")
	println("#2", got)
	if err != nil {
		t.Fatal("#2.1::" + err.Error())
	}
	if got {
		t.Fatal("#2.2 want=false, got=true")
	}
	
	got, err = o.Lock("test")
	println("#3", got)
	if err != nil {
		t.Fatal("#3.1::" + err.Error())
	}
	if got {
		t.Fatal("#3.2 want=false, got=true")
	}
	
	if err = o.UnLock("test"); err != nil {
		t.Fatal("#4::" + err.Error())
	}
	
	got, err = o.Lock("test")
	println("#5", got)
	if err != nil {
		t.Fatal("#5.1::" + err.Error())
	}
	if !got {
		t.Fatal("#5.2 want=true, got=false")
	}
	
	if err = o.UnLock("test"); err != nil {
		t.Fatal("#6::" + err.Error())
	}
	
}
