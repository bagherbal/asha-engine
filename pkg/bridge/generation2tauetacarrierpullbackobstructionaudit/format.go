package generation2tauetacarrierpullbackobstructionaudit

import (
	"fmt"
	"strings"
)

func FormatIntVec(v []int) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}

func FormatInherited(a InheritedGate555Audit) string {
	return fmt.Sprintf("selector=%t bminusl4to13=%t uniqueWeakPlane=%t tauPullback=%t tauSealed=%t contactBlocked=%t verdict=%q", a.SelectorTheoremProved, a.BMinusLFourToOnePlusThree, a.BMinusLUniqueWeakPlane, a.TauEtaPullbackValid, a.TauEtaSealedCapacity, a.ContactQuarticStillBlocked, a.Verdict)
}

func FormatType(a TypeClassificationAudit) string {
	return fmt.Sprintf("tau=%s abs=%s traceVector=%t nativeSpectrum=%t diagonalEndomorphism=%t character=%t coefficientVector=%t sealed=%t functional=%q operators=%v verdict=%q reason=%q", FormatIntVec(a.TauEta), FormatIntVec(a.AbsTauEta), a.IsTraceValueVector, a.IsSpectrumOfNativeOperator, a.IsDiagonalEndomorphism, a.IsCharacter, a.IsCoefficientVectorInNativeBasis, a.IsSealedBookkeepingDatum, a.SourceFunctional, a.SourceOperators, a.Verdict, a.Reason)
}

func FormatSource(a SourceAlgebraAudit) string {
	names := make([]string, len(a.Candidates))
	for i, c := range a.Candidates {
		names[i] = fmt.Sprintf("%s=%s native=%t hand=%t", c.Name, c.Presentation, c.FoundAsNativeProjectAlgebra, c.InsertedByHand)
	}
	return fmt.Sprintf("native=%t name=%q unit=%t recovered=%t candidates=[%s] verdict=%q reason=%q", a.NativeSourceAlgebraExists, a.NativeSourceAlgebraName, a.HasUnit, a.TraceOrEigenvalueDataRecovered, strings.Join(names, "; "), a.Verdict, a.Reason)
}

func FormatRepresentation(a RepresentationAudit) string {
	rows := make([]string, len(a.Candidates))
	for i, c := range a.Candidates {
		rows[i] = fmt.Sprintf("target=%s source=%t constructed=%t unit=%t rho1=%t compat=%t rejected=%q", c.Target, c.NativeSourceAvailable, c.RepresentationConstructed, c.UnitPreserving, c.RhoOneEqualsIdentity, c.CompatibilityKnown, c.RejectedReason)
	}
	return fmt.Sprintf("valid=%t target=%q rho1=%t candidates=[%s] verdict=%q reason=%q", a.AnyValidUnitPreservingRepresentation, a.ValidTarget, a.RhoOneIsIdentity, strings.Join(rows, "; "), a.Verdict, a.Reason)
}

func FormatSelector(a SelectorConsequenceAudit) string {
	return fmt.Sprintf("formalAbs=%s mult=%v formalComm=%s dim=%d formulaIfRep=%t validRep=%t nativeSelector=%t canonicalU12=%t basisDependent=%t verdict=%q reason=%q", FormatIntVec(a.FormalAbsTauEta), a.FormalMultiplicityPattern, a.FormalCommutant, a.FormalCommutantDimension, a.Gate555SelectorFormulaAppliesIfRepresentationExists, a.ValidRepresentationExists, a.ProducesNativeSelector, a.CanonicalU12Selected, a.BasisDependentIfForced, a.Verdict, a.Reason)
}

func FormatBMinusL(a BMinusLCompatibilityAudit) string {
	return fmt.Sprintf("actsOnWSpatial=%t validRep=%t restricted=%q formalCommZero=%t nativeVerified=%t verdict=%q reason=%q", a.ActsOnWSpatial, a.ValidRepresentationExists, a.BMinusLRestrictedToWSpatial, a.FormalCommutatorWithBMinusLZero, a.NativeCompatibilityVerified, a.Verdict, a.Reason)
}

func FormatSpectralTriple(a SpectralTripleCompatibilityAudit) string {
	return fmt.Sprintf("proposed=%t gamma=%t J=%t D=%t firstOrder=%t bminusl=%t missing=%v nativeAllowed=%t verdict=%q", a.ProposedNativeRepresentation, a.GammaCompatibilityAvailable, a.JCompatibilityAvailable, a.DCompatibilityAvailable, a.FirstOrderCompatibilityAvailable, a.BMinusLCompatibilityAvailable, a.MissingData, a.NativeSpectralTriplePromotionAllowed, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("weakIso=%t genMass=%t higgs=%t yukawa=%t ckm=%t observed=%t formalNative=%t diagNative=%t polluted=%t verdict=%q", a.PromotedToWeakIsospin, a.PromotedToGenerationMassHierarchy, a.PromotedToHiggs, a.PromotedToYukawa, a.PromotedToCKMPMNS, a.InsertedObservedFlavorData, a.InsertedFormalAlgebraAsNative, a.InsertedDiagonalMatrixAsNative, a.NativeRegistryPolluted, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("operator=%t traceOnly=%t nativeAlgebra=%t unitRep=%t canonical2plus1=%t next=%q verdict=%q", a.TauEtaOperator, a.TauEtaOnlyTraceVector, a.NativeSourceAlgebraExists, a.UnitPreservingRepresentationExists, a.CanonicalTwoPlusOneSelectorOnWSpatial, a.MissingNextTheorem, a.Verdict)
}
