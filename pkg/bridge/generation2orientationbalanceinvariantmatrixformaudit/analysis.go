// Package generation2orientationbalanceinvariantmatrixformaudit implements
// Gate 593: OrientationBalance Invariant Matrix Form Audit.
//
// Gate 590-592 identified the strongest current bridge-layer environmental
// relation
//
//	1 - 8*pi*epsilon_e ~= sin^2(theta13)/4 - J_CKM.
//
// Gate 593 rewrites this relation in invariant projector / Jarlskog language so
// that the missing CrossSectorOrientationIntertwiner is a precise mathematical
// target rather than a symbolic name.  This remains an environmental audit: it
// does not derive Koide, PMNS, CKM, Yukawa spectra, neutrino physics, or flavor
// texture from ASHA-native law.
package generation2orientationbalanceinvariantmatrixformaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2crosssectororientationintertwinerminimalityaudit"
)

const (
	AuditID = "GATE593-ORIENTATION-BALANCE-INVARIANT-MATRIX-FORM-AUDIT"

	StatusGate592Inherited              = "PASS_GATE592_ORIENTATION_BALANCE_SEAL_INHERITED"
	StatusRootSpaceMapDefined           = "PASS_CHARGED_LEPTON_ROOT_SPACE_EPSILON_MAP_DEFINED"
	StatusPMNSProjectorTraceDefined     = "PASS_PMNS_REACTOR_PROJECTOR_TRACE_FORM_DEFINED"
	StatusCKMJarlskogFormsRecorded      = "PASS_CKM_JARLSKOG_REPHASING_AND_COMMUTATOR_FORM_RECORDED"
	StatusInvariantBalanceWritten       = "PASS_ORIENTATION_BALANCE_INVARIANT_MATRIX_FORM_WRITTEN"
	StatusBasisLabelAuditComplete       = "PASS_BASIS_AND_LABEL_DEPENDENCE_AUDITED"
	StatusASHAAvailabilityAudited       = "PASS_CURRENT_ASHA_AVAILABILITY_AUDITED"
	StatusInvariantFormSharpensTarget   = "CONDITIONAL_SUPPORT_INVARIANT_FORM_SHARPENS_OPERATOR_TARGET"
	StatusSealRewrittenInvariantly      = "CONDITIONAL_SUPPORT_ORIENTATIONBALANCESEAL_REWRITTEN_AS_PROJECTOR_COMMUTATOR_BALANCE"
	StatusRootSpectrumNotNative         = "FAILED_ROUTE_ROOT_SPECTRUM_MAP_NOT_NATIVE_GATE352"
	StatusNoNativeEpsilon               = "FAILED_ROUTE_NO_NATIVE_EPSILON_OF_YE_OPERATOR"
	StatusNoNativePMNSProjectorOperator = "FAILED_ROUTE_PMNS_PROJECTOR_TRACE_IS_OBSERVED_LEDGER_NOT_NATIVE_OPERATOR"
	StatusNoNativeCKMCommutatorMap      = "FAILED_ROUTE_NO_NATIVE_FLAVOR_COMMUTATOR_TO_KOIDE_WALL_MAP"
	StatusNoCrossSectorTraceOperator    = "FAILED_ROUTE_NO_CROSS_SECTOR_TRACE_ORIENTATION_BALANCE_OPERATOR"
	StatusNoIntertwiner                 = "FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER"
	StatusSealEnvironmental             = "FAILED_ROUTE_ORIENTATIONBALANCESEAL_REMAINS_ENVIRONMENTAL"
	StatusNoFlavorDerivation            = "FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_YUKAWA_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedDataFirewalled        = "FIREWALL_PRESERVED_PMNS_CKM_AND_YUKAWA_INPUTS_REMAIN_OBSERVED_ENVIRONMENTAL_DATA"
	StatusGate352Preserved              = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusNoNewCarrierSelector          = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate593Boundary               = "FIREWALL_PRESERVED_GATE593_ORIENTATION_BALANCE_INVARIANT_FORM_BOUNDARY"
)

