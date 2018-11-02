package main

import (
	"testing"
	"time"
)

func TestParseTimeString(t *testing.T) {
	{
		a := parseTimeString("123s")
		e := time.Second * 123
		if e != a {
			t.Errorf("parse 123s fail: %s != %s", e, a)
		}
	}
	{
		a := parseTimeString("123seconds")
		e := time.Second * 123
		if e != a {
			t.Errorf("parse 123seconds fail: %s != %s", e, a)
		}
	}
	{
		a := parseTimeString("3h")
		e := time.Hour * 3
		if e != a {
			t.Errorf("parse 3h fail: %s != %s", e, a)
		}
	}
	{
		a := parseTimeString("2y+3d+5h+20m+6s")
		e := time.Hour*(((365*2+3)*24)+5) + time.Minute*20 + time.Second*6
		if e != a {
			t.Errorf("parse 2y+3d+5h+20m+6s fail: %s != %s", e, a)
		}
	}
}
