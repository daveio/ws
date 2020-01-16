package cmd

import (
	"github.com/spf13/cobra"
	"reflect"
	"testing"
)

func TestSwitchRun(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSwitchValidArgs(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwitchValidArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwitchValidArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
