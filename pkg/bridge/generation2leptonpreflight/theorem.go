package generation2leptonpreflight

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2LeptonSectorRankCompletePreflightTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 lepton-sector rank-complete preflight audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate475 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate474 PMNS frontier", Passed: a.Inheritance.Executed && a.Inheritance.Gate474NoNativeIK && a.Inheritance.PMNSBridgeFrontier && a.Inheritance.Gate465AirlockAvailable && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines rank-complete e/nu lepton ledger schema", Passed: a.Schema.Executed && a.Schema.RequiresISpecIK && a.Schema.RequiresBranchTags && a.Schema.RequiresEigenbasisConvention && a.Schema.RequiresNeutrinoOrderingPolicy && a.Schema.RequiresAbsoluteNeutrinoScalePolicy && a.Schema.RequiresBridgeOnly && !a.Schema.AllowsPMNSAsRayInput && !a.Schema.ComputesNow, Detail: FormatSchema(a.Schema)},
			{Name: "rejects rank-incomplete and PMNS-as-coordinate probes", Passed: a.Sieve.Executed && a.Sieve.AcceptedBridgeRows == 1 && !a.Sieve.ComputesPMNSResidual && !a.Sieve.ComputesIK && a.Sieve.Probes[0].Verdict == StatusFailedMissingENuSectors && a.Sieve.Probes[2].Verdict == StatusFailedPMNSAsCoordinate, Detail: FormatSieve(a.Sieve)},
			{Name: "accepts only complete synthetic bridge preflight", Passed: a.Sieve.Probes[3].Accepted && a.Sieve.Probes[3].ERow && a.Sieve.Probes[3].NuRow && a.Sieve.Probes[3].ISpec && a.Sieve.Probes[3].IK && a.Sieve.Probes[3].BranchTags && a.Sieve.Probes[3].BridgeOnly && !a.Sieve.Probes[3].PMNSAsRayInput, Detail: FormatProbe(a.Sieve.Probes[3])},
			{Name: "preserves PMNS/native firewall", Passed: a.Firewall.Executed && !a.Firewall.LeptonDataImported && !a.Firewall.PMNSMatrixComputed && !a.Firewall.PMNSNativePrediction && !a.Firewall.IKNativeSelectorFound && !a.Firewall.IKHalfDerived && !a.Firewall.NativeRegistryWritten && !a.Firewall.CKMNativePrediction && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate474Inherited, StatusPreflightDefined, StatusFailedPMNSAsCoordinate, StatusFailedPMNSNativePrediction, StatusFailedNativePromotion, StatusFirewallPreserved, a.Truth}}
	}}
}
