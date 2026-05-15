// Package generation2nativeupdownsource implements Gate 488:
// Native Up/Down Operator Source Search.
//
// Gate 487 proved that the Gate485 null-C3 spectrum does not determine a
// physical CKM/Jarlskog commutator: equal null spectra can be assigned arbitrary
// relative eigenbases. Gate 488 therefore searches the already-native ASHA
// structures for the missing object: an algebraic source of sector-specific
// up/down family operators O_u,O_d whose diagonalizers would be physical before
// the empirical airlock.
//
// The result is fail-closed but informative. Native electroweak/Higgs data do
// distinguish up-type and down-type slots, while color separates quarks from
// leptons and C3/K_gen provide universal family structure. None of these native
// candidates simultaneously supplies an up/down split, generation-aware family
// eigenbasis, two operator diagonalizers, and two rephasing-invariant CKM
// polynomial constraints. The Standard Model Yukawa matrices remain sealed
// bridge/environmental coefficients, not native Clifford outputs.
package generation2nativeupdownsource

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE488-NATIVE-UP-DOWN-OPERATOR-SOURCE-SEARCH"

	StatusGate487Inherited                    = "CONDITIONAL_SUPPORT_GATE487_COMMUTATOR_OBSTRUCTION_INHERITED"
	StatusSourceLedgerConstructed             = "CONDITIONAL_SUPPORT_NATIVE_SOURCE_LEDGER_CONSTRUCTED"
	StatusNativeUpDownSectorLabelsFound       = "CONDITIONAL_SUPPORT_NATIVE_UP_DOWN_SECTOR_LABELS_FOUND"
	StatusNativeQuarkLeptonSeparatorFound     = "CONDITIONAL_SUPPORT_NATIVE_QUARK_LEPTON_SEPARATOR_FOUND"
	StatusNativeUniversalFamilyAxisFound      = "CONDITIONAL_SUPPORT_NATIVE_UNIVERSAL_FAMILY_AXIS_FOUND"
	StatusGenerationBlindNativeSources        = "FAILED_ROUTE_AVAILABLE_NATIVE_SOURCES_ARE_GENERATION_BLIND_OR_SECTOR_NEUTRAL"
	StatusNoNativeUpDownEigenbasisSource      = "FAILED_ROUTE_NATIVE_UP_DOWN_FAMILY_EIGENBASIS_SOURCE_NOT_FOUND"
	StatusNoNativeUpDownOperatorsDerived      = "FAILED_ROUTE_NATIVE_UP_DOWN_CLIFFORD_OPERATORS_NOT_DERIVED"
	StatusNoCKMInvariantConstraintsDerived    = "FAILED_ROUTE_CKM_REPHASING_INVARIANT_CONSTRAINTS_STILL_ZERO"
	StatusYukawaMatricesRemainSealed          = "FAILED_ROUTE_YUKAWA_MATRIX_ENTRIES_REMAIN_SEALED_BRIDGE_ENVIRONMENTAL_DATA"
	StatusCKMOrientationQuarantined           = "FAILED_ROUTE_CKM_EIGENBASIS_ORIENTATION_REMAINS_QUARANTINED"
	StatusFirewallBlockedNativeOperatorWrite  = "FIREWALL_BLOCKED_NATIVE_UP_DOWN_OPERATOR_REGISTRY_WRITE"
	StatusGate489YukawaAirlockDecisionDefined = "CONDITIONAL_SUPPORT_GATE489_YUKAWA_AIRLOCK_BOUNDARY_DECISION_DEFINED"
)

const (
	NativeFlavorDim                 = 13
	KXYCoeffDim                     = 9
	RequiredCKMInvariantConstraints = 2
	DerivedCKMInvariantConstraints  = 0
)

type Inheritance struct {
	Executed                       bool
	Gate485NullC3BaselineInherited bool
	Gate486NullMirrorBridgeOnly    bool
	Gate487CommutatorObstruction   bool
	Gate487NullSpectrumOnly        bool
	Gate487RequiredConstraints     int
	Gate487DerivedConstraints      int
	NoObservedCKMImported          bool
	Verdict                        string
	Reason                         string
}

