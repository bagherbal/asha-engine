package generation2chargedleptonrootextensionbranchchambermonodromyaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate599) string {
	return fmt.Sprintf("traceRing=%t chi=%t algebraicExt=%t epsilonPolynomial=%t H14Native=%t bflavNative=%t seal=%q verdict=%q", a.TraceRingDefined, a.CharacteristicPolynomial, a.AlgebraicExtension, a.EpsilonNativePolynomial, a.HEOneFourthNative, a.BFlavNative, a.MinimalSeal, a.Verdict)
}

func FormatSplitting(a TraceRingToSplittingFieldTable) string {
	return fmt.Sprintf("base=%q chi=%q eigenvalues=%q field=%q ordersRoots=%t typed=%t verdict=%q", a.BaseRing, a.CharacteristicPolynomial, a.Eigenvalues, a.SplittingField, a.TraceRingOrdersRoots, a.Typed, a.Verdict)
}

func FormatMonodromy(a DiscriminantAndMonodromyAudit) string {
	return fmt.Sprintf("Delta=%q meaning=%q monodromy=%q square=%q nativeBranch=%t nativeOrdering=%t verdict=%q", a.Discriminant, a.DiscriminantMeaning, a.GenericMonodromy, a.SquareDiscriminant, a.NativeBranchSelector, a.NativeOrdering, a.Verdict)
}

func FormatFourthRoot(a FourthRootBranchAudit) string {
	return fmt.Sprintf("extension=%q sheets=%d positiveUnique=%t requiresPositivity=%t positivityNative=%t fourthRootNative=%t verdict=%q", a.Extension, a.ComplexSheetsPerEigenvalue, a.PositiveRealBranchUnique, a.RequiresPositivity, a.PositivityNative, a.FourthRootNative, a.Verdict)
}

func FormatChamber(a ChamberOrderingAudit) string {
	return fmt.Sprintf("order=%q chamber=%q wall=%q cyclic=%q traceWall=%t discWall=%t monodromyOrder=%t nativeChamber=%t verdict=%q", a.RequiredOrder, a.PositiveChamber, a.Wall, a.FourierCyclicOrder, a.TraceRingSelectsWall, a.DiscriminantSelectsWall, a.MonodromySelectsOrder, a.NativeChamberSelector, a.Verdict)
}

func FormatBranchSeal(a MinimalBranchSeal) string {
	return fmt.Sprintf("name=%q components=%q algebraic=%t native=%t environmental=%t verdict=%q", a.Name, strings.Join(a.Components, "; "), a.AlgebraicOverTrace, a.Native, a.Environmental, a.Verdict)
}

func FormatBFlav(a BFlavStatus) string {
	return fmt.Sprintf("expr=%q traceRing=%t splitting=%t fourthRoot=%t chamber=%t native=%t environmental=%t decision=%q verdict=%q", a.Expression, a.ChargedLeptonSideTraceRing, a.ChargedLeptonSideSplittingField, a.ChargedLeptonSideFourthRootBranch, a.ChargedLeptonSideChamberSeal, a.ChargedLeptonSideNative, a.EnvironmentalOnly, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t pmnsCkmNu=%t H14Native=%t chamberNative=%t bflav=%t carrier=%t selector=%t constants=%t gate352=%t gate596=%t gate599=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNSCKMNeutrino, a.PromotesHEOneFourthNative, a.PromotesChamberNative, a.PromotesBFlavZero, a.AddsCarrier, a.AddsSelector, a.SearchesNewConstants, a.PreservesGate352, a.PreservesGate596, a.PreservesGate599, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("splitting=%t traceOrders=%t nativeBranch=%t nativePositiveFourth=%t nativeChamber=%t epsilonAlgebraic=%t bflavNative=%t seal=%q decision=%q verdict=%q", a.SplittingFieldTyped, a.TraceRingOrdersSpectrum, a.NativeEigenvalueBranch, a.NativePositiveFourthRoot, a.NativeChamberSelector, a.EpsilonBranchAlgebraic, a.BFlavNative, a.MinimalSeal, a.Decision, a.Verdict)
}
