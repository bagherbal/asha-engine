// Package generation2flavorspectralbalancefunctionaltypeadmissibilityaudit implements
// Gate 595: Flavor Spectral Balance Functional Type-Admissibility Audit.
//
// Gate 594 constructed the observed environmental balance functional
//
//	B_flav(H_e,H_nu,H_u,H_d)
//	  = 1 - 8*pi*epsilon(H_e) - (1/4)Tr(P_e P_3^nu) + J(H_u,H_d).
//
// Gate 595 does not fit new numbers.  It type-checks whether this functional is
// admissible inside ASHA's current spectral framework, locates the primary
// native obstruction, and lists the exact theorem requirements for promoting
// the observed environmental balance to native law.
package generation2flavorspectralbalancefunctionaltypeadmissibilityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2flavorspectralorientationbalancefunctionalaudit"
)

const (
	AuditID = "GATE595-FLAVOR-SPECTRAL-BALANCE-FUNCTIONAL-TYPE-ADMISSIBILITY-AUDIT"

	StatusGate594Inherited             = "PASS_GATE594_B_FLAV_FUNCTIONAL_INHERITED"
	StatusTermTypingComplete           = "PASS_B_FLAV_TERM_TYPING_COMPLETE"
	StatusEnvironmentalWellDefined     = "PASS_B_FLAV_WELL_DEFINED_AS_ENVIRONMENTAL_SPECTRAL_FUNCTIONAL"
	StatusPolynomialAdmissible         = "PASS_POLYNOMIAL_SPECTRAL_INVARIANTS_ADMISSIBLE"
	StatusDetLogPfaffianAdmissible     = "PASS_DETERMINANT_LOGDETERMINANT_PFAFFIAN_INVARIANTS_ADMISSIBLE"
	StatusProjectorsConditionallyAdmit = "CONDITIONAL_SUPPORT_SPECTRAL_PROJECTORS_ADMISSIBLE_AS_OBSERVED_LEDGER"
	StatusCKMConditionallyAdmit        = "CONDITIONAL_SUPPORT_NORMALIZED_CKM_COMMUTATOR_ADMISSIBLE_AS_OBSERVED_LEDGER"
	StatusPMNSConditionallyAdmit       = "CONDITIONAL_SUPPORT_PMNS_PROJECTOR_OVERLAP_ADMISSIBLE_AS_OBSERVED_LEDGER"
	StatusCrossSectorEnvironmental     = "CONDITIONAL_SUPPORT_CROSS_SECTOR_BALANCE_ADMISSIBLE_AS_ENVIRONMENTAL_EQUATION_ONLY"
	StatusBFlavSharpensTypeTarget      = "CONDITIONAL_SUPPORT_B_FLAV_TYPE_AUDIT_SHARPENS_PROMOTION_REQUIREMENTS"
	StatusNoNativeFourthRootHE         = "FAILED_ROUTE_NO_NATIVE_H_E_FOURTH_ROOT_FUNCTIONAL"
	StatusNoNativeRootTrace            = "FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_ROOT_SPECTRUM_FUNCTIONAL"
	StatusNoNativeChamberWall          = "FAILED_ROUTE_NO_NATIVE_CHARGED_LEPTON_CHAMBER_WALL_FUNCTIONAL"
	StatusNoNativePMNSTerm             = "FAILED_ROUTE_PMNS_PROJECTOR_TERM_NOT_NATIVE_WITHOUT_PMNS_THEOREM"
	StatusNoNativeCKMTerm              = "FAILED_ROUTE_CKM_COMMUTATOR_TERM_NOT_NATIVE_WITHOUT_YUKAWA_CKM_THEOREM"
	StatusNoNativeCrossSectorEquation  = "FAILED_ROUTE_NO_NATIVE_CROSS_SECTOR_SCALAR_BALANCE_EQUATION"
	StatusNoNativeBFlavZero            = "FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM"
	StatusPrimaryObstructionEpsilon    = "FAILED_ROUTE_PRIMARY_NATIVE_OBSTRUCTION_IS_EPSILON_OF_H_E"
	StatusBFlavEnvironmental           = "FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL_TYPE_ONLY"
	StatusGate352Preserved             = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusNoFlavorDerivation           = "FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_YUKAWA_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedLedgersFirewalled    = "FIREWALL_PRESERVED_OBSERVED_FLAVOR_LEDGERS_NOT_PROMOTED_TO_NATIVE_LAW"
	StatusNoResidualFit                = "FIREWALL_PRESERVED_NO_NEW_NUMERICAL_RESIDUAL_FIT"
	StatusNoNewCarrierSelector         = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate595Boundary              = "FIREWALL_PRESERVED_GATE595_TYPE_ADMISSIBILITY_BOUNDARY"
)

