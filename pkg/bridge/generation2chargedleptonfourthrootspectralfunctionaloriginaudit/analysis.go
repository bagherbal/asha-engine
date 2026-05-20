// Package generation2chargedleptonfourthrootspectralfunctionaloriginaudit implements
// Gate 596: Charged-Lepton Fourth-Root Spectral Functional Origin Audit.
//
// Gate 595 located the primary native obstruction in the environmental
// flavor-balance functional at epsilon(H_e). Gate 596 asks whether ASHA
// currently contains, permits, or obstructs a native fourth-root spectral
// functional capable of producing the charged-lepton Koide chamber-wall
// coordinate. It is a type-origin/firewall audit, not a numerical fitting gate.
package generation2chargedleptonfourthrootspectralfunctionaloriginaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2flavorspectralbalancefunctionaltypeadmissibilityaudit"
)

const (
	AuditID = "GATE596-CHARGED-LEPTON-FOURTH-ROOT-SPECTRAL-FUNCTIONAL-ORIGIN-AUDIT"

	StatusGate595Inherited                = "PASS_GATE595_TYPE_ADMISSIBILITY_RESULT_INHERITED"
	StatusRootFunctionalTyped             = "PASS_CHARGED_LEPTON_ROOT_FUNCTIONAL_TYPED"
	StatusEpsilonEnvironmentalWellDefined = "PASS_EPSILON_H_E_WELL_DEFINED_AS_ENVIRONMENTAL_SPECTRAL_FUNCTIONAL"
	StatusPolynomialAdmissible            = "PASS_POLYNOMIAL_TRACES_ADMISSIBLE"
	StatusDetLogPfaffianAdmissible        = "PASS_DETERMINANT_LOGDETERMINANT_PFAFFIAN_STRUCTURES_ADMISSIBLE"
	StatusHeatKernelAdmissible            = "PASS_HEAT_KERNEL_SPECTRAL_ACTION_TERMS_ADMISSIBLE"
	StatusZetaEtaExistingLane             = "CONDITIONAL_SUPPORT_ZETA_ETA_EXIST_AS_SPECTRAL_LANES_NOT_FLAVOR_ROOT_THEOREM"
	StatusFunctionalSealDefined           = "CONDITIONAL_SUPPORT_FUNCTIONAL_CALCULUS_SEAL_DEFINED_FOR_H_E_ONE_FOURTH"
	StatusMinimalSealDefined              = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_ROOT_CHAMBER_SEAL_DEFINED"
	StatusClosestRouteFunctionalSeal      = "CONDITIONAL_SUPPORT_CLOSEST_LAWFUL_ROUTE_IS_EXPLICIT_FUNCTIONAL_CALCULUS_SEAL"
	StatusNoNativeFourthRoot              = "FAILED_ROUTE_NO_NATIVE_H_E_ONE_FOURTH_FUNCTIONAL"
	StatusNoNativeFractionalPowers        = "FAILED_ROUTE_NO_NATIVE_FLAVOR_FRACTIONAL_POWER_CALCULUS"
	StatusNoNativeRootTrace               = "FAILED_ROUTE_NO_NATIVE_TR_H_E_ONE_FOURTH_ROOT_TRACE"
	StatusNoNativeRootSpectrum            = "FAILED_ROUTE_NO_NATIVE_ROOT_SPECTRUM_CHAMBER_FUNCTIONAL"
	StatusNoNativeFourierWall             = "FAILED_ROUTE_NO_NATIVE_FOURIER_CIRCULANT_CHAMBER_WALL_COORDINATE"
	StatusNoNativeSpectralZeta            = "FAILED_ROUTE_NO_NATIVE_FINITE_FLAVOR_SPECTRAL_ZETA_AT_S_ONE_FOURTH"
	StatusCharacteristicRouteBlocked      = "FAILED_ROUTE_CHARACTERISTIC_POLYNOMIAL_STILL_REQUIRES_FOURTH_ROOT_DATA"
	StatusNoAbsoluteDirac                 = "FAILED_ROUTE_NO_ABSOLUTE_DIRAC_OPERATOR_WITH_SQRT_YUKAWA_SPECTRUM"
	StatusNoCirculantCarrier              = "FAILED_ROUTE_NO_NATIVE_GENERATION_CIRCULANT_CARRIER_SELECTING_X_E"
	StatusEpsilonEnvironmental            = "FAILED_ROUTE_EPSILON_H_E_REMAINS_ENVIRONMENTAL"
	StatusNoNativeBFlav                   = "FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL_WITHOUT_EPSILON_NATIVE_PROMOTION"
	StatusGate352Preserved                = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusNoKoideDerivation               = "FIREWALL_PRESERVED_NO_KOIDE_OR_CHARGED_LEPTON_MASS_DERIVATION"
	StatusNoFlavorDerivation              = "FIREWALL_PRESERVED_NO_PMNS_CKM_YUKAWA_NEUTRINO_OR_FLAVOR_TEXTURE_DERIVATION"
	StatusNoNewCarrierSelector            = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusNoBFlavNative                   = "FIREWALL_PRESERVED_NO_B_FLAV_ZERO_PROMOTION"
	StatusNoNumericalSearch               = "FIREWALL_PRESERVED_NO_NEW_NUMERICAL_CONSTANT_SEARCH"
	StatusGate596Boundary                 = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_FUNCTIONAL_BOUNDARY"
)

