package generation2leptocolorfiniteactionprojectorledgeraudit

import (
	"strings"
	"testing"
)

func TestGate838WCarrierAndBMinusLInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.W.Dim != 4 || a.W.P1Rank != 1 || a.W.P3Rank != 3 || !a.W.P1P3Orthogonal || !a.W.P1PlusP3CompletesW || !a.W.BMinusLTraceZero || !a.W.P3WIsM3Fundamental {
		t.Fatalf("bad W carrier: %s", FormatW(a.W))
	}
	if a.W.ColorAtomsCanonical || a.W.CanonicalColorFrame {
		t.Fatalf("color atoms over-certified: %s", FormatW(a.W))
	}
}

func TestGate838RhoActionSealConsistency(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Rho.M3ActsOnP3W || !a.Rho.M3TrivialOnP1W || !a.Rho.HActsOnLeftDoubleSocket || !a.Rho.CActsOnRightSocketPair || !a.Rho.ActionPreservesP1P3 || !a.Rho.ActionPreservesBMinusL || !a.Rho.RepresentationLawConsistentAtBlockLevel {
		t.Fatalf("rho action not consistent: %s", FormatRho(a.Rho))
	}
	if a.Rho.NativeDerivationCertified || a.Rho.ExplicitMatricesCertified || a.Rho.FirstOrderConditionCertified || a.Rho.BimoduleCommutantProof || a.Rho.CompleteRhoFActionLedger {
		t.Fatalf("rho action over-certified: %s", FormatRho(a.Rho))
	}
	if !containsAll(a.Rho.Failures, []string{FailureRepresentationSealNotNative, FailureNoFullFiniteTripleProof, FailureNoExplicitMatrices}) {
		t.Fatalf("missing rho failures: %s", strings.Join(a.Rho.Failures, ","))
	}
}

func TestGate838CoarseProjectorLedgerRanks(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int{"Pi_R1": 2, "Pi_R3": 6, "Pi_L1": 2, "Pi_L3": 6}
	for _, p := range a.Projectors.ParticleProjectors {
		if want[p.Name] != p.Rank {
			t.Fatalf("bad projector rank for %s: got %d want %d", p.Name, p.Rank, want[p.Name])
		}
	}
	if a.Projectors.ParticleRankSum != 16 || a.Projectors.HFProjectorRankSum != 32 || !a.Projectors.Orthogonal || !a.Projectors.CompleteOnHPart || !a.Projectors.JCopyIncluded || !a.Projectors.CoarsePiSectorFSealCertified {
		t.Fatalf("bad projector ledger: %s", FormatProjectors(a.Projectors))
	}
	if a.Projectors.FullNativePiSectorFCertified || a.Projectors.TraceMagnitudeReadoutCertified || a.Projectors.FineColorAtomLedgerCertified {
		t.Fatalf("projector ledger over-certified: %s", FormatProjectors(a.Projectors))
	}
}

func TestGate838DFCompressionImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DF.SymbolicEdgeSupportAudited || !a.DF.CouplingGraphOnly || a.DF.NumericalMagnitudes || a.DF.UsesObservedMasses || a.DF.UsesCKM || a.DF.UsesPMNS {
		t.Fatalf("D_F edge skeleton invalid: %s", FormatDF(a.DF))
	}
	if !a.Compression.SectorBodyBeforeCompression || a.Compression.AggregateCompressionMapCertified || a.Compression.AggregateToSectorPullbackCertified || a.Compression.AggregateOperatorIsSectorLedger || a.Compression.TraceMagnitudeReadoutCertified {
		t.Fatalf("compression stance invalid: %s", FormatCompression(a.Compression))
	}
	if !a.Impact.RhoActionSealConstructed || !a.Impact.CoarseLedgerConstructed || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact invalid: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoCompressionMap || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoThreeGeneration || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate838 {
		t.Fatalf("firewall invalid: %+v", a.Firewalls)
	}
	res := Generation2LeptoColorFiniteActionProjectorLedgerAuditTheorem().Verify()
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
