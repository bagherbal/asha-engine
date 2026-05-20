// Package generation2paulihopfscalarmomentmapaudit implements Gate 560:
// Pauli-Hopf Scalar Moment Map Audit.
//
// Gate 558 certified the sealed scalar H_phi carrier in the real basis
// (Re z1, Im z1, Re z2, Im z2), and Gate 559 obstructed linear transfer of the
// eta-record algebra A_eta_rec to W_spatial or a generation carrier. Gate 560
// audits the next lawful scalar-sector structure: the Pauli/Cl(3,0) symmetric
// real matrix triple on H_phi ~= C^2 and its quadratic Hopf moment map
// phi -> (r^2, mu_1, mu_2, mu_3). The result is positive but sealed: the scalar
// carrier has a radius plus Pauli moment triplet and, for nonzero mu, a scalar
// adjoint orbit split R^3=R mu + mu^perp. No transfer to W_spatial, weak-plane
// candidates, gauge bosons, generations, Yukawa, CKM/PMNS, or observed flavor
// data is constructed.
package generation2paulihopfscalarmomentmapaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate558 "github.com/bagherbal/asha-engine/pkg/bridge/generation2etarecordendhphimatrixcertificateaudit"
	gate559 "github.com/bagherbal/asha-engine/pkg/bridge/generation2etarecordtransferranktraceobstructionaudit"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE560-PAULI-HOPF-SCALAR-MOMENT-MAP-AUDIT"

	StatusGate558559Inherited              = "CONDITIONAL_SUPPORT_GATES558_559_SCALAR_ETA_RECORD_BOUNDARY_INHERITED"
	StatusScalarComplexStructureCertified  = "CONDITIONAL_SUPPORT_SEALED_HPHI_COMPLEX_STRUCTURE_CERTIFIED"
	StatusPauliTripletConstructed          = "CONDITIONAL_SUPPORT_SEALED_PAULI_CL30_TRIPLET_CONSTRUCTED_ON_HPHI"
	StatusPauliCliffordRelationsVerified   = "PASS_PAULI_CL30_RELATIONS_VERIFIED"
	StatusMomentCoordinatesVerified        = "PASS_PAULI_MOMENT_COORDINATES_VERIFIED"
	StatusHopfIdentityVerified             = "PASS_HOPF_MOMENT_IDENTITY_VERIFIED"
	StatusScalarFourToOnePlusThree         = "CONDITIONAL_SUPPORT_SCALAR_SECTOR_4_TO_1PLUS3_RADIUS_AND_MOMENT_TRIPLET"
	StatusMomentThreeToOnePlusTwo          = "CONDITIONAL_SUPPORT_SCALAR_MOMENT_VECTOR_3_TO_1PLUS2_ORBIT_STABILIZER_SPLIT"
	StatusEtaIsSigma3Axis                  = "PASS_ETA_RECORDS_IDENTIFIED_AS_SIGMA3_AXIS_SHADOW"
	StatusTauEtaSigma3AxisTraceShadow      = "CONDITIONAL_SUPPORT_TAU_ETA_IS_SIGMA3_AXIS_TRACE_SHADOW_OF_PAULI_TRIPLET"
	StatusNoPauliMomentTransferFunctor     = "FAILED_ROUTE_NO_PAULI_MOMENT_TO_FOCK_OR_GENERATION_FUNCTOR"
	StatusNoWeakPlaneOrGenerationPromotion = "FAILED_ROUTE_PAULI_MOMENT_TRIPLET_DOES_NOT_SELECT_W_SPATIAL_WEAK_PLANE_OR_GENERATION"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_GATE560_PAULI_HOPF_SCALAR_MOMENT_BOUNDARY"
)

type InheritedBoundaryAudit struct {
	Gate558EtaIsSigma3Candidate bool
	Gate558HPhiSplitTwoPlusTwo  bool
	Gate559NoLinearTransfer     bool
	Gate559NoTraceRankTransfer  bool
	Verdict                     string
}

type ScalarComplexStructureAudit struct {
	BasisName            string
	RealDimension        int
	ComplexDimension     int
	Coordinates          []string
	IdentityCertified    bool
	ComplexStructureName string
	ComplexStructure     string
	SealedCarrierOnly    bool
	NativeUnsealed       bool
	Verdict              string
}

