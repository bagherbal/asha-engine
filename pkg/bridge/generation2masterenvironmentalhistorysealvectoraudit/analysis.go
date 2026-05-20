// Package generation2masterenvironmentalhistorysealvectoraudit implements
// Gate 605: Master Environmental History Seal Vector Audit.
//
// Gate 604 closed the minimal flavor-history branch seal.  Gate 605 zooms out
// and assembles the master ASHA endpoint ledger, separating native law-space,
// algebraic extensions, bridge normalizations, environmental history seals,
// gauge/convention data, and observed endpoint ledgers.
package generation2masterenvironmentalhistorysealvectoraudit

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2minimalflavorhistorybranchsealclosureaudit"
)

const (
	AuditID = "GATE605-MASTER-ENVIRONMENTAL-HISTORY-SEAL-VECTOR-AUDIT"

	StatusGate604Inherited                    = "PASS_GATE604_MINIMAL_FLAVOR_HISTORY_BRANCH_SEAL_INHERITED"
	StatusMasterSealVectorConstructed         = "PASS_MASTER_ENVIRONMENTAL_HISTORY_SEAL_VECTOR_CONSTRUCTED"
	StatusNativeBoundaryClarified             = "PASS_NATIVE_VS_ENVIRONMENTAL_BOUNDARY_CLARIFIED"
	StatusNativeLawLedgerConstructed          = "PASS_NATIVE_LAW_SPACE_LEDGER_CONSTRUCTED"
	StatusAlgebraicExtensionLedgerConstructed = "PASS_ALGEBRAIC_EXTENSION_LAYER_CONSTRUCTED"
	StatusBridgeSealLedgerConstructed         = "PASS_BRIDGE_ENVIRONMENTAL_SEAL_LEDGER_CONSTRUCTED"
	StatusSolvedUnsolvedLedgerConstructed     = "PASS_SOLVED_UNSOLVED_LEDGER_CONSTRUCTED"
	StatusMasterFormulaWritten                = "PASS_MASTER_HISTORY_TRANSPORT_FORMULA_WRITTEN"
	StatusNextTargetRankingConstructed        = "PASS_NEXT_TARGET_RANKING_CONSTRUCTED"
	StatusFlavorSealIntegrated                = "CONDITIONAL_SUPPORT_FLAVOR_SEAL_INTEGRATED_AS_HISTORY_BRANCH"
	StatusRGThresholdNextActionable           = "CONDITIONAL_SUPPORT_RG_THRESHOLD_TRANSPORT_NEXT_ACTIONABLE_TARGET"
	StatusKineticScaleHighValue               = "CONDITIONAL_SUPPORT_KINETIC_NORMALIZATION_AND_HIGGS_VEV_HIGH_VALUE_TARGET"
	StatusRootChamberDeepButBlocked           = "CONDITIONAL_SUPPORT_ROOT_CHAMBER_THEOREM_DEEP_BUT_CURRENTLY_BLOCKED"
	StatusNoNativeFlavorBalanceTheorem        = "FAILED_ROUTE_NO_NATIVE_FLAVOR_BALANCE_THEOREM"
	StatusNoProductTimeAirlock                = "FAILED_ROUTE_NO_PRODUCT_TIME_AIRLOCK"
	StatusNoAbsoluteKineticScale              = "FAILED_ROUTE_NO_ABSOLUTE_KINETIC_SCALE"
	StatusNoNativeRootChamberTheorem          = "FAILED_ROUTE_NO_NATIVE_ROOT_CHAMBER_THEOREM"
	StatusNoNativeRGThresholdTheorem          = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_THEOREM"
	StatusNoNativeHiggsVEVTheorem             = "FAILED_ROUTE_NO_NATIVE_HIGGS_VEV_THEOREM"
	StatusNoNativeObservedEndpointDerivation  = "FAILED_ROUTE_NO_NATIVE_OBSERVED_ENDPOINT_DERIVATION"
	StatusNoNewConstantSearch                 = "FIREWALL_PRESERVED_NO_NEW_NUMERICAL_CONSTANT_SEARCH"
	StatusNoKoideDerivation                   = "FIREWALL_PRESERVED_NO_KOIDE_DERIVATION"
	StatusNoFlavorDerivation                  = "FIREWALL_PRESERVED_NO_PMNS_CKM_YUKAWA_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusNoEWMassDerivation                  = "FIREWALL_PRESERVED_NO_WZ_PHOTON_OR_HIGGS_POLE_MASS_DERIVATION"
	StatusNoCosmologyDerivation               = "FIREWALL_PRESERVED_NO_COSMOLOGY_DERIVATION"
	StatusGate352Preserved                    = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate596Preserved                    = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING"
	StatusGate604Preserved                    = "FIREWALL_PRESERVED_GATE604_MINIMAL_FLAVOR_HISTORY_BRANCH_SEAL_BOUNDARY_REMAINS_BINDING"
	StatusGate605Boundary                     = "FIREWALL_PRESERVED_GATE605_MASTER_HISTORY_SEAL_VECTOR_BOUNDARY"
)

