package bibliotheekbe

import "testing"

func Test_getAudience(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test unknown",
			args: args{in: "noidea"},
			want: "noidea",
		},
		{
			name: "test known",
			args: args{in: "age12-14"},
			want: "ages 12-14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAudience(tt.args.in); got != tt.want {
				t.Errorf("getAudience() = %v, want %v", got, tt.want)
			}
		})
	}
}
