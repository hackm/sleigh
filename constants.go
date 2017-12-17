package main

const (
	// MB is constant value for Mega Byte
	MB uint = 1024 * 1024
	// BlockSize for ChecksumIndex
	BlockSize uint = 1 * MB

	MulticastAddr   = "224.0.0.1"
	MaxDatagramSize = 8192
	RetryMax        = 10
)
