// Package spontaneouscarrierseal implements Gate 256:
// Spontaneous Carrier Seal / Gauge-Fixed Embedding Axiom Audit.
//
// Gate 255 proved that the scalar/contact operator Y_phi and the left-doublet
// operator T3L cannot be natively merged into the total Fock carrier S_C by a
// derived functor. Gate 256 therefore records the carrier merge as a seal: a
// quarantined spontaneous gauge-fixing boundary condition.  The package splits
// the result into two ledgers, per GateResearcherMethod.md:
//
//  1. the native-search obstruction inherited from Gate 255; and
//  2. the sealed conditional schema that says exactly which additional data
//     would be required before any so(8), triality, or neutral-kernel
//     computation is lawful.
//
// The seal does not insert Standard Model charge tables, choose a weak plane,
// select a Spin(8) triality branch by desired outcome, or force a 3-plane.
package spontaneouscarrierseal

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/carrierintertwiner"
)

const (
	AuditID = "GATE256-SPONTANEOUS-CARRIER-SEAL-GAUGE-FIXED-EMBEDDING-AXIOM-AUDIT"

	StatusGate255NativeNoGoInherited       = "CONDITIONAL_SUPPORT_GATE255_NATIVE_CARRIER_NO_GO_INHERITED"
	StatusSpontaneousCarrierSealInstituted = "CONDITIONAL_SUPPORT_SPONTANEOUS_CARRIER_SEAL_INSTITUTED"
	StatusConditionalIntertwinerSchema     = "CONDITIONAL_SUPPORT_CONDITIONAL_INTERTWINER_SCHEMA_DEFINED"
	StatusSymbolicFockLedgerSchema         = "CONDITIONAL_SUPPORT_SYMBOLIC_FOCK_LEDGER_SCHEMA_DEFINED"
	StatusSymbolicWittSO8Schema            = "CONDITIONAL_SUPPORT_SYMBOLIC_WITT_SO8_SCHEMA_AVAILABLE"
	StatusSealQuarantined                  = "CONDITIONAL_SUPPORT_SEAL_QUARANTINED_FROM_FINITE_CORE"
	StatusEmbeddingValuesMissing           = "FAILED_ROUTE_SEALED_EMBEDDING_VALUES_NOT_SUPPLIED"
	StatusConcreteEWLedgerBlocked          = "FAILED_ROUTE_CONCRETE_T3L_Y_PHI_FOCK_LEDGERS_STILL_BLOCKED"
	StatusTrialityBranchBlocked            = "FAILED_ROUTE_TRIALITY_BRANCH_SELECTION_STILL_BLOCKED"
	StatusQ8VCKernelBlocked                = "FAILED_ROUTE_Q8VC_KERNEL_COMPUTATION_STILL_BLOCKED"
	StatusNeutral3PlaneBlocked             = "FAILED_ROUTE_NEUTRAL_3PLANE_STILL_BLOCKED"
	StatusYukawaStillSealed                = "FAILED_ROUTE_YUKAWA_TEXTURE_STILL_SEALED"
)

type InheritedGate255Audit struct {
	Gate254Inherited          bool
	SCCarrierKnown            bool
	LocalActionsAudited       bool
	CommonCarrierDerived      bool
	CarrierIntertwinerDerived bool
	UnifiedLedgerConstructed  bool
	T3LYPhiSO8Coordinates     bool
	TrialityBranchSelected    bool
	Q8vCConstructed           bool
	Neutral3PlaneDerived      bool
	VTauConstructed           bool
	YukawaTextureDerived      bool
	Status                    string
	TruthStatement            string
}

type NativeSearchLedger struct {
	Object                     string
	Carrier                    string
	NativeStatus               string
	DerivedCommonSC            bool
	DerivedNumberLedger        bool
	DerivedSO8Coordinates      bool
	Obstruction                string
	SealAllowedAsBoundaryData  bool
	BoundaryDataChangesTheorem bool
	Verdict                    string
}

type NativeSearchAudit struct {
	Ledgers                           []NativeSearchLedger
	Gate255NoGoInherited              bool
	NativeCommonIntertwinerExists     bool
	NativeUnifiedLedgerExists         bool
	NativePhysicalSO8CoordinatesExist bool
	NativeTrialityPullbackAvailable   bool
	Verdict                           string
}

