package wikisearchapi

import (
	"testing"
)

func TestGetSearchWiki(t *testing.T) {
	type args struct {
		langu string
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "WikiTest",
			args: args{
				langu: "de-DE",
				query: "Oben"},
			want:    "Oben",
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSearchWiki(tt.args.langu, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSearchWiki() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Query.Search[0].Title != tt.want {
				t.Errorf("GetSearchWiki() = %v, want %v", got.Query.Search[0].Title, tt.want)
			}
		})
	}
}