type InheritedRelation struct {
	EpsilonObsRad        float64
	EpsilonObsDeg        float64
	KappaObs             float64
	ReactorQuarter       float64
	ReactorTrace         float64
	JCKM                 float64
	OrientationCandidate float64
	Delta590             float64
	AbsDelta590          float64
	EpsilonPredRad       float64
	EpsilonResidualRad   float64
	ResidualInsideSigma  bool
	ResidualBelowRDefect bool
	ResidualBelowQDefect bool
	Verdict              string
}

type RootSpaceMapAudit struct {
	Source                         string
	SingularValueMap               string
	SquareRootVector               string
	FourierKoideChamberCoordinate  string
	EpsilonFunctional              string
	RequiresRootSpectrumOperation  bool
	NativeRootTraceOperatorPresent bool
	NativeAbsoluteDiracPresent     bool
	Gate352ObstructionPreserved    bool
	Verdict                        string
}

type PMNSProjectorAudit struct {
	Expression     string
	Projectors     []string
	RequiredLabels []string
	TraceValue     float64
	ReactorQuarter float64
	Convention     string
	NativeOperator bool
	ObservedLedger bool
	Verdict        string
}

type CKMOrientationAudit struct {
	RephasingInvariantExpression string
	JCKM                         float64
	CommutatorExpression         string
	CommutatorSignConvention     string
	RequiredLabels               []string
	BasisInvariantGivenSpectra   bool
	NativeOperator               bool
	ObservedLedger               bool
	Verdict                      string
}

type InvariantBalanceAudit struct {
	EquationProjector      string
	EquationFunctional     string
	LeftKappa              float64
	RightProjectorMinusCKM float64
	Residual               float64
	AbsResidual            float64
	EpsilonEquation        string
	EpsilonPredictionRad   float64
	EpsilonResidualRad     float64
	ResidualInsideSigma    bool
	Verdict                string
}

type LabelDependency struct {
	Object        string
	InvariantPart string
	RequiredSeal  string
	Reason        string
	Verdict       string
}

type LabelAudit struct {
	Labels                []LabelDependency
	AllLabelSealsExplicit bool
	Verdict               string
}

type AvailabilityItem struct {
	Object           string
	ObservedLedger   bool
	NativeOperator   bool
	CanSupplyBalance bool
	Reason           string
	Verdict          string
}

type AvailabilityAudit struct {
	Items                        []AvailabilityItem
	AnyNativeBalanceOperator     bool
	AnyNativeRootSpectrumMap     bool
	AnyNativeFlavorCommutatorMap bool
	Verdict                      string
}

type OperatorTarget struct {
	Name                     string
	Domain                   []string
	Codomain                 string
	ZeroCondition            string
	RequiredEquation         string
	MustHandleRootSpectrum   bool
	MustHandleProjectors     bool
	MustHandleJarlskogArea   bool
	MustBeRephasingInvariant bool
	NativePresent            bool
	Verdict                  string
}

type FirewallAudit struct {
	DerivesKoide           bool
	DerivesPMNS            bool
	DerivesCKM             bool
	DerivesYukawas         bool
	DerivesNeutrinoPhysics bool
	DerivesFlavorTexture   bool
	PromotesObservedData   bool
	AddsNewCarrier         bool
	AddsNewSelector        bool
	PreservesGate352       bool
	Verdict                string
}

type FinalVerdict struct {
	InvariantFormAvailable              bool
	RequiredLabels                      string
	NativeOperatorPresent               bool
	MissingOperatorTarget               string
	OrientationBalanceSealEnvironmental bool
	Decision                            string
	Verdict                             string
}

