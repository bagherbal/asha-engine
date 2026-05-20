package generation2complexd4trilinearinvariantandyukawareadoutobstructionaudit

import (
	"strings"
	"testing"
)

func TestGate802InheritanceAndComplexCarriers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.VolumeSquare != -1 || a.Inheritance.RealChiralityCertified || !strings.Contains(a.Inheritance.TrialityLevel, "T1") {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Carriers.Defined || a.Carriers.VectorDimC != 8 || a.Carriers.SpinorPlusDimC != 8 || a.Carriers.SpinorMinusDimC != 8 || a.Carriers.GenerationCopies || a.Carriers.NativeRealCarriers {
		t.Fatalf("bad complex carriers: %s", FormatCarriers(a.Carriers))
	}
}

func TestGate802TrilinearMultiplicityAndCovariance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Trilinear.Defined || !a.Trilinear.Equivariant || !a.Trilinear.NonZero || !containsAll(a.Trilinear.Failures, []string{StatusTD4NotSMYukawa, StatusTD4NotTraceLedger}) {
		t.Fatalf("bad trilinear: %s", FormatTrilinear(a.Trilinear))
	}
	if !a.Multiplicity.Audited || a.Multiplicity.HomDimension != 1 || !a.Multiplicity.CanonicalShapeUpToScale || a.Multiplicity.DeterminesEigenvalues || a.Multiplicity.DeterminesHierarchy {
		t.Fatalf("bad multiplicity: %s", FormatMultiplicity(a.Multiplicity))
	}
	if !a.Covariance.Audited || !a.Covariance.CyclicStable || a.Covariance.GenerationTriplication || a.Covariance.MixingReadout {
		t.Fatalf("bad covariance: %+v", a.Covariance)
	}
}

func TestGate802ReadoutObstructions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Readout.Defined || !containsAll(a.Readout.Failures, []string{StatusTrilinearNoSectorOps, StatusTrilinearNoAtoms, StatusTrilinearNoNEff}) {
		t.Fatalf("bad readout requirements: %s", FormatReadout(a.Readout))
	}
	for name, obs := range map[string]Obstruction{"sector": a.Sector, "generation": a.Generation, "positivity": a.Positivity, "top": a.TopDominance, "gj": a.GeorgiJarlskog} {
		if !obs.Audited || len(obs.Failures) == 0 {
			t.Fatalf("bad obstruction %s: %s", name, FormatObstruction(obs))
		}
	}
	if !a.RealForm.Preserved || !containsAll(a.RealForm.Failures, []string{StatusComplexTD4NotRealCL17, StatusNoRealDescentReadout}) {
		t.Fatalf("bad real-form obstruction: %s", FormatRealForm(a.RealForm))
	}
}

func TestGate802LawfulUseCHiggsBranchAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Lawful.Recorded || !containsAll(a.Lawful.Supports, []string{StatusTD4UsefulAirlocked}) || !containsAll(a.Lawful.Failures, []string{StatusTD4NotPhysicalScalarInput}) {
		t.Fatalf("bad lawful use: %s", FormatLawfulUse(a.Lawful))
	}
	if !a.CHiggs.Preserved || !containsAll(a.CHiggs.Failures, []string{StatusTD4NoCYukawaUpdate, StatusCHiggsStillLevelB}) {
		t.Fatalf("bad C_Higgs firewall: %s", FormatCHiggs(a.CHiggs))
	}
	if !strings.Contains(a.Branch.NextNative, "Gate 803") || !strings.Contains(a.Branch.NextNative, "Triality-to-Yukawa") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2ComplexD4TrilinearInvariantAndYukawaReadoutObstructionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
