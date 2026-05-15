package generation2nativeupdownsource

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2NativeUpDownOperatorSourceSearchTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 native up/down operator source search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate488 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate487 commutator obstruction", Passed: a.Inheritance.Executed && a.Inheritance.Gate485NullC3BaselineInherited && a.Inheritance.Gate486NullMirrorBridgeOnly && a.Inheritance.Gate487CommutatorObstruction && a.Inheritance.Gate487NullSpectrumOnly && a.Inheritance.Gate487RequiredConstraints == RequiredCKMInvariantConstraints && a.Inheritance.Gate487DerivedConstraints == 0 && a.Inheritance.NoObservedCKMImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "constructs native source ledger", Passed: a.Ledger.Executed && a.Ledger.CandidateCount >= 7 && a.Ledger.NativeUpDownLabelSources > 0 && a.Ledger.NativeQuarkLeptonSeparators > 0 && a.Ledger.NativeUniversalFamilyAxes > 0 && a.Ledger.GenerationAwareCandidates > 0 && a.Ledger.SourcesPassingAllRequirements == 0 && a.Ledger.OnlySlotsNotOperators && a.Ledger.YukawaEntriesSealed && a.Ledger.NoObservedDataImported, Detail: FormatLedger(a.Ledger)},
			{Name: "fails full CKM-source requirement sieve", Passed: a.Requirements.Executed && a.Requirements.RequiresUpDownSplit && a.Requirements.RequiresGenerationAwareness && a.Requirements.RequiresFamilyEigenbasis && a.Requirements.RequiresNativeUpOperator && a.Requirements.RequiresNativeDownOperator && a.Requirements.RequiresNativeDiagonalizers && a.Requirements.RequiresTwoInvariantConstraints && a.Requirements.CandidatesPassing == 0 && !a.Requirements.NativeUpDownOperatorsDerived && !a.Requirements.NativeDiagonalizersDerived && a.Requirements.CKMInvariantConstraintsDerived == 0 && !a.Requirements.NativeCKMSourceFound, Detail: FormatRequirements(a.Requirements)},
			{Name: "keeps O_u/O_d as sockets not populated operators", Passed: a.Socket.Executed && a.Socket.UpDownSectorLabelsNative && a.Socket.YukawaSlotsNative && !a.Socket.YukawaMatrixValuesNative && !a.Socket.FamilyEigenbasisNative && a.Socket.CanNameOuOdSlots && !a.Socket.CanPopulateOuOdNatively && !a.Socket.CanComputeUuDaggerUd && !a.Socket.CanComputeJarlskogInvariant && a.Socket.BridgeAirlockRequired, Detail: FormatSocket(a.Socket)},
			{Name: "preserves the 13-moduli firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedCKMImported && !a.Firewall.ObservedWolfensteinImported && !a.Firewall.ObservedQuarkMassesImported && !a.Firewall.ObservedYukawaEntriesImported && !a.Firewall.NativeUpOperatorWritten && !a.Firewall.NativeDownOperatorWritten && !a.Firewall.NativeDiagonalizersWritten && !a.Firewall.CKMMatrixNativePrediction && !a.Firewall.JarlskogNativePrediction && !a.Firewall.CKMInvariantConstraintNativeWrite && !a.Firewall.NativeRegistryWritten && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{StatusNativeUpDownSectorLabelsFound, StatusNativeUniversalFamilyAxisFound, StatusNoNativeUpDownEigenbasisSource, StatusNoNativeUpDownOperatorsDerived, StatusNoCKMInvariantConstraintsDerived, StatusYukawaMatricesRemainSealed, StatusFirewallBlockedNativeOperatorWrite, a.Truth}}
	}}
}
