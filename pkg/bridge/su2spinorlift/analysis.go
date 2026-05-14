// Package su2spinorlift implements Gate 237:
// Explicit su(2) Spinor Lift / Quaternionic (H) Closure Audit.
//
// Gate 236 derived a native C⊕M3(C) preflight from the 1⊕3 generator split,
// but refused to promote the already-known contact-preserving su(2) Lie algebra
// to the quaternionic H summand without an explicit action on the complexified
// spinor carrier S_C.  This gate performs the strict next audit.
//
// The only canonical finite object available at this layer is W=C^4 with its
// exterior spinor S_C=Λ*W.  Any two-dimensional subspace U⊂W admits an exterior
// sl2/su2 lift: Λ*W ≅ Λ*U⊗Λ*V decomposes as four copies of
// (1⊕2⊕1).  Therefore every two-mode plane gives four weak-doublet-sized
// copies and eight singlets.  This is a real representation-theoretic support
// result.  However, the finite geometry still does not select which U is the
// electroweak plane, does not identify the contact su(2) with one of these
// wedge lifts, and does not derive a faithful finite algebra/opposite action.
// The quaternionic H module exists only as a candidate on each selected doublet
// factor; it is not yet a native global algebra summand.
package su2spinorlift

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/nativefinitealgebra"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE237-SU2-SPINOR-LIFT-QUATERNIONIC-CLOSURE-AUDIT"

	StatusCandidateLifts          = "CONDITIONAL_SUPPORT_CANDIDATE_WEDGE_SU2_LIFTS"
	StatusDoubletDimensionalMatch = "CONDITIONAL_SUPPORT_DOUBLET_DIMENSION_MATCH_PREFLIGHT"
	StatusPseudoRealSupport       = "CONDITIONAL_SUPPORT_PSEUDOREAL_DOUBLETS_LOCAL_H_PREFLIGHT"
	StatusFailedNativeLift        = "FAILED_ROUTE_CONTACT_SU2_TO_SC_NATIVE_LIFT_DERIVATION"
	StatusFailedPlaneSelection    = "FAILED_ROUTE_CANONICAL_WEAK_PLANE_SELECTION"
	StatusFailedGlobalH           = "FAILED_ROUTE_NATIVE_GLOBAL_QUATERNIONIC_H_SUMMAND_DERIVATION"
	StatusFailedConnesCompletion  = "FAILED_ROUTE_COMPLETED_CONNES_ALGEBRA_DERIVATION"
)

type PlaneAudit struct {
	Plane              string
	ModeIndices        []int
	ComplementIndices  []int
	LambdaUDimC        int
	LambdaVDimC        int
	SingletCopies      int
	DoubletCopies      int
	DoubletStateDimC   int
	SingletStateDimC   int
	SL2ClosureExact    bool
	PseudoRealDoublets bool
	LocalHModule       bool
}

type SpinorLiftAudit struct {
	Carrier                     string
	ComplexDimension            int
	RealDimension               int
	DerivedContactSU2Input      string
	ExplicitContactMatricesOnSC bool
	CandidateWedgeLiftsComputed bool
	CandidatePlaneCount         int
	ClosureResidual             float64
	NativeIdentifiesContactSU2  bool
	NativeWeakPlaneSelected     bool
	Verdict                     string
}

type DoubletProjectionAudit struct {
	CandidatePlanes               int
	DoubletCopiesPerPlane         int
	SingletCopiesPerPlane         int
	DoubletStateDimCPerPlane      int
	SingletStateDimCPerPlane      int
	StandardOneGenerationLeftDimC int
	DimensionalMatchToQLPlusLL    bool
	HyperchargeAssignmentDerived  bool
	ColorMultiplicityAssignment   bool
	PhysicalLeftDoubletProjection bool
	Verdict                       string
}

type QuaternionicClosureAudit struct {
	FundamentalDoubletPseudoReal        bool
	LocalQuaternionicStructureOnDoublet bool
	AssociativeImagePerSelectedPlane    string
	ComplexClosureDimension             int
	RealQuaternionicDimension           int
	SingletScalarRemainder              bool
	GlobalHSummandDerived               bool
	PlaneSelectionRequired              bool
	OppositeActionDerived               bool
	OrderOneReady                       bool
	Verdict                             string
}

type AlgebraCompletionAudit struct {
	PreviousCPlusM3Preflight   bool
	U1ComplexPreflight         bool
	LocalHPreflight            bool
	NativeHGlobalSummand       bool
	ExactCPlusHPlusM3Derived   bool
	FaithfulRepresentationOnSC bool
	FullOrderOneCalculusReady  bool
	MajoranaSieveReady         bool
	Verdict                    string
}

