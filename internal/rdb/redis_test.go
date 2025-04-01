package rdb

import (
	"testing"
)

func setup() {
	ConnectRedis("10.123.1.31", 6379)
}

func TestSet(t *testing.T) {
	setup()

	Set("name", "zhangyan")

	value := Get("name")
	t.Log(value)

	_ = Del("name")
}
