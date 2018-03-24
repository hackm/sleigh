package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	gosync "github.com/Redundancy/go-sync"
)

func Test_uid(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uid(); got != tt.want {
				t.Errorf("uid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copy(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_intToBytes(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToBytes(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int64ToBytes(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := int64ToBytes(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("int64ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeRSync(t *testing.T) {
	type args struct {
		local  gosync.ReadSeekerAt
		remote string
		fs     gosync.FileSummary
	}
	tests := []struct {
		name       string
		args       args
		want       *gosync.RSync
		wantOutput string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			if got := makeRSync(tt.args.local, tt.args.remote, output, tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRSync() = %v, want %v", got, tt.want)
			}
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("makeRSync() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_createSummary(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    gosync.FileSummary
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createSummary(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulticast(t *testing.T) {
	type args struct {
		port int
		b    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Multicast(tt.args.port, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multicast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Multicast() = %v, want %v", got, tt.want)
			}
		})
	}
}
