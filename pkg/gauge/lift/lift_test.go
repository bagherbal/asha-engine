package lift

import "testing"

func TestBuildDefaultCompression(t *testing.T) {
	c, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if c.Contact.Dimension() != 7 {
		t.Fatalf("contact dimension = %d, want 7", c.Contact.Dimension())
	}
	if len(c.ExteriorGenerators) != 4 {
		t.Fatalf("lifted generators = %d, want 4", len(c.ExteriorGenerators))
	}
	if c.CompressedFrameRank == 0 || c.CompressedFrameRank > 4 {
		t.Fatalf("compressed rank = %d, want 1..4", c.CompressedFrameRank)
	}
}
