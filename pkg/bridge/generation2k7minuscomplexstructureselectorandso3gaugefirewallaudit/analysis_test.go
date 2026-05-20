package generation2k7minuscomplexstructureselectorandso3gaugefirewallaudit

import (
	"strings"
	"testing"
)

func TestGate712FamilyMapSO3CovarianceAndSelectors(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.U2SocketInherited || a.Inherited.K7MinusDimension != 3 || !a.Inherited.CanSelectJH || a.Inherited.CanonicalJHSelected || a.Inherited.PhysicalElectroweakU2 || a.Inherited.PhysicalHiggsDoubletMap || a.Inherited.HyperchargeAssignment {
		t.Fatalf("bad Gate711 inheritance: %+v", a.Inherited)
	}
	if a.FamilyMap.UnitSphereDimension != 2 || a.FamilyMap.K7MinusDimension != 3 || !a.FamilyMap.UnitDirectionWouldSelect || a.FamilyMap.NativeSelectorFound || !a.FamilyMap.FamilyValuedU2Socket || a.FamilyMap.CanonicalComplexStructure || !strings.Contains(a.FamilyMap.Verdict, StatusU2SocketFamilyValuedOverS2) {
		t.Fatalf("bad family map audit: %+v", a.FamilyMap)
	}
	if !a.SO3Covariance.PreservesOmega || !a.SO3Covariance.PreservesEta123 || !a.SO3Covariance.GaugeCovariantFrame || a.SO3Covariance.CanonicalOrderedFrame || a.SO3Covariance.SelectsSingleAxis || a.SO3Covariance.PhysicalGenerationAxis || !strings.Contains(a.SO3Covariance.Verdict, StatusFanoVolumeOrFrameDoesNotSelectSingleAxis) {
		t.Fatalf("bad SO3 covariance audit: %+v", a.SO3Covariance)
	}
	if len(a.Selectors.Candidates) != 7 || a.Selectors.NativeSelectorFound || !a.Selectors.BoundaryScalarsRejected || !a.Selectors.HistoryScalarsRejected || !a.Selectors.ExternalSealOnly || !strings.Contains(a.Selectors.Verdict, StatusNoNativeK7MinusUnitVectorSelector) {
		t.Fatalf("bad selector candidate audit: %+v", a.Selectors)
	}
}

func TestGate712FirewallsVacuumFlavorAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	f := a.GaugeFirewall
	if !f.FamilyValuedPreHiggsCarriers || f.UniqueHiggsComplexStructure || f.PhysicalHiggsDoubletMap || f.PhysicalGenerationSpaceMap || f.YukawaOperatorMap || f.FlavorHierarchyTheorem || f.CKMPMNSThereom || !strings.Contains(f.Verdict, StatusGate712K7MinusSelectorBoundary) {
		t.Fatalf("gauge firewall violated: %+v", f)
	}
	if !strings.Contains(a.VacuumSelector.MissingObject, "n_*") || a.VacuumSelector.NativeTheoremCertified || !a.VacuumSelector.EnvironmentalSealAllowed || !a.VacuumSelector.SealMustBeQuarantined {
		t.Fatalf("bad vacuum selector classification: %+v", a.VacuumSelector)
	}
	if !a.FlavorRelation.ResemblesFlavorAxis || a.FlavorRelation.GenerationCarrierMap || a.FlavorRelation.FlavorOrientationMap || a.FlavorRelation.YukawaEigenvalueMap || a.FlavorRelation.CKMPMNSMap {
		t.Fatalf("bad flavor relation firewall: %+v", a.FlavorRelation)
	}
	if len(a.Missing.Missing) != 5 || !strings.Contains(a.Missing.Verdict, StatusNoTypedK7MinusToPhysicalGenerationSpaceMap) || !strings.Contains(a.Missing.Verdict, StatusNoYukawaOperatorOrEigenvalueTheorem) {
		t.Fatalf("bad missing map ledger: %+v", a.Missing)
	}
	res := Generation2K7MinusComplexStructureSelectorAndSO3GaugeFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
