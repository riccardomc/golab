package secondlargest

import (
	"testing"

	. "github.com/riccardomc/golab/cake/validbst"
)

func TestSecondLargest(t *testing.T) {
	input :=
		&BST{50,
			&BST{30,
				&BST{20, nil, nil},
				&BST{40, nil, nil},
			},
			&BST{80,
				&BST{70, nil, nil},
				&BST{90, nil, nil},
			},
		}

	t.Run("Rightmost has no left subtree", func(t *testing.T) {
		expected := 80
		actual := secondLargest(input)
		if actual.Key != expected {
			t.Errorf("%d != %d", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})

	input =
		&BST{50,
			&BST{30,
				&BST{20, nil, nil},
				&BST{40, nil, nil},
			},
			&BST{80,
				&BST{70, nil, nil},
				&BST{90,
					&BST{85,
						&BST{82, nil, nil},
						nil,
					},
					&BST{87,
						&BST{86, nil, nil},
						nil,
					},
				},
			},
		}

	t.Run("Rightmost has subtree", func(t *testing.T) {
		expected := 87
		actual := secondLargest(input)
		if actual.Key != expected {
			t.Errorf("%d != %d", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})
}
