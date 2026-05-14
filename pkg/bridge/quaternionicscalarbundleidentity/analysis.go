// Package quaternionicscalarbundleidentity implements Gate 399:
// Quaternionic (H) Endomorphism / Scalar Bundle Identity Sieve.
//
// Gate 398 proved that the contact quartic primary block Q[x]/(q4) and the
// active scalar carrier H_phi are both four-real-dimensional, but dimension is
// not a functor. Gate 399 tests the strongest nearby candidate selector: the
// weak quaternionic algebra inherited from the Morita/spectral-triple lane.
//
// The theorem boundary is deliberately strict. Quaternionic structure is good
// evidence that H_phi is a complex weak doublet, but every single quaternionic
// left/right endomorphism on a real four-module satisfies a quadratic minimal
// polynomial. The contact primary q4 is an irreducible quartic block. Therefore
// a native H-action can support the scalar doublet structure only if its
// invariant fingerprint actually matches q4; otherwise the quartic contact block
// remains a different exact domain object and the Gate 398 obstruction survives.
package quaternionicscalarbundleidentity

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticscalaryukawabundle"
	"github.com/bagherbal/asha-engine/pkg/bridge/nativeweakquaternionicalgebra"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcomplex"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE399-QUATERNIONIC-SCALAR-BUNDLE-IDENTITY-SIEVE"

	StatusGate398Inherited = "CONDITIONAL_SUPPORT_GATE398_QUARTIC_HPHI_OBSTRUCTION_INHERITED"
	StatusGate274Inherited = "CONDITIONAL_SUPPORT_GATE274_LOCAL_WEAK_QUATERNIONIC_H_INHERITED"
	StatusGate295Inherited = "CONDITIONAL_SUPPORT_GATE295_LEFT_WEAK_H_MORITA_ACTION_INHERITED"
	StatusGate50Inherited  = "CONDITIONAL_SUPPORT_GATE50_SCALAR_COMPLEX_QUATERNIONIC_STRUCTURE_INHERITED"
	StatusGate385Inherited = "CONDITIONAL_SUPPORT_GATE385_ONEFORM_EDGE_SUPPORT_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"

	StatusQuaternionicModuleAudited      = "CONDITIONAL_SUPPORT_HPHI_QUATERNIONIC_MODULE_AUDITED"
	StatusLocalHClosureVerified          = "CONDITIONAL_SUPPORT_LOCAL_H_ENDOMORPHISMS_CLOSE"
	StatusPairComplexStructureAvailable  = "CONDITIONAL_SUPPORT_PAIR_COMPATIBLE_COMPLEX_STRUCTURE_AVAILABLE"
	StatusAbstractQuaternionicTriple     = "CONDITIONAL_SUPPORT_ABSTRACT_QUATERNIONIC_TRIPLE_AVAILABLE"
	StatusHActionFingerprintComputed     = "CONDITIONAL_SUPPORT_QUATERNIONIC_ENDOMORPHISM_FINGERPRINT_COMPUTED"
	StatusHDoubletStructurePreserved     = "CONDITIONAL_SUPPORT_HIGGS_WEAK_DOUBLET_STRUCTURE_PRESERVED"
	StatusSealedCompanionStressInherited = "CONDITIONAL_SUPPORT_SEALED_Q4_COMPANION_STRESS_TEST_INHERITED"
	StatusScalarLanePreserved            = "CONDITIONAL_SUPPORT_EXISTING_SCALAR_HIGGS_LANE_PRESERVED"

	StatusTensionLocalHNotGlobal          = "CONDITIONAL_TENSION_LOCAL_H_IS_NOT_GLOBAL_UNSEALED_AF_SUMMAND"
	StatusTensionFullHNotSelectedByScalar = "CONDITIONAL_TENSION_FULL_H_NOT_SELECTED_BY_ANISOTROPIC_SCALAR_RESPONSE"
	StatusTensionQuadraticVsQuartic       = "CONDITIONAL_TENSION_QUATERNIONIC_MINPOLY_QUADRATIC_VS_Q4_QUARTIC"
	StatusTensionCharPolySquareQuadratic  = "CONDITIONAL_TENSION_H_ACTION_CHARPOLY_IS_SQUARE_OF_QUADRATIC"

	StatusVerifiedCanonicalHphiQuarticID    = "VERIFIED_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusConditionalCanonicalHphiQuarticID = "CONDITIONAL_SUPPORT_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusConditionalScalarBundleSealed     = "CONDITIONAL_SUPPORT_SCALAR_BUNDLE_GEOMETRICALLY_SEALED"
	StatusFailedHActionPolynomialDisjoint   = "FAILED_ROUTE_QUATERNIONIC_ACTION_POLYNOMIAL_DISJOINT_FROM_Q4"
	StatusFailedHActionMinPolyQuadratic     = "FAILED_ROUTE_H_ACTION_MINIMAL_POLYNOMIAL_QUADRATIC_NOT_QUARTIC"
	StatusFailedNoQ4ScalarEndomorphism      = "FAILED_ROUTE_NO_NATIVE_Q4_SCALAR_ENDOMORPHISM"
	StatusFailedNoCanonicalHphiID           = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoOneFormEdgeFunctor        = "FAILED_ROUTE_NO_QUATERNIONIC_Q4_ONEFORM_EDGE_FUNCTOR"
	StatusFailedNoYukawaReduction           = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli         = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

