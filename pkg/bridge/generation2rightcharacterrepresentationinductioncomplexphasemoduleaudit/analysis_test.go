package generation2rightcharacterrepresentationinductioncomplexphasemoduleaudit

import (
	"strings"
	"testing"
)

func TestGate905MinimalConjugationClosedModule(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.MinimalModule
	if !m.TwoCharacterMinimal || !m.MatchesRhoRShape || m.NativeASHAIdentification || m.SelectsOrder {
		t.Fatalf("bad minimal module audit: %s", FormatMinimalModule(m))
	}
}

func TestGate905ProjectorSupportsReconstructRhoRButNeedOrientation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.ProjectorSupport
	if !p.ProjectorsRealized || !p.RhoRFormReconstructed || p.IdentifyELambdaAsEPlus || !p.NeedsPhaseOrientationChoice {
		t.Fatalf("bad projector support audit: %s", FormatProjectorSupport(p))
	}
}

func TestGate905HopfAndCL17RemainAbstractSources(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.HopfAction.S1ActsOnAbstractModule || !a.HopfAction.ReconstructsSplitIfIdentified || a.HopfAction.NativeRightSocketAction {
		t.Fatalf("bad Hopf action audit: %s", FormatHopfAction(a.HopfAction))
	}
	if !a.CL17Induction.SuppliesComplexStructure || !a.CL17Induction.IMinusISplitMatchesPair || a.CL17Induction.EigenSocketToCR2Map || a.CL17Induction.InducesRhoRAction {
		t.Fatalf("bad Cl17 induction audit: %s", FormatCL17Induction(a.CL17Induction))
	}
}

func TestGate905PairNotOrder(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	o := a.Order
	if !o.PairCertified || o.OrderCertified || o.PositivePhaseExposure || !strings.Contains(o.RemainingWound, "orientation") {
		t.Fatalf("order incorrectly certified: %s", FormatOrder(o))
	}
}

func TestGate905MissingObjectAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MissingObject.PairNoLongerArbitrary || a.MissingObject.NativeSolved {
		t.Fatalf("bad missing object: %s", FormatMissingObject(a.MissingObject))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || !near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate905Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate905Theorem(t *testing.T) {
	res := Generation2RightCharacterRepresentationInductionComplexPhaseModuleAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
