package generation2nativez2equivariantairlockfunctorandboundaryalphasourceaudit

import (
	"strings"
	"testing"
)

func TestGate909Z2WellDefinedness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	w := a.WellDefined
	if !w.TauMapsPunctures || !w.TauMapsDegreeOneTarget || !w.TauMapsDegreeTwoTarget || !w.IBZ2CommutesWithTau || !w.RankPairInvariant || !w.ClassLevelWellDefined || w.NativeFunctorTheorem {
		t.Fatalf("bad Z2 well-definedness: %s", FormatWellDefined(w))
	}
	if w.Plus.RankF1OverF0 != 3 || w.Plus.RankF2OverF0 != 7 || w.Minus.RankF1OverF0 != 3 || w.Minus.RankF2OverF0 != 7 {
		t.Fatalf("bad rank pair: %s", FormatWellDefined(w))
	}
}

func TestGate909ReducedB2Response(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Response
	if !r.ZeroOrderSuppressed || !r.CubicTermStopped || !r.Lambda3B2Zero || !r.CompatibleWithZ2Class || r.NativeBoundaryFunctional || r.NativeDegreeToFlagFunctor {
		t.Fatalf("bad reduced B2 response: %s", FormatResponse(r))
	}
	if r.DegreeOneRank != RankF1OverF0 || r.DegreeTwoRank != RankF2OverF0 {
		t.Fatalf("bad target ranks: %s", FormatResponse(r))
	}
}

func TestGate909CrossLaneAndAlphaClassSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.CrossLane
	if c.DegreeOneToDegreeTwoAllowed || c.DegreeTwoToDegreeOneAllowed || !c.ExcludedIfFunctorCertified || c.NativeCrossLaneTheorem || c.NativeLinearDomainClassExclusion || c.NativeQuadraticFaceClassExclusion {
		t.Fatalf("bad cross-lane audit: %s", FormatCrossLane(c))
	}
	alpha := a.Alpha
	if alpha.RepresentativeAlphaRequired || !alpha.Z2ClassAlphaSupported || alpha.NativeAlphaCertified || !alpha.RepresentativeIndependent || !alpha.SealWeakenedToClassSeal {
		t.Fatalf("bad alpha layer flags: %s", FormatAlpha(alpha))
	}
	if !near(alpha.LinearCoefficient, 0.3) || !near(alpha.QuadraticCoefficient, 7.0/72.0) || !near(alpha.ReconstructedAlpha, AlphaB) || alpha.Residual != 0 {
		t.Fatalf("bad alpha reconstruction: %s", FormatAlpha(alpha))
	}
}

func TestGate909R3ConsequenceFreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.R3.R3LedgerOnZ2AirlockClass || a.R3.PhaseSignBlocksTraceLedger || a.R3.NativeZ2AlphaFunctor || a.R3.NativeR3 || a.R3.FullAFDescent || a.R3.OfficialLedgerUpdate {
		t.Fatalf("bad R3 consequence: %s", FormatR3(a.R3))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate909Theorem(t *testing.T) {
	res := Generation2NativeZ2EquivariantAirlockFunctorAndBoundaryAlphaSourceAuditTheorem().Verify()
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
	if !strings.Contains(joined, FinalTruth) || !strings.Contains(joined, Classification) || !strings.Contains(joined, ShortStatus) {
		t.Fatalf("missing final classification/truth in notes: %s", joined)
	}
}
