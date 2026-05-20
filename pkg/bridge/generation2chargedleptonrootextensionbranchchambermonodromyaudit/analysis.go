// Package generation2chargedleptonrootextensionbranchchambermonodromyaudit implements
// Gate 600: Charged-Lepton Root-Extension Branch and Chamber Monodromy Audit.
//
// Gate 599 anchored epsilon(H_e) in the native charged-lepton trace ring only
// after adjoining positive fourth roots and a charged-lepton chamber seal. Gate
// 600 asks what branch, splitting-field, monodromy, positivity, and chamber data
// are required to pass from the native trace ring to the observed Koide chamber
// coordinate. It is a type/branch audit, not a numerical fitting gate and not a
// native derivation of Koide, charged-lepton masses, PMNS, CKM, neutrino physics,
// flavor texture, or B_flav=0.
package generation2chargedleptonrootextensionbranchchambermonodromyaudit

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2chargedleptontraceringalgebraicrootchamberaudit"
)

const (
	AuditID = "GATE600-CHARGED-LEPTON-ROOT-EXTENSION-BRANCH-CHAMBER-MONODROMY-AUDIT"

	StatusGate599Inherited                 = "PASS_TRACE_RING_TO_CHARACTERISTIC_POLYNOMIAL_INHERITED"
	StatusCubicSplittingFieldTyped         = "PASS_CUBIC_SPLITTING_FIELD_TYPED"
	StatusDiscriminantDefined              = "PASS_DISCRIMINANT_AND_MONODROMY_DATA_TYPED"
	StatusTraceRingUnorderedSpectrum       = "PASS_TRACE_RING_GIVES_UNORDERED_SPECTRUM_ONLY"
	StatusGenericS3MonodromyTyped          = "PASS_GENERIC_CUBIC_MONODROMY_S3_OR_SUBGROUP_IF_DISCRIMINANT_SQUARE_TYPED"
	StatusFourthRootBranchTyped            = "PASS_FOURTH_ROOT_BRANCH_STRUCTURE_TYPED"
	StatusPositiveObservedBranchTyped      = "CONDITIONAL_SUPPORT_POSITIVE_REAL_FOURTH_ROOT_BRANCH_DEFINED_AS_OBSERVED_SEAL"
	StatusChamberOrderingTyped             = "PASS_KOIDE_CHAMBER_ORDERING_DATA_TYPED"
	StatusEpsilonBranchAlgebraic           = "CONDITIONAL_SUPPORT_EPSILON_H_E_BRANCH_ALGEBRAIC_OVER_TRACE_RING"
	StatusRootBranchChamberSealDefined     = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_ROOT_BRANCH_CHAMBER_SEAL_DEFINED"
	StatusBFlavBranchAnchoredEnvironmental = "CONDITIONAL_SUPPORT_B_FLAV_CHARGED_LEPTON_SIDE_BRANCH_ANCHORED_BUT_ENVIRONMENTAL"
	StatusTraceRingNoOrdering              = "FAILED_ROUTE_TRACE_RING_DOES_NOT_SELECT_CHARGED_LEPTON_ORDERING"
	StatusNoNativeBranchOrderTheorem       = "FAILED_ROUTE_NO_NATIVE_EIGENVALUE_BRANCH_OR_ORDER_THEOREM"
	StatusNoNativePositiveFourthRoot       = "FAILED_ROUTE_NO_NATIVE_POSITIVE_FOURTH_ROOT_BRANCH_THEOREM"
	StatusNoNativeElectronWall             = "FAILED_ROUTE_NO_NATIVE_ELECTRON_WALL_OR_CHAMBER_SELECTOR"
	StatusNoNativeCyclicOrder              = "FAILED_ROUTE_NO_NATIVE_FOURIER_CYCLIC_ORDERING_SELECTOR"
	StatusNoNativeEpsilonBranchTheorem     = "FAILED_ROUTE_NO_NATIVE_EPSILON_BRANCH_THEOREM"
	StatusBFlavEnvironmental               = "FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL"
	StatusNoNativeBFlavZero                = "FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_FROM_BRANCH_DATA"
	StatusGate352Preserved                 = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate596Preserved                 = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING"
	StatusGate599Preserved                 = "FIREWALL_PRESERVED_GATE599_TRACE_RING_EXTENSION_BOUNDARY_REMAINS_BINDING"
	StatusNoKoideDerivation                = "FIREWALL_PRESERVED_NO_KOIDE_DERIVATION"
	StatusNoChargedLeptonMassDerivation    = "FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_DERIVATION"
	StatusNoPMNSCKMNeutrinoFlavorDeriv     = "FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusNoHEOneFourthNativePromotion     = "FIREWALL_PRESERVED_NO_H_E_ONE_FOURTH_NATIVE_PROMOTION"
	StatusNoChamberNativePromotion         = "FIREWALL_PRESERVED_NO_CHAMBER_NATIVE_PROMOTION"
	StatusNoBFlavNativePromotion           = "FIREWALL_PRESERVED_NO_B_FLAV_ZERO_NATIVE_PROMOTION"
	StatusNoNewNumericalFit                = "FIREWALL_PRESERVED_NO_NEW_NUMERICAL_CONSTANT_SEARCH"
	StatusNoNewCarrierSelector             = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate600Boundary                  = "FIREWALL_PRESERVED_GATE600_BRANCH_CHAMBER_MONODROMY_BOUNDARY"
)

