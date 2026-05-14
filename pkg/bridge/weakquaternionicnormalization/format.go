package weakquaternionicnormalization

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate272Inheritance) string {
	return fmt.Sprintf("bimodule=%t opposite=%t nonVac=%t xy=%t a2a4=%t higgs=%t firewall=%t amps=[%s] verdict=%q", a.BimoduleExtracted, a.OppositeActionConstructed, a.NonVacuousEdgesExist, a.XYRatioLocked, a.A2A4Derived, a.HiggsRatioDerived, a.FirewallPreserved, strings.Join(a.SurvivingAmplitudeLabels, ","), a.Verdict)
}

func FormatSubCandidate(c SubBimoduleCandidate) string {
	return fmt.Sprintf("%s source=%q weak=%t H=%t chiral=%t native=%t physical=%t edges=[%s] reason=%q", c.Name, c.Source, c.UsesWeakSU2, c.UsesQuaternionicH, c.UsesChiralGrading, c.NativeToCPlusM3, c.SelectsPhysicalSMHF, strings.Join(c.SelectedEdgeLabels, ","), c.RejectedReason)
}

func FormatSieve(a PhysicalSubBimoduleSieve) string {
	parts := make([]string, len(a.CandidateSelectors))
	for i, c := range a.CandidateSelectors {
		parts[i] = FormatSubCandidate(c)
	}
	return fmt.Sprintf("universalSummands=%d universalDim=%d retained=[%s] weakNative=%t physical=%t candidates={%s} verdict=%q", a.UniversalSummands, a.UniversalComplexDimension, strings.Join(a.ChiralOrderOneEdgesRetained, ","), a.WeakQuaternionicNative, a.PhysicalSMHilbertDerived, strings.Join(parts, " | "), a.Verdict)
}

func FormatEdgeNormalization(e EdgeNormalization) string {
	return fmt.Sprintf("%s edge=%s right=%s dim=%d rank=%d k2=%.6g k4=%.6g needsNorm=%t detail=%q", e.Label, e.Edge, e.SharedRightModule, e.SharedRightDimension, e.MinimalRank, e.KappaD2, e.KappaD4, e.RequiresEdgeMapNorm, e.Detail)
}

func FormatInnerProduct(a InnerProductAudit) string {
	parts := make([]string, len(a.EdgeNormalizations))
	for i, e := range a.EdgeNormalizations {
		parts[i] = FormatEdgeNormalization(e)
	}
	return fmt.Sprintf("formula=%q orthogonal=%t canonicalTrace=%t kC=%.6g kQ=%.6g qOverC=%.6g geometric=%t edgeNorms=%t edges={%s} verdict=%q", a.Formula, a.OrthogonalMoritaSummands, a.CanonicalTraceOnSimpleModules, a.KappaCRatio, a.KappaQRatio, a.MultiplicityRatioQOverC, a.MultiplicitiesGeometric, a.EdgeNormsDerived, strings.Join(parts, " | "), a.Verdict)
}

func FormatXYRatio(a XYRatioAudit) string {
	return fmt.Sprintf("constraint=%q weightsKnown=%t equalWould=%q equalDerived=%t xyLocked=%t xy=%.12g unknowns=[%s] verdict=%q", a.CandidateConstraint, a.MultiplicityWeightsKnown, a.WouldEqualContributionFixXY, a.EqualContributionIsDerived, a.XOverYLocked, a.XOverYValue, strings.Join(a.Unknowns, "; "), a.Verdict)
}

func FormatMoment(c SpectralMomentCandidate) string {
	return fmt.Sprintf("%s D2=%.12g D4=%.12g ratio=%.12g", c.Name, c.TraceD2, c.TraceD4, c.Ratio)
}

func FormatSpectralTrace(a SpectralTraceAudit) string {
	parts := make([]string, len(a.Candidates))
	for i, c := range a.Candidates {
		parts[i] = FormatMoment(c)
	}
	return fmt.Sprintf("D2=%q D4=%q dependsXY=%t stable=%t a2a4=%t higgs=%t candidates=[%s] verdict=%q", a.TraceD2Formula, a.TraceD4Formula, a.RatioDependsOnXOverY, a.StableInvariant, a.A2A4Derived, a.HiggsRatioDerived, strings.Join(parts, " | "), a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("empirical=%t ssb=%t noMass=%t noVEV=%t noCutoff=%t noSMH=%t noHiggs=%t multNotAmp=%t polluted=%t verdict=%q", a.EmpiricalYukawaSealPreserved, a.SpontaneousCarrierSealPreserved, a.NoObservedMassInserted, a.NoVEVInserted, a.NoCutoffScaleInserted, a.NoSMQuaternionImportedAsTheorem, a.NoHiggsPredictionClaim, a.MultiplicityNotAmplitude, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := []string{}
	for _, c := range a.Criteria {
		if c.Required && !c.Satisfied {
			missing = append(missing, c.Name)
		}
	}
	return fmt.Sprintf("criteria=%d missing=[%s] needH=%t needJ=%t needNorm=%t needHeat=%t next=%q verdict=%q", len(a.Criteria), strings.Join(missing, "; "), a.NeedNativeWeakQuaternionicAlgebra, a.NeedPhysicalChargeConjugationJ, a.NeedEdgeNormOrAmplitudeAction, a.NeedHeatKernelProjection, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate272=%t sieve=%t inner=%t weights=%t physical=%t edgeNorms=%t xy=%t canonicalDF=%t a2a4=%t higgs=%t firewall=%t status=%q next=%q comment=%q", a.Gate272Inherited, a.PhysicalSieveAudited, a.InnerProductBuilt, a.TraceWeightsComputed, a.PhysicalSMHilbertDerived, a.EdgeNormsDerived, a.XYRatioLocked, a.CanonicalDFDerived, a.A2A4Derived, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