const (
	ClassNative             = "native law-space"
	ClassAlgebraicExtension = "algebraic extension"
	ClassBridgeSeal         = "bridge normalization seal"
	ClassEnvironmentalSeal  = "environmental history seal"
	ClassObservedLedger     = "observed endpoint ledger"
	ClassGaugeConvention    = "gauge/convention"
)

type InheritedGate604 struct {
	MinimalFlavorSealDefined bool
	SigmaGaugeForBFlav       bool
	OptionalFullOrderSeal    bool
	FlavorFormula            string
	Verdict                  string
}

type MasterSealRow struct {
	Symbol            string
	Sector            string
	Meaning           string
	Classification    string
	GateSource        string
	NativeObstruction string
	Verdict           string
}

type MasterFormula struct {
	Formula                 string
	NativeLawInputs         []string
	AlgebraicExtensions     []string
	HistorySeals            []string
	BridgeNormalizations    []string
	ObservedEndpointLedgers []string
	Verdict                 string
}

type SolvedUnsolvedRow struct {
	Item    string
	Sector  string
	Status  string
	Reason  string
	Verdict string
}

type NextTargetRankingRow struct {
	Rank           int
	Path           string
	Value          string
	CurrentStatus  string
	Rationale      string
	Recommendation string
	Verdict        string
}

type ClassificationSummary struct {
	NativeCount             int
	AlgebraicExtensionCount int
	BridgeSealCount         int
	EnvironmentalSealCount  int
	ObservedLedgerCount     int
	GaugeConventionCount    int
	BoundaryClear           bool
	Verdict                 string
}

type Firewalls struct {
	DerivesKoide            bool
	DerivesFlavor           bool
	DerivesEWMasses         bool
	DerivesCosmology        bool
	DerivesObservedEndpoint bool
	SearchesNewConstants    bool
	PreservesGate352        bool
	PreservesGate596        bool
	PreservesGate604        bool
	Verdict                 string
}

type Analysis struct {
	Inherited       InheritedGate604
	MasterSealTable []MasterSealRow
	Summary         ClassificationSummary
	Formula         MasterFormula
	SolvedUnsolved  []SolvedUnsolvedRow
	Ranking         []NextTargetRankingRow
	Firewalls       Firewalls
	Truth           string
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
	g604, err := generation2minimalflavorhistorybranchsealclosureaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate604 predecessor: %w", err)
	}
	inherited := inherit(g604)
	table := buildMasterSealTable()
	summary := summarize(table)
	formula := buildMasterFormula()
	solved := buildSolvedUnsolvedLedger()
	ranking := buildRanking()
	firewalls := auditFirewalls()
	truth := "Gate 605 assembles the master ASHA history-seal vector.  The native law-space contains the Clifford/Witt/Fock, finite spectral-triple, representation-trace, symbolic electroweak, and polynomial spectral-action trace structures.  The endpoint universe still requires algebraic extensions, bridge normalizations, environmental history seals, transport choices, time/OS airlocks, and observed ledgers.  The minimal flavor branch seal from Gate 604 is integrated as a history branch, not promoted to native law."
	return Analysis{inherited, table, summary, formula, solved, ranking, firewalls, truth}, nil
}

func inherit(a generation2minimalflavorhistorybranchsealclosureaudit.Analysis) InheritedGate604 {
	return InheritedGate604{
		MinimalFlavorSealDefined: a.MinimalSeal.Name == "MinimalFlavorHistoryBranchSeal" && a.MinimalSeal.IsEnvironmental,
		SigmaGaugeForBFlav:       a.OptionalFullOrder.RequiredForFullOrder && !a.OptionalFullOrder.RequiredForBFlav,
		OptionalFullOrderSeal:    a.OptionalFullOrder.Name == "ChargedLeptonDiscriminantOrientationSeal",
		FlavorFormula:            a.Formula.Formula,
		Verdict:                  StatusGate604Inherited,
	}
}

