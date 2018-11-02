package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseTimeString(input string) time.Duration {
	items := strings.Split(input, "+")
	var ret time.Duration = 0
	for _, item := range items {
		ret += parseTimeStringItem(item)
	}
	return ret
}

func parseTimeStringItem(input string) time.Duration {
	input = strings.ToLower(input)
	reg, err := regexp.Compile("^(\\d+)(s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days|w|week|weeks|y|year|years)?")
	if err != nil {
		panic(err)
	}
	ret := reg.FindAllSubmatch([]byte(input), 2)
	if len(ret) < 1 {
		return 0
	}
	n, err := strconv.ParseInt(string(ret[0][1]), 10, 32)
	if err != nil {
		return 0
	}
	t := string(ret[0][2])
	switch t {
	case "s":
		return time.Second * time.Duration(n)
	case "second":
		return time.Second * time.Duration(n)
	case "seconds":
		return time.Second * time.Duration(n)
	case "m":
		return time.Minute * time.Duration(n)
	case "minute":
		return time.Minute * time.Duration(n)
	case "minutes":
		return time.Minute * time.Duration(n)
	case "h":
		return time.Hour * time.Duration(n)
	case "hour":
		return time.Hour * time.Duration(n)
	case "hours":
		return time.Hour * time.Duration(n)
	case "d":
		return time.Hour * 24 * time.Duration(n)
	case "day":
		return time.Hour * 24 * time.Duration(n)
	case "days":
		return time.Hour * 24 * time.Duration(n)
	case "w":
		return time.Hour * 24 * 7 * time.Duration(n)
	case "week":
		return time.Hour * 24 * 7 * time.Duration(n)
	case "weeks":
		return time.Hour * 24 * 7 * time.Duration(n)
	case "y":
		return time.Hour * 24 * 365 * time.Duration(n)
	case "year":
		return time.Hour * 24 * 365 * time.Duration(n)
	case "years":
		return time.Hour * 24 * 365 * time.Duration(n)
	}
	return 0
}