type InheritedBalance struct {
	Functional          string
	Kappa               float64
	PMNSProjectorTrace  float64
	PMNSQuarter         float64
	JCKM                float64
	BFlav               float64
	Delta590            float64
	ResidualInsideSigma bool
	Verdict             string
}

type TermType struct {
	Name                 string
	Expression           string
	Inputs               []string
	RequiresLabels       []string
	RequiresSpectralCalc bool
	RequiresFractional   bool
	RequiresObservedData bool
	NativePresent        bool
	PrimaryObstruction   string
	Verdict              string
}

type TermTypingAudit struct {
	Epsilon  TermType
	PMNS     TermType
	CKM      TermType
	Complete bool
	Verdict  string
}

type AdmissibilityItem struct {
	Object              string
	CurrentASHAAdmits   bool
	Native              bool
	EnvironmentalLedger bool
	Reason              string
	Verdict             string
}

type AdmissibilityAudit struct {
	Items                         []AdmissibilityItem
	PolynomialInvariantsAdmitted  bool
	DeterminantPfaffianAdmitted   bool
	SpectralProjectorsAdmitted    bool
	FractionalFourthRootAdmitted  bool
	RootTraceAdmitted             bool
	ChamberWallFunctionalAdmitted bool
	NormalizedCKMAdmitted         bool
	CrossSectorEquationAdmitted   bool
	Verdict                       string
}

type NativeObstructionAudit struct {
	PrimaryObstruction string
	EpsilonHEBlocked   bool
	PMNSMoreAdmissible bool
	CKMMoreAdmissible  bool
	BFlavZeroBlocked   bool
	Explanation        string
	Verdict            string
}

type PromotionRequirement struct {
	Requirement string
	WhyNeeded   string
	Present     bool
	Verdict     string
}

type PromotionRequirements struct {
	Items               []PromotionRequirement
	AllPresent          bool
	ExactMissingTheorem string
	Verdict             string
}

type FirewallAudit struct {
	FitsNewResiduals       bool
	DerivesKoide           bool
	DerivesPMNS            bool
	DerivesCKM             bool
	DerivesYukawas         bool
	DerivesNeutrinoPhysics bool
	DerivesFlavorTexture   bool
	PromotesLedgers        bool
	AddsCarrier            bool
	AddsSelector           bool
	PreservesGate352       bool
	Verdict                string
}

type FinalVerdict struct {
	BFlavEnvironmentalWellDefined        bool
	PrimaryNativeObstruction             string
	ProjectorAndCommutatorMoreAdmissible bool
	NativeBFlavZeroTheoremPresent        bool
	RequiredTheorem                      string
	Decision                             string
	Verdict                              string
}

