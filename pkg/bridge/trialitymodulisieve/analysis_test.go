package trialitymodulisieve

import "testing"

func TestDomainAdmissionRejectedWithoutNativeCarrier(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Domain.AbstractSpin8TrialityAvailable || a.Domain.DomainAdmitted || a.Domain.NativeTrialityCarrierFound || a.Domain.GenerationToTrialityFunctorDerived || !a.Domain.ManualGenerationRelabelingRejected {
		t.Fatalf("bad domain audit:\n%s", FormatDomain(a.Domain))
	}
}

func TestEquivariantCentralizerDimensions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c3 := centralizerByName(a.Centralizer.Cases, "exact C3 label-triality stress test")
	s3 := centralizerByName(a.Centralizer.Cases, "exact S3 label-triality stress test")
	if c3.GeneralComplexRealDim != 6 || c3.HermitianRealDim != 3 || c3.CKMMisalignmentCapacity || !c3.AllSectorTexturesCommute {
		t.Fatalf("bad C3 centralizer:\n%s", FormatCentralizerCase(c3))
	}
	if s3.GeneralComplexRealDim != 4 || s3.HermitianRealDim != 2 || !s3.HasOnePlusTwoDegeneracy || s3.CKMMisalignmentCapacity {
		t.Fatalf("bad S3 centralizer:\n%s", FormatCentralizerCase(s3))
	}
}

func TestNumberOperatorIsNonNativeHierarchyOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Number.NativeDerived || !a.Number.BridgeCompatible || !a.Number.SealedExternalExtension || !a.Number.BreaksExactTriality || !a.Number.ProducesDiagonalHierarchy || a.Number.ProducesMixing || a.Number.ProvidesTwoNoncommutingTextures {
		t.Fatalf("bad number audit:\n%s", FormatNumber(a.Number))
	}
}

func TestModuliFirewallPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.StartingChargedDim != 13 || a.Moduli.NativeReductionBelow13 || a.Moduli.BestNativeDim != 13 {
		t.Fatalf("bad moduli audit:\n%s", FormatModuli(a.Moduli))
	}
	native := scenarioByName(a.Moduli.Scenarios, "native ASHA after Gate 393")
	if native.ResultingDim != 13 || !native.Native || native.Conditional || native.Failed {
		t.Fatalf("native scenario should preserve 13:\n%s", FormatModuliScenario(native))
	}
}

func TestConditionalBranchesAreNotNativeSolutions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"sealed exact C3 label-triality", "sealed exact S3 label-triality", "sealed N=diag(0,1,2) hierarchy"} {
		s := scenarioByName(a.Moduli.Scenarios, name)
		if !s.Conditional || s.Native || !s.Failed || s.CKMMisalignmentPossible {
			t.Fatalf("sealed branch should not be promoted: %s\n%s", name, FormatModuliScenario(s))
		}
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusFailedDomainNotAdmitted, StatusFailedTrialityOnlyLabelSymmetry, StatusFailedExactTrialityDegeneracy, StatusFailedNoCKMMisalignment, StatusConditionalFockHierarchyCapacity, StatusFirewallPreserved13Moduli}
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

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := TrialityDomainAdmissionEquivariantYukawaCentralizerSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}
