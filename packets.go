package main

// Hey is first message packet through UDP multicast
type Hey struct {
	Hostname string `json:"hostname"`
	Tree     []Item `json:"tree"`
}

// ItemType is type for tree item
type ItemType int

const (
	// File for file item
	File ItemType = iota
	// Dir for directory item
	Dir
)

// Item is directory tree struct
type Item struct {
	Type ItemType `json:"type"`
	Name string   `json:"name"`
	Tree []Item   `json:"tree"`
}

type Notification struct {
	Hostname  string   `json:"hostname"`
	Event     string   `json:"event"`
	Type      ItemType `json:"type"`
	Path      string   `json:"path"`
	timestamp int64    `timestamp:"timestamp"`
}

type Event int

const (
	Create Event = iota
	Write
	Rename
	Delete
)
