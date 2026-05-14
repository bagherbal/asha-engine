package empiricalquarantineseal

import "testing"

func TestLandscapeLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Landscape.Cataloged || len(a.Landscape.Items) < 8 || !a.Landscape.ContainsWeakMixing || !a.Landscape.ContainsMoritaSplit || !a.Landscape.ContainsGenerationTriality || !a.Landscape.ContainsTrueBimodule || !a.Landscape.ContainsTraceEquivalence || !a.Landscape.ContainsThresholdJump || !a.Landscape.ContainsPfaffianHierarchy || !a.Landscape.ContainsAlphaEightPi {
		t.Fatalf("bad landscape: %s", FormatLandscape(a.Landscape))
	}
}

func TestProxyLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Proxies.Cataloged || !a.Proxies.Contains125TreeProxy || !a.Proxies.ContainsNative125Proxy || !a.Proxies.ContainsThresholdTransport || !a.Proxies.ContainsPrecisionPoleTarget || !a.Proxies.AllEmpiricalInputsQuarantined || a.Proxies.FinalMassClaimed {
		t.Fatalf("bad proxies: %s", FormatProxies(a.Proxies))
	}
}

func TestQuarantineLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Quarantine.Defined || a.Quarantine.MinimalSMVacuumDimension != 15 || a.Quarantine.ExtendedVacuumDimension != 25 || !a.Quarantine.ContainsYukawas || !a.Quarantine.ContainsCKM || !a.Quarantine.ContainsStrongCP || !a.Quarantine.ContainsGravityCutoff || !a.Quarantine.ContainsPoleScheme || !a.Quarantine.ContainsCosmologicalConstant || !a.Quarantine.ContainsFlavorProjectionMetric || a.Quarantine.AnyClosed {
		t.Fatalf("bad quarantine: %s", FormatQuarantine(a.Quarantine))
	}
}

func TestSeparationAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Audit.NoYukawaFitPromoted || !a.Audit.NoCKMInvented || !a.Audit.NoPoleSchemeChosen || !a.Audit.NoCosmologicalFitPromoted || !a.Audit.NoObservedMassInserted || !a.Audit.NoAlphaGUTFitNeededInFinal || !a.Audit.NoFinalTOEClaimed || !a.Audit.NoExactColliderClaimed || !a.Audit.LandscapeVacuumSeparated || a.Audit.FiniteCorePolluted {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusGrandUnifiedLedgerCompiled, StatusPhaseIIQuarantineSealed, StatusRigidLandscapeCataloged, StatusProxiesCataloged, StatusEmpiricalQuarantineDefined, StatusSeparationPreserved, StatusFailedVacuumNotDerived, StatusFailedYukawasQuarantined, StatusFailedCKMQuarantined, StatusFailedGravityCutoffQuarantined, StatusFailedPoleSchemeQuarantined, StatusFailedCosmologicalQuarantined, StatusFailedFinalTOE}
	for _, req := range required {
		found := false
		for _, s := range statuses {
			if s == req {
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
	res := EmpiricalQuarantineSealGrandUnifiedProjectLedgerTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
