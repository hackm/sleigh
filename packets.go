package main

// Hey is first message packet through UDP multicast
type Hey struct {
	Hostname string `json:"hostname"`
	Tree     []Item `json:"type"`
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
