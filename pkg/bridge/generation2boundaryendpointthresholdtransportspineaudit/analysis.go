// Package generation2boundaryendpointthresholdtransportspineaudit implements
// Gate 606: Boundary-to-Endpoint RG Threshold Transport Spine Audit.
//
// Gate 605 assembled the master history-seal vector and identified RG / threshold
// transport as the next actionable history spine. Gate 606 classifies the native
// boundary data, endpoint ledgers, RG approximations, threshold correction slots,
// kinetic normalization blockers, flavor coupling, and time firewalls without
// claiming any endpoint observable as ASHA-native.
package generation2boundaryendpointthresholdtransportspineaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2masterenvironmentalhistorysealvectoraudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE606-BOUNDARY-ENDPOINT-RG-THRESHOLD-TRANSPORT-SPINE-AUDIT"

	StatusGate605Inherited                  = "PASS_GATE605_MASTER_HISTORY_VECTOR_INHERITED"
	StatusNativeBoundaryClassified          = "PASS_NATIVE_BOUNDARY_CONDITIONS_CLASSIFIED"
	StatusEndpointLedgerBuilt               = "PASS_ENDPOINT_OBSERVED_LEDGER_BUILT"
	StatusGaugeRGSlotsDefined               = "PASS_GAUGE_RG_TRANSPORT_SLOTS_DEFINED"
	StatusScalarRGSlotsDefined              = "PASS_SCALAR_RG_TRANSPORT_SLOTS_DEFINED"
	StatusThresholdLedgerDefined            = "PASS_THRESHOLD_CORRECTION_LEDGER_DEFINED"
	StatusKineticBlockersClassified         = "PASS_KINETIC_NORMALIZATION_BLOCKERS_CLASSIFIED"
	StatusFlavorRelationRecorded            = "PASS_GATE604_FLAVOR_SEALS_RECORDED_AS_ENVIRONMENTAL_RG_INPUTS"
	StatusUpdatedFormulaWritten             = "PASS_UPDATED_HISTORY_TRANSPORT_FORMULA_WRITTEN"
	StatusRGThresholdNextSpine              = "CONDITIONAL_SUPPORT_RG_THRESHOLD_TRANSPORT_IS_NEXT_ACTIONABLE_HISTORY_SPINE"
	StatusGaugeMismatchThresholdNeeded      = "CONDITIONAL_SUPPORT_GAUGE_STRONG_MISMATCH_REQUIRES_THRESHOLD_OR_HIGHER_LOOP_LEDGER"
	StatusScalarCrossingApproxVisible       = "CONDITIONAL_SUPPORT_SCALAR_ZERO_CROSSING_VISIBLE_IN_V1_TOP_DOMINANT_APPROXIMATION"
	StatusEndpointLedgerBridgeOnly          = "CONDITIONAL_SUPPORT_ENDPOINT_LEDGER_IS_BRIDGE_ONLY_NOT_NATIVE_DERIVATION"
	StatusNoNativeRGThresholdTheorem        = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_THEOREM"
	StatusNoAbsoluteKineticScale            = "FAILED_ROUTE_NO_ABSOLUTE_KINETIC_SCALE"
	StatusNoHiggsVEVDerivation              = "FAILED_ROUTE_NO_HIGGS_VEV_DERIVATION"
	StatusNoLowEnergyWZPhotonDynamics       = "FAILED_ROUTE_NO_LOW_ENERGY_WZ_PHOTON_DYNAMICS_DERIVED"
	StatusNoFullGaugeUnificationClaim       = "FAILED_ROUTE_NO_FULL_GAUGE_UNIFICATION_CLAIM_G1_G2_ONLY"
	StatusNoScalarStabilityFinalClaim       = "FAILED_ROUTE_NO_FINAL_SCALAR_STABILITY_CLAIM_FROM_V1_RUNNING"
	StatusNoObservedEndpointDerivation      = "FAILED_ROUTE_NO_OBSERVED_ENDPOINT_DERIVATION"
	StatusNoFlavorFeedbackAsNativeRG        = "FAILED_ROUTE_FLAVOR_BALANCE_NOT_NATIVE_RG_LAW"
	StatusRGScaleNotProductTime             = "FIREWALL_PRESERVED_RG_SCALE_NOT_PRODUCT_TIME"
	StatusNoEndpointDerivationFirewall      = "FIREWALL_PRESERVED_NO_OBSERVED_ENDPOINT_DERIVATION"
	StatusNoWZPhotonDerivationFirewall      = "FIREWALL_PRESERVED_NO_PHYSICAL_WZ_PHOTON_DYNAMICS_DERIVATION"
	StatusNoKoideFlavorPromotionFirewall    = "FIREWALL_PRESERVED_FLAVOR_SEALS_REMAIN_ENVIRONMENTAL"
	StatusThresholdsSchemesExplicitFirewall = "FIREWALL_PRESERVED_THRESHOLDS_AND_SCHEMES_EXPLICITLY_LABELED"
	StatusGate606Boundary                   = "FIREWALL_PRESERVED_GATE606_RG_THRESHOLD_TRANSPORT_SPINE_BOUNDARY"
)

