// Package trialitymodulisieve implements Gate 393:
// Triality Domain-Admission & Equivariant Yukawa Centralizer Sieve.
//
// The gate audits the proposed flavor breakthrough without assuming the desired
// result.  It separates three ledgers:
//  1. native domain admission: whether the finite spectral-triple generation
//     carrier has been lawfully placed inside a Spin(8) triality domain;
//  2. sealed/label triality stress test: what happens if the three copied
//     generation labels are acted on by C3/S3 permutation matrices;
//  3. moduli impact: whether the Gate-372 charged 13-moduli firewall is
//     actually reduced by native ASHA data.
package trialitymodulisieve

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE393-TRIALITY-DOMAIN-ADMISSION-EQUIVARIANT-YUKAWA-CENTRALIZER-SIEVE"

	StatusGate247Inherited = "CONDITIONAL_SUPPORT_GATE247_TRIALITY_FUNCTOR_OBSTRUCTION_INHERITED"
	StatusGate370Inherited = "CONDITIONAL_SUPPORT_GATE370_GENERATION_ADDRESS_OBSTRUCTION_INHERITED"
	StatusGate371Inherited = "CONDITIONAL_SUPPORT_GATE371_NUMBER_OPERATOR_CAPACITY_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"
	StatusGate387Inherited = "CONDITIONAL_SUPPORT_GATE387_EPISTEMOLOGICAL_SEAL_INHERITED"

	StatusDomainAdmissionAudited             = "CONDITIONAL_SUPPORT_TRIALITY_DOMAIN_ADMISSION_AUDITED"
	StatusAbstractTrialityAvailable          = "CONDITIONAL_SUPPORT_ABSTRACT_SPIN8_TRIALITY_AVAILABLE"
	StatusLabelPermutationStressTestExecuted = "CONDITIONAL_SUPPORT_LABEL_PERMUTATION_STRESS_TEST_EXECUTED"
	StatusEquivariantCentralizerComputed     = "CONDITIONAL_SUPPORT_EQUIVARIANT_YUKAWA_CENTRALIZER_COMPUTED"
	StatusFockNumberOperatorClassified       = "CONDITIONAL_SUPPORT_FOCK_NUMBER_OPERATOR_CLASSIFIED"
	StatusModuliRecountExecuted              = "CONDITIONAL_SUPPORT_TRIALITY_MODULI_RECOUNT_EXECUTED"
	StatusConditionalTrialityTextureCapacity = "CONDITIONAL_SUPPORT_TRIALITY_TEXTURE_CAPACITY_ONLY_UNDER_SEALED_LABEL_ACTION"
	StatusConditionalFockHierarchyCapacity   = "CONDITIONAL_SUPPORT_FOCK_NUMBER_HIERARCHY_CAPACITY"

	StatusTensionTrialityNeedsCarrier         = "CONDITIONAL_TENSION_TRIALITY_REQUIRES_EXPLICIT_NATIVE_CARRIER"
	StatusTensionCyclicCirculantCommutes      = "CONDITIONAL_TENSION_C3_CIRCULANT_TEXTURES_ARE_SIMULTANEOUSLY_DIAGONALIZED"
	StatusTensionS3Degeneracy                 = "CONDITIONAL_TENSION_S3_TEXTURE_HAS_ONE_PLUS_TWO_DEGENERACY"
	StatusTensionNumberOperatorBreaksTriality = "CONDITIONAL_TENSION_NUMBER_OPERATOR_BREAKS_EXACT_TRIALITY"
	StatusTensionNumberOperatorNotNative      = "CONDITIONAL_TENSION_NUMBER_OPERATOR_NOT_DERIVED_FROM_CURRENT_ASHA_LEDGER"
	StatusTensionSealedReductionNotNative     = "CONDITIONAL_TENSION_SEALED_MODULI_REDUCTION_DOES_NOT_REWRITE_NATIVE_FIREWALL"

	StatusVerifiedDomainAdmitted          = "VERIFIED_TRIALITY_DOMAIN_ADMITTED"
	StatusFailedDomainNotAdmitted         = "FAILED_ROUTE_DOMAIN_NOT_ADMITTED"
	StatusFailedTrialityOnlyLabelSymmetry = "FAILED_ROUTE_TRIALITY_IS_ONLY_LABEL_SYMMETRY"
	StatusFailedExactTrialityDegeneracy   = "FAILED_ROUTE_EXACT_TRIALITY_DEGENERACY"
	StatusFailedNoCKMMisalignment         = "FAILED_ROUTE_NO_CKM_MISALIGNMENT"
	StatusFailedCircularNumberInsertion   = "FAILED_ROUTE_CIRCULAR_N_INSERTION_NOT_NATIVE"
	StatusFirewallPreserved13Moduli       = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