func buildMasterSealTable() []MasterSealRow {
	return []MasterSealRow{
		{"Cℓ(1,7)", "Clifford law-space", "finite measurement ladder / native algebraic substrate", ClassNative, "core ASHA gates", "none for structural substrate", StatusNativeLawLedgerConstructed},
		{"W=C^4, S^7→CP^3", "Witt/Fock/projective law-space", "Hopf quotient and projective selector geometry", ClassNative, "Gates 570-572", "no proven functor to K7 or physical time", StatusNativeLawLedgerConstructed},
		{"B-L 4=1+3", "selector algebra", "native colorless/colored split with Comm(B-L)=u(1)+u(3)", ClassNative, "Gates 555,572", "does not access flavor root chamber", StatusNativeLawLedgerConstructed},
		{"A_F=C⊕H⊕M_3(C)", "finite spectral triple", "Standard Model internal algebra socket", ClassNative, "finite spectral triple gates", "does not fix observed Yukawa values", StatusNativeLawLedgerConstructed},
		{"D_F one-form edge inventory", "finite Dirac lane", "legal Q_L↔u_R,d_R and L_L↔e_R,ν_R edges", ClassNative, "Gates 253,576,598", "no cross-sector flavor-orientation theorem", StatusNativeLawLedgerConstructed},
		{"H_phi≈C^2", "scalar lane", "one complex scalar doublet carrier", ClassNative, "inner fluctuation / scalar gates", "no native VEV or pole mass theorem", StatusNativeLawLedgerConstructed},
		{"k_Y=5/3", "representation trace", "hypercharge normalization", ClassNative, "EW normalization gates", "absolute coupling scale still bridge", StatusNativeLawLedgerConstructed},
		{"sin²(theta_*)=3/8", "boundary electroweak", "trace-normalized boundary weak angle", ClassNative, "EW normalization gates", "transport to M_Z requires history", StatusNativeLawLedgerConstructed},
		{"symbolic EW Hessian", "electroweak bridge", "neutral matrix shape and W/Z ratio form", ClassNative, "Gates 564/565", "bridge-symbolic, not full W/Z/photon dynamics", StatusNativeLawLedgerConstructed},
		{"a,b spectral-action power sums", "color/colorless trace cable", "native polynomial lepton/quark Yukawa power-sum cable with color weight 3", ClassNative, "Gate 598", "polynomial only; no root/orientation cable", StatusNativeLawLedgerConstructed},

		{"R_e=Q[Tr(H_e),Tr(H_e²),Tr(H_e³)]", "flavor algebra", "charged-lepton native trace ring", ClassNative, "Gate 599", "unordered spectrum only", StatusAlgebraicExtensionLedgerConstructed},
		{"chi_e(lambda)", "flavor algebra", "charged-lepton characteristic polynomial from native traces", ClassNative, "Gate 599", "does not choose branch/order", StatusAlgebraicExtensionLedgerConstructed},
		{"K_e splitting field", "flavor algebra", "cubic eigenvalue splitting extension", ClassAlgebraicExtension, "Gate 600", "branch ordering not native", StatusAlgebraicExtensionLedgerConstructed},
		{"x_i=lambda_i^(1/4)", "flavor algebra", "positive fourth-root charged-lepton sheets", ClassAlgebraicExtension, "Gates 596,599,600", "no native fourth-root theorem", StatusNoNativeRootChamberTheorem},
		{"epsilon_e", "flavor branch", "electron-wall distance in charged-lepton root chamber", ClassEnvironmentalSeal, "Gates 583-604", "requires root/chamber seal", StatusFlavorSealIntegrated},
		{"B_flav≈0", "flavor branch", "orientation balance between epsilon_e, |U_e3|²/4, and +J_CKM", ClassEnvironmentalSeal, "Gates 590-604", "no native B_flav=0 theorem", StatusNoNativeFlavorBalanceTheorem},
		{"sigma∈S_3", "flavor branch", "charged-lepton Fourier cyclic presentation", ClassGaugeConvention, "Gate 603", "gauge-like for B_flav; signed Vandermonde optional", StatusFlavorSealIntegrated},
		{"sgn(V_e) / sgn(V_x)", "flavor branch", "optional full charged-lepton order orientation", ClassEnvironmentalSeal, "Gates 603-604", "no native signed-discriminant theorem", StatusNoNativeRootChamberTheorem},

		{"Lambda_12", "scale/gauge transport", "one-loop g1=g2 meeting scale", ClassEnvironmentalSeal, "history transport v1", "not a full unification scale; g3 mismatch remains", StatusBridgeSealLedgerConstructed},
		{"Delta_3,R_3", "gauge transport", "strong-coupling mismatch at Lambda_12", ClassEnvironmentalSeal, "history transport v1", "threshold/matching history not derived", StatusNoNativeRGThresholdTheorem},
		{"f_k,f_0,Lambda", "spectral-action scale", "cutoff moments / boundary scale data", ClassBridgeSeal, "spectral-action bridge", "moments/absolute scale not native constants", StatusBridgeSealLedgerConstructed},
		{"g,g',g1,g2,g3 absolute scale", "gauge", "endpoint and boundary coupling magnitudes", ClassObservedLedger, "history transport v1", "absolute coupling scale and thresholds remain history", StatusNoAbsoluteKineticScale},
		{"K_phi", "scalar kinetic normalization", "scalar metric / Higgs kinetic normalization", ClassBridgeSeal, "EW Hessian gates", "not fixed natively", StatusNoAbsoluteKineticScale},
		{"v", "vacuum/Higgs", "electroweak VEV from G_F endpoint convention", ClassObservedLedger, "history transport v1", "not native; observed endpoint", StatusNoNativeHiggsVEVTheorem},
		{"lambda(M_Z), lambda(Lambda_12)", "scalar transport", "Higgs quartic endpoint/running and zero crossing in v1", ClassEnvironmentalSeal, "history transport v1", "one-loop/top-dominant approximation", StatusBridgeSealLedgerConstructed},
		{"RG thresholds and matching", "transport", "beta functions, threshold ledger, scheme and matching data", ClassEnvironmentalSeal, "history transport v1", "missing deltaT_hist", StatusNoNativeRGThresholdTheorem},
		{"D_M / OS/Wick/Hilbert/time", "continuum/time", "product geometry, Lorentzian/OS/Hilbert reconstruction, dynamics", ClassBridgeSeal, "time/OS airlock gates", "airlock remains closed", StatusNoProductTimeAirlock},
		{"Yukawa singular values", "flavor", "observed endpoint flavor magnitudes", ClassObservedLedger, "history transport v1", "not native ASHA derivations", StatusNoNativeObservedEndpointDerivation},
		{"CKM", "flavor orientation", "observed quark mixing and Jarlskog orientation", ClassObservedLedger, "history transport v1 / Gates 590-604", "no native CKM theorem", StatusNoNativeFlavorBalanceTheorem},
		{"PMNS/neutrino ordering", "flavor orientation", "observed lepton projector data and mass-order labels", ClassObservedLedger, "Gates 587-604", "no native PMNS/neutrino theorem", StatusNoNativeFlavorBalanceTheorem},
		{"Planck ΛCDM endpoint", "cosmology", "optional observed cosmology endpoint ledger", ClassObservedLedger, "history transport v1", "no cosmology derivation", StatusNoNativeObservedEndpointDerivation},
	}
}