type SpontaneousCarrierSeal struct {
	Name                         string
	AxiomID                      string
	ConditionalStatus            string
	TargetCarrier                string
	PhysicalRole                 string
	ExplicitAxiom                bool
	Quarantined                  bool
	RequiredByGate255            bool
	GaugeFixingRequired          bool
	VacuumOrientationRequired    bool
	WeakFrameRequired            bool
	LeftDoubletInjectionRequired bool
	ScalarEmbeddingRequired      bool
	TrialityBranchRequired       bool
	DerivedFromFiniteGeometry    bool
	UsesObservedMasses           bool
	UsesObservedYukawas          bool
	UsesObservedGaugeCouplings   bool
	OverridesNativeNoGo          bool
	PollutesFiniteCore           bool
	Verdict                      string
}

type AxiomDatum struct {
	Name               string
	Symbol             string
	Required           bool
	Provided           bool
	Derived            bool
	Quarantined        bool
	NeededFor          string
	MissingReason      string
	AdmissibleProvider string
}

type ConditionalIntertwinerAxiom struct {
	Name                          string
	SourceCarriers                []string
	TargetCarrier                 string
	SchemaDefined                 bool
	OperationalIntertwinerBuilt   bool
	AxiomData                     []AxiomDatum
	RequiredDataCount             int
	ProvidedDataCount             int
	DerivedDataCount              int
	AllRequiredDataProvided       bool
	MapsT3LIntoSC                 bool
	MapsYPhiIntoSC                bool
	IntertwinesLocalActions       bool
	ProducesFourModeNumberLedgers bool
	ChangesCarrierByTensorProduct bool
	UsesDirectSumAsIntertwiner    bool
	Verdict                       string
}

type SealedFockLedger struct {
	Name                     string
	Expression               string
	CoefficientSymbols       []string
	NumberOperatorBasis      []string
	SymbolicLedgerDefined    bool
	NumericCoefficients      []float64
	NumericCoefficientsSet   bool
	NormalizationConstraints []string
	SealRequired             bool
	PhysicalMeaning          string
	Verdict                  string
}

type UnifiedLedgerAudit struct {
	Carrier                           string
	Ledgers                           []SealedFockLedger
	SymbolicSchemaConstructed         bool
	ConcreteT3LNumberLedgerAvailable  bool
	ConcreteYPhiNumberLedgerAvailable bool
	ConcreteQNumberLedgerAvailable    bool
	OperationalUnifiedLedgerBuilt     bool
	Obstruction                       string
	Verdict                           string
}

type SymbolicSO8Coordinate struct {
	Name                        string
	FockExpression              string
	SO8Formula                  string
	CoordinateSymbols           []string
	CartanBivectors             []string
	SymbolicFormulaAvailable    bool
	ConcreteCoordinateAvailable bool
	SealRequired                bool
	Verdict                     string
}

type SO8TranslationAudit struct {
	WittDictionaryInherited bool
	Coordinates             []SymbolicSO8Coordinate
	SymbolicSchemaAvailable bool
	ConcreteT3LSO8          bool
	ConcreteYPhiSO8         bool
	ConcreteQSO8            bool
	Obstruction             string
	Verdict                 string
}

type TrialityKernelAudit struct {
	TrialityCandidatesKnown      bool
	BranchSelectionSchemaDefined bool
	PhysicalBranchSelected       bool
	SelectedBranch               string
	SelectedByOutcome            bool
	Q8vCConstructed              bool
	EigensystemComputed          bool
	KernelDimensionKnown         bool
	KernelComplexDimension       int
	ExactlyThree                 bool
	NeutralThreePlaneDerived     bool
	DiagnosticEquationsRecorded  bool
	MissingConcreteInputs        []string
	Verdict                      string
}

type FirewallAudit struct {
	SealExplicitInput               bool
	SealQuarantined                 bool
	NativeNoGoPreserved             bool
	InventedEmbeddingValues         bool
	ImportedSMHyperchargeConvention bool
	ForcedWeakPlane                 bool
	SelectedTrialityByKernel        bool
	ForcedKernelDim3                bool
	TreatedSealAsFiniteDerivation   bool
	TreatedTensorProductAsSC        bool
	TreatedDirectSumAsIntertwiner   bool
	InsertedYukawaTexture           bool
	ImportedObservedMasses          bool
	PollutedFiniteCore              bool
	Verdict                         string
}

