// Package generation2minimalflavorhistorybranchsealclosureaudit implements
// Gate 604: Minimal Flavor History Branch Seal Closure Audit.
//
// Gate 603 showed that B_flav selects the electron wall, the third neutrino
// projector, and the positive CKM orientation sign while treating the remaining
// charged-lepton sigma/cyclic Fourier presentation as invisible. Gate 604 closes
// that branch analysis by separating native trace data, algebraic extensions,
// environmental branch seals, and gauge/convention data.
package generation2minimalflavorhistorybranchsealclosureaudit

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2chargedleptonsigmadegeneracygaugeorientationaudit"
)

const (
	AuditID = "GATE604-MINIMAL-FLAVOR-HISTORY-BRANCH-SEAL-CLOSURE-AUDIT"

	StatusGate603Inherited                      = "PASS_GATE603_SIGMA_GAUGE_ORIENTATION_RESULT_INHERITED"
	StatusFlavorHistoryBranchStackConstructed   = "PASS_FLAVOR_HISTORY_BRANCH_STACK_CONSTRUCTED"
	StatusNativeVsExtensionSealGaugeClassified  = "PASS_NATIVE_EXTENSION_SEAL_GAUGE_CLASSIFICATION_COMPLETE"
	StatusMinimalFlavorHistoryBranchSealDefined = "PASS_MINIMAL_FLAVOR_HISTORY_BRANCH_SEAL_DEFINED"
	StatusSigmaGaugeLikeForBFlav                = "PASS_SIGMA_CLASSIFIED_AS_GAUGE_LIKE_FOR_B_FLAV"
	StatusOptionalDiscriminantSeal              = "CONDITIONAL_SUPPORT_OPTIONAL_DISCRIMINANT_ORIENTATION_SEAL_FOR_FULL_ORDER"
	StatusBFlavBranchCompatibilityFilter        = "CONDITIONAL_SUPPORT_B_FLAV_ACTS_AS_ENVIRONMENTAL_BRANCH_COMPATIBILITY_FILTER"
	StatusFlavorTransportUpdated                = "PASS_HISTORY_TRANSPORT_FLAVOR_FORMULA_UPDATED"
	StatusNoNativeBranchSelectionTheorem        = "FAILED_ROUTE_NO_NATIVE_BRANCH_SELECTION_THEOREM"
	StatusNoNativeBFlavZeroTheorem              = "FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM"
	StatusNoNativeFourthRootTheorem             = "FAILED_ROUTE_NO_NATIVE_FOURTH_ROOT_THEOREM"
	StatusNoNativeKoidePMNSCKMFlavorDerivation  = "FAILED_ROUTE_NO_NATIVE_KOIDE_PMNS_CKM_FLAVOR_DERIVATION"
	StatusNoNewCarrierSelector                  = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate352Preserved                      = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate596Preserved                      = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING"
	StatusGate599Preserved                      = "FIREWALL_PRESERVED_GATE599_TRACE_RING_EXTENSION_BOUNDARY_REMAINS_BINDING"
	StatusGate603Preserved                      = "FIREWALL_PRESERVED_GATE603_SIGMA_GAUGE_ORIENTATION_BOUNDARY_REMAINS_BINDING"
	StatusNoKoideDerivation                     = "FIREWALL_PRESERVED_NO_KOIDE_DERIVATION"
	StatusNoMassDerivation                      = "FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_OR_YUKAWA_DERIVATION"
	StatusNoPMNSCKMNeutrinoDerivation           = "FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_TEXTURE_DERIVATION"
	StatusNoBFlavNativePromotion                = "FIREWALL_PRESERVED_NO_B_FLAV_ZERO_NATIVE_PROMOTION"
	StatusGate604Boundary                       = "FIREWALL_PRESERVED_GATE604_MINIMAL_FLAVOR_HISTORY_BRANCH_SEAL_BOUNDARY"
)

const (
	epsilonE    = 0.039569756309433
	kappaE      = 0.00550355419157456
	sin2Theta13 = 0.02215
	jCKM        = 3.1169935287554706e-05
	bFlav       = -2.77587313788925e-06
)

type InheritedGate603 struct {
	SelectsElectronWall            bool
	SelectsP3Nu                    bool
	SelectsPositiveJ               bool
	SigmaGaugeForBFlav             bool
	OptionalSignedDiscriminantSeal bool
	BestResidual                   float64
	NextDistinctResidual           float64
	Verdict                        string
}