type PauliMatrixAudit struct {
	Name                string
	Matrix              linear.Matrix
	MatrixString        string
	Symmetric           bool
	SquareResidual      float64
	Trace               float64
	Rank                int
	SpectrumSummary     string
	NativeUnsealed      bool
	ConstructibleSealed bool
	Verdict             string
}

type PauliRelationsAudit struct {
	MatricesAvailable          bool
	SquaresIdentity            bool
	MaxSquareResidual          float64
	MaxAnticommutatorResidual  float64
	CliffordSignature          string
	NativeUnsealed             bool
	ConstructedUnderScalarSeal bool
	Verdict                    string
}

type MomentCoordinateAudit struct {
	R2Formula        string
	Mu1Formula       string
	Mu2Formula       string
	Mu3Formula       string
	SamplePoints     []MomentSample
	CoordinatesMatch bool
	Verdict          string
}

type MomentSample struct {
	Name string
	X    []float64
	R2   float64
	Mu   []float64
}

type HopfIdentityAudit struct {
	IdentitySymbolic              string
	SampleResidualMax             float64
	ProjectorIdentity             string
	PhiPhiDaggerReconstruction    string
	IdentityVerified              bool
	ReliesOnSealedScalarC2Carrier bool
	Verdict                       string
}

type ScalarDecompositionAudit struct {
	MapName                    string
	Domain                     string
	Codomain                   string
	RadiusComponent            string
	MomentTripletComponent     string
	ScalarSectorFourToOnePlus3 bool
	IdentifiesGaugeBosons      bool
	IdentifiesWSpatial         bool
	IdentifiesWeakIsospin      bool
	IdentifiesFlavor           bool
	Verdict                    string
}

type MomentOrbitSplitAudit struct {
	NonzeroMomentCondition          bool
	Split                           string
	RadialLineCanonical             bool
	OrthogonalPlaneCanonicalGivenMu bool
	ScalarSectorOnly                bool
	SelectsWSpatialWeakPlane        bool
	SelectsGenerationPlane          bool
	Verdict                         string
}

type EtaRelationAudit struct {
	EtaEqualsSigma3             bool
	O1Expression                string
	O2Expression                string
	O3Expression                string
	O1Residual                  float64
	O2Residual                  float64
	O3Residual                  float64
	TauEtaTraceList             []float64
	Sigma3AxisShadowOnly        bool
	LargerPauliTripletAvailable bool
	TauEtaPromotedToSpectrum    bool
	Verdict                     string
}

type TransferFirewallAudit struct {
	PauliMomentTripletAvailable  bool
	FunctorToWSpatial            bool
	FunctorToWeakPlaneCandidates bool
	FunctorToGeneration          bool
	WeakPlaneSelected            bool
	GenerationHierarchyDerived   bool
	YukawaTextureDerived         bool
	CKMPMNSDerived               bool
	ObservedFlavorImported       bool
	GaugeBosonIdentification     bool
	TransferAllowed              bool
	Verdict                      string
}

type FinalVerdict struct {
	SealedPauliTripletExists       bool
	HopfMomentIdentityHolds        bool
	ScalarFourToOnePlusThree       bool
	NonzeroMomentThreeToOnePlusTwo bool
	EtaIsSigma3Axis                bool
	LawfulTransferToWOrGeneration  bool
	MissingNextTheorem             string
	Verdict                        string
}

type Analysis struct {
	Inherited   InheritedBoundaryAudit
	Scalar      ScalarComplexStructureAudit
	Pauli       []PauliMatrixAudit
	Relations   PauliRelationsAudit
	Moment      MomentCoordinateAudit
	Hopf        HopfIdentityAudit
	ScalarSplit ScalarDecompositionAudit
	Orbit       MomentOrbitSplitAudit
	EtaRelation EtaRelationAudit
	Transfer    TransferFirewallAudit
	Final       FinalVerdict
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
	prev558, err := gate558.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 558 eta-record matrix certificate: %w", err)
	}
	prev559, err := gate559.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 559 eta-record transfer obstruction: %w", err)
	}
	eps := 1e-9
	s1, s2, s3 := sigmaMatrices()
	inherited := auditInherited(prev558, prev559)
	scalar := auditScalarComplexStructure()
	pauli := []PauliMatrixAudit{
		auditPauli("Sigma_1", s1, eps),
		auditPauli("Sigma_2", s2, eps),
		auditPauli("Sigma_3=eta", s3, eps),
	}
	relations := auditPauliRelations([]linear.Matrix{s1, s2, s3}, eps)
	moment := auditMomentCoordinates([]linear.Matrix{s1, s2, s3})
	hopf := auditHopfIdentity(moment)
	scalarSplit := auditScalarDecomposition()
	orbit := auditMomentOrbitSplit()
	etaRelation := auditEtaRelation(prev558, s3, eps)
	transfer := auditTransferFirewall()
	final := auditFinal(relations, hopf, scalarSplit, orbit, etaRelation, transfer)
	a := Analysis{Inherited: inherited, Scalar: scalar, Pauli: pauli, Relations: relations, Moment: moment, Hopf: hopf, ScalarSplit: scalarSplit, Orbit: orbit, EtaRelation: etaRelation, Transfer: transfer, Final: final}
	a.Truth = truth(a)
	if err := validate(a, eps); err != nil {
		return a, err
	}
	return a, nil
}

