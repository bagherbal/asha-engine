package matrixinvariantkoideaudit

import (
	"math"
	"testing"
)

func TestKoideFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Koide.Formalized || math.Abs(a.Koide.Target-2.0/3.0) > 1e-15 || a.Koide.RootAngleDegrees != 45 {
		t.Fatalf("bad koide formalization: %s", FormatKoide(a.Koide))
	}
}

func TestChargedLeptonKoideQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.EmpiricalSpectra[0]
	if !s.Quarantined || s.KoideK < 0.666 || s.KoideK > 0.667 || math.Abs(s.RelativeDeviation) > 1e-4 {
		t.Fatalf("charged lepton comparison should align but remain quarantined: %s", FormatSpectrum(s))
	}
}

func TestTrialityDoesNotMandateKoide(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	tr := a.Triality
	if !tr.Executed || tr.NativeTwoThirdsMandated || tr.MagnitudeSquaredKoide >= 0.5 || math.Abs(tr.MagnitudeSquaredKoide-0.36) > 1e-12 {
		t.Fatalf("triality should not mandate K=2/3: %s", FormatTriality(tr))
	}
}

func TestCharacteristicNeedsRootTraceOperator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Characteristic
	if !c.Audited || !c.OneConstraintCapacity || !c.RequiresRootTraceOperator || c.CharacteristicPolynomialLocked {
		t.Fatalf("bad characteristic audit: %s", FormatCharacteristic(c))
	}
}

func TestNoParameterReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Reduction
	if r.StartingVacuumInputs != 15 || r.ReductionProved != 0 || r.RemainingVacuumInputs != 15 || r.SevenSealTargetReached {
		t.Fatalf("bad reduction census: %s", FormatReduction(r))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{
		StatusKoideInvariantFormalized,
		StatusTrialityInvariantSieveExecuted,
		StatusEmpiricalKoideAlignmentCataloged,
		StatusCharacteristicPolynomialAudited,
		StatusParameterReductionAssessed,
		StatusFailedMatrixTraceInvariantNotDerived,
		StatusFailedKoideConstraintNotNative,
		StatusFailedNoVacuumReductionProved,
		StatusFailedSevenCoordinatesNotReached,
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
	res := MatrixInvariantKoideTypeTracePolynomialAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
