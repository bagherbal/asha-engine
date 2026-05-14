// Package ewcartanledger implements Gate 254:
// Electroweak Cartan Ledger Retrieval / Native T3L-Y_phi Coefficient Audit.
//
// Gate 253 derived the generic Witt/Fock dictionary from number operators N_k
// to Cartan bivectors in so(8).  Gate 254 searches the active theorem registry
// for the specific electroweak ledgers that would allow T3L and Y_phi to be fed
// through that dictionary.  It deliberately distinguishes nearby but different
// carriers: B-L and T0 are Fock-number ledgers; T3L is currently a derived
// left-doublet matrix; Y_phi/T_phi is scalar/contact data.  Those objects are
// mathematically useful, but they are not interchangeable coordinates.
package ewcartanledger

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/spinctwistedchirality"
	"github.com/bagherbal/asha-engine/pkg/bridge/wittso8coordinates"
	"github.com/bagherbal/asha-engine/pkg/matter/hypercharge"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/t3r"
)

const (
	AuditID = "GATE254-ELECTROWEAK-CARTAN-LEDGER-RETRIEVAL-AUDIT"

	StatusGate253Inherited             = "CONDITIONAL_SUPPORT_GATE253_WITT_DICTIONARY_INHERITED"
	StatusRegistrySearchCompleted      = "CONDITIONAL_SUPPORT_EW_LEDGER_REGISTRY_SEARCH_COMPLETED"
	StatusFockLedgersRetrieved         = "CONDITIONAL_SUPPORT_FOCK_NUMBER_LEDGERS_RETRIEVED"
	StatusMatterT0CoordinateReady      = "CONDITIONAL_SUPPORT_MATTER_T0_T3R_DIAGNOSTIC_COORDINATE_READY"
	StatusScalarYPhiTypedNonFock       = "CONDITIONAL_SUPPORT_Y_PHI_TYPED_AS_SCALAR_CONTACT_NOT_FOCK_LEDGER"
	StatusT3LTypedLeftDoubletNonFock   = "CONDITIONAL_SUPPORT_T3L_TYPED_AS_LEFT_DOUBLET_MATRIX_NOT_NATIVE_FOCK_LEDGER"
	StatusCandidateWeakCartansAudited  = "CONDITIONAL_SUPPORT_CANDIDATE_WEAK_PLANE_CARTANS_AUDITED"
	StatusT3LNumberLedgerMissing       = "FAILED_ROUTE_T3L_NATIVE_NUMBER_OPERATOR_LEDGER_MISSING"
	StatusYPhiNumberLedgerMissing      = "FAILED_ROUTE_Y_PHI_NATIVE_NUMBER_OPERATOR_LEDGER_MISSING"
	StatusPhysicalEWCoordinatesMissing = "FAILED_ROUTE_PHYSICAL_EW_SO8_COORDINATES_STILL_MISSING"
	StatusTrialityBranchStillBlocked   = "FAILED_ROUTE_TRIALITY_BRANCH_SELECTION_STILL_BLOCKED"
	StatusQ8VCStillBlocked             = "FAILED_ROUTE_Q8VC_NEUTRAL_3PLANE_STILL_BLOCKED"
	StatusYukawaStillBlocked           = "FAILED_ROUTE_YUKAWA_TEXTURE_STILL_BLOCKED"
)

type InheritedGate253Audit struct {
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
	TruthStatement                  string
}

type RetrievedLedger struct {
	Name                     string
	Source                   string
	Carrier                  string
	Expression               string
	NumberCoefficients       []float64
	CentralIdentityShift     float64
	BivectorCoefficients     []float64
	BivectorFormula          string
	NativeNumberOperatorForm bool
	CoordinateReady          bool
	PhysicalT3L              bool
	PhysicalYPhi             bool
	PhysicalHypercharge      bool
	Obstruction              string
	Verdict                  string
}

type RegistrySearchAudit struct {
	SourcesSearched               []string
	BMinusLRetrieved              bool
	NativeU1Retrieved             bool
	TemporalT0Retrieved           bool
	ScalarTPhiRetrieved           bool
	MatterT3RDiagnosticRetrieved  bool
	LeftDoubletT3LRetrieved       bool
	CandidateWeakCartansRetrieved bool
	T3LAsNativeNumberLedger       bool
	YPhiAsNativeNumberLedger      bool
	CompleteEWLedgerFound         bool
	CarrierMismatchDetected       bool
	Verdict                       string
}

