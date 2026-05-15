// Package generation2yukawaselectorairlock implements Gate 489:
// Yukawa Selector Airlock Boundary Decision.
//
// Gates 486-488 showed that CKM compression cannot be promoted from a null-C3
// coordinate socket to native physics unless ASHA first derives sector-specific
// up/down family operators O_u,O_d and two rephasing-invariant CKM polynomial
// constraints. Gate 488 identified the only remaining socket: the finite
// Dirac/Yukawa coefficient block. Gate 489 audits whether existing native
// variational or spectral-action structures select the Yukawa matrices.
//
// The result is a formal airlock decision. Native geometry supplies the finite
// graph, representation labels, first-order admissibility, gauge/Higgs shape,
// and universal family baselines. It does not supply the 3x3 complex Yukawa
// entries, rank-three sector matrices, their relative eigenbasis, or CKM/Jarlskog
// invariants. The CKM branch is therefore closed for native prediction at this
// layer and redirected to a sealed empirical comparator airlock.
package generation2yukawaselectorairlock

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE489-YUKAWA-SELECTOR-AIRLOCK-BOUNDARY-DECISION"

	StatusGate488Inherited                 = "CONDITIONAL_SUPPORT_GATE488_YUKAWA_SOCKET_INHERITED"
	StatusSelectorLedgerConstructed        = "CONDITIONAL_SUPPORT_YUKAWA_SELECTOR_LEDGER_CONSTRUCTED"
	StatusNativeYukawaSlotsConfirmed       = "CONDITIONAL_SUPPORT_NATIVE_YUKAWA_SLOTS_CONFIRMED"
	StatusSpectralActionGenerationBlind    = "FAILED_ROUTE_SPECTRAL_ACTION_DOES_NOT_SELECT_YUKAWA_TEXTURE"
	StatusVariationalSelectorAbsent        = "FAILED_ROUTE_NATIVE_VARIATIONAL_YUKAWA_SELECTOR_NOT_FOUND"
	StatusNoRankThreeMatricesDerived       = "FAILED_ROUTE_RANK_THREE_UP_DOWN_YUKAWA_MATRICES_NOT_DERIVED"
	StatusNoEigenbasisDerived              = "FAILED_ROUTE_UP_DOWN_EIGENBASIS_ORIENTATION_NOT_DERIVED"
	StatusNoCKMInvariantsDerived           = "FAILED_ROUTE_CKM_JARLSKOG_INVARIANTS_NOT_DERIVED"
	StatusYukawaAirlockClosedNative        = "FIREWALL_CLOSED_NATIVE_YUKAWA_SELECTOR_BRANCH"
	StatusCKMEnvironmentalQuarantineFormal = "FIREWALL_FORMAL_CKM_ORIENTATION_ENVIRONMENTAL_QUARANTINE"
	StatusGate490RedirectDefined           = "CONDITIONAL_SUPPORT_GATE490_NATIVE_WORK_REDIRECT_DEFINED"
)

const (
	NativeFlavorDim                 = 13
	KXYCoeffDim                     = 9
	RequiredCKMInvariantConstraints = 2
	DerivedCKMInvariantConstraints  = 0
)

type Inheritance struct {
	Executed                          bool
	Gate485NullC3BaselineInherited    bool
	Gate486CKMCompressionBlocked      bool
	Gate487CommutatorObstruction      bool
	Gate488YukawaSocketInherited      bool
	Gate488NativeUpDownLabelsFound    bool
	Gate488NativeUpDownOperatorsFound bool
	Gate488YukawaValuesDerived        bool
	NoObservedCKMImported             bool
	Verdict                           string
	Reason                            string
}

type SelectorCandidate struct {
	Name                                  string
	NativeLayer                           string
	Native                                bool
	TouchesYukawaSocket                   bool
	DistinguishesUpDown                   bool
	GenerationAware                       bool
	SectorSpecific                        bool
	SelectsCoefficientValues              bool
	SelectsRankThreeMatrices              bool
	SelectsRelativeEigenbasis             bool
	ProducesRephasingInvariantConstraints int
	RequiresEmpiricalInput                bool
	CanServeAsNativeSelector              bool
	Verdict                               string
	Reason                                string
}

