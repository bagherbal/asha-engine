// Package generation2vacuumtilt implements Gate 484:
// Vacuum Tilt Vector / C3 Elliptic Slice Flavor Compression Audit.
package generation2vacuumtilt

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE484-VACUUM-TILT-VECTOR-C3-ELLIPTIC-SLICE-FLAVOR-COMPRESSION-AUDIT"

	StatusAuditCompleted                = "CONDITIONAL_SUPPORT_VACUUM_TILT_AUDIT_COMPLETED"
	StatusC3OrbitBasisValidated         = "CONDITIONAL_SUPPORT_C3_ORBIT_BASIS_VALIDATED"
	StatusChargedLeptonKoideShadowFound = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_KOIDE_SHADOW_FOUND"
	StatusFirewallPreserved             = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_VACUUM_TILT_AUDIT"

	StatusFailedTiltSliceReparam         = "FAILED_ROUTE_TILTED_SLICE_REPARAMETERIZES_FLAVOR_MODULI"
	StatusFailedKoideNotAllSectors       = "FAILED_ROUTE_KOIDE_RELATION_NOT_NATIVE_FOR_ALL_SECTORS"
	StatusFailedUniversalTiltUnsupported = "FAILED_ROUTE_UNIVERSAL_TILT_VECTOR_NOT_SUPPORTED"
	StatusFailedNoNativeTiltRatio        = "FAILED_ROUTE_NATIVE_NULL_CONE_DOES_NOT_FIX_TILT_RATIO"
	StatusFailedCKMPMNSPrediction        = "FAILED_ROUTE_VACUUM_TILT_AS_CKM_PMNS_PREDICTION_REJECTED"
	StatusFailedNativePromotion          = "FAILED_ROUTE_VACUUM_TILT_NATIVE_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	twoPiOverThree  = 2 * math.Pi / 3
	koideTarget     = 2.0 / 3.0
	koideRatio      = math.Sqrt2
)

type Inheritance struct {
	Executed, Gate480NullBaseline, Gate481Cancellation, Gate482SourceAbsent, Gate483TopologyNoGo bool
	AlphaVac, IKVac                                                                              float64
	SectorPerturbationsUnsolved, NativeRegistryClean                                             bool
	Verdict                                                                                      string
}

type SectorInput struct {
	Name   string
	Masses [3]float64
	Labels [3]string
}

type SectorDecomposition struct {
	Name                              string
	Labels                            [3]string
	Masses                            [3]float64
	SqrtMasses                        [3]float64
	Theta                             [3]float64
	S, A, B, R, RoverS, Psi           float64
	KoideQ, KoideResidual             float64
	KoideRatioResidual                float64
	Reconstructed                     [3]float64
	MaxReconstructionError            float64
	PassesKoideLikeFixedRatio         bool
	CoordinatesAreExactC3BasisRewrite bool
}

type C3BasisAudit struct {
	Executed                                                  bool
	BasisVectors                                              [3][3]float64
	Sectors                                                   []SectorDecomposition
	AllSectorsExactlyRepresented, RepresentationModuliNeutral bool
	DataDOF, BasisCoefficientDOF                              int
	Verdict, Reason                                           string
}

type KoideAudit struct {
	Executed                                                      bool
	TargetQ, TargetRoverS                                         float64
	ChargedLeptonPasses, UpQuarkPasses, DownQuarkPasses           bool
	ChargedLeptonResidual, UpQuarkResidual, DownQuarkResidual     float64
	LeptonStrongerThanQuarks                                      bool
	NativeAllSectorRelationFound, NativeTiltRatioForcedByNullCone bool
	Verdict, Reason                                               string
}

type UniversalTiltAudit struct {
	Executed                                                                               bool
	IndependentSectorTiltAnglesRequired                                                    bool
	OneUniversalPsiSupported, OneUniversalRoverSSupported, OneUniversalTiltVectorSupported bool
	SectorPsiSpread, SectorRoverSSpread                                                    float64
	ExactFitDOF, UniversalTiltDOF, ChargedMassDataDOF                                      int
	ReducesModuli                                                                          bool
	PredictsNontrivialRelation                                                             bool
	Verdict, Reason                                                                        string
}

type CompressionAudit struct {
	Executed                                                                    bool
	ExactC3RepresentationOnlyCoordinateChange                                   bool
	PerSectorFreeParameters, ChargedMassObservables                             int
	KoideReducesLeptonIfAssumed, KoideNativeForAllSectors, UniversalTiltReduces bool
	FlavorModuliReducedByCurrentGate                                            bool
	Verdict, Reason                                                             string
}

