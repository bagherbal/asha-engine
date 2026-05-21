package generation2globalphasez2equivarianceorientationgaugeaudit

import (
	"strings"
	"testing"
)

func TestGate907Z2Operation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	z := a.Z2
	if !z.SwapsLambdaBarLambda || !z.SwapsProjectors || !z.FlipsQPhi || z.NativeTheorem {
		t.Fatalf("bad Z2 audit: %s", FormatZ2(z))
	}
}

func TestGate907AirlockAndEdgeEquivariance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Airlock.AlphaRanksInvariant || !a.Airlock.OrientationClassCandidate || a.Airlock.NativeEquivariance || !sameInts(a.Airlock.PlusFlagRanks, a.Airlock.MinusFlagRanks) {
		t.Fatalf("bad airlock audit: %s", FormatAirlock(a.Airlock))
	}
	if !a.Edge.MirrorRepresentativeExists || !a.Edge.RankKernelInvariant || a.Edge.NativeOperatorTheorem || !sameInts(a.Edge.CurrentRankPattern, a.Edge.MirrorRankPattern) || a.Edge.CurrentKernelCount != a.Edge.MirrorKernelCount {
		t.Fatalf("bad edge audit: %s", FormatEdge(a.Edge))
	}
}

func TestGate907TraceInvariance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	tr := a.Trace
	if !tr.TraceInvariant || !tr.SquareTraceInvariant || !tr.NEffInvariant || !tr.CYukawaInvariant || tr.NativeTraceTheorem {
		t.Fatalf("bad trace audit: %s", FormatTrace(tr))
	}
}

func TestGate907LabelsAndWound(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Labels.SocketLabelsChange || a.Labels.PhysicalLabelsCertified || a.Labels.TraceLedgerNeedsPhysical {
		t.Fatalf("bad labels: %s", FormatLabels(a.Labels))
	}
	if a.Wound.AbsoluteSignNeeded || a.Wound.NativeSolved || !strings.Contains(a.Wound.ReducedWound, "Z2-equivariant") {
		t.Fatalf("bad wound: %s", FormatWound(a.Wound))
	}
}

func TestGate907FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || !near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate907Theorem(t *testing.T) {
	res := Generation2GlobalPhaseZ2EquivarianceOrientationGaugeAuditTheorem().Verify()
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