type CandidateSource struct {
	Name                                  string
	NativeLayer                           string
	Native                                bool
	DistinguishesUpDown                   bool
	DistinguishesQuarkLepton              bool
	GenerationAware                       bool
	SectorNeutral                         bool
	SuppliesFamilyEigenbasis              bool
	SuppliesNativeUpOperator              bool
	SuppliesNativeDownOperator            bool
	SuppliesNativeDiagonalizers           bool
	RephasingInvariantConstraintsProduced int
	RequiresEmpiricalCoefficients         bool
	CanServeAsCKMSource                   bool
	Verdict                               string
	Reason                                string
}

type SourceLedger struct {
	Executed                      bool
	Candidates                    []CandidateSource
	CandidateCount                int
	NativeUpDownLabelSources      int
	NativeQuarkLeptonSeparators   int
	NativeUniversalFamilyAxes     int
	GenerationAwareCandidates     int
	SourcesPassingAllRequirements int
	OnlySlotsNotOperators         bool
	YukawaEntriesSealed           bool
	NoObservedDataImported        bool
	Verdict                       string
	Reason                        string
}

type RequirementAudit struct {
	Executed                        bool
	RequiresUpDownSplit             bool
	RequiresGenerationAwareness     bool
	RequiresFamilyEigenbasis        bool
	RequiresNativeUpOperator        bool
	RequiresNativeDownOperator      bool
	RequiresNativeDiagonalizers     bool
	RequiresTwoInvariantConstraints bool
	CandidatesPassing               int
	NativeUpDownOperatorsDerived    bool
	NativeDiagonalizersDerived      bool
	CKMInvariantConstraintsDerived  int
	NativeCKMSourceFound            bool
	Verdict                         string
	Reason                          string
}

type OperatorSocket struct {
	Executed                    bool
	UpDownSectorLabelsNative    bool
	YukawaSlotsNative           bool
	YukawaMatrixValuesNative    bool
	FamilyEigenbasisNative      bool
	CanNameOuOdSlots            bool
	CanPopulateOuOdNatively     bool
	CanComputeUuDaggerUd        bool
	CanComputeJarlskogInvariant bool
	BridgeAirlockRequired       bool
	Verdict                     string
	Reason                      string
}

type Firewall struct {
	Executed                          bool
	ObservedCKMImported               bool
	ObservedWolfensteinImported       bool
	ObservedQuarkMassesImported       bool
	ObservedYukawaEntriesImported     bool
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
	Inheritance  Inheritance
	Ledger       SourceLedger
	Requirements RequirementAudit
	Socket       OperatorSocket
	Firewall     Firewall
	Registry     RegistryUpdate
	Next         NextStep
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Ledger = buildSourceLedger()
	a.Requirements = buildRequirementAudit(a.Ledger)
	a.Socket = buildOperatorSocket(a.Ledger, a.Requirements)
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
		Executed:                       true,
		Gate485NullC3BaselineInherited: true,
		Gate486NullMirrorBridgeOnly:    true,
		Gate487CommutatorObstruction:   true,
		Gate487NullSpectrumOnly:        true,
		Gate487RequiredConstraints:     RequiredCKMInvariantConstraints,
		Gate487DerivedConstraints:      DerivedCKMInvariantConstraints,
		NoObservedCKMImported:          true,
		Verdict:                        StatusGate487Inherited,
		Reason:                         "Gate487 proved that a shared null-C3 spectrum does not determine the relative up/down eigenbasis; Gate488 must find a native operator source, not another spectral or coordinate shortcut",
	}
}