type BranchStackRow struct {
	Layer          string
	Item           string
	Role           string
	Classification string
	NeededForBFlav bool
	Native         bool
	Verdict        string
}

type ClassificationRow struct {
	Item                    string
	Native                  bool
	AlgebraicExtension      bool
	EnvironmentalBranchSeal bool
	GaugeConvention         bool
	ObservedLedger          bool
	NeededForBFlav          bool
	Explanation             string
	Verdict                 string
}

type MinimalityRow struct {
	Item                          string
	RequiredForBFlav              bool
	RequiredForFullOrderedHistory bool
	Reason                        string
	Verdict                       string
}

type MinimalFlavorHistoryBranchSeal struct {
	Name            string
	Components      []string
	SelectedByBFlav []string
	NotIncluded     []string
	IsNative        bool
	IsEnvironmental bool
	Verdict         string
}

type OptionalFullOrderSeal struct {
	Name                 string
	RequiredForBFlav     bool
	RequiredForFullOrder bool
	Data                 []string
	NativeTheoremPresent bool
	Statement            string
	Verdict              string
}

type UpdatedHistoryTransportFlavorFormula struct {
	Formula            string
	YCore              []string
	OmegaCore          []string
	TCore              []string
	RemainingRawInputs []string
	Verdict            string
}

type Firewalls struct {
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesBFlavZero           bool
	AddsCarrier                bool
	AddsSelector               bool
	PreservesGate352           bool
	PreservesGate596           bool
	PreservesGate599           bool
	PreservesGate603           bool
	Verdict                    string
}

type Analysis struct {
	Inherited         InheritedGate603
	BranchStack       []BranchStackRow
	Classification    []ClassificationRow
	Minimality        []MinimalityRow
	MinimalSeal       MinimalFlavorHistoryBranchSeal
	OptionalFullOrder OptionalFullOrderSeal
	Formula           UpdatedHistoryTransportFlavorFormula
	Firewalls         Firewalls
	Truth             string
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
	g603, err := generation2chargedleptonsigmadegeneracygaugeorientationaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate603 predecessor: %w", err)
	}
	inherited := inherit(g603)
	stack := buildBranchStack()
	classification := classifyItems()
	minimality := auditMinimality()
	seal := defineMinimalSeal()
	opt := defineOptionalFullOrderSeal()
	formula := updateFlavorFormula()
	firewalls := auditFirewalls()
	truth := "Gate 604 closes the flavor branch analysis by identifying the minimal environmental flavor-history branch seal.  B_flav requires the positive fourth-root charged-lepton branch through the electron-wall coordinate, the PMNS electron-to-third-neutrino projector overlap, and positive CKM orientation.  It does not require the full charged-lepton sigma/cyclic Fourier presentation, which remains gauge-like for B_flav; a signed Vandermonde orientation seal is optional only if full ordered-history reconstruction is demanded."
	return Analysis{inherited, stack, classification, minimality, seal, opt, formula, firewalls, truth}, nil
}

func inherit(a generation2chargedleptonsigmadegeneracygaugeorientationaudit.Analysis) InheritedGate603 {
	return InheritedGate603{
		SelectsElectronWall:            a.Inherited.SelectsElectronRow,
		SelectsP3Nu:                    a.Inherited.SelectsP3,
		SelectsPositiveJ:               a.Inherited.SelectsPositiveJ,
		SigmaGaugeForBFlav:             a.MinimalRemaining.SigmaGaugeForBFlav,
		OptionalSignedDiscriminantSeal: a.MinimalRemaining.PhysicalFullOrderingRequiresSeal,
		BestResidual:                   a.Inherited.BestResidual,
		NextDistinctResidual:           a.Inherited.NextDistinctResidual,
		Verdict:                        StatusGate603Inherited,
	}
}

