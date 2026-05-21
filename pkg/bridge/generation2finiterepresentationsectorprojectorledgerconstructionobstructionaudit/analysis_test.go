package generation2finiterepresentationsectorprojectorledgerconstructionobstructionaudit

import (
	"strings"
	"testing"
)

func TestGate835CentralSupportProjectorRecipeObstructed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Central.CentralIdempotentsOrthogonal || !a.Central.CentralIdempotentsSumI || !a.Central.SupportProjectorRecipeDefined || a.Central.SupportProjectorsInstantiated || a.Central.SupportRanksCertified || !a.Central.BasisIndependentAtCoarseLevel || a.Central.CompleteFiniteSectorLedger || a.Central.TraceMagnitudeCertified {
		t.Fatalf("central support recipe over/under-certified: %s", FormatCentral(a.Central))
	}
	if !containsAll(a.Central.Failures, []string{FailureNoCompleteFiniteRepresentationLedger, FailureNoCompleteRhoFSupportRankLedger, FailureCentralSupportsOnlyCoarseNotFullLedger}) {
		t.Fatalf("missing central failures: %s", strings.Join(a.Central.Failures, ","))
	}
}

func TestGate835RefinementsBimoduleAndDiracEdgesRemainUncertified(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Chirality.GammaFAvailable || !a.Chirality.JFAvailable || a.Chirality.GammaRefinementCertified || a.Chirality.JRefinementCertified || a.Chirality.LeftRightSplitCertified || a.Chirality.ParticleOppositeSplitCertified || a.Chirality.CompleteRefinementLedger {
		t.Fatalf("chirality/J refinement over-promoted: %s", FormatChirality(a.Chirality))
	}
	if !a.Bimodule.LeftActionRequired || !a.Bimodule.RightActionRequired || a.Bimodule.CommutantStableProjectorsCertified || a.Bimodule.FirstOrderStableCertified || a.Bimodule.BimoduleDecompositionCertified || a.Bimodule.CompleteTypedLedger {
		t.Fatalf("bimodule/commutant typing over-promoted: %s", FormatBimodule(a.Bimodule))
	}
	if !a.DiracEdges.RequiresProjectorLedger || !a.DiracEdges.UsesDF || !a.DiracEdges.EdgeSupportAudited || a.DiracEdges.EdgeSupportLedgerCertified || !a.DiracEdges.CouplingGraphOnly || a.DiracEdges.TraceMagnitudeReadoutCertified || a.DiracEdges.YukawaValuesCertified || a.DiracEdges.ObservedMassDataUsed {
		t.Fatalf("D_F edge support over-promoted: %s", FormatDiracEdges(a.DiracEdges))
	}
}

func TestGate835MatrixFirewallPullbackDeferralAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Matrix.MatrixUnitsExist || !a.Matrix.DiagonalProjectorsExist || a.Matrix.CanonicalColorFrameCertified || a.Matrix.BasisIndependentAtoms || a.Matrix.CanonicalColorAtomsCertified || a.Matrix.CompleteSectorLedger {
		t.Fatalf("matrix-unit color-frame firewall failed: %s", FormatMatrix(a.Matrix))
	}
	if a.Pullback.PiSectorFCodomainCertified || a.Pullback.PullbackAllowedToRun || a.Pullback.SigmaCertified || a.Pullback.TopI3PulledBack || a.Pullback.FockP1P3PulledBack || !a.Pullback.NonCircular {
		t.Fatalf("aggregate pullback was not deferred: %s", FormatPullback(a.Pullback))
	}
	if a.Impact.PiSectorFCertified || a.Impact.SigmaCertified || a.Impact.TraceMagnitudeReadoutCertified || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact over-promoted: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoCompletePiSectorF || !a.Firewalls.NoGammaRefinement || !a.Firewalls.NoJRefinement || !a.Firewalls.NoBimoduleLedger || !a.Firewalls.NoDFEdgeLedger || !a.Firewalls.PullbackPremature || !a.Firewalls.ProjectorsNotMagnitudes || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoCYukawaUpdate {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2FiniteRepresentationSectorProjectorLedgerConstructionObstructionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem construction failure: %+v", res)
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