const (
	ClassNative     = "native"
	ClassBridge     = "bridge seal"
	ClassEnv        = "environmental/history"
	ClassObserved   = "observed ledger"
	ClassBlocked    = "blocked native promotion"
	ClassConvention = "scheme/convention"
)

type InheritedGate605 struct {
	MasterVectorBuilt bool
	TopRecommendation string
	RGTopActionable   bool
	Verdict           string
}

type NativeBoundaryConditionRow struct {
	Symbol                 string
	Meaning                string
	Classification         string
	GateSource             string
	RequiredNormalizations []string
	PhysicalPromotionBlock string
	Verdict                string
}

type EndpointObservedLedgerRow struct {
	Symbol         string
	Value          float64
	Unit           string
	Scale          string
	Scheme         string
	Source         string
	Role           string
	Classification string
	Verdict        string
}

type GaugeTransportRow struct {
	Quantity       string
	BoundaryOrFlow string
	RuntimeValue   float64
	Formula        string
	Approximation  string
	Interpretation string
	Verdict        string
}

type ScalarTransportRow struct {
	Quantity       string
	RuntimeValue   float64
	Unit           string
	Formula        string
	Approximation  string
	Interpretation string
	Verdict        string
}

type ThresholdCorrectionSlotRow struct {
	Slot            string
	Sector          string
	Purpose         string
	CurrentStatus   string
	RuntimeResidual string
	Verdict         string
}

type KineticNormalizationBlockerRow struct {
	Blocker       string
	Sector        string
	WhyRequired   string
	CurrentStatus string
	Verdict       string
}

type FlavorSealRGRelationRow struct {
	Item            string
	RoleInTransport string
	NativeStatus    string
	Firewall        string
	Verdict         string
}

type ProductTimeFirewall struct {
	RGScaleIsProductTime bool
	RGScaleIsOSHilbert   bool
	RGScaleIsCosmoTime   bool
	Statement            string
	Verdict              string
}

type UpdatedFormula struct {
	Formula           string
	NativeBoundary    []string
	TransportSpine    []string
	ThresholdSlots    []string
	EndpointLedgers   []string
	BlockedPromotions []string
	Verdict           string
}

type Firewalls struct {
	ClaimsFullUnification bool
	DerivesEndpoint       bool
	DerivesKineticScale   bool
	DerivesVEV            bool
	DerivesFlavor         bool
	DerivesProductTime    bool
	ThresholdsExplicit    bool
	Verdict               string
}