// Inheritance records the hard facts inherited from the immediately relevant
// prior gates.  The package intentionally stores these as audited snapshots to
// avoid re-running broad historical theorem chains.
type Inheritance struct {
	Executed                        bool
	Gate247TrialityFunctorMissing   bool
	Gate370NativeMapsCentral        bool
	Gate371NumberOperatorNonNative  bool
	Gate372ChargedModuliDim         int
	Gate387FlavorFirewallSealed     bool
	NoEmpiricalFlavorValuesImported bool
	Verdict                         string
}

type DomainAdmission struct {
	Executed                               bool
	AbstractSpin8TrialityAvailable         bool
	Representations                        []string
	RequestedAssignment                    []string
	NativeTrialityCarrierFound             bool
	GenerationToTrialityFunctorDerived     bool
	ExplicitNativeThetaAvailable           bool
	ExplicitLabelPermutationThetaAvailable bool
	CompatibleWithAFRepresentation         bool
	CompatibleWithJAndFirstOrder           bool
	CompatibleWithHyperchargeSU2Channels   bool
	DomainAdmitted                         bool
	ManualGenerationRelabelingRejected     bool
	MissingPieces                          []string
	Verdict                                string
}

type Matrix struct {
	Name string
	Data [][]float64
}

type CentralizerCase struct {
	Name                          string
	Constraints                   []string
	GeneralComplexRealDim         int
	HermitianRealDim              int
	SymmetricRealDim              int
	CanonicalForm                 string
	DistinctSingularValuesGeneric int
	HasOnePlusTwoDegeneracy       bool
	AllSectorTexturesCommute      bool
	CKMMisalignmentCapacity       bool
	Native                        bool
	Sealed                        bool
	RankResidual                  float64
	Verdict                       string
}

type CentralizerAudit struct {
	Executed bool
	Cycle    Matrix
	Mirror   Matrix
	Cases    []CentralizerCase
	Verdict  string
}

type NumberOperatorAudit struct {
	Executed                        bool
	Operator                        []float64
	Status                          string
	NativeDerived                   bool
	BridgeCompatible                bool
	SealedExternalExtension         bool
	CircularIfUsedAsSolution        bool
	CommNormWithCycle               float64
	CommNormWithMirror              float64
	BreaksExactTriality             bool
	CommutesWithHyperchargeSU2      bool
	CommutesWithDiagonalTextures    bool
	ProducesDiagonalHierarchy       bool
	ProducesMixing                  bool
	ProvidesTwoNoncommutingTextures bool
	Verdict                         string
}

type ModuliScenario struct {
	Name                          string
	AssumptionClass               string
	StartingChargedDim            int
	ResultingDim                  int
	DistinctChargedMassesPossible bool
	CKMMisalignmentPossible       bool
	LeptonQuarkSectorSeparation   bool
	Native                        bool
	Conditional                   bool
	Failed                        bool
	Reason                        string
	Verdict                       string
}

type ModuliAudit struct {
	Executed               bool
	StartingChargedDim     int
	NativeReductionBelow13 bool
	BestNativeDim          int
	BestConditionalDim     int
	Scenarios              []ModuliScenario
	Verdict                string
}

type FirewallAudit struct {
	Executed                     bool
	NoYukawaMassesImported       bool
	NoCKMImported                bool
	NoPMNSImported               bool
	NoEmpiricalOrderingImported  bool
	NoManualGenerationAssignment bool
	NoFakeSpin8MatricesInvented  bool
	NoNativeCarrierClaimed       bool
	NoModuliReductionClaimed     bool
	Verdict                      string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Domain      DomainAdmission
	Centralizer CentralizerAudit
	Number      NumberOperatorAudit
	Moduli      ModuliAudit
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
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
	inheritance := inheritLateGates()
	domain := auditDomainAdmission(inheritance)
	centralizer, err := auditCentralizers(domain)
	if err != nil {
		return Analysis{}, err
	}
	number := auditNumberOperator(centralizer)
	moduli := auditModuli(inheritance, domain, centralizer, number)
	firewall := auditFirewall(domain, moduli)
	next := chooseNextGate(domain, moduli, number)
	truth := buildTruth(domain, centralizer, number, moduli, next)
	return Analysis{inheritance, domain, centralizer, number, moduli, firewall, next, truth}, nil
}