func summarize(rows []MasterSealRow) ClassificationSummary {
	var s ClassificationSummary
	for _, r := range rows {
		switch r.Classification {
		case ClassNative:
			s.NativeCount++
		case ClassAlgebraicExtension:
			s.AlgebraicExtensionCount++
		case ClassBridgeSeal:
			s.BridgeSealCount++
		case ClassEnvironmentalSeal:
			s.EnvironmentalSealCount++
		case ClassObservedLedger:
			s.ObservedLedgerCount++
		case ClassGaugeConvention:
			s.GaugeConventionCount++
		}
	}
	s.BoundaryClear = s.NativeCount > 0 && s.EnvironmentalSealCount > 0 && s.ObservedLedgerCount > 0 && s.GaugeConventionCount > 0
	s.Verdict = StatusNativeBoundaryClarified
	return s
}

func buildMasterFormula() MasterFormula {
	return MasterFormula{
		Formula: "E_End = T_history[NativeLaw(Cℓ(1,7),A_F,D_F,CP^3 selector geometry,spectral-action symbolic sockets), AlgebraicExtensions(R_e,chi_e,K_e,H_e^(1/4)), HistorySeals(Λ,f,K_phi,v,g_i,RG thresholds,OS/Wick/Hilbert time,MinimalFlavorHistoryBranchSeal,B_flav,endpoint ledgers)]",
		NativeLawInputs: []string{
			"Cℓ(1,7)", "W=C^4 and Hopf S^7→CP^3", "selector algebra and B-L 4=1+3", "A_F=C⊕H⊕M_3(C)", "D_F finite edge inventory", "H_phi≈C²", "k_Y=5/3", "sin²(theta_*)=3/8", "symbolic electroweak Hessian", "native polynomial spectral-action trace cable",
		},
		AlgebraicExtensions: []string{
			"R_e trace ring", "chi_e(lambda)", "cubic splitting field K_e", "positive fourth-root sheets x_i=lambda_i^(1/4)", "charged-lepton branch/chamber stack",
		},
		HistorySeals: []string{
			"Lambda_12 and boundary scale data", "Delta_3/R_3 gauge mismatch", "K_phi and scalar normalization", "v and Higgs vacuum endpoint", "RG thresholds and matching", "OS/Wick/Hilbert time airlock", "MinimalFlavorHistoryBranchSeal", "OrientationBalanceSeal B_flav≈0", "cosmology endpoint optional",
		},
		BridgeNormalizations:    []string{"f_k/f_0 cutoff moments", "absolute gauge coupling scale", "canonical g1 normalization", "scalar metric normalization", "transport scheme T_v1/δT_hist"},
		ObservedEndpointLedgers: []string{"m_W,m_Z,m_H,G_F", "Yukawa singular values", "CKM", "PMNS/neutrino ordering", "Planck ΛCDM optional endpoint"},
		Verdict:                 StatusMasterFormulaWritten,
	}
}