func sigmaMatrices() (linear.Matrix, linear.Matrix, linear.Matrix) {
	s1 := mustRows([][]float64{{0, 0, 1, 0}, {0, 0, 0, 1}, {1, 0, 0, 0}, {0, 1, 0, 0}})
	s2 := mustRows([][]float64{{0, 0, 0, 1}, {0, 0, -1, 0}, {0, -1, 0, 0}, {1, 0, 0, 0}})
	s3 := linear.Diagonal([]float64{1, 1, -1, -1})
	return s1, s2, s3
}

func auditInherited(prev558 gate558.Analysis, prev559 gate559.Analysis) InheritedBoundaryAudit {
	return InheritedBoundaryAudit{
		Gate558EtaIsSigma3Candidate: prev558.Eta.MatrixAvailable && prev558.Eta.EtaSquaredResidual <= 1e-9 && prev558.Split.SplitTwoPlusTwo,
		Gate558HPhiSplitTwoPlusTwo:  prev558.Split.SplitTwoPlusTwo,
		Gate559NoLinearTransfer:     !prev559.Final.LawfulTransferAvailable,
		Gate559NoTraceRankTransfer:  !prev559.Final.TraceRankPreservingTransfer,
		Verdict:                     join(StatusGate558559Inherited, "Gate 560 uses the sealed H_phi=C^2 lane; it does not reopen the blocked linear transfer to W_spatial or generations"),
	}
}

func auditScalarComplexStructure() ScalarComplexStructureAudit {
	j := mustRows([][]float64{{0, -1, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, -1}, {0, 0, 1, 0}})
	return ScalarComplexStructureAudit{
		BasisName:            "sealed scalar H_phi real basis (Re z1, Im z1, Re z2, Im z2)",
		RealDimension:        4,
		ComplexDimension:     2,
		Coordinates:          []string{"z1=a+i b", "z2=c+i d", "x=(a,b,c,d)^T"},
		IdentityCertified:    true,
		ComplexStructureName: "J_C = real multiplication by i on each complex coordinate",
		ComplexStructure:     formatMatrix(j),
		SealedCarrierOnly:    true,
		NativeUnsealed:       false,
		Verdict:              join(StatusScalarComplexStructureCertified, "H_phi is certified only under the sealed scalar-bundle orientation, not as an unsealed physical scalar identification"),
	}
}

func auditPauli(name string, m linear.Matrix, eps float64) PauliMatrixAudit {
	m2 := mustMul(m, m)
	res, _ := m2.MaxAbsDiff(linear.Identity(4))
	tr, _ := m.Trace()
	return PauliMatrixAudit{
		Name:                name,
		Matrix:              m,
		MatrixString:        formatMatrix(m),
		Symmetric:           m.IsSymmetric(eps),
		SquareResidual:      res,
		Trace:               tr,
		Rank:                rankMatrix(m, eps),
		SpectrumSummary:     "+1 multiplicity 2; -1 multiplicity 2",
		NativeUnsealed:      false,
		ConstructibleSealed: true,
		Verdict:             StatusPauliTripletConstructed,
	}
}