func buildBranchStack() []BranchStackRow {
	return []BranchStackRow{
		{"native layer", "R_e=Q[Tr(H_e),Tr(H_e^2),Tr(H_e^3)]", "symmetric charged-lepton trace ring", "native", true, true, StatusFlavorHistoryBranchStackConstructed},
		{"native layer", "chi_e(lambda)", "characteristic polynomial determined by trace ring", "native", true, true, StatusFlavorHistoryBranchStackConstructed},
		{"algebraic-extension layer", "cubic splitting field K_e", "unordered eigenvalue roots lambda_i", "algebraic extension", true, false, StatusFlavorHistoryBranchStackConstructed},
		{"algebraic-extension layer", "positive fourth-root sheets x_i=lambda_i^(1/4)", "root coordinates needed by epsilon(H_e)", "sealed algebraic extension", true, false, StatusNoNativeFourthRootTheorem},
		{"environmental branch layer", "electron wall / epsilon_e", "selected wall coordinate used by B_flav", "environmental branch seal", true, false, StatusMinimalFlavorHistoryBranchSealDefined},
		{"environmental branch layer", "P_3^nu with electron row", "PMNS projector overlap |U_e3|^2/4", "observed ledger + branch seal", true, false, StatusMinimalFlavorHistoryBranchSealDefined},
		{"environmental branch layer", "+J_CKM", "positive quark CP orientation area", "observed ledger + orientation seal", true, false, StatusMinimalFlavorHistoryBranchSealDefined},
		{"gauge/convention layer", "sigma in S_3", "Fourier cyclic presentation of charged-lepton chamber", "gauge-like for B_flav", false, false, StatusSigmaGaugeLikeForBFlav},
		{"gauge/convention layer", "PMNS/CKM phase conventions", "matrix-representation conventions preserving projector/J invariants", "convention", false, false, StatusNativeVsExtensionSealGaugeClassified},
	}
}

func classifyItems() []ClassificationRow {
	return []ClassificationRow{
		{"charged-lepton trace ring R_e", true, false, false, false, false, true, "native polynomial trace data", StatusNativeVsExtensionSealGaugeClassified},
		{"characteristic polynomial chi_e", true, false, false, false, false, true, "constructed from Newton identities over R_e", StatusNativeVsExtensionSealGaugeClassified},
		{"cubic splitting field", false, true, false, false, false, true, "adjoins eigenvalue branches not ordered by trace ring", StatusNativeVsExtensionSealGaugeClassified},
		{"positive fourth-root branch", false, true, true, false, true, true, "needed for x_i=lambda_i^(1/4); not native by Gate 596", StatusNoNativeFourthRootTheorem},
		{"electron wall epsilon_e", false, false, true, false, true, true, "selected by Gate 602/603 as the wall relevant to B_flav", StatusMinimalFlavorHistoryBranchSealDefined},
		{"third neutrino projector P_3^nu", false, false, true, false, true, true, "selected projector in the PMNS overlap", StatusMinimalFlavorHistoryBranchSealDefined},
		{"positive CKM orientation", false, false, true, false, true, true, "selected sign of J_CKM", StatusMinimalFlavorHistoryBranchSealDefined},
		{"sixfold sigma/cyclic Fourier presentation", false, false, false, true, false, false, "does not affect unsigned electron-wall B_flav", StatusSigmaGaugeLikeForBFlav},
		{"signed Vandermonde orientation", false, false, true, false, false, false, "optional for full cyclic/order history, not needed by B_flav", StatusOptionalDiscriminantSeal},
	}
}

func auditMinimality() []MinimalityRow {
	return []MinimalityRow{
		{"positive fourth-root sheet", true, true, "epsilon(H_e) cannot be formed from H_e traces without x_i=lambda_i^(1/4)", StatusNoNativeFourthRootTheorem},
		{"electron-wall coordinate epsilon_e", true, true, "the scalar wall distance entering 1-8*pi*epsilon_e", StatusMinimalFlavorHistoryBranchSealDefined},
		{"PMNS electron-to-third-neutrino projector overlap", true, true, "Gate 602 selected alpha=e and i=3", StatusMinimalFlavorHistoryBranchSealDefined},
		{"positive CKM orientation +J_CKM", true, true, "Gate 602 selected s_J=+1", StatusMinimalFlavorHistoryBranchSealDefined},
		{"full charged-lepton sigma/cyclic order", false, true, "B_flav uses unsigned wall distance and does not see sigma", StatusSigmaGaugeLikeForBFlav},
		{"signed Vandermonde orientation", false, true, "needed only if full ordered-history branch is requested", StatusOptionalDiscriminantSeal},
		{"exact Fourier presentation choice", false, false, "coordinate convention for the same selected wall distance", StatusSigmaGaugeLikeForBFlav},
	}
}

