package instantontracebridge

import "testing"

func TestBuildDefaultInstantonTraceBridge(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Input.Gate174ConditionalBranchAvailable || a.Input.Gate174StrictAbsoluteUDerived {
		t.Fatalf("expected inherited conditional-but-not-strict branch: %+v", a.Input)
	}
	if !a.Input.RelativeGaugeRatioClosed || !a.Input.WeakAngleSeedClosed {
		t.Fatalf("expected closed relative gauge data: %+v", a.Input)
	}
}

func TestContinuumIndexBridgeNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.StrictContinuumIndexBridgeDerived {
		t.Fatalf("continuum index bridge should not be derived: %+v", a.Firewall)
	}
	if a.Firewall.ContinuumIndexRequirements != 5 || a.Firewall.ContinuumIndexRequirementsMet >= a.Firewall.ContinuumIndexRequirements {
		t.Fatalf("expected unmet index requirements: %+v", a.Firewall)
	}
}

func TestTraceKineticBridgeNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.StrictTraceKineticBridgeDerived || a.TraceAudit.AbsoluteTraceScaleDerived || a.TraceAudit.KineticIntegralNormalization {
		t.Fatalf("trace/kinetic bridge should not be derived: firewall=%+v trace=%+v", a.Firewall, a.TraceAudit)
	}
	if !a.TraceAudit.RepresentationTraceRatioClosed || !a.TraceAudit.F0ConventionDependenceOpen {
		t.Fatalf("expected relative trace closure with f0 convention dependence: %+v", a.TraceAudit)
	}
}

func TestShortcutRoutesQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Routes) != 5 {
		t.Fatalf("expected five candidate routes, got %d", len(a.Routes))
	}
	if !noStrictRoute(a.Routes) || !hasConditionalRoute(a.Routes) || !hasObservedForbiddenRoute(a.Routes) {
		t.Fatalf("bad route quarantine: %s", FormatRoutes(a.Routes))
	}
}

func TestStrictNullityRemainsThreeConditionalTwo(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.ConditionalAbsoluteUPreserved || a.Firewall.ConditionalNullityAfter != 2 {
		t.Fatalf("conditional branch should be preserved with nullity 2: %+v", a.Firewall)
	}
	if a.Firewall.StrictAbsoluteUDerived || a.Firewall.StrictNullityBefore != 3 || a.Firewall.StrictNullityAfter != 3 {
		t.Fatalf("strict nullity should remain 3: %+v", a.Firewall)
	}
	if a.Firewall.PhysicalCouplingsDerived || a.Firewall.FineStructureDerived || a.Firewall.HiddenObservedInputUsed {
		t.Fatalf("physical constants should remain sealed: %+v", a.Firewall)
	}
}