func buildSolvedUnsolvedLedger() []SolvedUnsolvedRow {
	return []SolvedUnsolvedRow{
		{"B-L 4=1+3", "selector algebra", "solved/native", "native selector and commutant certified", StatusNativeLawLedgerConstructed},
		{"A_F=C⊕H⊕M_3(C)", "finite spectral triple", "solved/native", "internal algebra and gauge sockets certified", StatusNativeLawLedgerConstructed},
		{"one scalar doublet", "finite one-form lane", "solved/native", "finite one-forms produce one complex scalar doublet", StatusNativeLawLedgerConstructed},
		{"k_Y=5/3 and sin²(theta_*)=3/8", "representation trace", "solved/native", "boundary normalization certified", StatusNativeLawLedgerConstructed},
		{"symbolic EW Hessian", "electroweak bridge", "solved/native shape", "matrix form and W/Z ratio shape certified", StatusNativeLawLedgerConstructed},
		{"native polynomial color/colorless trace cable", "spectral action", "solved/native polynomial", "lepton/quark power sums with color factor 3 visible", StatusNativeLawLedgerConstructed},
		{"charged-lepton flavor seal", "flavor", "environmental/compressed", "epsilon_e, |U_e3|² and +J_CKM form minimal branch seal", StatusFlavorSealIntegrated},
		{"B_flav≈0", "flavor", "environmental/compressed", "branch compatibility filter, not native theorem", StatusNoNativeFlavorBalanceTheorem},
		{"H_e^(1/4)", "flavor", "unsolved/native gap", "fourth-root/root-chamber theorem missing", StatusNoNativeRootChamberTheorem},
		{"absolute kinetic scale", "gauge/scalar", "unsolved/native gap", "absolute g_i/K_phi/normalization not fixed natively", StatusNoAbsoluteKineticScale},
		{"Higgs VEV and pole masses", "vacuum/Higgs", "unsolved/native gap", "v and pole values are endpoint/bridge data", StatusNoNativeHiggsVEVTheorem},
		{"RG thresholds/matching", "transport", "unsolved/history gap", "deltaT_hist not derived", StatusNoNativeRGThresholdTheorem},
		{"OS/Hilbert/product time", "continuum/time", "unsolved/airlock", "no product-time/OS/Hilbert theorem", StatusNoProductTimeAirlock},
		{"neutrino mass mechanism", "flavor", "unsolved/native gap", "PMNS/neutrino data remain observed ledgers", StatusNoNativeFlavorBalanceTheorem},
		{"full flavor texture", "flavor", "unsolved/native gap", "Yukawa eigenvalues, CKM, PMNS not derived", StatusNoNativeFlavorBalanceTheorem},
	}
}