type CandidateWeakCartan struct {
	Name                 string
	Plane                string
	ModeIndices          []int
	NumberCoefficients   []float64
	BivectorCoefficients []float64
	BivectorFormula      string
	SpatialPreservingU1  bool
	SelectedPhysicalT3L  bool
	Obstruction          string
	Verdict              string
}

type TranslationAudit struct {
	TranslatedLedgers          []RetrievedLedger
	TranslatedLedgerCount      int
	BMinusLSO8Coordinate       bool
	TemporalT0SO8Coordinate    bool
	CandidateWeakSO8Coordinate bool
	T3LSO8Coordinate           bool
	YPhiSO8Coordinate          bool
	QSO8Coordinate             bool
	ZSO8Coordinate             bool
	Obstruction                string
	Verdict                    string
}

type CarrierTypingAudit struct {
	T3LBridgeKnown        bool
	T3LCarrier            string
	T3LDimension          int
	T3LNumberLedgerFound  bool
	T3LDirectSO8Found     bool
	YPhiBridgeKnown       bool
	YPhiCarrier           string
	YPhiDimension         int
	YPhiNumberLedgerFound bool
	YPhiDirectSO8Found    bool
	MatterT3RCarrier      string
	MatterT3RNumberLedger bool
	ConflationRejected    bool
	Verdict               string
}

type TrialityBranchAudit struct {
	CandidateBranchCount           int
	RepresentationWeightsAvailable bool
	T3LWeightsAvailable            bool
	YPhiWeightsAvailable           bool
	CanSelect8sTo8v                bool
	SelectedBranch                 string
	SelectedByOutcome              bool
	Obstruction                    string
	Verdict                        string
}

type Q8VCKernelAudit struct {
	Definition               string
	T3LCoordinatesAvailable  bool
	YPhiCoordinatesAvailable bool
	TrialityBranchAvailable  bool
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
	ImportedSMHyperchargeAsLedger bool
	ConflatedT3RWithT3L           bool
	ConflatedScalarYPhiWithFockY  bool
	ForcedWeakPlane               bool
	SelectedTrialityByKernel      bool
	ForcedKernelDim3              bool
	ConstructedVTauByHand         bool
	InsertedYukawaTexture         bool
	ImportedObservedMasses        bool
	PollutedFiniteCore            bool
	Verdict                       string
}