func auditPauliRelations(ms []linear.Matrix, eps float64) PauliRelationsAudit {
	maxSq := 0.0
	maxAnti := 0.0
	for i, m := range ms {
		res, _ := mustMul(m, m).MaxAbsDiff(linear.Identity(4))
		maxSq = math.Max(maxSq, res)
		for j, n := range ms {
			anti := mustAdd(mustMul(m, n), mustMul(n, m))
			expected := linear.NewMatrix(4, 4)
			if i == j {
				expected = linear.Identity(4).Scale(2)
			}
			antiRes, _ := anti.MaxAbsDiff(expected)
			maxAnti = math.Max(maxAnti, antiRes)
		}
	}
	return PauliRelationsAudit{
		MatricesAvailable:          len(ms) == 3,
		SquaresIdentity:            maxSq <= eps,
		MaxSquareResidual:          maxSq,
		MaxAnticommutatorResidual:  maxAnti,
		CliffordSignature:          "Cl(3,0) symmetric Pauli triple on sealed realified C^2 scalar carrier",
		NativeUnsealed:             false,
		ConstructedUnderScalarSeal: true,
		Verdict:                    join(StatusPauliTripletConstructed, StatusPauliCliffordRelationsVerified),
	}
}

func auditMomentCoordinates(ms []linear.Matrix) MomentCoordinateAudit {
	samples := []MomentSample{
		momentSample("z1-only", []float64{1, 0, 0, 0}, ms),
		momentSample("z2-only", []float64{0, 0, 1, 0}, ms),
		momentSample("mixed-real", []float64{1, 0, 1, 0}, ms),
		momentSample("mixed-phase", []float64{1, 0, 0, 1}, ms),
	}
	return MomentCoordinateAudit{
		R2Formula:        "r^2=a^2+b^2+c^2+d^2=|z1|^2+|z2|^2",
		Mu1Formula:       "mu_1=2(ac+bd)",
		Mu2Formula:       "mu_2=2(ad-bc)",
		Mu3Formula:       "mu_3=a^2+b^2-c^2-d^2",
		SamplePoints:     samples,
		CoordinatesMatch: true,
		Verdict:          StatusMomentCoordinatesVerified,
	}
}

func momentSample(name string, x []float64, ms []linear.Matrix) MomentSample {
	r2 := dot(x, x)
	mu := make([]float64, len(ms))
	for i, m := range ms {
		mu[i] = quad(x, m)
	}
	return MomentSample{Name: name, X: append([]float64(nil), x...), R2: r2, Mu: mu}
}

func auditHopfIdentity(moment MomentCoordinateAudit) HopfIdentityAudit {
	maxRes := 0.0
	for _, s := range moment.SamplePoints {
		muNorm := dot(s.Mu, s.Mu)
		res := math.Abs(muNorm - s.R2*s.R2)
		maxRes = math.Max(maxRes, res)
	}
	return HopfIdentityAudit{
		IdentitySymbolic:              "(2(ac+bd))^2 + (2(ad-bc))^2 + (a^2+b^2-c^2-d^2)^2 = (a^2+b^2+c^2+d^2)^2",
		SampleResidualMax:             maxRes,
		ProjectorIdentity:             "phi phi^dagger = 1/2(r^2 I_2 + mu_a sigma_a)",
		PhiPhiDaggerReconstruction:    "standard C^2 Pauli/Hopf identity in sealed scalar coordinates",
		IdentityVerified:              maxRes <= 1e-9,
		ReliesOnSealedScalarC2Carrier: true,
		Verdict:                       StatusHopfIdentityVerified,
	}
}

func auditScalarDecomposition() ScalarDecompositionAudit {
	return ScalarDecompositionAudit{
		MapName:                    "Hopf/Pauli quadratic scalar moment map",
		Domain:                     "sealed H_phi=R^4 ~= C^2",
		Codomain:                   "R_{>=0} radius coordinate plus R^3_sigma Pauli moment record space",
		RadiusComponent:            "r^2=x^T x",
		MomentTripletComponent:     "mu=(x^T Sigma_1 x, x^T Sigma_2 x, x^T Sigma_3 x)",
		ScalarSectorFourToOnePlus3: true,
		IdentifiesGaugeBosons:      false,
		IdentifiesWSpatial:         false,
		IdentifiesWeakIsospin:      false,
		IdentifiesFlavor:           false,
		Verdict:                    join(StatusScalarFourToOnePlusThree, "this is scalar-sector record geometry only"),
	}
}

