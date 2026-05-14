package fermionicroottracesieve

import (
	"math"
	"testing"
)

func TestPfaffianRootDeterminantNotRootTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Pfaffian
	if !p.Executed || p.RootDeterminant <= 0 || p.RootTrace <= 0 || p.PfaffianCanGenerateKoide {
		t.Fatalf("bad pfaffian sieve: %s", FormatPfaffian(p))
	}
	if math.Abs(p.KoideK-2.0/3.0) > 1e-4 {
		t.Fatalf("charged lepton Koide comparison should be close but quarantined: %s", FormatPfaffian(p))
	}
}

func TestRootTraceOperatorNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.RootTrace
	if !r.Audited || !r.BosonicEvenTraceBarrier || r.RootTraceNative {
		t.Fatalf("root trace should remain non-native: %s", FormatRootTrace(r))
	}
}

func TestDixmierTraceDoesNotLockFiniteYukawaRootTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Dixmier
	if !d.Audited || !d.FiniteRankYukawaDixmierZero || d.LocksYukawaRootTrace || d.ContactVolumeF0 != 7 {
		t.Fatalf("bad Dixmier audit: %s", FormatDixmier(d))
	}
}

func TestNoKoidePromotionOrParameterReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.KoidePromotion.EmpiricalAlignment || a.KoidePromotion.NativePromotion {
		t.Fatalf("Koide should align empirically but not promote: %s", FormatKoidePromotion(a.KoidePromotion))
	}
	if a.Census.AdditionalReduction != 0 || a.Census.RemainingVacuumInputs != 15 || a.Census.SevenSealTargetReached {
		t.Fatalf("parameter census should remain unchanged: %s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{
		StatusFermionicEffectiveActionFormalized,
		StatusPfaffianSieveExecuted,
		StatusRootTraceOperatorAudited,
		StatusDixmierContactTraceAudited,
		StatusKoidePromotionSieveExecuted,
		StatusFailedRootTraceNotDerived,
		StatusFailedFermionicPfaffianNoKoide,
		StatusFailedMatrixInvariantNotPromoted,
		StatusFailedNoAdditionalReduction,
		StatusFailedSevenCoordinatesNotProved,
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
	res := FermionicEffectiveActionRootTracePfaffianSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
