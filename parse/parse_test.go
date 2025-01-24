package parse

import "testing"

func TestIDFromURL(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			"fail because not http or https",
			args{target: "ftp:localhost"},
			"",
			"",
			true,
		},
		{
			"success aoj",
			args{target: "https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ITP1_1_A"},
			"ITP1_1_A",
			"judge.u-aizu.ac.jp",
			false,
		},
		{
			"fail aoj without id query",
			args{target: "https://judge.u-aizu.ac.jp/onlinejudge/description.jsp"},
			"",
			"",
			true,
		},
		{
			"success atcoder",
			args{target: "https://atcoder.jp/contests/APG4b/tasks/APG4b_g"},
			"APG4b_g",
			"atcoder.jp",
			false,
		},
		{
			"fail atcoder without task",
			args{target: "https://atcoder.jp/contests/APG4b"},
			"",
			"",
			true,
		},
		{
			"fail atcoder invalid",
			args{target: "https://atcoder.jp/contests/APG4b/invalid/APG4b_g"},
			"",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, got1, err := IDFromURL(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("IDFromURL() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("IDFromURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IDFromURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