type InheritedGate599 struct {
	TraceRingDefined         bool
	CharacteristicPolynomial bool
	AlgebraicExtension       bool
	EpsilonNativePolynomial  bool
	HEOneFourthNative        bool
	BFlavNative              bool
	MinimalSeal              string
	Verdict                  string
}

type TraceRingToSplittingFieldTable struct {
	BaseRing                 string
	CharacteristicPolynomial string
	Eigenvalues              string
	SplittingField           string
	TraceRingOrdersRoots     bool
	Typed                    bool
	Verdict                  string
}

type DiscriminantAndMonodromyAudit struct {
	Discriminant         string
	DiscriminantMeaning  string
	GenericMonodromy     string
	SquareDiscriminant   string
	NativeBranchSelector bool
	NativeOrdering       bool
	Verdict              string
}

type FourthRootBranchAudit struct {
	Extension                  string
	ComplexSheetsPerEigenvalue int
	PositiveRealBranchUnique   bool
	RequiresPositivity         bool
	PositivityNative           bool
	FourthRootNative           bool
	Verdict                    string
}

type ChamberOrderingAudit struct {
	RequiredOrder           string
	PositiveChamber         string
	Wall                    string
	FourierCyclicOrder      string
	TraceRingSelectsWall    bool
	DiscriminantSelectsWall bool
	MonodromySelectsOrder   bool
	NativeChamberSelector   bool
	Verdict                 string
}

type MinimalBranchSeal struct {
	Name               string
	Components         []string
	AlgebraicOverTrace bool
	Native             bool
	Environmental      bool
	Verdict            string
}

type BFlavStatus struct {
	Expression                        string
	ChargedLeptonSideTraceRing        bool
	ChargedLeptonSideSplittingField   bool
	ChargedLeptonSideFourthRootBranch bool
	ChargedLeptonSideChamberSeal      bool
	ChargedLeptonSideNative           bool
	EnvironmentalOnly                 bool
	Decision                          string
	Verdict                           string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNSCKMNeutrino     bool
	PromotesHEOneFourthNative  bool
	PromotesChamberNative      bool
	PromotesBFlavZero          bool
	AddsCarrier                bool
	AddsSelector               bool
	SearchesNewConstants       bool
	PreservesGate352           bool
	PreservesGate596           bool
	PreservesGate599           bool
	Verdict                    string
}

type FinalVerdict struct {
	SplittingFieldTyped      bool
	TraceRingOrdersSpectrum  bool
	NativeEigenvalueBranch   bool
	NativePositiveFourthRoot bool
	NativeChamberSelector    bool
	EpsilonBranchAlgebraic   bool
	BFlavNative              bool
	MinimalSeal              string
	Decision                 string
	Verdict                  string
}