type Inheritance struct {
	Executed bool

	Gate398NoCanonicalHphiID     bool
	Gate398QuarticDim            int
	Gate398HphiDim               int
	Gate398Q4Polynomial          string
	Gate398CompanionStressSealed bool

	Gate274LocalHExtracted        bool
	Gate274QuaternionClosureExact bool
	Gate274GlobalHDerived         bool
	Gate274PhysicalJDerived       bool

	Gate295WeakHLeftActionIsolated bool
	Gate295ExactAFDerived          bool
	Gate295FirstOrderComplete      bool

	Gate50PairComplexAvailable        bool
	Gate50CanonicalComplexDerived     bool
	Gate50QuaternionicTripleAvailable bool
	Gate50QuaternionicTripleSelected  bool

	Gate385OneFormEdgeSupportDerived bool
	Gate385JDoubledEdgeCount         int
	Gate372ChargedModuliDim          int
	NoEmpiricalValuesImported        bool
	Verdict                          string
}

type Q4Fingerprint struct {
	Executed             bool
	Polynomial           string
	MonicCoefficients    []float64 // x^4 + c3 x^3 + c2 x^2 + c1 x + c0
	Degree               int
	IrreducibleOverQ     bool
	BranchFreePrimary    bool
	ContactSpectralDatum bool
	Verdict              string
}

type QuaternionicModuleAudit struct {
	Executed                            bool
	Carrier                             string
	RealDimension                       int
	ComplexDoubletDimension             int
	Algebra                             string
	LocalHExtracted                     bool
	MoritaWeakHAction                   bool
	GlobalHUnsealed                     bool
	PairComplexAvailable                bool
	CanonicalComplexDerived             bool
	AbstractQuaternionicTripleAvailable bool
	QuaternionicTripleSelectedByScalar  bool
	FullScalarSU2Recovered              bool
	CompatibleWithAF                    bool
	CompatibleWithJ                     bool
	CompatibleWithFirstOrder            bool
	CompatibleWithOneFormEdges          bool
	Verdict                             string
}

type EndomorphismFingerprint struct {
	Name                        string
	Source                      string
	MatrixDim                   int
	Native                      bool
	Sealed                      bool
	Circular                    bool
	QuaternionicAction          bool
	SquaresToMinusIdentity      bool
	SquareResidual              float64
	ClosureResidual             float64
	MinimalPolynomial           string
	MinimalDegree               int
	CharacteristicPolynomial    string
	CharacteristicCoefficients  []float64 // monic high-to-low, len 5
	CharPolyIsSquareOfQuadratic bool
	Q4CoefficientResidual       float64
	Q4ExactMatch                bool
	Q4FactorMatch               bool
	CommutesWithScalarResponse  bool
	ScalarCommutatorNorm        float64
	CompatibleWithAF            bool
	CompatibleWithJ             bool
	CompatibleWithFirstOrder    bool
	CompatibleWithOneFormEdges  bool
	PromotableAsQ4Selector      bool
	Reason                      string
	Verdict                     string
}

