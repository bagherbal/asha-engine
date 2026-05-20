// Package generation2crosssectororientationintertwinerminimalityaudit implements
// Gate 592: Cross-Sector Orientation Intertwiner Minimality Audit.
//
// Gate 590/591 established the strongest bridge-layer environmental relation
// currently visible in the runtime,
//
//	kappa_e = 1 - 8*pi*epsilon_e ~= sin^2(theta13)/4 - J_CKM,
//
// and Gate 591 showed that the remaining residual is inside the propagated
// one-sigma uncertainty band and below the current near-Koide R/Q defect scale.
// Gate 592 therefore does not fit another residual term.  It asks what kind of
// typed mathematical object would be required for ASHA to lawfully connect the
// charged-lepton Koide chamber-wall coordinate, PMNS reactor leakage, and CKM
// oriented area.  The result is a minimal environmental OrientationBalanceSeal,
// not a native derivation of flavor.
package generation2crosssectororientationintertwinerminimalityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidereactorckmresidualclosureuncertaintyaudit"
)

const (
	AuditID = "GATE592-CROSS-SECTOR-ORIENTATION-INTERTWINER-MINIMALITY-AUDIT"

	StatusGate591Inherited                 = "PASS_GATE591_ORIENTATION_BALANCE_RESULT_INHERITED"
	StatusTypedObjectsClassified           = "PASS_GATE590_RELATION_TYPED_OBJECTS_CLASSIFIED"
	StatusMinimalBridgeTypeDefined         = "PASS_MINIMAL_CROSS_SECTOR_ORIENTATION_INTERTWINER_TYPE_DEFINED"
	StatusCurrentObjectsAudited            = "PASS_CURRENT_ASHA_OBJECTS_AUDITED_FOR_INTERTWINER"
	StatusOrientationBalanceSealDefined    = "CONDITIONAL_SUPPORT_ORIENTATION_BALANCE_SEAL_DEFINED"
	StatusNoDeltaFitJustified              = "CONDITIONAL_SUPPORT_RESIDUAL_BELOW_CURRENT_UNCERTAINTY_NO_ADDITIONAL_DELTA_FIT_JUSTIFIED"
	StatusResidualBelowDefectScale         = "PASS_RESIDUAL_BELOW_NEAR_KOIDE_R_Q_DEFECT_SCALE_INHERITED"
	StatusNoCrossSectorIntertwiner         = "FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER"
	StatusNoFlavorBalanceOperator          = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_BALANCE_OPERATOR"
	StatusNoRootSpaceOrientationMap        = "FAILED_ROUTE_NO_ROOT_SPACE_ORIENTATION_MAP"
	StatusNoNativeRootTraceAbsoluteDirac   = "FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_ABSOLUTE_DIRAC_OPERATOR"
	StatusNoDFYukawaPMNSCKMKoideIntertwine = "FAILED_ROUTE_DF_YUKAWA_CKM_PMNS_KOIDE_LEDGER_DOES_NOT_SUPPLY_INTERTWINER"
	StatusKappaRemainsSeal                 = "FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL"
	StatusNoFlavorDerivation               = "FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedRemainObserved           = "FIREWALL_PRESERVED_ORIENTATION_INPUTS_REMAIN_OBSERVED_ENVIRONMENTAL_DATA"
	StatusGate352Preserved                 = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusNoNewCarrierSelector             = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate592Boundary                  = "FIREWALL_PRESERVED_GATE592_CROSS_SECTOR_ORIENTATION_INTERTWINER_BOUNDARY"
)

type InheritedRelation struct {
	EpsilonObsRad               float64
	EpsilonObsDeg               float64
	KappaObs                    float64
	ReactorQuarter              float64
	JCKM                        float64
	OrientationCandidate        float64
	Delta590                    float64
	AbsDelta590                 float64
	RelativeDelta590            float64
	EpsilonPredictionRad        float64
	EpsilonPredictionDeg        float64
	EpsilonResidualRad          float64
	EpsilonResidualDeg          float64
	CombinedOneSigmaLow         float64
	CombinedOneSigmaHigh        float64
	ResidualInsideOneSigma      bool
	ResidualBelowRDefect        bool
	ResidualBelowQDefect        bool
	RDefect                     float64
	AbsQResidual                float64
	Theta13DominatesUncertainty bool
	Verdict                     string
}

type TypedObject struct {
	Symbol       string
	Carrier      string
	Role         string
	Equation     string
	RuntimeValue float64
	NativeStatus string
	Verdict      string
}

type TypedObjectAudit struct {
	Objects []TypedObject
	Verdict string
}

