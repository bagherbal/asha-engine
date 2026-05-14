package yukawaintertwiner

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ChargeCompatibleYukawaChannelsDerived {
		t.Fatalf("expected charge-compatible channels to be derived")
	}
	if a.MinimalChannelCount != 8 {
		t.Fatalf("expected 8 minimal channels, got %d", a.MinimalChannelCount)
	}
	if a.FiberEntryCount != 16 {
		t.Fatalf("expected 16 scalar-fiber entries, got %d", a.FiberEntryCount)
	}
	if a.HyperchargeResidualMax > 1e-10 {
		t.Fatalf("hypercharge residual too large: %g", a.HyperchargeResidualMax)
	}
	if a.UpChannels != 3 || a.DownChannels != 3 || a.NeutrinoChannels != 1 || a.ElectronChannels != 1 {
		t.Fatalf("unexpected channel counts: up=%d down=%d nu=%d e=%d", a.UpChannels, a.DownChannels, a.NeutrinoChannels, a.ElectronChannels)
	}
}
