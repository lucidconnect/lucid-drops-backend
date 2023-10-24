package ledger

import (
	"testing"
)

func Test_isNegativeInt(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isNegativeInt",
			args: args{
				value: -1,
			},
			want: true,
		},
		{
			name: "isNotNegativeInt",
			args: args{
				value: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNegativeInt(tt.args.value); got != tt.want {
				t.Errorf("isNegativeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toNegativeInt(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "isNegativeInt",
			args: args{
				value: -1,
			},
			want: -1,
		},
		{
			name: "isNotNegativeInt",
			args: args{
				value: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toNegativeInt(tt.args.value); got != tt.want {
				t.Errorf("toNegativeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGreaterThanBalance(t *testing.T) {
	type args struct {
		value   int64
		balance int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isNotNegativeInt",
			args: args{
				value: 10,
				balance: 100,
			},
			want: false,
		},
		{
			name: "isNotNegativeInt",
			args: args{
				value: 1000,
				balance: 100,
			},
			want: true,
		},
		{
			name: "isNotNegativeInt",
			args: args{
				value: 101,
				balance: 100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGreaterThanBalance(tt.args.value, tt.args.balance); got != tt.want {
				t.Errorf("isGreaterThanBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
