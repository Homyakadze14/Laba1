package exptree

import "testing"

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name      string
		root      *Node
		want      int
		wantErr   bool
		errString string
	}{
		{
			name:      "nil tree",
			root:      nil,
			want:      0,
			wantErr:   true,
			errString: ErrEmptyTree.Error(),
		},
		{
			name:    "leaf constant",
			root:    &Node{Value: 42},
			want:    42,
			wantErr: false,
		},
		{
			name: "addition",
			root: &Node{
				Operation: "+",
				Left:      &Node{Value: 5},
				Right:     &Node{Value: 3},
			},
			want: 8,
		},
		{
			name: "subtraction",
			root: &Node{
				Operation: "-",
				Left:      &Node{Value: 10},
				Right:     &Node{Value: 4},
			},
			want: 6,
		},
		{
			name: "multiplication",
			root: &Node{
				Operation: "*",
				Left:      &Node{Value: 7},
				Right:     &Node{Value: 6},
			},
			want: 42,
		},
		{
			name: "division",
			root: &Node{
				Operation: "/",
				Left:      &Node{Value: 20},
				Right:     &Node{Value: 5},
			},
			want: 4,
		},
		{
			name: "division by zero",
			root: &Node{
				Operation: "/",
				Left:      &Node{Value: 10},
				Right:     &Node{Value: 0},
			},
			want:      0,
			wantErr:   true,
			errString: "division by zero",
		},
		{
			name: "power",
			root: &Node{
				Operation: "^",
				Left:      &Node{Value: 2},
				Right:     &Node{Value: 3},
			},
			want: 8,
		},
		{
			name: "negative exponent",
			root: &Node{
				Operation: "^",
				Left:      &Node{Value: 2},
				Right:     &Node{Value: -1},
			},
			want:      0,
			wantErr:   true,
			errString: "negative exponent not supported for integers",
		},
		{
			name: "complex expression",
			root: &Node{
				Operation: "*",
				Left: &Node{
					Operation: "+",
					Left:      &Node{Value: 1},
					Right:     &Node{Value: 2},
				},
				Right: &Node{
					Operation: "-",
					Left:      &Node{Value: 5},
					Right:     &Node{Value: 3},
				},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.root)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Evaluate() expected error, got nil")
				} else if tt.errString != "" && err.Error() != tt.errString {
					t.Errorf("Evaluate() error = %v, want %v", err, tt.errString)
				}
				return
			}
			if err != nil {
				t.Errorf("Evaluate() unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Evaluate() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestBuild(t *testing.T) {
	tests := []struct {
		name      string
		tokens    []string
		wantErr   bool
		errType   error
		evalValue int
	}{
		{
			name:      "simple addition",
			tokens:    []string{"3", "4", "+"},
			wantErr:   false,
			evalValue: 7,
		},
		{
			name:      "multiplication and addition",
			tokens:    []string{"2", "3", "4", "+", "*"},
			wantErr:   false,
			evalValue: 14,
		},
		{
			name:      "division and power",
			tokens:    []string{"8", "2", "/", "3", "^"},
			wantErr:   false,
			evalValue: 64,
		},
		{
			name:      "even number of tokens",
			tokens:    []string{"1", "2", "+", "3"},
			wantErr:   true,
			errType:   ErrTokCount,
			evalValue: 0,
		},
		{
			name:      "insufficient operands",
			tokens:    []string{"1", "+"},
			wantErr:   true,
			errType:   ErrTokCount,
			evalValue: 0,
		},
		{
			name:      "invalid token",
			tokens:    []string{"1", "2", "$"},
			wantErr:   true,
			errType:   ErrBadTokens,
			evalValue: 0,
		},
		{
			name:      "extra tokens left in stack",
			tokens:    []string{"1", "2", "3", "+"},
			wantErr:   true,
			errType:   ErrTokCount,
			evalValue: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree, err := Build(tt.tokens)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Build() expected error, got nil")
				} else if tt.errType != nil && err != tt.errType {
					t.Errorf("Build() error = %v, want %v", err, tt.errType)
				}
				return
			}
			if err != nil {
				t.Errorf("Build() unexpected error: %v", err)
				return
			}
			val, err := Evaluate(tree.Parent)
			if err != nil {
				t.Errorf("Evaluate() error: %v", err)
			}
			if val != tt.evalValue {
				t.Errorf("Evaluate() = %d, want %d", val, tt.evalValue)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected int
	}{
		{
			name:     "nil tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "One node",
			root:     &Node{Value: 42},
			expected: 1,
		},
		{
			name: "One operation",
			root: &Node{
				Operation: "+",
				Left:      &Node{Value: 1},
				Right:     &Node{Value: 2},
			},
			expected: 2,
		},
		{
			name: "left-skewed tree",
			root: &Node{
				Operation: "*",
				Left: &Node{
					Operation: "+",
					Left:      &Node{Value: 1},
					Right:     &Node{Value: 2},
				},
				Right: &Node{Value: 3},
			},
			expected: 3,
		},
		{
			name: "balanced tree",
			root: &Node{
				Operation: "^",
				Left: &Node{
					Operation: "+",
					Left:      &Node{Value: 1},
					Right:     &Node{Value: 2},
				},
				Right: &Node{
					Operation: "*",
					Left:      &Node{Value: 3},
					Right:     &Node{Value: 4},
				},
			},
			expected: 3,
		},
		{
			name: "chain",
			root: &Node{
				Operation: "neg",
				Right: &Node{
					Operation: "neg",
					Right: &Node{
						Operation: "neg",
						Right:     &Node{Value: 10},
					},
				},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Height(tt.root)
			if got != tt.expected {
				t.Errorf("Height() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestPostfixString(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want string
	}{
		{
			name: "nil",
			root: nil,
			want: "",
		},
		{
			name: "leaf",
			root: &Node{Value: 123},
			want: "123",
		},
		{
			name: "addition",
			root: &Node{
				Operation: "+",
				Left:      &Node{Value: 3},
				Right:     &Node{Value: 4},
			},
			want: "3 4 +",
		},
		{
			name: "multiplication and addition",
			root: &Node{
				Operation: "*",
				Left: &Node{
					Operation: "+",
					Left:      &Node{Value: 2},
					Right:     &Node{Value: 3},
				},
				Right: &Node{Value: 5},
			},
			want: "2 3 + 5 *",
		},
		{
			name: "power and division",
			root: &Node{
				Operation: "^",
				Left: &Node{
					Operation: "/",
					Left:      &Node{Value: 8},
					Right:     &Node{Value: 2},
				},
				Right: &Node{Value: 3},
			},
			want: "8 2 / 3 ^",
		},
		{
			name: "nested left",
			root: &Node{
				Operation: "-",
				Left: &Node{
					Operation: "-",
					Left:      &Node{Value: 10},
					Right:     &Node{Value: 4},
				},
				Right: &Node{Value: 2},
			},
			want: "10 4 - 2 -",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostfixString(tt.root)
			if got != tt.want {
				t.Errorf("PostfixString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPrefixString(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want string
	}{
		{
			name: "nil",
			root: nil,
			want: "",
		},
		{
			name: "leaf",
			root: &Node{Value: 7},
			want: "7",
		},
		{
			name: "addition",
			root: &Node{
				Operation: "+",
				Left:      &Node{Value: 1},
				Right:     &Node{Value: 2},
			},
			want: "+ 1 2",
		},
		{
			name: "multiplication and addition",
			root: &Node{
				Operation: "*",
				Left: &Node{
					Operation: "+",
					Left:      &Node{Value: 3},
					Right:     &Node{Value: 4},
				},
				Right: &Node{Value: 5},
			},
			want: "* + 3 4 5",
		},
		{
			name: "power and division",
			root: &Node{
				Operation: "^",
				Left: &Node{
					Operation: "/",
					Left:      &Node{Value: 16},
					Right:     &Node{Value: 2},
				},
				Right: &Node{Value: 2},
			},
			want: "^ / 16 2 2",
		},
		{
			name: "nested left",
			root: &Node{
				Operation: "-",
				Left: &Node{
					Operation: "-",
					Left:      &Node{Value: 10},
					Right:     &Node{Value: 3},
				},
				Right: &Node{Value: 1},
			},
			want: "- - 10 3 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrefixString(tt.root)
			if got != tt.want {
				t.Errorf("PrefixString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCountOperations(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want int
	}{
		{
			name: "nil",
			root: nil,
			want: 0,
		},
		{
			name: "leaf",
			root: &Node{Value: 42},
			want: 0,
		},
		{
			name: "single operator",
			root: &Node{
				Operation: "+",
				Left:      &Node{Value: 1},
				Right:     &Node{Value: 2},
			},
			want: 1,
		},
		{
			name: "two operators",
			root: &Node{
				Operation: "*",
				Left: &Node{
					Operation: "+",
					Left:      &Node{Value: 3},
					Right:     &Node{Value: 4},
				},
				Right: &Node{Value: 5},
			},
			want: 2,
		},
		{
			name: "three operators",
			root: &Node{
				Operation: "-",
				Left: &Node{
					Operation: "/",
					Left: &Node{
						Operation: "*",
						Left:      &Node{Value: 2},
						Right:     &Node{Value: 3},
					},
					Right: &Node{Value: 4},
				},
				Right: &Node{Value: 1},
			},
			want: 3,
		},
		{
			name: "full binary tree of height 3",
			root: &Node{
				Operation: "+",
				Left: &Node{
					Operation: "*",
					Left:      &Node{Value: 1},
					Right:     &Node{Value: 2},
				},
				Right: &Node{
					Operation: "-",
					Left:      &Node{Value: 3},
					Right:     &Node{Value: 4},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOperations(tt.root)
			if got != tt.want {
				t.Errorf("CountOperations() = %d, want %d", got, tt.want)
			}
		})
	}
}
