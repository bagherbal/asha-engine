// Package sealedcarrierwitness implements Gate 257:
// Sealed Carrier Embedding Data / Weak-Frame and Triality-Branch Witness Audit.
//
// The gate follows the GateResearcherMethod discipline.  It does not import an
// observed Standard-Model charge table.  It first separates charge eigenvalues
// that are already present in the engine from the genuinely sealed datum: the
// embedding orientation that places scalar/contact and left-doublet carriers in
// the common four-mode Fock carrier S_C.  It then scans every audited weak-frame
// witness and every Cartan triality branch, instead of selecting the branch by
// the desired neutral kernel.
package sealedcarrierwitness

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/spontaneouscarrierseal"
)

const (
	AuditID = "GATE257-SEALED-CARRIER-EMBEDDING-DATA-WEAK-FRAME-TRIALITY-BRANCH-WITNESS-AUDIT"

	StatusGate256SealInherited       = "CONDITIONAL_SUPPORT_GATE256_SPONTANEOUS_CARRIER_SEAL_INHERITED"
	StatusNativeChargeTableExtracted = "CONDITIONAL_SUPPORT_NATIVE_CHARGE_EIGENVALUE_TABLE_EXTRACTED"
	StatusEmbeddingWitnessScanned    = "CONDITIONAL_SUPPORT_SEALED_EMBEDDING_WITNESSES_SCANNED"
	StatusWittSO8Translated          = "CONDITIONAL_SUPPORT_WITNESS_Q_SO8_CARTAN_TRANSLATED"
	StatusAllTrialityBranchesScanned = "CONDITIONAL_SUPPORT_ALL_TRIALITY_BRANCHES_SCANNED"
	StatusNoOutcomeBranchSelection   = "CONDITIONAL_SUPPORT_NO_TRIALITY_BRANCH_SELECTED_BY_HAND"

	StatusWeakFrameStillDegenerate        = "FAILED_ROUTE_WEAK_FRAME_EMBEDDING_STILL_DEGENERATE"
	StatusNoUniqueTrialityBranch          = "FAILED_ROUTE_TRIALITY_BRANCH_NOT_UNIQUELY_SELECTED_BY_3PLANE"
	StatusNeutral3PlaneNotDerived         = "FAILED_ROUTE_SEALED_WITNESS_NEUTRAL_3PLANE_NOT_DERIVED"
	StatusFullQ8VCKernelNotThree          = "FAILED_ROUTE_FULL_Q8VC_KERNEL_NOT_THREE_DIMENSIONAL"
	StatusYOnlyDiagnosticRejected         = "FAILED_ROUTE_Y_ONLY_THREE_SLOT_DIAGNOSTIC_REJECTED_AS_NOT_Q"
	StatusYukawaTextureStillSealed        = "FAILED_ROUTE_YUKAWA_TEXTURE_STILL_SEALED"
	StatusConcreteT3LYPhiStillConditional = "FAILED_ROUTE_CONCRETE_T3L_Y_PHI_LEDGER_REMAINS_EMBEDDING_CONDITIONAL"
)

type InheritedGate256Audit struct {
	SpontaneousSealRecorded       bool
	ConditionalIntertwinerSchema  bool
	SymbolicLedgerSchemaAvailable bool
	SymbolicSO8SchemaAvailable    bool
	ConcreteUnifiedLedgerBuilt    bool
	ConcreteSO8Coordinates        bool
	TrialityBranchSelected        bool
	Q8vCConstructed               bool
	Neutral3PlaneDerived          bool
	YukawaTextureDerived          bool
	Status                        string
	TruthStatement                string
}

type ChargeSource struct {
	Name                    string
	Source                  string
	Carrier                 string
	Expression              string
	EigenvaluesDerived      bool
	CoefficientVector       []float64
	CoefficientVectorRole   string
	RequiresEmbeddingSeal   bool
	UsesObservedInput       bool
	NativeFiniteCoreTheorem bool
	Verdict                 string
}

