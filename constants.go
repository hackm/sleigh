package main

const (
	// MB is constant value for Mega Byte
	MB uint = 1024 * 1024
	// BlockSize for ChecksumIndex
	BlockSize uint = 1 * MB
	// MulticastAddr is IPv4 Multiacst Address, refer rfc5771
	MulticastAddr = "224.0.0.1"
	// MaxDatagramSize is limitatioion on UDP Multicast (support for NSF or old systems)
	MaxDatagramSize = 8192
	// RetryMax is maximum retry count
	RetryMax = 10
)
