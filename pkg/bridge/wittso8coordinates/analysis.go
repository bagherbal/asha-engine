// Package wittso8coordinates implements Gate 253:
// Witt Decomposition / Fock-to-so(8) Bivector Coordinate Audit.
//
// Gate 252 correctly refused to push representation labels through
// infinitesimal triality: T3L and Y_phi were known only as bridge names, not as
// typed coordinates in so(8)=Λ²R⁸.  Gate 253 reads the native four-mode Fock
// dictionary backwards.  It exposes the Cartan bivectors carried by the number
// operators N_k and audits exactly what this unlocks and exactly what remains
// blocked.
package wittso8coordinates

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/lietrialitypullback"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE253-WITT-DECOMPOSITION-FOCK-TO-SO8-BIVECTOR-COORDINATE-AUDIT"

	StatusWittPairingRetrieved             = "CONDITIONAL_SUPPORT_NATIVE_WITT_PAIRING_RETRIEVED"
	StatusNumberOperatorsAsCartanBivectors = "CONDITIONAL_SUPPORT_NUMBER_OPERATORS_HAVE_SO8_CARTAN_COORDINATES"
	StatusKnownFockLedgersCoordinateReady  = "CONDITIONAL_SUPPORT_KNOWN_FOCK_NUMBER_LEDGERS_COORDINATE_READY"
	StatusTrialityCartanCandidatesAudited  = "CONDITIONAL_SUPPORT_D4_CARTAN_TRIALITY_CANDIDATES_AUDITED"
	StatusT3LYPhiLedgerMissing             = "FAILED_ROUTE_T3L_Y_PHI_NUMBER_OPERATOR_LEDGER_MISSING"
	StatusExplicitTrialityStillUnselected  = "FAILED_ROUTE_EXPLICIT_SPINOR_TO_VECTOR_TRIALITY_AUTOMORPHISM_STILL_UNSELECTED"
	StatusQ8VCStillBlocked                 = "FAILED_ROUTE_Q8VC_CONSTRUCTION_STILL_BLOCKED"
	StatusNeutral3PlaneStillBlocked        = "FAILED_ROUTE_NEUTRAL_3PLANE_STILL_BLOCKED"
	StatusVTauStillBlocked                 = "FAILED_ROUTE_V_TAU_CONSTRUCTION_STILL_BLOCKED"
	StatusYukawaStillBlocked               = "FAILED_ROUTE_YUKAWA_TEXTURE_DERIVATION_STILL_BLOCKED"
)

type InheritedGate252Audit struct {
	InfinitesimalTrialityCapacity bool
	SpinorEWBridgeKnown           bool
	SpinorSO8Coordinates          bool
	ExplicitLieTrialityMap        bool
	VectorEWMatriciesDerived      bool
	Q8vCConstructed               bool
	Neutral3PlaneDerived          bool
	JCompatibleTransport          bool
	VTauConstructed               bool
	TrialityUnblocked             bool
	YukawaTextureDerived          bool
	TruthStatement                string
}

type WittPairAudit struct {
	ModeIndex           int
	CreationName        string
	AnnihilationName    string
	Kind                string
	RealPlane           string
	Bivector            string
	CreationFormula     string
	AnnihilationFormula string
	NativePairing       bool
}

type WittBasisAudit struct {
	SourceGate           string
	RealDimension        int
	ComplexModeCount     int
	Pairs                []WittPairAudit
	AllPairsNative       bool
	TemporalSpatialSplit string
	Convention           string
	Retrieved            bool
	Verdict              string
}

type NumberOperatorCoordinate struct {
	ModeIndex              int
	NumberOperator         string
	CentralIdentityShift   float64
	ImaginaryBivectorCoeff float64
	Bivector               string
	LieCoordinateFormula   string
	CentralShiftRemoved    bool
}

type NumberOperatorExpansionAudit struct {
	Formula                   string
	Coordinates               []NumberOperatorCoordinate
	CoordinateCount           int
	CartanBivectors           []string
	MaximalTorusDimension     int
	AllPureBivectorAfterShift bool
	CentralPartRejectedBySO8  bool
	Derived                   bool
	Verdict                   string
}