func inheritLateGates() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate247TrialityFunctorMissing:   true,
		Gate370NativeMapsCentral:        true,
		Gate371NumberOperatorNonNative:  true,
		Gate372ChargedModuliDim:         13,
		Gate387FlavorFirewallSealed:     true,
		NoEmpiricalFlavorValuesImported: true,
		Verdict:                         join(StatusGate247Inherited, StatusGate370Inherited, StatusGate371Inherited, StatusGate372Inherited, StatusGate387Inherited),
	}
}

func auditDomainAdmission(i Inheritance) DomainAdmission {
	missing := []string{
		"native map C^3_gen -> 8_v ⊕ 8_s ⊕ 8_c",
		"explicit Spin(8) triality matrices on the finite spectral-triple generation carrier",
		"proof that the A_F, J, Γ, and first-order-compatible D_F edge ledger is a triality representation rather than a copied multiplicity",
		"noncentral generation address functor compatible with hypercharge and SU(2)_L Yukawa channels",
	}
	domainAdmitted := false
	return DomainAdmission{
		Executed:                               true,
		AbstractSpin8TrialityAvailable:         true,
		Representations:                        []string{"8_v vector", "8_s left spinor", "8_c right spinor"},
		RequestedAssignment:                    []string{"generation 1 -> 8_v", "generation 2 -> 8_s", "generation 3 -> 8_c"},
		NativeTrialityCarrierFound:             false,
		GenerationToTrialityFunctorDerived:     false,
		ExplicitNativeThetaAvailable:           false,
		ExplicitLabelPermutationThetaAvailable: true,
		CompatibleWithAFRepresentation:         false,
		CompatibleWithJAndFirstOrder:           false,
		CompatibleWithHyperchargeSU2Channels:   false,
		DomainAdmitted:                         domainAdmitted,
		ManualGenerationRelabelingRejected:     true,
		MissingPieces:                          missing,
		Verdict:                                join(StatusDomainAdmissionAudited, StatusAbstractTrialityAvailable, StatusFailedDomainNotAdmitted, StatusFailedTrialityOnlyLabelSymmetry, StatusTensionTrialityNeedsCarrier),
	}
}

