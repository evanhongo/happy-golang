package string

import "testing"

func TestCleanStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TesttCleanStrSuccess",
			args: args{
				str: "e  v   a    n",
			},
			want: "e v a n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanStr(tt.args.str); got != tt.want {
				t.Errorf("CleanStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
