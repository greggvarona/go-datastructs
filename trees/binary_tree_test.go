package trees

import (
	"reflect"
	"testing"

	"github.com/greggvarona/go-datastructs/model"
)

var nilTreeNode *BinaryTreeNode

func TestNewBinaryTreeNode(t *testing.T) {
	type args struct {
		data interface{}
		l    *BinaryTreeNode
		r    *BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
		want *BinaryTreeNode
	}{
		{
			name: "Creates new BinaryTreeNode without children.",
			args: args{
				data: 1,
				l:    nil,
				r:    nil,
			},
			want: &BinaryTreeNode{
				TreeNode{
					data:         1,
					neighborhood: []model.Node{nilTreeNode, nilTreeNode},
				},
			},
		},
		{
			name: "Creates new BinaryTreeNode with children.",
			args: args{
				data: 1,
				l:    NewBinaryTreeNode(2, nil, nil),
				r:    NewBinaryTreeNode(3, nil, nil),
			},
			want: &BinaryTreeNode{
				TreeNode{
					data:         1,
					neighborhood: []model.Node{NewBinaryTreeNode(2, nil, nil), NewBinaryTreeNode(3, nil, nil)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryTreeNode(tt.args.data, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