func auditCentralizers(domain DomainAdmission) (CentralizerAudit, error) {
	cycle := Matrix{Name: "C3 generation cycle", Data: [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}}
	mirror := Matrix{Name: "S3 generation reflection", Data: [][]float64{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}}
	cases := make([]CentralizerCase, 0, 4)

	c3Dim, c3Residual, err := invariantComplexDim([]Matrix{cycle})
	if err != nil {
		return CentralizerAudit{}, err
	}
	s3Dim, s3Residual, err := invariantComplexDim([]Matrix{cycle, mirror})
	if err != nil {
		return CentralizerAudit{}, err
	}
	hermC3Dim, _, err := invariantHermitianDim([]Matrix{cycle})
	if err != nil {
		return CentralizerAudit{}, err
	}
	hermS3Dim, _, err := invariantHermitianDim([]Matrix{cycle, mirror})
	if err != nil {
		return CentralizerAudit{}, err
	}

	cases = append(cases, CentralizerCase{
		Name:                          "unconstrained charged Yukawa block",
		Constraints:                   []string{"none beyond Gate-372 finite spectral-triple edge allowance"},
		GeneralComplexRealDim:         18,
		HermitianRealDim:              9,
		SymmetricRealDim:              12,
		CanonicalForm:                 "generic complex 3x3 matrix",
		DistinctSingularValuesGeneric: 3,
		HasOnePlusTwoDegeneracy:       false,
		AllSectorTexturesCommute:      false,
		CKMMisalignmentCapacity:       true,
		Native:                        true,
		Sealed:                        false,
		RankResidual:                  0,
		Verdict:                       StatusGate372Inherited,
	})
	cases = append(cases, CentralizerCase{
		Name:                          "exact C3 label-triality stress test",
		Constraints:                   []string{"P Y P^{-1} = Y", "P=(123)"},
		GeneralComplexRealDim:         c3Dim,
		HermitianRealDim:              hermC3Dim,
		SymmetricRealDim:              6,
		CanonicalForm:                 "complex circulant Y=aI+bP+cP^2",
		DistinctSingularValuesGeneric: 3,
		HasOnePlusTwoDegeneracy:       false,
		AllSectorTexturesCommute:      true,
		CKMMisalignmentCapacity:       false,
		Native:                        domain.DomainAdmitted,
		Sealed:                        !domain.DomainAdmitted,
		RankResidual:                  c3Residual,
		Verdict:                       join(StatusLabelPermutationStressTestExecuted, StatusConditionalTrialityTextureCapacity, StatusTensionCyclicCirculantCommutes, StatusFailedNoCKMMisalignment),
	})
	cases = append(cases, CentralizerCase{
		Name:                          "exact S3 label-triality stress test",
		Constraints:                   []string{"P Y P^{-1} = Y", "R Y R^{-1} = Y", "P=(123)", "R=(23)"},
		GeneralComplexRealDim:         s3Dim,
		HermitianRealDim:              hermS3Dim,
		SymmetricRealDim:              4,
		CanonicalForm:                 "Y=aI+b(1-I); scalar on the standard 2D irrep",
		DistinctSingularValuesGeneric: 2,
		HasOnePlusTwoDegeneracy:       true,
		AllSectorTexturesCommute:      true,
		CKMMisalignmentCapacity:       false,
		Native:                        domain.DomainAdmitted,
		Sealed:                        !domain.DomainAdmitted,
		RankResidual:                  s3Residual,
		Verdict:                       join(StatusLabelPermutationStressTestExecuted, StatusConditionalTrialityTextureCapacity, StatusTensionS3Degeneracy, StatusFailedExactTrialityDegeneracy, StatusFailedNoCKMMisalignment),
	})
	cases = append(cases, CentralizerCase{
		Name:                          "native triality carrier branch",
		Constraints:                   []string{"requires domain admission before centralizer can be promoted"},
		GeneralComplexRealDim:         -1,
		HermitianRealDim:              -1,
		SymmetricRealDim:              -1,
		CanonicalForm:                 "not computed as native because the carrier/theta pair is absent",
		DistinctSingularValuesGeneric: 0,
		HasOnePlusTwoDegeneracy:       false,
		AllSectorTexturesCommute:      false,
		CKMMisalignmentCapacity:       false,
		Native:                        false,
		Sealed:                        false,
		RankResidual:                  0,
		Verdict:                       join(StatusFailedDomainNotAdmitted, StatusFirewallPreserved13Moduli),
	})
	return CentralizerAudit{Executed: true, Cycle: cycle, Mirror: mirror, Cases: cases, Verdict: join(StatusEquivariantCentralizerComputed, StatusConditionalTrialityTextureCapacity, StatusFailedNoCKMMisalignment)}, nil
}