type FockLedgerCoordinate struct {
	Name                 string
	Expression           string
	NumberCoefficients   []float64
	CentralIdentityShift float64
	BivectorCoefficients []float64
	BivectorFormula      string
	CoordinateDerived    bool
	PhysicalEWGenerator  bool
	Verdict              string
}

type KnownFockLedgerAudit struct {
	Ledgers                      []FockLedgerCoordinate
	BMinusLCoordinatesDerived    bool
	TemporalT0CoordinatesDerived bool
	WeakPlaneCandidateDerived    bool
	AllDerivedFromNumberOps      bool
	Verdict                      string
}

type ElectroweakGeneratorCoordinateAudit struct {
	RequestedGenerators            []string
	T3LBridgeNameKnown             bool
	YPhiBridgeNameKnown            bool
	T3LNumberOperatorCoefficients  bool
	YPhiNumberOperatorCoefficients bool
	T3LSO8CoordinatesDerived       bool
	YPhiSO8CoordinatesDerived      bool
	QSO8CoordinatesDerived         bool
	ZSO8CoordinatesDerived         bool
	CandidateFockLedgersAvailable  []string
	Obstruction                    string
	Verdict                        string
}

type TrialityCartanCandidate struct {
	Name              string
	Matrix            [][]float64
	Orthogonal        bool
	Involutive        bool
	DetAbsOne         bool
	MapsD4RootLattice bool
	Selected          bool
	Verdict           string
}

type TrialityAudit struct {
	OuterAutomorphismGroup              string
	CandidateCount                      int
	Candidates                          []TrialityCartanCandidate
	SpecificSpinorToVectorChoiceDerived bool
	UsesWrongChoiceRiskAudited          bool
	CanApplyToPhysicalEW                bool
	Obstruction                         string
	Verdict                             string
}

type Q8VCKernelAudit struct {
	Definition               string
	T3LCoordinatesAvailable  bool
	YPhiCoordinatesAvailable bool
	TrialityChoiceAvailable  bool
	Q8vCConstructed          bool
	EigensystemComputed      bool
	KernelDimensionKnown     bool
	KernelComplexDimension   int
	ExactlyThree             bool
	ThreePlaneDerived        bool
	DiagnosticOnlyReason     string
	Verdict                  string
}

type DownstreamAudit struct {
	Neutral3PlaneAvailable bool
	TauEta                 []int
	VTauConstructed        bool
	YukawaTextureDerived   bool
	CKMPMNSDerived         bool
	FermionMassesDerived   bool
	Verdict                string
}

type FirewallAudit struct {
	InventedWittPairing       bool
	InventedT3LCoefficients   bool
	InventedYPhiCoefficients  bool
	SelectedTrialityByOutcome bool
	ForcedKernelDim3          bool
	ConstructedVTauByHand     bool
	InsertedYukawaTexture     bool
	ImportedObservedMasses    bool
	PollutedFiniteCore        bool
	Verdict                   string
}

type Summary struct {
	WittPairingRetrieved            bool
	NumberSO8Coordinates            bool
	KnownFockLedgersCoordinateReady bool
	T3LYPhiSO8Coordinates           bool
	ExplicitTrialitySelected        bool
	Q8vCConstructed                 bool
	Neutral3PlaneDerived            bool
	VTauConstructed                 bool
	YukawaTextureDerived            bool
	Status                          string
	NextGate                        string
	Comment                         string
}

