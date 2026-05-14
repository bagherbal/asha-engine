package flavororientationoperator

import (
	"math"
	"testing"
)

func TestBasisAndOperatorSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Basis.Formalized || len(a.Basis.NormalizedSource) != 3 || math.Abs(a.Basis.TrialityNorm-3) > 1e-12 {
		t.Fatalf("bad basis: %s", FormatBasis(a.Basis))
	}
	if !a.Operator.SieveFormalized || a.Operator.InstalledNativeUnitary || a.Operator.JSwapActsOnFlavor || a.Operator.DoubledSpaceActsOnFlavor || a.Operator.BimoduleOverlapActsOnFlavor {
		t.Fatalf("bad operator audit: %s", FormatOperator(a.Operator))
	}
}

func TestNullspaceSuppressionCapacity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Nullspace.Computed || a.Nullspace.Dimension != 2 || !a.Nullspace.AllBasisVectorsOrthogonal || !a.Nullspace.TopSuppressionPossible || a.Nullspace.UniquePhysicalTopVector {
		t.Fatalf("bad nullspace audit: %s", FormatNullspace(a.Nullspace))
	}
	for _, v := range a.Nullspace.Basis {
		if math.Abs(dot(a.Basis.NormalizedSource, v)) > 1e-12 {
			t.Fatalf("nullspace vector not orthogonal: source=%s v=%s", FormatVector(a.Basis.NormalizedSource), FormatVector(v))
		}
	}
}

func TestFlavorCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Candidates) != 5 {
		t.Fatalf("expected 5 candidates, got %d", len(a.Candidates))
	}
	if math.Abs(a.Candidates[0].TopFraction-4.0/9.0) > 1e-12 || math.Abs(a.Candidates[2].TopFraction-1.0/9.0) > 1e-12 {
		t.Fatalf("bad direct fractions: %s | %s", FormatCandidate(a.Candidates[0]), FormatCandidate(a.Candidates[2]))
	}
	if a.Candidates[3].TopFraction > 1e-12 || a.Candidates[4].TopFraction > 1e-12 || !a.Candidates[3].SuppressesTopBoundary || !a.Candidates[4].SuppressesTopBoundary {
		t.Fatalf("bad null candidates: %s | %s", FormatCandidate(a.Candidates[3]), FormatCandidate(a.Candidates[4]))
	}
	if a.Candidates[3].Native || a.Candidates[4].Native || a.Candidates[3].Unique || a.Candidates[4].Unique {
		t.Fatalf("null candidates must remain non-native/non-unique: %s | %s", FormatCandidate(a.Candidates[3]), FormatCandidate(a.Candidates[4]))
	}
}

func TestRGCompatibilityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.RG.Audited || !a.RG.NullRotationPreservesGate322 || a.RG.NativeJustificationInstalled || a.RG.PhysicalLaneAuthorized {
		t.Fatalf("bad RG audit: %s", FormatRG(a.RG))
	}
	if math.Abs(a.RG.NullTopMassGeV-gate322RunningMassGeV) > 1e-9 {
		t.Fatalf("unexpected null-top mass: %s", FormatRG(a.RG))
	}
	if !a.Firewalls.NoCKMImported || !a.Firewalls.NoObservedTopMassInserted || !a.Firewalls.NoFlavorTextureInvented || !a.Firewalls.NoPoleMassClaimed || !a.Firewalls.NoTwoLoopClaimed || !a.Firewalls.NoColliderMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
	if !a.Summary.NullspaceCapacityProved || a.Summary.NativeFlavorOperatorDerived || a.Summary.TopBoundarySuppressionJustified || a.Summary.Gate322PhysicalLaneAuthorized || !a.Summary.FirewallsPreserved || a.Summary.FinalMassClaimed {
		t.Fatalf("bad summary: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FlavorOrientationOperatorTrialityToMassEigenstateTextureAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