type InheritedTypeAudit struct {
	BFlavExpression    string
	BFlav              float64
	PrimaryObstruction string
	RequiredTheorem    string
	Gate595Verdict     string
	Verdict            string
}

type RootFunctionalSpec struct {
	Input                    string
	EigenvalueConvention     string
	RootCoordinates          string
	KoideFourierForm         string
	ChamberWallCoordinate    string
	RequiresFourthRoot       bool
	RequiresOrderedChamber   bool
	RequiresObservedLedger   bool
	NativePresent            bool
	EnvironmentalWellDefined bool
	Verdict                  string
}

type NativeSpectralOperation struct {
	Operation            string
	CurrentASHAAdmits    bool
	NativeForCurrentLane bool
	SuppliesEpsilonHE    bool
	Reason               string
	Verdict              string
}

type NativeSpectralAudit struct {
	Operations                       []NativeSpectralOperation
	PolynomialAdmissible             bool
	DeterminantLogPfaffianAdmissible bool
	HeatKernelAdmissible             bool
	ZetaEtaLanePresent               bool
	FractionalPowersNative           bool
	FourthRootTraceNative            bool
	OrderedChamberNative             bool
	FourierWallNative                bool
	Verdict                          string
}

type RouteAudit struct {
	Name            string
	Mechanism       string
	Status          string
	NativePromotion bool
	BridgeSeal      bool
	Reason          string
	Verdict         string
}

type RouteComparison struct {
	Routes             []RouteAudit
	ClosestLawfulRoute string
	AnyNativeRoute     bool
	Verdict            string
}

type MinimalSeal struct {
	Name          string
	Components    []string
	MayEnterBFlav bool
	NativeLaw     bool
	Reason        string
	Verdict       string
}

type FirewallAudit struct {
	FitsNewConstants           bool
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesYukawas             bool
	DerivesNeutrinos           bool
	DerivesFlavorTexture       bool
	AddsCarrier                bool
	AddsSelector               bool
	PromotesBFlavZero          bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	EpsilonEnvironmentalWellDefined bool
	NativeFourthRootPresent         bool
	ClosestPromotionRoute           string
	MinimalSealName                 string
	BFlavStillEnvironmental         bool
	RequiredTheorem                 string
	Decision                        string
	Verdict                         string
}