type DownstreamAudit struct {
	Neutral3PlaneAvailable bool
	TauEta                 []int
	VTauConstructed        bool
	TrialityTextureOpened  bool
	YukawaTextureDerived   bool
	CKMPMNSDerived         bool
	FermionMassesDerived   bool
	Verdict                string
}

type Summary struct {
	Gate255NoGoInherited          bool
	SpontaneousSealRecorded       bool
	ConditionalIntertwinerSchema  bool
	SealedAxiomValuesProvided     bool
	SymbolicLedgerSchemaAvailable bool
	ConcreteUnifiedLedgerBuilt    bool
	SymbolicSO8SchemaAvailable    bool
	ConcreteSO8Coordinates        bool
	TrialityBranchSelected        bool
	Q8vCConstructed               bool
	Neutral3PlaneDerived          bool
	YukawaTextureDerived          bool
	Status                        string
	NextGate                      string
	Comment                       string
}

type Analysis struct {
	PreviousGate255 InheritedGate255Audit
	NativeSearch    NativeSearchAudit
	Seal            SpontaneousCarrierSeal
	Intertwiner     ConditionalIntertwinerAxiom
	UnifiedLedger   UnifiedLedgerAudit
	SO8             SO8TranslationAudit
	TrialityKernel  TrialityKernelAudit
	Firewall        FirewallAudit
	Downstream      DownstreamAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prevRaw, err := carrierintertwiner.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 255 predecessor: %w", err)
			return
		}
		prev := inheritGate255(prevRaw)
		native := auditNativeSearch(prev, prevRaw)
		seal := defineSeal(native)
		intertwiner := defineConditionalIntertwiner(seal)
		ledger := defineUnifiedLedger(intertwiner)
		so8 := translateSymbolicSO8(prev, ledger)
		kernel := auditTrialityKernel(intertwiner, so8)
		fw := auditFirewall(native, seal, intertwiner, ledger, so8, kernel)
		down := auditDownstream(kernel)
		summary := summarize(native, seal, intertwiner, ledger, so8, kernel, down)
		truth := buildTruth(native, seal, intertwiner, ledger, so8, kernel)
		defaultA = Analysis{PreviousGate255: prev, NativeSearch: native, Seal: seal, Intertwiner: intertwiner, UnifiedLedger: ledger, SO8: so8, TrialityKernel: kernel, Firewall: fw, Downstream: down, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate255(a carrierintertwiner.Analysis) InheritedGate255Audit {
	return InheritedGate255Audit{
		Gate254Inherited:          a.Summary.Gate254Inherited,
		SCCarrierKnown:            a.Summary.SCCarrierKnown,
		LocalActionsAudited:       a.Summary.LocalActionsAudited,
		CommonCarrierDerived:      a.Summary.CommonCarrierDerived,
		CarrierIntertwinerDerived: a.Summary.CarrierIntertwinerDerived,
		UnifiedLedgerConstructed:  a.Summary.UnifiedLedgerConstructed,
		T3LYPhiSO8Coordinates:     a.Summary.T3LYPhiSO8Coordinates,
		TrialityBranchSelected:    a.Summary.TrialityBranchSelected,
		Q8vCConstructed:           a.Summary.Q8vCConstructed,
		Neutral3PlaneDerived:      a.Summary.Neutral3PlaneDerived,
		VTauConstructed:           a.Summary.VTauConstructed,
		YukawaTextureDerived:      a.Summary.YukawaTextureDerived,
		Status:                    a.Summary.Status,
		TruthStatement:            a.TruthStatement,
	}
}

func auditNativeSearch(prev InheritedGate255Audit, gate255 carrierintertwiner.Analysis) NativeSearchAudit {
	ledgers := []NativeSearchLedger{
		{
			Object:                     "T3L left-doublet action",
			Carrier:                    "Q_L⊕L_L local left-doublet carrier",
			NativeStatus:               gate255.Carriers.ObjectsAudited[1].Verdict,
			DerivedCommonSC:            false,
			DerivedNumberLedger:        false,
			DerivedSO8Coordinates:      false,
			Obstruction:                gate255.Carriers.ObjectsAudited[1].Obstruction,
			SealAllowedAsBoundaryData:  true,
			BoundaryDataChangesTheorem: false,
			Verdict:                    "may be related to S_C only after an explicit gauge-fixed state-index injection is supplied as sealed boundary data",
		},
		{
			Object:                     "Y_phi scalar/contact action",
			Carrier:                    "H_phi scalar/contact carrier",
			NativeStatus:               gate255.Carriers.ObjectsAudited[2].Verdict,
			DerivedCommonSC:            false,
			DerivedNumberLedger:        false,
			DerivedSO8Coordinates:      false,
			Obstruction:                gate255.Carriers.ObjectsAudited[2].Obstruction,
			SealAllowedAsBoundaryData:  true,
			BoundaryDataChangesTheorem: false,
			Verdict:                    "may be related to S_C only after a scalar trivialization and gauge orientation are supplied as sealed boundary data",
		},
	}
	return NativeSearchAudit{
		Ledgers:                           ledgers,
		Gate255NoGoInherited:              prev.Gate254Inherited && prev.SCCarrierKnown && prev.LocalActionsAudited && !prev.CommonCarrierDerived && !prev.UnifiedLedgerConstructed,
		NativeCommonIntertwinerExists:     false,
		NativeUnifiedLedgerExists:         false,
		NativePhysicalSO8CoordinatesExist: false,
		NativeTrialityPullbackAvailable:   false,
		Verdict:                           "Gate 256 preserves the Gate 255 native no-go: no finite-core functor embeds both local electroweak observables into S_C without a spontaneous carrier seal.",
	}
}

func defineSeal(native NativeSearchAudit) SpontaneousCarrierSeal {
	return SpontaneousCarrierSeal{
		Name:                         "SpontaneousCarrierSeal",
		AxiomID:                      "SEAL-SSB-CARRIER-GAUGE-FIXED-SC-EMBEDDING",
		ConditionalStatus:            StatusSpontaneousCarrierSealInstituted,
		TargetCarrier:                "S_C = Λ*(C^4)",
		PhysicalRole:                 "records that scalar/contact and left-doublet observables can be compared on S_C only after a vacuum/gauge frame is chosen",
		ExplicitAxiom:                native.Gate255NoGoInherited,
		Quarantined:                  true,
		RequiredByGate255:            native.Gate255NoGoInherited,
		GaugeFixingRequired:          true,
		VacuumOrientationRequired:    true,
		WeakFrameRequired:            true,
		LeftDoubletInjectionRequired: true,
		ScalarEmbeddingRequired:      true,
		TrialityBranchRequired:       true,
		DerivedFromFiniteGeometry:    false,
		UsesObservedMasses:           false,
		UsesObservedYukawas:          false,
		UsesObservedGaugeCouplings:   false,
		OverridesNativeNoGo:          false,
		PollutesFiniteCore:           false,
		Verdict:                      "seal instituted as an explicit SSB/gauge-fixing boundary condition; it permits a future conditional embedding but does not itself provide the missing numerical or discrete ledger values",
	}
}

func defineConditionalIntertwiner(seal SpontaneousCarrierSeal) ConditionalIntertwinerAxiom {
	data := []AxiomDatum{
		{Name: "scalar/contact trivialization", Symbol: "ι_phi:H_phi→S_C", Required: true, Provided: false, Derived: false, Quarantined: true, NeededFor: "Y_phi common-carrier action", MissingReason: "Gate 255 found no native H_phi→S_C embedding", AdmissibleProvider: "SpontaneousCarrierSeal branch data"},
		{Name: "left-doublet occupation injection", Symbol: "ι_L:Q_L⊕L_L→S_C", Required: true, Provided: false, Derived: false, Quarantined: true, NeededFor: "T3L common-carrier action", MissingReason: "Gate 255 found no native state-index map into the sixteen Fock occupations", AdmissibleProvider: "SpontaneousCarrierSeal branch data"},
		{Name: "weak SU(2) frame / plane", Symbol: "U_L⊂{N_0,N_1,N_2,N_3}", Required: true, Provided: false, Derived: false, Quarantined: true, NeededFor: "T3L coefficient vector", MissingReason: "candidate weak Cartans exist, but no physical plane selector is derived", AdmissibleProvider: "gauge-fixed frame axiom or later finite selector theorem"},
		{Name: "Higgs/scalar charge orientation", Symbol: "Y_phi^seal", Required: true, Provided: false, Derived: false, Quarantined: true, NeededFor: "Y_phi coefficient vector", MissingReason: "scalar/contact hypercharge has not been identified with a four-mode Fock ledger", AdmissibleProvider: "gauge-fixed scalar orientation axiom or later finite selector theorem"},
		{Name: "spinor-to-vector triality branch", Symbol: "τ_{s→v}", Required: true, Provided: false, Derived: false, Quarantined: true, NeededFor: "Q_8vC construction", MissingReason: "Gate 253/254/255 audited τ_even/τ_odd risk but selected no branch", AdmissibleProvider: "representation-weight theorem or explicit branch seal"},
	}
	provided, derived := 0, 0
	for _, d := range data {
		if d.Provided {
			provided++
		}
		if d.Derived {
			derived++
		}
	}
	all := provided == len(data)
	return ConditionalIntertwinerAxiom{
		Name:                          "conditional gauge-fixed carrier intertwiner schema",
		SourceCarriers:                []string{"H_phi", "Q_L⊕L_L"},
		TargetCarrier:                 seal.TargetCarrier,
		SchemaDefined:                 seal.ExplicitAxiom && seal.Quarantined,
		OperationalIntertwinerBuilt:   all,
		AxiomData:                     data,
		RequiredDataCount:             len(data),
		ProvidedDataCount:             provided,
		DerivedDataCount:              derived,
		AllRequiredDataProvided:       all,
		MapsT3LIntoSC:                 all,
		MapsYPhiIntoSC:                all,
		IntertwinesLocalActions:       all,
		ProducesFourModeNumberLedgers: all,
		ChangesCarrierByTensorProduct: false,
		UsesDirectSumAsIntertwiner:    false,
		Verdict:                       "the lawful conditional schema is now explicit, but no operational intertwiner is built until every sealed embedding datum is supplied or derived",
	}
}

func defineUnifiedLedger(ax ConditionalIntertwinerAxiom) UnifiedLedgerAudit {
	basis := []string{"N_0", "N_1", "N_2", "N_3"}
	ledgers := []SealedFockLedger{
		{
			Name:                     "T3L^seal",
			Expression:               "Σ_k t_k N_k on S_C after ι_L is supplied",
			CoefficientSymbols:       []string{"t_0", "t_1", "t_2", "t_3"},
			NumberOperatorBasis:      basis,
			SymbolicLedgerDefined:    ax.SchemaDefined,
			NumericCoefficientsSet:   false,
			NormalizationConstraints: []string{"must reproduce the local left-doublet T3L spectrum on im(ι_L)", "must preserve the chosen weak SU(2) frame", "must not be inferred from T3R or B-L alone"},
			SealRequired:             true,
			PhysicalMeaning:          "sealed common-carrier representative of the left weak Cartan",
			Verdict:                  "symbolic ledger schema only; concrete t_k are not supplied by Gate 256",
		},
		{
			Name:                     "Y_phi^seal",
			Expression:               "Σ_k y_k N_k on S_C after ι_phi is supplied",
			CoefficientSymbols:       []string{"y_0", "y_1", "y_2", "y_3"},
			NumberOperatorBasis:      basis,
			SymbolicLedgerDefined:    ax.SchemaDefined,
			NumericCoefficientsSet:   false,
			NormalizationConstraints: []string{"must reproduce the scalar/contact Y_phi action on im(ι_phi)", "must use the sealed scalar orientation", "must not be imported as an untyped Standard Model convention"},
			SealRequired:             true,
			PhysicalMeaning:          "sealed common-carrier representative of the scalar hypercharge/contact operator",
			Verdict:                  "symbolic ledger schema only; concrete y_k are not supplied by Gate 256",
		},
		{
			Name:                     "Q^seal=T3L^seal+Y_phi^seal",
			Expression:               "Σ_k (t_k+y_k) N_k",
			CoefficientSymbols:       []string{"t_0+y_0", "t_1+y_1", "t_2+y_2", "t_3+y_3"},
			NumberOperatorBasis:      basis,
			SymbolicLedgerDefined:    ax.SchemaDefined,
			NumericCoefficientsSet:   false,
			NormalizationConstraints: []string{"requires both concrete T3L and Y_phi ledgers", "must be evaluated only after the triality branch is selected"},
			SealRequired:             true,
			PhysicalMeaning:          "sealed electromagnetic generator candidate on common S_C carrier",
			Verdict:                  "formal sum is typed, but no numeric Q_8vC matrix is available",
		},
	}
	return UnifiedLedgerAudit{
		Carrier:                           ax.TargetCarrier,
		Ledgers:                           ledgers,
		SymbolicSchemaConstructed:         ax.SchemaDefined,
		ConcreteT3LNumberLedgerAvailable:  false,
		ConcreteYPhiNumberLedgerAvailable: false,
		ConcreteQNumberLedgerAvailable:    false,
		OperationalUnifiedLedgerBuilt:     false,
		Obstruction:                       "the SpontaneousCarrierSeal names the required boundary data but does not supply t_k, y_k, state injections, or a weak plane; the unified ledger remains symbolic and conditional",
		Verdict:                           "Gate 256 upgrades the problem from an illegal carrier merge to an explicit sealed ledger schema, not to a concrete electroweak Fock ledger",
	}
}

func translateSymbolicSO8(prev InheritedGate255Audit, ledger UnifiedLedgerAudit) SO8TranslationAudit {
	bivs := []string{"e0∧e1", "e2∧e3", "e4∧e5", "e6∧e7"}
	coords := make([]SymbolicSO8Coordinate, 0, len(ledger.Ledgers))
	for _, l := range ledger.Ledgers {
		terms := make([]string, 0, 4)
		for i, c := range l.CoefficientSymbols {
			terms = append(terms, fmt.Sprintf("(i/2)(%s)%s", c, bivs[i]))
		}
		coords = append(coords, SymbolicSO8Coordinate{
			Name:                        l.Name,
			FockExpression:              l.Expression,
			SO8Formula:                  strings.Join(terms, " + "),
			CoordinateSymbols:           append([]string(nil), l.CoefficientSymbols...),
			CartanBivectors:             append([]string(nil), bivs...),
			SymbolicFormulaAvailable:    ledger.SymbolicSchemaConstructed,
			ConcreteCoordinateAvailable: l.NumericCoefficientsSet,
			SealRequired:                true,
			Verdict:                     "Witt dictionary gives the formal Cartan expression; concrete so(8) coordinates still require sealed coefficient values",
		})
	}
	concrete := false
	return SO8TranslationAudit{
		WittDictionaryInherited: prev.SCCarrierKnown,
		Coordinates:             coords,
		SymbolicSchemaAvailable: ledger.SymbolicSchemaConstructed && prev.SCCarrierKnown,
		ConcreteT3LSO8:          concrete,
		ConcreteYPhiSO8:         concrete,
		ConcreteQSO8:            concrete,
		Obstruction:             "symbolic Witt translation is available, but no concrete numeric coefficient vector exists for T3L^seal, Y_phi^seal, or Q^seal",
		Verdict:                 "the seal makes the formula type-correct while preserving the numerical/branch firewall",
	}
}

func auditTrialityKernel(ax ConditionalIntertwinerAxiom, so8 SO8TranslationAudit) TrialityKernelAudit {
	missing := []string{}
	for _, d := range ax.AxiomData {
		if d.Required && !d.Provided {
			missing = append(missing, d.Symbol)
		}
	}
	return TrialityKernelAudit{
		TrialityCandidatesKnown:      true,
		BranchSelectionSchemaDefined: true,
		PhysicalBranchSelected:       false,
		SelectedBranch:               "",
		SelectedByOutcome:            false,
		Q8vCConstructed:              false,
		EigensystemComputed:          false,
		KernelDimensionKnown:         false,
		KernelComplexDimension:       0,
		ExactlyThree:                 false,
		NeutralThreePlaneDerived:     false,
		DiagnosticEquationsRecorded:  so8.SymbolicSchemaAvailable,
		MissingConcreteInputs:        missing,
		Verdict:                      "no Q_8vC eigensystem is computed because the triality branch and concrete sealed coefficients are absent; the 3-plane remains a target condition, not a result",
	}
}

func auditFirewall(native NativeSearchAudit, seal SpontaneousCarrierSeal, ax ConditionalIntertwinerAxiom, ledger UnifiedLedgerAudit, so8 SO8TranslationAudit, kernel TrialityKernelAudit) FirewallAudit {
	return FirewallAudit{
		SealExplicitInput:               seal.ExplicitAxiom,
		SealQuarantined:                 seal.Quarantined,
		NativeNoGoPreserved:             native.Gate255NoGoInherited && !native.NativeCommonIntertwinerExists && !seal.OverridesNativeNoGo,
		InventedEmbeddingValues:         ax.ProvidedDataCount > 0,
		ImportedSMHyperchargeConvention: false,
		ForcedWeakPlane:                 false,
		SelectedTrialityByKernel:        kernel.SelectedByOutcome,
		ForcedKernelDim3:                kernel.ExactlyThree && !kernel.KernelDimensionKnown,
		TreatedSealAsFiniteDerivation:   seal.DerivedFromFiniteGeometry,
		TreatedTensorProductAsSC:        ax.ChangesCarrierByTensorProduct,
		TreatedDirectSumAsIntertwiner:   ax.UsesDirectSumAsIntertwiner,
		InsertedYukawaTexture:           false,
		ImportedObservedMasses:          seal.UsesObservedMasses,
		PollutedFiniteCore:              seal.PollutesFiniteCore || so8.ConcreteQSO8 || ledger.OperationalUnifiedLedgerBuilt,
		Verdict:                         "Gate 256 records a conditional seal without changing the native theorem status or importing Standard Model tables, masses, couplings, Yukawas, or a desired kernel dimension",
	}
}

func auditDownstream(kernel TrialityKernelAudit) DownstreamAudit {
	return DownstreamAudit{
		Neutral3PlaneAvailable: kernel.NeutralThreePlaneDerived,
		TauEta:                 []int{2, -2, 1},
		VTauConstructed:        false,
		TrialityTextureOpened:  false,
		YukawaTextureDerived:   false,
		CKMPMNSDerived:         false,
		FermionMassesDerived:   false,
		Verdict:                "tau_eta remains available as a downstream generation-breaking trace, but it is not applied because the neutral 3-plane is still unavailable",
	}
}

func summarize(native NativeSearchAudit, seal SpontaneousCarrierSeal, ax ConditionalIntertwinerAxiom, ledger UnifiedLedgerAudit, so8 SO8TranslationAudit, kernel TrialityKernelAudit, down DownstreamAudit) Summary {
	return Summary{
		Gate255NoGoInherited:          native.Gate255NoGoInherited,
		SpontaneousSealRecorded:       seal.ExplicitAxiom && seal.Quarantined,
		ConditionalIntertwinerSchema:  ax.SchemaDefined,
		SealedAxiomValuesProvided:     ax.AllRequiredDataProvided,
		SymbolicLedgerSchemaAvailable: ledger.SymbolicSchemaConstructed,
		ConcreteUnifiedLedgerBuilt:    ledger.OperationalUnifiedLedgerBuilt,
		SymbolicSO8SchemaAvailable:    so8.SymbolicSchemaAvailable,
		ConcreteSO8Coordinates:        so8.ConcreteT3LSO8 && so8.ConcreteYPhiSO8 && so8.ConcreteQSO8,
		TrialityBranchSelected:        kernel.PhysicalBranchSelected,
		Q8vCConstructed:               kernel.Q8vCConstructed,
		Neutral3PlaneDerived:          kernel.NeutralThreePlaneDerived,
		YukawaTextureDerived:          down.YukawaTextureDerived,
		Status:                        StatusEmbeddingValuesMissing,
		NextGate:                      "Gate 257 — Sealed Carrier Embedding Data / Weak-Frame and Triality-Branch Witness Audit",
		Comment:                       "The SpontaneousCarrierSeal is now explicit and quarantined. It defines the lawful schema for a gauge-fixed embedding into S_C, but it does not supply the concrete state injections, weak frame, charge coefficients, or triality branch needed to compute Q_8vC.",
	}
}

func buildTruth(native NativeSearchAudit, seal SpontaneousCarrierSeal, ax ConditionalIntertwinerAxiom, ledger UnifiedLedgerAudit, so8 SO8TranslationAudit, kernel TrialityKernelAudit) string {
	return fmt.Sprintf("Gate 256 preserves the native Gate-255 no-go (%t) and institutes %s as explicit quarantined SSB boundary data. The seal defines a typed conditional schema H_phi,Q_L⊕L_L -> %s and symbolic ledgers T3L=Σt_kN_k, Y_phi=Σy_kN_k; the Witt dictionary then gives symbolic so(8) Cartan formulas. However %d/%d required axiom data are supplied, so no concrete T3L/Y_phi coordinates, no selected τ_{s→v}, no Q_8vC eigensystem, and no neutral 3-plane are derived.", native.Gate255NoGoInherited, seal.Name, seal.TargetCarrier, ax.ProvidedDataCount, ax.RequiredDataCount)
}