type Analysis struct {
	Inherited           InheritedGate605
	NativeBoundaryTable []NativeBoundaryConditionRow
	EndpointLedger      []EndpointObservedLedgerRow
	GaugeTransport      []GaugeTransportRow
	ScalarTransport     []ScalarTransportRow
	ThresholdSlots      []ThresholdCorrectionSlotRow
	KineticBlockers     []KineticNormalizationBlockerRow
	FlavorRelation      []FlavorSealRGRelationRow
	ProductTimeFirewall ProductTimeFirewall
	Formula             UpdatedFormula
	Firewalls           Firewalls
	Truth               string
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
	g605, err := generation2masterenvironmentalhistorysealvectoraudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate605 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport bundle: %w", err)
	}
	inherited := inherit(g605)
	boundary := buildNativeBoundaryTable()
	endpoint := buildEndpointLedger(bundle)
	gauge := buildGaugeTransport(bundle)
	scalar := buildScalarTransport(bundle)
	thresholds := buildThresholdSlots(bundle)
	blockers := buildKineticBlockers()
	flavor := buildFlavorRelation()
	pt := buildProductTimeFirewall()
	formula := buildUpdatedFormula()
	firewalls := auditFirewalls()
	truth := "Gate 606 audits the boundary-to-endpoint transport spine. ASHA supplies native boundary ratios and symbolic sockets; the current runtime supplies one-loop endpoint transport, visible gauge/scalar residuals, and explicit threshold/matching gaps. The gate defines correction slots without fitting them, preserves the kinetic/VEV and product-time firewalls, and treats flavor seals as environmental inputs rather than RG law."
	return Analysis{inherited, boundary, endpoint, gauge, scalar, thresholds, blockers, flavor, pt, formula, firewalls, truth}, nil
}

func inherit(a generation2masterenvironmentalhistorysealvectoraudit.Analysis) InheritedGate605 {
	return InheritedGate605{
		MasterVectorBuilt: len(a.MasterSealTable) > 0 && a.Summary.BoundaryClear,
		TopRecommendation: a.Ranking[0].Path,
		RGTopActionable:   len(a.Ranking) > 0 && a.Ranking[0].Path == "RG / threshold transport",
		Verdict:           StatusGate605Inherited,
	}
}

func buildNativeBoundaryTable() []NativeBoundaryConditionRow {
	return []NativeBoundaryConditionRow{
		{"k_Y=5/3", "hypercharge trace normalization g1^2=(5/3)gY^2", ClassNative, "representation-trace normalization gates", []string{"canonical U(1) normalization", "finite representation trace"}, "absolute coupling scale and threshold history remain bridge data", StatusNativeBoundaryClassified},
		{"sin²(theta_*)=3/8", "boundary weak-angle normalization when g1=g2", ClassNative, "EW normalization gates", []string{"g1=g2 boundary convention", "g1^2=(5/3)gY^2"}, "transport to endpoint weak angle requires RG/history", StatusNativeBoundaryClassified},
		{"g'^2/g²=3/5", "scalar-Hessian abelian/weak ratio at canonical boundary", ClassBridge, "Gate 565 boundary Hessian ratio", []string{"canonical boundary", "Hessian coupling convention"}, "not an endpoint value; low-energy matching required", StatusNativeBoundaryClassified},
		{"m_W²/m_Z²=5/8", "symbolic boundary ratio shape from Hessian and trace normalization", ClassBridge, "Gate 565", []string{"K_phi", "v", "g,g' convention"}, "physical pole masses require K_phi, v, thresholds, and continuum matching", StatusNoLowEnergyWZPhotonDynamics},
		{"symbolic EW Hessian", "M_neutral²=(K_phi v²/4)[[g²,-gg'],[-gg',g'²]] and W/Z ratio socket", ClassNative, "Gates 564/565", []string{"K_phi", "v", "g,g' convention", "finite scalar doublet"}, "symbolic shape only; no endpoint W/Z/photon dynamics without transport and matching", StatusNativeBoundaryClassified},
		{"neutral null socket", "massless neutral direction in symbolic electroweak Hessian", ClassNative, "Gate 564/565", []string{"finite one-form scalar lane", "gauge socket representation"}, "not full photon dynamics without continuum/time/gauge-field reconstruction", StatusNoLowEnergyWZPhotonDynamics},
		{"A_F=C⊕H⊕M_3(C)", "finite algebra gauge sockets", ClassNative, "finite spectral triple gates", []string{"internal algebra representation"}, "does not fix endpoint couplings or thresholds", StatusNativeBoundaryClassified},
		{"H_phi≈C²", "scalar doublet socket from finite one-forms", ClassNative, "finite one-form/scalar gates", []string{"finite Dirac edge lane"}, "does not supply VEV or kinetic normalization", StatusNativeBoundaryClassified},
		{"a,b Yukawa power-sum trace cable", "native polynomial color/colorless trace coefficients with quark color factor 3", ClassNative, "Gate 598", []string{"finite spectral-action trace lane"}, "polynomial cable only; not root/orientation cable and not full flavor theorem", StatusNativeBoundaryClassified},
	}
}

