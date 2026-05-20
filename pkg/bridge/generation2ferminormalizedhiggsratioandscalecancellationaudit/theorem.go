package generation2ferminormalizedhiggsratioandscalecancellationaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FermiNormalizedHiggsRatioAndScaleCancellationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 779 — Fermi-Normalized Higgs Ratio and Scale-Cancellation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate779FermiNormalizedHiggsRatioBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 778 electroweak scale airlock", Passed: a.Gate778.Inherited && a.Gate778.TreeTowerFormula == "m_H_tree=(v/2)sqrt(C_Higgs)" && a.Gate778.FermiVEVSeal == "FermiVEVScaleSeal: v=(sqrt(2)G_F)^(-1/2)" && closeRel(a.Gate778.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Gate778.VEVGeV, 246.2196508, 1e-15) && closeRel(a.Gate778.DilationFactor, 1.0184402389953279, 1e-15) && closeRel(a.Gate778.TreeMassGeV, 125.38000000304908, 1e-15) && closeRel(a.Gate778.EquivalentGFGeVMinus2, 1.1663786999444556e-05, 1e-15) && !a.Gate778.NativeFermiTheorem && !a.Gate778.NativeElectroweakScale && !a.Gate778.PoleMassTheorem, Detail: FormatGate778(a.Gate778)},
			{Name: "define Fermi-normalized ratio", Passed: a.Ratio.Defined && a.Ratio.SquaredTreeFormula == "m_H_tree^2=(v^2/4)C_Higgs" && a.Ratio.VEVCancelledFormula == "m_H_tree^2/v^2=C_Higgs/4" && a.Ratio.FermiConvention == "1/v^2=sqrt(2)G_F" && a.Ratio.NormalizedIdentity == "4sqrt(2)G_F m_H_tree^2=C_Higgs" && a.Ratio.DimensionlessLeft == "4sqrt(2)G_F m_H_tree^2" && a.Ratio.DimensionlessRight == "C_Higgs" && a.Ratio.UsesExternalGFSeal && !a.Ratio.DerivesGF && !a.Ratio.DerivesVEV, Detail: FormatRatio(a.Ratio)},
			{Name: "compute VEV scale cancellation", Passed: a.Cancellation.Computed && closeRel(a.Cancellation.TreeMassSquaredOverVSquared, 0.2593051301012151, 1e-15) && closeRel(a.Cancellation.CHiggsOverFour, 0.2593051301012151, 1e-15) && closeRel(a.Cancellation.Sqrt2GFTreeMassSquared, 0.2593051301012151, 1e-15) && closeRel(a.Cancellation.FourSqrt2GFTreeMassSquared, 1.0372205204048603, 1e-15) && a.Cancellation.MatchesCHiggs && a.Cancellation.ScaleCancelledToDimensionless, Detail: FormatCancellation(a.Cancellation)},
			{Name: "record numerical ratio ledger", Passed: a.Ledger.Finite && closeRel(a.Ledger.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Ledger.VEVGeV, 246.2196508, 1e-15) && closeRel(a.Ledger.EquivalentGFGeVMinus2, 1.1663786999444556e-05, 1e-15) && closeRel(a.Ledger.TreeMassGeV, 125.38000000304908, 1e-15) && closeRel(a.Ledger.TreeMassSquaredGeV2, 15720.144400764586, 1e-15) && closeRel(a.Ledger.TreeMassOverVEV, 0.5092201194976639, 1e-15) && closeRel(a.Ledger.Sqrt2GFTreeMassSquared, 0.2593051301012151, 1e-15) && closeRel(a.Ledger.FourSqrt2GFTreeMassSquared, 1.0372205204048603, 1e-15), Detail: FormatLedger(a.Ledger)},
			{Name: "separate dimensionless and scale tasks", Passed: a.Tasks.Separated && a.Tasks.DimensionlessTask == "derive or reduce C_Higgs natively" && a.Tasks.ScaleTask == "derive or seal G_F / v" && a.Tasks.RequiresBothForMass && a.Tasks.RatioDoesNotDeriveGF && a.Tasks.RatioDoesNotDeriveVEV, Detail: FormatTasks(a.Tasks)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.RatioPoleMassTheorem && !a.Firewalls.GFAShaNativeInput && !a.Firewalls.FermiNormalizedRatioMeasuredPrediction && !a.Firewalls.CHiggsNativeHiggsTheorem && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.DimensionlessRatioElectroweakScale && !a.Firewalls.YukawaOperatorOrEigenvalue && a.Firewalls.Verdict == StatusGate779FermiNormalizedHiggsRatioBoundary, Detail: FormatFirewalls(a.Firewalls)},
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

func containsStatuses(notes []string, wants []string) bool {
	joined := "\x00" + strings.Join(notes, "\x00") + "\x00"
	for _, w := range wants {
		if !strings.Contains(joined, "\x00"+w+"\x00") {
			return false
		}
	}
	return true
}