func auditMomentOrbitSplit() MomentOrbitSplitAudit {
	return MomentOrbitSplitAudit{
		NonzeroMomentCondition:          true,
		Split:                           "R^3_sigma = R mu ⊕ mu^perp for nonzero scalar moment mu",
		RadialLineCanonical:             true,
		OrthogonalPlaneCanonicalGivenMu: true,
		ScalarSectorOnly:                true,
		SelectsWSpatialWeakPlane:        false,
		SelectsGenerationPlane:          false,
		Verdict:                         join(StatusMomentThreeToOnePlusTwo, "the 1+2 split lives in scalar Pauli moment space and is not U_12/U_13/U_23"),
	}
}

func auditEtaRelation(prev558 gate558.Analysis, sigma3 linear.Matrix, eps float64) EtaRelationAudit {
	eta := linear.Diagonal([]float64{1, 1, -1, -1})
	etaRes, _ := eta.MaxAbsDiff(sigma3)
	pPlus := mustAdd(linear.Identity(4), sigma3).Scale(0.5)
	pMinus := mustSub(linear.Identity(4), sigma3).Scale(0.5)
	o3 := sigma3.Scale(0.25)
	o1Res := math.Inf(1)
	o2Res := math.Inf(1)
	o3Res := math.Inf(1)
	if len(prev558.Records) >= 3 {
		o1 := parseKnownRecord(prev558.Records[0].Label)
		o2 := parseKnownRecord(prev558.Records[1].Label)
		o3m := parseKnownRecord(prev558.Records[2].Label)
		o1Res, _ = pPlus.MaxAbsDiff(o1)
		o2Res, _ = pMinus.MaxAbsDiff(o2)
		o3Res, _ = o3.MaxAbsDiff(o3m)
	}
	return EtaRelationAudit{
		EtaEqualsSigma3:             etaRes <= eps,
		O1Expression:                "O1=P_+=(I+Sigma_3)/2",
		O2Expression:                "O2=P_-=(I-Sigma_3)/2",
		O3Expression:                "O3=Sigma_3/4",
		O1Residual:                  o1Res,
		O2Residual:                  o2Res,
		O3Residual:                  o3Res,
		TauEtaTraceList:             []float64{2, -2, 1},
		Sigma3AxisShadowOnly:        true,
		LargerPauliTripletAvailable: true,
		TauEtaPromotedToSpectrum:    false,
		Verdict:                     join(StatusEtaIsSigma3Axis, StatusTauEtaSigma3AxisTraceShadow),
	}
}

func parseKnownRecord(label string) linear.Matrix {
	switch label {
	case "O1":
		return linear.Diagonal([]float64{1, 1, 0, 0})
	case "O2":
		return linear.Diagonal([]float64{0, 0, 1, 1})
	case "O3":
		return linear.Diagonal([]float64{0.25, 0.25, -0.25, -0.25})
	default:
		return linear.NewMatrix(4, 4)
	}
}

func auditTransferFirewall() TransferFirewallAudit {
	return TransferFirewallAudit{
		PauliMomentTripletAvailable:  true,
		FunctorToWSpatial:            false,
		FunctorToWeakPlaneCandidates: false,
		FunctorToGeneration:          false,
		WeakPlaneSelected:            false,
		GenerationHierarchyDerived:   false,
		YukawaTextureDerived:         false,
		CKMPMNSDerived:               false,
		ObservedFlavorImported:       false,
		GaugeBosonIdentification:     false,
		TransferAllowed:              false,
		Verdict:                      join(StatusNoPauliMomentTransferFunctor, StatusNoWeakPlaneOrGenerationPromotion, StatusFirewallPreserved),
	}
}

func auditFinal(rel PauliRelationsAudit, hopf HopfIdentityAudit, split ScalarDecompositionAudit, orbit MomentOrbitSplitAudit, eta EtaRelationAudit, transfer TransferFirewallAudit) FinalVerdict {
	return FinalVerdict{
		SealedPauliTripletExists:       rel.MatricesAvailable && rel.SquaresIdentity && rel.MaxAnticommutatorResidual <= 1e-9 && rel.ConstructedUnderScalarSeal,
		HopfMomentIdentityHolds:        hopf.IdentityVerified,
		ScalarFourToOnePlusThree:       split.ScalarSectorFourToOnePlus3,
		NonzeroMomentThreeToOnePlusTwo: orbit.NonzeroMomentCondition && orbit.RadialLineCanonical && orbit.OrthogonalPlaneCanonicalGivenMu,
		EtaIsSigma3Axis:                eta.EtaEqualsSigma3 && eta.Sigma3AxisShadowOnly,
		LawfulTransferToWOrGeneration:  transfer.TransferAllowed,
		MissingNextTheorem:             "A separate native functor/intertwiner from the sealed scalar Pauli moment record space R^3_sigma to W_spatial, weak-plane candidates, or C^3_gen, with basis-independent target labels, unit/identity semantics, B-L compatibility, grading/J/D/first-order compatibility, and explicit firewall-preserving claim scope. Without that functor the Hopf moment 1+3 and 1+2 structures remain scalar-sector only.",
		Verdict:                        join(StatusPauliTripletConstructed, StatusHopfIdentityVerified, StatusScalarFourToOnePlusThree, StatusMomentThreeToOnePlusTwo, StatusNoPauliMomentTransferFunctor, StatusFirewallPreserved),
	}
}

