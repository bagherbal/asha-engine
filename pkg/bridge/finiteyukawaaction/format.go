package finiteyukawaaction

import "fmt"

func FormatInheritance(a Gate262Inheritance) string {
	return fmt.Sprintf("bilinear=%t tauSource=%t offdiag=%t hermitianBasis=%t rawPartner=%t prevQualified=%t prevTexture=%t tau=%v real=%q phase=%q commNorms=(%d,%d) verdict=%q", a.BilinearCarrierDefined, a.TauEtaDiagonalSourceOpened, a.OffDiagonalComplementExposed, a.HermitianTrialityBasisExposed, a.RawNonCommutingPartnerExists, a.PreviousQualifiedPartnerFound, a.PreviousPhysicalYukawaTexture, a.TauEtaEigenvalues, a.RealBasisName, a.PhaseBasisName, a.RealCommutatorNormSquared, a.PhaseCommutatorNormSquared, a.Verdict)
}

func FormatTraceAudit(a TraceFunctionalAudit) string {
	return fmt.Sprintf("name=%q formula=%q applies=%t exact=%t real=%d phase=%d cross=%d tauReal=%d tauPhase=%d nonzero=%t distinguishes=%t selectsCoeff=%t promotable=%t verdict=%q", a.FunctionalName, a.Formula, a.AppliesToM3, a.ExactEvaluation, a.RealBasisValue, a.PhaseBasisValue, a.CrossValue, a.TauRealCrossValue, a.TauPhaseCrossValue, a.NonZeroOnBasis, a.DistinguishesRealAndPhase, a.SelectsAmplitudeCoefficient, a.PromotableToYukawaAction, a.Verdict)
}

func FormatActionCandidate(a NativeActionCandidate) string {
	return fmt.Sprintf("name=%q source=%q available=%t canonical=%t actsOnM3=%t finiteTrace=%t evalsTriality=%t coeffs=%t relAmp=%t bGapMap=%t hopfMap=%t missing=%q texture=%t verdict=%q", a.Name, a.Source, a.Available, a.Canonical, a.ActsOnM3BilinearCarrier, a.HasFiniteTraceOrVariation, a.EvaluatesTrialityBasis, a.AssignsNonzeroCoefficients, a.SelectsRelativeAmplitude, a.MapsBGapToOffDiagonal, a.MapsHopfPhaseToPhaseBasis, a.RequiresMissingIngredient, a.PhysicalYukawaTextureDerived, a.Verdict)
}

func FormatScalarPhase(a ScalarPhaseIntegrationAudit) string {
	return fmt.Sprintf("bGapScale=%t bGapCoeff=%t bGapEndo=%t bGapWeights=%t hopfLedger=%t hopfProjection=%t hopfCP=%t integrated=%t verdict=%q", a.BGapAvailableAsScale, a.BGapActionCoefficientDerived, a.BGapGenerationEndomorphismDerived, a.BGapCanWeightTrialityBasis, a.HopfPhaseLedgerAvailable, a.HopfProjectionToKTrialityDerived, a.HopfCanFixCPPhase, a.ScalarPhaseIntegrationDerived, a.Verdict)
}

func FormatTexture(a TextureConstructionAudit) string {
	return fmt.Sprintf("tau=%t offdiagBasis=%t traceMetric=%t coeffRule=%t relPhase=%t scale=%t sector=%t formula=%q free=%v texture=%t seal=%t verdict=%q", a.DiagonalTauSourceAvailable, a.HermitianOffDiagonalBasisExists, a.TraceMetricAvailable, a.FiniteActionCoefficientRule, a.RelativeRealPhaseWeightSelected, a.OverallYukawaScaleSelected, a.FermionKindDependenceSelected, a.CandidateFormula, a.FreeParameters, a.PhysicalTextureConstructed, a.EmpiricalYukawaSealRequired, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate262=%t noTraceDyn=%t noSymAmp=%t noMasses=%t noAngles=%t noBGap=%t noHopf=%t noSpectralTriple=%t empirical=%t polluted=%t verdict=%q", a.Gate262RawBasisPreserved, a.DoesNotPromoteTraceMetricToDynamics, a.DoesNotPromoteSymmetryToAmplitude, a.DoesNotUseObservedMasses, a.DoesNotUseObservedMixingAngles, a.DoesNotUseBGapWithoutMap, a.DoesNotUseHopfWithoutProjection, a.DoesNotClaimSpectralTripleComplete, a.EmpiricalYukawaSealPreserved, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate262=%t traceEval=%t traceDegenerate=%t actionCandidates=%d qualified=%t scalarPhase=%t finiteAction=%t texture=%t ckm=%t masses=%t status=%q next=%q comment=%q", a.Gate262Inherited, a.TraceFunctionalsEvaluated, a.TraceMetricDegenerate, a.NativeActionCandidateCount, a.ActionCandidateQualified, a.ScalarPhaseIntegrationDerived, a.FiniteYukawaActionDerived, a.PhysicalYukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