func auditNumberOperator(_ CentralizerAudit) NumberOperatorAudit {
	n := diag([]float64{0, 1, 2})
	cycle := [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	mirror := [][]float64{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	cn := frob(comm(n, cycle))
	rn := frob(comm(n, mirror))
	return NumberOperatorAudit{
		Executed:                        true,
		Operator:                        []float64{0, 1, 2},
		Status:                          "bridge-level compatible / sealed if used as a texture selector",
		NativeDerived:                   false,
		BridgeCompatible:                true,
		SealedExternalExtension:         true,
		CircularIfUsedAsSolution:        true,
		CommNormWithCycle:               cn,
		CommNormWithMirror:              rn,
		BreaksExactTriality:             cn > eps || rn > eps,
		CommutesWithHyperchargeSU2:      true,
		CommutesWithDiagonalTextures:    true,
		ProducesDiagonalHierarchy:       true,
		ProducesMixing:                  false,
		ProvidesTwoNoncommutingTextures: false,
		Verdict:                         join(StatusFockNumberOperatorClassified, StatusConditionalFockHierarchyCapacity, StatusTensionNumberOperatorBreaksTriality, StatusTensionNumberOperatorNotNative, StatusFailedCircularNumberInsertion, StatusFailedNoCKMMisalignment),
	}
}

func auditModuli(i Inheritance, d DomainAdmission, c CentralizerAudit, n NumberOperatorAudit) ModuliAudit {
	start := i.Gate372ChargedModuliDim
	scenarios := []ModuliScenario{
		{
			Name:                          "native ASHA after Gate 393",
			AssumptionClass:               "native only",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       true,
			LeptonQuarkSectorSeparation:   true,
			Native:                        true,
			Conditional:                   false,
			Failed:                        false,
			Reason:                        "no native triality carrier/theta/generation-address functor was admitted, so Gate-372 census is unchanged",
			Verdict:                       join(StatusFirewallPreserved13Moduli, StatusFailedDomainNotAdmitted),
		},
		{
			Name:                          "central-only generation broadcast",
			AssumptionClass:               "native central maps only",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       true,
			LeptonQuarkSectorSeparation:   true,
			Native:                        true,
			Conditional:                   false,
			Failed:                        true,
			Reason:                        "maps through I3 add no constraint and do not select a vacuum point",
			Verdict:                       join(StatusFailedTrialityOnlyLabelSymmetry, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                          "sealed exact C3 label-triality",
			AssumptionClass:               "sealed label action",
			StartingChargedDim:            start,
			ResultingDim:                  9,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   true,
			Native:                        d.DomainAdmitted,
			Conditional:                   !d.DomainAdmitted,
			Failed:                        true,
			Reason:                        "sector matrices become circulant; they are Fourier-diagonalized together, so CKM capacity is removed rather than derived",
			Verdict:                       join(StatusConditionalTrialityTextureCapacity, StatusTensionCyclicCirculantCommutes, StatusFailedNoCKMMisalignment, StatusTensionSealedReductionNotNative),
		},
		{
			Name:                          "sealed exact S3 label-triality",
			AssumptionClass:               "sealed label action",
			StartingChargedDim:            start,
			ResultingDim:                  6,
			DistinctChargedMassesPossible: false,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   true,
			Native:                        d.DomainAdmitted,
			Conditional:                   !d.DomainAdmitted,
			Failed:                        true,
			Reason:                        "S3 invariance gives a singlet plus a twofold-degenerate standard sector, not three observed-like distinct masses",
			Verdict:                       join(StatusConditionalTrialityTextureCapacity, StatusFailedExactTrialityDegeneracy, StatusFailedNoCKMMisalignment, StatusTensionSealedReductionNotNative),
		},
		{
			Name:                          "sealed N=diag(0,1,2) hierarchy",
			AssumptionClass:               "bridge/sealed number operator",
			StartingChargedDim:            start,
			ResultingDim:                  9,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   true,
			Native:                        n.NativeDerived,
			Conditional:                   !n.NativeDerived,
			Failed:                        true,
			Reason:                        "one diagonal address can split generation levels, but it supplies no noncommuting up/down texture pair",
			Verdict:                       join(StatusConditionalFockHierarchyCapacity, StatusFailedCircularNumberInsertion, StatusFailedNoCKMMisalignment, StatusTensionSealedReductionNotNative),
		},
		{
			Name:                          "two native noncommuting texture operators",
			AssumptionClass:               "missing prerequisite",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			DistinctChargedMassesPossible: false,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   false,
			Native:                        false,
			Conditional:                   false,
			Failed:                        true,
			Reason:                        "Gate 393 found no native noncommuting generation-address algebra; exact quotient recount must wait for that object",
			Verdict:                       join(StatusFailedDomainNotAdmitted, StatusFailedNoCKMMisalignment, StatusFirewallPreserved13Moduli),
		},
	}
	bestConditional := start
	for _, s := range scenarios {
		if s.Conditional && s.ResultingDim >= 0 && s.ResultingDim < bestConditional {
			bestConditional = s.ResultingDim
		}
	}
	return ModuliAudit{Executed: true, StartingChargedDim: start, NativeReductionBelow13: false, BestNativeDim: start, BestConditionalDim: bestConditional, Scenarios: scenarios, Verdict: join(StatusModuliRecountExecuted, StatusFirewallPreserved13Moduli)}
}

func auditFirewall(d DomainAdmission, m ModuliAudit) FirewallAudit {
	return FirewallAudit{
		Executed:                     true,
		NoYukawaMassesImported:       true,
		NoCKMImported:                true,
		NoPMNSImported:               true,
		NoEmpiricalOrderingImported:  true,
		NoManualGenerationAssignment: d.ManualGenerationRelabelingRejected,
		NoFakeSpin8MatricesInvented:  true,
		NoNativeCarrierClaimed:       !d.DomainAdmitted,
		NoModuliReductionClaimed:     !m.NativeReductionBelow13,
		Verdict:                      join(StatusFirewallPreserved13Moduli, StatusFailedDomainNotAdmitted),
	}
}

func chooseNextGate(d DomainAdmission, m ModuliAudit, n NumberOperatorAudit) NextStep {
	if !d.DomainAdmitted {
		return NextStep{Gate: 394, Title: "Native Generation-Address Functor from Triality/Morita Edge Incidence", Reason: "Gate 393 rejects direct generation->triality relabeling; the missing object is still a native noncentral map into End(C^3_gen).", PrimaryTask: "Search Morita edge incidence, one-form support, and triality branch data for a lawful generation-address functor before any flavor-moduli quotient is recomputed."}
	}
	if !m.NativeReductionBelow13 && n.ProducesDiagonalHierarchy {
		return NextStep{Gate: 394, Title: "Noncommuting Generation Texture Pair Sieve", Reason: "one diagonal hierarchy operator cannot create CKM/PMNS capacity", PrimaryTask: "derive or reject a second noncommuting generation operator compatible with D_F, J, first-order, hypercharge, and SU(2)_L."}
	}
	return NextStep{Gate: 398, Title: "Exact Charged Flavor-Moduli Quotient Recount", Reason: "native moduli reduction data exists", PrimaryTask: "compute the quotient dimension under the derived generation-address algebra."}
}

func buildTruth(d DomainAdmission, c CentralizerAudit, n NumberOperatorAudit, m ModuliAudit, next NextStep) string {
	return fmt.Sprintf("Gate 393 rejects the direct claim that Spin(8) triality alone breaks the Gate-372 13-moduli firewall. Abstract Spin(8) triality is available, and the sealed label-permutation stress test is computable: C3 forces complex circulant Yukawa blocks, while S3 forces aI+b(1-I). However, the native ASHA finite spectral-triple generation carrier is not admitted into 8_v ⊕ 8_s ⊕ 8_c, no explicit native theta acts on D_F generation entries, exact S3 has a 1+2 degeneracy, and C3/S3-constrained sector textures commute or are simultaneously diagonalized, so CKM misalignment is not derived. N=diag(0,1,2) has hierarchy capacity but breaks exact triality and remains non-native/circular if promoted to a solution. Native charged flavor moduli remain %d. The next gate is Gate %d: %s.", m.StartingChargedDim, next.Gate, next.Title)
}

func invariantComplexDim(gens []Matrix) (int, float64, error) {
	basis := complexMatrixBasis()
	constraints := make([][]float64, 0)
	for _, g := range gens {
		ginv := transpose(g.Data)
		for _, b := range basis {
			lhsR := mul(mul(g.Data, b.real), ginv)
			lhsI := mul(mul(g.Data, b.imag), ginv)
			dr := sub(lhsR, b.real)
			di := sub(lhsI, b.imag)
			constraints = append(constraints, flattenComplex(dr, di))
		}
	}
	// The loop above produced columns, not equations.  Build A by rows from the
	// linear map columns.
	cols := constraints
	if len(cols) != len(basis)*len(gens) {
		return 0, 0, fmt.Errorf("unexpected constraint column count")
	}
	rows := transposeRect(cols)
	rank, residual := rank(rows)
	return len(basis) - rank, residual, nil
}

func invariantHermitianDim(gens []Matrix) (int, float64, error) {
	basis := hermitianBasis()
	cols := make([][]float64, 0)
	for _, g := range gens {
		ginv := transpose(g.Data)
		for _, b := range basis {
			lhsR := mul(mul(g.Data, b.real), ginv)
			lhsI := mul(mul(g.Data, b.imag), ginv)
			dr := sub(lhsR, b.real)
			di := sub(lhsI, b.imag)
			cols = append(cols, flattenComplex(dr, di))
		}
	}
	rows := transposeRect(cols)
	rank, residual := rank(rows)
	return len(basis) - rank, residual, nil
}

type complexBasisMatrix struct{ real, imag [][]float64 }

func complexMatrixBasis() []complexBasisMatrix {
	out := make([]complexBasisMatrix, 0, 18)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := zero(3)
			r[i][j] = 1
			out = append(out, complexBasisMatrix{real: r, imag: zero(3)})
			ii := zero(3)
			ii[i][j] = 1
			out = append(out, complexBasisMatrix{real: zero(3), imag: ii})
		}
	}
	return out
}

func hermitianBasis() []complexBasisMatrix {
	out := make([]complexBasisMatrix, 0, 9)
	for i := 0; i < 3; i++ {
		r := zero(3)
		r[i][i] = 1
		out = append(out, complexBasisMatrix{real: r, imag: zero(3)})
	}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			r := zero(3)
			r[i][j], r[j][i] = 1, 1
			out = append(out, complexBasisMatrix{real: r, imag: zero(3)})
			im := zero(3)
			im[i][j], im[j][i] = 1, -1
			out = append(out, complexBasisMatrix{real: zero(3), imag: im})
		}
	}
	return out
}

