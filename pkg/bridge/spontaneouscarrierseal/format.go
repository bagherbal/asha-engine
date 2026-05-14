package spontaneouscarrierseal

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate255Audit) string {
	return fmt.Sprintf("gate254=%t SC=%t local=%t common=%t intertwiner=%t ledger=%t so8=%t triality=%t Q=%t plane=%t vtau=%t yukawa=%t status=%q", a.Gate254Inherited, a.SCCarrierKnown, a.LocalActionsAudited, a.CommonCarrierDerived, a.CarrierIntertwinerDerived, a.UnifiedLedgerConstructed, a.T3LYPhiSO8Coordinates, a.TrialityBranchSelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.VTauConstructed, a.YukawaTextureDerived, a.Status)
}

func FormatNativeLedger(x NativeSearchLedger) string {
	return fmt.Sprintf("%s carrier=%q native=%q commonSC=%t ledger=%t so8=%t obstruction=%q sealAllowed=%t changesTheorem=%t verdict=%q", x.Object, x.Carrier, x.NativeStatus, x.DerivedCommonSC, x.DerivedNumberLedger, x.DerivedSO8Coordinates, x.Obstruction, x.SealAllowedAsBoundaryData, x.BoundaryDataChangesTheorem, x.Verdict)
}

func FormatNativeSearch(a NativeSearchAudit) string {
	parts := make([]string, len(a.Ledgers))
	for i, x := range a.Ledgers {
		parts[i] = FormatNativeLedger(x)
	}
	return fmt.Sprintf("gate255=%t nativeIntertwiner=%t nativeLedger=%t nativeSO8=%t nativeTriality=%t verdict=%q ledgers=[%s]", a.Gate255NoGoInherited, a.NativeCommonIntertwinerExists, a.NativeUnifiedLedgerExists, a.NativePhysicalSO8CoordinatesExist, a.NativeTrialityPullbackAvailable, a.Verdict, strings.Join(parts, "; "))
}

func FormatSeal(a SpontaneousCarrierSeal) string {
	return fmt.Sprintf("%s id=%q status=%q target=%q explicit=%t quarantined=%t required=%t gauge=%t vacuum=%t weakFrame=%t leftInjection=%t scalarEmbedding=%t triality=%t finiteDerived=%t masses=%t yukawas=%t couplings=%t overridesNoGo=%t polluted=%t verdict=%q", a.Name, a.AxiomID, a.ConditionalStatus, a.TargetCarrier, a.ExplicitAxiom, a.Quarantined, a.RequiredByGate255, a.GaugeFixingRequired, a.VacuumOrientationRequired, a.WeakFrameRequired, a.LeftDoubletInjectionRequired, a.ScalarEmbeddingRequired, a.TrialityBranchRequired, a.DerivedFromFiniteGeometry, a.UsesObservedMasses, a.UsesObservedYukawas, a.UsesObservedGaugeCouplings, a.OverridesNativeNoGo, a.PollutesFiniteCore, a.Verdict)
}

func FormatAxiomDatum(x AxiomDatum) string {
	return fmt.Sprintf("%s symbol=%q required=%t provided=%t derived=%t quarantined=%t neededFor=%q missing=%q provider=%q", x.Name, x.Symbol, x.Required, x.Provided, x.Derived, x.Quarantined, x.NeededFor, x.MissingReason, x.AdmissibleProvider)
}

func FormatIntertwiner(a ConditionalIntertwinerAxiom) string {
	parts := make([]string, len(a.AxiomData))
	for i, x := range a.AxiomData {
		parts[i] = FormatAxiomDatum(x)
	}
	return fmt.Sprintf("%s sources=%v target=%q schema=%t operational=%t required=%d provided=%d derived=%d all=%t mapsT3=%t mapsY=%t intertwines=%t ledgers=%t tensor=%t directSum=%t verdict=%q data=[%s]", a.Name, a.SourceCarriers, a.TargetCarrier, a.SchemaDefined, a.OperationalIntertwinerBuilt, a.RequiredDataCount, a.ProvidedDataCount, a.DerivedDataCount, a.AllRequiredDataProvided, a.MapsT3LIntoSC, a.MapsYPhiIntoSC, a.IntertwinesLocalActions, a.ProducesFourModeNumberLedgers, a.ChangesCarrierByTensorProduct, a.UsesDirectSumAsIntertwiner, a.Verdict, strings.Join(parts, "; "))
}

