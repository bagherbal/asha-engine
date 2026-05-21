package generation2rightcharacterorientationandsocketordersourceaudit

import (
	"strings"
	"testing"
)

func TestGate898RightCharacterPairNeedsPhaseOrientation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.RightCharacter
	if !r.PairTyped || !r.ComplexOrientationCandidate || r.NativePhaseOrientation || r.SelectsPlusWithoutOrientation || !r.RequiresPhaseOrientationSeal {
		t.Fatalf("bad right character audit: %s", FormatRightCharacter(r))
	}
	if !containsAll(r.Characters, []string{RightCharacterPlus, RightCharacterMinus}) || !containsAll(r.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureCharacterConjugationNeedsPhase}) {
		t.Fatalf("missing right character statuses: %s", FormatRightCharacter(r))
	}
}

func TestGate898ComplexOrientationCanStateButNotSourceOrder(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.RightCharacter
	if !r.EPlusAsLambda || !r.EMinusAsBarLambda || !r.SocketOrderStatedGivenOrientation || r.SelectsPlusWithoutOrientation {
		t.Fatalf("complex orientation leaked: %s", FormatRightCharacter(r))
	}
	if !containsAll(r.Supports, []string{SupportSocketOrderGivenComplexOrientation, SupportPhaseOrientationCanSelectPlus, SupportAirlockOrderFollowsPhaseSeal}) || !containsAll(r.Failures, []string{FailureComplexOrientationNotNativeSelector, FailureLambdaBarLabelingConvention}) {
		t.Fatalf("missing complex orientation statuses: %s", FormatRightCharacter(r))
	}
}

func TestGate898OneFormBoundaryJAndBMinusLDoNotBreakZ2(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.OneFormArrow.MatchesEPlusPunctureOrder || a.OneFormArrow.DerivesOrderIndependently || !a.OneFormArrow.RestatesSocketOrder || a.OneFormArrow.NativeArrowDirectionSelector {
		t.Fatalf("one-form route leaked: %s", FormatOneFormArrow(a.OneFormArrow))
	}
	if !a.BoundaryDegree.IndexesFlagLevels || a.BoundaryDegree.SelectsSocketSign || a.BoundaryDegree.BreaksZ2 {
		t.Fatalf("boundary route leaked: %s", FormatBoundaryDegree(a.BoundaryDegree))
	}
	if a.JChirality.KOSignCertified || a.JChirality.JMirrorBreaksZ2 || a.JChirality.ChiralitySelectsPlus {
		t.Fatalf("J/chirality route leaked: %s", FormatJChirality(a.JChirality))
	}
	if !a.BMinusL.CompensationWorks || a.BMinusL.BreaksZ2 || a.BMinusL.PlusPunctureCharge != BMinusLLepton || a.BMinusL.MinusPunctureCharge != BMinusLLepton {
		t.Fatalf("B-L route leaked: %s", FormatBMinusL(a.BMinusL))
	}
}

func TestGate898PhaseOrientationIsNextFrontier(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Phase
	if !p.Z2IsPhaseAmbiguity || !p.RightPhaseSealCandidate || p.NativeTheorem || !p.SelectsPlusIfSealed || p.NextFrontier != NextFrontier {
		t.Fatalf("bad phase audit: %s", FormatPhase(p))
	}
	if !containsAll(p.Supports, []string{SupportPhaseOrientationSealRequired, SupportZ2IsPhaseAmbiguity, SupportWoundSharpensToPhaseSource}) || !containsAll(p.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionSigmaPlus}) {
		t.Fatalf("missing phase statuses: %s", FormatPhase(p))
	}
}

func TestGate898FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate898Theorem(t *testing.T) {
	res := Generation2RightCharacterOrientationAndSocketOrderSourceAuditTheorem().Verify()
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