type Firewall struct {
	Executed                                                                               bool
	ObservedMassesUsedAsBridgeAuditData                                                    bool
	CKMImported, PMNSImported                                                              bool
	VacuumIKNativeBaseline                                                                 bool
	TiltRatioNative, UniversalTiltNative, SectorPerturbationsNative                        bool
	PhysicalDUDComputed, PhysicalDENuComputed, CKMMatrixConstructed, PMNSMatrixConstructed bool
	NativeRegistryWritten                                                                  bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                 int
	Verdict, Reason                                                                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance   Inheritance
	C3BasisAudit  C3BasisAudit
	KoideAudit    KoideAudit
	UniversalTilt UniversalTiltAudit
	Compression   CompressionAudit
	Firewall      Firewall
	Next          NextStep
	Truth         string
}

func BuildDefault() (Analysis, error) {
	inheritance := buildInheritance()
	basis := buildC3BasisAudit(defaultSectorInputs())
	koide := buildKoideAudit(basis)
	universal := buildUniversalTiltAudit(basis, koide)
	compression := buildCompressionAudit(basis, koide, universal)
	analysis := Analysis{Inheritance: inheritance, C3BasisAudit: basis, KoideAudit: koide, UniversalTilt: universal, Compression: compression}
	analysis.Firewall = buildFirewall(analysis)
	analysis.Next = buildNext()
	analysis.Truth = truth(analysis)
	if err := validate(analysis); err != nil {
		return Analysis{}, err
	}
	return analysis, nil
}

func defaultSectorInputs() []SectorInput {
	return []SectorInput{
		{Name: "up-type quarks", Labels: [3]string{"u", "c", "t"}, Masses: [3]float64{2.16, 1270.0, 172500.0}},
		{Name: "down-type quarks", Labels: [3]string{"d", "s", "b"}, Masses: [3]float64{4.67, 93.0, 4180.0}},
		{Name: "charged leptons", Labels: [3]string{"e", "mu", "tau"}, Masses: [3]float64{0.511, 105.6, 1776.0}},
	}
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate480NullBaseline: true, Gate481Cancellation: true, Gate482SourceAbsent: true, Gate483TopologyNoGo: true, AlphaVac: 1, IKVac: 0.5, SectorPerturbationsUnsolved: true, NativeRegistryClean: true, Verdict: StatusAuditCompleted}
}

func buildC3BasisAudit(inputs []SectorInput) C3BasisAudit {
	basisVectors := [3][3]float64{
		{1, 1, 1},
		{math.Cos(0), math.Cos(twoPiOverThree), math.Cos(2 * twoPiOverThree)},
		{math.Sin(0), math.Sin(twoPiOverThree), math.Sin(2 * twoPiOverThree)},
	}
	sectors := make([]SectorDecomposition, 0, len(inputs))
	allExact := true
	for _, in := range inputs {
		d := decomposeSector(in)
		if d.MaxReconstructionError > 1e-9 {
			allExact = false
		}
		sectors = append(sectors, d)
	}
	dataDOF := 3 * len(inputs)
	basisDOF := 3 * len(inputs)
	return C3BasisAudit{Executed: true, BasisVectors: basisVectors, Sectors: sectors, AllSectorsExactlyRepresented: allExact, RepresentationModuliNeutral: dataDOF == basisDOF, DataDOF: dataDOF, BasisCoefficientDOF: basisDOF, Verdict: StatusC3OrbitBasisValidated, Reason: "the C3 orbit basis {1, cos(theta_i), sin(theta_i)} exactly reconstructs every three-generation square-root mass vector, but this exactness is only a full-rank coordinate decomposition"}
}

func decomposeSector(in SectorInput) SectorDecomposition {
	theta := [3]float64{0, twoPiOverThree, 2 * twoPiOverThree}
	var x [3]float64
	for i, m := range in.Masses {
		x[i] = math.Sqrt(m)
	}
	S := (x[0] + x[1] + x[2]) / 3
	A := 0.0
	B := 0.0
	for i := range x {
		A += x[i] * math.Cos(theta[i])
		B += x[i] * math.Sin(theta[i])
	}
	A *= 2.0 / 3.0
	B *= 2.0 / 3.0
	R := math.Hypot(A, B)
	psi := math.Atan2(B, A)
	var rec [3]float64
	maxErr := 0.0
	for i := range x {
		rec[i] = S + A*math.Cos(theta[i]) + B*math.Sin(theta[i])
		if err := math.Abs(rec[i] - x[i]); err > maxErr {
			maxErr = err
		}
	}
	mSum := in.Masses[0] + in.Masses[1] + in.Masses[2]
	xSum := x[0] + x[1] + x[2]
	Q := mSum / (xSum * xSum)
	ratio := R / S
	return SectorDecomposition{Name: in.Name, Labels: in.Labels, Masses: in.Masses, SqrtMasses: x, Theta: theta, S: S, A: A, B: B, R: R, RoverS: ratio, Psi: psi, KoideQ: Q, KoideResidual: Q - koideTarget, KoideRatioResidual: ratio - koideRatio, Reconstructed: rec, MaxReconstructionError: maxErr, PassesKoideLikeFixedRatio: math.Abs(Q-koideTarget) < 1e-4 && math.Abs(ratio-koideRatio) < 1e-3, CoordinatesAreExactC3BasisRewrite: maxErr < 1e-9}
}

