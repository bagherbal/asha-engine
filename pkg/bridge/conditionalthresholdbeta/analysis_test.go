package conditionalthresholdbeta

import (
	"testing"

	"github.com/bagherbal/asha-engine/pkg/bridge/yukawaamplitudeseal"
)

func TestExactSectorRows(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	rows := map[yukawaamplitudeseal.YukawaKind]SectorRow{}
	for _, r := range a.Ledger.SectorRows {
		rows[r.Sector] = r
	}
	cases := []struct {
		k            yukawaamplitudeseal.YukawaKind
		u1, su2, su3 Rational
	}{
		{yukawaamplitudeseal.YukawaUp, NewRational(17, 30), NewRational(1, 2), NewRational(2, 3)},
		{yukawaamplitudeseal.YukawaDown, NewRational(1, 6), NewRational(1, 2), NewRational(2, 3)},
		{yukawaamplitudeseal.YukawaElectron, NewRational(1, 2), NewRational(1, 6), NewRational(0, 1)},
		{yukawaamplitudeseal.YukawaNeutrino, NewRational(1, 10), NewRational(1, 6), NewRational(0, 1)},
	}
	for _, c := range cases {
		r, ok := rows[c.k]
		if !ok {
			t.Fatalf("missing sector row %s", c.k)
		}
		if !r.U1YContribution.Equal(c.u1) || !r.SU2LContribution.Equal(c.su2) || !r.SU3CContribution.Equal(c.su3) || !r.ExactRational || !r.SplitRequiresBrokenPhaseVEV {
			t.Fatalf("bad row %s: %s", c.k, FormatSectorRow(r))
		}
	}
}

func TestLedgerReconstructsFermionInventory(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	l := a.Ledger
	if l.ThresholdRows != 12 || !l.AllRowsExactRational || !l.AllRowsConditionalOnTexture || !l.AllRowsConditionalOnVEV || !l.AllRowsConditionalOnScheme {
		t.Fatalf("bad ledger conditionals: %s", FormatLedger(l))
	}
	if !l.FermionContributionU1Y.Equal(NewRational(4, 1)) || !l.FermionContributionSU2L.Equal(NewRational(4, 1)) || !l.FermionContributionSU3C.Equal(NewRational(4, 1)) {
		t.Fatalf("bad fermion sums: %s", FormatLedger(l))
	}
	if l.AnyObservedMassUsed || l.AnyNativeFiniteThresholdRowDerived || l.NativeFiniteBetaTheoremDerived {
		t.Fatalf("ledger leaked finite/observed claim: %s", FormatLedger(l))
	}
}

func TestSchemeSealAndMatchingAnswer(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.SchemeSeal
	if !s.ExplicitBoundaryConvention || !s.Quarantined || !s.RequiredForThresholdRows || s.DerivedFromFiniteAlgebra || s.SharpStepDerivedNatively || !s.SmoothRegulatorSearched || s.SmoothRegulatorDerived || s.SchemeSelectedNatively || s.MSbarDerived || s.MOMDerived || !s.TreeLevelContinuityEnforced || s.FiniteMatchingCorrectionsDerived || !s.DownstreamMustDeclareSeal {
		t.Fatalf("bad scheme seal: %s", FormatScheme(s))
	}
	p := a.Piecewise
	if !p.PiecewiseSymbolicTreeBuilt || p.EvaluatedNumerically || p.ThresholdOrderingKnown || p.BoundaryScaleDerived || p.BoundaryCouplingDerived || p.FiniteMatchingCorrectionsDerived || !p.EnforcesTreeLevelContinuity || !p.SchemeDependentCorrectionsSealed {
		t.Fatalf("bad piecewise matching: %s", FormatPiecewise(p))
	}
}

func TestLowEnergyAndAbsoluteFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Domain
	if !d.FermionThresholdSymbolsAvailable || d.GaugeBosonThresholdsAvailable || !d.WZThresholdsBlocked || d.DeepInfraredFlowDefined || d.RunToMZAllowed {
		t.Fatalf("bad low-energy domain: %s", FormatDomain(d))
	}
	f := a.Firewall
	if !f.ConditionalBetaRowsActivated || f.FiniteThresholdRowsDerived || f.NativeSmoothRegulatorDerived || f.SchemeConventionDerivedFromFinite || !f.PiecewiseRGTreeConstructed || f.PhysicalRGFlowEvaluated || f.NumericalThresholdsKnown || f.ThresholdOrderingKnown || f.LowEnergyWZDomainKnown || f.GaugeBosonThresholdsDerived || f.AbsoluteBoundaryScaleDerived || f.AbsoluteBoundaryCouplingDerived || f.GaugeCouplingsDerived || f.TopologicalEightPiSquaredImported || f.FiniteToContinuumScaleDerived || f.ObservedInputsImported {
		t.Fatalf("firewall leaked: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalSchemeNullityBefore != 1 || f.ConditionalSchemeNullityAfter != 0 || f.ConditionalBetaRowNullityBefore != 1 || f.ConditionalBetaRowNullityAfter != 0 || f.ConditionalRGEvaluationNullityBefore != f.ConditionalRGEvaluationNullityAfter {
		t.Fatalf("bad nullity ledger: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := ConditionalThresholdBetaRowActivationDecouplingSchemeFirewallTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