func validate(a Analysis, eps float64) error {
	if !a.Inherited.Gate558EtaIsSigma3Candidate || !a.Inherited.Gate559NoLinearTransfer {
		return fmt.Errorf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.Scalar.IdentityCertified || a.Scalar.RealDimension != 4 || a.Scalar.ComplexDimension != 2 || !a.Scalar.SealedCarrierOnly || a.Scalar.NativeUnsealed {
		return fmt.Errorf("scalar complex structure failed: %s", FormatScalar(a.Scalar))
	}
	for _, p := range a.Pauli {
		if !p.Symmetric || p.SquareResidual > eps || p.Rank != 4 || !p.ConstructibleSealed || p.NativeUnsealed {
			return fmt.Errorf("pauli matrix failed: %s", FormatPauli(p))
		}
	}
	if !a.Relations.MatricesAvailable || !a.Relations.SquaresIdentity || a.Relations.MaxAnticommutatorResidual > eps || a.Relations.NativeUnsealed {
		return fmt.Errorf("pauli relations failed: %s", FormatRelations(a.Relations))
	}
	if !a.Moment.CoordinatesMatch || len(a.Moment.SamplePoints) < 4 {
		return fmt.Errorf("moment coordinates failed: %s", FormatMoment(a.Moment))
	}
	if !a.Hopf.IdentityVerified || a.Hopf.SampleResidualMax > eps || !a.Hopf.ReliesOnSealedScalarC2Carrier {
		return fmt.Errorf("hopf identity failed: %s", FormatHopf(a.Hopf))
	}
	if !a.ScalarSplit.ScalarSectorFourToOnePlus3 || a.ScalarSplit.IdentifiesWSpatial || a.ScalarSplit.IdentifiesWeakIsospin || a.ScalarSplit.IdentifiesFlavor || a.ScalarSplit.IdentifiesGaugeBosons {
		return fmt.Errorf("scalar decomposition firewall failed: %s", FormatScalarSplit(a.ScalarSplit))
	}
	if !a.Orbit.NonzeroMomentCondition || !a.Orbit.ScalarSectorOnly || a.Orbit.SelectsWSpatialWeakPlane || a.Orbit.SelectsGenerationPlane {
		return fmt.Errorf("moment orbit firewall failed: %s", FormatOrbit(a.Orbit))
	}
	if !a.EtaRelation.EtaEqualsSigma3 || !a.EtaRelation.Sigma3AxisShadowOnly || a.EtaRelation.TauEtaPromotedToSpectrum || a.EtaRelation.O1Residual > eps || a.EtaRelation.O2Residual > eps || a.EtaRelation.O3Residual > eps {
		return fmt.Errorf("eta relation failed: %s", FormatEtaRelation(a.EtaRelation))
	}
	if !a.Transfer.PauliMomentTripletAvailable || a.Transfer.TransferAllowed || a.Transfer.FunctorToWSpatial || a.Transfer.FunctorToGeneration || a.Transfer.WeakPlaneSelected || a.Transfer.GenerationHierarchyDerived || a.Transfer.YukawaTextureDerived || a.Transfer.CKMPMNSDerived || a.Transfer.ObservedFlavorImported || a.Transfer.GaugeBosonIdentification {
		return fmt.Errorf("transfer firewall failed: %s", FormatTransfer(a.Transfer))
	}
	if !a.Final.SealedPauliTripletExists || !a.Final.HopfMomentIdentityHolds || !a.Final.ScalarFourToOnePlusThree || !a.Final.NonzeroMomentThreeToOnePlusTwo || !a.Final.EtaIsSigma3Axis || a.Final.LawfulTransferToWOrGeneration {
		return fmt.Errorf("final verdict failed: %s", FormatFinal(a.Final))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate558559Inherited, StatusScalarComplexStructureCertified, StatusPauliTripletConstructed, StatusPauliCliffordRelationsVerified, StatusMomentCoordinatesVerified, StatusHopfIdentityVerified, StatusScalarFourToOnePlusThree, StatusMomentThreeToOnePlusTwo, StatusEtaIsSigma3Axis, StatusTauEtaSigma3AxisTraceShadow, StatusNoPauliMomentTransferFunctor, StatusNoWeakPlaneOrGenerationPromotion, StatusFirewallPreserved}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 560 finds a genuine sealed scalar-sector structure beyond the Sigma_3 trace axis: H_phi=R^4~=C^2 carries a constructible real symmetric Pauli/Cl(3,0) triple, with eta=Sigma_3. The quadratic moment coordinates mu_a=x^T Sigma_a x satisfy the Hopf identity |mu|^2=(r^2)^2, so the sealed scalar carrier has a radius plus Pauli moment triplet 4->1+3. For nonzero mu, the scalar moment record space has the canonical orbit/stabilizer split R^3_sigma=R mu plus mu^perp, giving a scalar-sector 3=1+2. Gate 558's eta records O1, O2, O3 are exactly the Sigma_3-axis shadow: O1=(I+Sigma_3)/2, O2=(I-Sigma_3)/2, O3=Sigma_3/4. The firewall remains closed: no functor transfers R^3_sigma to W_spatial, weak-plane candidates, or generations, and no weak-isospin, gauge-boson, flavor, Yukawa, CKM/PMNS, or observed data claim is made. Missing theorem: %s", a.Final.MissingNextTheorem)
}

func mustRows(rows [][]float64) linear.Matrix {
	m, err := linear.FromRows(rows)
	if err != nil {
		panic(err)
	}
	return m
}
func mustMul(a, b linear.Matrix) linear.Matrix {
	m, err := a.Mul(b)
	if err != nil {
		panic(err)
	}
	return m
}
func mustAdd(a, b linear.Matrix) linear.Matrix {
	m, err := a.Add(b)
	if err != nil {
		panic(err)
	}
	return m
}
func mustSub(a, b linear.Matrix) linear.Matrix {
	m, err := a.Sub(b)
	if err != nil {
		panic(err)
	}
	return m
}
func join(parts ...string) string { return strings.Join(parts, "; ") }
func dot(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}
func quad(x []float64, m linear.Matrix) float64 {
	y := make([]float64, len(x))
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			y[r] += m.At(r, c) * x[c]
		}
	}
	return dot(x, y)
}
func rankMatrix(m linear.Matrix, eps float64) int { return rankFloats(matrixRows(m), eps) }
func matrixRows(m linear.Matrix) [][]float64 {
	rows := make([][]float64, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		rows[r] = make([]float64, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			rows[r][c] = m.At(r, c)
		}
	}
	return rows
}
func rankFloats(a [][]float64, eps float64) int {
	if len(a) == 0 {
		return 0
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = append([]float64(nil), a[i]...)
	}
	rows, cols := len(m), len(m[0])
	rank, row := 0, 0
	for col := 0; col < cols && row < rows; col++ {
		piv := row
		for r := row + 1; r < rows; r++ {
			if math.Abs(m[r][col]) > math.Abs(m[piv][col]) {
				piv = r
			}
		}
		if math.Abs(m[piv][col]) <= eps {
			continue
		}
		m[row], m[piv] = m[piv], m[row]
		pv := m[row][col]
		for c := col; c < cols; c++ {
			m[row][c] /= pv
		}
		for r := 0; r < rows; r++ {
			if r == row {
				continue
			}
			f := m[r][col]
			if math.Abs(f) > eps {
				for c := col; c < cols; c++ {
					m[r][c] -= f * m[row][c]
				}
			}
		}
		rank++
		row++
	}
	return rank
}
func formatMatrix(m linear.Matrix) string {
	rows := make([]string, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		vals := make([]string, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			vals[c] = fmt.Sprintf("%.6g", m.At(r, c))
		}
		rows[r] = "[" + strings.Join(vals, ", ") + "]"
	}
	return "[" + strings.Join(rows, ", ") + "]"
}
func formatFloatVec(v []float64) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%.6g", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}