type Analysis struct {
	Inherited    InheritedRelation
	RootSpace    RootSpaceMapAudit
	PMNS         PMNSProjectorAudit
	CKM          CKMOrientationAudit
	Balance      InvariantBalanceAudit
	Labels       LabelAudit
	Availability AvailabilityAudit
	Target       OperatorTarget
	Firewalls    FirewallAudit
	Final        FinalVerdict
	Truth        string
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
	g592, err := generation2crosssectororientationintertwinerminimalityaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate592 predecessor: %w", err)
	}
	inherited := inheritGate592(g592)
	root := auditRootSpaceMap()
	pmns := auditPMNSProjector(inherited)
	ckm := auditCKMOrientation(inherited)
	balance := writeInvariantBalance(inherited)
	labels := auditLabels()
	availability := auditAvailability()
	target := defineOperatorTarget()
	firewalls := auditFirewalls()
	final := compileFinal(labels, availability, target)
	truth := "Gate 593 rewrites the Gate 590/592 orientation balance in matrix/projector form: 1-8*pi*epsilon(Y_e) ?= (1/4)Tr(P_e U_PMNS P_3^nu U_PMNS^dagger) - J(Y_u,Y_d).  The rewrite makes the missing object precise: a cross-sector trace/orientation balance operator able to connect charged-lepton root-space chamber geometry, PMNS reactor projectors, and CKM Jarlskog area.  The current ASHA runtime has observed ledgers for these quantities but no native root-spectrum, PMNS-projector, CKM-commutator-to-Koide-wall, or cross-sector orientation operator."
	return Analysis{Inherited: inherited, RootSpace: root, PMNS: pmns, CKM: ckm, Balance: balance, Labels: labels, Availability: availability, Target: target, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritGate592(g592 generation2crosssectororientationintertwinerminimalityaudit.Analysis) InheritedRelation {
	trace := 4.0 * g592.Inherited.ReactorQuarter
	return InheritedRelation{
		EpsilonObsRad:        g592.Inherited.EpsilonObsRad,
		EpsilonObsDeg:        g592.Inherited.EpsilonObsDeg,
		KappaObs:             g592.Inherited.KappaObs,
		ReactorQuarter:       g592.Inherited.ReactorQuarter,
		ReactorTrace:         trace,
		JCKM:                 g592.Inherited.JCKM,
		OrientationCandidate: g592.Inherited.OrientationCandidate,
		Delta590:             g592.Inherited.Delta590,
		AbsDelta590:          g592.Inherited.AbsDelta590,
		EpsilonPredRad:       g592.Inherited.EpsilonPredictionRad,
		EpsilonResidualRad:   g592.Inherited.EpsilonResidualRad,
		ResidualInsideSigma:  g592.Inherited.ResidualInsideOneSigma,
		ResidualBelowRDefect: g592.Inherited.ResidualBelowRDefect,
		ResidualBelowQDefect: g592.Inherited.ResidualBelowQDefect,
		Verdict:              StatusGate592Inherited,
	}
}

func auditRootSpaceMap() RootSpaceMapAudit {
	return RootSpaceMapAudit{
		Source:                         "charged-lepton Yukawa singular values y_e,y_mu,y_tau from observed endpoint ledger",
		SingularValueMap:               "Y_e -> spec_sing(Y_e) = (y_e,y_mu,y_tau)",
		SquareRootVector:               "x_e=(sqrt(y_e),sqrt(y_mu),sqrt(y_tau)) in R^3_+",
		FourierKoideChamberCoordinate:  "x_j=A[1+sqrt(2)R cos(delta+2*pi*j/3)], canonical chamber (e,mu,tau), epsilon(Y_e)=135deg-delta",
		EpsilonFunctional:              "epsilon(Y_e) is a root-spectrum / chamber functional; kappa(Y_e)=1-8*pi*epsilon(Y_e)",
		RequiresRootSpectrumOperation:  true,
		NativeRootTraceOperatorPresent: false,
		NativeAbsoluteDiracPresent:     false,
		Gate352ObstructionPreserved:    true,
		Verdict:                        strings.Join([]string{StatusRootSpaceMapDefined, StatusRootSpectrumNotNative, StatusNoNativeEpsilon, StatusGate352Preserved}, ";"),
	}
}

func auditPMNSProjector(in InheritedRelation) PMNSProjectorAudit {
	return PMNSProjectorAudit{
		Expression:     "sin^2(theta13)=|U_PMNS[e,3]|^2=Tr(P_e U_PMNS P_3^nu U_PMNS^dagger)",
		Projectors:     []string{"P_e=electron flavor projector", "P_3^nu=third neutrino mass-eigenstate projector"},
		RequiredLabels: []string{"charged-lepton flavor basis label e", "neutrino mass eigenstate label 3", "PMNS standard convention", "mass ordering seal"},
		TraceValue:     in.ReactorTrace,
		ReactorQuarter: in.ReactorQuarter,
		Convention:     "NuFIT 6.0 normal-ordering theta13 convention inherited from Gates 587-591",
		NativeOperator: false,
		ObservedLedger: true,
		Verdict:        strings.Join([]string{StatusPMNSProjectorTraceDefined, StatusNoNativePMNSProjectorOperator}, ";"),
	}
}

func auditCKMOrientation(in InheritedRelation) CKMOrientationAudit {
	return CKMOrientationAudit{
		RephasingInvariantExpression: "J_CKM=Im(V_us V_cb V_ub^* V_cs^*)",
		JCKM:                         in.JCKM,
		CommutatorExpression:         "det([H_u,H_d]) = 2 i J_CKM prod_{i<j}(y_{u_i}^2-y_{u_j}^2) prod_{i<j}(y_{d_i}^2-y_{d_j}^2), up to sign/convention",
		CommutatorSignConvention:     "orientation sign depends on quark-generation ordering and CKM phase convention; |J| is rephasing invariant, signed J requires orientation convention",
		RequiredLabels:               []string{"up-quark generation ordering", "down-quark generation ordering", "CKM orientation sign", "Y_u,Y_d Hermitian products H_u,H_d"},
		BasisInvariantGivenSpectra:   true,
		NativeOperator:               false,
		ObservedLedger:               true,
		Verdict:                      strings.Join([]string{StatusCKMJarlskogFormsRecorded, StatusNoNativeCKMCommutatorMap}, ";"),
	}
}

func writeInvariantBalance(in InheritedRelation) InvariantBalanceAudit {
	return InvariantBalanceAudit{
		EquationProjector:      "kappa_e ?= (1/4)Tr(P_e U_PMNS P_3^nu U_PMNS^dagger) - J_CKM",
		EquationFunctional:     "1 - 8*pi*epsilon(Y_e) ?= (1/4)Tr(P_e U_PMNS P_3^nu U_PMNS^dagger) - J(Y_u,Y_d)",
		LeftKappa:              in.KappaObs,
		RightProjectorMinusCKM: in.OrientationCandidate,
		Residual:               in.Delta590,
		AbsResidual:            in.AbsDelta590,
		EpsilonEquation:        "epsilon(Y_e) ?= (1/(8*pi))[1 - (1/4)Tr(P_e U_PMNS P_3^nu U_PMNS^dagger) + J(Y_u,Y_d)]",
		EpsilonPredictionRad:   in.EpsilonPredRad,
		EpsilonResidualRad:     in.EpsilonResidualRad,
		ResidualInsideSigma:    in.ResidualInsideSigma,
		Verdict:                strings.Join([]string{StatusInvariantBalanceWritten, StatusSealRewrittenInvariantly}, ";"),
	}
}

func auditLabels() LabelAudit {
	labels := []LabelDependency{
		{Object: "epsilon(Y_e)", InvariantPart: "root-spectrum chamber coordinate after singular values are ordered", RequiredSeal: "charged-lepton chamber/order seal (e,mu,tau) and electron-zero wall choice", Reason: "Koide Fourier chamber walls are S3-permutation dependent; the observed wall offset is canonical only after ordering is fixed.", Verdict: "PASS_LABEL_CHARGED_LEPTON_CHAMBER_REQUIRED"},
		{Object: "P_e", InvariantPart: "rank-one flavor projector once charged-lepton flavor basis is selected", RequiredSeal: "electron flavor label", Reason: "The trace |U_e3|^2 requires the electron row; without P_e it is not a scalar target.", Verdict: "PASS_LABEL_ELECTRON_PROJECTOR_REQUIRED"},
		{Object: "P_3^nu", InvariantPart: "rank-one mass-eigenstate projector", RequiredSeal: "third neutrino mass-eigenstate label and mass ordering", Reason: "The reactor entry is the third PMNS column in the chosen ordering/convention.", Verdict: "PASS_LABEL_NEUTRINO_PROJECTOR_REQUIRED"},
		{Object: "J_CKM", InvariantPart: "rephasing-invariant area magnitude; signed orientation needs convention", RequiredSeal: "quark generation ordering and CKM orientation sign", Reason: "Changing generation orientation can flip the signed J used in the balance.", Verdict: "PASS_LABEL_CKM_ORIENTATION_REQUIRED"},
		{Object: "1/(8*pi)", InvariantPart: "dimensionless loop-sized angular unit", RequiredSeal: "bridge normalization convention", Reason: "The loop unit is typed as an environmental angular scale, not native source theorem.", Verdict: "PASS_LABEL_LOOP_UNIT_CONVENTION_REQUIRED"},
	}
	return LabelAudit{Labels: labels, AllLabelSealsExplicit: true, Verdict: StatusBasisLabelAuditComplete}
}

func auditAvailability() AvailabilityAudit {
	items := []AvailabilityItem{
		{Object: "charged-lepton root-space epsilon(Y_e)", ObservedLedger: true, NativeOperator: false, CanSupplyBalance: false, Reason: "Gates 577-586 compute it from observed singular values; Gate 352 blocks a native root-trace/absolute-Dirac source.", Verdict: StatusRootSpectrumNotNative},
		{Object: "PMNS projector trace Tr(P_e U P_3 U^dagger)", ObservedLedger: true, NativeOperator: false, CanSupplyBalance: false, Reason: "NuFIT reactor data supplies the trace value, but ASHA does not derive PMNS projectors or theta13.", Verdict: StatusNoNativePMNSProjectorOperator},
		{Object: "CKM Jarlskog invariant / commutator area", ObservedLedger: true, NativeOperator: false, CanSupplyBalance: false, Reason: "Runtime CKM ledger supplies J; no native commutator operator maps it into Koide root-space.", Verdict: StatusNoNativeCKMCommutatorMap},
		{Object: "finite spectral triple D_F and one-form edges", ObservedLedger: false, NativeOperator: true, CanSupplyBalance: false, Reason: "They type weak/scalar carriers but do not construct a CKM-PMNS-Koide orientation balance.", Verdict: StatusNoCrossSectorTraceOperator},
		{Object: "quaternionic weak socket H / Im(H)", ObservedLedger: false, NativeOperator: true, CanSupplyBalance: false, Reason: "It identifies the weak socket, not the charged-lepton root-space wall or cross-sector Jarlskog map.", Verdict: StatusNoCrossSectorTraceOperator},
		{Object: "CrossSectorOrientationIntertwiner", ObservedLedger: false, NativeOperator: false, CanSupplyBalance: false, Reason: "This is the missing target isolated by Gates 592-593.", Verdict: StatusNoIntertwiner},
	}
	return AvailabilityAudit{Items: items, AnyNativeBalanceOperator: false, AnyNativeRootSpectrumMap: false, AnyNativeFlavorCommutatorMap: false, Verdict: strings.Join([]string{StatusASHAAvailabilityAudited, StatusNoCrossSectorTraceOperator, StatusNoIntertwiner}, ";")}
}

func defineOperatorTarget() OperatorTarget {
	return OperatorTarget{
		Name:                     "CrossSectorOrientationIntertwiner",
		Domain:                   []string{"charged-lepton root-space chamber functional epsilon(Y_e)", "PMNS reactor projector trace Tr(P_e U_PMNS P_3^nu U_PMNS^dagger)", "CKM Jarlskog commutator area J(Y_u,Y_d)", "loop angular unit 1/(8*pi)"},
		Codomain:                 "scalar orientation balance residual",
		ZeroCondition:            "I_orient=0 iff 1-8*pi*epsilon(Y_e) = (1/4)|U_e3|^2 - J_CKM",
		RequiredEquation:         "I_orient := 1 - 8*pi*epsilon(Y_e) - (1/4)Tr(P_e U_PMNS P_3^nu U_PMNS^dagger) + J(Y_u,Y_d)",
		MustHandleRootSpectrum:   true,
		MustHandleProjectors:     true,
		MustHandleJarlskogArea:   true,
		MustBeRephasingInvariant: true,
		NativePresent:            false,
		Verdict:                  strings.Join([]string{StatusInvariantFormSharpensTarget, StatusNoCrossSectorTraceOperator, StatusNoIntertwiner}, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesPMNS: false, DerivesCKM: false, DerivesYukawas: false, DerivesNeutrinoPhysics: false, DerivesFlavorTexture: false, PromotesObservedData: false, AddsNewCarrier: false, AddsNewSelector: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedDataFirewalled, StatusGate352Preserved, StatusNoNewCarrierSelector, StatusGate593Boundary}, ";")}
}