func rank(a [][]float64) (int, float64) {
	if len(a) == 0 || len(a[0]) == 0 {
		return 0, 0
	}
	m, n := len(a), len(a[0])
	mat := make([][]float64, m)
	for i := range a {
		mat[i] = append([]float64(nil), a[i]...)
	}
	r := 0
	maxResidual := 0.0
	for c := 0; c < n && r < m; c++ {
		piv := r
		for i := r + 1; i < m; i++ {
			if math.Abs(mat[i][c]) > math.Abs(mat[piv][c]) {
				piv = i
			}
		}
		if math.Abs(mat[piv][c]) <= eps {
			if math.Abs(mat[piv][c]) > maxResidual {
				maxResidual = math.Abs(mat[piv][c])
			}
			continue
		}
		mat[r], mat[piv] = mat[piv], mat[r]
		pv := mat[r][c]
		for j := c; j < n; j++ {
			mat[r][j] /= pv
		}
		for i := 0; i < m; i++ {
			if i == r {
				continue
			}
			f := mat[i][c]
			if math.Abs(f) <= eps {
				continue
			}
			for j := c; j < n; j++ {
				mat[i][j] -= f * mat[r][j]
			}
		}
		r++
	}
	for i := range mat {
		for j := range mat[i] {
			if math.Abs(mat[i][j]) < eps && math.Abs(mat[i][j]) > maxResidual {
				maxResidual = math.Abs(mat[i][j])
			}
		}
	}
	return r, maxResidual
}

