package exponentialtauetatexture

import (
	"math"
	"testing"
)

func TestCanonicalExponentialTexture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.Formalized || len(a.Sieve.Results) < 4 {
		t.Fatalf("bad sieve: %s", FormatSieve(a.Sieve))
	}
	r := a.Sieve.Results[0]
	expected := math.Exp(2 * BGap)
	if !r.RankPreserved || !r.KineticSafe {
		t.Fatalf("canonical C12 not safe: %s", FormatTexture(r))
	}
	if math.Abs(r.FirstSecondRatio-expected) > 1e-9 {
		t.Fatalf("unexpected canonical 1-2 ratio: got %.12f want %.12f detail=%s", r.FirstSecondRatio, expected, FormatTexture(r))
	}
	if r.ObservedScaleMatch {
		t.Fatalf("canonical generator should not match observed hierarchy: %s", FormatTexture(r))
	}
}

func TestAmplifiedWitnessNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Sieve.Results[1]
	expected := math.Exp(10 * BGap)
	if a.Sieve.Generators[1].Canonical {
		t.Fatalf("amplified witness should not be canonical")
	}
	if math.Abs(r.FirstSecondRatio-expected) > 1e-8 {
		t.Fatalf("unexpected amplified 1-2 ratio: got %.12f want %.12f detail=%s", r.FirstSecondRatio, expected, FormatTexture(r))
	}
	if r.FirstSecondRatio < 2 {
		t.Fatalf("amplified witness should split more than canonical: %s", FormatTexture(r))
	}
}

func TestRequiredCoefficientsAreLarge(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Sieve.Results[0]
	if r.RequiredCoeffFor17 <= 10 || r.RequiredCoeffFor44 <= 15 || r.RequiredCoeffFor136 <= 20 || r.RequiredCoeffFor207 <= 25 {
		t.Fatalf("required coefficients not large enough: %s", FormatTexture(r))
	}
}

func TestCKMNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CKM.Audited || a.CKM.CKMDerived || a.CKM.NativeSectorChoice {
		t.Fatalf("bad CKM audit: %s", FormatCKM(a.CKM))
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
		StatusExponentialMapFormalized,
		StatusExponentialHierarchySieve,
		StatusSignInterferenceVerified,
		StatusTensionCanonicalGeneratorMild,
		StatusTensionLargeGeneratorNeeded,
		StatusFailedNativeGeneratorMagnitude,
		StatusFailedHierarchyNotDerived,
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
	res := ExponentialTauEtaTextureBGapMixingHierarchyAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