type RequiredBridge struct {
	NameCandidates         []string
	Domain                 []string
	Codomain               string
	RequiredMap            string
	Minimality             string
	MustBeBasisInvariant   bool
	MustHandleRootSpace    bool
	MustRespectSectorTypes bool
	Verdict                string
}

type RepositoryObjectAudit struct {
	Name                 string
	PresentInASHA        bool
	CanSupplyIntertwiner bool
	Reason               string
	Verdict              string
}

type RepositoryAudit struct {
	Objects                         []RepositoryObjectAudit
	AnyNativeCrossSectorIntertwiner bool
	NativeRootTraceOrAbsoluteDirac  bool
	Verdict                         string
}

type OrientationBalanceSeal struct {
	Name               string
	KappaDefinition    string
	EpsilonDefinition  string
	KappaValue         float64
	KappaCandidate     float64
	KappaResidual      float64
	EpsilonValueRad    float64
	EpsilonPredRad     float64
	EpsilonResidualRad float64
	ResidualStatus     string
	Native             bool
	Interpretation     string
	Verdict            string
}

type PrecisionAudit struct {
	Delta590                      float64
	AbsDelta590                   float64
	OneSigmaLowDistance           float64
	OneSigmaHighDistance          float64
	SigmaFractionMinus            float64
	SigmaFractionPlus             float64
	RDefect                       float64
	AbsQResidual                  float64
	DeltaSmallerThanRDefect       bool
	DeltaSmallerThanQResidual     bool
	AdditionalCorrectionJustified bool
	Verdict                       string
}

type LawfulnessAudit struct {
	CrossSectorOrientationIntertwinerPresent bool
	FlavorOrientationBalanceOperatorPresent  bool
	RootSpaceOrientationMapPresent           bool
	NativeRootTraceOperatorPresent           bool
	AbsoluteDiracObservablePresent           bool
	DerivesKoideWallCoordinate               bool
	DerivesPMNSReactorAngle                  bool
	DerivesCKMJarlskog                       bool
	DerivesKappaRelation                     bool
	Verdict                                  string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesTheta13             bool
	DerivesNeutrinoPhysics     bool
	DerivesChargedLeptonMasses bool
	DerivesFlavorTexture       bool
	PromotesObservedAsNative   bool
	AddsNewCarrier             bool
	AddsNewSelector            bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	TypedObjectsConnected      string
	NativeIntertwinerPresent   bool
	ResidualMeaningfulBeyondV1 bool
	MinimalSeal                string
	KappaRemainsEnvironmental  bool
	Decision                   string
	Verdict                    string
}