type Analysis struct {
	PreviousGate252 InheritedGate252Audit
	WittBasis       WittBasisAudit
	NumberOperators NumberOperatorExpansionAudit
	FockLedgers     KnownFockLedgerAudit
	Electroweak     ElectroweakGeneratorCoordinateAudit
	Triality        TrialityAudit
	Kernel          Q8VCKernelAudit
	Downstream      DownstreamAudit
	Firewall        FirewallAudit
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
		prevRaw, err := lietrialitypullback.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate252(prevRaw)
		witt, err := auditWittBasis()
		if err != nil {
			defaultErr = err
			return
		}
		numbers := auditNumberOperators(witt)
		ledgers := auditKnownFockLedgers(numbers)
		ew := auditElectroweakCoordinates(prev, ledgers)
		triality := auditTriality(ew)
		kernel := auditKernel(ew, triality)
		down := auditDownstream(kernel)
		fw := auditFirewall()
		summary := summarize(witt, numbers, ledgers, ew, triality, kernel, down)
		truth := buildTruth(prev, witt, numbers, ledgers, ew, triality, kernel)
		defaultA = Analysis{PreviousGate252: prev, WittBasis: witt, NumberOperators: numbers, FockLedgers: ledgers, Electroweak: ew, Triality: triality, Kernel: kernel, Downstream: down, Firewall: fw, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate252(a lietrialitypullback.Analysis) InheritedGate252Audit {
	return InheritedGate252Audit{
		InfinitesimalTrialityCapacity: a.Summary.InfinitesimalTrialityCapacity,
		SpinorEWBridgeKnown:           a.Summary.SpinorEWBridgeKnown,
		SpinorSO8Coordinates:          a.Summary.SpinorSO8Coordinates,
		ExplicitLieTrialityMap:        a.Summary.ExplicitLieTrialityMap,
		VectorEWMatriciesDerived:      a.Summary.VectorEWMatriciesDerived,
		Q8vCConstructed:               a.Summary.Q8vCConstructed,
		Neutral3PlaneDerived:          a.Summary.Neutral3PlaneDerived,
		JCompatibleTransport:          a.Summary.JCompatibleTransport,
		VTauConstructed:               a.Summary.VTauConstructed,
		TrialityUnblocked:             a.Summary.TrialityUnblocked,
		YukawaTextureDerived:          a.Summary.YukawaTextureDerived,
		TruthStatement:                a.TruthStatement,
	}
}

func auditWittBasis() (WittBasisAudit, error) {
	w, err := spinor.NativeWittDecomposition(4)
	if err != nil {
		return WittBasisAudit{}, err
	}
	pairs := make([]WittPairAudit, 0, len(w.Pairs))
	allNative := true
	temporal, spatial := 0, 0
	for _, p := range w.Pairs {
		native := p.RealBasisA == 2*p.ModeIndex && p.RealBasisB == 2*p.ModeIndex+1
		allNative = allNative && native
		if p.Kind == spinor.TemporalMode {
			temporal++
		}
		if p.Kind == spinor.SpatialMode {
			spatial++
		}
		pairs = append(pairs, WittPairAudit{
			ModeIndex:           p.ModeIndex,
			CreationName:        p.CreationName,
			AnnihilationName:    p.AnnihilationName,
			Kind:                string(p.Kind),
			RealPlane:           fmt.Sprintf("span{e%d,e%d}", p.RealBasisA, p.RealBasisB),
			Bivector:            p.BivectorLabel,
			CreationFormula:     p.CreationFormula,
			AnnihilationFormula: p.AnnihilationFormula,
			NativePairing:       native,
		})
	}
	return WittBasisAudit{
		SourceGate:           "Gate 14 / SPINOR-WITT-FOCK-16 native four-mode Fock bookkeeping, now made explicit as a pair dictionary",
		RealDimension:        w.RealDimension,
		ComplexModeCount:     w.ComplexModeCount,
		Pairs:                pairs,
		AllPairsNative:       allNative,
		TemporalSpatialSplit: fmt.Sprintf("%d temporal + %d spatial modes", temporal, spatial),
		Convention:           w.Convention,
		Retrieved:            w.RealDimension == 8 && w.ComplexModeCount == 4 && allNative,
		Verdict:              "the engine now has an explicit native dictionary from Fock mode k to Cartan bivector e_{2k}∧e_{2k+1}; this is data, not a label guess",
	}, nil
}

func auditNumberOperators(w WittBasisAudit) NumberOperatorExpansionAudit {
	coords := make([]NumberOperatorCoordinate, 0, len(w.Pairs))
	bivs := make([]string, 0, len(w.Pairs))
	pure := w.Retrieved
	for _, p := range w.Pairs {
		coords = append(coords, NumberOperatorCoordinate{
			ModeIndex:              p.ModeIndex,
			NumberOperator:         fmt.Sprintf("N_%d=a†_%d a_%d", p.ModeIndex, p.ModeIndex, p.ModeIndex),
			CentralIdentityShift:   0.5,
			ImaginaryBivectorCoeff: 0.5,
			Bivector:               p.Bivector,
			LieCoordinateFormula:   fmt.Sprintf("N_%d - 1/2 I ↦ (i/2)%s", p.ModeIndex, p.Bivector),
			CentralShiftRemoved:    true,
		})
		bivs = append(bivs, p.Bivector)
	}
	return NumberOperatorExpansionAudit{
		Formula:                   "N_k = 1/2 I + (i/2) e_{2k}∧e_{2k+1}; only N_k-1/2 I is an so(8) Cartan coordinate",
		Coordinates:               coords,
		CoordinateCount:           len(coords),
		CartanBivectors:           bivs,
		MaximalTorusDimension:     4,
		AllPureBivectorAfterShift: pure && len(coords) == 4,
		CentralPartRejectedBySO8:  true,
		Derived:                   pure && len(coords) == 4,
		Verdict:                   "all four diagonal number operators have explicit coordinates in the Cartan torus of so(8) after removing the central identity shift",
	}
}

func auditKnownFockLedgers(n NumberOperatorExpansionAudit) KnownFockLedgerAudit {
	ledgers := []FockLedgerCoordinate{
		makeLedger("B-L", "-N_0 + (1/3)(N_1+N_2+N_3)", []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}, n, false),
		makeLedger("T0 temporal polarization", "1/2 I - N_0", []float64{-1, 0, 0, 0}, n, false),
		makeLedger("conditional weak-plane Cartan T3_U12", "(1/2)(N_1-N_2) if U={a†_1,a†_2} is later lawfully selected", []float64{0, 0.5, -0.5, 0}, n, false),
	}
	return KnownFockLedgerAudit{
		Ledgers:                      ledgers,
		BMinusLCoordinatesDerived:    ledgers[0].CoordinateDerived,
		TemporalT0CoordinatesDerived: ledgers[1].CoordinateDerived,
		WeakPlaneCandidateDerived:    ledgers[2].CoordinateDerived,
		AllDerivedFromNumberOps:      ledgers[0].CoordinateDerived && ledgers[1].CoordinateDerived && ledgers[2].CoordinateDerived,
		Verdict:                      "the Witt dictionary now coordinates any audited number-operator ledger, including B-L, temporal T0, and a conditional weak-plane Cartan candidate; these are not promoted to physical T3L/Y_phi without their own coefficient ledger",
	}
}

func makeLedger(name, expr string, coeffs []float64, n NumberOperatorExpansionAudit, physicalEW bool) FockLedgerCoordinate {
	central := 0.0
	biv := make([]float64, len(coeffs))
	for i, c := range coeffs {
		central += 0.5 * c
		biv[i] = 0.5 * c
	}
	return FockLedgerCoordinate{
		Name:                 name,
		Expression:           expr,
		NumberCoefficients:   append([]float64(nil), coeffs...),
		CentralIdentityShift: central,
		BivectorCoefficients: biv,
		BivectorFormula:      formatBivectorFormula(biv, n.CartanBivectors),
		CoordinateDerived:    n.Derived && len(coeffs) == n.CoordinateCount,
		PhysicalEWGenerator:  physicalEW,
		Verdict:              "coordinate is derived as a Cartan so(8) bivector for this number-operator expression; physical interpretation remains whatever the source ledger proves",
	}
}

func auditElectroweakCoordinates(prev InheritedGate252Audit, ledgers KnownFockLedgerAudit) ElectroweakGeneratorCoordinateAudit {
	candidateNames := make([]string, 0, len(ledgers.Ledgers))
	for _, l := range ledgers.Ledgers {
		if l.CoordinateDerived {
			candidateNames = append(candidateNames, l.Name)
		}
	}
	return ElectroweakGeneratorCoordinateAudit{
		RequestedGenerators:            []string{"T3L", "Y_phi", "Q=T3L+Y_phi", "Z=T3L-Y_phi"},
		T3LBridgeNameKnown:             prev.SpinorEWBridgeKnown,
		YPhiBridgeNameKnown:            prev.SpinorEWBridgeKnown,
		T3LNumberOperatorCoefficients:  false,
		YPhiNumberOperatorCoefficients: false,
		T3LSO8CoordinatesDerived:       false,
		YPhiSO8CoordinatesDerived:      false,
		QSO8CoordinatesDerived:         false,
		ZSO8CoordinatesDerived:         false,
		CandidateFockLedgersAvailable:  candidateNames,
		Obstruction:                    "the project now knows how to coordinate any diagonal Fock number-operator expression, but it still does not contain a native theorem identifying the bridge-level T3L and scalar/contact Y_phi with concrete coefficient vectors over (N_0,N_1,N_2,N_3)",
		Verdict:                        "Gate 253 removes the generic Fock-to-so(8) dictionary obstruction but does not erase the separate ontology problem: T3L/Y_phi must be retrieved as number-operator ledgers or as another typed so(8) representative before Q_8vC can be physical",
	}
}

func auditTriality(ew ElectroweakGeneratorCoordinateAudit) TrialityAudit {
	candidates := []TrialityCartanCandidate{
		makeTrialityCandidate("D4 Hadamard Cartan candidate τ_even (vector ↔ even spinor)", [][]float64{{0.5, 0.5, 0.5, 0.5}, {0.5, 0.5, -0.5, -0.5}, {0.5, -0.5, 0.5, -0.5}, {0.5, -0.5, -0.5, 0.5}}),
		makeTrialityCandidate("D4 Hadamard Cartan candidate τ_odd (vector ↔ odd spinor, orientation variant)", [][]float64{{0.5, 0.5, 0.5, -0.5}, {0.5, 0.5, -0.5, 0.5}, {0.5, -0.5, 0.5, 0.5}, {-0.5, 0.5, 0.5, 0.5}}),
	}
	canApply := ew.T3LSO8CoordinatesDerived && ew.YPhiSO8CoordinatesDerived
	return TrialityAudit{
		OuterAutomorphismGroup:              "Out(Spin(8)) ≅ S3; Cartan representatives can be written by D4 Hadamard transforms but the physical branch must be selected by native representation data",
		CandidateCount:                      len(candidates),
		Candidates:                          candidates,
		SpecificSpinorToVectorChoiceDerived: false,
		UsesWrongChoiceRiskAudited:          true,
		CanApplyToPhysicalEW:                canApply,
		Obstruction:                         "two triality branch risks remain: the project has not selected the exact 8_s→8_v outer automorphism, and physical T3L/Y_phi coordinates are still missing",
		Verdict:                             "triality candidates are audited as D4 Cartan maps, but no outcome-driven selection is made and no physical electroweak generator is transported",
	}
}

func makeTrialityCandidate(name string, m [][]float64) TrialityCartanCandidate {
	orth := isOrthogonal(m, 1e-12)
	involutive := isIdentity(mul(m, m), 1e-12)
	det := determinant4(m)
	return TrialityCartanCandidate{
		Name:              name,
		Matrix:            cloneMatrix(m),
		Orthogonal:        orth,
		Involutive:        involutive,
		DetAbsOne:         math.Abs(math.Abs(det)-1) < 1e-12,
		MapsD4RootLattice: orth && math.Abs(math.Abs(det)-1) < 1e-12,
		Selected:          false,
		Verdict:           "lawful abstract Cartan triality candidate; not selected as the physical spinor-to-vector branch in this gate",
	}
}

func auditKernel(ew ElectroweakGeneratorCoordinateAudit, t TrialityAudit) Q8VCKernelAudit {
	constructed := ew.QSO8CoordinatesDerived && t.SpecificSpinorToVectorChoiceDerived
	return Q8VCKernelAudit{
		Definition:               "Q_8vC = i R_8v(τ(T3L + Y_phi)); neutral plane = ker(Q_8vC)",
		T3LCoordinatesAvailable:  ew.T3LSO8CoordinatesDerived,
		YPhiCoordinatesAvailable: ew.YPhiSO8CoordinatesDerived,
		TrialityChoiceAvailable:  t.SpecificSpinorToVectorChoiceDerived,
		Q8vCConstructed:          constructed,
		EigensystemComputed:      constructed,
		KernelDimensionKnown:     constructed,
		KernelComplexDimension:   0,
		ExactlyThree:             false,
		ThreePlaneDerived:        false,
		DiagnosticOnlyReason:     "without physical T3L/Y_phi coordinates and a selected spinor-to-vector triality branch, any computed kernel would be a candidate-kernel, not a theorem",
		Verdict:                  "the neutral 3-plane is still blocked; this gate refuses to manufacture Q_8vC from available nonphysical diagnostic ledgers",
	}
}

func auditDownstream(k Q8VCKernelAudit) DownstreamAudit {
	return DownstreamAudit{
		Neutral3PlaneAvailable: k.ThreePlaneDerived,
		TauEta:                 []int{2, -2, 1},
		VTauConstructed:        false,
		YukawaTextureDerived:   false,
		CKMPMNSDerived:         false,
		FermionMassesDerived:   false,
		Verdict:                "tau_eta keeps generation-breaking capacity, but v_tau/Yukawa/CKM/PMNS remain sealed until the neutral vector three-plane is derived",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		InventedWittPairing:       false,
		InventedT3LCoefficients:   false,
		InventedYPhiCoefficients:  false,
		SelectedTrialityByOutcome: false,
		ForcedKernelDim3:          false,
		ConstructedVTauByHand:     false,
		InsertedYukawaTexture:     false,
		ImportedObservedMasses:    false,
		PollutedFiniteCore:        false,
		Verdict:                   "firewall preserved: the gate derives the dictionary and known number-ledger coordinates only; it does not invent electroweak coefficient vectors, select triality by desired kernel dimension, or insert flavor data",
	}
}

func summarize(w WittBasisAudit, n NumberOperatorExpansionAudit, l KnownFockLedgerAudit, ew ElectroweakGeneratorCoordinateAudit, t TrialityAudit, k Q8VCKernelAudit, d DownstreamAudit) Summary {
	status := strings.Join([]string{
		StatusWittPairingRetrieved,
		StatusNumberOperatorsAsCartanBivectors,
		StatusKnownFockLedgersCoordinateReady,
		StatusTrialityCartanCandidatesAudited,
		StatusT3LYPhiLedgerMissing,
		StatusExplicitTrialityStillUnselected,
		StatusQ8VCStillBlocked,
		StatusNeutral3PlaneStillBlocked,
		StatusVTauStillBlocked,
		StatusYukawaStillBlocked,
	}, ";")
	return Summary{
		WittPairingRetrieved:            w.Retrieved,
		NumberSO8Coordinates:            n.Derived,
		KnownFockLedgersCoordinateReady: l.AllDerivedFromNumberOps,
		T3LYPhiSO8Coordinates:           ew.T3LSO8CoordinatesDerived && ew.YPhiSO8CoordinatesDerived,
		ExplicitTrialitySelected:        t.SpecificSpinorToVectorChoiceDerived,
		Q8vCConstructed:                 k.Q8vCConstructed,
		Neutral3PlaneDerived:            k.ThreePlaneDerived,
		VTauConstructed:                 d.VTauConstructed,
		YukawaTextureDerived:            d.YukawaTextureDerived,
		Status:                          status,
		NextGate:                        "Gate 254 — retrieve or derive the actual T3L/Y_phi coefficient ledger over N_k, or derive their direct Spin(8) bivector representatives, then select the native 8_s→8_v triality branch by representation weights rather than by kernel outcome",
		Comment:                         "Gate 253 succeeds at the Witt/Fock-to-Cartan so(8) dictionary but keeps the physical electroweak generator and triality-branch obstructions separate.",
	}
}

func buildTruth(prev InheritedGate252Audit, w WittBasisAudit, n NumberOperatorExpansionAudit, l KnownFockLedgerAudit, ew ElectroweakGeneratorCoordinateAudit, t TrialityAudit, k Q8VCKernelAudit) string {
	return fmt.Sprintf("Gate 253 reads the native Fock dictionary backwards. It retrieves %d Witt pairs on %s and derives %d number-operator Cartan coordinates %s. This closes the generic 'names are not coordinates' gap for any diagonal Fock number ledger; B-L/T0/conditional weak-plane ledgers are coordinate-ready=%t. However, the physical bridge generators T3L and Y_phi are still not present as coefficient vectors over N_k, and the 8_s→8_v triality branch is not selected. Therefore Q_8vC and the neutral three-plane remain un-derived. inheritedGate252Triality=%t witt=%t numbers=%t T3Y=%t triality=%t Q=%t kernel3=%t", len(w.Pairs), w.TemporalSpatialSplit, n.CoordinateCount, strings.Join(n.CartanBivectors, ","), l.AllDerivedFromNumberOps, prev.InfinitesimalTrialityCapacity, w.Retrieved, n.Derived, ew.T3LSO8CoordinatesDerived && ew.YPhiSO8CoordinatesDerived, t.SpecificSpinorToVectorChoiceDerived, k.Q8vCConstructed, k.ExactlyThree)
}

func formatBivectorFormula(coeffs []float64, labels []string) string {
	parts := []string{}
	for i, c := range coeffs {
		if math.Abs(c) < 1e-12 || i >= len(labels) {
			continue
		}
		parts = append(parts, fmt.Sprintf("%.10g i·%s", c, labels[i]))
	}
	if len(parts) == 0 {
		return "0"
	}
	return strings.Join(parts, " + ")
}

func cloneMatrix(m [][]float64) [][]float64 {
	out := make([][]float64, len(m))
	for i := range m {
		out[i] = append([]float64(nil), m[i]...)
	}
	return out
}

func transpose(m [][]float64) [][]float64 {
	out := make([][]float64, len(m[0]))
	for i := range out {
		out[i] = make([]float64, len(m))
	}
	for r := range m {
		for c := range m[r] {
			out[c][r] = m[r][c]
		}
	}
	return out
}

func mul(a, b [][]float64) [][]float64 {
	out := make([][]float64, len(a))
	for r := range out {
		out[r] = make([]float64, len(b[0]))
		for c := range out[r] {
			for k := range b {
				out[r][c] += a[r][k] * b[k][c]
			}
		}
	}
	return out
}

func isOrthogonal(m [][]float64, eps float64) bool {
	return isIdentity(mul(transpose(m), m), eps)
}

func isIdentity(m [][]float64, eps float64) bool {
	for r := range m {
		for c := range m[r] {
			want := 0.0
			if r == c {
				want = 1
			}
			if math.Abs(m[r][c]-want) > eps {
				return false
			}
		}
	}
	return true
}

func determinant4(m [][]float64) float64 {
	if len(m) != 4 || len(m[0]) != 4 {
		return math.NaN()
	}
	det := 0.0
	for c := 0; c < 4; c++ {
		det += sign(c) * m[0][c] * determinant3(minor4(m, 0, c))
	}
	return det
}

func determinant3(m [][]float64) float64 {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) - m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) + m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

func minor4(m [][]float64, skipR, skipC int) [][]float64 {
	out := [][]float64{}
	for r := 0; r < 4; r++ {
		if r == skipR {
			continue
		}
		row := []float64{}
		for c := 0; c < 4; c++ {
			if c == skipC {
				continue
			}
			row = append(row, m[r][c])
		}
		out = append(out, row)
	}
	return out
}

func sign(i int) float64 {
	if i%2 == 0 {
		return 1
	}
	return -1
}