type Summary struct {
	Gate253DictionaryInherited bool
	RegistrySearchCompleted    bool
	FockLedgersRetrieved       bool
	T3LNumberLedgerRetrieved   bool
	YPhiNumberLedgerRetrieved  bool
	T3LYPhiSO8Coordinates      bool
	TrialityBranchSelected     bool
	Q8vCConstructed            bool
	Neutral3PlaneDerived       bool
	VTauConstructed            bool
	YukawaTextureDerived       bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	PreviousGate253 InheritedGate253Audit
	RegistrySearch  RegistrySearchAudit
	Ledgers         []RetrievedLedger
	WeakCartans     []CandidateWeakCartan
	Translation     TranslationAudit
	CarrierTyping   CarrierTypingAudit
	Triality        TrialityBranchAudit
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
		prevRaw, err := wittso8coordinates.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 253 predecessor: %w", err)
			return
		}
		prev := inheritGate253(prevRaw)

		t3rA, err := t3r.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build matter T3R audit: %w", err)
			return
		}
		hypA, err := hypercharge.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar hypercharge audit: %w", err)
			return
		}
		su2A, err := su2lgauge.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build finite SU2L gauge audit: %w", err)
			return
		}
		spincA, err := spinctwistedchirality.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build native U1 twist audit: %w", err)
			return
		}

		ledgers := auditLedgers(prevRaw, t3rA, hypA, su2A, spincA)
		weak := auditCandidateWeakCartans(prevRaw, spincA)
		search := auditRegistrySearch(ledgers, weak)
		translation := auditTranslation(ledgers, weak)
		carrier := auditCarrierTyping(t3rA, hypA, su2A, translation)
		triality := auditTriality(translation, carrier)
		kernel := auditKernel(translation, triality)
		down := auditDownstream(kernel)
		fw := auditFirewall()
		summary := summarize(prev, search, translation, carrier, triality, kernel, down)
		truth := buildTruth(prev, search, translation, carrier, triality, kernel)
		defaultA = Analysis{PreviousGate253: prev, RegistrySearch: search, Ledgers: ledgers, WeakCartans: weak, Translation: translation, CarrierTyping: carrier, Triality: triality, Kernel: kernel, Downstream: down, Firewall: fw, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate253(a wittso8coordinates.Analysis) InheritedGate253Audit {
	return InheritedGate253Audit{
		WittPairingRetrieved:            a.Summary.WittPairingRetrieved,
		NumberSO8Coordinates:            a.Summary.NumberSO8Coordinates,
		KnownFockLedgersCoordinateReady: a.Summary.KnownFockLedgersCoordinateReady,
		T3LYPhiSO8Coordinates:           a.Summary.T3LYPhiSO8Coordinates,
		ExplicitTrialitySelected:        a.Summary.ExplicitTrialitySelected,
		Q8vCConstructed:                 a.Summary.Q8vCConstructed,
		Neutral3PlaneDerived:            a.Summary.Neutral3PlaneDerived,
		VTauConstructed:                 a.Summary.VTauConstructed,
		YukawaTextureDerived:            a.Summary.YukawaTextureDerived,
		Status:                          a.Summary.Status,
		TruthStatement:                  a.TruthStatement,
	}
}

func auditLedgers(prev wittso8coordinates.Analysis, t t3r.Analysis, h hypercharge.Analysis, s su2lgauge.Analysis, sp spinctwistedchirality.Analysis) []RetrievedLedger {
	labels := prev.NumberOperators.CartanBivectors
	out := []RetrievedLedger{}
	out = append(out, makeNumberLedger("B-L", "Gate 16 / Gate 253 Fock charge polarization", "S_C=Λ*(C^4) Fock carrier", "-N_0 + (1/3)(N_1+N_2+N_3)", []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}, labels, false, false, false, "native 1⊕3 baryon-minus-lepton ledger; useful abelian input but not T3L or Y_phi"))
	out = append(out, makeNumberLedger("Y_native diagonal u(1)", "Gate 240 Spin^c twisted chirality sieve", "S_C=Λ*(C^4) Fock carrier", sp.U1.WeightFormula, sp.U1.ModeWeights, labels, false, false, false, "same native diagonal charge class as B-L; a class sieve, not Standard Model hypercharge"))
	out = append(out, makeNumberLedger("T0 temporal polarization", "Gate 21 / matter T3R diagnostic", "S_C=Λ*(C^4) Fock carrier", "1/2 I - N_0", []float64{-1, 0, 0, 0}, labels, false, false, false, fmt.Sprintf("vectorlike temporal seed exists with Tr(T0^2)=%.10g; chiral restrictions are projectors, not pure N_k ledgers, and T3R is not T3L", t.TemporalTraceSquared)))
	out = append(out, RetrievedLedger{
		Name:                     "chiral-restricted T3_R branch",
		Source:                   "Gate 21 / Gate 22 matter-side hypercharge audit",
		Carrier:                  "chiral sub-blocks of the 16-state Fock carrier",
		Expression:               "P_odd(1/2 I - N_0) or P_even(1/2 I - N_0)",
		NativeNumberOperatorForm: false,
		CoordinateReady:          false,
		PhysicalT3L:              false,
		PhysicalYPhi:             false,
		PhysicalHypercharge:      t.HyperchargeCandidateConstructed,
		Obstruction:              "physical branch uses a parity projector P_even/P_odd in addition to N_0; it is matter-side T3_R, not the left SU(2) Cartan T3L",
		Verdict:                  StatusMatterT0CoordinateReady + "; parity-restricted physical use is not a pure Cartan number ledger",
	})
	out = append(out, RetrievedLedger{
		Name:                     "T_phi / Y_phi scalar-contact charge",
		Source:                   "Gate 20 scalar hypercharge bridge / scalar-contact doublet",
		Carrier:                  "4D active scalar/contact factor H_phi, not S_C Fock number space",
		Expression:               "diag(+1/2,+1/2,-1/2,-1/2) on H_phi; equivalently phase-rotation action in scalar covariant templates",
		NativeNumberOperatorForm: false,
		CoordinateReady:          false,
		PhysicalT3L:              false,
		PhysicalYPhi:             true,
		PhysicalHypercharge:      false,
		Obstruction:              fmt.Sprintf("scalar dimension=%d and trace2=%.10g; no theorem identifies this scalar/contact operator with a coefficient vector over Fock N_0..N_3", h.ScalarDimension, h.ScalarChargeTrace2),
		Verdict:                  StatusScalarYPhiTypedNonFock,
	})
	out = append(out, RetrievedLedger{
		Name:                     "T3L finite left-doublet Cartan",
		Source:                   "Gate 24 finite SU(2)_L gauge doublet audit",
		Carrier:                  "derived 8-state left-doublet table Q_L⊕L_L, not native full S_C Fock basis",
		Expression:               "diag(+1/2,-1/2) on each selected weak doublet; ladder closure [T+,T-]=2T3 verified",
		NativeNumberOperatorForm: false,
		CoordinateReady:          false,
		PhysicalT3L:              s.NonabelianSU2LGeneratorsDerived,
		PhysicalYPhi:             false,
		PhysicalHypercharge:      false,
		Obstruction:              fmt.Sprintf("left-doublet dimension=%d; the matrix is derived after charge-table selection and has no native coefficient vector over all four Fock occupation numbers", s.Dimension),
		Verdict:                  StatusT3LTypedLeftDoubletNonFock,
	})
	return out
}

