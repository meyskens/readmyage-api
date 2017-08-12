package bibliotheekbe

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		term string
	}
	tests := []struct {
		name    string
		args    args
		want    SearchResponse
		wantErr bool
	}{
		{
			name: "test isbn",
			args: args{
				term: "9789026119453",
			},
			wantErr: false,
			want:    testInfoReuzenperzik,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