func buildSourceLedger() SourceLedger {
	candidates := []CandidateSource{
		{
			Name: "weak isospin and hypercharge charge table", NativeLayer: "finite spectral triple / electroweak representation", Native: true,
			DistinguishesUpDown: true, DistinguishesQuarkLepton: true, GenerationAware: false, SectorNeutral: false,
			SuppliesFamilyEigenbasis: false, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: false, CanServeAsCKMSource: false,
			Verdict: StatusNativeUpDownSectorLabelsFound,
			Reason:  "T3/Y identify up-type and down-type representation slots, but the action is identical across the three families and supplies no family eigenbasis orientation",
		},
		{
			Name: "Higgs one-form edge orientation", NativeLayer: "inner fluctuation / finite one-form graph", Native: true,
			DistinguishesUpDown: true, DistinguishesQuarkLepton: false, GenerationAware: false, SectorNeutral: false,
			SuppliesFamilyEigenbasis: false, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: true, CanServeAsCKMSource: false,
			Verdict: StatusYukawaMatricesRemainSealed,
			Reason:  "the finite graph names up/down Yukawa slots, but its 3x3 coefficient matrices are not fixed by the native graph and remain sealed bridge/environmental inputs",
		},
		{
			Name: "SU(3) color and QCD dressing topology", NativeLayer: "gauge/color sector", Native: true,
			DistinguishesUpDown: false, DistinguishesQuarkLepton: true, GenerationAware: false, SectorNeutral: false,
			SuppliesFamilyEigenbasis: false, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: false, CanServeAsCKMSource: false,
			Verdict: StatusNativeQuarkLeptonSeparatorFound,
			Reason:  "color separates quarks from leptons and explains why quark mass shadows are dressed, but it is generation-blind and does not orient up against down",
		},
		{
			Name: "Gate485 null-C3 Koide baseline", NativeLayer: "C3 mass-shadow null boundary", Native: true,
			DistinguishesUpDown: false, DistinguishesQuarkLepton: false, GenerationAware: true, SectorNeutral: true,
			SuppliesFamilyEigenbasis: false, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: false, CanServeAsCKMSource: false,
			Verdict: StatusGenerationBlindNativeSources,
			Reason:  "the null-C3 baseline fixes the universal spectral shape R/S=sqrt(2), but it is sector-neutral and says nothing about the relative up/down eigenbasis",
		},
		{
			Name: "K_gen primitive family axis", NativeLayer: "generation structural axis", Native: true,
			DistinguishesUpDown: false, DistinguishesQuarkLepton: false, GenerationAware: true, SectorNeutral: true,
			SuppliesFamilyEigenbasis: true, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: false, CanServeAsCKMSource: false,
			Verdict: StatusNativeUniversalFamilyAxisFound,
			Reason:  "K_gen gives a universal family reference axis, but a common axis for all sectors produces no CKM misalignment and no sector-specific pair O_u,O_d",
		},
		{
			Name: "finite Dirac/Yukawa block", NativeLayer: "finite Dirac operator coefficient slots", Native: true,
			DistinguishesUpDown: true, DistinguishesQuarkLepton: true, GenerationAware: true, SectorNeutral: false,
			SuppliesFamilyEigenbasis: false, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: true, CanServeAsCKMSource: false,
			Verdict: StatusYukawaMatricesRemainSealed,
			Reason:  "the block has exactly the right socket type for O_u and O_d, but the entries are free Yukawa data unless a further native selector theorem is supplied",
		},
		{
			Name: "triality / Spin(8) family-cycle intuition", NativeLayer: "Cℓ(1,7) representation symmetry", Native: true,
			DistinguishesUpDown: false, DistinguishesQuarkLepton: false, GenerationAware: true, SectorNeutral: true,
			SuppliesFamilyEigenbasis: false, SuppliesNativeUpOperator: false, SuppliesNativeDownOperator: false, SuppliesNativeDiagonalizers: false,
			RephasingInvariantConstraintsProduced: 0, RequiresEmpiricalCoefficients: false, CanServeAsCKMSource: false,
			Verdict: StatusGenerationBlindNativeSources,
			Reason:  "triality motivates a threefold family grammar, but this gate finds no native triality map that distinguishes up-sector and down-sector family operators",
		},
	}
	updown, ql, familyAxis, genAware, passing := 0, 0, 0, 0, 0
	yukawaSealed := false
	for _, c := range candidates {
		if c.DistinguishesUpDown {
			updown++
		}
		if c.DistinguishesQuarkLepton {
			ql++
		}
		if c.GenerationAware {
			genAware++
		}
		if c.GenerationAware && c.SectorNeutral {
			familyAxis++
		}
		if c.RequiresEmpiricalCoefficients {
			yukawaSealed = true
		}
		if sourcePasses(c) {
			passing++
		}
	}
	return SourceLedger{
		Executed:                      true,
		Candidates:                    candidates,
		CandidateCount:                len(candidates),
		NativeUpDownLabelSources:      updown,
		NativeQuarkLeptonSeparators:   ql,
		NativeUniversalFamilyAxes:     familyAxis,
		GenerationAwareCandidates:     genAware,
		SourcesPassingAllRequirements: passing,
		OnlySlotsNotOperators:         true,
		YukawaEntriesSealed:           yukawaSealed,
		NoObservedDataImported:        true,
		Verdict:                       StatusNoNativeUpDownEigenbasisSource,
		Reason:                        "native ledgers contain up/down labels and universal family structure, but no candidate simultaneously gives sector-specific family operators, diagonalizers, and rephasing-invariant CKM constraints",
	}
}

