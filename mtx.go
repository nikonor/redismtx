package redismtx

import (
	"time"
	
	"github.com/go-redis/redis"
)

type Typ struct {
	rc     *redis.Client
	prefix string
	maxTTL time.Duration
}

func Init(rc *redis.Client, prefix string, maxTTL time.Duration) *Typ {
	return &Typ{
		rc:     rc,
		prefix: prefix,
		maxTTL: maxTTL,
	}
}

// Lock - возврат false говорит, что ключ занят
func (t Typ) Lock(suffix string) (bool, error) {
	k := t.key(suffix)
	err := t.rc.GetSet(k, time.Now().String()).Err()
	switch {
	case err != nil && err == redis.Nil:
		return true, t.rc.Expire(k, t.maxTTL).Err()
	case err != nil && err != redis.Nil:
		return false, err
	default:
		return false, nil
	}
}

// UnLock - отпускаем ключ
func (t Typ) UnLock(suffix string) error {
	err := t.rc.Del(t.key(suffix), time.Now().String()).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return nil
}

func (t Typ) key(suffix string) string {
	return t.prefix + "::" + suffix
}