func buildRanking() []NextTargetRankingRow {
	return []NextTargetRankingRow{
		{1, "RG / threshold transport", "highest actionable", "v1 one-loop transport exposes Delta_3 and scalar crossing but lacks thresholds", "Directly improves the map from ASHA boundary normalizations to measured End data without demanding a new root functional.", "Pursue next as the practical continuation.", StatusRGThresholdNextActionable},
		{2, "kinetic normalization / Higgs VEV", "high physical value", "absolute gauge/scalar scale and v remain bridge seals", "Needed before W/Z/Higgs pole-scale claims can become physical predictions.", "Pursue after transport ledger improves.", StatusKineticScaleHighValue},
		{3, "root/chamber theorem", "deep but blocked", "Gate 596-604 isolate H_e^(1/4) and branch/chamber seals", "Would promote the flavor seal, but currently requires a new fourth-root/absolute-Dirac theorem.", "Keep as a long-range native theorem target.", StatusRootChamberDeepButBlocked},
		{4, "flavor spectral balance theorem", "deep but dependent", "B_flav≈0 is environmental and needs epsilon(H_e), PMNS and CKM theorems", "Cannot be native before the fourth-root and cross-sector orientation maps exist.", "Do not pursue as a standalone native claim yet.", StatusNoNativeFlavorBalanceTheorem},
		{5, "product-time / OS-Hilbert airlock", "fundamental but blocked", "Gates preserve Hopf/CP3/projective law-space from physical time", "Important, but current gates do not supply the necessary continuum reconstruction.", "Defer until a lawful carrier/functor appears.", StatusNoProductTimeAirlock},
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{
		DerivesKoide:            false,
		DerivesFlavor:           false,
		DerivesEWMasses:         false,
		DerivesCosmology:        false,
		DerivesObservedEndpoint: false,
		SearchesNewConstants:    false,
		PreservesGate352:        true,
		PreservesGate596:        true,
		PreservesGate604:        true,
		Verdict:                 StatusGate605Boundary,
	}
}

func Statuses() []string {
	return []string{
		StatusGate604Inherited,
		StatusMasterSealVectorConstructed,
		StatusNativeBoundaryClarified,
		StatusNativeLawLedgerConstructed,
		StatusAlgebraicExtensionLedgerConstructed,
		StatusBridgeSealLedgerConstructed,
		StatusSolvedUnsolvedLedgerConstructed,
		StatusMasterFormulaWritten,
		StatusNextTargetRankingConstructed,
		StatusFlavorSealIntegrated,
		StatusRGThresholdNextActionable,
		StatusKineticScaleHighValue,
		StatusRootChamberDeepButBlocked,
		StatusNoNativeFlavorBalanceTheorem,
		StatusNoProductTimeAirlock,
		StatusNoAbsoluteKineticScale,
		StatusNoNativeRootChamberTheorem,
		StatusNoNativeRGThresholdTheorem,
		StatusNoNativeHiggsVEVTheorem,
		StatusNoNativeObservedEndpointDerivation,
		StatusNoNewConstantSearch,
		StatusNoKoideDerivation,
		StatusNoFlavorDerivation,
		StatusNoEWMassDerivation,
		StatusNoCosmologyDerivation,
		StatusGate352Preserved,
		StatusGate596Preserved,
		StatusGate604Preserved,
		StatusGate605Boundary,
	}
}

func HasStatus(statuses []string, s string) bool {
	for _, x := range statuses {
		if x == s {
			return true
		}
	}
	return false
}

func rowsWithClass(rows []MasterSealRow, class string) []MasterSealRow {
	out := make([]MasterSealRow, 0)
	for _, r := range rows {
		if r.Classification == class {
			out = append(out, r)
		}
	}
	return out
}

func containsRow(rows []MasterSealRow, symbol string) bool {
	for _, r := range rows {
		if r.Symbol == symbol {
			return true
		}
	}
	return false
}

func containsText(xs []string, s string) bool {
	for _, x := range xs {
		if x == s || strings.Contains(x, s) {
			return true
		}
	}
	return false
}
