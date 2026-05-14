package topologicalamplifierflavorsector

import (
	"math"
	"testing"
)

func TestTraceCapacityAmplifierMagnitude(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Sieve.Textures[0]
	want := math.Exp(2 * BGap * CTrace)
	if math.Abs(r.SplitPairRatio-want) > 1e-8 {
		t.Fatalf("bad split: got %.12f want %.12f detail=%s", r.SplitPairRatio, want, FormatTexture(r))
	}
	if !(r.MatchesTopCharm && r.MatchesMuonElectron) {
		t.Fatalf("C_trace branch should match O(10^2) hierarchy band: %s", FormatTexture(r))
	}
}

func TestEightPiAmplifierMagnitude(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Sieve.Textures[3]
	want := math.Exp(2 * BGap * 8 * math.Pi)
	if math.Abs(r.SplitPairRatio-want) > 1e-8 {
		t.Fatalf("bad 8pi split: got %.12f want %.12f detail=%s", r.SplitPairRatio, want, FormatTexture(r))
	}
	if r.SplitPairRatio < 100 || r.SplitPairRatio > 250 {
		t.Fatalf("8pi branch should be in O(10^2) band: %s", FormatTexture(r))
	}
}

func TestAmplifierNotPromotedToFlavorNorm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, amp := range a.Sieve.Amplifiers {
		if amp.NativeAsFlavorNorm {
			t.Fatalf("amplifier should not be promoted to flavor norm: %s", FormatAmplifier(amp))
		}
	}
}

func TestSectorAssignmentNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sector.Audited || !a.Sector.CKMCapacity || a.Sector.NativeAssignment || a.Sector.CKMDerived {
		t.Fatalf("bad sector audit: %s", FormatSector(a.Sector))
	}
}

func TestCensusNoReduction(t *testing.T) {
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
	required := []string{StatusAmplifierFormalized, StatusTrace25Audited, StatusEightPiAudited, StatusAmplifierMatchesScale, StatusTensionAmplifierFlavorCoupling, StatusFailedTopologicalAmplifierDerived, StatusFailedSectorGeneratorsDerived, StatusFailedVacuumReduced}
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
	res := TopologicalAmplifierBimoduleFlavorSectorSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