func buildEndpointLedger(b historytransport.Bundle) []EndpointObservedLedgerRow {
	m := b.Inputs.Measured
	rows := []EndpointObservedLedgerRow{
		{"gY(M_Z)", b.EndVector.GY, "dimensionless", "M_Z", "extracted from m_W,m_Z,G_F in v1", "historytransport:end_vector", "scalar-Hessian abelian coupling at endpoint", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"g1(M_Z)", b.EndVector.G1, "dimensionless", "M_Z", "canonical g1=sqrt(5/3)gY", "historytransport:end_vector", "canonical hypercharge coupling at endpoint", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"g2(M_Z)", b.EndVector.G2, "dimensionless", "M_Z", "from 2m_W/v", "historytransport:end_vector", "weak coupling endpoint proxy", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"g3(M_Z)", b.EndVector.G3, "dimensionless", "M_Z", "from alpha_s(M_Z)", "historytransport:end_vector", "strong coupling endpoint proxy", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"sin²(theta_End)", b.EndVector.Sin2Theta, "dimensionless", "M_Z/on-shell v1", "1-m_W²/m_Z²", "historytransport:end_vector", "transported weak-angle endpoint", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"lambda(M_Z)", b.EndVector.Lambda, "dimensionless", "M_Z endpoint proxy", "m_H²/(2v²)", "historytransport:end_vector", "scalar quartic endpoint proxy", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"v", b.EndVector.VGeV, "GeV", "endpoint", "v=(sqrt(2)G_F)^(-1/2)", m["G_F"].SourceID, "VEV extracted from observed Fermi constant", ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"G_F", m["G_F"].Value, m["G_F"].Unit, m["G_F"].Scale, m["G_F"].Scheme, m["G_F"].SourceID, m["G_F"].Role, ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"m_W", m["m_W"].Value, m["m_W"].Unit, m["m_W"].Scale, m["m_W"].Scheme, m["m_W"].SourceID, m["m_W"].Role, ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"m_Z", m["m_Z"].Value, m["m_Z"].Unit, m["m_Z"].Scale, m["m_Z"].Scheme, m["m_Z"].SourceID, m["m_Z"].Role, ClassObserved, StatusEndpointLedgerBridgeOnly},
		{"m_H", m["m_H"].Value, m["m_H"].Unit, m["m_H"].Scale, m["m_H"].Scheme, m["m_H"].SourceID, m["m_H"].Role, ClassObserved, StatusEndpointLedgerBridgeOnly},
	}
	return rows
}

