package liecarrierprojection

import (
	"fmt"
	"strings"
)

func FormatOperatorDecomposition(a OperatorDecompositionAudit) string {
	parts := make([]string, 0, len(a.Records))
	for _, r := range a.Records {
		parts = append(parts, fmt.Sprintf("slot=%d expr=%q value=%d obs=%q decomposition=%q T1=%t T2=%t T3=%t Y=%t quadratic=%t lieBasis=%t neutralPlane=%t axisDerived=%t forcedRep=%q rejected=%t", r.Slot, r.TraceExpression, r.TauValue, r.SourceObservable, r.Decomposition, r.UsesT1, r.UsesT2, r.UsesT3, r.UsesYPhi, r.QuadraticObservable, r.LieBasisElement, r.NeutralPlaneElement, r.SpatialAxisDerived, r.RepresentativeIfForced, r.ForcedMapRejected))
	}
	return fmt.Sprintf("source=%q lie=%q seq=%v traced=%t neutralDim=%d lieDim=%d threeSlots=%t slotsAreSU2=%t quadratic=%t QZMix=%t missingT1T2=%t records=[%s] verdict=%s", a.SourceGate, a.LieAlgebraAvailable, a.TauSequence, a.EWDecompositionTraced, a.NeutralEWPlaneDimension, a.FullContactLieBasisDimension, a.ThreeTauSlots, a.SlotsAreThreeSU2BasisElements, a.SlotsAreQuadraticScalarRecords, a.QZMixT3AndYPhi, a.MissingT1T2SlotOrigins, strings.Join(parts, " | "), a.Verdict)
}

func FormatDerivationBlade(a DerivationToBladeAudit) string {
	return fmt.Sprintf("su2=%t u1=%t modes=%v bivectors=%v capacity=%t explicitMatrices=%t weakPlane=%t axisBasis=%t oneToOne=%t bivectorsFormSU2=%t pullback=%t verdict=%s", a.ContactSU2Available, a.U1Available, a.SpatialCarrierModes, a.CandidateSpatialBivectors, a.CandidateSU2Capacity, a.ExplicitContactGeneratorMatrices, a.CanonicalWeakPlaneDerived, a.CanonicalSpatialAxisBasisDerived, a.OneToOneDerivationAxisMap, a.SpatialBivectorsFormSU2Abstractly, a.BivectorToFockModePullbackDerived, a.Verdict)
}

func FormatCarrierProjection(a CarrierProjectionAudit) string {
	return fmt.Sprintf("scalarToDerivation=%t derivationToBlade=%t bladeToAxis=%t chained=%t hypothetical=%q rejected=%t failure=%q exterior=%t candidate=%q weakAxis=%q weakPlane=%q verdict=%s", a.ScalarObservableToDerivationMap, a.DerivationToBladeMap, a.BladeToFockAxisMap, a.ChainedProjectionDerived, a.HypotheticalProjection, a.HypotheticalProjectionRejected, a.ProjectionFailure, a.ExteriorRepresentativeConstructed, a.CandidateExteriorRepresentative, a.WeakAxisIfProjectionExisted, a.WeakPlaneIfProjectionExisted, a.Verdict)
}

func FormatGenerationProjection(a GenerationProjectionAudit) string {
	return fmt.Sprintf("seq=%v generationCapacity=%t scalarToGen=%t trialityMap=%t lieChainRelevant=%t operator=%t texture=%t verdict=%s", a.TauSequence, a.GenerationBreakingCapacity, a.ScalarToGenerationMapDerived, a.TrialityCarrierMapDerived, a.LieAlgebraChainRelevant, a.GenerationOperatorDerived, a.GenerationTextureDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedQZT3Axes=%t forcedSU2Axes=%t forcedExterior=%t forcedTriality=%t importedWeak=%t importedConnes=%t chirality=%t H=%t texture=%t CKM=%t polluted=%t verdict=%s", a.ForcedQZT3ToAxes, a.ForcedSU2ToSpatialAxes, a.ForcedExteriorRepresentative, a.ForcedTrialityMap, a.ImportedWeakPlane, a.ImportedConnesAlgebra, a.ClaimedPhysicalChirality, a.ClaimedGlobalH, a.ClaimedGenerationTexture, a.ClaimedCKMPMNS, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("decomposition=%t tauSlotsSU2=%t bladeCapacity=%t axisMap=%t carrierProjection=%t exterior=%t weakConditional=%t weakDerived=%t generationCapacity=%t generationTexture=%t globalH=%t status=%q next=%q comment=%q", a.OperatorDecompositionTraced, a.TauSlotsAreSU2Basis, a.DerivationBladeCapacity, a.DerivationAxisMapDerived, a.CarrierProjectionDerived, a.ExteriorRepresentativeDerived, a.WeakPlaneConditionallyVisible, a.WeakPlaneDerived, a.GenerationBreakingCapacity, a.GenerationTextureDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