func buildKoideAudit(b C3BasisAudit) KoideAudit {
	var up, down, lepton SectorDecomposition
	for _, s := range b.Sectors {
		switch s.Name {
		case "up-type quarks":
			up = s
		case "down-type quarks":
			down = s
		case "charged leptons":
			lepton = s
		}
	}
	leptonStrong := math.Abs(lepton.KoideResidual) < math.Abs(up.KoideResidual) && math.Abs(lepton.KoideResidual) < math.Abs(down.KoideResidual)
	allNative := lepton.PassesKoideLikeFixedRatio && up.PassesKoideLikeFixedRatio && down.PassesKoideLikeFixedRatio
	return KoideAudit{Executed: true, TargetQ: koideTarget, TargetRoverS: koideRatio, ChargedLeptonPasses: lepton.PassesKoideLikeFixedRatio, UpQuarkPasses: up.PassesKoideLikeFixedRatio, DownQuarkPasses: down.PassesKoideLikeFixedRatio, ChargedLeptonResidual: lepton.KoideResidual, UpQuarkResidual: up.KoideResidual, DownQuarkResidual: down.KoideResidual, LeptonStrongerThanQuarks: leptonStrong, NativeAllSectorRelationFound: allNative, NativeTiltRatioForcedByNullCone: false, Verdict: StatusFailedKoideNotAllSectors, Reason: "the charged-lepton square-root vector is extremely close to the Koide R/S=sqrt(2) circle condition, but the up and down quark sectors are not; the current null-cone ledger does not prove a universal native tilt ratio"}
}

func buildUniversalTiltAudit(b C3BasisAudit, k KoideAudit) UniversalTiltAudit {
	minPsi, maxPsi := math.Inf(1), math.Inf(-1)
	minRatio, maxRatio := math.Inf(1), math.Inf(-1)
	for _, s := range b.Sectors {
		if s.Psi < minPsi {
			minPsi = s.Psi
		}
		if s.Psi > maxPsi {
			maxPsi = s.Psi
		}
		if s.RoverS < minRatio {
			minRatio = s.RoverS
		}
		if s.RoverS > maxRatio {
			maxRatio = s.RoverS
		}
	}
	psiSpread := maxPsi - minPsi
	ratioSpread := maxRatio - minRatio
	onePsi := psiSpread < 1e-3
	oneRatio := ratioSpread < 1e-3
	oneVector := onePsi && oneRatio
	return UniversalTiltAudit{Executed: true, IndependentSectorTiltAnglesRequired: !oneVector, OneUniversalPsiSupported: onePsi, OneUniversalRoverSSupported: oneRatio, OneUniversalTiltVectorSupported: oneVector, SectorPsiSpread: psiSpread, SectorRoverSSpread: ratioSpread, ExactFitDOF: b.BasisCoefficientDOF, UniversalTiltDOF: 5, ChargedMassDataDOF: b.DataDOF, ReducesModuli: false, PredictsNontrivialRelation: false, Verdict: StatusFailedUniversalTiltUnsupported, Reason: fmt.Sprintf("sector tilt ratios and phases differ (spread R/S=%s, psi=%s rad); a single universal vacuum tilt vector with only sector scales cannot reproduce the three charged-sector mass shadows", fmtFloat(ratioSpread), fmtFloat(psiSpread))}
}