type Analysis struct {
	Inherited      InheritedTypeAudit
	RootFunctional RootFunctionalSpec
	NativeSpectral NativeSpectralAudit
	Routes         RouteComparison
	Seal           MinimalSeal
	Firewalls      FirewallAudit
	Final          FinalVerdict
	Truth          string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g595, err := generation2flavorspectralbalancefunctionaltypeadmissibilityaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate595 predecessor: %w", err)
	}
	inherited := inheritGate595(g595)
	root := defineRootFunctional()
	native := auditNativeSpectralOperations()
	routes := compareRoutes()
	seal := defineMinimalSeal()
	firewalls := auditFirewalls()
	final := compileFinal(inherited, root, native, routes, seal)
	truth := "Gate 596 certifies epsilon(H_e) as a well-defined environmental spectral functional requiring eig(H_e)^(1/4), a charged-lepton root spectrum, and a canonical S3 chamber-wall coordinate. Current ASHA admits polynomial, determinant/log-determinant/Pfaffian, and heat-kernel spectral lanes, and has zeta/eta lanes in other contexts, but none supplies a native finite flavor H_e^(1/4), Tr(H_e^(1/4)), ordered root chamber, Fourier wall coordinate, |D_e| operator with sqrt-Yukawa spectrum, or generation-circulant carrier. The closest lawful route is an explicit ChargedLeptonRootChamberSeal; B_flav remains environmental and Gate352 remains binding."
	return Analysis{Inherited: inherited, RootFunctional: root, NativeSpectral: native, Routes: routes, Seal: seal, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritGate595(g595 generation2flavorspectralbalancefunctionaltypeadmissibilityaudit.Analysis) InheritedTypeAudit {
	return InheritedTypeAudit{
		BFlavExpression:    g595.Inherited.Functional,
		BFlav:              g595.Inherited.BFlav,
		PrimaryObstruction: g595.Final.PrimaryNativeObstruction,
		RequiredTheorem:    g595.Final.RequiredTheorem,
		Gate595Verdict:     g595.Final.Verdict,
		Verdict:            StatusGate595Inherited,
	}
}

func defineRootFunctional() RootFunctionalSpec {
	return RootFunctionalSpec{
		Input:                    "positive Hermitian charged-lepton bilinear H_e=Y_eY_e^dagger",
		EigenvalueConvention:     "lambda_i=eig_i(H_e)=y_i^2 with canonical charged-lepton ordering (e,mu,tau)",
		RootCoordinates:          "x_i=lambda_i^(1/4)=sqrt(y_i)",
		KoideFourierForm:         "x_j=A[1+sqrt(2)R cos(delta+2*pi*j/3)]",
		ChamberWallCoordinate:    "epsilon(H_e)=135 degrees - delta in the positive canonical (e,mu,tau) chamber",
		RequiresFourthRoot:       true,
		RequiresOrderedChamber:   true,
		RequiresObservedLedger:   true,
		NativePresent:            false,
		EnvironmentalWellDefined: true,
		Verdict:                  strings.Join([]string{StatusRootFunctionalTyped, StatusEpsilonEnvironmentalWellDefined, StatusNoNativeFourthRoot, StatusNoNativeRootSpectrum}, ";"),
	}
}

func auditNativeSpectralOperations() NativeSpectralAudit {
	ops := []NativeSpectralOperation{
		{Operation: "polynomial traces Tr(H_e^n)", CurrentASHAAdmits: true, NativeForCurrentLane: true, SuppliesEpsilonHE: false, Reason: "Polynomial traces are admissible but cannot produce ordered fourth-root chamber coordinates without additional root data.", Verdict: StatusPolynomialAdmissible},
		{Operation: "determinant det(H_e) and elementary symmetric polynomials", CurrentASHAAdmits: true, NativeForCurrentLane: true, SuppliesEpsilonHE: false, Reason: "They encode products/symmetric data of y_i^2, not the ordered fourth-root vector x_i=eig_i(H_e)^(1/4).", Verdict: StatusDetLogPfaffianAdmissible},
		{Operation: "log determinant Tr log H_e", CurrentASHAAdmits: true, NativeForCurrentLane: true, SuppliesEpsilonHE: false, Reason: "Log-product information remains symmetric and does not supply the Koide Fourier chamber wall.", Verdict: StatusDetLogPfaffianAdmissible},
		{Operation: "Pfaffian / fermionic determinant structures", CurrentASHAAdmits: true, NativeForCurrentLane: true, SuppliesEpsilonHE: false, Reason: "Gate352 already blocks promotion from determinant/Pfaffian root-determinant behavior to linear root-trace/root-chamber Koide data.", Verdict: StatusDetLogPfaffianAdmissible},
		{Operation: "heat-kernel spectral-action terms", CurrentASHAAdmits: true, NativeForCurrentLane: true, SuppliesEpsilonHE: false, Reason: "Heat-kernel moments are native in spectral-action lanes, but current terms do not define charged-lepton H_e^(1/4) chamber coordinates.", Verdict: StatusHeatKernelAdmissible},
		{Operation: "zeta/eta spectral invariants", CurrentASHAAdmits: true, NativeForCurrentLane: true, SuppliesEpsilonHE: false, Reason: "ASHA has zeta/eta-style spectral lanes, but not a finite flavor spectral-zeta theorem evaluating Tr(H_e^s) at s=1/4 as a native Koide-wall observable.", Verdict: StatusZetaEtaExistingLane},
		{Operation: "fractional powers H_e^s", CurrentASHAAdmits: false, NativeForCurrentLane: false, SuppliesEpsilonHE: false, Reason: "No native finite flavor functional calculus currently admits H_e^s for noninteger s as an ASHA law object.", Verdict: StatusNoNativeFractionalPowers},
		{Operation: "fourth-root traces Tr(H_e^(1/4))", CurrentASHAAdmits: false, NativeForCurrentLane: false, SuppliesEpsilonHE: false, Reason: "The exact Koide root-trace family is the known Gate352 obstruction.", Verdict: StatusNoNativeRootTrace},
		{Operation: "ordered eigenvalue chamber functionals", CurrentASHAAdmits: false, NativeForCurrentLane: false, SuppliesEpsilonHE: false, Reason: "epsilon(H_e) requires ordering and an electron wall in an S3 chamber; current ASHA supplies no native chamber selector.", Verdict: StatusNoNativeRootSpectrum},
		{Operation: "Fourier/circulant chamber-wall coordinate epsilon(H_e)", CurrentASHAAdmits: false, NativeForCurrentLane: false, SuppliesEpsilonHE: false, Reason: "No native generation-plane/circulant carrier selects the charged-lepton root vector and its wall offset.", Verdict: StatusNoNativeFourierWall},
	}
	return NativeSpectralAudit{
		Operations:                       ops,
		PolynomialAdmissible:             true,
		DeterminantLogPfaffianAdmissible: true,
		HeatKernelAdmissible:             true,
		ZetaEtaLanePresent:               true,
		FractionalPowersNative:           false,
		FourthRootTraceNative:            false,
		OrderedChamberNative:             false,
		FourierWallNative:                false,
		Verdict:                          strings.Join([]string{StatusPolynomialAdmissible, StatusDetLogPfaffianAdmissible, StatusHeatKernelAdmissible, StatusZetaEtaExistingLane, StatusNoNativeFourthRoot, StatusNoNativeRootTrace, StatusNoNativeFourierWall}, ";"),
	}
}

func compareRoutes() RouteComparison {
	routes := []RouteAudit{
		{Name: "Route A — Functional calculus seal", Mechanism: "Explicitly seal H_e^(1/4), ordered eigenvalues, and epsilon(H_e) as environmental spectral calculus.", Status: "bridge-sealed", NativePromotion: false, BridgeSeal: true, Reason: "This is mathematically well-defined and sufficient for B_flav as environmental data, but it is not an ASHA-native derivation.", Verdict: strings.Join([]string{StatusFunctionalSealDefined, StatusMinimalSealDefined}, ";")},
		{Name: "Route B — Spectral-zeta route", Mechanism: "Try to define Tr(H_e^s) by finite flavor zeta/heat-kernel continuation at s=1/4.", Status: "blocked", NativePromotion: false, BridgeSeal: false, Reason: "Current ASHA has no finite flavor spectral-zeta theorem that makes s=1/4 a native Koide chamber observable.", Verdict: StatusNoNativeSpectralZeta},
		{Name: "Route C — Characteristic-polynomial route", Mechanism: "Express the root-vector chamber coordinate through symmetric functions of x_i.", Status: "blocked", NativePromotion: false, BridgeSeal: false, Reason: "Because x_i=lambda_i^(1/4), the route still requires fourth roots and then additionally ordered chamber data.", Verdict: StatusCharacteristicRouteBlocked},
		{Name: "Route D — Absolute-Dirac route", Mechanism: "Seek an operator |D_e| whose eigenvalues are sqrt(y_i) directly.", Status: "missing", NativePromotion: false, BridgeSeal: false, Reason: "No native absolute-Dirac charged-lepton operator with spectrum sqrt(y_i) is currently constructed.", Verdict: StatusNoAbsoluteDirac},
		{Name: "Route E — Circulant generation-plane operator", Mechanism: "Seek a native generation-plane/circulant operator whose eigenvector is the charged-lepton root vector x_e.", Status: "missing", NativePromotion: false, BridgeSeal: false, Reason: "No native generation/circulant carrier or selector currently supplies x_e or epsilon(H_e).", Verdict: StatusNoCirculantCarrier},
	}
	return RouteComparison{Routes: routes, ClosestLawfulRoute: "Route A — explicit ChargedLeptonRootChamberSeal / functional-calculus seal", AnyNativeRoute: false, Verdict: strings.Join([]string{StatusClosestRouteFunctionalSeal, StatusNoNativeSpectralZeta, StatusCharacteristicRouteBlocked, StatusNoAbsoluteDirac, StatusNoCirculantCarrier}, ";")}
}

func defineMinimalSeal() MinimalSeal {
	return MinimalSeal{
		Name: "ChargedLeptonRootChamberSeal",
		Components: []string{
			"observed environmental ledger H_e=Y_eY_e^dagger",
			"root coordinates x_i=eig_i(H_e)^(1/4)=sqrt(y_i)",
			"canonical charged-lepton chamber ordering (e,mu,tau)",
			"Fourier Koide form x_j=A[1+sqrt(2)R cos(delta+2*pi*j/3)]",
			"electron wall coordinate epsilon(H_e)=135 degrees-delta",
		},
		MayEnterBFlav: true,
		NativeLaw:     false,
		Reason:        "This is the minimal explicit datum needed for the charged-lepton side of B_flav after native fourth-root/root-chamber promotion fails.",
		Verdict:       strings.Join([]string{StatusMinimalSealDefined, StatusEpsilonEnvironmental}, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{
		FitsNewConstants:           false,
		DerivesKoide:               false,
		DerivesChargedLeptonMasses: false,
		DerivesPMNS:                false,
		DerivesCKM:                 false,
		DerivesYukawas:             false,
		DerivesNeutrinos:           false,
		DerivesFlavorTexture:       false,
		AddsCarrier:                false,
		AddsSelector:               false,
		PromotesBFlavZero:          false,
		PreservesGate352:           true,
		Verdict:                    strings.Join([]string{StatusNoNumericalSearch, StatusNoKoideDerivation, StatusNoFlavorDerivation, StatusNoNewCarrierSelector, StatusNoBFlavNative, StatusGate352Preserved, StatusGate596Boundary}, ";"),
	}
}

func compileFinal(in InheritedTypeAudit, root RootFunctionalSpec, native NativeSpectralAudit, routes RouteComparison, seal MinimalSeal) FinalVerdict {
	nativeFourthRoot := native.FractionalPowersNative || native.FourthRootTraceNative || root.NativePresent
	required := "ChargedLeptonFourthRootSpectralFunctionalTheorem: construct a native finite flavor spectral calculus admitting H_e^(1/4), a root-spectrum/root-trace observable, a canonical charged-lepton S3 chamber/electron-wall selector, and a proof that epsilon(H_e) may enter B_flav without environmental sealing."
	decision := fmt.Sprintf("Gate 596 finds epsilon(H_e) well-defined environmentally but not native. It requires %s and the current ASHA spectral lanes do not provide H_e^(1/4), Tr(H_e^(1/4)), a root-chamber selector, an absolute-Dirac sqrt-Yukawa operator, or a generation-circulant carrier. The closest lawful route is %s. Therefore %s enters B_flav only as an environmental seal and B_flav remains environmental.", root.RootCoordinates, routes.ClosestLawfulRoute, seal.Name)
	return FinalVerdict{
		EpsilonEnvironmentalWellDefined: root.EnvironmentalWellDefined,
		NativeFourthRootPresent:         nativeFourthRoot,
		ClosestPromotionRoute:           routes.ClosestLawfulRoute,
		MinimalSealName:                 seal.Name,
		BFlavStillEnvironmental:         true,
		RequiredTheorem:                 required,
		Decision:                        decision,
		Verdict:                         strings.Join([]string{StatusEpsilonEnvironmentalWellDefined, StatusNoNativeFourthRoot, StatusNoNativeRootTrace, StatusMinimalSealDefined, StatusNoNativeBFlav, StatusGate596Boundary}, ";"),
	}
}

func Statuses() []string {
	return []string{StatusGate595Inherited, StatusRootFunctionalTyped, StatusEpsilonEnvironmentalWellDefined, StatusPolynomialAdmissible, StatusDetLogPfaffianAdmissible, StatusHeatKernelAdmissible, StatusZetaEtaExistingLane, StatusFunctionalSealDefined, StatusMinimalSealDefined, StatusClosestRouteFunctionalSeal, StatusNoNativeFourthRoot, StatusNoNativeFractionalPowers, StatusNoNativeRootTrace, StatusNoNativeRootSpectrum, StatusNoNativeFourierWall, StatusNoNativeSpectralZeta, StatusCharacteristicRouteBlocked, StatusNoAbsoluteDirac, StatusNoCirculantCarrier, StatusEpsilonEnvironmental, StatusNoNativeBFlav, StatusGate352Preserved, StatusNoKoideDerivation, StatusNoFlavorDerivation, StatusNoNewCarrierSelector, StatusNoBFlavNative, StatusNoNumericalSearch, StatusGate596Boundary}
}

func almostEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
