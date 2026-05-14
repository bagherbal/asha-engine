package nativeweakquaternionicalgebra

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate273Inheritance) string {
	return fmt.Sprintf("innerProduct=%t traceWeights=%t κC=%.6g κQ=%.6g edgeNorms=%t xyLocked=%t a2a4=%t firewall=%t verdict=%s", a.InnerProductBuilt, a.TraceWeightsComputed, a.KappaC, a.KappaQ, a.EdgeNormsDerived, a.XYRatioLocked, a.A2A4Derived, a.FirewallPreserved, a.Verdict)
}

func FormatQuaternionic(a QuaternionClosureAudit) string {
	return fmt.Sprintf("plane=%q authority=%q localH=%t nativeSelected=%t unsealed=%t globalH=%t residuals=[I2 %.3g J2 %.3g K2 %.3g IJ-K %.3g JI+K %.3g anti %.3g] verdict=%s", a.SelectedPlane, a.SelectionAuthority, a.LocalHExtracted, a.NativeToSelectedDoublet, a.SelectionUnsealed, a.GlobalHSummandDerived, a.ISquareResidual, a.JSquareResidual, a.KSquareResidual, a.IJMinusKResidual, a.JIMinusNegativeKResidual, a.AntiCommutatorResidual, a.Verdict)
}

func FormatAlgebra(a FullAlgebraAudit) string {
	return fmt.Sprintf("C=%t M3=%t localH=%t candidate=%q realDim=%d complexEnvelopeDim=%d underSelector=%t exactSM=%t faithful=%t opposite=%t orderOne=%t verdict=%s", a.ComplexSummandInherited, a.ColorM3Inherited, a.LocalQuaternionicH, a.CandidateAlgebra, a.CandidateRealDimension, a.CandidateComplexEnvelopeDim, a.AssembledOnlyUnderSelector, a.ExactSMFiniteAlgebraDerived, a.FaithfulRepresentationReady, a.OppositeActionReady, a.OrderOneReady, a.Verdict)
}

func FormatHilbert(a PhysicalHilbertAudit) string {
	rows := []string{}
	for _, s := range a.CandidateSectors {
		rows = append(rows, fmt.Sprintf("%s[L=%s R=%s dim=%d derived=%t conditional=%t missing=%s]", s.Label, s.LeftAction, s.RightAction, s.ComplexDim, s.Derived, s.Conditional, s.Missing))
	}
	return fmt.Sprintf("universal=%t sectors={%s} Hleft=%t Cright=%t color=%t chirality=%t hypercharge=%t J=%t exactHF=%t verdict=%s", a.UniversalMoritaLedgerInherited, strings.Join(rows, "; "), a.LeftDoubletHActionAvailable, a.RightSingletCActionAvailable, a.ColorActionAvailable, a.ChiralGradingPhysical, a.HyperchargeAttachmentDerived, a.OppositeActionJDerived, a.ExactPhysicalHFDerived, a.Verdict)
}

func FormatAmplitude(a AmplitudeLockingAudit) string {
	return fmt.Sprintf("κC=%.6g κQ=%.6g HleftFactor=%.6g multiplicityUpdated=%t normC=%t normQ=%t xyLocked=%t equalNormsDerived=%t candidate=%q reason=%q verdict=%s", a.KappaC, a.KappaQ, a.QuaternionicLeftDoubletFactor, a.MultiplicityWeightsUpdated, a.EdgeNormCSelected, a.EdgeNormQSelected, a.XOverYLocked, a.EqualEdgeNormsDerived, a.CandidateIfEqualEdgeNorms, a.Reason, a.Verdict)
}

func FormatSpectralTrace(a SpectralTraceAudit) string {
	rows := []string{}
	for _, c := range a.Candidates {
		rows = append(rows, fmt.Sprintf("%s:D2=%.12g D4=%.12g ratio=%.12g", c.Name, c.TraceD2, c.TraceD4, c.Ratio))
	}
	return fmt.Sprintf("D2=%q D4=%q candidates=[%s] depends=%t stable=%t a2a4=%t higgs=%t verdict=%s", a.FormulaD2, a.FormulaD4, strings.Join(rows, "; "), a.RatioDependsOnXOverY, a.StableInvariant, a.A2A4Derived, a.HiggsRatioDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("noConnesImport=%t noWeakPlaneUnsealed=%t noMass=%t noYukawaAmp=%t noVEV=%t noHiggsClaim=%t localHNotGlobal=%t multiplicityNotAmp=%t polluted=%t verdict=%s", a.NoConnesAlgebraImportedAsTheorem, a.NoWeakPlaneUnsealed, a.NoObservedMassInserted, a.NoYukawaAmplitudeInserted, a.NoVEVInserted, a.NoHiggsPredictionClaimed, a.LocalHNotPromotedToGlobalH, a.MultiplicityNotAmplitude, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := []string{}
	for _, c := range a.Criteria {
		if c.Required && !c.Satisfied {
			missing = append(missing, c.Name+": "+c.Detail)
		}
	}
	return fmt.Sprintf("criteria=%d missing=[%s] weakPlane=%t physicalHF=%t J=%t edgeNorm=%t heat=%t next=%q verdict=%q", len(a.Criteria), strings.Join(missing, "; "), a.NeedUnsealedWeakPlane, a.NeedPhysicalFiniteHF, a.NeedPhysicalJ, a.NeedEdgeNormAction, a.NeedHeatKernelProjection, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("inherit=%t localH=%t closureExact=%t candidateAlg=%t exactSM=%t physicalHF=%t J=%t ampLocked=%t xy=%t a2a4=%t higgs=%t firewall=%t status=%s next=%q comment=%q", a.Gate273Inherited, a.LocalHExtracted, a.QuaternionClosureExact, a.CandidateAlgebraBuilt, a.ExactSMAlgebraDerived, a.PhysicalHFDerived, a.PhysicalJDerived, a.EdgeAmplitudesLocked, a.XYRatioLocked, a.A2A4Derived, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
