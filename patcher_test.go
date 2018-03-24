package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TO struct {
	Path   string
	Binary []byte
}

func TestJSON(t *testing.T) {
	// fmt.Println("Hello")
	// file, err := os.Open(`C:\Users\iwate\Downloads\dotnet-sdk-2.0.2-win-gs-x64.exe`)
	// if err != nil {
	// 	t.Fatal("Cannot open file")
	// }
	// defer file.Close()
	// stat, err := os.Stat(`C:\Users\iwate\Downloads\dotnet-sdk-2.0.2-win-gs-x64.exe`)
	// if err != nil {
	// 	t.Fatal("Cannot read stat")
	// }
	// b, err := encodeChecksumIndex(file, stat.Size(), BlockSize)
	// if err != nil {
	// 	t.Fatal("Cannot encode")
	// }
	// b64 := base64.RawURLEncoding.EncodeToString(b)

	bs, _ := json.Marshal(TO{
		Path:   "hoge/fuga/piyo",
		Binary: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	})

	fmt.Printf("%s", string(bs))
}

func TestNewPatcher(t *testing.T) {
	type args struct {
		n                chan Notification
		remoteURLResolve remoteURLResolver
		localPathResolve LocalPathResolver
	}
	tests := []struct {
		name string
		args args
		want *Patcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPatcher(tt.args.n, tt.args.remoteURLResolve, tt.args.localPathResolve); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPatcher_equals(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name    string
		p       Patcher
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.equals(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Patcher.equals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Patcher.equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPatcher_patch(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name    string
		p       Patcher
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.patch(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Patcher.patch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPatcher_download(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name    string
		p       Patcher
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.download(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Patcher.download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPatcher_remove(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name    string
		p       Patcher
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.remove(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Patcher.remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPatcher_rename(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name    string
		p       Patcher
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.rename(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Patcher.rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPatcher_try(t *testing.T) {
	type args struct {
		fn func() error
	}
	tests := []struct {
		name string
		p    *Patcher
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.try(tt.args.fn)
		})
	}
}

func TestPatcher_Patch(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name string
		p    *Patcher
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Patch(tt.args.n)
		})
	}
}

func TestPatcher_Download(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name string
		p    *Patcher
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Download(tt.args.n)
		})
	}
}

func TestPatcher_Remove(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name string
		p    *Patcher
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Remove(tt.args.n)
		})
	}
}

func TestPatcher_Rename(t *testing.T) {
	type args struct {
		n Notification
	}
	tests := []struct {
		name string
		p    *Patcher
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Rename(tt.args.n)
		})
	}
}

func TestPatcher_Start(t *testing.T) {
	tests := []struct {
		name string
		p    *Patcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Start()
		})
	}
}