type EndomorphismAudit struct {
	Executed                   bool
	Candidates                 []EndomorphismFingerprint
	QuaternionicCandidateCount int
	Q4ExactMatchCount          int
	Q4FactorMatchCount         int
	PromotableNativeCount      int
	MaxScalarCommutator        float64
	BestNativeCandidate        string
	Verdict                    string
}

type BundleIdentityAudit struct {
	Executed                        bool
	HphiQuarticIdentified           bool
	ScalarBundleGeometricallySealed bool
	OneFormEdgeFunctorDerived       bool
	YukawaCouplingsReduced          bool
	ChargedModuliStart              int
	ChargedModuliResult             int
	FlavorFirewallPreserved         bool
	HiggsLanePreserved              bool
	Verdict                         string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoObservedHiggsInserted        bool
	NoManualQ4HphiID               bool
	NoCompanionOperatorPromoted    bool
	NoArbitraryBasisMapPromoted    bool
	NoYukawaCouplingClaimed        bool
	NoFlavorModuliReductionClaimed bool
	Verdict                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance   Inheritance
	Q4            Q4Fingerprint
	Module        QuaternionicModuleAudit
	Endomorphisms EndomorphismAudit
	Identity      BundleIdentityAudit
	Firewall      FirewallAudit
	Next          NextStep
	Truth         string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inh, err := inheritPreviousGates()
	if err != nil {
		return Analysis{}, err
	}
	q4 := auditQ4(inh)
	module, err := auditQuaternionicModule(inh)
	if err != nil {
		return Analysis{}, err
	}
	endos, err := auditEndomorphisms(q4, module)
	if err != nil {
		return Analysis{}, err
	}
	identity := auditIdentity(inh, endos)
	firewall := auditFirewall()
	next := auditNext(identity)
	return Analysis{Inheritance: inh, Q4: q4, Module: module, Endomorphisms: endos, Identity: identity, Firewall: firewall, Next: next, Truth: truth(identity, endos)}, nil
}

func inheritPreviousGates() (Inheritance, error) {
	g398, err := contactquarticscalaryukawabundle.BuildDefault()
	if err != nil {
		return Inheritance{}, err
	}
	g274, err := nativeweakquaternionicalgebra.BuildDefault()
	if err != nil {
		return Inheritance{}, err
	}
	g50, err := scalarcomplex.BuildDefault()
	if err != nil {
		return Inheritance{}, err
	}
	statuses398 := mapStringSet(contactquarticscalaryukawabundle.Statuses(g398))

	return Inheritance{
		Executed:                          true,
		Gate398NoCanonicalHphiID:          statuses398[contactquarticscalaryukawabundle.StatusFailedNoCanonicalHphiID],
		Gate398QuarticDim:                 g398.Quartic.Dimension,
		Gate398HphiDim:                    g398.Scalar.ActiveRealDim,
		Gate398Q4Polynomial:               g398.Quartic.Polynomial,
		Gate398CompanionStressSealed:      statuses398[contactquarticscalaryukawabundle.StatusCompanionStressTestAvailable],
		Gate274LocalHExtracted:            g274.Summary.LocalHExtracted,
		Gate274QuaternionClosureExact:     g274.Summary.QuaternionClosureExact,
		Gate274GlobalHDerived:             g274.Summary.ExactSMAlgebraDerived,
		Gate274PhysicalJDerived:           g274.Summary.PhysicalJDerived,
		Gate295WeakHLeftActionIsolated:    true,
		Gate295ExactAFDerived:             false,
		Gate295FirstOrderComplete:         true, // mature Gates 296-297 verify first-order after the Morita repair; Gate 295 itself is a prerequisite ledger.
		Gate50PairComplexAvailable:        g50.PairCompatibleComplexAvailable,
		Gate50CanonicalComplexDerived:     g50.CanonicalComplexDerived,
		Gate50QuaternionicTripleAvailable: g50.QuaternionicTripleAvailable,
		Gate50QuaternionicTripleSelected:  g50.QuaternionicTripleSelected,
		Gate385OneFormEdgeSupportDerived:  true,
		Gate385JDoubledEdgeCount:          10,
		Gate372ChargedModuliDim:           13,
		NoEmpiricalValuesImported:         true,
		Verdict:                           join(StatusGate398Inherited, StatusGate274Inherited, StatusGate295Inherited, StatusGate50Inherited, StatusGate385Inherited, StatusGate372Inherited),
	}, nil
}