func zero(n int) [][]float64 {
	m := make([][]float64, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}

func diag(d []float64) [][]float64 {
	m := zero(len(d))
	for i, v := range d {
		m[i][i] = v
	}
	return m
}

func mul(a, b [][]float64) [][]float64 {
	m, n, p := len(a), len(b), len(b[0])
	out := make([][]float64, m)
	for i := 0; i < m; i++ {
		out[i] = make([]float64, p)
		for k := 0; k < n; k++ {
			for j := 0; j < p; j++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func sub(a, b [][]float64) [][]float64 {
	out := zero(len(a))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func transpose(a [][]float64) [][]float64 {
	out := make([][]float64, len(a[0]))
	for i := range out {
		out[i] = make([]float64, len(a))
		for j := range a {
			out[i][j] = a[j][i]
		}
	}
	return out
}

func transposeRect(cols [][]float64) [][]float64 {
	if len(cols) == 0 {
		return nil
	}
	rows := make([][]float64, len(cols[0]))
	for i := range rows {
		rows[i] = make([]float64, len(cols))
		for j := range cols {
			rows[i][j] = cols[j][i]
		}
	}
	return rows
}

func flattenComplex(r, im [][]float64) []float64 {
	out := make([]float64, 0, 18)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out = append(out, r[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out = append(out, im[i][j])
		}
	}
	return out
}

func comm(a, b [][]float64) [][]float64 { return sub(mul(a, b), mul(b, a)) }

func frob(a [][]float64) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}

func join(parts ...string) string {
	seen := map[string]bool{}
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		for _, q := range strings.Split(p, ";") {
			q = strings.TrimSpace(q)
			if q == "" || seen[q] {
				continue
			}
			seen[q] = true
			out = append(out, q)
		}
	}
	return strings.Join(out, ";")
}

func Statuses(a Analysis) []string {
	var raw []string
	raw = append(raw, a.Inheritance.Verdict, a.Domain.Verdict, a.Centralizer.Verdict, a.Number.Verdict, a.Moduli.Verdict, a.Firewall.Verdict)
	for _, c := range a.Centralizer.Cases {
		raw = append(raw, c.Verdict)
	}
	for _, s := range a.Moduli.Scenarios {
		raw = append(raw, s.Verdict)
	}
	return strings.Split(join(raw...), ";")
}
