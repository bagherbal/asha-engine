package generation2boundaryscalarprojectorselectorfactorizationfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate687Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SupportMinimalityInherited || a.Inherited.SelectedProjector != "P_K7" || !a.Inherited.PriorFirewallPreserved {
		t.Fatalf("bad Gate686 inheritance: %+v", a.Inherited)
	}
	if !a.ScalarAction.CentralAction || a.ScalarAction.CarriesProjectorDirection || a.ScalarAction.CanDistinguishPK7FromPW7 {
		t.Fatalf("bad scalar centrality audit: %+v", a.ScalarAction)
	}
	if !a.Factorization.FactorizationRequired || a.Factorization.NativeCouplingProved {
		t.Fatalf("bad factorization audit: %+v", a.Factorization)
	}
}

func TestScalarActionCannotImposeSupport(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ScalarAction.CommutesWithPB || !a.ScalarAction.CommutesWithPG || !a.ScalarAction.CommutesWithAnyProjector {
		t.Fatalf("scalar action must commute with projector algebra: %+v", a.ScalarAction)
	}
	if a.ScalarAction.CanImposeBooleanSupport || a.ScalarAction.CanImposeOctonionicSupport || a.NoGo.BoundaryScalarImposesSupport {
		t.Fatalf("scalar action cannot impose Boolean-octonionic support: scalar=%+v nogo=%+v", a.ScalarAction, a.NoGo)
	}
	if !a.NoGo.NoGoCertified || !strings.Contains(a.NoGo.Verdict, StatusSSplitAloneDoesNotImposeSupport) {
		t.Fatalf("expected no-go verdict: %+v", a.NoGo)
	}
}

func TestNativeSupportSelectorRemainsSeparate(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SupportSelection.ImageInBooleanSector || !a.SupportSelection.ImageInOctonionicSector || !a.SupportSelection.ImageInIntersection {
		t.Fatalf("support selector must force the intersection carrier: %+v", a.SupportSelection)
	}
	if a.SupportSelection.IntersectionDimension != 7 || !a.SupportSelection.RankEqualsIntersection || a.SupportSelection.SelectedProjector != "P_K7" {
		t.Fatalf("support selector should select P_K7: %+v", a.SupportSelection)
	}
	if !a.SupportSelection.IndependentOfSSplit {
		t.Fatalf("support selection should be independent of the scalar at this theorem level: %+v", a.SupportSelection)
	}
}

func TestThreeSealFactorization(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ThreeSeal.BoundaryScalarControlsAmplitude || !a.ThreeSeal.ProjectorSelectorControlsIdentity || !a.ThreeSeal.TraceControlsScalarResponse {
		t.Fatalf("bad three-seal decomposition: %+v", a.ThreeSeal)
	}
	if a.Factorization.SSplitAloneSelectsIdentity || !a.Factorization.ProjectorIdentitySupportSealed || a.Factorization.NativeCouplingProved {
		t.Fatalf("factorization firewall violated: %+v", a.Factorization)
	}
}

func TestCandidateScalarIndistinguishability(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Indistinguishability.Candidates) != 3 || !a.Indistinguishability.AllRankSevenScaled || a.Indistinguishability.ScalarSeparatesCandidates {
		t.Fatalf("scalar should scale but not separate rank-seven candidates: %+v", a.Indistinguishability)
	}
	if !a.Indistinguishability.SupportSeparatesCandidates || !a.Indistinguishability.PK7SelectedBySupport || !a.Indistinguishability.PW7RejectedBySupport {
		t.Fatalf("native support should separate candidates: %+v", a.Indistinguishability)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryScalarProjectorSelectorFactorizationFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