func buildCompressionAudit(b C3BasisAudit, k KoideAudit, u UniversalTiltAudit) CompressionAudit {
	return CompressionAudit{Executed: true, ExactC3RepresentationOnlyCoordinateChange: b.RepresentationModuliNeutral, PerSectorFreeParameters: b.BasisCoefficientDOF, ChargedMassObservables: b.DataDOF, KoideReducesLeptonIfAssumed: k.ChargedLeptonPasses, KoideNativeForAllSectors: k.NativeAllSectorRelationFound, UniversalTiltReduces: u.ReducesModuli, FlavorModuliReducedByCurrentGate: false, Verdict: StatusFailedTiltSliceReparam, Reason: "without a native fixed tilt ratio or a universal sector-shared tilt vector, the tilted-slice model uses independent S,R,psi per sector and exactly reparametrizes the charged mass ledger rather than reducing it"}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, ObservedMassesUsedAsBridgeAuditData: true, CKMImported: false, PMNSImported: false, VacuumIKNativeBaseline: true, TiltRatioNative: false, UniversalTiltNative: false, SectorPerturbationsNative: false, PhysicalDUDComputed: false, PhysicalDENuComputed: false, CKMMatrixConstructed: false, PMNSMatrixConstructed: false, NativeRegistryWritten: false, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate484 treats mass fingerprints only as bridge audit data; no C3 tilt coefficient, Koide constraint, CKM/PMNS residual, or sector perturbation is promoted into native law-space"}
}