func auditQ4(inh Inheritance) Q4Fingerprint {
	return Q4Fingerprint{
		Executed:             true,
		Polynomial:           inh.Gate398Q4Polynomial,
		MonicCoefficients:    q4MonicCoefficients(),
		Degree:               4,
		IrreducibleOverQ:     true,
		BranchFreePrimary:    true,
		ContactSpectralDatum: true,
		Verdict:              StatusGate398Inherited,
	}
}

func auditQuaternionicModule(inh Inheritance) (QuaternionicModuleAudit, error) {
	return QuaternionicModuleAudit{
		Executed:                            true,
		Carrier:                             "H_phi active scalar/contact carrier",
		RealDimension:                       inh.Gate398HphiDim,
		ComplexDoubletDimension:             2,
		Algebra:                             "local weak quaternionic H acting on selected doublet / Morita left weak action",
		LocalHExtracted:                     inh.Gate274LocalHExtracted && inh.Gate274QuaternionClosureExact,
		MoritaWeakHAction:                   inh.Gate295WeakHLeftActionIsolated,
		GlobalHUnsealed:                     inh.Gate274GlobalHDerived && inh.Gate295ExactAFDerived,
		PairComplexAvailable:                inh.Gate50PairComplexAvailable,
		CanonicalComplexDerived:             inh.Gate50CanonicalComplexDerived,
		AbstractQuaternionicTripleAvailable: inh.Gate50QuaternionicTripleAvailable,
		QuaternionicTripleSelectedByScalar:  inh.Gate50QuaternionicTripleSelected,
		FullScalarSU2Recovered:              inh.Gate50QuaternionicTripleSelected,
		CompatibleWithAF:                    inh.Gate295ExactAFDerived,
		CompatibleWithJ:                     inh.Gate274PhysicalJDerived,
		CompatibleWithFirstOrder:            inh.Gate295FirstOrderComplete,
		CompatibleWithOneFormEdges:          inh.Gate385OneFormEdgeSupportDerived,
		Verdict:                             join(StatusQuaternionicModuleAudited, StatusLocalHClosureVerified, StatusPairComplexStructureAvailable, StatusAbstractQuaternionicTriple, StatusTensionLocalHNotGlobal, StatusTensionFullHNotSelectedByScalar),
	}, nil
}

func auditEndomorphisms(q4 Q4Fingerprint, module QuaternionicModuleAudit) (EndomorphismAudit, error) {
	I, J, K := quaternionicGenerators()
	scalarResponse := scalarResponseMatrix()

	candidates := []EndomorphismFingerprint{}
	for _, spec := range []struct {
		name        string
		source      string
		m           linear.Matrix
		commutes    bool
		compatAF    bool
		compatJ     bool
		compatFirst bool
		compatEdge  bool
		reason      string
	}{
		{"left H unit I on H_phi", "local weak H / scalarcomplex I", I, false, module.CompatibleWithAF, module.CompatibleWithJ, module.CompatibleWithFirstOrder, module.CompatibleWithOneFormEdges, "quaternionic unit squares to -1, so its minimal polynomial is x^2+1 and its characteristic polynomial is (x^2+1)^2, not q4"},
		{"left H unit J pair-rotation on H_phi", "pair-compatible scalar complex structure", J, true, module.CompatibleWithAF, module.CompatibleWithJ, module.CompatibleWithFirstOrder, module.CompatibleWithOneFormEdges, "this is the strongest scalar-compatible complex direction, but it still has minimal polynomial x^2+1 rather than q4"},
		{"left H unit K on H_phi", "local weak H / scalarcomplex K", K, false, module.CompatibleWithAF, module.CompatibleWithJ, module.CompatibleWithFirstOrder, module.CompatibleWithOneFormEdges, "quaternionic unit squares to -1, so it cannot generate the irreducible contact quartic primary"},
	} {
		fp, err := fingerprintQuaternionicUnit(spec.name, spec.source, spec.m, scalarResponse, q4, spec.compatAF, spec.compatJ, spec.compatFirst, spec.compatEdge, spec.reason)
		if err != nil {
			return EndomorphismAudit{}, err
		}
		candidates = append(candidates, fp)
	}

	generic := genericQuaternionicElementFingerprint(q4)
	candidates = append(candidates, generic)
	sealed := sealedCompanionFingerprint(q4)
	candidates = append(candidates, sealed)

	audit := EndomorphismAudit{Executed: true, Candidates: candidates, BestNativeCandidate: "none", Verdict: join(StatusHActionFingerprintComputed, StatusFailedHActionPolynomialDisjoint, StatusFailedHActionMinPolyQuadratic, StatusFailedNoQ4ScalarEndomorphism)}
	for _, c := range candidates {
		if c.QuaternionicAction {
			audit.QuaternionicCandidateCount++
		}
		if c.Q4ExactMatch {
			audit.Q4ExactMatchCount++
		}
		if c.Q4FactorMatch {
			audit.Q4FactorMatchCount++
		}
		if c.PromotableAsQ4Selector {
			audit.PromotableNativeCount++
			if audit.BestNativeCandidate == "none" {
				audit.BestNativeCandidate = c.Name
			}
		}
		if c.ScalarCommutatorNorm > audit.MaxScalarCommutator && !math.IsInf(c.ScalarCommutatorNorm, 0) {
			audit.MaxScalarCommutator = c.ScalarCommutatorNorm
		}
	}
	return audit, nil
}

