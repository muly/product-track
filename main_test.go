package main

import "testing"

func Test_shouldNotify(t *testing.T) {
	type args struct {
		i input
		p product
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "checking product availability",
			args: args{i: input{typeOfRequest: "AVAILABILITY"},
				p: product{availability: true}},
			want: true,
		},
		{
			name: "checking product unavailability",
			args: args{i: input{typeOfRequest: "AVAILABILITY"},
				p: product{availability: false}},
			want: false,
		},
		{
			name: "less than minimum threshold limit ",
			args: args{},
			want: true,
		},
		{
			name: "greater than minimum threshold limit",
			args: args{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldNotify(tt.args.i, tt.args.p); got != tt.want {
				t.Errorf("shouldNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}
