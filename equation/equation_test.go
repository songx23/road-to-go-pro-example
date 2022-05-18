package equation

import "testing"

func TestSolveLinear(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "equation 1",
			args: args{
				a: 2.0,
				b: -2.0,
			},
			want: 1.0,
		},
		{
			name: "equation 2",
			args: args{
				a: 2.0,
				b: -1.0,
			},
			want: 0.5,
		},
		{
			name: "equation 3",
			args: args{
				a: 3.0,
				b: 1.0,
			},
			want: -1.0 / 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveLinear(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SolveLinear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveQuadratic(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "equation 1",
			args: args{
				a: 1,
				b: -1,
				c: -6,
			},
			want:  3,
			want1: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SolveQuadratic(tt.args.a, tt.args.b, tt.args.c)
			if got != tt.want {
				t.Errorf("SolveQuadratic() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SolveQuadratic() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