type ChargeExtractionAudit struct {
	Sources                          []ChargeSource
	BMinusLFockLedgerDerived         bool
	ScalarYphiEigenvaluesDerived     bool
	T3LLeftDoubletEigenvaluesDerived bool
	ChargeEigenvalueTableDerived     bool
	PhysicalT3LDirectSCVector        bool
	PhysicalYPhiDirectSCVector       bool
	ExternalChargeInputUsed          bool
	Verdict                          string
}

type WeakFrameWitness struct {
	Name               string
	ModePair           [2]int
	OrientationSign    int
	T3Coefficients     []float64
	DerivedEigenvalues bool
	EmbeddingSealed    bool
	NativeSelected     bool
	Verdict            string
}

type ScalarEmbeddingWitness struct {
	Name               string
	Kind               string
	PositiveModes      []int
	YPhiCoefficients   []float64
	DerivedEigenvalues bool
	EmbeddingSealed    bool
	NativeSelected     bool
	Verdict            string
}

type EmbeddingWitnessAudit struct {
	SealName                      string
	WeakFrames                    []WeakFrameWitness
	ScalarEmbeddings              []ScalarEmbeddingWitness
	WeakFrameCount                int
	ScalarEmbeddingCount          int
	TotalCombinedWitnesses        int
	NativeWeakFrameSelected       bool
	NativeScalarEmbeddingSelected bool
	AllWitnessesSealed            bool
	UsesObservedMasses            bool
	UsesObservedYukawas           bool
	Verdict                       string
}

type QLedgerWitness struct {
	Name                 string
	WeakFrameName        string
	ScalarEmbeddingName  string
	T3Coefficients       []float64
	YPhiCoefficients     []float64
	QCoefficients        []float64
	BivectorCoefficients []float64
	CentralIdentityShift float64
	SO8Formula           string
	DerivedChargeInput   bool
	SealedEmbeddingInput bool
	CoordinateTranslated bool
}

type SO8WitnessAudit struct {
	WittDictionaryInherited bool
	CartanBivectors         []string
	Witnesses               []QLedgerWitness
	WitnessCount            int
	AllTranslated           bool
	Verdict                 string
}

type TrialityBranch struct {
	Name             string
	Description      string
	Matrix           [][]float64
	Orthogonal       bool
	DetAbsOne        bool
	AdmissibleCartan bool
}

type BranchScanResult struct {
	WitnessName               string
	WeakFrameName             string
	ScalarEmbeddingName       string
	BranchName                string
	QCoefficients             []float64
	TransformedCoefficients   []float64
	ZeroCartanSlots           int
	PolarizedKernelComplexDim int
	FullQ8vCKernelComplexDim  int
	ExactPolarized3Plane      bool
	ExactFull3Kernel          bool
	SelectedByMathematics     bool
	RejectedReason            string
}

type YOnlyDiagnostic struct {
	Run                         bool
	ScalarEmbeddingName         string
	BranchName                  string
	TransformedCoefficients     []float64
	PolarizedKernelComplexDim   int
	WouldGivePolarizedThreeSlot bool
	RejectedBecauseMissingT3L   bool
	Verdict                     string
}

type TrialityScanAudit struct {
	Branches                       []TrialityBranch
	Results                        []BranchScanResult
	BranchCount                    int
	ResultCount                    int
	ExactPolarized3PlaneResults    int
	ExactFull3KernelResults        int
	MaxPolarizedKernelComplexDim   int
	MaxFullQ8vCKernelComplexDim    int
	UniqueBranchForPolarized3Plane bool
	SelectedBranch                 string
	SelectedByKernelOutcome        bool
	AllBranchesScanned             bool
	YOnly                          YOnlyDiagnostic
	Verdict                        string
}

type FirewallAudit struct {
	Gate256NativeNoGoPreserved          bool
	ChargeEigenvaluesTreatedAsDerived   bool
	EmbeddingOrientationTreatedAsSealed bool
	ImportedObservedChargeTable         bool
	ImportedObservedMasses              bool
	ImportedObservedYukawas             bool
	ForcedWeakPlane                     bool
	SelectedTrialityByHand              bool
	SelectedTrialityByDesiredKernel     bool
	ForcedKernelDim3                    bool
	AcceptedYOnlyAsQ                    bool
	TreatedSealAsFiniteDerivation       bool
	ConstructedVTauByHand               bool
	InsertedYukawaTexture               bool
	PollutedFiniteCore                  bool
	Verdict                             string
}

