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

// Notification is packet for notify diff
type Notification struct {
	Hostname  string   `json:"hostname"`
	Event     string   `json:"event"`
	Type      ItemType `json:"type"`
	Path      string   `json:"path"`
	timestamp int64    `timestamp:"timestamp"`
}

// Event for file change
type Event int

const (
	// Create file|dir
	Create Event = iota
	// Write file
	Write
	// Rename file|dir
	Rename
	// Delete file|dir
	Delete
)