type SelectorLedger struct {
	Executed                      bool
	Candidates                    []SelectorCandidate
	CandidateCount                int
	NativeYukawaSlotCandidates    int
	UpDownAwareCandidates         int
	GenerationAwareCandidates     int
	CoefficientSelectors          int
	RankThreeMatrixSelectors      int
	EigenbasisSelectors           int
	NativeSelectorsPassing        int
	RephasingInvariantConstraints int
	ObservedDataImported          bool
	Verdict                       string
	Reason                        string
}

type VariationalAudit struct {
	Executed                       bool
	FiniteSpectralActionAudited    bool
	FirstOrderConditionAudited     bool
	HiggsOneFormGraphAudited       bool
	KGenFamilyAxisAudited          bool
	GaugeKineticHessianAudited     bool
	NativeYukawaSlotsExist         bool
	NativeYukawaValuesDerived      bool
	RankThreeUpMatrixDerived       bool
	RankThreeDownMatrixDerived     bool
	RelativeEigenbasisDerived      bool
	CKMInvariantConstraintsDerived int
	SelectorFound                  bool
	Verdict                        string
	Reason                         string
}

type AirlockDecision struct {
	Executed                         bool
	NativeYukawaSelectorBranchClosed bool
	YukawaEntriesEnvironmental       bool
	CKMOrientationEnvironmental      bool
	CKMMatrixEnvironmental           bool
	JarlskogEnvironmental            bool
	AllowedBridgeComparator          bool
	NativeCKMPredictionAllowed       bool
	FutureEmpiricalUseRequiresLabel  bool
	FutureEmpiricalUseRequiresScale  bool
	FutureEmpiricalUseRequiresSource bool
	Verdict                          string
	Reason                           string
}

type Firewall struct {
	Executed                          bool
	ObservedCKMImported               bool
	ObservedWolfensteinImported       bool
	ObservedQuarkMassesImported       bool
	ObservedYukawaEntriesImported     bool
	NativeYukawaMatrixWritten         bool
	NativeUpOperatorWritten           bool
	NativeDownOperatorWritten         bool
	NativeDiagonalizersWritten        bool
	CKMMatrixNativePrediction         bool
	JarlskogNativePrediction          bool
	CKMInvariantConstraintNativeWrite bool
	NativeRegistryWritten             bool
	NativeFlavorDimAfter              int
	KXYCoeffDimAfter                  int
	Verdict                           string
	Reason                            string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Ledger      SelectorLedger
	Variational VariationalAudit
	Airlock     AirlockDecision
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Ledger = buildSelectorLedger()
	a.Variational = buildVariationalAudit(a.Ledger)
	a.Airlock = buildAirlockDecision(a.Variational)
	a.Firewall = buildFirewall(a)
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate485NullC3BaselineInherited:    true,
		Gate486CKMCompressionBlocked:      true,
		Gate487CommutatorObstruction:      true,
		Gate488YukawaSocketInherited:      true,
		Gate488NativeUpDownLabelsFound:    true,
		Gate488NativeUpDownOperatorsFound: false,
		Gate488YukawaValuesDerived:        false,
		NoObservedCKMImported:             true,
		Verdict:                           StatusGate488Inherited,
		Reason:                            "Gate488 left exactly one candidate socket: the finite Dirac/Yukawa coefficient block. Gate489 may audit selector principles, but cannot import CKM, quark masses, or Yukawa entries.",
	}
}