func sourcePasses(c CandidateSource) bool {
	return c.Native && c.DistinguishesUpDown && c.GenerationAware && c.SuppliesFamilyEigenbasis && c.SuppliesNativeUpOperator && c.SuppliesNativeDownOperator && c.SuppliesNativeDiagonalizers && c.RephasingInvariantConstraintsProduced >= RequiredCKMInvariantConstraints && !c.RequiresEmpiricalCoefficients && c.CanServeAsCKMSource
}

func buildRequirementAudit(l SourceLedger) RequirementAudit {
	return RequirementAudit{
		Executed:                        true,
		RequiresUpDownSplit:             true,
		RequiresGenerationAwareness:     true,
		RequiresFamilyEigenbasis:        true,
		RequiresNativeUpOperator:        true,
		RequiresNativeDownOperator:      true,
		RequiresNativeDiagonalizers:     true,
		RequiresTwoInvariantConstraints: true,
		CandidatesPassing:               l.SourcesPassingAllRequirements,
		NativeUpDownOperatorsDerived:    false,
		NativeDiagonalizersDerived:      false,
		CKMInvariantConstraintsDerived:  DerivedCKMInvariantConstraints,
		NativeCKMSourceFound:            false,
		Verdict:                         StatusNoNativeUpDownOperatorsDerived,
		Reason:                          fmt.Sprintf("%d native candidates were audited; %d name up/down slots and %d are generation-aware, but %d pass the full CKM-source requirement", l.CandidateCount, l.NativeUpDownLabelSources, l.GenerationAwareCandidates, l.SourcesPassingAllRequirements),
	}
}