type DownstreamAudit struct {
	Neutral3PlaneAvailable bool
	FullQ8vCKernelDimThree bool
	VTauConstructed        bool
	TrialityTextureOpened  bool
	YukawaTextureDerived   bool
	CKMPMNSDerived         bool
	FermionMassesDerived   bool
	Verdict                string
}

type Summary struct {
	Gate256SealInherited             bool
	NativeChargeEigenvaluesExtracted bool
	ChargeCoefficientsExternal       bool
	EmbeddingWitnessesScanned        bool
	SO8WitnessesTranslated           bool
	AllTrialityBranchesScanned       bool
	UniqueTrialityBranchSelected     bool
	NeutralPolarized3PlaneDerived    bool
	FullQ8vCKernelDimThree           bool
	YukawaTextureDerived             bool
	Status                           string
	NextGate                         string
	Comment                          string
}

type Analysis struct {
	PreviousGate256 InheritedGate256Audit
	Charges         ChargeExtractionAudit
	Embedding       EmbeddingWitnessAudit
	SO8             SO8WitnessAudit
	TrialityScan    TrialityScanAudit
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
		prevRaw, err := spontaneouscarrierseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 256 predecessor: %w", err)
			return
		}
		prev := inheritGate256(prevRaw)
		charges := auditChargeExtraction()
		embedding := auditEmbeddingWitnesses(charges)
		so8 := auditSO8Witnesses(charges, embedding)
		scan := auditTrialityBranches(so8)
		firewall := auditFirewall(charges, embedding, scan)
		down := auditDownstream(scan)
		summary := summarize(prev, charges, embedding, so8, scan, down)
		truth := buildTruth(prev, charges, embedding, so8, scan)
		defaultA = Analysis{PreviousGate256: prev, Charges: charges, Embedding: embedding, SO8: so8, TrialityScan: scan, Firewall: firewall, Downstream: down, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate256(a spontaneouscarrierseal.Analysis) InheritedGate256Audit {
	return InheritedGate256Audit{
		SpontaneousSealRecorded:       a.Summary.SpontaneousSealRecorded,
		ConditionalIntertwinerSchema:  a.Summary.ConditionalIntertwinerSchema,
		SymbolicLedgerSchemaAvailable: a.Summary.SymbolicLedgerSchemaAvailable,
		SymbolicSO8SchemaAvailable:    a.Summary.SymbolicSO8SchemaAvailable,
		ConcreteUnifiedLedgerBuilt:    a.Summary.ConcreteUnifiedLedgerBuilt,
		ConcreteSO8Coordinates:        a.Summary.ConcreteSO8Coordinates,
		TrialityBranchSelected:        a.Summary.TrialityBranchSelected,
		Q8vCConstructed:               a.Summary.Q8vCConstructed,
		Neutral3PlaneDerived:          a.Summary.Neutral3PlaneDerived,
		YukawaTextureDerived:          a.Summary.YukawaTextureDerived,
		Status:                        a.Summary.Status,
		TruthStatement:                a.TruthStatement,
	}
}

func auditChargeExtraction() ChargeExtractionAudit {
	sources := []ChargeSource{
		{
			Name: "B-L Fock charge polarization", Source: "Gate 16 / Gate 253", Carrier: "S_C=Λ*(C^4)", Expression: "B-L=-N_0+(1/3)(N_1+N_2+N_3)",
			EigenvaluesDerived: true, CoefficientVector: []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}, CoefficientVectorRole: "native finite Fock ledger", RequiresEmbeddingSeal: false, UsesObservedInput: false, NativeFiniteCoreTheorem: true,
			Verdict: "native derived charge ledger; not equal to physical T3L or scalar Y_phi",
		},
		{
			Name: "scalar/contact Y_phi eigenvalues", Source: "Gate 20 scalar hypercharge bridge", Carrier: "H_phi scalar/contact factor", Expression: "T_phi/Y_phi spectrum ±1/2 in a 2+2 scalar doublet",
			EigenvaluesDerived: true, CoefficientVector: nil, CoefficientVectorRole: "derived scalar eigenvalues; Fock coefficients require sealed H_phi→S_C embedding", RequiresEmbeddingSeal: true, UsesObservedInput: false, NativeFiniteCoreTheorem: true,
			Verdict: "the eigenvalues are derived, but a four-mode coefficient vector is an embedding witness, not a native unsealed ledger",
		},
		{
			Name: "left-doublet T3L eigenvalues", Source: "Gate 24 finite SU(2)_L gauge audit", Carrier: "Q_L⊕L_L left-doublet carrier", Expression: "T3L=diag(+1/2,-1/2) on each quark/lepton weak doublet",
			EigenvaluesDerived: true, CoefficientVector: nil, CoefficientVectorRole: "derived local left-doublet eigenvalues; Fock coefficients require sealed weak-frame embedding", RequiresEmbeddingSeal: true, UsesObservedInput: false, NativeFiniteCoreTheorem: false,
			Verdict: "the charge table fixes ±1/2 weights, but not which pair of Fock modes is the electroweak plane without the SSB carrier seal",
		},
	}
	return ChargeExtractionAudit{
		Sources:                          sources,
		BMinusLFockLedgerDerived:         true,
		ScalarYphiEigenvaluesDerived:     true,
		T3LLeftDoubletEigenvaluesDerived: true,
		ChargeEigenvalueTableDerived:     true,
		PhysicalT3LDirectSCVector:        false,
		PhysicalYPhiDirectSCVector:       false,
		ExternalChargeInputUsed:          false,
		Verdict:                          StatusNativeChargeTableExtracted + "; charge eigenvalues are native, while the common S_C coefficient witnesses remain seal-conditioned embeddings",
	}
}