func buildSelectorLedger() SelectorLedger {
	candidates := []SelectorCandidate{
		{
			Name: "finite Dirac/Yukawa coefficient block", NativeLayer: "finite spectral triple coefficient socket", Native: true,
			TouchesYukawaSocket: true, DistinguishesUpDown: true, GenerationAware: true, SectorSpecific: true,
			SelectsCoefficientValues: false, SelectsRankThreeMatrices: false, SelectsRelativeEigenbasis: false,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: true, CanServeAsNativeSelector: false,
			Verdict: StatusNativeYukawaSlotsConfirmed,
			Reason:  "the block has the right arity and sector labels, but its entries are free coefficients unless an independent native selector fixes them",
		},
		{
			Name: "Chamseddine-Connes spectral action traces", NativeLayer: "almost-commutative spectral action", Native: true,
			TouchesYukawaSocket: true, DistinguishesUpDown: false, GenerationAware: false, SectorSpecific: false,
			SelectsCoefficientValues: false, SelectsRankThreeMatrices: false, SelectsRelativeEigenbasis: false,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: false, CanServeAsNativeSelector: false,
			Verdict: StatusSpectralActionGenerationBlind,
			Reason:  "trace invariants can measure already-supplied Yukawa data, but the action does not choose the matrix entries or their up/down relative eigenbasis",
		},
		{
			Name: "first-order condition and admissible Dirac graph", NativeLayer: "finite geometry admissibility sieve", Native: true,
			TouchesYukawaSocket: true, DistinguishesUpDown: true, GenerationAware: false, SectorSpecific: false,
			SelectsCoefficientValues: false, SelectsRankThreeMatrices: false, SelectsRelativeEigenbasis: false,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: false, CanServeAsNativeSelector: false,
			Verdict: StatusVariationalSelectorAbsent,
			Reason:  "the order-one graph says which edges are legal; it does not assign the complex 3x3 weights on those edges",
		},
		{
			Name: "Higgs one-form edge measure", NativeLayer: "inner fluctuation / Higgs-as-one-form", Native: true,
			TouchesYukawaSocket: true, DistinguishesUpDown: true, GenerationAware: false, SectorSpecific: false,
			SelectsCoefficientValues: false, SelectsRankThreeMatrices: false, SelectsRelativeEigenbasis: false,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: false, CanServeAsNativeSelector: false,
			Verdict: StatusSpectralActionGenerationBlind,
			Reason:  "the Higgs edge normalizes a universal scalar channel and tree-level proxy, but it is generation-blind and cannot orient O_u against O_d",
		},
		{
			Name: "K_gen/null-C3 family baseline", NativeLayer: "family axis and null mass-shadow geometry", Native: true,
			TouchesYukawaSocket: false, DistinguishesUpDown: false, GenerationAware: true, SectorSpecific: false,
			SelectsCoefficientValues: false, SelectsRankThreeMatrices: false, SelectsRelativeEigenbasis: false,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: false, CanServeAsNativeSelector: false,
			Verdict: StatusVariationalSelectorAbsent,
			Reason:  "the family baseline fixes a universal C3 shape constraint, not sector-specific Yukawa matrices or CKM orientation",
		},
		{
			Name: "gauge kinetic Hessian and representation traces", NativeLayer: "gauge normalization / Hessian lane", Native: true,
			TouchesYukawaSocket: false, DistinguishesUpDown: false, GenerationAware: false, SectorSpecific: false,
			SelectsCoefficientValues: false, SelectsRankThreeMatrices: false, SelectsRelativeEigenbasis: false,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: false, CanServeAsNativeSelector: false,
			Verdict: StatusSpectralActionGenerationBlind,
			Reason:  "gauge traces fix gauge normalization data, not flavor texture entries or eigenbasis mismatch",
		},
		{
			Name: "empirical Yukawa seal", NativeLayer: "environmental airlock", Native: false,
			TouchesYukawaSocket: true, DistinguishesUpDown: true, GenerationAware: true, SectorSpecific: true,
			SelectsCoefficientValues: true, SelectsRankThreeMatrices: true, SelectsRelativeEigenbasis: true,
			ProducesRephasingInvariantConstraints: 0, RequiresEmpiricalInput: true, CanServeAsNativeSelector: false,
			Verdict: StatusCKMEnvironmentalQuarantineFormal,
			Reason:  "empirical rows can populate the socket for comparison only; they are not a theorem source and cannot back-write into the native registry",
		},
	}
	ledger := SelectorLedger{Executed: true, Candidates: candidates, CandidateCount: len(candidates), ObservedDataImported: false, Verdict: StatusVariationalSelectorAbsent}
	for _, c := range candidates {
		if c.Native && c.TouchesYukawaSocket {
			ledger.NativeYukawaSlotCandidates++
		}
		if c.DistinguishesUpDown {
			ledger.UpDownAwareCandidates++
		}
		if c.GenerationAware {
			ledger.GenerationAwareCandidates++
		}
		if c.SelectsCoefficientValues {
			ledger.CoefficientSelectors++
		}
		if c.SelectsRankThreeMatrices {
			ledger.RankThreeMatrixSelectors++
		}
		if c.SelectsRelativeEigenbasis {
			ledger.EigenbasisSelectors++
		}
		ledger.RephasingInvariantConstraints += c.ProducesRephasingInvariantConstraints
		if selectorPasses(c) {
			ledger.NativeSelectorsPassing++
		}
	}
	ledger.Reason = fmt.Sprintf("%d selector candidates were audited; %d native candidates touch the Yukawa socket, but %d native candidates select coefficients, rank-three matrices, eigenbasis, and two invariants", ledger.CandidateCount, ledger.NativeYukawaSlotCandidates, ledger.NativeSelectorsPassing)
	return ledger
}

