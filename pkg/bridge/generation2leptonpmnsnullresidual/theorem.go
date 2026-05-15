package generation2leptonpmnsnullresidual

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2LeptonSectorSyntheticPMNSNullResidualTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 lepton-sector synthetic PMNS null residual"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate476 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate475 lepton preflight", Passed: a.Inheritance.Executed && a.Inheritance.Gate475LeptonPreflightDefined && a.Inheritance.Gate475RequiresENuSectors && a.Inheritance.Gate475RequiresISpecIK && a.Inheritance.Gate475RequiresBranchTags && a.Inheritance.Gate475RequiresNeutrinoPolicies && a.Inheritance.Gate475RejectsPMNSAsRayInput && a.Inheritance.Gate475RejectsNativePromotion && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines PMNS-null phase-cylinder socket", Passed: a.Map.Executed && a.Map.RequiresERow && a.Map.RequiresNuRow && a.Map.RequiresGate475Preflight && a.Map.RequiresISpecIK && a.Map.RequiresBranchTags && a.Map.RequiresNeutrinoOrdering && a.Map.RequiresAbsoluteNuScale && a.Map.RequiresSyntheticMode && a.Map.RequiresBridgeOnly && a.Map.AllowsPMNSTarget && !a.Map.AllowsPMNSAsRayInput && a.Map.ComputesDENu && !a.Map.ComputesPMNSMatrix && !a.Map.ComputesPMNSEntry && a.Map.StructurallyIdenticalToQuarkMap, Detail: FormatMap(a.Map)},
			{Name: "computes synthetic e/nu residual only", Passed: a.Sieve.ValidSyntheticResidualAccepted && a.Output.SyntheticDENu > 0 && a.Output.SyntheticPMNSTarget > 0 && a.Output.SyntheticResidual >= 0 && a.Sieve.AllAcceptedBridgeOnlySynthetic, Detail: FormatOutput(a.Output)},
			{Name: "rejects unsafe PMNS routes", Passed: a.Sieve.MissingENuRejected && a.Sieve.MissingPreflightRejected && a.Sieve.MissingRankLedgerRejected && a.Sieve.MissingNeutrinoPoliciesRejected && a.Sieve.ObservedPMNSRejected && a.Sieve.PMNSAsRayInputRejected && a.Sieve.PMNSNativePredictionRejected && a.Sieve.PMNSMatrixExportRejected && a.Sieve.NativeResidualPromotionRejected && a.Sieve.ProjectiveDomainRejected && a.Sieve.PhaseDomainRejected && a.Sieve.CausticRejected, Detail: FormatSieve(a.Sieve)},
			{Name: "preserves PMNS/native firewall", Passed: a.Firewall.Executed && a.Firewall.PMNSNullResidualAdapterDefined && a.Firewall.PMNSNullResidualMayRunBridgeOnly && a.Firewall.SyntheticLeptonDataEvaluated && !a.Firewall.ObservedLeptonDataImported && !a.Firewall.ObservedPMNSImported && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.PMNSEntryComputed && !a.Firewall.PMNSNativePrediction && !a.Firewall.DENuNativePrediction && !a.Firewall.NativeRegistryWritten && !a.Firewall.CKMNativePrediction && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate475Inherited, StatusPMNSNullMapDefined, StatusSyntheticLeptonLedgerAccepted, StatusPMNSNullResidualComputed, StatusPMNSNullResidualFirewallValidated, StatusLeptonSocketStructurallyIdentical, StatusFailedObservedPMNSImport, StatusFailedPMNSAsRayInput, StatusFailedPMNSNativePrediction, StatusFailedPMNSMatrixExport, StatusFailedNativeResidualPromotion, StatusFirewallPreserved, a.Truth}}
	}}
}
