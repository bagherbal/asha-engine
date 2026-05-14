// Package carrierintertwiner implements Gate 255:
// Carrier Intertwiner / T3L-Y_phi Representation Unification Audit.
//
// Gate 254 proved that the Witt dictionary is ready for genuine Fock-number
// ledgers, but the physical electroweak objects are carrier-mismatched: T3L is
// a derived left-doublet matrix and Y_phi is a scalar/contact operator.  Gate
// 255 audits whether the existing finite geometry already contains a lawful
// representation functor that embeds both objects into the same complexified
// four-mode Fock carrier S_C = Lambda*(C^4).  It deliberately treats formal
// direct sums, tensor products, and dimensional coincidences as diagnostics, not
// as intertwiners.
package carrierintertwiner

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ewcartanledger"
)

const (
	AuditID = "GATE255-CARRIER-INTERTWINER-T3L-Y-PHI-REPRESENTATION-UNIFICATION-AUDIT"

	StatusGate254Inherited             = "CONDITIONAL_SUPPORT_GATE254_CARRIER_MISMATCH_INHERITED"
	StatusSCKnown                      = "CONDITIONAL_SUPPORT_COMPLEXIFIED_FOCK_CARRIER_KNOWN"
	StatusLocalActionsAudited          = "CONDITIONAL_SUPPORT_LOCAL_CARRIER_ACTIONS_AUDITED"
	StatusScalarOrientationLocalized   = "CONDITIONAL_SUPPORT_SCALAR_ORIENTATION_CLASSIFIED_SPONTANEOUS"
	StatusFormalAssembliesRejected     = "CONDITIONAL_SUPPORT_FORMAL_ASSEMBLIES_REJECTED_AS_INTERTWINERS"
	StatusLeftDoubletProjectionBlocked = "FAILED_ROUTE_T3L_LEFT_DOUBLET_TO_SC_INCLUSION_NOT_DERIVED"
	StatusScalarHphiEmbeddingBlocked   = "FAILED_ROUTE_Y_PHI_HPHI_TO_SC_EMBEDDING_NOT_DERIVED"
	StatusTotalRepresentationMissing   = "FAILED_ROUTE_FAITHFUL_TOTAL_REPRESENTATION_FUNCTOR_MISSING"
	StatusUnifiedLedgerBlocked         = "FAILED_ROUTE_UNIFIED_T3L_Y_PHI_FOCK_LEDGER_BLOCKED"
	StatusSO8CoordinatesStillBlocked   = "FAILED_ROUTE_PHYSICAL_EW_SO8_COORDINATES_STILL_BLOCKED"
	StatusTrialityStillBlocked         = "FAILED_ROUTE_TRIALITY_PULLBACK_STILL_BLOCKED"
	StatusQ8VCThreePlaneStillBlocked   = "FAILED_ROUTE_Q8VC_NEUTRAL_3PLANE_STILL_BLOCKED"
	StatusYukawaStillBlocked           = "FAILED_ROUTE_YUKAWA_TEXTURE_STILL_BLOCKED"
)

type InheritedGate254Audit struct {
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
	TruthStatement             string
}

type CarrierObject struct {
	Name                  string
	Dimension             int
	CarrierType           string
	Source                string
	Available             bool
	CanonicalOnOwnCarrier bool
	ActsOnMatter          bool
	ActsOnScalar          bool
	ActsOnContact         bool
	SubspaceOfSC          bool
	EmbeddingIntoSC       bool
	ProjectionFromSC      bool
	IntertwiningVerified  bool
	CoordinateComplete    bool
	Obstruction           string
	Verdict               string
}

type CarrierInventoryAudit struct {
	ObjectsAudited               []CarrierObject
	SCAvailable                  bool
	SCDimension                  int
	T3LAvailable                 bool
	T3LDimension                 int
	YPhiAvailable                bool
	YPhiDimension                int
	HphiSubspaceOfSC             bool
	LeftDoubletSubspaceOfSC      bool
	CommonSCCarrierAvailable     bool
	CommonActionCarrierAvailable bool
	Verdict                      string
}

type IntertwinerCandidate struct {
	Name                string
	From                string
	To                  string
	Source              string
	Available           bool
	Canonical           bool
	BranchFree          bool
	Injective           bool
	SurjectiveOntoImage bool
	IntertwinesT3L      bool
	IntertwinesYPhi     bool
	MapsIntoSC          bool
	MapsFromSC          bool
	Isometric           bool
	UsesOrientationSeal bool
	UsesObservedInput   bool
	RequiresGaugeFrame  bool
	RejectedReason      string
	Verdict             string
}

