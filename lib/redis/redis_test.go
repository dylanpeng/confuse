package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

var cachePool *Pool
var user *User
var ctx context.Context

type User struct {
	Name    string    `json:"name"`
	Gender  int       `json:"gender"`
	Courses []*Course `json:"courses"`
}

func (u *User) String() string {
	return fmt.Sprintf("%+v", *u)
}

type Course struct {
	CourseName string `json:"course_name"`
	Score      int    `json:"score"`
}

func (c *Course) String() string {
	return fmt.Sprintf("%+v", *c)
}

func GetKey() string {
	return "user:test:course"
}

func TestInitPool(t *testing.T) {
	cachePool = NewPool()

	conf := &Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "",
	}

	cachePool.Add("test", conf)

	user = &User{
		Name:    "user",
		Gender:  1,
		Courses: make([]*Course, 0),
	}

	user.Courses = append(user.Courses,
		&Course{
			CourseName: "english",
			Score:      80,
		},
		&Course{
			CourseName: "math",
			Score:      95,
		})

	ctx = context.TODO()
}

func TestNewPool(t *testing.T) {
	cache, err := cachePool.Get("test")

	if err != nil {
		t.Fatalf("get cache failed. err: %s", err)
	}

	userMarshal, err := json.Marshal(user)

	if err != nil {
		t.Fatalf("marshal user failed. err: %s", err)
	}

	result, err := cache.Set(ctx, GetKey(), userMarshal, 0).Result()

	if err != nil {
		t.Fatalf("set cache key: %s faile. | result: %s | err: %s", GetKey(), result, err)
	}

	t.Logf("set user success. result: %s", result)

	cacheUser := &User{}

	result, err = cache.Get(ctx, GetKey()).Result()

	if err != nil {
		t.Fatalf("set cache key: %s faile. | result: %s | err: %s", GetKey(), result, err)
	}

	err = json.Unmarshal([]byte(result), cacheUser)

	if err != nil {
		t.Fatalf("convert fail. err: %s", err)
	}

	t.Logf("cache get success. result: %s", cacheUser)
}