func buildGaugeTransport(b historytransport.Bundle) []GaugeTransportRow {
	return []GaugeTransportRow{
		{"beta_g_i", "flow law", math.NaN(), "dg_i/dlnmu=b_i g_i^3/(16*pi^2)", "one-loop SM v1", "slots defined; two-loop and threshold matching absent", StatusGaugeRGSlotsDefined},
		{"Lambda_12", "g1=g2 crossing", b.GaugeBoundary.Lambda12GeV, "ln(Lambda12/M_Z)=8pi^2(g1^-2-g2^-2)/(b1-b2)", "one-loop SM v1", "boundary normalization scale solved; not full unification", StatusGaugeRGSlotsDefined},
		{"log(Lambda_12/M_Z)", "transport interval", b.GaugeBoundary.LogLambda12Mu0, "t=ln(Lambda_12/M_Z)", "one-loop SM v1", "history interval for current transport spine", StatusGaugeRGSlotsDefined},
		{"g_star", "boundary value", b.GaugeBoundary.GStar, "g_star=g1(Lambda12)=g2(Lambda12)", "one-loop SM v1", "electroweak canonical crossing value", StatusGaugeRGSlotsDefined},
		{"g3(Lambda_12)", "strong mismatch", b.GaugeBoundary.G3Lambda, "run g3 with b3=-7", "one-loop SM v1", "does not meet g_star in v1", StatusGaugeMismatchThresholdNeeded},
		{"Delta_3", "strong inverse-coupling residual", b.GaugeBoundary.Delta3, "g3^-2(Lambda12)-g_star^-2", "one-loop SM v1", "threshold or higher-loop ledger needed", StatusGaugeMismatchThresholdNeeded},
		{"R_3", "strong ratio residual", b.GaugeBoundary.R3, "g3(Lambda12)/g_star", "one-loop SM v1", "strong coupling about 5% above g_star in v1", StatusGaugeMismatchThresholdNeeded},
		{"Delta_sin²", "weak-angle transport residual", b.WeakAngleTransport.DeltaSin2, "sin²(theta_End)-3/8", "endpoint/on-shell v1", "transport required from boundary angle to endpoint", StatusGaugeRGSlotsDefined},
	}
}

func buildScalarTransport(b historytransport.Bundle) []ScalarTransportRow {
	zero := math.NaN()
	if b.ScalarTransport.ZeroCrossingScaleGeV != nil {
		zero = *b.ScalarTransport.ZeroCrossingScaleGeV
	}
	return []ScalarTransportRow{
		{"lambda(M_Z)", b.ScalarTransport.LambdaMZ, "dimensionless", "m_H²/(2v²)", b.ScalarTransport.Approximation, "observed endpoint scalar quartic proxy", StatusScalarRGSlotsDefined},
		{"y_t(M_Z)", b.ScalarTransport.YT_MZ, "dimensionless", "sqrt(2)m_t/v in v1 ledger", b.ScalarTransport.Approximation, "dominant Yukawa contribution to scalar beta", StatusScalarRGSlotsDefined},
		{"beta_lambda(M_Z)", b.ScalarTransport.BetaLambdaMZ, "dimensionless per log scale", "one-loop beta_lambda top-dominant", b.ScalarTransport.Approximation, "negative at endpoint in v1", StatusScalarCrossingApproxVisible},
		{"lambda(Lambda_12)", b.ScalarTransport.LambdaLambda12, "dimensionless", "integrated v1 one-loop scalar transport", b.ScalarTransport.Approximation, "negative in v1; precision claim blocked", StatusScalarCrossingApproxVisible},
		{"y_t(Lambda_12)", b.ScalarTransport.YT_Lambda12, "dimensionless", "integrated v1 top Yukawa transport", b.ScalarTransport.Approximation, "transported top-dominant Yukawa", StatusScalarRGSlotsDefined},
		{"zero_crossing_scale", zero, "GeV", "lambda(mu)=0 if present", b.ScalarTransport.Approximation, "visible crossing in v1, threshold/scheme sensitive", StatusScalarCrossingApproxVisible},
	}
}

