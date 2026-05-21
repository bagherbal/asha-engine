package generation2finitetriplerepresentationcompletionandprojectorledgerdataaudit

import (
	"strings"
	"testing"
)

func TestGate836FiniteTripleDataIsMissingAndNotInvented(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Data.AlgebraKnown || a.Data.ExplicitHF || a.Data.ExplicitRhoF || a.Data.ExplicitJF || a.Data.ExplicitGammaF || a.Data.ExplicitDF || a.Data.CompletePackageCertified || a.Data.CanConstructPiSectorF || a.Data.ObservedDataUsed {
		t.Fatalf("finite triple data over/under-certified: %s", FormatData(a.Data))
	}
	if !containsAll(a.Data.Failures, []string{FailureNoCompleteFiniteTripleRepresentationData, FailureNoExplicitHFCarrier, FailureNoExplicitRhoFRepresentation, FailureNoExplicitJFRealStructure, FailureNoExplicitGammaFChirality, FailureNoExplicitDFOperator}) {
		t.Fatalf("missing finite triple data failures: %s", strings.Join(a.Data.Failures, ","))
	}
}

func TestGate836CentralRanksRefinementsBimoduleAndEdgesAreBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Central.CentralIdempotentsOrthogonal || !a.Central.CentralIdempotentsSumI || a.Central.RhoFExplicit || a.Central.SupportProjectorsInstantiated || a.Central.SupportRanksCertified || a.Central.OrthogonalityCertified || a.Central.CompletenessCertified || a.Central.RankLedgerCertified || !a.Central.CoarseRecipeOnly || a.Central.CompleteLedger {
		t.Fatalf("central rank ledger should remain unavailable: %s", FormatCentral(a.Central))
	}
	if a.Chirality.GammaFExplicit || a.Chirality.JFExplicit || a.Chirality.ChiralityProjectorsInstantiated || a.Chirality.RealStructureImagesInstantiated || a.Chirality.LeftRightSplitCertified || a.Chirality.ParticleOppositeSplitCertified || a.Chirality.YukawaMagnitudeCertified || a.Chirality.ObservedParticleAssignment {
		t.Fatalf("chirality/J refinement over-promoted: %s", FormatChirality(a.Chirality))
	}
	if !a.Bimodule.RequiresLeftAction || !a.Bimodule.RequiresRightAction || a.Bimodule.RhoFExplicit || a.Bimodule.JFExplicit || a.Bimodule.DFExplicit || a.Bimodule.BimoduleStabilityCertified || a.Bimodule.FirstOrderCompatibilityCertified || a.Bimodule.TypedProjectorLedgerCertified {
		t.Fatalf("bimodule/first-order typing over-promoted: %s", FormatBimodule(a.Bimodule))
	}
	if !a.Edges.RequiresPiSectorF || !a.Edges.RequiresDF || a.Edges.PiSectorFExists || a.Edges.DFExplicit || a.Edges.EdgeBlocksComputed || a.Edges.EdgeSupportGraphCertified || !a.Edges.CouplingGraphOnly || a.Edges.TraceMagnitudeReadoutCertified || a.Edges.YukawaValuesCertified || a.Edges.ObservedDataUsed {
		t.Fatalf("D_F edge graph over-promoted: %s", FormatEdges(a.Edges))
	}
}

func TestGate836ColorFrameImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ColorFrame.M3MatrixUnitsExist || !a.ColorFrame.DiagonalProjectorsExist || a.ColorFrame.CanonicalFrameCertified || a.ColorFrame.BasisIndependentAtoms || a.ColorFrame.ColorAtomLedgerCertified || a.ColorFrame.GaugeFrameChoiceUsed {
		t.Fatalf("M3 color-frame firewall failed: %s", FormatColorFrame(a.ColorFrame))
	}
	if a.Impact.FiniteTripleDataComplete || a.Impact.PiSectorFConstructible || a.Impact.PiSectorFCertified || a.Impact.SigmaAllowed || a.Impact.TraceMagnitudeAllowed || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || !strings.Contains(a.Impact.NextRequiredObject, "FiniteRepresentationDataSeal") {
		t.Fatalf("impact over-promoted: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoCompleteFiniteTripleData || !a.Firewalls.NoExplicitHF || !a.Firewalls.NoExplicitRhoF || !a.Firewalls.NoExplicitJF || !a.Firewalls.NoExplicitGammaF || !a.Firewalls.NoExplicitDF || !a.Firewalls.NoPiSectorF || !a.Firewalls.PullbackPremature || !a.Firewalls.ProjectorsNotMagnitudes || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoCYukawaUpdate || a.Firewalls.Verdict != StatusFirewallGate836 {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2FiniteTripleRepresentationCompletionAndProjectorLedgerDataAuditTheorem().Verify()
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