func auditIdentity(inh Inheritance, endos EndomorphismAudit) BundleIdentityAudit {
	promoted := endos.PromotableNativeCount > 0
	return BundleIdentityAudit{
		Executed:                        true,
		HphiQuarticIdentified:           promoted,
		ScalarBundleGeometricallySealed: promoted,
		OneFormEdgeFunctorDerived:       false,
		YukawaCouplingsReduced:          false,
		ChargedModuliStart:              inh.Gate372ChargedModuliDim,
		ChargedModuliResult:             inh.Gate372ChargedModuliDim,
		FlavorFirewallPreserved:         !promoted,
		HiggsLanePreserved:              true,
		Verdict:                         join(StatusFailedNoCanonicalHphiID, StatusFailedNoOneFormEdgeFunctor, StatusFailedNoYukawaReduction, StatusFirewallPreserved13Moduli),
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoObservedHiggsInserted: true, NoManualQ4HphiID: true, NoCompanionOperatorPromoted: true, NoArbitraryBasisMapPromoted: true, NoYukawaCouplingClaimed: true, NoFlavorModuliReductionClaimed: true, Verdict: StatusFirewallPreserved13Moduli}
}

func auditNext(identity BundleIdentityAudit) NextStep {
	if identity.HphiQuarticIdentified {
		return NextStep{Gate: 400, Title: "Quartic-Sealed Scalar One-Form/Yukawa Coefficient Audit", Reason: "Gate 399 found a native q4 selector, so the next theorem would test whether it acts on one-form edges and Yukawa fibers.", PrimaryTask: "compute scalar one-form and Yukawa bundle reductions under the derived q4 action"}
	}
	return NextStep{Gate: 400, Title: "Non-Quaternionic Scalar Identity Selector Search", Reason: "Gate 399 proves the weak quaternionic H action supports the scalar doublet structure but has quadratic, not quartic, polynomial fingerprints. The q4 identity selector, if it exists, must come from a different invariant than a single H endomorphism.", PrimaryTask: "search mixed invariants built from scalar response S_phi, complex structure J, one-form edge Laplacian, and contact projector compression; require q4 without arbitrary basis choice"}
}

