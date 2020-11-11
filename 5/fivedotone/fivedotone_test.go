package fivedotone

import "testing"

func TestInsertion(t *testing.T) {
	type args struct {
		n int
		m int
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1024-19-2-6", args{1024, 19, 2, 6}, 1100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insertion(tt.args.n, tt.args.m, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}
