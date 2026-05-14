package flavorprojectionmetric

import (
	"math"
	"testing"
)

func TestProjectionMetricAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Metric.Formalized || len(a.Metric.TrialityVector) != 3 || math.Abs(vectorNorm(a.Metric.TrialityVector)-3) > 1e-12 {
		t.Fatalf("bad metric audit: %s", FormatMetric(a.Metric))
	}
	if math.Abs(a.Metric.PositiveEigenMin-1.0/9.0) > 1e-12 || math.Abs(a.Metric.PositiveEigenMax-4.0/9.0) > 1e-12 {
		t.Fatalf("bad positive eigenvalues: %s", FormatMetric(a.Metric))
	}
	if a.Metric.PhysicalMetricSelected || a.Metric.SignedRank != 1 || a.Metric.SignedNullity != 2 {
		t.Fatalf("bad metric selection/rank: %s", FormatMetric(a.Metric))
	}
}

func TestPositiveTraceMetricForbidsNulling(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Positive.Audited || a.Positive.ExactNullingPossible || !a.Positive.VariationalMinimumUnique {
		t.Fatalf("bad positive metric sieve: %s", FormatPositive(a.Positive))
	}
	if math.Abs(a.Positive.MinimumFraction-1.0/9.0) > 1e-12 {
		t.Fatalf("positive minimum must be 1/9: %s", FormatPositive(a.Positive))
	}
	if math.Abs(a.Positive.PredictedMassAtMinimumGeV-gate323UniqueLowTopMassGeV) > 1e-9 {
		t.Fatalf("unexpected inherited positive-lane mass: %s", FormatPositive(a.Positive))
	}
}

func TestSignedMetricAllowsButDoesNotSelectNulling(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Signed.Audited || !a.Signed.ExactNullingPossible || a.Signed.NullspaceDimension != 2 || a.Signed.VariationalMinimumUnique {
		t.Fatalf("bad signed metric sieve: %s", FormatSigned(a.Signed))
	}
	for _, v := range a.Signed.NullspaceBasis {
		if math.Abs(dot(a.Metric.NormalizedTriality, v)) > 1e-12 {
			t.Fatalf("signed nullspace vector not orthogonal: tau=%s v=%s", FormatVector(a.Metric.NormalizedTriality), FormatVector(v))
		}
	}
	if math.Abs(a.Signed.NullTopMassGeV-gate322RunningMassGeV) > 1e-9 {
		t.Fatalf("unexpected null-top mass: %s", FormatSigned(a.Signed))
	}
}

func TestVariationalFallbackAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Variational.Executed || a.Variational.PositiveLaneAuthorizesGate322 || a.Variational.SignedLaneAuthorizesGate322 || a.Variational.NativeSelectorInstalled || a.Variational.DynamicalPotentialInstalled {
		t.Fatalf("bad variational audit: %s", FormatVariational(a.Variational))
	}
	if !a.CKMFallback.Formalized || !a.CKMFallback.RequiresEmpiricalTexture || !a.CKMFallback.RequiresVacuumSelectorPotential {
		t.Fatalf("bad CKM fallback: %s", FormatCKM(a.CKMFallback))
	}
	if !a.RG.Audited || a.RG.PositiveMetricPreservesGate322 || !a.RG.SignedMetricPreservesGate322 || a.RG.PhysicalLaneAuthorized {
		t.Fatalf("bad RG audit: %s", FormatRG(a.RG))
	}
	if !a.Firewalls.NoCKMImported || !a.Firewalls.NoObservedTopMassInserted || !a.Firewalls.NoFlavorTextureInvented || !a.Firewalls.NoProjectionMetricForced || !a.Firewalls.NoPoleMassClaimed || !a.Firewalls.NoTwoLoopClaimed || !a.Firewalls.NoColliderMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
	if !a.Summary.PositiveMetricForbidsNulling || !a.Summary.SignedMetricAllowsNulling || a.Summary.NativeMetricDerived || a.Summary.NativeFlavorVacuumSelected || a.Summary.TopBoundarySuppressionJustified || a.Summary.Gate322PhysicalLaneAuthorized || !a.Summary.FirewallsPreserved || a.Summary.FinalMassClaimed {
		t.Fatalf("bad summary: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FlavorProjectionMetricVariationalVacuumSelectorAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
