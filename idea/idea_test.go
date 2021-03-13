package idea

import "testing"

func TestThing(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Thing(tt.args.a); got != tt.want {
				t.Errorf("Thing() = %v, want %v", got, tt.want)
			}
		})
	}
}
