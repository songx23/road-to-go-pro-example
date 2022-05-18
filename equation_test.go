package road_to_go_pro_example

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