func buildThresholdSlots(b historytransport.Bundle) []ThresholdCorrectionSlotRow {
	return []ThresholdCorrectionSlotRow{
		{"delta_i^gauge", "gauge", "collect gauge threshold, matching, and higher-loop corrections to g1,g2,g3 transport", "open slot", fmt.Sprintf("Delta_3=%.15g, R_3=%.15g", b.GaugeBoundary.Delta3, b.GaugeBoundary.R3), StatusThresholdLedgerDefined},
		{"delta_lambda", "scalar", "correct scalar quartic matching and higher-loop vacuum-stability transport", "open slot", fmt.Sprintf("lambda(Lambda12)=%.15g", b.ScalarTransport.LambdaLambda12), StatusThresholdLedgerDefined},
		{"delta_yukawa", "flavor/scalar", "replace diagonal/top-dominant v1 Yukawa running with full matrix and threshold transport", "open slot", "v1 uses diagonal/top-dominant approximations", StatusThresholdLedgerDefined},
		{"delta_K_phi", "scalar kinetic", "match finite scalar kinetic metric normalization to continuum endpoint convention", "open slot", "K_phi not natively fixed", StatusThresholdLedgerDefined},
		{"delta_v", "vacuum", "match G_F-derived VEV / vacuum scale to finite scalar-vacuum data", "open slot", fmt.Sprintf("v=%.12g GeV endpoint ledger", b.EndVector.VGeV), StatusThresholdLedgerDefined},
		{"delta_pole_MSbar", "matching scheme", "convert pole masses/on-shell inputs and MSbar quantities consistently", "open slot", "v1 mixes explicitly labeled endpoint conventions", StatusThresholdLedgerDefined},
		{"delta_boundary", "boundary", "threshold corrections at Lambda_12 or spectral cutoff boundary", "open slot", "g1=g2 only; g3 mismatch remains", StatusThresholdLedgerDefined},
	}
}

func buildKineticBlockers() []KineticNormalizationBlockerRow {
	return []KineticNormalizationBlockerRow{
		{"K_phi", "scalar/Higgs", "sets scalar kinetic normalization entering W/Z symbolic Hessian", "bridge seal; no native value", StatusNoAbsoluteKineticScale},
		{"v", "vacuum", "needed for physical W/Z/Higgs masses", "G_F-derived observed endpoint ledger", StatusNoHiggsVEVDerivation},
		{"absolute g scale", "gauge", "needed beyond trace-normalized ratios for endpoint couplings", "observed/transported; not native", StatusNoAbsoluteKineticScale},
		{"f0 / cutoff moments", "spectral action", "set absolute gauge/gravity/scalar action normalizations", "bridge moments not native constants", StatusNoAbsoluteKineticScale},
		{"finite Yukawa trace a", "scalar normalization", "enters NCG scalar normalization lanes in continuum matching", "polynomial trace cable native, observed Yukawa values not native", StatusNoAbsoluteKineticScale},
		{"continuum matching", "field normalization", "maps finite sockets to physical pole/on-shell/MSbar observables", "missing threshold/matching theorem", StatusNoNativeRGThresholdTheorem},
	}
}

func buildFlavorRelation() []FlavorSealRGRelationRow {
	return []FlavorSealRGRelationRow{
		{"MinimalFlavorHistoryBranchSeal", "environmental flavor endpoint branch entering Y_core", "not native", "not used as RG beta-function law", StatusFlavorRelationRecorded},
		{"epsilon_e", "charged-lepton root chamber wall coordinate", "environmental fourth-root branch seal", "Gate 596/604 preserved", StatusNoFlavorFeedbackAsNativeRG},
		{"B_flav≈0", "orientation balance among epsilon_e, |U_e3|²/4, and +J_CKM", "environmental compatibility filter", "not fed back as native transport equation", StatusNoFlavorFeedbackAsNativeRG},
		{"Yukawa singular values", "observed endpoint and v1 transported flavor magnitudes", "observed ledger", "full matrix RGE and thresholds remain open", StatusFlavorRelationRecorded},
		{"CKM/PMNS", "observed orientation ledgers", "observed ledger", "no native CKM/PMNS derivation", StatusFlavorRelationRecorded},
	}
}

func buildProductTimeFirewall() ProductTimeFirewall {
	return ProductTimeFirewall{
		RGScaleIsProductTime: false,
		RGScaleIsOSHilbert:   false,
		RGScaleIsCosmoTime:   false,
		Statement:            "RG scale mu and t=ln(mu/M_Z) are transport parameters, not Lorentzian time, OS/Hilbert dynamics, or cosmological history time.",
		Verdict:              StatusRGScaleNotProductTime,
	}
}

