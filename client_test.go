package main

import "testing"

func TestNotify(t *testing.T) {
	type args struct {
		ip   string
		port int
		n    *Notification
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
			if err := Notify(tt.args.ip, tt.args.port, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