type Analysis struct {
	Inherited     InheritedBalance
	TermTyping    TermTypingAudit
	Admissibility AdmissibilityAudit
	Obstruction   NativeObstructionAudit
	Requirements  PromotionRequirements
	Firewalls     FirewallAudit
	Final         FinalVerdict
	Truth         string
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
	g594, err := generation2flavorspectralorientationbalancefunctionalaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate594 predecessor: %w", err)
	}
	inherited := inheritGate594(g594)
	terms := typeTerms()
	admissibility := auditAdmissibility()
	obstruction := auditNativeObstruction(terms, admissibility)
	requirements := definePromotionRequirements()
	firewalls := auditFirewalls()
	final := compileFinal(inherited, obstruction, requirements)
	truth := "Gate 595 type-checks B_flav as a well-defined environmental spectral functional while refusing native promotion.  Polynomial, determinant, log-determinant, and Pfaffian-style spectral invariants are admissible in the existing ASHA lanes, and PMNS projector / CKM normalized-commutator terms are conditionally admissible as observed spectral ledgers.  The native blockage is sharper: epsilon(H_e) requires eig(H_e)^(1/4), a root-spectrum/root-trace chamber wall functional, and a charged-lepton S3 chamber seal.  Gate352 remains binding, no B_flav=0 theorem exists, and the exact missing theorem is a native FlavorSpectralBalanceAdmissibilityAndZeroTheorem constructing epsilon(H_e), PMNS/CKM orientation data, and the cross-sector scalar balance from ASHA law."
	return Analysis{Inherited: inherited, TermTyping: terms, Admissibility: admissibility, Obstruction: obstruction, Requirements: requirements, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritGate594(g594 generation2flavorspectralorientationbalancefunctionalaudit.Analysis) InheritedBalance {
	return InheritedBalance{
		Functional:          g594.Balance.Definition,
		Kappa:               g594.Balance.LeftKappa,
		PMNSProjectorTrace:  4 * g594.Balance.ProjectorQuarter,
		PMNSQuarter:         g594.Balance.ProjectorQuarter,
		JCKM:                g594.Balance.JCKM,
		BFlav:               g594.Balance.BFlav,
		Delta590:            g594.Balance.Delta590,
		ResidualInsideSigma: g594.Balance.ResidualInsideSigma,
		Verdict:             StatusGate594Inherited,
	}
}

func typeTerms() TermTypingAudit {
	eps := TermType{
		Name:                 "epsilon(H_e)",
		Expression:           "H_e=Y_eY_e^dagger, eig(H_e)=y_i^2, x_i=eig_i(H_e)^(1/4), epsilon(H_e)=electron-wall coordinate of the charged-lepton Koide Fourier chamber",
		Inputs:               []string{"spectrum of H_e", "fourth-root spectral coordinates", "canonical charged-lepton chamber (e,mu,tau)", "electron-zero wall"},
		RequiresLabels:       []string{"charged-lepton generation ordering", "electron wall", "positive S3 chamber"},
		RequiresSpectralCalc: true,
		RequiresFractional:   true,
		RequiresObservedData: true,
		NativePresent:        false,
		PrimaryObstruction:   "fourth-root/root-spectrum chamber functional blocked by Gate352",
		Verdict:              strings.Join([]string{StatusNoNativeFourthRootHE, StatusNoNativeRootTrace, StatusNoNativeChamberWall, StatusPrimaryObstructionEpsilon}, ";"),
	}
	pmns := TermType{
		Name:                 "Tr(P_e P_3^nu)",
		Expression:           "Tr(P_e P_3^nu)=Tr(P_e U_PMNS P_3^nu U_PMNS^dagger)=|U_e3|^2=sin^2(theta13)",
		Inputs:               []string{"spectral projector P_e of H_e", "spectral projector P_3^nu of H_nu", "PMNS matrix / neutrino ordering ledger"},
		RequiresLabels:       []string{"electron projector", "third neutrino mass eigenstate", "normal-ordering convention", "PMNS convention"},
		RequiresSpectralCalc: true,
		RequiresFractional:   false,
		RequiresObservedData: true,
		NativePresent:        false,
		PrimaryObstruction:   "observed PMNS ledger only; no native PMNS projector theorem",
		Verdict:              strings.Join([]string{StatusPMNSConditionallyAdmit, StatusNoNativePMNSTerm}, ";"),
	}
	ckm := TermType{
		Name:                 "J(H_u,H_d)",
		Expression:           "normalized Jarlskog commutator area det([H_u,H_d])/(2i Vandermonde_u Vandermonde_d), equivalently Im(V_us V_cb V_ub^* V_cs^*) up to sign convention",
		Inputs:               []string{"H_u=Y_uY_u^dagger", "H_d=Y_dY_d^dagger", "nondegenerate spectra", "CKM orientation ledger"},
		RequiresLabels:       []string{"up-generation ordering", "down-generation ordering", "CKM orientation sign"},
		RequiresSpectralCalc: true,
		RequiresFractional:   false,
		RequiresObservedData: true,
		NativePresent:        false,
		PrimaryObstruction:   "observed CKM/Yukawa ledger only; no native Yukawa/CKM theorem",
		Verdict:              strings.Join([]string{StatusCKMConditionallyAdmit, StatusNoNativeCKMTerm}, ";"),
	}
	return TermTypingAudit{Epsilon: eps, PMNS: pmns, CKM: ckm, Complete: true, Verdict: StatusTermTypingComplete}
}

func auditAdmissibility() AdmissibilityAudit {
	items := []AdmissibilityItem{
		{Object: "polynomial spectral invariants of H_f", CurrentASHAAdmits: true, Native: true, EnvironmentalLedger: false, Reason: "Trace/determinant-type polynomial tests are already native-compatible in theorem-gated spectral lanes.", Verdict: StatusPolynomialAdmissible},
		{Object: "determinant/log-determinant/Pfaffian invariants", CurrentASHAAdmits: true, Native: true, EnvironmentalLedger: false, Reason: "Existing determinant/Pfaffian lanes are admissible, but they do not supply linear root traces or fourth-root chamber coordinates.", Verdict: StatusDetLogPfaffianAdmissible},
		{Object: "spectral projectors", CurrentASHAAdmits: true, Native: false, EnvironmentalLedger: true, Reason: "Projectors are mathematically valid once observed spectral ledgers and labels are supplied; no native PMNS theorem derives them.", Verdict: StatusProjectorsConditionallyAdmit},
		{Object: "fractional powers H_e^(1/4)", CurrentASHAAdmits: false, Native: false, EnvironmentalLedger: true, Reason: "The charged-lepton Koide coordinate needs eig(H_e)^(1/4), a fractional spectral-calculus operation not supplied by native ASHA algebra.", Verdict: StatusNoNativeFourthRootHE},
		{Object: "root traces / root-spectrum functionals", CurrentASHAAdmits: false, Native: false, EnvironmentalLedger: true, Reason: "Gate352 blocks promotion of root-trace/root-spectrum Koide functionals from current determinant/Pfaffian structures.", Verdict: StatusNoNativeRootTrace},
		{Object: "charged-lepton Fourier chamber-wall functional", CurrentASHAAdmits: false, Native: false, EnvironmentalLedger: true, Reason: "Requires fourth-root spectrum plus canonical S3 chamber and electron-wall label; no native chamber selector exists.", Verdict: StatusNoNativeChamberWall},
		{Object: "normalized CKM commutator/Jarlskog functional", CurrentASHAAdmits: true, Native: false, EnvironmentalLedger: true, Reason: "Well-defined from observed H_u,H_d with nondegenerate spectra and sign convention; no native Yukawa/CKM theorem supplies it.", Verdict: StatusCKMConditionallyAdmit},
		{Object: "cross-sector scalar balance equation B_flav=0", CurrentASHAAdmits: false, Native: false, EnvironmentalLedger: true, Reason: "No ASHA object mixes charged-lepton root chamber, PMNS projector overlap, and CKM commutator area into a native zero theorem.", Verdict: StatusNoNativeCrossSectorEquation},
	}
	return AdmissibilityAudit{
		Items:                         items,
		PolynomialInvariantsAdmitted:  true,
		DeterminantPfaffianAdmitted:   true,
		SpectralProjectorsAdmitted:    true,
		FractionalFourthRootAdmitted:  false,
		RootTraceAdmitted:             false,
		ChamberWallFunctionalAdmitted: false,
		NormalizedCKMAdmitted:         true,
		CrossSectorEquationAdmitted:   false,
		Verdict:                       strings.Join([]string{StatusEnvironmentalWellDefined, StatusProjectorsConditionallyAdmit, StatusCKMConditionallyAdmit, StatusPrimaryObstructionEpsilon, StatusNoNativeCrossSectorEquation}, ";"),
	}
}

func auditNativeObstruction(terms TermTypingAudit, adm AdmissibilityAudit) NativeObstructionAudit {
	primary := "epsilon(H_e): eig(H_e)^(1/4) root-spectrum chamber-wall functional"
	return NativeObstructionAudit{
		PrimaryObstruction: primary,
		EpsilonHEBlocked:   terms.Epsilon.RequiresFractional && !terms.Epsilon.NativePresent && !adm.FractionalFourthRootAdmitted && !adm.RootTraceAdmitted,
		PMNSMoreAdmissible: adm.SpectralProjectorsAdmitted && !terms.PMNS.NativePresent,
		CKMMoreAdmissible:  adm.NormalizedCKMAdmitted && !terms.CKM.NativePresent,
		BFlavZeroBlocked:   !adm.CrossSectorEquationAdmitted,
		Explanation:        "Projector overlap and normalized commutator terms are standard observed spectral invariants once labels are sealed.  The charged-lepton epsilon term is less native because it requires H_e^(1/4), root-spectrum extraction, and an S3 chamber-wall functional, all blocked by Gate352 in current ASHA.",
		Verdict:            strings.Join([]string{StatusPrimaryObstructionEpsilon, StatusNoNativeFourthRootHE, StatusNoNativeRootTrace, StatusNoNativeBFlavZero}, ";"),
	}
}

func definePromotionRequirements() PromotionRequirements {
	items := []PromotionRequirement{
		{Requirement: "native finite flavor spectral algebra", WhyNeeded: "B_flav must be constructed from ASHA law-space rather than observed ledgers.", Present: false, Verdict: "FAILED_ROUTE_NO_NATIVE_FLAVOR_SPECTRAL_ALGEBRA"},
		{Requirement: "root-spectrum / fourth-root functional on H_e", WhyNeeded: "epsilon(H_e) requires eig(H_e)^(1/4) and root-space Koide coordinates.", Present: false, Verdict: StatusNoNativeFourthRootHE},
		{Requirement: "charged-lepton chamber/orientation selector", WhyNeeded: "epsilon(H_e) is chamber/wall dependent and needs the canonical (e,mu,tau) electron-zero wall lawfully selected.", Present: false, Verdict: StatusNoNativeChamberWall},
		{Requirement: "PMNS projector theorem", WhyNeeded: "Tr(P_eP_3^nu) must be derived, including neutrino ordering and electron projector, not imported.", Present: false, Verdict: StatusNoNativePMNSTerm},
		{Requirement: "CKM normalized commutator theorem", WhyNeeded: "J(H_u,H_d) must be supplied by native Yukawa/CKM structure rather than observed CKM ledger.", Present: false, Verdict: StatusNoNativeCKMTerm},
		{Requirement: "cross-sector orientation balance principle", WhyNeeded: "ASHA must prove why charged-lepton root chamber, PMNS leakage, and CKM area satisfy B_flav=0.", Present: false, Verdict: StatusNoNativeCrossSectorEquation},
	}
	return PromotionRequirements{Items: items, AllPresent: false, ExactMissingTheorem: "FlavorSpectralBalanceAdmissibilityAndZeroTheorem: construct native H_f spectral algebra, epsilon(H_e), PMNS projector overlap, normalized CKM commutator area, and prove B_flav(H_e,H_nu,H_u,H_d)=0 with all label/chamber seals accounted for.", Verdict: strings.Join([]string{StatusBFlavSharpensTypeTarget, StatusNoNativeBFlavZero}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{
		FitsNewResiduals:       false,
		DerivesKoide:           false,
		DerivesPMNS:            false,
		DerivesCKM:             false,
		DerivesYukawas:         false,
		DerivesNeutrinoPhysics: false,
		DerivesFlavorTexture:   false,
		PromotesLedgers:        false,
		AddsCarrier:            false,
		AddsSelector:           false,
		PreservesGate352:       true,
		Verdict:                strings.Join([]string{StatusNoResidualFit, StatusNoFlavorDerivation, StatusObservedLedgersFirewalled, StatusGate352Preserved, StatusNoNewCarrierSelector, StatusGate595Boundary}, ";"),
	}
}

func compileFinal(in InheritedBalance, obstruction NativeObstructionAudit, req PromotionRequirements) FinalVerdict {
	native := req.AllPresent
	decision := fmt.Sprintf("Gate 595 certifies B_flav as a well-defined environmental spectral functional with B_flav=%.15g, but not as a native ASHA theorem.  The primary native obstruction is %s.  PMNS projector and CKM commutator terms are more type-admissible as observed spectral ledgers; epsilon(H_e) requires fourth-root/root-spectrum chamber calculus blocked by Gate352.  Native promotion requires %s", in.BFlav, obstruction.PrimaryObstruction, req.ExactMissingTheorem)
	return FinalVerdict{
		BFlavEnvironmentalWellDefined:        true,
		PrimaryNativeObstruction:             obstruction.PrimaryObstruction,
		ProjectorAndCommutatorMoreAdmissible: obstruction.PMNSMoreAdmissible && obstruction.CKMMoreAdmissible,
		NativeBFlavZeroTheoremPresent:        native,
		RequiredTheorem:                      req.ExactMissingTheorem,
		Decision:                             decision,
		Verdict:                              strings.Join([]string{StatusEnvironmentalWellDefined, StatusPrimaryObstructionEpsilon, StatusNoNativeBFlavZero, StatusBFlavEnvironmental, StatusGate595Boundary}, ";"),
	}
}

func Statuses() []string {
	return []string{StatusGate594Inherited, StatusTermTypingComplete, StatusEnvironmentalWellDefined, StatusPolynomialAdmissible, StatusDetLogPfaffianAdmissible, StatusProjectorsConditionallyAdmit, StatusCKMConditionallyAdmit, StatusPMNSConditionallyAdmit, StatusCrossSectorEnvironmental, StatusBFlavSharpensTypeTarget, StatusNoNativeFourthRootHE, StatusNoNativeRootTrace, StatusNoNativeChamberWall, StatusNoNativePMNSTerm, StatusNoNativeCKMTerm, StatusNoNativeCrossSectorEquation, StatusNoNativeBFlavZero, StatusPrimaryObstructionEpsilon, StatusBFlavEnvironmental, StatusGate352Preserved, StatusNoFlavorDerivation, StatusObservedLedgersFirewalled, StatusNoResidualFit, StatusNoNewCarrierSelector, StatusGate595Boundary}
}

func almostEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
