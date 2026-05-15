package generation2finiteactionsecondvariation

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate495FiniteActionSecondVariation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.TraceNormalizationDoesNotSelectKappa || !a.Inheritance.FiniteActionSecondVariationRequested {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Candidate.SecondVariationComputed || !a.Candidate.BrokenDiag114 || !a.Candidate.KappaSixSelectedInCandidate || a.Candidate.KappaU1 != 6 {
		t.Fatalf("bad canonical candidate: %+v", a.Candidate)
	}
	if !a.Candidate.FullGaugeHessianPositive || a.Candidate.FullGaugeHessianRank != 4 || a.Candidate.PhysicalCouplingsDerived || a.Candidate.PhysicalMassesDerived {
		t.Fatalf("bad Hessian candidate/firewall: %+v", a.Candidate)
	}
	if a.Provenance.NativeActionProvenanceClosed || a.Provenance.NativeKappaSelectionClosed || a.Provenance.NativeGaugeHessianSelectionClosed {
		t.Fatalf("provenance over-promoted candidate: %+v", a.Provenance)
	}
	if !a.Provenance.CanonicalCandidateUsesDiagnosticDphi || !a.Provenance.CanonicalCandidateUsesI4Metric || !a.Provenance.CanonicalCandidateUsesChosenVacuum {
		t.Fatalf("expected diagnostic provenance dependencies: %+v", a.Provenance)
	}
	if !a.Boundary.DimensionlessSecondVariationCandidate || a.Boundary.NativeKappaSelected || a.Boundary.NativeGaugeCouplingsDerived || a.Boundary.NativeWZMassesDerived {
		t.Fatalf("bad boundary: %+v", a.Boundary)
	}
	if a.Firewall.WeakAngleImported || a.Firewall.NativeKappaWritten || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate495RenderAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 495 Registry Audit",
		StatusCanonicalActionCandidateFound,
		StatusKappaSixSelectedInsideCandidate,
		StatusFailedCanonicalActionNotNativeClosed,
		StatusNativeRegistryWriteBlocked,
		"diag(1,1,4)",
		"kappa_U1 = 6",
		"Gate 496",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
