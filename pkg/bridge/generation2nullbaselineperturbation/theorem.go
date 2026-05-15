package generation2nullbaselineperturbation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2NullBaselinePerturbationLedgerSectorTransportAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 null-baseline perturbation ledger sector transport audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate481 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate480 null-vacuum baseline", Passed: a.Inheritance.Executed && a.Inheritance.Gate480NullBaseline && a.Inheritance.AlphaVac == 1 && a.Inheritance.IKVac == 0.5 && a.Inheritance.Gate480FirewallPreserved && a.Inheritance.NativeRegistryClean, Detail: "alpha_vac=1 I_K,vac=1/2 inherited as common vacuum baseline"},
			{Name: "defines bridge-only transport chart", Passed: a.Transport.Executed && a.Transport.SharedBaselineAssumed && !a.Transport.TransportPreviouslyNative && !a.Transport.PerturbationsNative && !a.Transport.IKVacCanReplaceSectorIK, Detail: a.Transport.Reason},
			{Name: "proves common baseline cancellation", Passed: a.Proof.Executed && a.Proof.BaselineAlphaCancels && a.Proof.BaselinePhiCancels && a.Proof.OnlyPerturbationsRemain, Detail: a.Proof.ReducedFormula},
			{Name: "runs synthetic perturbation diagnostics", Passed: a.Ledger.Executed && a.Ledger.AllRowsBridgeOnly && a.Ledger.AllRowsSynthetic && a.Ledger.QuarkDistanceComputed && a.Ledger.LeptonDistanceComputed && a.Ledger.Pairs[0].Distance > 0 && a.Ledger.Pairs[1].Distance > 0, Detail: FormatLedger(a)},
			{Name: "rejects physical-sector and CKM/PMNS promotion", Passed: a.Firewall.Executed && !a.Firewall.ObservedMassImported && !a.Firewall.CKMImported && !a.Firewall.PMNSImported && !a.Firewall.VacuumIKPhysicalSectorCoordinate && !a.Firewall.SectorIKSolvedByBaseline && !a.Firewall.PerturbationsNative && !a.Firewall.DUDNativePrediction && !a.Firewall.DENuNativePrediction && !a.Firewall.NativeRegistryWritten && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: a.Firewall.Reason},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusLedgerDefined, StatusBaselineCancellation, StatusSyntheticTransportOK, StatusFailedTransportNotNative, StatusFailedPerturbationsUnforced, StatusFailedIKVacAsSectorIK, StatusFailedCKMPMNSPrediction, StatusFailedNativePromotion, StatusFirewallPreserved, a.Truth}}
	}}
}

func FormatLedger(a Analysis) string {
	return "synthetic d_ud=" + fmtFloat(a.Ledger.Pairs[0].Distance) + " synthetic d_eν=" + fmtFloat(a.Ledger.Pairs[1].Distance) + " baseline_cancelled=true bridge_only=true"
}