func compileFinal(labels LabelAudit, availability AvailabilityAudit, target OperatorTarget) FinalVerdict {
	native := availability.AnyNativeBalanceOperator || availability.AnyNativeRootSpectrumMap || availability.AnyNativeFlavorCommutatorMap || target.NativePresent
	decision := "Gate 593 writes the Gate 590 relation in invariant matrix/projector form and therefore turns the missing CrossSectorOrientationIntertwiner into a concrete operator target: I_orient := 1 - 8*pi*epsilon(Y_e) - (1/4)Tr(P_e U_PMNS P_3^nu U_PMNS^dagger) + J(Y_u,Y_d).  Current ASHA has observed ledgers for the ingredients but no native root-spectrum epsilon operator, no PMNS projector derivation, no CKM commutator-to-Koide-wall map, and no cross-sector orientation balance operator.  OrientationBalanceSeal remains environmental."
	return FinalVerdict{InvariantFormAvailable: true, RequiredLabels: fmt.Sprintf("%d explicit label/seal dependencies", len(labels.Labels)), NativeOperatorPresent: native, MissingOperatorTarget: target.Name, OrientationBalanceSealEnvironmental: !native, Decision: decision, Verdict: strings.Join([]string{StatusInvariantBalanceWritten, StatusNoCrossSectorTraceOperator, StatusSealEnvironmental, StatusGate593Boundary}, ";")}
}

func Statuses() []string {
	return []string{StatusGate592Inherited, StatusRootSpaceMapDefined, StatusPMNSProjectorTraceDefined, StatusCKMJarlskogFormsRecorded, StatusInvariantBalanceWritten, StatusBasisLabelAuditComplete, StatusASHAAvailabilityAudited, StatusInvariantFormSharpensTarget, StatusSealRewrittenInvariantly, StatusRootSpectrumNotNative, StatusNoNativeEpsilon, StatusNoNativePMNSProjectorOperator, StatusNoNativeCKMCommutatorMap, StatusNoCrossSectorTraceOperator, StatusNoIntertwiner, StatusSealEnvironmental, StatusNoFlavorDerivation, StatusObservedDataFirewalled, StatusGate352Preserved, StatusNoNewCarrierSelector, StatusGate593Boundary}
}

func rad2deg(x float64) float64 { return x * 180.0 / math.Pi }