func buildOperatorSocket(l SourceLedger, r RequirementAudit) OperatorSocket {
	return OperatorSocket{
		Executed:                    true,
		UpDownSectorLabelsNative:    l.NativeUpDownLabelSources > 0,
		YukawaSlotsNative:           true,
		YukawaMatrixValuesNative:    false,
		FamilyEigenbasisNative:      false,
		CanNameOuOdSlots:            true,
		CanPopulateOuOdNatively:     false,
		CanComputeUuDaggerUd:        false,
		CanComputeJarlskogInvariant: false,
		BridgeAirlockRequired:       !r.NativeCKMSourceFound,
		Verdict:                     StatusCKMOrientationQuarantined,
		Reason:                      "ASHA can name the O_u/O_d sockets through finite electroweak-Higgs structure, but cannot populate the 3x3 family operators or diagonalize them without sealed Yukawa data",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                          true,
		ObservedCKMImported:               false,
		ObservedWolfensteinImported:       false,
		ObservedQuarkMassesImported:       false,
		ObservedYukawaEntriesImported:     false,
		NativeUpOperatorWritten:           false,
		NativeDownOperatorWritten:         false,
		NativeDiagonalizersWritten:        false,
		CKMMatrixNativePrediction:         false,
		JarlskogNativePrediction:          false,
		CKMInvariantConstraintNativeWrite: false,
		NativeRegistryWritten:             false,
		NativeFlavorDimAfter:              NativeFlavorDim,
		KXYCoeffDimAfter:                  KXYCoeffDim,
		Verdict:                           StatusFirewallBlockedNativeOperatorWrite,
		Reason:                            "Gate488 writes no native O_u/O_d matrices, diagonalizers, CKM matrix, Jarlskog value, or invariant polynomial constraints; Yukawa entries remain behind the bridge/environmental airlock",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"native electroweak/Higgs representation data can label up-type and down-type slots",
			"native color topology separates quark/lepton sectors but remains generation-blind",
			"native K_gen/null-C3 family structure remains universal and sector-neutral",
		},
		BridgeEntries: []string{
			"O_u/O_d may be named as bridge sockets attached to finite Yukawa blocks",
			"future synthetic operator tests may use the Gate487 commutator sieve, but not as native predictions",
		},
		EnvironmentalEntries: []string{
			"Yukawa matrix entries, quark masses, CKM matrix, Wolfenstein parameters, and CP phase remain quarantined comparator data",
		},
		FailedRoutes: []string{
			StatusGenerationBlindNativeSources,
			StatusNoNativeUpDownEigenbasisSource,
			StatusNoNativeUpDownOperatorsDerived,
			StatusNoCKMInvariantConstraintsDerived,
			StatusYukawaMatricesRemainSealed,
			StatusCKMOrientationQuarantined,
		},
		OpenTheorems: []string{
			StatusGate489YukawaAirlockDecisionDefined,
			"decide whether to search for a native Yukawa coefficient selector, or formally close CKM orientation as environmental input beyond the finite core",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        489,
		Title:       "Yukawa Selector Airlock Boundary Decision",
		Reason:      "Gate488 finds native up/down labels but no native family operators. The only remaining CKM source socket is the finite Dirac/Yukawa coefficient block, whose entries are sealed.",
		PrimaryTask: "audit whether any native variational or spectral-action principle selects Yukawa matrices; if not, formally mark CKM orientation as environmental bridge data and redirect native work away from flavor fitting",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate485NullC3BaselineInherited || !a.Inheritance.Gate486NullMirrorBridgeOnly || !a.Inheritance.Gate487CommutatorObstruction || !a.Inheritance.Gate487NullSpectrumOnly || a.Inheritance.Gate487RequiredConstraints != RequiredCKMInvariantConstraints || a.Inheritance.Gate487DerivedConstraints != 0 || !a.Inheritance.NoObservedCKMImported {
		return fmt.Errorf("Gate488 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Ledger.Executed || a.Ledger.CandidateCount < 7 || a.Ledger.NativeUpDownLabelSources == 0 || a.Ledger.NativeQuarkLeptonSeparators == 0 || a.Ledger.NativeUniversalFamilyAxes == 0 || a.Ledger.GenerationAwareCandidates == 0 || a.Ledger.SourcesPassingAllRequirements != 0 || !a.Ledger.OnlySlotsNotOperators || !a.Ledger.YukawaEntriesSealed || !a.Ledger.NoObservedDataImported {
		return fmt.Errorf("Gate488 source ledger invalid: %+v", a.Ledger)
	}
	if !a.Requirements.Executed || !a.Requirements.RequiresUpDownSplit || !a.Requirements.RequiresGenerationAwareness || !a.Requirements.RequiresFamilyEigenbasis || !a.Requirements.RequiresNativeUpOperator || !a.Requirements.RequiresNativeDownOperator || !a.Requirements.RequiresNativeDiagonalizers || !a.Requirements.RequiresTwoInvariantConstraints || a.Requirements.CandidatesPassing != 0 || a.Requirements.NativeUpDownOperatorsDerived || a.Requirements.NativeDiagonalizersDerived || a.Requirements.CKMInvariantConstraintsDerived != 0 || a.Requirements.NativeCKMSourceFound {
		return fmt.Errorf("Gate488 requirement audit invalid: %+v", a.Requirements)
	}
	if !a.Socket.Executed || !a.Socket.UpDownSectorLabelsNative || !a.Socket.YukawaSlotsNative || a.Socket.YukawaMatrixValuesNative || a.Socket.FamilyEigenbasisNative || !a.Socket.CanNameOuOdSlots || a.Socket.CanPopulateOuOdNatively || a.Socket.CanComputeUuDaggerUd || a.Socket.CanComputeJarlskogInvariant || !a.Socket.BridgeAirlockRequired {
		return fmt.Errorf("Gate488 operator socket invalid: %+v", a.Socket)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedCKMImported || a.Firewall.ObservedWolfensteinImported || a.Firewall.ObservedQuarkMassesImported || a.Firewall.ObservedYukawaEntriesImported || a.Firewall.NativeUpOperatorWritten || a.Firewall.NativeDownOperatorWritten || a.Firewall.NativeDiagonalizersWritten || a.Firewall.CKMMatrixNativePrediction || a.Firewall.JarlskogNativePrediction || a.Firewall.CKMInvariantConstraintNativeWrite || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate488 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate488 finds the exact missing wall: ASHA natively names up/down representation slots and universal family axes, but no existing native source couples them into sector-specific 3x3 family operators. Therefore CKM orientation cannot be derived from the null cone, color, K_gen, or Higgs-edge topology alone; Yukawa matrices remain sealed bridge/environmental data until a new native selector theorem is found. Audited candidates=%d; full CKM-source candidates=%d; derived CKM invariant constraints=%d.", a.Ledger.CandidateCount, a.Ledger.SourcesPassingAllRequirements, a.Requirements.CKMInvariantConstraintsDerived)
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: Gate485=%t Gate486_bridge_only=%t Gate487_commutator_obstruction=%t required_constraints=%d derived_constraints=%d; %s", x.Verdict, x.Gate485NullC3BaselineInherited, x.Gate486NullMirrorBridgeOnly, x.Gate487CommutatorObstruction, x.Gate487RequiredConstraints, x.Gate487DerivedConstraints, x.Reason)
}

func FormatLedger(x SourceLedger) string {
	return fmt.Sprintf("%s: candidates=%d updown_label_sources=%d quark_lepton_separators=%d universal_family_axes=%d generation_aware=%d full_ckm_sources=%d; %s", x.Verdict, x.CandidateCount, x.NativeUpDownLabelSources, x.NativeQuarkLeptonSeparators, x.NativeUniversalFamilyAxes, x.GenerationAwareCandidates, x.SourcesPassingAllRequirements, x.Reason)
}

func FormatRequirements(x RequirementAudit) string {
	return fmt.Sprintf("%s: passing=%d native_operators=%t diagonalizers=%t invariant_constraints=%d/%d; %s", x.Verdict, x.CandidatesPassing, x.NativeUpDownOperatorsDerived, x.NativeDiagonalizersDerived, x.CKMInvariantConstraintsDerived, RequiredCKMInvariantConstraints, x.Reason)
}

func FormatSocket(x OperatorSocket) string {
	return fmt.Sprintf("%s: labels_native=%t slots_native=%t matrix_values_native=%t family_eigenbasis_native=%t can_name_OuOd=%t can_populate=%t airlock_required=%t; %s", x.Verdict, x.UpDownSectorLabelsNative, x.YukawaSlotsNative, x.YukawaMatrixValuesNative, x.FamilyEigenbasisNative, x.CanNameOuOdSlots, x.CanPopulateOuOdNatively, x.BridgeAirlockRequired, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_CKM=%t observed_Yukawa=%t native_Ou=%t native_Od=%t native_CKM=%t native_J=%t native_write=%t dim=%d KXY=%d; %s", x.Verdict, x.ObservedCKMImported, x.ObservedYukawaEntriesImported, x.NativeUpOperatorWritten, x.NativeDownOperatorWritten, x.CKMMatrixNativePrediction, x.JarlskogNativePrediction, x.NativeRegistryWritten, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 488 Registry Audit — Native Up/Down Operator Source Search\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusNativeUpDownSectorLabelsFound,
		StatusNativeUniversalFamilyAxisFound,
		StatusNoNativeUpDownEigenbasisSource,
		StatusNoNativeUpDownOperatorsDerived,
		StatusNoCKMInvariantConstraintsDerived,
		StatusYukawaMatricesRemainSealed,
		StatusCKMOrientationQuarantined,
		StatusFirewallBlockedNativeOperatorWrite,
	} {
		b.WriteString("- `" + v + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate485 gives a native null-C3 spectral baseline. Gate486 and Gate487 prove that this baseline does not by itself become a physical CKM theorem, because CKM lives in the relative up/down eigenbasis quotient.\n\n")
	b.WriteString("## Native source ledger\n\n")
	b.WriteString(FormatLedger(a.Ledger) + "\n\n")
	b.WriteString("| Candidate | Native layer | Up/down? | Generation-aware? | Family eigenbasis? | Native O_u/O_d? | Invariant constraints | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Ledger.Candidates {
		ops := c.SuppliesNativeUpOperator && c.SuppliesNativeDownOperator
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %t | %t | %d | `%s` |\n", c.Name, c.NativeLayer, c.DistinguishesUpDown, c.GenerationAware, c.SuppliesFamilyEigenbasis, ops, c.RephasingInvariantConstraintsProduced, c.Verdict))
	}
	b.WriteString("\n## Requirement sieve\n\n")
	b.WriteString(FormatRequirements(a.Requirements) + "\n\n")
	b.WriteString("A CKM-native source must satisfy all gates simultaneously: up/down split, generation awareness, family eigenbasis, native O_u, native O_d, native diagonalizers, and two rephasing-invariant polynomial constraints. No audited source passes.\n\n")
	b.WriteString("## Operator socket audit\n\n")
	b.WriteString(FormatSocket(a.Socket) + "\n\n")
	b.WriteString("The finite spectral triple can name Yukawa/operator sockets. It cannot populate their 3x3 family entries from native Clifford geometry at this gate. Therefore `V_CKM = U_u^† U_d`, `[O_u,O_d]`, and the Jarlskog invariant remain non-computable natively.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No observed CKM, Wolfenstein, quark mass, or Yukawa entry data were imported. No native CKM matrix, Jarlskog value, O_u/O_d matrix, diagonalizer, or invariant polynomial was written.\n\n")
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