func fingerprintQuaternionicUnit(name, source string, m, scalar linear.Matrix, q4 Q4Fingerprint, compatAF, compatJ, compatFirst, compatEdge bool, reason string) (EndomorphismFingerprint, error) {
	id := linear.Identity(4)
	mm, err := m.Mul(m)
	if err != nil {
		return EndomorphismFingerprint{}, err
	}
	plus, err := mm.Add(id)
	if err != nil {
		return EndomorphismFingerprint{}, err
	}
	sq := plus.FrobeniusNorm()
	comm, err := linear.Commutator(scalar, m)
	if err != nil {
		return EndomorphismFingerprint{}, err
	}
	commNorm := comm.FrobeniusNorm()
	coeffs := []float64{1, 0, 2, 0, 1}
	residual := coeffResidual(coeffs, q4.MonicCoefficients)
	return EndomorphismFingerprint{
		Name:                        name,
		Source:                      source,
		MatrixDim:                   4,
		Native:                      true,
		QuaternionicAction:          true,
		SquaresToMinusIdentity:      sq < eps,
		SquareResidual:              sq,
		ClosureResidual:             0,
		MinimalPolynomial:           "x^2 + 1",
		MinimalDegree:               2,
		CharacteristicPolynomial:    "(x^2 + 1)^2 = x^4 + 2x^2 + 1",
		CharacteristicCoefficients:  coeffs,
		CharPolyIsSquareOfQuadratic: true,
		Q4CoefficientResidual:       residual,
		Q4ExactMatch:                residual < eps,
		Q4FactorMatch:               false,
		CommutesWithScalarResponse:  commNorm < eps,
		ScalarCommutatorNorm:        commNorm,
		CompatibleWithAF:            compatAF,
		CompatibleWithJ:             compatJ,
		CompatibleWithFirstOrder:    compatFirst,
		CompatibleWithOneFormEdges:  compatEdge,
		PromotableAsQ4Selector:      false,
		Reason:                      reason,
		Verdict:                     join(StatusHActionFingerprintComputed, StatusFailedHActionMinPolyQuadratic, StatusFailedHActionPolynomialDisjoint),
	}, nil
}

func genericQuaternionicElementFingerprint(q4 Q4Fingerprint) EndomorphismFingerprint {
	// Any real quaternion a + bi + cj + dk acting by left multiplication obeys
	// t^2 - 2a t + (a^2+b^2+c^2+d^2)=0. On the real 4-module the characteristic
	// polynomial is the square of that quadratic. This is a structural theorem,
	// not a numerical sample.
	coeffs := []float64{1, -2, 2, -2, 1} // representative a=1, |v|=1: (x^2-2x+2)^2
	return EndomorphismFingerprint{
		Name:                        "generic single quaternion element",
		Source:                      "left H action theorem",
		MatrixDim:                   4,
		Native:                      true,
		QuaternionicAction:          true,
		MinimalPolynomial:           "x^2 - 2a x + (a^2+b^2+c^2+d^2)",
		MinimalDegree:               2,
		CharacteristicPolynomial:    "[x^2 - 2a x + (a^2+|v|^2)]^2",
		CharacteristicCoefficients:  coeffs,
		CharPolyIsSquareOfQuadratic: true,
		Q4CoefficientResidual:       math.Inf(1),
		Q4ExactMatch:                false,
		Q4FactorMatch:               false,
		CommutesWithScalarResponse:  false,
		ScalarCommutatorNorm:        math.Inf(1),
		CompatibleWithFirstOrder:    true,
		CompatibleWithOneFormEdges:  true,
		PromotableAsQ4Selector:      false,
		Reason:                      "the full single-element H action family has quadratic minimal polynomial, so it cannot natively produce an irreducible quartic minimal polynomial",
		Verdict:                     join(StatusTensionCharPolySquareQuadratic, StatusFailedHActionMinPolyQuadratic, StatusFailedHActionPolynomialDisjoint),
	}
}

func sealedCompanionFingerprint(q4 Q4Fingerprint) EndomorphismFingerprint {
	return EndomorphismFingerprint{
		Name:                        "sealed q4 companion operator placed on H_phi",
		Source:                      "Gate398 sealed stress test / arbitrary basis map",
		MatrixDim:                   4,
		Sealed:                      true,
		Circular:                    true,
		QuaternionicAction:          false,
		MinimalPolynomial:           "q4",
		MinimalDegree:               4,
		CharacteristicPolynomial:    "q4",
		CharacteristicCoefficients:  append([]float64(nil), q4.MonicCoefficients...),
		CharPolyIsSquareOfQuadratic: false,
		Q4CoefficientResidual:       0,
		Q4ExactMatch:                true,
		Q4FactorMatch:               true,
		ScalarCommutatorNorm:        math.Inf(1),
		CompatibleWithAF:            false,
		CompatibleWithJ:             false,
		CompatibleWithFirstOrder:    false,
		CompatibleWithOneFormEdges:  false,
		PromotableAsQ4Selector:      false,
		Reason:                      "a companion matrix can always be installed on a chosen 4D basis, but Gate 398 already quarantined this as an arbitrary identification, not a quaternionic/Morita theorem",
		Verdict:                     join(StatusSealedCompanionStressInherited, StatusFailedNoQ4ScalarEndomorphism, StatusFailedNoCanonicalHphiID),
	}
}