type Analysis struct {
	Inherited  InheritedGate599
	Splitting  TraceRingToSplittingFieldTable
	Monodromy  DiscriminantAndMonodromyAudit
	FourthRoot FourthRootBranchAudit
	Chamber    ChamberOrderingAudit
	BranchSeal MinimalBranchSeal
	BFlav      BFlavStatus
	Firewalls  FirewallAudit
	Final      FinalVerdict
	Truth      string
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
	g599, err := generation2chargedleptontraceringalgebraicrootchamberaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate599 predecessor: %w", err)
	}
	inherited := inheritGate599(g599)
	splitting := constructSplittingField()
	monodromy := auditDiscriminantAndMonodromy()
	fourthRoot := auditFourthRootBranch()
	chamber := auditChamberOrdering()
	branchSeal := defineMinimalBranchSeal(splitting, monodromy, fourthRoot, chamber)
	bflav := updateBFlavStatus(branchSeal)
	firewalls := auditFirewalls()
	final := compileFinal(splitting, monodromy, fourthRoot, chamber, branchSeal, bflav, firewalls)
	truth := "Gate 600 decomposes epsilon(H_e) into a native trace ring, a cubic splitting field, a positive fourth-root branch, and a charged-lepton chamber/wall seal. Law gives chi_e(lambda); history supplies the root branch, positivity, ordering, cyclic chamber, and electron-wall choice. B_flav is therefore branch-anchored but remains environmental."
	return Analysis{Inherited: inherited, Splitting: splitting, Monodromy: monodromy, FourthRoot: fourthRoot, Chamber: chamber, BranchSeal: branchSeal, BFlav: bflav, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritGate599(a generation2chargedleptontraceringalgebraicrootchamberaudit.Analysis) InheritedGate599 {
	return InheritedGate599{
		TraceRingDefined:         a.Final.TraceRingDefined,
		CharacteristicPolynomial: a.Final.CharacteristicPolynomial,
		AlgebraicExtension:       a.Final.AlgebraicExtension,
		EpsilonNativePolynomial:  a.Final.EpsilonNativePolynomial,
		HEOneFourthNative:        a.Final.HEOneFourthNative,
		BFlavNative:              a.Final.BFlavNative,
		MinimalSeal:              a.Final.MinimalSeal,
		Verdict:                  StatusGate599Inherited,
	}
}

func constructSplittingField() TraceRingToSplittingFieldTable {
	return TraceRingToSplittingFieldTable{
		BaseRing:                 "R_e=Q[p1,p2,p3], p_k=Tr(H_e^k)",
		CharacteristicPolynomial: "chi_e(lambda)=lambda^3-e1*lambda^2+e2*lambda-e3",
		Eigenvalues:              "lambda_i are roots of chi_e(lambda)",
		SplittingField:           "K_e=Frac(R_e)(lambda_1,lambda_2,lambda_3)",
		TraceRingOrdersRoots:     false,
		Typed:                    true,
		Verdict:                  StatusCubicSplittingFieldTyped,
	}
}

func auditDiscriminantAndMonodromy() DiscriminantAndMonodromyAudit {
	return DiscriminantAndMonodromyAudit{
		Discriminant:         "Delta_e=prod_{i<j}(lambda_i-lambda_j)^2",
		DiscriminantMeaning:  "Delta_e detects eigenvalue collision walls and controls whether the generic cubic Galois/monodromy group is S3 or a subgroup.",
		GenericMonodromy:     "generic S3 over the trace ring; A3 or smaller only under additional square-discriminant/specialization data",
		SquareDiscriminant:   "not supplied natively; no square-discriminant theorem in current ASHA flavor ledger",
		NativeBranchSelector: false,
		NativeOrdering:       false,
		Verdict:              StatusDiscriminantDefined,
	}
}

func auditFourthRootBranch() FourthRootBranchAudit {
	return FourthRootBranchAudit{
		Extension:                  "K_e(lambda_i^(1/4)); x_i^4=lambda_i",
		ComplexSheetsPerEigenvalue: 4,
		PositiveRealBranchUnique:   true,
		RequiresPositivity:         true,
		PositivityNative:           false,
		FourthRootNative:           false,
		Verdict:                    StatusFourthRootBranchTyped,
	}
}

func auditChamberOrdering() ChamberOrderingAudit {
	return ChamberOrderingAudit{
		RequiredOrder:           "canonical charged-lepton order (e,mu,tau)",
		PositiveChamber:         "positive Koide S3 chamber containing the observed charged-lepton ray",
		Wall:                    "electron-zero wall delta=135 degrees; epsilon(H_e)=135 degrees-delta",
		FourierCyclicOrder:      "the cyclic order used in x_j=A[1+sqrt(2)R cos(delta+2*pi*j/3)]",
		TraceRingSelectsWall:    false,
		DiscriminantSelectsWall: false,
		MonodromySelectsOrder:   false,
		NativeChamberSelector:   false,
		Verdict:                 StatusChamberOrderingTyped,
	}
}

func defineMinimalBranchSeal(s TraceRingToSplittingFieldTable, m DiscriminantAndMonodromyAudit, r FourthRootBranchAudit, c ChamberOrderingAudit) MinimalBranchSeal {
	return MinimalBranchSeal{
		Name: "ChargedLeptonRootBranchChamberSeal",
		Components: []string{
			"cubic splitting branch lambda_e,lambda_mu,lambda_tau of chi_e(lambda)",
			"positive fourth roots x_i=lambda_i^(1/4)",
			"canonical chamber order (e,mu,tau)",
			"Fourier cyclic chamber orientation",
			"electron-zero wall delta=135 degrees",
			"epsilon(H_e)=135 degrees-delta",
		},
		AlgebraicOverTrace: s.Typed && !s.TraceRingOrdersRoots && !m.NativeBranchSelector && r.PositiveRealBranchUnique && !r.FourthRootNative && !c.NativeChamberSelector,
		Native:             false,
		Environmental:      true,
		Verdict:            StatusRootBranchChamberSealDefined,
	}
}

func updateBFlavStatus(seal MinimalBranchSeal) BFlavStatus {
	return BFlavStatus{
		Expression:                        "B_flav=1-8*pi*epsilon_branch(R_e)-(1/4)Tr(P_eP_3^nu)+J(H_u,H_d)",
		ChargedLeptonSideTraceRing:        true,
		ChargedLeptonSideSplittingField:   true,
		ChargedLeptonSideFourthRootBranch: true,
		ChargedLeptonSideChamberSeal:      true,
		ChargedLeptonSideNative:           false,
		EnvironmentalOnly:                 true,
		Decision:                          "B_flav is refined from observed epsilon(H_e) to epsilon_branch(R_e): native trace-ring data plus explicit algebraic branch/chamber seals. This improves typing but does not supply a native zero theorem.",
		Verdict:                           StatusBFlavBranchAnchoredEnvironmental,
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{
		DerivesKoide:               false,
		DerivesChargedLeptonMasses: false,
		DerivesPMNSCKMNeutrino:     false,
		PromotesHEOneFourthNative:  false,
		PromotesChamberNative:      false,
		PromotesBFlavZero:          false,
		AddsCarrier:                false,
		AddsSelector:               false,
		SearchesNewConstants:       false,
		PreservesGate352:           true,
		PreservesGate596:           true,
		PreservesGate599:           true,
		Verdict:                    StatusGate600Boundary,
	}
}

func compileFinal(s TraceRingToSplittingFieldTable, m DiscriminantAndMonodromyAudit, r FourthRootBranchAudit, c ChamberOrderingAudit, seal MinimalBranchSeal, b BFlavStatus, f FirewallAudit) FinalVerdict {
	return FinalVerdict{
		SplittingFieldTyped:      s.Typed,
		TraceRingOrdersSpectrum:  s.TraceRingOrdersRoots,
		NativeEigenvalueBranch:   m.NativeBranchSelector && m.NativeOrdering,
		NativePositiveFourthRoot: r.FourthRootNative && r.PositivityNative,
		NativeChamberSelector:    c.NativeChamberSelector,
		EpsilonBranchAlgebraic:   seal.AlgebraicOverTrace,
		BFlavNative:              !b.EnvironmentalOnly,
		MinimalSeal:              seal.Name,
		Decision:                 "Outcome: epsilon(H_e) is branch-algebraic over R_e after adjoining a cubic splitting branch, positive fourth roots, and the charged-lepton chamber/wall seal. The trace ring does not select the observed ordering or wall; B_flav remains environmental.",
		Verdict:                  StatusEpsilonBranchAlgebraic,
	}
}

func Statuses() []string {
	return []string{
		StatusGate599Inherited,
		StatusCubicSplittingFieldTyped,
		StatusDiscriminantDefined,
		StatusTraceRingUnorderedSpectrum,
		StatusGenericS3MonodromyTyped,
		StatusFourthRootBranchTyped,
		StatusPositiveObservedBranchTyped,
		StatusChamberOrderingTyped,
		StatusEpsilonBranchAlgebraic,
		StatusRootBranchChamberSealDefined,
		StatusBFlavBranchAnchoredEnvironmental,
		StatusTraceRingNoOrdering,
		StatusNoNativeBranchOrderTheorem,
		StatusNoNativePositiveFourthRoot,
		StatusNoNativeElectronWall,
		StatusNoNativeCyclicOrder,
		StatusNoNativeEpsilonBranchTheorem,
		StatusBFlavEnvironmental,
		StatusNoNativeBFlavZero,
		StatusGate352Preserved,
		StatusGate596Preserved,
		StatusGate599Preserved,
		StatusNoKoideDerivation,
		StatusNoChargedLeptonMassDerivation,
		StatusNoPMNSCKMNeutrinoFlavorDeriv,
		StatusNoHEOneFourthNativePromotion,
		StatusNoChamberNativePromotion,
		StatusNoBFlavNativePromotion,
		StatusNoNewNumericalFit,
		StatusNoNewCarrierSelector,
		StatusGate600Boundary,
	}
}

func HasStatus(statuses []string, want string) bool {
	for _, got := range statuses {
		if got == want || strings.Contains(got, want) {
			return true
		}
	}
	return false
}