func makeNumberLedger(name, source, carrier, expr string, coeffs []float64, labels []string, t3l, yphi, hyper bool, obstruction string) RetrievedLedger {
	central := 0.0
	biv := make([]float64, len(coeffs))
	for i, c := range coeffs {
		central += 0.5 * c
		biv[i] = 0.5 * c
	}
	return RetrievedLedger{
		Name:                     name,
		Source:                   source,
		Carrier:                  carrier,
		Expression:               expr,
		NumberCoefficients:       append([]float64(nil), coeffs...),
		CentralIdentityShift:     central,
		BivectorCoefficients:     biv,
		BivectorFormula:          formatBivectorFormula(biv, labels),
		NativeNumberOperatorForm: len(coeffs) == 4,
		CoordinateReady:          len(coeffs) == 4,
		PhysicalT3L:              t3l,
		PhysicalYPhi:             yphi,
		PhysicalHypercharge:      hyper,
		Obstruction:              obstruction,
		Verdict:                  "valid native number-operator ledger and therefore so(8)-Cartan coordinate-ready; physical role is limited to the source theorem",
	}
}

func auditCandidateWeakCartans(prev wittso8coordinates.Analysis, sp spinctwistedchirality.Analysis) []CandidateWeakCartan {
	labels := prev.NumberOperators.CartanBivectors
	preserving := map[string]bool{}
	for _, p := range sp.Planes {
		if p.SurvivesU1CommutantSieve {
			preserving[fmt.Sprintf("%d-%d", p.ModeIndices[0], p.ModeIndices[1])] = true
		}
	}
	pairs := [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	out := make([]CandidateWeakCartan, 0, len(pairs))
	for _, pair := range pairs {
		coeff := []float64{0, 0, 0, 0}
		coeff[pair[0]] = 0.5
		coeff[pair[1]] = -0.5
		biv := make([]float64, len(coeff))
		for i, c := range coeff {
			biv[i] = 0.5 * c
		}
		spatial := preserving[fmt.Sprintf("%d-%d", pair[0], pair[1])]
		verdict := "candidate exterior SU(2) Cartan coordinate only; not selected as physical T3L"
		if spatial {
			verdict = StatusCandidateWeakCartansAudited + "; survives native U1 pure-spatial sieve but remains one of a degenerate family"
		}
		out = append(out, CandidateWeakCartan{
			Name:                 fmt.Sprintf("T3_U%d%d", pair[0], pair[1]),
			Plane:                fmt.Sprintf("U={a†_%d,a†_%d}", pair[0], pair[1]),
			ModeIndices:          []int{pair[0], pair[1]},
			NumberCoefficients:   coeff,
			BivectorCoefficients: biv,
			BivectorFormula:      formatBivectorFormula(biv, labels),
			SpatialPreservingU1:  spatial,
			SelectedPhysicalT3L:  false,
			Obstruction:          "the project derives candidate wedge su(2) lifts for every two-mode plane and a U1 sieve that leaves three pure-spatial planes; no theorem selects one plane as electroweak",
			Verdict:              verdict,
		})
	}
	return out
}

func auditRegistrySearch(ledgers []RetrievedLedger, weak []CandidateWeakCartan) RegistrySearchAudit {
	search := RegistrySearchAudit{
		SourcesSearched: []string{
			"Gate 16/253 B-L Fock number ledger",
			"Gate 21/22 matter T3R and hypercharge branch audits",
			"Gate 20 scalar-contact T_phi/Y_phi bridge",
			"Gate 24 finite SU(2)_L left-doublet generators",
			"Gate 237/240 candidate weak planes and native U1 sieve",
		},
		CandidateWeakCartansRetrieved: len(weak) == 6,
		CarrierMismatchDetected:       true,
	}
	for _, l := range ledgers {
		switch l.Name {
		case "B-L":
			search.BMinusLRetrieved = l.CoordinateReady
		case "Y_native diagonal u(1)":
			search.NativeU1Retrieved = l.CoordinateReady
		case "T0 temporal polarization":
			search.TemporalT0Retrieved = l.CoordinateReady
		case "T_phi / Y_phi scalar-contact charge":
			search.ScalarTPhiRetrieved = l.PhysicalYPhi
		case "chiral-restricted T3_R branch":
			search.MatterT3RDiagnosticRetrieved = true
		case "T3L finite left-doublet Cartan":
			search.LeftDoubletT3LRetrieved = l.PhysicalT3L
		}
		search.T3LAsNativeNumberLedger = search.T3LAsNativeNumberLedger || (l.PhysicalT3L && l.NativeNumberOperatorForm && l.CoordinateReady)
		search.YPhiAsNativeNumberLedger = search.YPhiAsNativeNumberLedger || (l.PhysicalYPhi && l.NativeNumberOperatorForm && l.CoordinateReady)
	}
	search.CompleteEWLedgerFound = search.T3LAsNativeNumberLedger && search.YPhiAsNativeNumberLedger
	search.Verdict = "registry search completed: nearby electroweak objects are present, but the complete native Fock-Cartan ledgers for physical T3L and Y_phi are not present"
	return search
}

func auditTranslation(ledgers []RetrievedLedger, weak []CandidateWeakCartan) TranslationAudit {
	translated := []RetrievedLedger{}
	for _, l := range ledgers {
		if l.CoordinateReady {
			translated = append(translated, l)
		}
	}
	candidateReady := len(weak) > 0
	for _, w := range weak {
		candidateReady = candidateReady && len(w.BivectorCoefficients) == 4
	}
	t3l := false
	yphi := false
	for _, l := range ledgers {
		t3l = t3l || (l.PhysicalT3L && l.CoordinateReady)
		yphi = yphi || (l.PhysicalYPhi && l.CoordinateReady)
	}
	return TranslationAudit{
		TranslatedLedgers:          translated,
		TranslatedLedgerCount:      len(translated),
		BMinusLSO8Coordinate:       findReady(translated, "B-L"),
		TemporalT0SO8Coordinate:    findReady(translated, "T0 temporal polarization"),
		CandidateWeakSO8Coordinate: candidateReady,
		T3LSO8Coordinate:           t3l,
		YPhiSO8Coordinate:          yphi,
		QSO8Coordinate:             t3l && yphi,
		ZSO8Coordinate:             t3l && yphi,
		Obstruction:                "Witt translation succeeds for true Fock number ledgers and for candidate weak-plane Cartans, but not for physical T3L/Y_phi because their current carrier data are not native four-mode number coefficients",
		Verdict:                    "translation is partially successful and type-safe; physical Q=T3L+Y_phi remains blocked",
	}
}

func auditCarrierTyping(t t3r.Analysis, h hypercharge.Analysis, s su2lgauge.Analysis, tr TranslationAudit) CarrierTypingAudit {
	return CarrierTypingAudit{
		T3LBridgeKnown:        s.NonabelianSU2LGeneratorsDerived,
		T3LCarrier:            "derived left-doublet representation from Gate 24, dimension 8",
		T3LDimension:          s.Dimension,
		T3LNumberLedgerFound:  tr.T3LSO8Coordinate,
		T3LDirectSO8Found:     false,
		YPhiBridgeKnown:       h.ScalarChargeBridgeConstructed,
		YPhiCarrier:           "scalar/contact active factor H_phi, dimension 4",
		YPhiDimension:         h.ScalarDimension,
		YPhiNumberLedgerFound: tr.YPhiSO8Coordinate,
		YPhiDirectSO8Found:    false,
		MatterT3RCarrier:      fmt.Sprintf("Fock temporal seed exists; matter-side T3R diagnostic=%t mirrorAmbiguity=%t", t.MatterSideOperatorFound, t.MirrorAmbiguity),
		MatterT3RNumberLedger: true,
		ConflationRejected:    true,
		Verdict:               "T3R/T3L/Y_phi carriers are separated: T0 is Fock-Cartan coordinate-ready, T3L is a left-doublet matrix, and Y_phi is scalar/contact. No carrier equality theorem is present.",
	}
}

func auditTriality(tr TranslationAudit, c CarrierTypingAudit) TrialityBranchAudit {
	can := tr.T3LSO8Coordinate && tr.YPhiSO8Coordinate && c.T3LDirectSO8Found && c.YPhiDirectSO8Found
	return TrialityBranchAudit{
		CandidateBranchCount:           2,
		RepresentationWeightsAvailable: false,
		T3LWeightsAvailable:            tr.T3LSO8Coordinate,
		YPhiWeightsAvailable:           tr.YPhiSO8Coordinate,
		CanSelect8sTo8v:                can,
		SelectedBranch:                 "",
		SelectedByOutcome:              false,
		Obstruction:                    "the branch cannot be selected by desired kernel dimension; it requires physical T3L/Y_phi weights as so(8) coordinates on a specified spinor representation",
		Verdict:                        "triality branch remains unselected; candidate D4 Cartan maps from Gate 253 stay diagnostic only",
	}
}

func auditKernel(tr TranslationAudit, tb TrialityBranchAudit) Q8VCKernelAudit {
	constructed := tr.QSO8Coordinate && tb.CanSelect8sTo8v
	return Q8VCKernelAudit{
		Definition:               "Q_8vC = i R_8v(τ(T3L + Y_phi)); neutral plane = ker(Q_8vC)",
		T3LCoordinatesAvailable:  tr.T3LSO8Coordinate,
		YPhiCoordinatesAvailable: tr.YPhiSO8Coordinate,
		TrialityBranchAvailable:  tb.CanSelect8sTo8v,
		Q8vCConstructed:          constructed,
		EigensystemComputed:      constructed,
		KernelDimensionKnown:     constructed,
		KernelComplexDimension:   0,
		ExactlyThree:             false,
		ThreePlaneDerived:        false,
		DiagnosticOnlyReason:     "the prerequisites for physical Q_8vC are absent; computing a candidate kernel from B-L/T0/weak-plane guesses would be a type error",
		Verdict:                  "neutral 3-plane remains un-derived in Gate 254",
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
		Verdict:                "tau_eta remains a valid generation-breaking capacity record, but no v_tau/Yukawa/CKM/PMNS/mass derivation is opened without the neutral vector three-plane",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ImportedSMHyperchargeAsLedger: false,
		ConflatedT3RWithT3L:           false,
		ConflatedScalarYPhiWithFockY:  false,
		ForcedWeakPlane:               false,
		SelectedTrialityByKernel:      false,
		ForcedKernelDim3:              false,
		ConstructedVTauByHand:         false,
		InsertedYukawaTexture:         false,
		ImportedObservedMasses:        false,
		PollutedFiniteCore:            false,
		Verdict:                       "firewall preserved: the audit retrieves real ledgers and classifies carrier mismatches, but does not turn nearby diagnostics into physical electroweak so(8) coordinates",
	}
}

func summarize(prev InheritedGate253Audit, rs RegistrySearchAudit, tr TranslationAudit, c CarrierTypingAudit, tb TrialityBranchAudit, k Q8VCKernelAudit, d DownstreamAudit) Summary {
	status := strings.Join([]string{
		StatusGate253Inherited,
		StatusRegistrySearchCompleted,
		StatusFockLedgersRetrieved,
		StatusMatterT0CoordinateReady,
		StatusScalarYPhiTypedNonFock,
		StatusT3LTypedLeftDoubletNonFock,
		StatusCandidateWeakCartansAudited,
		StatusT3LNumberLedgerMissing,
		StatusYPhiNumberLedgerMissing,
		StatusPhysicalEWCoordinatesMissing,
		StatusTrialityBranchStillBlocked,
		StatusQ8VCStillBlocked,
		StatusYukawaStillBlocked,
	}, ";")
	return Summary{
		Gate253DictionaryInherited: prev.WittPairingRetrieved && prev.NumberSO8Coordinates,
		RegistrySearchCompleted:    rs.BMinusLRetrieved && rs.TemporalT0Retrieved && rs.ScalarTPhiRetrieved && rs.LeftDoubletT3LRetrieved,
		FockLedgersRetrieved:       tr.BMinusLSO8Coordinate && tr.TemporalT0SO8Coordinate,
		T3LNumberLedgerRetrieved:   c.T3LNumberLedgerFound,
		YPhiNumberLedgerRetrieved:  c.YPhiNumberLedgerFound,
		T3LYPhiSO8Coordinates:      tr.T3LSO8Coordinate && tr.YPhiSO8Coordinate,
		TrialityBranchSelected:     tb.CanSelect8sTo8v,
		Q8vCConstructed:            k.Q8vCConstructed,
		Neutral3PlaneDerived:       k.ThreePlaneDerived,
		VTauConstructed:            d.VTauConstructed,
		YukawaTextureDerived:       d.YukawaTextureDerived,
		Status:                     status,
		NextGate:                   "Gate 255 — Carrier Intertwiner / T3L-Y_phi Representation Unification Audit: derive a native map from scalar/contact H_phi and derived left-doublet SU(2)L data into the same Spin(8) representation carrier, or prove the required operator cannot live as a pure Fock Cartan ledger",
		Comment:                    "Gate 254 completes the ledger search and retrieves nearby coordinates, but the physical electroweak pair is still carrier-mismatched rather than merely undocumented.",
	}
}

func buildTruth(prev InheritedGate253Audit, rs RegistrySearchAudit, tr TranslationAudit, c CarrierTypingAudit, tb TrialityBranchAudit, k Q8VCKernelAudit) string {
	return fmt.Sprintf("Gate 254 inherits the Gate-253 Witt dictionary=%t and searches the active electroweak registry. It retrieves Fock-number coordinates for B-L and T0 and audits %d candidate weak-plane Cartans, but it also proves the key carrier mismatch: physical T3L is currently a %d-dimensional left-doublet matrix, while Y_phi is a %d-dimensional scalar/contact operator. Neither is a native coefficient vector over (N_0,N_1,N_2,N_3), and neither has a direct so(8) representative theorem. Therefore T3L/Y_phi so(8) coordinates=%t, trialityBranch=%t, Q8vC=%t, neutral3Plane=%t. completeEWLedger=%t carrierMismatch=%t", prev.NumberSO8Coordinates, len(tr.TranslatedLedgers), c.T3LDimension, c.YPhiDimension, tr.T3LSO8Coordinate && tr.YPhiSO8Coordinate, tb.CanSelect8sTo8v, k.Q8vCConstructed, k.ThreePlaneDerived, rs.CompleteEWLedgerFound, rs.CarrierMismatchDetected)
}

func findReady(xs []RetrievedLedger, name string) bool {
	for _, x := range xs {
		if x.Name == name && x.CoordinateReady {
			return true
		}
	}
	return false
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
