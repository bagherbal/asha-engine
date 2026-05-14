package faithfuloppositeactionrep

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate269Inheritance) string {
	return fmt.Sprintf("orderOne=%t sieve=%t family=%q dimC=%d nonVacuous=%t canonical=%t higgs=%t firewall=%t next=%q verdict=%q", a.OrderOneDefined, a.ModeLevelSieveReduced, a.AllowedFamilyFormula, a.SurvivingFamilyDimC, a.NonVacuousCalculus, a.CanonicalDFDerived, a.HiggsRatioDerived, a.FirewallPreserved, a.RecommendedNextGate, a.Verdict)
}

func FormatLift(a FaithfulLiftAudit) string {
	return fmt.Sprintf("target=%q dim=%d available=%q availDim=%d candidate=%q candDim=%d fullSC=%t grading=%t candidateFaithful=%t importedConnes=%t verdict=%q", a.TargetCarrier, a.TargetComplexDimension, a.AvailableCarrier, a.AvailableComplexDimension, a.ModePreflightCarrier, a.ModePreflightComplexDimension, a.FullSCRepresentationDerived, a.ChiralGradingRespected, a.AlgebraFaithfulOnCandidate, a.ImportedConnesRepresentation, a.Verdict)
}

func FormatOpposite(a OppositeActionAudit) string {
	return fmt.Sprintf("J=%q antiLinear=%t particleSemantics=%t opposite=%t candidateSwap=%t physical=%t verdict=%q", a.JFormula, a.AntiLinearJDerived, a.ParticleAntiparticleSemantics, a.OppositeActionDerived, a.CandidateSwapActionAvailable, a.CandidateIsPhysical, a.Verdict)
}

func FormatChiral(a ChiralActionPreflight) string {
	return fmt.Sprintf("carrier=%q algebra=%q left=%q right=%q character=%q D=%q differs=%t nonVacuousPossible=%t fullSC=%t verdict=%q", a.Carrier, a.Algebra, a.LeftAction, a.RightAction, a.CenterCharacter, a.DiracBlock, a.LeftRightActionsDiffer, a.NonVacuousPossible, a.FullSCPhysical, a.Verdict)
}

func FormatProbe(p Probe) string {
	return fmt.Sprintf("%s λ=%.6g Bdiag=%v χ=%.6g", p.Name, p.Lambda, p.BDiag, p.Character)
}

func FormatOneForm(a OneFormAudit) string {
	return fmt.Sprintf("probe={%s} x=%.6g y=%.6g Pdiag=%v normSq=%.12g nonzero=%t centralVanishes=%t physical=%t verdict=%q", FormatProbe(a.ProbeA), a.X, a.Y, a.SpatialOneFormDiag, a.FrobeniusNormSq, a.NonZero, a.CentralProbeVanishes, a.PhysicalOneForm, a.Verdict)
}

func FormatResidual(a OrderOneResidualAudit) string {
	return fmt.Sprintf("a={%s} b={%s} residual=%v normSq=%.12g zero=%t candidatePasses=%t fullOrderOne=%t verdict=%q", FormatProbe(a.ProbeA), FormatProbe(a.ProbeB), a.ResidualDiag, a.FrobeniusNormSq, a.ResidualZero, a.CandidatePasses, a.FullOrderOneProved, a.Verdict)
}

func FormatMomentRow(r MomentRow) string {
	return fmt.Sprintf("%s x=%.6g y=%.6g TrD2=%.12g TrD4=%.12g ratio=%.12g comment=%q", r.Name, r.X, r.Y, r.TraceD2, r.TraceD4, r.Ratio, r.Comment)
}

func FormatRatio(a InvariantRatioAudit) string {
	rows := make([]string, len(a.Rows))
	for i, r := range a.Rows {
		rows[i] = FormatMomentRow(r)
	}
	return fmt.Sprintf("family=%q selected=%t stable=%t dependsOnXY=%t gaugeProjection=%t scalarMap=%t higgs=%t rows=[%s] verdict=%q", a.FamilyFormula, a.XToYSelected, a.RatioStableAcrossFamily, a.DependsOnXY, a.GaugeProjectionDerived, a.ScalarFluctuationMapDerived, a.HiggsRatioDerived, strings.Join(rows, " | "), a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("empiricalSeal=%t ssbSeal=%t noMass=%t noVEV=%t noCutoff=%t noConnes=%t candidateNotPromoted=%t noHiggs=%t polluted=%t verdict=%q", a.EmpiricalYukawaSealPreserved, a.SpontaneousCarrierSealPreserved, a.NoObservedMassInserted, a.NoVEVInserted, a.NoCutoffScaleInserted, a.NoConnesRepresentationImported, a.CandidateNotPromoted, a.NoHiggsPredictionClaim, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := make([]string, 0, len(a.Obligations))
	for _, o := range a.Obligations {
		if o.Required && !o.Satisfied {
			missing = append(missing, o.Name)
		}
	}
	return fmt.Sprintf("obligations=%d missing=[%s] fullSC=%t J=%t bimodule=%t oneForms=%t xy=%t heat=%t next=%q verdict=%q", len(a.Obligations), strings.Join(missing, "; "), a.NeedFullSCRepresentation, a.NeedPhysicalJ, a.NeedOrderOnePassingBimod, a.NeedNonVacuousCalculus, a.NeedCanonicalXYSelector, a.NeedHeatKernelProjection, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("inherited=%t lift=%t fullSC=%t opposite=%t oneForms=%t orderOne=%t canonical=%t ratioStable=%t higgs=%t firewall=%t status=%q next=%q comment=%q", a.Gate269Inherited, a.FaithfulLiftAudited, a.FullSCRepresentation, a.PhysicalOppositeAction, a.CandidateOneFormsNonzero, a.CandidateOrderOnePasses, a.CanonicalDFDerived, a.RatioStable, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
