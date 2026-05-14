package moritabimodulesearch

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate271Inheritance) string {
	return fmt.Sprintf("carrier=%t car=%t gammaFail=%t dGammaFail=%t oneParticle=%t fullSC=%t J=%t higgs=%t firewall=%t verdict=%q", a.FullCarrierEnumerated, a.CARPassed, a.GammaFailedAdditivity, a.DGammaFailedAssociativity, a.OneParticleActionAvailable, a.FullSCRepDerived, a.PhysicalJDerived, a.HiggsRatioDerived, a.FirewallPreserved, a.Verdict)
}

func FormatObstruction(a ObstructionClassification) string {
	return fmt.Sprintf("algebra=%q fock=%q Gamma(mult=%t,add=%t) dGamma(add=%t,mult=%t) oneParticleAssoc=%t livesOnHF=%t fockSecondQuantized=%t verdict=%q", a.Algebra, a.FullFockCarrier, a.GammaMultiplicative, a.GammaAdditive, a.DGammaAdditive, a.DGammaMultiplicative, a.OneParticleAssociative, a.SpectralTripleLivesOnHF, a.FockIsSecondQuantizedKinematics, a.Verdict)
}

func FormatSummand(s BimoduleSummand) string {
	return fmt.Sprintf("%s dim=%d left=%s right=%s", s.Label, s.ComplexDimension, s.Left.Label, s.Right.Label)
}

func FormatBimodule(a BimoduleExtraction) string {
	parts := make([]string, len(a.Summands))
	for i, s := range a.Summands {
		parts[i] = FormatSummand(s)
	}
	return fmt.Sprintf("summands=[%s] dim=%d chiralDouble=%d particleAntiDouble=%d leftFaithful=%t rightFaithful=%t linear=%t unital=%t assoc=%t star=%t commute=%t fullFock=%t verdict=%q", strings.Join(parts, "; "), a.TotalComplexDimension, a.ChiralDoubleDimension, a.ParticleAntiParticleDoubleDimension, a.LeftActionFaithful, a.RightOppositeActionFaithful, a.Linear, a.Unital, a.Associative, a.StarCompatible, a.LeftRightCommute, a.FullFockCarrierUsed, a.Verdict)
}

func FormatOpposite(a OppositeActionAudit) string {
	return fmt.Sprintf("formula=%q constructed=%t antiLinearJRequired=%t faithfulOpposite=%t physicalSemantics=%t verdict=%q", a.Formula, a.Constructed, a.AntiLinearJRequiredForPhysicalChargeConjugation, a.AlgebraicOppositeActionFaithful, a.ParticleAntiParticleSemanticsDerived, a.Verdict)
}

func FormatEdge(e DiracEdge) string {
	return fmt.Sprintf("%s↔%s sameL=%t sameR=%t nonVac=%t orderOne=%t amp=%q reason=%q", e.From, e.To, e.SameLeft, e.SameRight, e.NonVacuousOneForm, e.OrderOneAllowed, e.AmplitudeLabel, e.Reason)
}

func FormatOrderOne(a OrderOneAudit) string {
	parts := make([]string, len(a.Edges))
	for i, e := range a.Edges {
		parts[i] = FormatEdge(e)
	}
	return fmt.Sprintf("rule=%q allowedNonVac=%d allowedVac=%d rejected=%d orderOneAllowed=%t nonVac=%t canonicalDF=%t amps=[%s] xy=%t edges=[%s] verdict=%q", a.Rule, a.NonVacuousAllowedEdges, a.VacuousAllowedEdges, a.RejectedEdges, a.OrderOneSatisfiedForAllowedEdges, a.NonVacuousOneFormsAvailable, a.CanonicalDFSelected, strings.Join(a.SurvivingAmplitudeLabels, ","), a.XYRatioLocked, strings.Join(parts, " | "), a.Verdict)
}

func FormatRatio(a SpectralRatioAudit) string {
	return fmt.Sprintf("DF=%q D2=%q D4=%q ratio=%q dependsXY=%t a2a4=%t higgs=%t missing=%q verdict=%q", a.CandidateDFShape, a.TraceD2Formula, a.TraceD4Formula, a.RatioFormula, a.DependsOnXOverY, a.A2A4Derived, a.HiggsRatioDerived, a.MissingSelector, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("empirical=%t ssb=%t noMass=%t noVEV=%t noCutoff=%t noConnesSM=%t notPromoted=%t noHiggs=%t polluted=%t verdict=%q", a.EmpiricalYukawaSealPreserved, a.SpontaneousCarrierSealPreserved, a.NoObservedMassInserted, a.NoVEVInserted, a.NoCutoffScaleInserted, a.NoConnesSMAlgebraImported, a.BimoduleNotPromotedToSM, a.NoHiggsPredictionClaim, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := []string{}
	for _, c := range a.Criteria {
		if c.Required && !c.Satisfied {
			missing = append(missing, c.Name)
		}
	}
	return fmt.Sprintf("criteria=%d missing=[%s] weakSelector=%t innerProduct=%t spectralProjection=%t ampSelector=%t next=%q verdict=%q", len(a.Criteria), strings.Join(missing, "; "), a.NeedWeakQuaternionicOrChiralSelector, a.NeedCanonicalInnerProductNormalization, a.NeedFiniteSpectralActionProjection, a.NeedAmplitudeSelector, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate271=%t obstruction=%t bimodule=%t opposite=%t nonVacOrderOne=%t physicalSM=%t canonicalDF=%t xy=%t a2a4=%t higgs=%t firewall=%t status=%q next=%q comment=%q", a.Gate271Inherited, a.ObstructionClassified, a.BimoduleExtracted, a.FaithfulOppositeAction, a.NonVacuousOrderOneEdges, a.PhysicalSMHilbertDerived, a.CanonicalDFDerived, a.XYRatioLocked, a.A2A4Derived, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
