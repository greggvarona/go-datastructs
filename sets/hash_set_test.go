package sets

import (
	"testing"

	"github.com/greggvarona/go-datastructs/model"
)

type args struct {
	d Hashable
}

type testHashable struct {
	testData string
}

func newTestHashable(code string) Hashable {
	return &testHashable{
		testData: code,
	}
}

func (t *testHashable) HashCode() string {
	return t.testData
}

func TestHashSet_Put(t *testing.T) {
	tests := []struct {
		name string
		hs   *HashSet
		args args
		want int
	}{
		{
			name: "should put hashable data and return 1 for the size",
			hs: &HashSet{
				data: map[string]model.Void{},
			},
			args: args{
				d: newTestHashable("1"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hs.Put(tt.args.d)

			if got := tt.hs.Size(); got != tt.want {
				t.Errorf("HashSet.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Contains(t *testing.T) {
	tests := []struct {
		name string
		hs   *HashSet
		args args
		want bool
	}{
		{
			name: "should not contain a key value of 2",
			hs: &HashSet{
				data: map[string]model.Void{
					"0": model.Void{},
				},
			},
			args: args{
				d: newTestHashable("2"),
			},
			want: false,
		},
		{
			name: "should contain a key value of 2",
			hs: &HashSet{
				data: map[string]model.Void{
					"2": model.Void{},
				},
			},
			args: args{
				d: newTestHashable("2"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hs.Contains(tt.args.d); got != tt.want {
				t.Errorf("HashSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Delete(t *testing.T) {
	tests := []struct {
		name string
		hs   *HashSet
		args args
		want bool
	}{
		{
			name: "should delete element with a key value of 2",
			hs: &HashSet{
				data: map[string]model.Void{
					"2": model.Void{},
				},
			},
			args: args{
				d: newTestHashable("2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hs.Delete(tt.args.d)

			if got := tt.hs.Contains(tt.args.d); got != tt.want {
				t.Errorf("HashSet.Contains() = %v", got)
			}
		})
	}
}