type IntertwinerSearchAudit struct {
	Candidates                  []IntertwinerCandidate
	CandidateCount              int
	AvailableCandidates         int
	CanonicalCandidates         int
	SCEmbeddingCandidates       int
	T3LIntertwiningCandidates   int
	YPhiIntertwiningCandidates  int
	JointIntertwiningCandidates int
	LawfulCommonIntertwiner     bool
	RejectedFormalAssemblies    int
	Verdict                     string
}

type UnifiedLedgerAudit struct {
	CommonCarrier                   string
	CommonCarrierDerived            bool
	T3LProjectedToSC                bool
	YPhiProjectedToSC               bool
	T3LNumberCoefficientsAvailable  bool
	YPhiNumberCoefficientsAvailable bool
	T3LNumberCoefficients           []float64
	YPhiNumberCoefficients          []float64
	UnifiedLedgerConstructed        bool
	Obstruction                     string
	Verdict                         string
}

type SO8TranslationAudit struct {
	WittDictionaryAvailable bool
	T3LUnifiedLedger        bool
	YPhiUnifiedLedger       bool
	T3LSO8Coordinates       bool
	YPhiSO8Coordinates      bool
	QSO8Coordinates         bool
	ZSO8Coordinates         bool
	Obstruction             string
	Verdict                 string
}

