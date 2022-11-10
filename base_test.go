package model

import "testing"

func TestJoinQueryFields(t *testing.T) {
	type args struct {
		columns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"a",
			args{
				[]string{"a", "b"},
			},
			" and a = ? and b = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinQueryFields(tt.args.columns); got != tt.want {
				t.Errorf("JoinQueryFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
