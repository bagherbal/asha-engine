package generation2chargedleptontraceringalgebraicrootchamberaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate598) string {
	return fmt.Sprintf("traceCable=%t rootMissing=%t missing=%q verdict=%q", a.NativeTraceCableVisible, a.RootOrientationMissing, a.MissingObject, a.Verdict)
}

func FormatTraceRing(a TraceRingTable) string {
	return fmt.Sprintf("ring=%q generators=%q nativePolynomial=%t admissible=%t verdict=%q", a.Ring, strings.Join(a.Generators, ","), a.NativePolynomial, a.Admissible, a.Verdict)
}

func FormatCharacteristic(a CharacteristicPolynomialAudit) string {
	return fmt.Sprintf("p=(%s,%s,%s) e1=%q e2=%q e3=%q chi=%q traceBuilt=%t nativePolynomial=%t verdict=%q", a.P1, a.P2, a.P3, a.ElementarySymmetricE1, a.ElementarySymmetricE2, a.ElementarySymmetricE3, a.Polynomial, a.BuiltFromTraceRing, a.NativePolynomial, a.Verdict)
}

func FormatRootExtension(a RootExtensionAudit) string {
	return fmt.Sprintf("eigenvalues=%q roots=%q positive=%t algebraic=%t requiresFourthRoot=%t native=%t avoidsGate596=%t closest=%q verdict=%q", a.EigenvalueDefinition, a.RootCoordinateDefinition, a.PositiveBranch, a.AlgebraicOverTraceRing, a.RequiresFourthRoot, a.Native, a.AvoidsGate596Obstruction, a.ClosestPromotionRoute, a.Verdict)
}

func FormatChamber(a ChamberFunctionalAudit) string {
	return fmt.Sprintf("form=%q epsilon=%q ordering=%t chamberSeal=%t chamber=%q algebraicRoot=%t nativePolynomial=%t verdict=%q", a.FourierForm, a.EpsilonDefinition, a.RequiresOrdering, a.RequiresChamberSeal, a.CanonicalChamber, a.AlgebraicOverRootExt, a.NativePolynomial, a.Verdict)
}

func FormatEpsilon(a EpsilonStatus) string {
	return fmt.Sprintf("wellDefined=%t nativePolynomial=%t algebraicTrace=%t fourthRootSeal=%t chamberSeal=%t rawInsertion=%t decision=%q verdict=%q", a.WellDefinedEnvironmental, a.NativePolynomial, a.AlgebraicOverTraceRing, a.RequiresFourthRootSeal, a.RequiresChamberSeal, a.PurelyRawInsertion, a.Decision, a.Verdict)
}

func FormatBFlav(a BFlavStatus) string {
	return fmt.Sprintf("expr=%q traceAnchored=%t chargedNative=%t ledgers=%t nativeZero=%t environmental=%t decision=%q verdict=%q", a.Expression, a.ChargedLeptonSideTraceAnchored, a.ChargedLeptonSideNative, a.PMNSCKMSidesEnvironmentalLedgers, a.NativeZeroTheorem, a.EnvironmentalOnly, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t pmnsCkmNu=%t H14Native=%t bflav=%t carrier=%t selector=%t constants=%t gate352=%t gate596=%t gate598=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNSCKMNeutrino, a.PromotesHEOneFourthNative, a.PromotesBFlavZero, a.AddsCarrier, a.AddsSelector, a.SearchesNewConstants, a.PreservesGate352, a.PreservesGate596, a.PreservesGate598, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("traceRing=%t chi=%t algebraicExt=%t epsilonPolynomial=%t H14Native=%t bflavNative=%t seal=%q decision=%q verdict=%q", a.TraceRingDefined, a.CharacteristicPolynomial, a.AlgebraicExtension, a.EpsilonNativePolynomial, a.HEOneFourthNative, a.BFlavNative, a.MinimalSeal, a.Decision, a.Verdict)
}
