package higgsprecisionroutesieve

import "testing"

func TestBuildGate337PrecisionRepairRouteSieve(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error: %v", err)
	}
	if a.Inputs.RequiredRePiGeV2.RatString() != "504067437/11560000" {
		t.Fatalf("unexpected RePi target: %s", a.Inputs.RequiredRePiGeV2.RatString())
	}
	if len(a.Routes.Routes) != 3 {
		t.Fatalf("expected 3 repair routes, got %d", len(a.Routes.Routes))
	}
	if a.Routes.Routes[0].AllowedByNativeCore {
		t.Fatalf("contact-shape deformation must remain forbidden")
	}
	if !a.Routes.Routes[2].AllowedByNativeCore {
		t.Fatalf("pole correction route should be allowed as precision route")
	}
	if !nearlyFloat(a.Kernel.RawKernelGeV2, -991.5670298916105, 1e-9) {
		t.Fatalf("unexpected raw kernel: %s", dec(a.Kernel.RawKernelGeV2, 30))
	}
	if !nearlyFloat(a.Counterterm.FiniteRemainderGeV2, 1035.1714794590845, 1e-9) {
		t.Fatalf("unexpected remainder: %s", dec(a.Counterterm.FiniteRemainderGeV2, 30))
	}
	if !a.Firewalls.NoColliderClaim || !a.Firewalls.NoShapeFit {
		t.Fatalf("firewalls not preserved: %+v", a.Firewalls)
	}
}
