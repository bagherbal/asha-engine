package finiteyukawaaction

import (
	"strings"
	"testing"
)

func TestGate263FiniteYukawaActionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.HermitianTrialityBasisExposed || !a.Inheritance.RawNonCommutingPartnerExists || a.Inheritance.PreviousQualifiedPartnerFound {
		t.Fatalf("bad Gate 262 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if len(a.TraceAudits) != 3 || !a.Summary.TraceFunctionalsEvaluated {
		t.Fatalf("trace audits missing: %s", FormatSummary(a.Summary))
	}
	if a.TraceAudits[0].NonZeroOnBasis {
		t.Fatalf("linear trace should vanish on off-diagonal basis: %s", FormatTraceAudit(a.TraceAudits[0]))
	}
	if a.TraceAudits[1].RealBasisValue != 6 || a.TraceAudits[1].PhaseBasisValue != 6 || a.TraceAudits[1].CrossValue != 0 || a.TraceAudits[1].DistinguishesRealAndPhase {
		t.Fatalf("bad Hilbert-Schmidt audit: %s", FormatTraceAudit(a.TraceAudits[1]))
	}
	if a.TraceAudits[2].RealBasisValue != 52 || a.TraceAudits[2].PhaseBasisValue != 52 || a.TraceAudits[2].DistinguishesRealAndPhase {
		t.Fatalf("bad commutator norm audit: %s", FormatTraceAudit(a.TraceAudits[2]))
	}
	if a.Summary.NativeActionCandidateCount != 5 || a.Summary.ActionCandidateQualified || a.Summary.FiniteYukawaActionDerived {
		t.Fatalf("bad action candidate summary: %s", FormatSummary(a.Summary))
	}
	if a.ScalarPhase.BGapCanWeightTrialityBasis || a.ScalarPhase.HopfCanFixCPPhase || a.ScalarPhase.ScalarPhaseIntegrationDerived {
		t.Fatalf("scalar/phase integration should remain blocked: %s", FormatScalarPhase(a.ScalarPhase))
	}
	if a.Texture.FiniteActionCoefficientRule || a.Texture.PhysicalTextureConstructed || !a.Texture.EmpiricalYukawaSealRequired {
		t.Fatalf("texture should remain an unselected ansatz: %s", FormatTexture(a.Texture))
	}
	if !a.Firewall.DoesNotPromoteTraceMetricToDynamics || !a.Firewall.DoesNotUseBGapWithoutMap || !a.Firewall.DoesNotUseHopfWithoutProjection || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate263StatusLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	statuses := StatusLines(a)
	for _, want := range []string{
		StatusGate262Inherited,
		StatusTraceFunctionalsEvaluated,
		StatusTraceMetricDegenerate,
		StatusCanonicalActionNoMixing,
		StatusSpectralActionNotReady,
		StatusBGapNoActionMap,
		StatusHopfNoProjection,
		StatusNoFiniteYukawaAction,
		StatusEmpiricalYukawaSealPreserved,
		StatusPhysicalTextureStillBlocked,
		StatusCKMPMNSMassesStillBlocked,
	} {
		if !containsLine(statuses, want) {
			t.Fatalf("status %q missing from ledger:\n%s", want, statuses)
		}
	}
}

func containsLine(haystack, needle string) bool {
	for _, line := range strings.Split(haystack, "\n") {
		if line == needle {
			return true
		}
	}
	return false
}
