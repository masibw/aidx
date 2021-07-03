package cmd

import "testing"

func TestGenIndex(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		wantIdx string
	}{
		{
			name:    "Return Index for simple query.",
			query:   "SELECT * FROM users WHERE name = 'hoge';",
			wantIdx: "CREATE INDEX idx ON users(name);",
		},
		{
			name:    "Return Index for simple query with AND expression.",
			query:   "SELECT * FROM users WHERE name = 'hoge' AND age = 20;",
			wantIdx: "CREATE INDEX idx ON users(name, age);",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIdx := GenIndex(tt.query); gotIdx != tt.wantIdx {
				t.Errorf("GenIndex() = %v, want %v", gotIdx, tt.wantIdx)
			}
		})
	}
}
