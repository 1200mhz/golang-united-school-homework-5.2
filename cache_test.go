package cache

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCacheGetWithoutDeadline(t *testing.T) {
	cache := NewCache()

	key, value := "hello", "world"
	cache.Put(key, value)

	actual, ok := cache.Get(key)
	expected := value

	if !reflect.DeepEqual(ok, true) || !reflect.DeepEqual(actual, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}

func TestCacheGetWithActualDeadline(t *testing.T) {
	cache := NewCache()

	key, value := "hello", "world"
	cache.PutTill(key, value, time.Now().Add(time.Duration(1)*time.Minute))

	actual, ok := cache.Get(key)
	expected := value

	if !reflect.DeepEqual(ok, true) || !reflect.DeepEqual(actual, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}
func TestCacheGetWithNotActualDeadline(t *testing.T) {
	cache := NewCache()

	key, value := "hello", "world"
	cache.PutTill(key, value, time.Now().AddDate(0, 0, -1))

	actual, ok := cache.Get(key)
	expected := ""

	if !reflect.DeepEqual(ok, false) || !reflect.DeepEqual(actual, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}

func TestCacheGetKeys(t *testing.T) {
	cache := NewCache()

	key1, key2, key3, key4, key5 := "k1", "k2", "k3", "k4", "k5"

	cache.Put(key1, "some value")
	cache.PutTill(key2, "some value", time.Now().AddDate(0, 0, -1))
	cache.PutTill(key3, "some value", time.Now().AddDate(0, 0, -1))
	cache.Put(key4, "some value")
	cache.Put(key5, "some value")

	actual := cache.Keys()
	expected := []string{key1, key4, key5}

	if !reflect.DeepEqual(len(actual), len(expected)) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}