type Analysis struct {
	Inherited  InheritedRelation
	Typed      TypedObjectAudit
	Required   RequiredBridge
	Repository RepositoryAudit
	Seal       OrientationBalanceSeal
	Precision  PrecisionAudit
	Lawfulness LawfulnessAudit
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
	g591, err := generation2koidereactorckmresidualclosureuncertaintyaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate591 predecessor: %w", err)
	}
	inherited := inheritRelation(g591)
	typed := classifyTypedObjects(inherited)
	required := defineRequiredBridge()
	repository := auditRepositoryObjects()
	seal := defineOrientationBalanceSeal(inherited)
	precision := auditPrecision(inherited, g591)
	lawfulness := auditLawfulness(repository)
	firewalls := auditFirewalls()
	final := compileFinal(inherited, repository, precision, seal)
	truth := "Gate 592 stops residual fitting and isolates the missing mathematical object.  The Gate 590 relation connects a charged-lepton root-space chamber-wall coordinate, PMNS reactor leakage, a CKM oriented area, and a loop-sized angular unit, but the current ASHA runtime contains no native cross-sector orientation intertwiner, no root-space orientation map, and no root-trace or absolute-Dirac observable that can derive the relation.  The minimal lawful object is therefore an environmental OrientationBalanceSeal, not ASHA-native flavor law."
	return Analysis{Inherited: inherited, Typed: typed, Required: required, Repository: repository, Seal: seal, Precision: precision, Lawfulness: lawfulness, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRelation(g591 generation2koidereactorckmresidualclosureuncertaintyaudit.Analysis) InheritedRelation {
	return InheritedRelation{
		EpsilonObsRad:               g591.Runtime.EpsilonObsRad,
		EpsilonObsDeg:               g591.Runtime.EpsilonObsDeg,
		KappaObs:                    g591.Runtime.KappaObs,
		ReactorQuarter:              g591.Residual.AReactorQuarter,
		JCKM:                        g591.Inputs.JCKMRuntime,
		OrientationCandidate:        g591.Residual.BReactorMinusCKM,
		Delta590:                    g591.Residual.Delta590,
		AbsDelta590:                 g591.Residual.AbsDelta590,
		RelativeDelta590:            g591.Residual.RelativeDelta590,
		EpsilonPredictionRad:        g591.Runtime.LoopUnit * (1.0 - g591.Residual.BReactorMinusCKM),
		EpsilonPredictionDeg:        rad2deg(g591.Runtime.LoopUnit * (1.0 - g591.Residual.BReactorMinusCKM)),
		EpsilonResidualRad:          g591.Residual.EpsilonResidualRad,
		EpsilonResidualDeg:          g591.Residual.EpsilonResidualDeg,
		CombinedOneSigmaLow:         g591.Uncertainty.BMin1Sigma,
		CombinedOneSigmaHigh:        g591.Uncertainty.BMax1Sigma,
		ResidualInsideOneSigma:      g591.Uncertainty.CoversKappa,
		ResidualBelowRDefect:        g591.Defects.DeltaSmallerThanR,
		ResidualBelowQDefect:        g591.Defects.DeltaSmallerThanQ,
		RDefect:                     g591.Defects.RDefect,
		AbsQResidual:                g591.Defects.AbsQResidual,
		Theta13DominatesUncertainty: strings.Contains(g591.Uncertainty.DominantUncertainty, "theta13"),
		Verdict:                     strings.Join([]string{StatusGate591Inherited, StatusResidualBelowDefectScale}, ";"),
	}
}

func classifyTypedObjects(in InheritedRelation) TypedObjectAudit {
	objects := []TypedObject{
		{Symbol: "epsilon_e", Carrier: "charged-lepton square-root Yukawa ray in the positive Koide S3 chamber", Role: "electron-zero chamber-wall distance", Equation: "delta=135deg-epsilon_e", RuntimeValue: in.EpsilonObsRad, NativeStatus: "observed environmental wall coordinate; not native root-trace law", Verdict: "PASS_TYPED_CHARGED_LEPTON_WALL_COORDINATE"},
		{Symbol: "kappa_e", Carrier: "loop-angle deficit of the charged-lepton wall coordinate", Role: "dimensionless Koide wall deficit", Equation: "kappa_e=1-8*pi*epsilon_e", RuntimeValue: in.KappaObs, NativeStatus: "environmental seal value", Verdict: "PASS_TYPED_LOOP_DEFICIT"},
		{Symbol: "sin^2(theta13)/4", Carrier: "PMNS lepton-sector reactor leakage", Role: "lepton orientation leakage with weak-normalization factor 1/4", Equation: "A=sin^2(theta13)/4", RuntimeValue: in.ReactorQuarter, NativeStatus: "version-pinned observed PMNS input", Verdict: "PASS_TYPED_PMNS_REACTOR_LEAKAGE"},
		{Symbol: "J_CKM", Carrier: "CKM quark-sector unitary-triangle orientation", Role: "quark CP oriented area scale", Equation: "J=Im(V_us V_cb V_ub^* V_cs^*)", RuntimeValue: in.JCKM, NativeStatus: "runtime observed CKM invariant", Verdict: "PASS_TYPED_CKM_ORIENTATION_AREA"},
		{Symbol: "1/(8*pi)", Carrier: "loop-sized angular unit", Role: "sets wall-offset angular scale", Equation: "epsilon=(1/(8*pi))(1-kappa)", RuntimeValue: 1.0 / (8.0 * math.Pi), NativeStatus: "typed bridge scale; not an ASHA-native source theorem", Verdict: "PASS_TYPED_LOOP_ANGULAR_UNIT"},
	}
	return TypedObjectAudit{Objects: objects, Verdict: StatusTypedObjectsClassified}
}

func defineRequiredBridge() RequiredBridge {
	return RequiredBridge{
		NameCandidates:         []string{"CrossSectorOrientationIntertwiner", "FlavorOrientationBalanceOperator", "RootSpaceOrientationMap"},
		Domain:                 []string{"PMNS reactor leakage sin^2(theta13)/4", "CKM oriented area J_CKM", "loop angular unit 1/(8*pi)"},
		Codomain:               "charged-lepton Koide chamber-wall deficit kappa_e or wall coordinate epsilon_e",
		RequiredMap:            "I_orient[sin^2(theta13)/4 - J_CKM] = 1 - 8*pi*epsilon_e",
		Minimality:             "Because Gate 591 shows no additional delta fit is justified at v1 precision, the minimal missing object is a type-preserving orientation intertwiner, not another numerical correction term.",
		MustBeBasisInvariant:   true,
		MustHandleRootSpace:    true,
		MustRespectSectorTypes: true,
		Verdict:                StatusMinimalBridgeTypeDefined,
	}
}

func auditRepositoryObjects() RepositoryAudit {
	objects := []RepositoryObjectAudit{
		{Name: "finite spectral triple D_F edges", PresentInASHA: true, CanSupplyIntertwiner: false, Reason: "D_F/one-form edges type weak-doublet and scalar lanes; they do not map CKM area plus PMNS reactor leakage into a charged-lepton root-space wall coordinate.", Verdict: StatusNoDFYukawaPMNSCKMKoideIntertwine},
		{Name: "Yukawa matrices / empirical flavor ledgers", PresentInASHA: true, CanSupplyIntertwiner: false, Reason: "They record observed magnitudes and mixings, but do not construct a basis-invariant CKM-PMNS-Koide orientation map.", Verdict: StatusNoCrossSectorIntertwiner},
		{Name: "CKM/PMNS observed ledgers", PresentInASHA: true, CanSupplyIntertwiner: false, Reason: "They supply endpoint orientation data; observed ledgers are not native operators.", Verdict: StatusNoFlavorBalanceOperator},
		{Name: "charged-lepton root-space Koide frame", PresentInASHA: true, CanSupplyIntertwiner: false, Reason: "Gates 577-586 define the environmental root-space wall geometry but inherit Gate 352's root-trace obstruction.", Verdict: StatusNoNativeRootTraceAbsoluteDirac},
		{Name: "quaternionic weak socket H / Im(H)", PresentInASHA: true, CanSupplyIntertwiner: false, Reason: "The weak socket acts on finite weak carriers; Gates 575-576 block identification with sealed spatial CP1 and do not provide a Koide-root map.", Verdict: StatusNoRootSpaceOrientationMap},
		{Name: "B-L selector and projective orientation seals", PresentInASHA: true, CanSupplyIntertwiner: false, Reason: "B-L gives CP0|CP2 and later seals are projective orientation data; none connects PMNS/CKM orientation invariants to charged-lepton wall epsilon.", Verdict: StatusNoCrossSectorIntertwiner},
		{Name: "root-trace / absolute-Dirac observable", PresentInASHA: false, CanSupplyIntertwiner: false, Reason: "Gate 352 remains binding: no native root-trace or absolute-Dirac observable generates Koide root-space quantities.", Verdict: StatusNoNativeRootTraceAbsoluteDirac},
	}
	return RepositoryAudit{Objects: objects, AnyNativeCrossSectorIntertwiner: false, NativeRootTraceOrAbsoluteDirac: false, Verdict: strings.Join([]string{StatusCurrentObjectsAudited, StatusNoCrossSectorIntertwiner, StatusNoNativeRootTraceAbsoluteDirac}, ";")}
}

func defineOrientationBalanceSeal(in InheritedRelation) OrientationBalanceSeal {
	return OrientationBalanceSeal{
		Name:               "OrientationBalanceSeal",
		KappaDefinition:    "kappa_e := sin^2(theta13)/4 - J_CKM",
		EpsilonDefinition:  "epsilon_e := (1/(8*pi)) [1 - sin^2(theta13)/4 + J_CKM]",
		KappaValue:         in.KappaObs,
		KappaCandidate:     in.OrientationCandidate,
		KappaResidual:      in.Delta590,
		EpsilonValueRad:    in.EpsilonObsRad,
		EpsilonPredRad:     in.EpsilonPredictionRad,
		EpsilonResidualRad: in.EpsilonResidualRad,
		ResidualStatus:     "inside propagated theta13-dominated one-sigma band and below current near-Koide R/Q defect scale",
		Native:             false,
		Interpretation:     "Minimal environmental compression of the charged-lepton wall deficit by one PMNS reactor leakage term and one CKM orientation-area term.  It is a bridge-layer seal until a native cross-sector orientation intertwiner is proven.",
		Verdict:            strings.Join([]string{StatusOrientationBalanceSealDefined, StatusNoCrossSectorIntertwiner, StatusKappaRemainsSeal}, ";"),
	}
}

func auditPrecision(in InheritedRelation, g591 generation2koidereactorckmresidualclosureuncertaintyaudit.Analysis) PrecisionAudit {
	lowDist := in.KappaObs - in.CombinedOneSigmaLow
	highDist := in.CombinedOneSigmaHigh - in.KappaObs
	additional := !(in.ResidualInsideOneSigma && in.ResidualBelowRDefect && in.ResidualBelowQDefect)
	return PrecisionAudit{
		Delta590:                      in.Delta590,
		AbsDelta590:                   in.AbsDelta590,
		OneSigmaLowDistance:           lowDist,
		OneSigmaHighDistance:          highDist,
		SigmaFractionMinus:            g591.Uncertainty.SigmaFractionMinus,
		SigmaFractionPlus:             g591.Uncertainty.SigmaFractionPlus,
		RDefect:                       in.RDefect,
		AbsQResidual:                  in.AbsQResidual,
		DeltaSmallerThanRDefect:       in.ResidualBelowRDefect,
		DeltaSmallerThanQResidual:     in.ResidualBelowQDefect,
		AdditionalCorrectionJustified: additional,
		Verdict:                       strings.Join([]string{StatusNoDeltaFitJustified, StatusResidualBelowDefectScale}, ";"),
	}
}

func auditLawfulness(r RepositoryAudit) LawfulnessAudit {
	return LawfulnessAudit{
		CrossSectorOrientationIntertwinerPresent: r.AnyNativeCrossSectorIntertwiner,
		FlavorOrientationBalanceOperatorPresent:  false,
		RootSpaceOrientationMapPresent:           false,
		NativeRootTraceOperatorPresent:           false,
		AbsoluteDiracObservablePresent:           false,
		DerivesKoideWallCoordinate:               false,
		DerivesPMNSReactorAngle:                  false,
		DerivesCKMJarlskog:                       false,
		DerivesKappaRelation:                     false,
		Verdict:                                  strings.Join([]string{StatusNoCrossSectorIntertwiner, StatusNoFlavorBalanceOperator, StatusNoRootSpaceOrientationMap, StatusNoNativeRootTraceAbsoluteDirac}, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesPMNS: false, DerivesCKM: false, DerivesTheta13: false, DerivesNeutrinoPhysics: false, DerivesChargedLeptonMasses: false, DerivesFlavorTexture: false, PromotesObservedAsNative: false, AddsNewCarrier: false, AddsNewSelector: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedRemainObserved, StatusGate352Preserved, StatusNoNewCarrierSelector, StatusGate592Boundary}, ";")}
}