func defineMinimalSeal() MinimalFlavorHistoryBranchSeal {
	return MinimalFlavorHistoryBranchSeal{
		Name:            "MinimalFlavorHistoryBranchSeal",
		Components:      []string{"positive fourth-root charged-lepton branch", "electron-wall coordinate epsilon_e", "PMNS projector overlap Tr(P_eP_3^nu)=|U_e3|^2", "positive CKM orientation +J_CKM", "OrientationBalanceSeal B_flav≈0"},
		SelectedByBFlav: []string{"electron wall / alpha=e", "third neutrino projector P_3^nu", "positive CKM sign s_J=+1"},
		NotIncluded:     []string{"full sigma/cyclic charged-lepton Fourier presentation", "signed Vandermonde orientation", "exact coordinate presentation of the Fourier chamber"},
		IsNative:        false,
		IsEnvironmental: true,
		Verdict:         strings.Join([]string{StatusMinimalFlavorHistoryBranchSealDefined, StatusBFlavBranchCompatibilityFilter, StatusNoNativeBranchSelectionTheorem}, ";"),
	}
}

func defineOptionalFullOrderSeal() OptionalFullOrderSeal {
	return OptionalFullOrderSeal{
		Name:                 "ChargedLeptonDiscriminantOrientationSeal",
		RequiredForBFlav:     false,
		RequiredForFullOrder: true,
		Data:                 []string{"sign(V_e)=sign(prod_{i<j}(lambda_j-lambda_i))", "or sign(V_x)=sign(prod_{i<j}(x_j-x_i))", "full cyclic order of (e,mu,tau)"},
		NativeTheoremPresent: false,
		Statement:            "Optional for B_flav, required only if the full charged-lepton cyclic/order branch is treated as physical rather than Fourier-coordinate gauge.",
		Verdict:              strings.Join([]string{StatusOptionalDiscriminantSeal, StatusNoNativeBranchSelectionTheorem}, ";"),
	}
}

func updateFlavorFormula() UpdatedHistoryTransportFlavorFormula {
	return UpdatedHistoryTransportFlavorFormula{
		Formula:            "E_flavor(M_Z)=T_flavor[native trace ring R_e, MinimalFlavorHistoryBranchSeal, OrientationBalanceSeal, remaining raw Yukawa/PMNS/CKM inputs]",
		YCore:              []string{"R_e and chi_e(lambda)", "positive fourth-root charged-lepton branch", "epsilon_e", "Tr(P_eP_3^nu)", "J(H_u,H_d)", "remaining Yukawa singular values as observed ledgers"},
		OmegaCore:          []string{"electron wall alpha=e", "third neutrino projector P_3^nu", "positive CKM orientation sign", "PMNS/CKM convention labels", "optional signed Vandermonde orientation only for full order"},
		TCore:              []string{"bridge-layer flavor transport", "no native derivation of epsilon_e, PMNS, CKM, Yukawa eigenvalues, or B_flav=0", "sigma treated as gauge-like for B_flav"},
		RemainingRawInputs: []string{"Yukawa singular values", "PMNS central/uncertainty ledger", "CKM/Jarlskog ledger", "neutrino ordering and convention data"},
		Verdict:            StatusFlavorTransportUpdated,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, false, true, true, true, true, StatusGate604Boundary}
}

func Statuses() []string {
	return []string{StatusGate603Inherited, StatusFlavorHistoryBranchStackConstructed, StatusNativeVsExtensionSealGaugeClassified, StatusMinimalFlavorHistoryBranchSealDefined, StatusSigmaGaugeLikeForBFlav, StatusOptionalDiscriminantSeal, StatusBFlavBranchCompatibilityFilter, StatusFlavorTransportUpdated, StatusNoNativeBranchSelectionTheorem, StatusNoNativeBFlavZeroTheorem, StatusNoNativeFourthRootTheorem, StatusNoNativeKoidePMNSCKMFlavorDerivation, StatusNoNewCarrierSelector, StatusGate352Preserved, StatusGate596Preserved, StatusGate599Preserved, StatusGate603Preserved, StatusNoKoideDerivation, StatusNoMassDerivation, StatusNoPMNSCKMNeutrinoDerivation, StatusNoBFlavNativePromotion, StatusGate604Boundary}
}
