package generation2hopfchiralityphaseanchortypingfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate903HopfTypingShapeButNoTransport(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	h := a.HopfTyping
	if !h.HasS1Phase || !h.HasPositiveConjugatePair || !h.LabelsEPlusIfSealed || !h.LabelsEMinusIfSealed || h.TypedActionOnCR2Certified || h.NativeRepresentationMap {
		t.Fatalf("bad Hopf typing: %s", FormatHopfTyping(h))
	}
}

func TestGate903CL17TypingShapeButNoRightAction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.CL17Typing
	if !c.OmegaSquaredMinusOne || !c.SuppliesIOverMinusI || !c.CorrectConjugationShape || c.TypedToRightCharacter || c.SelectsSocketOrder {
		t.Fatalf("bad CL17 typing: %s", FormatCL17Typing(c))
	}
}

func TestGate903AlignmentAndShapeFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	al := a.Alignment
	if !al.CompatiblePhaseTypes || !al.StrongestBridgeCandidate || al.NativeAlignmentMap || al.TransportToRhoR || al.CanSourceRightCharacterAnchor {
		t.Fatalf("alignment incorrectly promoted: %s", FormatAlignment(al))
	}
	sf := a.ShapeFirewall
	if !sf.SamePhaseShape || sf.TypedTransportCertified || !sf.ConjugationResonanceOnly || sf.NativeLambdaSelection {
		t.Fatalf("shape firewall leak: %s", FormatShapeFirewall(sf))
	}
}

func TestGate903TransportTarget(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	tt := a.TransportTarget
	if tt.MissingObject != PhaseTransportMap || !tt.Sharpened || tt.NativeMapCertified {
		t.Fatalf("bad transport target: %s", FormatTransportTarget(tt))
	}
	if !strings.Contains(tt.Codomain, "rho_R") || !strings.Contains(tt.RequiredAction, "lambda socket") {
		t.Fatalf("transport target not typed enough: %s", FormatTransportTarget(tt))
	}
}

func TestGate903FreezeAndFirewalls(t *testing.T) {
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

func TestGate903Theorem(t *testing.T) {
	res := Generation2HopfChiralityPhaseAnchorTypingFirewallAuditTheorem().Verify()
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