func buildUpdatedFormula() UpdatedFormula {
	return UpdatedFormula{
		Formula:           "E_End(M_Z)=T_RG+threshold[NativeBoundary(k_Y=5/3,sin²θ*=3/8,g1=g2,EW Hessian sockets), EndpointLedgers(g_i,lambda,v,Y,CKM,PMNS), ThresholdSlots(delta_i^gauge,delta_lambda,delta_yukawa,delta_K_phi,delta_v), Firewalls(no endpoint/native promotion)]",
		NativeBoundary:    []string{"k_Y=5/3", "sin²(theta_*)=3/8", "g1=g2 boundary relation", "g'^2/g²=3/5", "symbolic EW Hessian", "A_F gauge sockets", "H_phi scalar doublet socket"},
		TransportSpine:    []string{"one-loop SM gauge running v1", "one-loop top-dominant scalar running v1", "diagonal Yukawa/flavor transport v1", "transport interval ln(Lambda12/M_Z)", "explicit threshold/matching gaps"},
		ThresholdSlots:    []string{"delta_i^gauge", "delta_lambda", "delta_yukawa", "delta_K_phi", "delta_v", "delta_pole_MSbar", "delta_boundary"},
		EndpointLedgers:   []string{"g1,g2,g3 at M_Z", "sin²(theta_End)", "lambda(M_Z)", "v from G_F", "m_W,m_Z,m_H", "Yukawa/CKM/PMNS ledgers"},
		BlockedPromotions: []string{"full gauge unification", "absolute kinetic scale", "Higgs VEV derivation", "physical W/Z/photon dynamics", "RG scale as product time", "observed endpoint derivation"},
		Verdict:           StatusUpdatedFormulaWritten,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, true, StatusGate606Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate605Inherited,
		StatusNativeBoundaryClassified,
		StatusEndpointLedgerBuilt,
		StatusGaugeRGSlotsDefined,
		StatusScalarRGSlotsDefined,
		StatusThresholdLedgerDefined,
		StatusKineticBlockersClassified,
		StatusFlavorRelationRecorded,
		StatusUpdatedFormulaWritten,
		StatusRGThresholdNextSpine,
		StatusGaugeMismatchThresholdNeeded,
		StatusScalarCrossingApproxVisible,
		StatusEndpointLedgerBridgeOnly,
		StatusNoNativeRGThresholdTheorem,
		StatusNoAbsoluteKineticScale,
		StatusNoHiggsVEVDerivation,
		StatusNoLowEnergyWZPhotonDynamics,
		StatusNoFullGaugeUnificationClaim,
		StatusNoScalarStabilityFinalClaim,
		StatusNoObservedEndpointDerivation,
		StatusNoFlavorFeedbackAsNativeRG,
		StatusRGScaleNotProductTime,
		StatusNoEndpointDerivationFirewall,
		StatusNoWZPhotonDerivationFirewall,
		StatusNoKoideFlavorPromotionFirewall,
		StatusThresholdsSchemesExplicitFirewall,
		StatusGate606Boundary,
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

func containsBoundary(rows []NativeBoundaryConditionRow, symbol string) bool {
	for _, r := range rows {
		if r.Symbol == symbol {
			return true
		}
	}
	return false
}

func containsEndpoint(rows []EndpointObservedLedgerRow, symbol string) bool {
	for _, r := range rows {
		if r.Symbol == symbol {
			return true
		}
	}
	return false
}

func containsGauge(rows []GaugeTransportRow, quantity string) bool {
	for _, r := range rows {
		if r.Quantity == quantity {
			return true
		}
	}
	return false
}

func containsScalar(rows []ScalarTransportRow, quantity string) bool {
	for _, r := range rows {
		if r.Quantity == quantity {
			return true
		}
	}
	return false
}

func containsThreshold(rows []ThresholdCorrectionSlotRow, slot string) bool {
	for _, r := range rows {
		if r.Slot == slot {
			return true
		}
	}
	return false
}

func containsBlocker(rows []KineticNormalizationBlockerRow, blocker string) bool {
	for _, r := range rows {
		if r.Blocker == blocker {
			return true
		}
	}
	return false
}

func containsString(xs []string, s string) bool {
	for _, x := range xs {
		if x == s || strings.Contains(x, s) {
			return true
		}
	}
	return false
}