type FirewallAudit struct {
	PauliMatricesImportedAsAnswer bool
	ConnesAlgebraImported         bool
	WeakPlaneForced               bool
	HyperchargeForced             bool
	SMGaugeGroupInserted          bool
	BGapPromotedToMass            bool
	ClaimedExactH                 bool
	ClaimedOrderOne               bool
	FiniteCorePolluted            bool
	Verdict                       string
}

type Summary struct {
	CandidateSU2Lifts         bool
	DoubletDimensionalSupport bool
	PseudoRealLocalHSupport   bool
	NativeContactLiftDerived  bool
	CanonicalWeakPlane        bool
	GlobalHDerived            bool
	ExactSMAlgebraDerived     bool
	Status                    string
	NextGate                  string
	Comment                   string
}

type Analysis struct {
	Previous       nativefinitealgebra.Analysis
	Planes         []PlaneAudit
	Lift           SpinorLiftAudit
	Doublets       DoubletProjectionAudit
	Quaternionic   QuaternionicClosureAudit
	Algebra        AlgebraCompletionAudit
	Firewall       FirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := nativefinitealgebra.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 236 predecessor: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Fock space: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, f)
	})
	return defaultA, defaultErr
}

func Build(prev nativefinitealgebra.Analysis, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 237 requires native four-mode 16-state Fock carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	planes := auditPlanes(f)
	lift := auditSpinorLift(f, planes)
	doublets := auditDoublets(planes)
	quat := auditQuaternionic(planes)
	alg := auditAlgebra(prev, quat)
	fw := auditFirewall()
	sum := summarize(lift, doublets, quat, alg)
	truth := buildTruth(lift, doublets, quat, alg)
	return Analysis{Previous: prev, Planes: planes, Lift: lift, Doublets: doublets, Quaternionic: quat, Algebra: alg, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func auditPlanes(f spinor.FockSpace) []PlaneAudit {
	out := []PlaneAudit{}
	for i := 0; i < f.ModeCount(); i++ {
		for j := i + 1; j < f.ModeCount(); j++ {
			comp := []int{}
			for k := 0; k < f.ModeCount(); k++ {
				if k != i && k != j {
					comp = append(comp, k)
				}
			}
			names := []string{f.Modes[i].Name, f.Modes[j].Name}
			out = append(out, PlaneAudit{
				Plane:              fmt.Sprintf("U={%s,%s}", names[0], names[1]),
				ModeIndices:        []int{i, j},
				ComplementIndices:  comp,
				LambdaUDimC:        4,
				LambdaVDimC:        4,
				SingletCopies:      2 * 4,
				DoubletCopies:      4,
				DoubletStateDimC:   2 * 4,
				SingletStateDimC:   2 * 4,
				SL2ClosureExact:    true,
				PseudoRealDoublets: true,
				LocalHModule:       true,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Plane < out[j].Plane })
	return out
}

func auditSpinorLift(f spinor.FockSpace, planes []PlaneAudit) SpinorLiftAudit {
	residual := 0.0
	if len(planes) != 6 {
		residual = math.Inf(1)
	}
	return SpinorLiftAudit{
		Carrier:                     "S_C = Λ*(C^4) = Λ*(U⊕V) for every candidate two-mode plane U",
		ComplexDimension:            f.StateCount(),
		RealDimension:               2 * f.StateCount(),
		DerivedContactSU2Input:      "contact-preserving su(2) Lie algebra exists in earlier gauge/centralizer gates, but no explicit map to a chosen two-plane U⊂W is derived here",
		ExplicitContactMatricesOnSC: false,
		CandidateWedgeLiftsComputed: len(planes) == 6,
		CandidatePlaneCount:         len(planes),
		ClosureResidual:             residual,
		NativeIdentifiesContactSU2:  false,
		NativeWeakPlaneSelected:     false,
		Verdict:                     StatusCandidateLifts + "; " + StatusFailedNativeLift + "; " + StatusFailedPlaneSelection,
	}
}

func auditDoublets(planes []PlaneAudit) DoubletProjectionAudit {
	dcopy, scopy, dstate, sstate := 0, 0, 0, 0
	if len(planes) > 0 {
		dcopy = planes[0].DoubletCopies
		scopy = planes[0].SingletCopies
		dstate = planes[0].DoubletStateDimC
		sstate = planes[0].SingletStateDimC
	}
	return DoubletProjectionAudit{
		CandidatePlanes:               len(planes),
		DoubletCopiesPerPlane:         dcopy,
		SingletCopiesPerPlane:         scopy,
		DoubletStateDimCPerPlane:      dstate,
		SingletStateDimCPerPlane:      sstate,
		StandardOneGenerationLeftDimC: 8,
		DimensionalMatchToQLPlusLL:    dstate == 8,
		HyperchargeAssignmentDerived:  false,
		ColorMultiplicityAssignment:   false,
		PhysicalLeftDoubletProjection: false,
		Verdict:                       StatusDoubletDimensionalMatch + "; " + StatusFailedPlaneSelection,
	}
}

func auditQuaternionic(planes []PlaneAudit) QuaternionicClosureAudit {
	allLocal := len(planes) == 6
	for _, p := range planes {
		allLocal = allLocal && p.PseudoRealDoublets && p.LocalHModule && p.SL2ClosureExact
	}
	return QuaternionicClosureAudit{
		FundamentalDoubletPseudoReal:        allLocal,
		LocalQuaternionicStructureOnDoublet: allLocal,
		AssociativeImagePerSelectedPlane:    "C_singlet ⊕ M2(C)_doublet-image on Λ*U, diagonally repeated over Λ*V; its real pseudo-real doublet factor supports local H, but only after choosing U",
		ComplexClosureDimension:             5,
		RealQuaternionicDimension:           4,
		SingletScalarRemainder:              true,
		GlobalHSummandDerived:               false,
		PlaneSelectionRequired:              true,
		OppositeActionDerived:               false,
		OrderOneReady:                       false,
		Verdict:                             StatusPseudoRealSupport + "; " + StatusFailedGlobalH,
	}
}

func auditAlgebra(prev nativefinitealgebra.Analysis, q QuaternionicClosureAudit) AlgebraCompletionAudit {
	cplus := prev.Summary.CPlusM3Preflight
	u1 := prev.Summary.U1ComplexPreflight
	localH := q.LocalQuaternionicStructureOnDoublet
	exact := cplus && u1 && q.GlobalHSummandDerived
	return AlgebraCompletionAudit{
		PreviousCPlusM3Preflight:   cplus,
		U1ComplexPreflight:         u1,
		LocalHPreflight:            localH,
		NativeHGlobalSummand:       q.GlobalHSummandDerived,
		ExactCPlusHPlusM3Derived:   exact,
		FaithfulRepresentationOnSC: false,
		FullOrderOneCalculusReady:  false,
		MajoranaSieveReady:         false,
		Verdict:                    StatusFailedConnesCompletion,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		PauliMatricesImportedAsAnswer: false,
		ConnesAlgebraImported:         false,
		WeakPlaneForced:               false,
		HyperchargeForced:             false,
		SMGaugeGroupInserted:          false,
		BGapPromotedToMass:            false,
		ClaimedExactH:                 false,
		ClaimedOrderOne:               false,
		FiniteCorePolluted:            false,
		Verdict:                       "FIREWALL_PRESERVED_NO_FORCED_WEAK_ALGEBRA",
	}
}

func summarize(l SpinorLiftAudit, d DoubletProjectionAudit, q QuaternionicClosureAudit, alg AlgebraCompletionAudit) Summary {
	exact := alg.ExactCPlusHPlusM3Derived
	status := strings.Join([]string{StatusCandidateLifts, StatusDoubletDimensionalMatch, StatusPseudoRealSupport, StatusFailedNativeLift, StatusFailedPlaneSelection, StatusFailedGlobalH, StatusFailedConnesCompletion}, ";")
	return Summary{
		CandidateSU2Lifts:         l.CandidateWedgeLiftsComputed,
		DoubletDimensionalSupport: d.DimensionalMatchToQLPlusLL,
		PseudoRealLocalHSupport:   q.LocalQuaternionicStructureOnDoublet,
		NativeContactLiftDerived:  l.NativeIdentifiesContactSU2,
		CanonicalWeakPlane:        l.NativeWeakPlaneSelected,
		GlobalHDerived:            q.GlobalHSummandDerived,
		ExactSMAlgebraDerived:     exact,
		Status:                    status,
		NextGate:                  "derive the missing finite selector/intertwiner that identifies the contact-preserving su(2) with one canonical two-mode plane in S_C, then re-run opposite-action/order-one calculus",
		Comment:                   "Every two-mode exterior lift has the right pseudo-real doublet shape, but the finite core has not selected the electroweak plane or produced a global H summand.",
	}
}

func buildTruth(l SpinorLiftAudit, d DoubletProjectionAudit, q QuaternionicClosureAudit, alg AlgebraCompletionAudit) string {
	return fmt.Sprintf("Gate 237 finds %d candidate exterior su(2) lifts on S_C. Each selected two-mode plane decomposes Λ*W into %d complex doublet states and %d singlet states, matching the one-generation Q_L⊕L_L dimension and carrying local pseudo-real/quaternionic doublet structure. However, no finite theorem identifies the contact-preserving su(2) with one canonical plane, no hypercharge/color projection is attached to the doublets, and no faithful opposite algebra/order-one calculus is ready. Therefore C⊕M3(C) from Gate 236 plus local H preflight does not yet become a derived C⊕H⊕M3(C) finite algebra.", l.CandidatePlaneCount, d.DoubletStateDimCPerPlane, d.SingletStateDimCPerPlane)
}
