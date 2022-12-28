package myerr

import (
	"errors"
	"testing"
)

func TestMyErr_Is(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields *MyErr
		args   args
		want   bool
	}{
		{
			name:   "same myerr",
			fields: SUCCESS,
			args: args{
				SUCCESS,
			},
			want: true,
		},
		{
			name:   "diff myerr",
			fields: SUCCESS,
			args: args{
				MODIFY_CONFIG_ERROR,
			},
			want: false,
		},
		{
			name:   "myerr vs goerror",
			fields: SUCCESS,
			args: args{
				err: errors.New("go error"),
			},
			want: false,
		},
		{
			name:   "myerr vs nil",
			fields: SUCCESS,
			args: args{
				err: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.fields
			if got := e.Is(tt.args.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