func selectorPasses(c SelectorCandidate) bool {
	return c.Native && c.TouchesYukawaSocket && c.DistinguishesUpDown && c.GenerationAware && c.SectorSpecific && c.SelectsCoefficientValues && c.SelectsRankThreeMatrices && c.SelectsRelativeEigenbasis && c.ProducesRephasingInvariantConstraints >= RequiredCKMInvariantConstraints && !c.RequiresEmpiricalInput && c.CanServeAsNativeSelector
}

func buildVariationalAudit(l SelectorLedger) VariationalAudit {
	return VariationalAudit{
		Executed:                       true,
		FiniteSpectralActionAudited:    true,
		FirstOrderConditionAudited:     true,
		HiggsOneFormGraphAudited:       true,
		KGenFamilyAxisAudited:          true,
		GaugeKineticHessianAudited:     true,
		NativeYukawaSlotsExist:         l.NativeYukawaSlotCandidates > 0,
		NativeYukawaValuesDerived:      false,
		RankThreeUpMatrixDerived:       false,
		RankThreeDownMatrixDerived:     false,
		RelativeEigenbasisDerived:      false,
		CKMInvariantConstraintsDerived: DerivedCKMInvariantConstraints,
		SelectorFound:                  l.NativeSelectorsPassing > 0,
		Verdict:                        StatusVariationalSelectorAbsent,
		Reason:                         "native variational structures constrain admissibility, scalar/gauge normalization, and universal baselines, but no audited action extremizes or uniquely selects the complex 3x3 Yukawa matrices",
	}
}