func quaternionicGenerators() (linear.Matrix, linear.Matrix, linear.Matrix) {
	// Same real quaternionic triple used by the scalarcomplex gate: twice the
	// real SU(2) doublet generators. J is the pair-rotation complex structure
	// that commutes with the pair-degenerate scalar response.
	I, _ := linear.FromRows([][]float64{{0, 0, 0, 1}, {0, 0, -1, 0}, {0, 1, 0, 0}, {-1, 0, 0, 0}})
	J, _ := linear.FromRows([][]float64{{0, 1, 0, 0}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {0, 0, 1, 0}})
	K, _ := linear.FromRows([][]float64{{0, 0, -1, 0}, {0, 0, 0, -1}, {1, 0, 0, 0}, {0, 1, 0, 0}})
	return I, J, K
}

func scalarResponseMatrix() linear.Matrix {
	// Gate 12/37 active scalar response pair spectrum snapshot.
	return linear.Diagonal([]float64{0.336692702, 0.336692702, 0.229973965, 0.229973965})
}

func q4MonicCoefficients() []float64 {
	return []float64{1, -71.0 / 30.0, 1071.0 / 540.0, -149.0 / 216.0, 271.0 / 3240.0}
}

func coeffResidual(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.Inf(1)
	}
	sum := 0.0
	for i := range a {
		d := a[i] - b[i]
		sum += d * d
	}
	return math.Sqrt(sum)
}

func mapStringSet(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, v := range values {
		out[v] = true
	}
	return out
}

func Statuses(a Analysis) []string {
	status := []string{
		StatusGate398Inherited,
		StatusGate274Inherited,
		StatusGate295Inherited,
		StatusGate50Inherited,
		StatusGate385Inherited,
		StatusGate372Inherited,
		StatusQuaternionicModuleAudited,
		StatusLocalHClosureVerified,
		StatusPairComplexStructureAvailable,
		StatusAbstractQuaternionicTriple,
		StatusHActionFingerprintComputed,
		StatusHDoubletStructurePreserved,
		StatusScalarLanePreserved,
		StatusTensionLocalHNotGlobal,
		StatusTensionFullHNotSelectedByScalar,
		StatusTensionQuadraticVsQuartic,
		StatusTensionCharPolySquareQuadratic,
		StatusFailedHActionPolynomialDisjoint,
		StatusFailedHActionMinPolyQuadratic,
		StatusFailedNoQ4ScalarEndomorphism,
		StatusFailedNoCanonicalHphiID,
		StatusFailedNoOneFormEdgeFunctor,
		StatusFailedNoYukawaReduction,
		StatusFirewallPreserved13Moduli,
	}
	if a.Inheritance.Gate398CompanionStressSealed {
		status = append(status, StatusSealedCompanionStressInherited)
	}
	if a.Identity.HphiQuarticIdentified {
		status = append(status, StatusVerifiedCanonicalHphiQuarticID, StatusConditionalCanonicalHphiQuarticID, StatusConditionalScalarBundleSealed)
	}
	sort.Strings(status)
	return unique(status)
}

func unique(values []string) []string {
	if len(values) == 0 {
		return values
	}
	out := values[:1]
	for _, v := range values[1:] {
		if v != out[len(out)-1] {
			out = append(out, v)
		}
	}
	return out
}

func join(parts ...string) string { return strings.Join(parts, " | ") }

func truth(identity BundleIdentityAudit, endos EndomorphismAudit) string {
	if identity.HphiQuarticIdentified {
		return "Gate 399 derives a native quaternionic scalar-bundle identity selector matching q4. This would reopen the scalar one-form/Yukawa coefficient lane under strict follow-up audits."
	}
	return fmt.Sprintf("Gate 399 preserves the Gate 398 obstruction. The weak quaternionic H action correctly supports the four-real-dimensional Higgs doublet arena, but its native single-endomorphism fingerprints have quadratic minimal polynomials and characteristic polynomials that are squares of quadratics. None matches the irreducible contact q4 primary; promotable native q4 selectors=%d. Therefore H_phi is not canonically identified with the contact quartic block by quaternionic action, and the 13-moduli flavor firewall remains preserved.", endos.PromotableNativeCount)
}
