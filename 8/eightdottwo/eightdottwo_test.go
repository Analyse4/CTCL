package eightdottwo

import (
	"reflect"
	"testing"

	"github.com/Analyse4/mew/adt/stack"
)

func TestRobotSearchPath(t *testing.T) {
	want1 := []*position{
		{7, 6},
		{7, 5},
		{7, 4},
		{7, 3},
		{7, 2},
		{7, 1},
		{6, 1},
		{6, 0},
		{5, 0},
		{4, 0},
		{3, 0},
		{2, 0},
		{1, 0},
		{0, 0},
	}
	want2 := []*position{
		{7, 6},
		{7, 5},
		{7, 4},
		{7, 3},
		{7, 2},
		{7, 1},
		{6, 1},
		{5, 1},
		{4, 1},
		{3, 1},
		{2, 1},
		{1, 1},
		{0, 1},
		{0, 0},
	}
	s1 := stack.New()
	s2 := stack.New()

	type args struct {
		r          int
		c          int
		banPoint   []*position
		path       *stack.Stack
		faildPoint []*position
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []*position
	}{
		{"1-test", args{7, 7, []*position{{7, 0}}, s1, make([]*position, 0)}, true, want1},
		{"2-test", args{7, 7, []*position{{1, 0}}, s2, make([]*position, 0)}, true, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, _ := RobotSearchPath(tt.args.r, tt.args.c, tt.args.banPoint, tt.args.path, tt.args.faildPoint)
			if got != tt.want {
				t.Errorf("RobotSearchPath() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1.Size(), len(tt.want1)) {
				t.Errorf("RobotSearchPath() got1 = %v, want %v", got1.Size(), len(tt.want1))
				t.Error("got: ")
				size := got1.Size()
				for i := 0; i < size; i++ {
					p := got1.Pop().(*position)
					t.Errorf("row: %v, cow: %v\n", p.x, p.y)
				}
				t.Error("want: ")
				for _, p := range want1 {
					t.Errorf("row: %v, cow: %v\n", p.x, p.y)
				}
			}
		})
	}
}