func buildAirlockDecision(v VariationalAudit) AirlockDecision {
	closed := !v.SelectorFound
	return AirlockDecision{
		Executed:                         true,
		NativeYukawaSelectorBranchClosed: closed,
		YukawaEntriesEnvironmental:       closed,
		CKMOrientationEnvironmental:      closed,
		CKMMatrixEnvironmental:           closed,
		JarlskogEnvironmental:            closed,
		AllowedBridgeComparator:          closed,
		NativeCKMPredictionAllowed:       false,
		FutureEmpiricalUseRequiresLabel:  true,
		FutureEmpiricalUseRequiresScale:  true,
		FutureEmpiricalUseRequiresSource: true,
		Verdict:                          StatusCKMEnvironmentalQuarantineFormal,
		Reason:                           "with no native Yukawa selector, future CKM/Yukawa work may only enter as explicitly labeled bridge/environmental comparator rows with scheme, scale, and provenance metadata",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                          true,
		ObservedCKMImported:               false,
		ObservedWolfensteinImported:       false,
		ObservedQuarkMassesImported:       false,
		ObservedYukawaEntriesImported:     false,
		NativeYukawaMatrixWritten:         false,
		NativeUpOperatorWritten:           false,
		NativeDownOperatorWritten:         false,
		NativeDiagonalizersWritten:        false,
		CKMMatrixNativePrediction:         false,
		JarlskogNativePrediction:          false,
		CKMInvariantConstraintNativeWrite: false,
		NativeRegistryWritten:             false,
		NativeFlavorDimAfter:              NativeFlavorDim,
		KXYCoeffDimAfter:                  KXYCoeffDim,
		Verdict:                           StatusYukawaAirlockClosedNative,
		Reason:                            "Gate489 writes no Yukawa matrix, O_u/O_d matrix, diagonalizer, CKM matrix, Jarlskog value, or CKM invariant constraint to the native registry",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"finite Dirac/Higgs geometry supplies admissible Yukawa sockets only",
			"spectral action and gauge/Higgs Hessian lanes remain generation-blind with respect to Yukawa texture selection",
		},
		BridgeEntries: []string{
			"Yukawa and CKM comparator rows are allowed only through an explicit airlock with sector, scheme, scale, and provenance labels",
			"synthetic Yukawa matrices may be used to test algorithms, but never as native predictions",
		},
		EnvironmentalEntries: []string{
			"Yukawa matrix entries, quark masses, CKM matrix, Wolfenstein parameters, CP phase, and Jarlskog value are environmental/bridge data at this layer",
		},
		FailedRoutes: []string{
			StatusSpectralActionGenerationBlind,
			StatusVariationalSelectorAbsent,
			StatusNoRankThreeMatricesDerived,
			StatusNoEigenbasisDerived,
			StatusNoCKMInvariantsDerived,
			StatusYukawaAirlockClosedNative,
		},
		OpenTheorems: []string{
			StatusGate490RedirectDefined,
			"search for native non-flavor consequences of the accepted finite law-space rather than fitting flavor moduli",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        490,
		Title:       "Native Frontier Redirect After Flavor Airlock Closure",
		Reason:      "Gate489 closes the current CKM/Yukawa native-prediction branch. The next valid work should move away from flavor fitting and toward native invariant consequences already supported by ASHA.",
		PrimaryTask: "select a non-flavor theorem lane where the finite core still has native leverage, such as anomaly/topological charge ledgers, scalar-edge stability, or continuum matching permissions",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate485NullC3BaselineInherited || !a.Inheritance.Gate486CKMCompressionBlocked || !a.Inheritance.Gate487CommutatorObstruction || !a.Inheritance.Gate488YukawaSocketInherited || !a.Inheritance.Gate488NativeUpDownLabelsFound || a.Inheritance.Gate488NativeUpDownOperatorsFound || a.Inheritance.Gate488YukawaValuesDerived || !a.Inheritance.NoObservedCKMImported {
		return fmt.Errorf("Gate489 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Ledger.Executed || a.Ledger.CandidateCount < 7 || a.Ledger.NativeYukawaSlotCandidates == 0 || a.Ledger.UpDownAwareCandidates == 0 || a.Ledger.GenerationAwareCandidates == 0 || a.Ledger.NativeSelectorsPassing != 0 || a.Ledger.RephasingInvariantConstraints != 0 || a.Ledger.ObservedDataImported {
		return fmt.Errorf("Gate489 selector ledger invalid: %+v", a.Ledger)
	}
	if !a.Variational.Executed || !a.Variational.FiniteSpectralActionAudited || !a.Variational.FirstOrderConditionAudited || !a.Variational.HiggsOneFormGraphAudited || !a.Variational.KGenFamilyAxisAudited || !a.Variational.GaugeKineticHessianAudited || !a.Variational.NativeYukawaSlotsExist || a.Variational.NativeYukawaValuesDerived || a.Variational.RankThreeUpMatrixDerived || a.Variational.RankThreeDownMatrixDerived || a.Variational.RelativeEigenbasisDerived || a.Variational.CKMInvariantConstraintsDerived != 0 || a.Variational.SelectorFound {
		return fmt.Errorf("Gate489 variational audit invalid: %+v", a.Variational)
	}
	if !a.Airlock.Executed || !a.Airlock.NativeYukawaSelectorBranchClosed || !a.Airlock.YukawaEntriesEnvironmental || !a.Airlock.CKMOrientationEnvironmental || !a.Airlock.CKMMatrixEnvironmental || !a.Airlock.JarlskogEnvironmental || !a.Airlock.AllowedBridgeComparator || a.Airlock.NativeCKMPredictionAllowed || !a.Airlock.FutureEmpiricalUseRequiresLabel || !a.Airlock.FutureEmpiricalUseRequiresScale || !a.Airlock.FutureEmpiricalUseRequiresSource {
		return fmt.Errorf("Gate489 airlock invalid: %+v", a.Airlock)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedCKMImported || a.Firewall.ObservedWolfensteinImported || a.Firewall.ObservedQuarkMassesImported || a.Firewall.ObservedYukawaEntriesImported || a.Firewall.NativeYukawaMatrixWritten || a.Firewall.NativeUpOperatorWritten || a.Firewall.NativeDownOperatorWritten || a.Firewall.NativeDiagonalizersWritten || a.Firewall.CKMMatrixNativePrediction || a.Firewall.JarlskogNativePrediction || a.Firewall.CKMInvariantConstraintNativeWrite || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate489 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate489 closes the current CKM/Yukawa native-prediction branch. ASHA has native Yukawa sockets and admissible Higgs/Dirac edge structure, but no native variational or spectral-action selector for the complex 3x3 up/down matrices, no relative eigenbasis, and no CKM/Jarlskog invariant constraints. Therefore Yukawa entries and CKM orientation are formally quarantined as bridge/environmental data. Audited candidates=%d; native selectors passing=%d; derived CKM constraints=%d.", a.Ledger.CandidateCount, a.Ledger.NativeSelectorsPassing, a.Variational.CKMInvariantConstraintsDerived)
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: Gate485=%t Gate486_blocked=%t Gate487_obstruction=%t Gate488_socket=%t updown_labels=%t native_OuOd=%t Yukawa_values=%t; %s", x.Verdict, x.Gate485NullC3BaselineInherited, x.Gate486CKMCompressionBlocked, x.Gate487CommutatorObstruction, x.Gate488YukawaSocketInherited, x.Gate488NativeUpDownLabelsFound, x.Gate488NativeUpDownOperatorsFound, x.Gate488YukawaValuesDerived, x.Reason)
}

func FormatLedger(x SelectorLedger) string {
	return fmt.Sprintf("%s: candidates=%d native_socket_candidates=%d updown_aware=%d generation_aware=%d coefficient_selectors=%d rank3_selectors=%d eigenbasis_selectors=%d native_selectors_passing=%d invariant_constraints=%d; %s", x.Verdict, x.CandidateCount, x.NativeYukawaSlotCandidates, x.UpDownAwareCandidates, x.GenerationAwareCandidates, x.CoefficientSelectors, x.RankThreeMatrixSelectors, x.EigenbasisSelectors, x.NativeSelectorsPassing, x.RephasingInvariantConstraints, x.Reason)
}

func FormatVariational(x VariationalAudit) string {
	return fmt.Sprintf("%s: spectral_action=%t first_order=%t higgs_edge=%t K_gen=%t gauge_hessian=%t slots=%t values=%t rank3_up=%t rank3_down=%t eigenbasis=%t constraints=%d/%d selector=%t; %s", x.Verdict, x.FiniteSpectralActionAudited, x.FirstOrderConditionAudited, x.HiggsOneFormGraphAudited, x.KGenFamilyAxisAudited, x.GaugeKineticHessianAudited, x.NativeYukawaSlotsExist, x.NativeYukawaValuesDerived, x.RankThreeUpMatrixDerived, x.RankThreeDownMatrixDerived, x.RelativeEigenbasisDerived, x.CKMInvariantConstraintsDerived, RequiredCKMInvariantConstraints, x.SelectorFound, x.Reason)
}

func FormatAirlock(x AirlockDecision) string {
	return fmt.Sprintf("%s: native_branch_closed=%t Yukawa_env=%t CKM_orientation_env=%t CKM_matrix_env=%t J_env=%t bridge_comparator_allowed=%t native_CKM_allowed=%t metadata(label=%t scale=%t source=%t); %s", x.Verdict, x.NativeYukawaSelectorBranchClosed, x.YukawaEntriesEnvironmental, x.CKMOrientationEnvironmental, x.CKMMatrixEnvironmental, x.JarlskogEnvironmental, x.AllowedBridgeComparator, x.NativeCKMPredictionAllowed, x.FutureEmpiricalUseRequiresLabel, x.FutureEmpiricalUseRequiresScale, x.FutureEmpiricalUseRequiresSource, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_CKM=%t observed_Yukawa=%t native_Yukawa=%t native_Ou=%t native_Od=%t native_CKM=%t native_J=%t invariant_write=%t registry_write=%t dim=%d KXY=%d; %s", x.Verdict, x.ObservedCKMImported, x.ObservedYukawaEntriesImported, x.NativeYukawaMatrixWritten, x.NativeUpOperatorWritten, x.NativeDownOperatorWritten, x.CKMMatrixNativePrediction, x.JarlskogNativePrediction, x.CKMInvariantConstraintNativeWrite, x.NativeRegistryWritten, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 489 Registry Audit — Yukawa Selector Airlock Boundary Decision\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusNativeYukawaSlotsConfirmed,
		StatusSpectralActionGenerationBlind,
		StatusVariationalSelectorAbsent,
		StatusNoRankThreeMatricesDerived,
		StatusNoEigenbasisDerived,
		StatusNoCKMInvariantsDerived,
		StatusYukawaAirlockClosedNative,
		StatusCKMEnvironmentalQuarantineFormal,
	} {
		b.WriteString("- `" + v + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate485 derived only the null-C3 Koide baseline. Gate486 blocked CKM 4->2 as native. Gate487 proved null spectra do not determine commutator/Jarlskog structure. Gate488 found native up/down labels and a Yukawa socket but no native O_u/O_d matrices.\n\n")
	b.WriteString("## Yukawa selector ledger\n\n")
	b.WriteString(FormatLedger(a.Ledger) + "\n\n")
	b.WriteString("| Candidate | Native layer | Native? | Socket? | Up/down? | Gen-aware? | Values? | Rank-3? | Eigenbasis? | Constraints | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Ledger.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %t | %t | %t | %t | %t | %d | `%s` |\n", c.Name, c.NativeLayer, c.Native, c.TouchesYukawaSocket, c.DistinguishesUpDown, c.GenerationAware, c.SelectsCoefficientValues, c.SelectsRankThreeMatrices, c.SelectsRelativeEigenbasis, c.ProducesRephasingInvariantConstraints, c.Verdict))
	}
	b.WriteString("\n## Variational and spectral-action audit\n\n")
	b.WriteString(FormatVariational(a.Variational) + "\n\n")
	b.WriteString("The spectral action can evaluate trace expressions once a finite Dirac/Yukawa block is supplied, but this is not a coefficient-selection theorem. Admissibility and one-form graph constraints define legal edges; they do not choose the complex 3x3 texture or the relative up/down eigenbasis.\n\n")
	b.WriteString("## Airlock decision\n\n")
	b.WriteString(FormatAirlock(a.Airlock) + "\n\n")
	b.WriteString("The CKM/Yukawa branch is closed for native prediction at this layer. Future rows may enter only as bridge/environmental comparators with explicit sector, scheme, scale, and provenance metadata.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No observed CKM, Wolfenstein, quark-mass, or Yukawa-entry data were imported. No native Yukawa matrix, O_u/O_d operator, diagonalizer, CKM matrix, Jarlskog invariant, or CKM polynomial constraint was written.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("**Gate %d — %s.** %s Primary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