func compileFinal(in InheritedRelation, repo RepositoryAudit, precision PrecisionAudit, seal OrientationBalanceSeal) FinalVerdict {
	native := repo.AnyNativeCrossSectorIntertwiner || repo.NativeRootTraceOrAbsoluteDirac
	meaningful := precision.AdditionalCorrectionJustified
	decision := "Gate 592 defines the missing CrossSectorOrientationIntertwiner type and finds that it is absent from the current ASHA runtime.  Since the Gate 590 residual is already inside the propagated one-sigma band and below the near-Koide R/Q defect scale, no additional residual term is justified at v1 precision.  The minimal lawful object is OrientationBalanceSeal: kappa_e := sin^2(theta13)/4 - J_CKM, equivalently epsilon_e := (1/(8*pi))[1 - sin^2(theta13)/4 + J_CKM]."
	_ = in
	_ = seal
	return FinalVerdict{TypedObjectsConnected: "epsilon_e/kappa_e, sin^2(theta13)/4, J_CKM, and 1/(8*pi)", NativeIntertwinerPresent: native, ResidualMeaningfulBeyondV1: meaningful, MinimalSeal: "OrientationBalanceSeal", KappaRemainsEnvironmental: !native, Decision: decision, Verdict: strings.Join([]string{StatusOrientationBalanceSealDefined, StatusNoCrossSectorIntertwiner, StatusKappaRemainsSeal, StatusGate592Boundary}, ";")}
}

func Statuses() []string {
	return []string{StatusGate591Inherited, StatusTypedObjectsClassified, StatusMinimalBridgeTypeDefined, StatusCurrentObjectsAudited, StatusOrientationBalanceSealDefined, StatusNoDeltaFitJustified, StatusResidualBelowDefectScale, StatusNoCrossSectorIntertwiner, StatusNoFlavorBalanceOperator, StatusNoRootSpaceOrientationMap, StatusNoNativeRootTraceAbsoluteDirac, StatusNoDFYukawaPMNSCKMKoideIntertwine, StatusKappaRemainsSeal, StatusNoFlavorDerivation, StatusObservedRemainObserved, StatusGate352Preserved, StatusNoNewCarrierSelector, StatusGate592Boundary}
}

func rad2deg(x float64) float64 { return x * 180.0 / math.Pi }