func FormatSealedLedger(x SealedFockLedger) string {
	return fmt.Sprintf("%s expr=%q coeffs=%v basis=%v symbolic=%t numericSet=%t numeric=%v constraints=%v seal=%t meaning=%q verdict=%q", x.Name, x.Expression, x.CoefficientSymbols, x.NumberOperatorBasis, x.SymbolicLedgerDefined, x.NumericCoefficientsSet, x.NumericCoefficients, x.NormalizationConstraints, x.SealRequired, x.PhysicalMeaning, x.Verdict)
}

func FormatUnifiedLedger(a UnifiedLedgerAudit) string {
	parts := make([]string, len(a.Ledgers))
	for i, x := range a.Ledgers {
		parts[i] = FormatSealedLedger(x)
	}
	return fmt.Sprintf("carrier=%q symbolic=%t T3concrete=%t Yconcrete=%t Qconcrete=%t operational=%t obstruction=%q verdict=%q ledgers=[%s]", a.Carrier, a.SymbolicSchemaConstructed, a.ConcreteT3LNumberLedgerAvailable, a.ConcreteYPhiNumberLedgerAvailable, a.ConcreteQNumberLedgerAvailable, a.OperationalUnifiedLedgerBuilt, a.Obstruction, a.Verdict, strings.Join(parts, "; "))
}

func FormatSO8Coordinate(x SymbolicSO8Coordinate) string {
	return fmt.Sprintf("%s fock=%q so8=%q symbols=%v bivectors=%v symbolic=%t concrete=%t seal=%t verdict=%q", x.Name, x.FockExpression, x.SO8Formula, x.CoordinateSymbols, x.CartanBivectors, x.SymbolicFormulaAvailable, x.ConcreteCoordinateAvailable, x.SealRequired, x.Verdict)
}

func FormatSO8(a SO8TranslationAudit) string {
	parts := make([]string, len(a.Coordinates))
	for i, x := range a.Coordinates {
		parts[i] = FormatSO8Coordinate(x)
	}
	return fmt.Sprintf("witt=%t symbolic=%t T3=%t Y=%t Q=%t obstruction=%q verdict=%q coords=[%s]", a.WittDictionaryInherited, a.SymbolicSchemaAvailable, a.ConcreteT3LSO8, a.ConcreteYPhiSO8, a.ConcreteQSO8, a.Obstruction, a.Verdict, strings.Join(parts, "; "))
}

func FormatTrialityKernel(a TrialityKernelAudit) string {
	return fmt.Sprintf("candidates=%t branchSchema=%t branch=%t selected=%q outcome=%t Q=%t eig=%t known=%t dim=%d exact3=%t plane=%t equations=%t missing=%v verdict=%q", a.TrialityCandidatesKnown, a.BranchSelectionSchemaDefined, a.PhysicalBranchSelected, a.SelectedBranch, a.SelectedByOutcome, a.Q8vCConstructed, a.EigensystemComputed, a.KernelDimensionKnown, a.KernelComplexDimension, a.ExactlyThree, a.NeutralThreePlaneDerived, a.DiagnosticEquationsRecorded, a.MissingConcreteInputs, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("seal=%t quarantine=%t nativeNoGo=%t invented=%t SMY=%t weakPlane=%t trialityByKernel=%t kernel3=%t sealAsFinite=%t tensor=%t directSum=%t yukawa=%t masses=%t polluted=%t verdict=%q", a.SealExplicitInput, a.SealQuarantined, a.NativeNoGoPreserved, a.InventedEmbeddingValues, a.ImportedSMHyperchargeConvention, a.ForcedWeakPlane, a.SelectedTrialityByKernel, a.ForcedKernelDim3, a.TreatedSealAsFiniteDerivation, a.TreatedTensorProductAsSC, a.TreatedDirectSumAsIntertwiner, a.InsertedYukawaTexture, a.ImportedObservedMasses, a.PollutedFiniteCore, a.Verdict)
}

func FormatDownstream(a DownstreamAudit) string {
	return fmt.Sprintf("plane=%t tau=%v vtau=%t trialityTexture=%t yukawa=%t CKM=%t masses=%t verdict=%q", a.Neutral3PlaneAvailable, a.TauEta, a.VTauConstructed, a.TrialityTextureOpened, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate255=%t seal=%t schema=%t values=%t symbolicLedger=%t concreteLedger=%t symbolicSO8=%t concreteSO8=%t triality=%t Q=%t plane=%t yukawa=%t status=%q next=%q comment=%q", a.Gate255NoGoInherited, a.SpontaneousSealRecorded, a.ConditionalIntertwinerSchema, a.SealedAxiomValuesProvided, a.SymbolicLedgerSchemaAvailable, a.ConcreteUnifiedLedgerBuilt, a.SymbolicSO8SchemaAvailable, a.ConcreteSO8Coordinates, a.TrialityBranchSelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
