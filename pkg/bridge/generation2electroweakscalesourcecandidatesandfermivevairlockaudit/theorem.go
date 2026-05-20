package generation2electroweakscalesourcecandidatesandfermivevairlockaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ElectroweakScaleSourceCandidatesAndFermiVEVAirlockAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 778 — Electroweak Scale Source Candidates and Fermi-VEV Airlock Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate778ElectroweakScaleSourceBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 777 VEV scale airlock", Passed: a.Gate777.Inherited && a.Gate777.TreeTowerFormula == "m_H_tree=(v/2)sqrt(C_Higgs)" && closeRel(a.Gate777.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Gate777.DilationFactor, 1.0184402389953279, 1e-15) && closeRel(a.Gate777.VEVScaleGeV, 246.2196508, 1e-15) && closeRel(a.Gate777.TreeMassGeV, 125.38000000304908, 1e-15) && a.Gate777.DimensionlessCorrectionExists && a.Gate777.DimensionfulScaleSeal == "v" && !a.Gate777.NativeVEVTheorem && !a.Gate777.PoleMassTheorem, Detail: FormatGate777(a.Gate777)},
			{Name: "audit Fermi-VEV convention lane", Passed: a.FermiLane.Audited && a.FermiLane.SealName == "FermiVEVScaleSeal" && a.FermiLane.Formula == "v=(sqrt(2)G_F)^(-1/2)" && a.FermiLane.Input == "G_F" && a.FermiLane.Output == "v" && closeRel(a.FermiLane.VEVGeV, 246.2196508, 1e-15) && closeRel(a.FermiLane.EquivalentGFGeVMinus2, 1.1663786999444556e-05, 1e-15) && !a.FermiLane.NativeFermiTheorem && !a.FermiLane.NativeVEVTheorem && a.FermiLane.LawfulExternalAirlock, Detail: FormatFermiLane(a.FermiLane)},
			{Name: "audit W-mass / gauge-coupling lane", Passed: a.WLane.Audited && a.WLane.Formula == "v=2m_W/g" && a.WLane.RequiresAbsoluteWeakG && a.WLane.RequiresWMass && a.WLane.GaugeRatiosOrganized && !a.WLane.AbsoluteWeakScaleNative && !a.WLane.WMassNative && a.WLane.LaneSealed, Detail: FormatWLane(a.WLane)},
			{Name: "audit potential stationarity lane", Passed: a.PotentialLane.Audited && a.PotentialLane.Formula == "v^2=-mu^2/lambda_H" && a.PotentialLane.LambdaAirlockExists && !a.PotentialLane.MuSquaredIndependentlySourced && a.PotentialLane.CircularWithoutMuSource && !a.PotentialLane.NativeMuSquaredSource && !a.PotentialLane.DeterminesVEV, Detail: FormatPotentialLane(a.PotentialLane)},
			{Name: "audit spectral-action scale candidate", Passed: a.SpectralLane.Audited && a.SpectralLane.Candidate == "dimensionful spectral-action scale or cutoff" && a.SpectralLane.DimensionfulScaleCouldSetUnits && a.SpectralLane.CurrentBridgeDimensionlessOnly && !a.SpectralLane.MapsSpectralScaleToVEV && a.SpectralLane.LaneCandidateOnly, Detail: FormatSpectralLane(a.SpectralLane)},
			{Name: "audit boundary/RG scale lane", Passed: a.BoundaryLane.Audited && a.BoundaryLane.BoundaryScaleSealExists && a.BoundaryLane.ScalarWallDataExists && !a.BoundaryLane.DeterminesElectroweakVEV && !a.BoundaryLane.BoundaryScaleEqualsVEVTheorem, Detail: FormatBoundaryLane(a.BoundaryLane)},
			{Name: "record source ranking and blocked shortcuts", Passed: a.Ranking.Recorded && a.Ranking.BestCurrentLawfulSource == "FermiVEVScaleSeal: v=(sqrt(2)G_F)^(-1/2)" && containsAll(a.Ranking.BestFutureNativeTargets, []string{"mu^2 source theorem", "absolute electroweak scale theorem"}) && containsAll(a.Ranking.BlockedShortcuts, []string{"C_Higgs does not determine v", "lambda_runtime_eff does not determine v without mu^2", "P_rad does not determine v", "HistoryLoopUnit does not determine v", "7/72 does not determine v", "1/(8pi) does not determine v"}), Detail: FormatRanking(a.Ranking)},
			{Name: "record derived scale ledger", Passed: a.Ledger.Finite && closeRel(a.Ledger.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Ledger.DilationFactor, 1.0184402389953279, 1e-15) && closeRel(a.Ledger.VEVGeV, 246.2196508, 1e-15) && closeRel(a.Ledger.EquivalentGF, 1.1663786999444556e-05, 1e-15) && closeRel(a.Ledger.VHalfGeV, 123.1098254, 1e-15) && closeRel(a.Ledger.TreeMassGeV, 125.38000000304908, 1e-15), Detail: FormatLedger(a.Ledger)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.FermiScaleNativeTheorem && !a.Firewalls.VEVDerivedFromCHiggs && !a.Firewalls.VEVDerivedFromLambdaRuntimeOnly && !a.Firewalls.MuSquaredBridgeNativeSource && !a.Firewalls.WRelationNativeWithoutInputs && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.DimensionlessTowerMassScaleTheorem && !a.Firewalls.YukawaOperatorOrEigenvalue && a.Firewalls.Verdict == StatusGate778ElectroweakScaleSourceBoundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if want == 0 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}

func containsAll(haystack, needles []string) bool {
	joined := "\x00" + strings.Join(haystack, "\x00") + "\x00"
	for _, n := range needles {
		if !strings.Contains(joined, "\x00"+n+"\x00") {
			return false
		}
	}
	return true
}
