package main

import (
	"github.com/upyun/go-sdk/upyun"
	"path/filepath"
	"time"
)

const (
	TIME_NOT_SET = iota
	TIME_BEFORE
	TIME_AFTER
	TIME_INTERVAL
)

const (
	ITEM_NOT_SET = iota
	DIR
	FILE
)

type MatchConfig struct {
	Wildcard string

	TimeType int
	Before   time.Time
	After    time.Time

	ItemType int
}

func IsMatched(upInfo *upyun.FileInfo, mc *MatchConfig) bool {
	if mc.Wildcard != "" {
		if same, _ := filepath.Match(mc.Wildcard, upInfo.Name); !same {
			return false
		}
	}

	switch mc.TimeType {
	case TIME_BEFORE:
		if !upInfo.Time.Before(mc.Before) {
			return false
		}
	case TIME_AFTER:
		if !upInfo.Time.After(mc.After) {
			return false
		}
	case TIME_INTERVAL:
		if !upInfo.Time.Before(mc.Before) {
			return false
		}
		if !upInfo.Time.After(mc.After) {
			return false
		}
	}

	switch mc.ItemType {
	case DIR:
		if !upInfo.IsDir {
			return false
		}
	case FILE:
		if upInfo.IsDir {
			return false
		}
	}

	return true
}
