package tauetargtexture

import (
	"math"
	"testing"
)

func TestSeedAndNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seed.Formalized || a.Seed.AbsoluteTau != [3]float64{2, 2, 1} || a.Seed.SignsAffectRG {
		t.Fatalf("bad seed: %s", FormatSeed(a.Seed))
	}
	if !a.Normalization.Formalized || !a.Normalization.OverallScaleFree || a.Normalization.XDerived {
		t.Fatalf("bad normalization: %s", FormatNormalization(a.Normalization))
	}
	got := (a.Normalization.ExampleYu0*a.Normalization.ExampleYu0 + a.Normalization.ExampleYd0*a.Normalization.ExampleYd0) / (a.Normalization.ExampleYe0 * a.Normalization.ExampleYe0)
	if math.Abs(got-rPlus) > 1e-12 {
		t.Fatalf("r+ witness not satisfied: got %.15f want %.15f", got, rPlus)
	}
}

func TestRGPreservesDegeneracyAndOrdering(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.RG.Executed || !a.RG.DegeneracyPreserved || a.RG.OrderingInverted {
		t.Fatalf("bad RG audit: %s", FormatRG(a.RG))
	}
	for _, run := range a.RG.Runs {
		if !run.Perturbative {
			t.Fatalf("nonperturbative run: %s", FormatRun(run))
		}
		if math.Abs(run.UpFirstSecondSplit) > 1e-10 || math.Abs(run.DownFirstSecondSplit) > 1e-10 || math.Abs(run.LeptonFirstSecondSplit) > 1e-10 {
			t.Fatalf("first/second degeneracy not preserved: %s", FormatRun(run))
		}
		if run.UpHighLowRatio < 1.5 || run.UpHighLowRatio > 2.5 {
			t.Fatalf("unexpected up high/low ratio: %s", FormatRun(run))
		}
	}
}

func TestNoObservedHierarchyGenerated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.RG.MatchesOrderOfMagnitude || a.RG.BestHighLowRatio >= 3 || a.RG.ParameterReduction != 0 || a.RG.ReductionProved {
		t.Fatalf("hierarchy should not be generated: %s", FormatRG(a.RG))
	}
}

func TestSignTextureRequiresOffDiagonalOperator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.SignTexture
	if !s.Formalized || s.SignVisibleInDiagonalRG || !s.NeedsOffDiagonalTexture || s.CKMReductionProved {
		t.Fatalf("bad sign texture audit: %s", FormatSignTexture(s))
	}
}

func TestCensusStillFifteen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.StartingVacuumInputs != 15 || a.Census.TotalReduction != 0 || a.Census.RemainingInputs != 15 || a.Census.SevenSealReached {
		t.Fatalf("bad census: %s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{
		StatusSeedFormalized,
		StatusRGTextureExecuted,
		StatusDegeneracyPreserved,
		StatusFailedHierarchyNotGenerated,
		StatusFailedFirstSecondNotSplit,
		StatusFailedCKMTextureMissing,
		StatusFailedNoReduction,
	}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := TauEtaDiagonalTextureRGEvolutionMassHierarchyAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