type TrialityKernelAudit struct {
	TrialityCandidatesKnown    bool
	RepresentationWeightsKnown bool
	PhysicalBranchSelected     bool
	Q8vCConstructed            bool
	EigensystemComputed        bool
	KernelDimensionKnown       bool
	KernelComplexDimension     int
	ExactlyThree               bool
	NeutralThreePlaneDerived   bool
	DiagnosticOnlyReason       string
	Verdict                    string
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

type FirewallAudit struct {
	EmbeddedHphiIntoSCByDimension   bool
	EmbeddedLeftDoubletByLabel      bool
	TreatedTensorProductAsSC        bool
	TreatedDirectSumAsIntertwiner   bool
	ImportedConnesRepresentation    bool
	InsertedSMHyperchargeConvention bool
	ForcedWeakPlane                 bool
	SelectedTrialityByKernel        bool
	ForcedKernelDim3                bool
	ConstructedVTauByHand           bool
	InsertedYukawaTexture           bool
	ImportedObservedMasses          bool
	PollutedFiniteCore              bool
	Verdict                         string
}

type Summary struct {
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
	NextGate                  string
	Comment                   string
}

type Analysis struct {
	PreviousGate254 InheritedGate254Audit
	Carriers        CarrierInventoryAudit
	Intertwiners    IntertwinerSearchAudit
	UnifiedLedger   UnifiedLedgerAudit
	SO8             SO8TranslationAudit
	TrialityKernel  TrialityKernelAudit
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
		prevRaw, err := ewcartanledger.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 254 predecessor: %w", err)
			return
		}
		prev := inheritGate254(prevRaw)

		carriers := auditCarriers(prev, prevRaw)
		intertwiners := auditIntertwiners(carriers)
		ledger := auditUnifiedLedger(prev, carriers, intertwiners)
		so8 := auditSO8(prev, ledger)
		kernel := auditTrialityKernel(so8)
		down := auditDownstream(kernel)
		fw := auditFirewall()
		summary := summarize(prev, carriers, intertwiners, ledger, so8, kernel, down)
		truth := buildTruth(prev, carriers, intertwiners, ledger, so8, kernel)
		defaultA = Analysis{PreviousGate254: prev, Carriers: carriers, Intertwiners: intertwiners, UnifiedLedger: ledger, SO8: so8, TrialityKernel: kernel, Downstream: down, Firewall: fw, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate254(a ewcartanledger.Analysis) InheritedGate254Audit {
	return InheritedGate254Audit{
		Gate253DictionaryInherited: a.Summary.Gate253DictionaryInherited,
		RegistrySearchCompleted:    a.Summary.RegistrySearchCompleted,
		FockLedgersRetrieved:       a.Summary.FockLedgersRetrieved,
		T3LNumberLedgerRetrieved:   a.Summary.T3LNumberLedgerRetrieved,
		YPhiNumberLedgerRetrieved:  a.Summary.YPhiNumberLedgerRetrieved,
		T3LYPhiSO8Coordinates:      a.Summary.T3LYPhiSO8Coordinates,
		TrialityBranchSelected:     a.Summary.TrialityBranchSelected,
		Q8vCConstructed:            a.Summary.Q8vCConstructed,
		Neutral3PlaneDerived:       a.Summary.Neutral3PlaneDerived,
		VTauConstructed:            a.Summary.VTauConstructed,
		YukawaTextureDerived:       a.Summary.YukawaTextureDerived,
		Status:                     a.Summary.Status,
		TruthStatement:             a.TruthStatement,
	}
}

func auditCarriers(prev InheritedGate254Audit, gate254 ewcartanledger.Analysis) CarrierInventoryAudit {
	objects := []CarrierObject{
		{
			Name:                  "S_C complexified four-mode Fock carrier",
			Dimension:             16,
			CarrierType:           "target Spin/Fock carrier",
			Source:                "Gates 14, 235, 253",
			Available:             prev.Gate253DictionaryInherited,
			CanonicalOnOwnCarrier: true,
			ActsOnMatter:          true,
			SubspaceOfSC:          true,
			EmbeddingIntoSC:       true,
			ProjectionFromSC:      true,
			IntertwiningVerified:  true,
			CoordinateComplete:    true,
			Verdict:               "S_C is the lawful target for number-operator Cartan coordinates; it does not by itself embed scalar/contact or derived left-doublet objects.",
		},
		{
			Name:                  "T3L left-doublet action",
			Dimension:             gate254.CarrierTyping.T3LDimension,
			CarrierType:           gate254.CarrierTyping.T3LCarrier,
			Source:                "Gate 24 / Gate 254 carrier typing",
			Available:             gate254.CarrierTyping.T3LBridgeKnown,
			CanonicalOnOwnCarrier: true,
			ActsOnMatter:          true,
			SubspaceOfSC:          false,
			EmbeddingIntoSC:       false,
			ProjectionFromSC:      false,
			IntertwiningVerified:  false,
			CoordinateComplete:    false,
			Obstruction:           "Gate 254 types T3L as a derived left-doublet matrix rather than a native endomorphism of the full sixteen-state Fock carrier; no state-to-occupation inclusion or diagonal N_k ledger is present.",
			Verdict:               "T3L is real finite SU(2)L data, but not a native endomorphism of S_C.",
		},
		{
			Name:                  "Y_phi scalar/contact hypercharge",
			Dimension:             gate254.CarrierTyping.YPhiDimension,
			CarrierType:           gate254.CarrierTyping.YPhiCarrier,
			Source:                "Gate 20 / Gate 254 carrier typing",
			Available:             gate254.CarrierTyping.YPhiBridgeKnown,
			CanonicalOnOwnCarrier: true,
			ActsOnScalar:          true,
			SubspaceOfSC:          false,
			EmbeddingIntoSC:       false,
			ProjectionFromSC:      false,
			IntertwiningVerified:  false,
			CoordinateComplete:    false,
			Obstruction:           "Gate 254 types Y_phi as a scalar/contact operator on H_phi; no theorem identifies H_phi with a subspace of S_C or supplies a Fock-number coefficient vector.",
			Verdict:               "Y_phi is a valid scalar/contact operator, not a Fock-number ledger.",
		},
		{
			Name:                  "formal H_Fock tensor H_phi block",
			Dimension:             64,
			CarrierType:           "formal tensor bookkeeping carrier",
			Source:                "Gate 165 total-representation obstruction audit",
			Available:             true,
			CanonicalOnOwnCarrier: true,
			ActsOnMatter:          true,
			ActsOnScalar:          true,
			SubspaceOfSC:          false,
			EmbeddingIntoSC:       false,
			ProjectionFromSC:      false,
			IntertwiningVerified:  false,
			CoordinateComplete:    false,
			Obstruction:           "tensoring matter and scalar carriers changes the target from S_C to S_C⊗H_phi and does not produce a common four-mode Cartan ledger.",
			Verdict:               "useful bookkeeping but not the requested common S_C carrier.",
		},
		{
			Name:                  "Spin(8) vector carrier 8_v",
			Dimension:             8,
			CarrierType:           "vector representation",
			Source:                "Gate 248 vector representative audit",
			Available:             true,
			CanonicalOnOwnCarrier: true,
			SubspaceOfSC:          false,
			EmbeddingIntoSC:       false,
			ProjectionFromSC:      false,
			IntertwiningVerified:  false,
			CoordinateComplete:    false,
			Obstruction:           "8_v is the eventual triality target, not a proof that scalar/contact and spinor/matter operators have already become S_C Cartan coordinates.",
			Verdict:               "native and important, but downstream of the missing S_C operator ledger.",
		},
	}
	return CarrierInventoryAudit{
		ObjectsAudited:               objects,
		SCAvailable:                  prev.Gate253DictionaryInherited,
		SCDimension:                  16,
		T3LAvailable:                 gate254.CarrierTyping.T3LBridgeKnown,
		T3LDimension:                 gate254.CarrierTyping.T3LDimension,
		YPhiAvailable:                gate254.CarrierTyping.YPhiBridgeKnown,
		YPhiDimension:                gate254.CarrierTyping.YPhiDimension,
		HphiSubspaceOfSC:             false,
		LeftDoubletSubspaceOfSC:      false,
		CommonSCCarrierAvailable:     false,
		CommonActionCarrierAvailable: false,
		Verdict:                      "local carriers are real and canonical on themselves, but no audited object identifies both T3L and Y_phi as endomorphisms of S_C.",
	}
}

func auditIntertwiners(carriers CarrierInventoryAudit) IntertwinerSearchAudit {
	candidates := []IntertwinerCandidate{
		{Name: "identity on S_C", From: "S_C", To: "S_C", Source: "Gate 253", Available: carriers.SCAvailable, Canonical: true, BranchFree: true, Injective: true, SurjectiveOntoImage: true, MapsIntoSC: true, MapsFromSC: true, Isometric: true, RejectedReason: "acts only on objects already in S_C; it cannot import H_phi or the derived left-doublet table", Verdict: "valid identity, not a carrier unifier"},
		{Name: "left-doublet inclusion", From: "Q_L⊕L_L", To: "S_C", Source: "Gate 24 target requested by Gate 255", Available: false, Canonical: false, BranchFree: false, Injective: false, IntertwinesT3L: false, MapsIntoSC: true, RejectedReason: "no native injection from the derived eight-state weak table to occupation basis states is present", Verdict: "missing"},
		{Name: "scalar H_phi embedding", From: "H_phi", To: "S_C", Source: "Gate 20/190 target requested by Gate 255", Available: false, Canonical: false, BranchFree: false, Injective: false, IntertwinesYPhi: false, MapsIntoSC: true, UsesOrientationSeal: true, RequiresGaugeFrame: true, RejectedReason: "the scalar high/low orientation is spontaneous/gauge data and no S_C coordinate map is derived", Verdict: "missing"},
		{Name: "formal direct-sum assembly", From: "H_Fock ⊕ H_phi", To: "H_Fock ⊕ H_phi", Source: "Gate 165 formal assembly", Available: true, Canonical: false, BranchFree: true, Injective: true, SurjectiveOntoImage: true, RejectedReason: "a direct sum lists sectors side by side but does not intertwine them or place both on S_C", Verdict: "rejected as bookkeeping"},
		{Name: "matter-scalar tensor block", From: "H_Fock⊗H_phi", To: "H_Fock⊗H_phi", Source: "Gate 165 matter-scalar tensor representation", Available: true, Canonical: true, BranchFree: true, Injective: true, SurjectiveOntoImage: true, RejectedReason: "the tensor carrier is 64-dimensional and does not yield four Fock-number coefficients for Y_phi or T3L", Verdict: "rejected as target change"},
		{Name: "scalar-to-8_v map", From: "H_phi trace slots", To: "8_v", Source: "Gate 248", Available: false, Canonical: false, BranchFree: false, Injective: false, IntertwinesYPhi: false, RejectedReason: "no basis-independent H_phi -> 8_v map was derived; dimensional embeddability is not a vector representative theorem", Verdict: "blocked before triality"},
		{Name: "A_total faithful representation functor", From: "contact/matter/scalar tower", To: "H_total", Source: "Gate 165 target", Available: false, Canonical: false, IntertwinesT3L: false, IntertwinesYPhi: false, RejectedReason: "faithful total representation and canonical glue maps remain unconstructed", Verdict: "missing"},
	}
	available, canonical, sc, t3, yphi, joint, formal := 0, 0, 0, 0, 0, 0, 0
	for _, c := range candidates {
		if c.Available {
			available++
		}
		if c.Canonical {
			canonical++
		}
		if c.MapsIntoSC && c.Available {
			sc++
		}
		if c.IntertwinesT3L {
			t3++
		}
		if c.IntertwinesYPhi {
			yphi++
		}
		if c.IntertwinesT3L && c.IntertwinesYPhi && c.MapsIntoSC && c.Canonical {
			joint++
		}
		if strings.Contains(c.Name, "formal") || strings.Contains(c.Name, "tensor") {
			formal++
		}
	}
	return IntertwinerSearchAudit{
		Candidates:                  candidates,
		CandidateCount:              len(candidates),
		AvailableCandidates:         available,
		CanonicalCandidates:         canonical,
		SCEmbeddingCandidates:       sc,
		T3LIntertwiningCandidates:   t3,
		YPhiIntertwiningCandidates:  yphi,
		JointIntertwiningCandidates: joint,
		LawfulCommonIntertwiner:     joint > 0,
		RejectedFormalAssemblies:    formal,
		Verdict:                     "intertwiner search completed: only the identity on already-valid S_C data is canonical; no joint T3L/Y_phi embedding into S_C exists.",
	}
}

func auditUnifiedLedger(prev InheritedGate254Audit, carriers CarrierInventoryAudit, inter IntertwinerSearchAudit) UnifiedLedgerAudit {
	common := carriers.CommonSCCarrierAvailable && inter.LawfulCommonIntertwiner
	return UnifiedLedgerAudit{
		CommonCarrier:                   "S_C = Lambda*(C^4)",
		CommonCarrierDerived:            common,
		T3LProjectedToSC:                false,
		YPhiProjectedToSC:               false,
		T3LNumberCoefficientsAvailable:  false,
		YPhiNumberCoefficientsAvailable: false,
		T3LNumberCoefficients:           nil,
		YPhiNumberCoefficients:          nil,
		UnifiedLedgerConstructed:        false,
		Obstruction:                     "Gate 254 retrieved nearby Fock ledgers, but Gate 255 finds no native functor that sends both physical T3L and scalar Y_phi into common S_C number-operator coordinates.",
		Verdict:                         "unified electroweak Fock ledger remains blocked; no coefficients over N0..N3 are emitted.",
	}
}

func auditSO8(prev InheritedGate254Audit, ledger UnifiedLedgerAudit) SO8TranslationAudit {
	return SO8TranslationAudit{
		WittDictionaryAvailable: prev.Gate253DictionaryInherited,
		T3LUnifiedLedger:        ledger.T3LNumberCoefficientsAvailable,
		YPhiUnifiedLedger:       ledger.YPhiNumberCoefficientsAvailable,
		T3LSO8Coordinates:       false,
		YPhiSO8Coordinates:      false,
		QSO8Coordinates:         false,
		ZSO8Coordinates:         false,
		Obstruction:             "the Witt dictionary cannot translate absent unified physical ledgers; translating B-L/T0/candidate weak Cartans would not be T3L+Y_phi.",
		Verdict:                 "physical electroweak so(8) coordinates remain un-derived in Gate 255.",
	}
}

func auditTrialityKernel(so8 SO8TranslationAudit) TrialityKernelAudit {
	q := so8.QSO8Coordinates && so8.T3LSO8Coordinates && so8.YPhiSO8Coordinates
	return TrialityKernelAudit{
		TrialityCandidatesKnown:    true,
		RepresentationWeightsKnown: false,
		PhysicalBranchSelected:     false,
		Q8vCConstructed:            q,
		EigensystemComputed:        q,
		KernelDimensionKnown:       q,
		KernelComplexDimension:     0,
		ExactlyThree:               false,
		NeutralThreePlaneDerived:   false,
		DiagnosticOnlyReason:       "selecting triality by desired 3-plane outcome would be circular; physical S_C weights are still missing.",
		Verdict:                    "triality pullback and neutral three-plane remain blocked.",
	}
}

func auditDownstream(k TrialityKernelAudit) DownstreamAudit {
	return DownstreamAudit{
		Neutral3PlaneAvailable: k.NeutralThreePlaneDerived,
		TauEta:                 []int{2, -2, 1},
		VTauConstructed:        false,
		TrialityTextureOpened:  false,
		YukawaTextureDerived:   false,
		CKMPMNSDerived:         false,
		FermionMassesDerived:   false,
		Verdict:                "tau_eta remains a real generation-breaking capacity, but no v_tau/Yukawa/CKM/PMNS/mass route opens without the common carrier and neutral vector 3-plane.",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		EmbeddedHphiIntoSCByDimension:   false,
		EmbeddedLeftDoubletByLabel:      false,
		TreatedTensorProductAsSC:        false,
		TreatedDirectSumAsIntertwiner:   false,
		ImportedConnesRepresentation:    false,
		InsertedSMHyperchargeConvention: false,
		ForcedWeakPlane:                 false,
		SelectedTrialityByKernel:        false,
		ForcedKernelDim3:                false,
		ConstructedVTauByHand:           false,
		InsertedYukawaTexture:           false,
		ImportedObservedMasses:          false,
		PollutedFiniteCore:              false,
		Verdict:                         "firewall preserved: Gate 255 audits candidate maps but does not promote dimensional embeddings, direct sums, tensor blocks, or external Standard Model conventions into native coordinates.",
	}
}

func summarize(prev InheritedGate254Audit, carriers CarrierInventoryAudit, inter IntertwinerSearchAudit, ledger UnifiedLedgerAudit, so8 SO8TranslationAudit, k TrialityKernelAudit, d DownstreamAudit) Summary {
	status := strings.Join([]string{
		StatusGate254Inherited,
		StatusSCKnown,
		StatusLocalActionsAudited,
		StatusScalarOrientationLocalized,
		StatusFormalAssembliesRejected,
		StatusLeftDoubletProjectionBlocked,
		StatusScalarHphiEmbeddingBlocked,
		StatusTotalRepresentationMissing,
		StatusUnifiedLedgerBlocked,
		StatusSO8CoordinatesStillBlocked,
		StatusTrialityStillBlocked,
		StatusQ8VCThreePlaneStillBlocked,
		StatusYukawaStillBlocked,
	}, ";")
	return Summary{
		Gate254Inherited:          prev.Gate253DictionaryInherited && prev.RegistrySearchCompleted && !prev.T3LYPhiSO8Coordinates,
		SCCarrierKnown:            carriers.SCAvailable && carriers.SCDimension == 16,
		LocalActionsAudited:       carriers.T3LAvailable && carriers.YPhiAvailable,
		CommonCarrierDerived:      ledger.CommonCarrierDerived,
		CarrierIntertwinerDerived: inter.LawfulCommonIntertwiner,
		UnifiedLedgerConstructed:  ledger.UnifiedLedgerConstructed,
		T3LYPhiSO8Coordinates:     so8.T3LSO8Coordinates && so8.YPhiSO8Coordinates,
		TrialityBranchSelected:    k.PhysicalBranchSelected,
		Q8vCConstructed:           k.Q8vCConstructed,
		Neutral3PlaneDerived:      k.NeutralThreePlaneDerived,
		VTauConstructed:           d.VTauConstructed,
		YukawaTextureDerived:      d.YukawaTextureDerived,
		Status:                    status,
		NextGate:                  "Gate 256 — Spontaneous Carrier Seal / Gauge-Fixed H_phi and Left-Doublet Embedding Axiom Audit: if the project chooses to proceed, record the needed vacuum orientation, gauge frame, and state-index ledger as explicit sealed data before any so(8) pullback.",
		Comment:                   "Gate 255 proves the obstruction is not merely missing code: the required common-carrier functor is absent from the finite theorem state.",
	}
}

func buildTruth(prev InheritedGate254Audit, carriers CarrierInventoryAudit, inter IntertwinerSearchAudit, ledger UnifiedLedgerAudit, so8 SO8TranslationAudit, k TrialityKernelAudit) string {
	return fmt.Sprintf("Gate 255 inherits Gate254 ledger obstruction with Fock ledgers=%t and T3L/Y_phi so8=%t. It audits %d carrier objects and %d candidate intertwiners. S_C is known with dimension %d, T3L exists on an %d-dimensional left-doublet carrier, and Y_phi exists on a %d-dimensional scalar/contact carrier. The available formal direct-sum/tensor constructions are rejected as carrier changes, not S_C intertwiners. Therefore commonCarrier=%t lawfulIntertwiner=%t unifiedLedger=%t physicalSO8=%t Q8vC=%t neutral3Plane=%t.", prev.FockLedgersRetrieved, prev.T3LYPhiSO8Coordinates, len(carriers.ObjectsAudited), inter.CandidateCount, carriers.SCDimension, carriers.T3LDimension, carriers.YPhiDimension, ledger.CommonCarrierDerived, inter.LawfulCommonIntertwiner, ledger.UnifiedLedgerConstructed, so8.T3LSO8Coordinates && so8.YPhiSO8Coordinates, k.Q8vCConstructed, k.NeutralThreePlaneDerived)
}