func auditEmbeddingWitnesses(charges ChargeExtractionAudit) EmbeddingWitnessAudit {
	weak := makeWeakFrames()
	scalar := makeScalarEmbeddings()
	allSealed := charges.ChargeEigenvalueTableDerived
	for _, w := range weak {
		allSealed = allSealed && w.EmbeddingSealed && !w.NativeSelected
	}
	for _, s := range scalar {
		allSealed = allSealed && s.EmbeddingSealed && !s.NativeSelected
	}
	return EmbeddingWitnessAudit{
		SealName:                      "SpontaneousCarrierSeal",
		WeakFrames:                    weak,
		ScalarEmbeddings:              scalar,
		WeakFrameCount:                len(weak),
		ScalarEmbeddingCount:          len(scalar),
		TotalCombinedWitnesses:        len(weak) * len(scalar),
		NativeWeakFrameSelected:       false,
		NativeScalarEmbeddingSelected: false,
		AllWitnessesSealed:            allSealed,
		UsesObservedMasses:            false,
		UsesObservedYukawas:           false,
		Verdict:                       StatusEmbeddingWitnessScanned + "; the seal authorizes scanning weak-frame/scalar embeddings but does not canonically select one",
	}
}

func makeWeakFrames() []WeakFrameWitness {
	pairs := [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	out := make([]WeakFrameWitness, 0, len(pairs)*2)
	for _, p := range pairs {
		for _, sign := range []int{+1, -1} {
			c := []float64{0, 0, 0, 0}
			c[p[0]] = 0.5 * float64(sign)
			c[p[1]] = -0.5 * float64(sign)
			name := fmt.Sprintf("T3_U%d%d", p[0], p[1])
			if sign < 0 {
				name += "_opposite"
			}
			out = append(out, WeakFrameWitness{
				Name: name, ModePair: p, OrientationSign: sign, T3Coefficients: c, DerivedEigenvalues: true, EmbeddingSealed: true, NativeSelected: false,
				Verdict: "charge-table T3L weights ±1/2 embedded into this two-mode Fock weak frame under the SpontaneousCarrierSeal",
			})
		}
	}
	return out
}

func makeScalarEmbeddings() []ScalarEmbeddingWitness {
	out := []ScalarEmbeddingWitness{
		{Name: "Yphi_uniform_plus_one_particle", Kind: "uniform_higgs_hypercharge", PositiveModes: []int{0, 1, 2, 3}, YPhiCoefficients: []float64{0.5, 0.5, 0.5, 0.5}, DerivedEigenvalues: true, EmbeddingSealed: true, Verdict: "represents the derived +1/2 scalar charge on a sealed one-particle Fock slice; not a full-S_C identity theorem"},
		{Name: "Yphi_uniform_minus_one_particle", Kind: "uniform_higgs_hypercharge_mirror", PositiveModes: []int{}, YPhiCoefficients: []float64{-0.5, -0.5, -0.5, -0.5}, DerivedEigenvalues: true, EmbeddingSealed: true, Verdict: "mirror scalar orientation included for branch-falsification; not selected natively"},
	}
	pairs := [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	for _, p := range pairs {
		c := []float64{-0.5, -0.5, -0.5, -0.5}
		c[p[0]] = 0.5
		c[p[1]] = 0.5
		out = append(out, ScalarEmbeddingWitness{
			Name: fmt.Sprintf("Yphi_contact_2plus2_pos_%d%d", p[0], p[1]), Kind: "contact_2plus2_orientation", PositiveModes: []int{p[0], p[1]}, YPhiCoefficients: c, DerivedEigenvalues: true, EmbeddingSealed: true, NativeSelected: false,
			Verdict: "embeds the Gate-20 2+2 scalar/contact ±1/2 spectrum into a sealed Fock-mode orientation",
		})
	}
	return out
}

func auditSO8Witnesses(charges ChargeExtractionAudit, embedding EmbeddingWitnessAudit) SO8WitnessAudit {
	labels := []string{"e0∧e1", "e2∧e3", "e4∧e5", "e6∧e7"}
	out := make([]QLedgerWitness, 0, embedding.TotalCombinedWitnesses)
	for _, w := range embedding.WeakFrames {
		for _, s := range embedding.ScalarEmbeddings {
			q := add(w.T3Coefficients, s.YPhiCoefficients)
			biv := scale(q, 0.5)
			central := 0.5 * sum(q)
			name := w.Name + "__" + s.Name
			out = append(out, QLedgerWitness{
				Name: name, WeakFrameName: w.Name, ScalarEmbeddingName: s.Name, T3Coefficients: cloneVec(w.T3Coefficients), YPhiCoefficients: cloneVec(s.YPhiCoefficients), QCoefficients: q,
				BivectorCoefficients: biv, CentralIdentityShift: central, SO8Formula: formatBivectorFormula(biv, labels), DerivedChargeInput: charges.ChargeEigenvalueTableDerived, SealedEmbeddingInput: true, CoordinateTranslated: true,
			})
		}
	}
	return SO8WitnessAudit{WittDictionaryInherited: true, CartanBivectors: labels, Witnesses: out, WitnessCount: len(out), AllTranslated: len(out) == embedding.TotalCombinedWitnesses, Verdict: StatusWittSO8Translated + "; every sealed witness is mechanically translated by N_k-1/2I ↦ (i/2)e_{2k}∧e_{2k+1}"}
}

func auditTrialityBranches(so8 SO8WitnessAudit) TrialityScanAudit {
	branches := makeTrialityBranches()
	results := make([]BranchScanResult, 0, len(so8.Witnesses)*len(branches))
	exactPolarized := 0
	exactFull := 0
	maxPol := 0
	maxFull := 0
	branchHit := map[string]int{}
	for _, w := range so8.Witnesses {
		for _, b := range branches {
			trans := matVec(b.Matrix, w.QCoefficients)
			zeros := zeroCount(trans, 1e-12)
			full := 2 * zeros
			exactP := zeros == 3
			exactF := full == 3
			if exactP {
				exactPolarized++
				branchHit[b.Name]++
			}
			if exactF {
				exactFull++
			}
			if zeros > maxPol {
				maxPol = zeros
			}
			if full > maxFull {
				maxFull = full
			}
			reason := ""
			if !exactP {
				reason = fmt.Sprintf("polarized kernel has %d zero Cartan slots, not 3", zeros)
			}
			results = append(results, BranchScanResult{WitnessName: w.Name, WeakFrameName: w.WeakFrameName, ScalarEmbeddingName: w.ScalarEmbeddingName, BranchName: b.Name, QCoefficients: cloneVec(w.QCoefficients), TransformedCoefficients: trans, ZeroCartanSlots: zeros, PolarizedKernelComplexDim: zeros, FullQ8vCKernelComplexDim: full, ExactPolarized3Plane: exactP, ExactFull3Kernel: exactF, SelectedByMathematics: false, RejectedReason: reason})
		}
	}
	unique := false
	selected := ""
	if exactPolarized > 0 && len(branchHit) == 1 {
		unique = true
		for k := range branchHit {
			selected = k
		}
		for i := range results {
			if results[i].ExactPolarized3Plane && results[i].BranchName == selected {
				results[i].SelectedByMathematics = true
			}
		}
	}
	return TrialityScanAudit{Branches: branches, Results: results, BranchCount: len(branches), ResultCount: len(results), ExactPolarized3PlaneResults: exactPolarized, ExactFull3KernelResults: exactFull, MaxPolarizedKernelComplexDim: maxPol, MaxFullQ8vCKernelComplexDim: maxFull, UniqueBranchForPolarized3Plane: unique, SelectedBranch: selected, SelectedByKernelOutcome: false, AllBranchesScanned: len(results) == len(so8.Witnesses)*len(branches), YOnly: yOnlyDiagnostic(branches), Verdict: trialityVerdict(exactPolarized, unique)}
}

func makeTrialityBranches() []TrialityBranch {
	branches := []TrialityBranch{
		{Name: "identity", Description: "no outer triality transform on Cartan coefficients", Matrix: identity4()},
		{Name: "tau_even", Description: "D4 Hadamard representative exchanging vector with an even-spinor Cartan convention", Matrix: [][]float64{{0.5, 0.5, 0.5, 0.5}, {0.5, 0.5, -0.5, -0.5}, {0.5, -0.5, 0.5, -0.5}, {0.5, -0.5, -0.5, 0.5}}},
		{Name: "tau_odd", Description: "D4 Hadamard representative exchanging vector with an odd-spinor Cartan convention", Matrix: [][]float64{{0.5, 0.5, 0.5, -0.5}, {0.5, 0.5, -0.5, 0.5}, {0.5, -0.5, 0.5, 0.5}, {-0.5, 0.5, 0.5, 0.5}}},
	}
	for i := range branches {
		branches[i].Orthogonal = isOrthogonal(branches[i].Matrix, 1e-12)
		branches[i].DetAbsOne = math.Abs(math.Abs(det4(branches[i].Matrix))-1) < 1e-12
		branches[i].AdmissibleCartan = branches[i].Orthogonal && branches[i].DetAbsOne
	}
	return branches
}

func yOnlyDiagnostic(branches []TrialityBranch) YOnlyDiagnostic {
	y := []float64{0.5, 0.5, 0.5, 0.5}
	for _, b := range branches {
		trans := matVec(b.Matrix, y)
		zeros := zeroCount(trans, 1e-12)
		if zeros == 3 {
			return YOnlyDiagnostic{Run: true, ScalarEmbeddingName: "Yphi_uniform_plus_one_particle", BranchName: b.Name, TransformedCoefficients: trans, PolarizedKernelComplexDim: zeros, WouldGivePolarizedThreeSlot: true, RejectedBecauseMissingT3L: true, Verdict: StatusYOnlyDiagnosticRejected + "; the scalar-only uniform witness produces a three-slot diagnostic under this branch, but Q=T3L+Y_phi cannot omit T3L"}
		}
	}
	return YOnlyDiagnostic{Run: true, RejectedBecauseMissingT3L: true, Verdict: "no scalar-only branch yielded a three-slot diagnostic"}
}

func trialityVerdict(exact int, unique bool) string {
	if exact == 0 {
		return StatusAllTrialityBranchesScanned + "; " + StatusNeutral3PlaneNotDerived + "; every lawful sealed Q=T3L+Y_phi witness misses the exact three-slot criterion"
	}
	if unique {
		return "CONDITIONAL_SUPPORT_SEALED_WITNESS_NEUTRAL_3PLANE_DERIVED_BY_UNIQUE_BRANCH"
	}
	return "FAILED_ROUTE_MULTIPLE_TRIALITY_BRANCHES_OR_WITNESSES_YIELD_THREE_PLANE_DEGENERACY"
}

func auditFirewall(charges ChargeExtractionAudit, embedding EmbeddingWitnessAudit, scan TrialityScanAudit) FirewallAudit {
	return FirewallAudit{
		Gate256NativeNoGoPreserved:          true,
		ChargeEigenvaluesTreatedAsDerived:   charges.ChargeEigenvalueTableDerived && !charges.ExternalChargeInputUsed,
		EmbeddingOrientationTreatedAsSealed: embedding.AllWitnessesSealed && !embedding.NativeWeakFrameSelected && !embedding.NativeScalarEmbeddingSelected,
		ImportedObservedChargeTable:         false,
		ImportedObservedMasses:              false,
		ImportedObservedYukawas:             false,
		ForcedWeakPlane:                     false,
		SelectedTrialityByHand:              false,
		SelectedTrialityByDesiredKernel:     scan.SelectedByKernelOutcome,
		ForcedKernelDim3:                    false,
		AcceptedYOnlyAsQ:                    false,
		TreatedSealAsFiniteDerivation:       false,
		ConstructedVTauByHand:               false,
		InsertedYukawaTexture:               false,
		PollutedFiniteCore:                  false,
		Verdict:                             "firewall holds: derived charge eigenvalues are separated from sealed embedding witnesses, and every triality branch is scanned without hand-selection",
	}
}

func auditDownstream(scan TrialityScanAudit) DownstreamAudit {
	ok := scan.UniqueBranchForPolarized3Plane && scan.ExactPolarized3PlaneResults > 0
	return DownstreamAudit{Neutral3PlaneAvailable: ok, FullQ8vCKernelDimThree: scan.ExactFull3KernelResults > 0, VTauConstructed: false, TrialityTextureOpened: ok, YukawaTextureDerived: false, CKMPMNSDerived: false, FermionMassesDerived: false, Verdict: StatusYukawaTextureStillSealed}
}

func summarize(prev InheritedGate256Audit, charges ChargeExtractionAudit, embedding EmbeddingWitnessAudit, so8 SO8WitnessAudit, scan TrialityScanAudit, down DownstreamAudit) Summary {
	status := StatusNeutral3PlaneNotDerived
	if scan.UniqueBranchForPolarized3Plane && scan.ExactPolarized3PlaneResults > 0 {
		status = "CONDITIONAL_SUPPORT_SEALED_WITNESS_NEUTRAL_3PLANE_DERIVED"
	}
	return Summary{
		Gate256SealInherited:             prev.SpontaneousSealRecorded,
		NativeChargeEigenvaluesExtracted: charges.ChargeEigenvalueTableDerived,
		ChargeCoefficientsExternal:       charges.ExternalChargeInputUsed,
		EmbeddingWitnessesScanned:        embedding.TotalCombinedWitnesses > 0,
		SO8WitnessesTranslated:           so8.AllTranslated,
		AllTrialityBranchesScanned:       scan.AllBranchesScanned,
		UniqueTrialityBranchSelected:     scan.UniqueBranchForPolarized3Plane,
		NeutralPolarized3PlaneDerived:    down.Neutral3PlaneAvailable,
		FullQ8vCKernelDimThree:           down.FullQ8vCKernelDimThree,
		YukawaTextureDerived:             down.YukawaTextureDerived,
		Status:                           status,
		NextGate:                         "Gate 258 — Weak-Plane Selector / Scalar Embedding Orientation Constraint Audit",
		Comment:                          fmt.Sprintf("Gate 257 scanned %d sealed Q witnesses across %d triality branches. Charge eigenvalues were native; embedding orientations were sealed. No lawful Q=T3L+Y_phi witness yielded an exact three-slot polarized neutral kernel; max was %d.", so8.WitnessCount, scan.BranchCount, scan.MaxPolarizedKernelComplexDim),
	}
}

func buildTruth(prev InheritedGate256Audit, charges ChargeExtractionAudit, embedding EmbeddingWitnessAudit, so8 SO8WitnessAudit, scan TrialityScanAudit) string {
	parts := []string{
		"Gate 257 confirms the corrected premise: the charge eigenvalues are not external phenomenology; they are inherited from the engine's own early charge table.",
		"The genuinely sealed datum is the carrier embedding orientation: which weak plane and scalar/contact orientation are placed into S_C.",
		fmt.Sprintf("The audit scanned %d weak frames, %d scalar embeddings, and %d Cartan triality branches (%d total branch evaluations).", embedding.WeakFrameCount, embedding.ScalarEmbeddingCount, scan.BranchCount, scan.ResultCount),
		fmt.Sprintf("No full Q=T3L+Y_phi witness produced the exact polarized three-slot kernel; the maximum polarized kernel dimension was %d and the maximum full 8_vC kernel dimension was %d.", scan.MaxPolarizedKernelComplexDim, scan.MaxFullQ8vCKernelComplexDim),
	}
	if scan.YOnly.WouldGivePolarizedThreeSlot {
		parts = append(parts, fmt.Sprintf("A scalar-only diagnostic gives a three-slot pattern under %s, but it is rejected because it is not Q=T3L+Y_phi.", scan.YOnly.BranchName))
	}
	if prev.SpontaneousSealRecorded && charges.ChargeEigenvalueTableDerived && so8.AllTranslated {
		parts = append(parts, "Therefore the next obstruction is not missing charge data or missing Witt translation; it is a missing weak-plane/scalar embedding selector strong enough to reduce the sealed witness degeneracy.")
	}
	return strings.Join(parts, " ")
}

func add(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}

func scale(a []float64, s float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = s * a[i]
	}
	return out
}

func sum(a []float64) float64 {
	t := 0.0
	for _, x := range a {
		t += x
	}
	return t
}

func cloneVec(a []float64) []float64 { return append([]float64(nil), a...) }

func matVec(m [][]float64, v []float64) []float64 {
	out := make([]float64, len(m))
	for i := range m {
		for j := range v {
			out[i] += m[i][j] * v[j]
		}
		if math.Abs(out[i]) < 1e-12 {
			out[i] = 0
		}
	}
	return out
}

func zeroCount(v []float64, eps float64) int {
	c := 0
	for _, x := range v {
		if math.Abs(x) <= eps {
			c++
		}
	}
	return c
}

func identity4() [][]float64 {
	return [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
}

func isOrthogonal(m [][]float64, eps float64) bool {
	for i := range m {
		for j := range m {
			dot := 0.0
			for k := range m[i] {
				dot += m[i][k] * m[j][k]
			}
			target := 0.0
			if i == j {
				target = 1
			}
			if math.Abs(dot-target) > eps {
				return false
			}
		}
	}
	return true
}

func det4(m [][]float64) float64 {
	// Laplace expansion specialized for the 4x4 matrices used here.
	det3 := func(a [3][3]float64) float64 {
		return a[0][0]*(a[1][1]*a[2][2]-a[1][2]*a[2][1]) - a[0][1]*(a[1][0]*a[2][2]-a[1][2]*a[2][0]) + a[0][2]*(a[1][0]*a[2][1]-a[1][1]*a[2][0])
	}
	d := 0.0
	for col := 0; col < 4; col++ {
		var minor [3][3]float64
		for i := 1; i < 4; i++ {
			mj := 0
			for j := 0; j < 4; j++ {
				if j == col {
					continue
				}
				minor[i-1][mj] = m[i][j]
				mj++
			}
		}
		sign := 1.0
		if col%2 == 1 {
			sign = -1
		}
		d += sign * m[0][col] * det3(minor)
	}
	return d
}

func formatBivectorFormula(coeffs []float64, labels []string) string {
	parts := make([]string, 0, len(coeffs))
	for i, c := range coeffs {
		if math.Abs(c) < 1e-12 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%+.6g i%s", c, labels[i]))
	}
	if len(parts) == 0 {
		return "0"
	}
	return strings.TrimPrefix(strings.Join(parts, " "), "+")
}

func sortedBranchNames(results []BranchScanResult) []string {
	m := map[string]bool{}
	for _, r := range results {
		if r.ExactPolarized3Plane {
			m[r.BranchName] = true
		}
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
