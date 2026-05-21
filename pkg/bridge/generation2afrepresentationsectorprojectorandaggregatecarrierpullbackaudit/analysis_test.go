package generation2afrepresentationsectorprojectorandaggregatecarrierpullbackaudit

import (
	"strings"
	"testing"
)

func TestGate834CentralIdempotentsAndRepresentationRequirement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Central.Orthogonal || !a.Central.SumToIdentity || !a.Central.CoarseSectorBlocks || a.Central.SummandCount != 3 || a.Central.SectorLedgerCertified || a.Central.TraceMagnitudeCertified {
		t.Fatalf("central idempotents over/under-certified: %s", FormatCentral(a.Central))
	}
	if !a.Representation.UsesHF || !a.Representation.UsesRhoF || !a.Representation.PartialPredataAvailable || a.Representation.CompletePackageCertified || a.Representation.RepresentationInducedProjectorsCertified || !a.Representation.CanSourceCoarseProjectorCandidates || a.Representation.CanSourceSectorLedger || a.Representation.CanSourceTraceMagnitudes {
		t.Fatalf("representation requirement over/under-certified: %s", FormatRepresentation(a.Representation))
	}
	if !containsAll(a.Representation.Failures, []string{FailureNoCompleteRhoFRepresentationCertified, FailureNoCompleteFiniteHilbertPackage}) {
		t.Fatalf("missing representation failures: %s", strings.Join(a.Representation.Failures, ","))
	}
}

func TestGate834MatrixUnitsAndAggregatePullbackObstructed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MatrixUnits.MatrixUnitsExist || !a.MatrixUnits.DiagonalProjectorsExist || a.MatrixUnits.CanonicalColorFrameCertified || a.MatrixUnits.BasisIndependent || a.MatrixUnits.CanonicalColorAtomsCertified || !a.MatrixUnits.SuppliesCarrierProjectors || a.MatrixUnits.SuppliesSectorLedger {
		t.Fatalf("matrix-unit firewall failed: %s", FormatMatrixUnits(a.MatrixUnits))
	}
	if !containsAll(a.MatrixUnits.Failures, []string{FailureM3MatrixUnitsNotCanonicalColorAtomsWithoutFrame, FailureNoCanonicalColorFrame}) {
		t.Fatalf("missing matrix-unit failures: %s", strings.Join(a.MatrixUnits.Failures, ","))
	}
	if !a.Pullback.CentralBlocksAvailable || !a.Pullback.RepresentationProjectorRecipe || !a.Pullback.NonCircular || a.Pullback.PullbackCertified || a.Pullback.TopI3PulledBack || a.Pullback.FockP1P3PulledBack || a.Pullback.M3P3IntertwinerCertified {
		t.Fatalf("aggregate pullback over-promoted: %s", FormatPullback(a.Pullback))
	}
	if !containsAll(a.Pullback.Failures, []string{FailureNoAggregateCarrierToRepresentationProjectorPullback, FailureNoSigmaMap, FailureTopI3NotPulledBackToRepresentationSector, FailureFockP1P3NotPulledBackToRepresentationSector}) {
		t.Fatalf("missing pullback failures: %s", strings.Join(a.Pullback.Failures, ","))
	}
}

func TestGate834ImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Impact.CentralProjectorSource || !a.Impact.RepresentationProjectorRecipe || a.Impact.AggregatePullbackCertified || a.Impact.SectorProjectorMapCertified || a.Impact.SectorTraceLedgerCertified || a.Impact.TraceMagnitudeReadoutCertified {
		t.Fatalf("impact over-promoted: %s", FormatImpact(a.Impact))
	}
	if a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger update or R3/R4 promotion allowed: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AFAloneNotLedger || !a.Firewalls.RequiresRepresentation || !a.Firewalls.NoCompleteRhoF || !a.Firewalls.MatrixUnitsBasisDependent || !a.Firewalls.NoAggregatePullback || !a.Firewalls.ProjectorsNotMagnitudes || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoCYukawaUpdate {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2AFRepresentationSectorProjectorAndAggregateCarrierPullbackAuditTheorem().Verify()
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
