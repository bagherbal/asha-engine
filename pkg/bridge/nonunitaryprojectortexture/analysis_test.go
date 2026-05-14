package nonunitaryprojectortexture

import (
	"math"
	"testing"
)

func TestProjectorCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Search.Formalized || len(a.Search.Candidates) != 3 {
		t.Fatalf("bad projector search: %s", FormatSearch(a.Search))
	}
	ray := a.Search.Candidates[0]
	if ray.Name != "signed tau ray projector" || ray.Rank != 1 || !ray.RankDefect || ray.KineticSafe {
		t.Fatalf("bad ray projector: %s", FormatCandidate(ray))
	}
	q := a.Search.Candidates[1]
	if q.Name != "signed tau null-complement projector" || q.Rank != 2 || !q.RankDefect || q.SingularValues[2] != 0 {
		t.Fatalf("bad null complement: %s", FormatCandidate(q))
	}
}

func TestProjectorsAreNotKineticSafe(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Safety.Audited || a.Safety.NativeRepairDerived {
		t.Fatalf("bad kinetic safety audit: %s", FormatSafety(a.Safety))
	}
	if !a.Search.AnyRankDefect || a.Search.AnyKineticSafeHierarchy {
		t.Fatalf("bad rank-defect/hierarchy flags: %s", FormatSearch(a.Search))
	}
}

func TestNumericalProjectorRatios(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	ray := a.Search.Candidates[0]
	if math.Abs(ray.SingularValues[0]-math.Sqrt(33)/3) > 1e-12 {
		t.Fatalf("unexpected ray singular value: %s", FormatCandidate(ray))
	}
	q := a.Search.Candidates[1]
	if math.Abs(q.SingularValues[0]-2) > 1e-12 || math.Abs(q.SingularValues[1]-2/math.Sqrt(3)) > 1e-12 {
		t.Fatalf("unexpected null complement values: %s", FormatCandidate(q))
	}
}

func TestNoParameterReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.TotalReduction != 0 || a.Census.RemainingInputs != 15 || a.Census.SevenSealReached {
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
		StatusProjectorSearchFormalized,
		StatusTauRayProjectorAudited,
		StatusTauNullComplementAudited,
		StatusRankDefectDetected,
		StatusKineticSafetyAudited,
		StatusTensionProjectorsSplitButDestroyRank,
		StatusFailedKineticSafeTextureNotDerived,
		StatusFailedNativeProjectorTextureNotDerived,
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
	res := NonUnitaryProjectorKineticSafeFlavorTextureSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
