package util

import (
	"reflect"
	"testing"
)

func TestDetectShell(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   error
		wantShell Shell
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, gotShell := DetectShell()
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("DetectShell() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotShell, tt.wantShell) {
				t.Errorf("DetectShell() gotShell = %v, want %v", gotShell, tt.wantShell)
			}
		})
	}
}

func TestGetHomedir(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPath := GetHomedir(); gotPath != tt.wantPath {
				t.Errorf("GetHomedir() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func TestShellDetectionError_Error(t *testing.T) {
	type fields struct {
		EnvVar string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ShellDetectionError{
				EnvVar: tt.fields.EnvVar,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
