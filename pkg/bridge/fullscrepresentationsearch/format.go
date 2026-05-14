package fullscrepresentationsearch

import (
	"fmt"
	"sort"
	"strings"
)

func FormatInheritance(a Gate270Inheritance) string {
	return fmt.Sprintf("candidateOneForms=%t candidateOrderOne=%t fullSC=%t physicalOpposite=%t xy=%t higgs=%t firewall=%t next=%q verdict=%q", a.CandidateOneFormsExposed, a.CandidateOrderOnePasses, a.FullSCRepresentation, a.PhysicalOppositeAction, a.XYRatioSelected, a.HiggsRatioDerived, a.FirewallPreserved, a.RecommendedNextGate, a.Verdict)
}

func FormatCarrier(a FockCarrierAudit) string {
	keys := make([]int, 0, len(a.GradeHistogram))
	for k := range a.GradeHistogram {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = fmt.Sprintf("grade%d=%d", k, a.GradeHistogram[k])
	}
	return fmt.Sprintf("modes=%d dimSC=%d dimDoubled=%d masks=%d grades={%s} even=%d odd=%d create=%t annihilate=%t carResidual=%.12g carPassed=%t verdict=%q", a.ModeCount, a.BaseComplexDimension, a.DoubledComplexDimension, a.BasisMasksEnumerated, strings.Join(parts, ","), a.ParityEvenStates, a.ParityOddStates, a.CreationOperatorsAvailable, a.AnnihilationOperatorsAvailable, a.CARMaxResidual, a.CARPassed, a.Verdict)
}

func FormatLiftCandidate(c LiftCandidate) string {
	return fmt.Sprintf("%s formula=%q fullSC=%t usesCAR=%t faithfulΛ1=%t additive=%t multiplicative=%t unital=%t star=%t assocRep=%t defect=%.12g detail=%q verdict=%q", c.Name, c.Formula, c.ActsOnFullSC, c.UsesCreationAnnihilation, c.FaithfulOnOneParticle, c.LinearAdditive, c.Multiplicative, c.Unital, c.StarCompatible, c.AssociativeAlgebraRep, c.DiagnosticDefect, c.DefectDetail, c.Verdict)
}

func FormatRepresentation(a RepresentationSearchAudit) string {
	parts := make([]string, len(a.Candidates))
	for i, c := range a.Candidates {
		parts[i] = FormatLiftCandidate(c)
	}
	return fmt.Sprintf("algebra=%q target=%q validFullRep=%t blocked=%t best=%q candidates=[%s] verdict=%q", a.FiniteAlgebra, a.TargetCarrier, a.ValidFullAssociativeRepFound, a.FullSCPromotionBlocked, a.BestNativeOperatorCalculus, strings.Join(parts, "; "), a.Verdict)
}

func FormatOpposite(a OppositeActionAudit) string {
	return fmt.Sprintf("requiresLeft=%t J=%q antiLinear=%t physicalSemantics=%t oppositeConstructed=%t orderOnePhysical=%t verdict=%q", a.RequiresValidLeftRepresentation, a.CandidateJFormula, a.CandidateJAntiLinear, a.CandidateJPhysicalSemantics, a.OppositeActionConstructed, a.OrderOneCanBeEvaluatedPhysically, a.Verdict)
}

func FormatOrderOne(a OrderOneReevaluationAudit) string {
	return fmt.Sprintf("D=%q fullLeft=%t physicalOpposite=%t oneForms=%t orderOne=%t spectralTriple=%t inheritedToyResidual=%.12g verdict=%q", a.DiracFamilyFormula, a.FullSCLeftRepAvailable, a.PhysicalOppositeRepAvailable, a.NonVacuousOneFormsDerived, a.OrderOneSatisfied, a.ReevaluatedAsSpectralTriple, a.Gate270ToyResidualInherited, a.Verdict)
}

func FormatRatio(a RatioAudit) string {
	return fmt.Sprintf("xy=%t traceStable=%t gaugeProjection=%t scalarMap=%t heatKernel=%t higgs=%t verdict=%q", a.XYRatioSelected, a.TraceRatioStable, a.GaugeProjectionDerived, a.ScalarFluctuationMapDerived, a.HeatKernelNormalizationDerived, a.HiggsRatioDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("empYukawa=%t spontaneous=%t noMass=%t noVEV=%t noCutoff=%t noConnes=%t noPromotion=%t noHiggs=%t polluted=%t verdict=%q", a.EmpiricalYukawaSealPreserved, a.SpontaneousCarrierSealPreserved, a.NoObservedMassInserted, a.NoVEVInserted, a.NoCutoffScaleInserted, a.NoConnesModelImported, a.NoCandidatePromoted, a.NoHiggsPredictionClaim, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	parts := make([]string, len(a.Obligations))
	for i, o := range a.Obligations {
		parts[i] = fmt.Sprintf("%s required=%t satisfied=%t detail=%q", o.Name, o.Required, o.Satisfied, o.Detail)
	}
	return fmt.Sprintf("needRep=%t needJ=%t needOrderOneNonVacuous=%t needXY=%t needProjection=%t next=%q obligations=[%s] verdict=%q", a.NeedAssociativeFullSCRep, a.NeedPhysicalJ, a.NeedOrderOnePassingNonVacuous, a.NeedCanonicalXYSelector, a.NeedSpectralActionProjection, a.RecommendedNextGate, strings.Join(parts, "; "), a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate270=%t fullCarrier=%t car=%t lifts=%t validRep=%t J=%t orderOne=%t oneForms=%t xy=%t higgs=%t firewall=%t status=%q next=%q comment=%q", a.Gate270Inherited, a.FullCarrierEnumerated, a.CARPreflightPassed, a.NativeOperatorLiftsAudited, a.ValidFullSCRepDerived, a.PhysicalJDerived, a.FullOrderOneProved, a.NonVacuousOneFormsProved, a.XYRatioSelected, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
