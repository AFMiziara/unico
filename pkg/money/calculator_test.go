package money

import (
	"reflect"
	"testing"
)

func Test_calculator_round(t *testing.T) {
	type args struct {
		a *Amount
		e int
	}
	tests := []struct {
		name string
		c    *calculator
		args args
		want *Amount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.round(tt.args.a, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculator.round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator_negative(t *testing.T) {
	type args struct {
		a *Amount
	}
	tests := []struct {
		name string
		c    *calculator
		args args
		want *Amount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.negative(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculator.negative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator_allocate(t *testing.T) {
	type args struct {
		a *Amount
		r int
		s int
	}
	tests := []struct {
		name string
		c    *calculator
		args args
		want *Amount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.allocate(tt.args.a, tt.args.r, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculator.allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}