func buildNext() NextStep {
	return NextStep{Gate: 485, Title: "Koide constraint provenance or closure", Reason: "Gate484 finds a strong charged-lepton Koide shadow but no native universal tilt ratio and no cross-sector compression.", PrimaryTask: "audit whether R/S=sqrt(2) can be derived from a finite C3/null-cone constraint, rather than imposed as an empirical charged-lepton relation"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate480NullBaseline || !a.Inheritance.Gate481Cancellation || !a.Inheritance.Gate482SourceAbsent || !a.Inheritance.Gate483TopologyNoGo || a.Inheritance.AlphaVac != 1 || a.Inheritance.IKVac != 0.5 || !a.Inheritance.SectorPerturbationsUnsolved || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate484 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.C3BasisAudit.Executed || !a.C3BasisAudit.AllSectorsExactlyRepresented || !a.C3BasisAudit.RepresentationModuliNeutral || a.C3BasisAudit.DataDOF != 9 || a.C3BasisAudit.BasisCoefficientDOF != 9 || len(a.C3BasisAudit.Sectors) != 3 {
		return fmt.Errorf("Gate484 C3 basis invalid: %+v", a.C3BasisAudit)
	}
	if !a.KoideAudit.Executed || !a.KoideAudit.ChargedLeptonPasses || a.KoideAudit.UpQuarkPasses || a.KoideAudit.DownQuarkPasses || !a.KoideAudit.LeptonStrongerThanQuarks || a.KoideAudit.NativeAllSectorRelationFound || a.KoideAudit.NativeTiltRatioForcedByNullCone {
		return fmt.Errorf("Gate484 Koide audit invalid: %+v", a.KoideAudit)
	}
	if !a.UniversalTilt.Executed || !a.UniversalTilt.IndependentSectorTiltAnglesRequired || a.UniversalTilt.OneUniversalTiltVectorSupported || a.UniversalTilt.ReducesModuli || a.UniversalTilt.PredictsNontrivialRelation {
		return fmt.Errorf("Gate484 universal tilt audit invalid: %+v", a.UniversalTilt)
	}
	if !a.Compression.Executed || !a.Compression.ExactC3RepresentationOnlyCoordinateChange || a.Compression.PerSectorFreeParameters != a.Compression.ChargedMassObservables || !a.Compression.KoideReducesLeptonIfAssumed || a.Compression.KoideNativeForAllSectors || a.Compression.UniversalTiltReduces || a.Compression.FlavorModuliReducedByCurrentGate {
		return fmt.Errorf("Gate484 compression audit invalid: %+v", a.Compression)
	}
	if !a.Firewall.Executed || !a.Firewall.ObservedMassesUsedAsBridgeAuditData || a.Firewall.CKMImported || a.Firewall.PMNSImported || !a.Firewall.VacuumIKNativeBaseline || a.Firewall.TiltRatioNative || a.Firewall.UniversalTiltNative || a.Firewall.SectorPerturbationsNative || a.Firewall.PhysicalDUDComputed || a.Firewall.PhysicalDENuComputed || a.Firewall.CKMMatrixConstructed || a.Firewall.PMNSMatrixConstructed || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate484 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate484 result: C3 tilted-slice coordinates exactly represent the u/d/e square-root mass fingerprints, and charged leptons nearly satisfy Koide (Q residual %s), but no native universal tilt ratio or cross-sector tilt vector is found; the model reparametrizes rather than reduces the flavor moduli.", fmtFloat(a.KoideAudit.ChargedLeptonResidual))
}

func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 484 Registry Audit — Vacuum Tilt Vector / C3 Elliptic Slice Flavor Compression Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusChargedLeptonKoideShadowFound + "\n")
	b.WriteString(StatusFailedTiltSliceReparam + "\n")
	b.WriteString(StatusFailedKoideNotAllSectors + "\n")
	b.WriteString(StatusFailedUniversalTiltUnsupported + "\n")
	b.WriteString("```\n\n")
	b.WriteString("Gate 484 validates the C3 tilted-slice coordinate system, but it does not reduce the flavor firewall. The charged-lepton Koide shadow is real in the supplied bridge data; it is not yet native for all sectors.\n\n")

	b.WriteString("## Inherited boundary\n\n")
	b.WriteString(fmt.Sprintf("- α_vac = `%s`\n", fmtFloat(a.Inheritance.AlphaVac)))
	b.WriteString(fmt.Sprintf("- I_K,vac = `%s`\n", fmtFloat(a.Inheritance.IKVac)))
	b.WriteString("- Gate 481 cancellation: common null baselines cancel from relative distances.\n")
	b.WriteString("- Gate 483 no-go: native color/winding topology separates quark/lepton sectors but is generation-blind.\n\n")

	b.WriteString("## C3 tilted-slice decomposition\n\n")
	b.WriteString("The basis is `x_i = sqrt(m_i) = S + A cos(θ_i) + B sin(θ_i)`, with `θ_i = 0, 2π/3, 4π/3`.\n\n")
	b.WriteString("| sector | S | R/S | ψ rad | Q=(Σm)/(Σ√m)^2 | Q-2/3 | max reconstruction error |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|\n")
	for _, s := range a.C3BasisAudit.Sectors {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n", s.Name, fmtFloat(s.S), fmtFloat(s.RoverS), fmtFloat(s.Psi), fmtFloat(s.KoideQ), fmtFloat(s.KoideResidual), fmtFloat(s.MaxReconstructionError)))
	}
	b.WriteString("\nThis exact reconstruction is not a theorem of mass hierarchy. It is a full-rank C3 coordinate decomposition of three input numbers.\n\n")

	b.WriteString("## Koide/fixed-ratio audit\n\n")
	b.WriteString(fmt.Sprintf("- Koide target: `Q = 2/3 = %s`.\n", fmtFloat(a.KoideAudit.TargetQ)))
	b.WriteString(fmt.Sprintf("- Equivalent C3 tilt target: `R/S = sqrt(2) = %s`.\n", fmtFloat(a.KoideAudit.TargetRoverS)))
	b.WriteString(fmt.Sprintf("- Charged-lepton residual: `%s` — passes the bridge-level Koide shadow test.\n", fmtFloat(a.KoideAudit.ChargedLeptonResidual)))
	b.WriteString(fmt.Sprintf("- Up-quark residual: `%s` — fails the same fixed-ratio condition.\n", fmtFloat(a.KoideAudit.UpQuarkResidual)))
	b.WriteString(fmt.Sprintf("- Down-quark residual: `%s` — fails the same fixed-ratio condition.\n", fmtFloat(a.KoideAudit.DownQuarkResidual)))
	b.WriteString("\nThe current Cℓ(1,7) null-cone ledger fixes the vacuum baseline α_vac = 1, but it does not prove `R/S = sqrt(2)` for all charged sectors.\n\n")

	b.WriteString("## Universal tilt-vector audit\n\n")
	b.WriteString(fmt.Sprintf("- Sector `R/S` spread: `%s`.\n", fmtFloat(a.UniversalTilt.SectorRoverSSpread)))
	b.WriteString(fmt.Sprintf("- Sector phase `ψ` spread: `%s` rad.\n", fmtFloat(a.UniversalTilt.SectorPsiSpread)))
	b.WriteString(fmt.Sprintf("- Exact per-sector tilted-slice DOF: `%d` for `%d` charged mass observables.\n", a.UniversalTilt.ExactFitDOF, a.UniversalTilt.ChargedMassDataDOF))
	b.WriteString(fmt.Sprintf("- Universal tilt ansatz DOF: `%d`, but it does not fit the sector shadows under the current audit.\n\n", a.UniversalTilt.UniversalTiltDOF))

	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n")
	b.WriteString("physical d_ud = undefined\n")
	b.WriteString("physical d_eν = undefined\n")
	b.WriteString("CKM/PMNS = not constructed\n")
	b.WriteString("native registry write = false\n")
	b.WriteString("```\n\n")
	b.WriteString("Gate 484 may motivate a later Koide-provenance gate, but it does not by itself collapse the 13 charged flavor moduli.\n\n")

	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s. %s\n\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
