package ashafinalclosingtheorem

import "testing"

func TestGate374ClosingAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.HighestInheritedGate != 373 || a.Inheritance.ChargedFiniteDiracModuli != 13 || a.Inheritance.ExternalMinimalLedger != 15 || a.Inheritance.HolographicReductionFound {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Ledger.BoundaryCount != 4 || a.Ledger.ProxyCount != 2 || a.Ledger.FreeModuliCount != 1 {
		t.Fatalf("bad ledger partition: %+v", a.Ledger)
	}
	if a.Moduli.ChargedFlavorModuli != 13 || a.Moduli.ExternalLedger != 15 || !a.Moduli.FlatDirections || a.Moduli.PureGeometrySelectsPoint {
		t.Fatalf("bad moduli seal: %+v", a.Moduli)
	}
	if len(a.Routes.Routes) < 8 {
		t.Fatalf("expected closed route ledger, got %d", len(a.Routes.Routes))
	}
	if !a.Firewall.NoObservedYukawaValues || !a.Firewall.NoObservedCKMValues || !a.Firewall.NoManualTauEtaHamiltonian || !a.Firewall.NoFinalVacuumClaim {
		t.Fatalf("firewall breach: %+v", a.Firewall)
	}
	if !a.Closing.KinematicsComplete || !a.Closing.DynamicsOfFlavorUnselected || !a.Closing.VacuumFree {
		t.Fatalf("bad closing theorem: %+v", a.Closing)
	}
}

func TestNativeBoundaryValues(t *testing.T) {
	v := NativeBoundaryValues()
	if v["sin2_thetaW_boundary"] != 0.375 {
		t.Fatalf("bad weak angle boundary: %v", v["sin2_thetaW_boundary"])
	}
	if v["lambdaH_over_gstar2"] <= 0.25 || v["lambdaH_over_gstar2"] >= 0.27 {
		t.Fatalf("bad lambda ratio: %v", v["lambdaH_over_gstar2"])
	}
	if v["v_over_MP_hierarchy"] <= 0 || v["v_over_MP_hierarchy"] >= 1e-16 {
		t.Fatalf("bad hierarchy: %v", v["v_over_MP_hierarchy"])
	}
}

func TestTheoremPasses(t *testing.T) {
	res := AshaFinalClosingTheoremThirteenModuliVacuumManifoldTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
