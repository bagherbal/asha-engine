#!/usr/bin/env python3
"""
Reproducible calculation sheet for the source-realization framework.
This script intentionally separates: established physical constants, realization seals,
source hypotheses, predictions/retrodictions, and reference comparisons.
"""
import math, cmath, json
from pathlib import Path

pi=math.pi
sqrt=math.sqrt
GF = 1.1663787e-5 # GeV^-2
v_ref = 1.0 / math.sqrt(math.sqrt(2)*GF)
MP_CODATA_GeV = 2.4353234600842885e18
MP_SOURCE_GeV = 2.4350000000000000e18
hbarc_eVm = 1.973269804e-7
c_si = 299792458.0

refs = {
    "v": v_ref,
    "mH": 125.20,
    "me_MeV": 0.51099895000,
    "mmu_MeV": 105.6583755,
    "mtau_MeV": 1776.86,
    "dm21": 7.49e-5,
    "dm31": 2.513e-3,
    "s2theta12": 0.307,
    "s2theta13": 0.02203,
    "Vus": 0.2243,
    "Vcb": 0.0415,
    "Vub": 0.00382,
}
H0_km_s_Mpc = 67.4
Omega_m = 0.315
Omega_L = 1.0 - Omega_m
Mpc_m = 3.0856775814913673e22
H0_si = H0_km_s_Mpc*1000.0/Mpc_m
Lambda_ref = 3*Omega_L*H0_si**2/c_si**2
rho_ref = Lambda_ref * hbarc_eVm**2 * (MP_CODATA_GeV*1e9)**2
rho_quarter_ref_meV = (rho_ref**0.25)*1e3

L = 1/(8*pi)
rank_Mnu = 2
V8_dim = 8
F3_dim = 3
EndF3_dim = F3_dim**2
Lambda2X4_dim = 6
C7_dim = V8_dim - 1
S = L*(rank_Mnu*C7_dim)/(Lambda2X4_dim*V8_dim*EndF3_dim - 1)
epsilon = ((9/5)*S)**0.25
phase_q = 3*pi/8
phase_pmns = 4*phase_q
alpha_eff = (2*pi/3 - 2*phase_pmns) % (2*pi)
theta_l = rank_Mnu/EndF3_dim
Phi_bb = 2*pi/3

def compute_branch(MP_GeV):
    A_EW = -12*pi + sqrt(3)/2 + 2*S + 148*S**2
    v = MP_GeV*math.exp(A_EW)
    lam = (3/8)*(1+L)*(1/3 - S)
    mH = v*math.sqrt(2*lam)
    At = L - 5*S + 2*(72-3)*S**2
    y_t = math.exp(-At)
    Ab = 4*pi/3 - 56*S + 2*(56-3)*S**2
    y_b = math.exp(-Ab)
    Atau = 4*pi/3 + 3/10 + 7/72 - S + 0.5*(72+27)*S**2
    y_tau = math.exp(-Atau)
    mtau_MeV = v/math.sqrt(2)*y_tau*1000
    s12q = epsilon*math.sqrt(1+epsilon**2)
    s23q = sqrt(3)/2*epsilon**2
    s13q = 1/(2*sqrt(2))*epsilon**3
    deltaq = phase_q
    c12q = math.sqrt(1-s12q**2); c23q=math.sqrt(1-s23q**2); c13q=math.sqrt(1-s13q**2)
    eimd = cmath.exp(-1j*deltaq); eipd=cmath.exp(1j*deltaq)
    Vckm = [
        [c12q*c13q, s12q*c13q, s13q*eimd],
        [-s12q*c23q-c12q*s23q*s13q*eipd, c12q*c23q-s12q*s23q*s13q*eipd, s23q*c13q],
        [s12q*s23q-c12q*c23q*s13q*eipd, -c12q*s23q-s12q*c23q*s13q*eipd, c23q*c13q]
    ]
    Vabs=[[abs(x) for x in row] for row in Vckm]
    Jq = (c12q*c23q*c13q**2*s12q*s23q*s13q*math.sin(deltaq))
    Vud,Vus,Vub = Vckm[0]
    Vcd,Vcs,Vcb = Vckm[1]
    Vtd,Vts,Vtb = Vckm[2]
    alpha = cmath.phase(-Vtd*Vtb.conjugate()/(Vud*Vub.conjugate()))
    beta = cmath.phase(-Vcd*Vcb.conjugate()/(Vtd*Vtb.conjugate()))
    gamma = cmath.phase(-Vud*Vub.conjugate()/(Vcd*Vcb.conjugate()))
    def pos_deg(x):
        deg=math.degrees(x)
        if deg<0: deg+=360
        return deg
    alpha_deg,beta_deg,gamma_deg = map(pos_deg,[alpha,beta,gamma])
    theta=theta_l
    amps=[1+sqrt(2)*math.cos(theta+2*pi*k/3) for k in range(3)]
    A = math.sqrt(mtau_MeV)/amps[0]
    masses_lep=[(A*a)**2 for a in amps]
    me_pred, mmu_pred = masses_lep[1], masses_lep[2]
    Koide=(sum(masses_lep))/(sum(math.sqrt(x) for x in masses_lep)**2)
    theta12 = pi/6 + 48*S
    theta13 = 4*L - 7*S + 4*S**2
    theta23_minus = pi/4 - 48*S
    theta23_plus = pi/4 + 48*S
    s2theta12 = math.sin(theta12)**2
    s2theta13 = math.sin(theta13)**2
    s2theta23m = math.sin(theta23_minus)**2
    s2theta23p = math.sin(theta23_plus)**2
    MR3 = (math.sqrt(2*pi)+49*S+90*S**2)*math.sqrt(v*MP_GeV)
    m3_eV = (v**2*math.exp(-2*Atau)/(2*MR3))*1e9
    m2_eV = (4*L+10*S)*m3_eV
    MR2 = MR3*math.exp(-4*pi/3)/(4*L+10*S)
    dm21=m2_eV**2
    dm31=m3_eV**2
    a = m2_eV*(math.sin(theta12)**2)*(math.cos(theta13)**2)
    b = m3_eV*(math.sin(theta13)**2)
    m_bb = math.sqrt(a*a+b*b-a*b)
    rho = m_bb**4
    Lambda = (rho/(MP_GeV*1e9)**2)/(hbarc_eVm**2)
    s12=math.sin(theta12); c12=math.cos(theta12); s13=math.sin(theta13); c13=math.cos(theta13); s23=math.sin(theta23_minus); c23=math.cos(theta23_minus); delta=phase_pmns
    eimd=cmath.exp(-1j*delta); eipd=cmath.exp(1j*delta)
    U=[
        [c12*c13, s12*c13, s13*eimd],
        [-s12*c23-c12*s23*s13*eipd, c12*c23-s12*s23*s13*eipd, s23*c13],
        [s12*s23-c12*c23*s13*eipd, -c12*s23-s12*c23*s13*eipd, c23*c13]
    ]
    Uabs=[[abs(x) for x in row] for row in U]
    return locals()

branch=compute_branch(MP_CODATA_GeV)
branch_src=compute_branch(MP_SOURCE_GeV)

def rel(pred, ref):
    return (pred-ref)/ref

def fmt(x, sig=6):
    if x == 0: return "0"
    ax=abs(x)
    if ax>=1e5 or ax<1e-3:
        return f"{x:.{sig}e}"
    return f"{x:.{sig}g}"

def texnum(x, sig=6):
    s=fmt(x,sig)
    if 'e' in s:
        base, exp=s.split('e')
        return base + r"\times 10^{" + str(int(exp)) + "}"
    return s

def row(name, pred, ref, status):
    r=rel(pred,ref) if ref not in (None,0) else None
    return {"name":name,"pred":pred,"ref":ref,"rel":r,"status":status}

rows=[
    row(r"$v$ [GeV]", branch['v'], refs['v'], "source-level; CODATA Planck seal"),
    row(r"$m_H$ [GeV]", branch['mH'], refs['mH'], "tree-level bridge; loop transport required"),
    row(r"$m_\tau$ [MeV]", branch['mtau_MeV'], refs['mtau_MeV'], "source-level edge value"),
    row(r"$m_e$ [MeV]", branch['me_pred'], refs['me_MeV'], "Koide-cone source value; residual may be transport"),
    row(r"$m_\mu$ [MeV]", branch['mmu_pred'], refs['mmu_MeV'], "Koide-cone source value; residual may be transport"),
    row(r"$\sin^2\theta_{12}$", branch['s2theta12'], refs['s2theta12'], "NuFIT comparison; source skeleton"),
    row(r"$\sin^2\theta_{13}$", branch['s2theta13'], refs['s2theta13'], "NuFIT comparison; source skeleton"),
    row(r"$\Delta m^2_{21}$ [eV$^2$]", branch['dm21'], refs['dm21'], "NuFIT comparison"),
    row(r"$\Delta m^2_{31}$ [eV$^2$]", branch['dm31'], refs['dm31'], "NuFIT comparison"),
    row(r"$\Lambda$ [m$^{-2}$]", branch['Lambda'], Lambda_ref, "Planck-2018 comparison; cosmological uncertainty dominates"),
    row(r"$\rho_\Lambda^{1/4}$ [meV]", branch['rho']**0.25*1e3, rho_quarter_ref_meV, "dark-energy fourth-root scale"),
    row(r"$|V_{us}|$", branch['Vabs'][0][1], refs['Vus'], "representative CKM comparison"),
    row(r"$|V_{cb}|$", branch['Vabs'][1][2], refs['Vcb'], "representative CKM comparison"),
    row(r"$|V_{ub}|$", branch['Vabs'][0][2], refs['Vub'], "representative CKM comparison"),
]

predictions=[
    (r"Lightest neutrino mass", r"$m_1=0$", "Can be falsified by absolute-mass data incompatible with rank two."),
    (r"Sum of neutrino masses", fr"$\sum m_\nu={texnum(branch['m2_eV']+branch['m3_eV'],7)}\,\mathrm{{eV}}$", "Within current cosmological upper bounds; future tightening tests it."),
    (r"Atmospheric octant branch", fr"$\theta_{{23}}={math.degrees(branch['theta23_minus']):.6f}^\circ$; $\sin^2\theta_{{23}}={branch['s2theta23m']:.6f}$", "Killed if the upper octant is established decisively."),
    (r"Leptonic CP phase", r"$\delta_{\rm PMNS}=3\pi/2$", "Future long-baseline experiments can test/exclude."),
    (r"Effective Majorana mass", fr"$m_{{\beta\beta}}={branch['m_bb']*1e3:.6f}\,\mathrm{{meV}}$", "Below near-term $0\nu\beta\beta$ reach; sharp long-term prediction."),
    (r"Heavy Majorana scales", fr"$M_{{R2}}={texnum(branch['MR2'],6)}\,\mathrm{{GeV}},\;M_{{R3}}={texnum(branch['MR3'],6)}\,\mathrm{{GeV}}$", "Model-internal high-scale prediction; indirectly testable."),
    (r"Vacuum-boundary relation", r"$\rho_\Lambda=|M_{\nu,ee}|^4$", "Killed if improved neutrino and cosmology data break the relation."),
    (r"CKM source phase", r"$\delta_q=3\pi/8$", "Already close to CKM fits; refined global fits can stress it."),
]

calc_table=[]
for rr in rows:
    pred=texnum(rr['pred'],7)
    ref=texnum(rr['ref'],7) if rr['ref'] is not None else "--"
    rels=(texnum(rr['rel'],4) if rr['rel'] is not None else "--")
    calc_table.append(f"{rr['name']} & ${pred}$ & ${ref}$ & ${rels}$ & {rr['status']} \\\\")

pred_table=[]
for p in predictions:
    pred_table.append(f"{p[0]} & {p[1]} & {p[2]} \\\\")

def matrix_tex(M, sig=6):
    lines=[]
    for rowv in M:
        lines.append(" & ".join(texnum(x, sig) for x in rowv))
    return "\\begin{pmatrix}\n" + "\\\\\n".join(lines) + "\n\\end{pmatrix}"

outdir=Path('/mnt/data/lawful_toe_build'); outdir.mkdir(exist_ok=True)
(outdir/'calculation_table.tex').write_text("\n".join(calc_table))
(outdir/'prediction_table.tex').write_text("\n".join(pred_table))
(outdir/'ckm_matrix.tex').write_text(matrix_tex(branch['Vabs'],6))
(outdir/'pmns_matrix.tex').write_text(matrix_tex(branch['Uabs'],6))

selected={
    "L":L,"S_star":S,"epsilon":epsilon,"theta_l":theta_l,"delta_q_rad":phase_q,"delta_pmns_rad":phase_pmns,"alpha_eff_rad":alpha_eff,
    "v_CODATA":branch['v'],"lambda":branch['lam'],"mH":branch['mH'],"y_t":branch['y_t'],"y_b":branch['y_b'],"y_tau":branch['y_tau'],"m_tau_MeV":branch['mtau_MeV'],"m_e_MeV":branch['me_pred'],"m_mu_MeV":branch['mmu_pred'],
    "MR2":branch['MR2'],"MR3":branch['MR3'],"m2_eV":branch['m2_eV'],"m3_eV":branch['m3_eV'],"m_bb_eV":branch['m_bb'],"rho_eV4":branch['rho'],"Lambda_m2":branch['Lambda'],
    "Vckm_abs":branch['Vabs'],"Jq":branch['Jq'],"UT_angles_deg":[branch['alpha_deg'],branch['beta_deg'],branch['gamma_deg']],
    "PMNS_abs":branch['Uabs'],"Lambda_ref":Lambda_ref,"rho_ref_eV4":rho_ref,
    "source_normalization_v":branch_src['v'],"source_normalization_mH":branch_src['mH'],"source_normalization_mtau":branch_src['mtau_MeV']
}
(outdir/'calculation_data.json').write_text(json.dumps(selected,indent=2))
macros = fr"""
\newcommand{{\Lval}}{{{texnum(L,10)}}}
\newcommand{{\Sval}}{{{texnum(S,10)}}}
\newcommand{{\epsval}}{{{texnum(epsilon,10)}}}
\newcommand{{\vval}}{{{texnum(branch['v'],10)}}}
\newcommand{{\lambdaval}}{{{texnum(branch['lam'],10)}}}
\newcommand{{\mHval}}{{{texnum(branch['mH'],10)}}}
\newcommand{{\ytval}}{{{texnum(branch['y_t'],10)}}}
\newcommand{{\ybval}}{{{texnum(branch['y_b'],10)}}}
\newcommand{{\ytauval}}{{{texnum(branch['y_tau'],10)}}}
\newcommand{{\meval}}{{{texnum(branch['me_pred'],10)}}}
\newcommand{{\mmuval}}{{{texnum(branch['mmu_pred'],10)}}}
\newcommand{{\mtauval}}{{{texnum(branch['mtau_MeV'],10)}}}
\newcommand{{\MRtwoval}}{{{texnum(branch['MR2'],10)}}}
\newcommand{{\MRthreeval}}{{{texnum(branch['MR3'],10)}}}
\newcommand{{\mtwoval}}{{{texnum(branch['m2_eV'],10)}}}
\newcommand{{\mthreeval}}{{{texnum(branch['m3_eV'],10)}}}
\newcommand{{\mbbval}}{{{texnum(branch['m_bb']*1e3,10)}}}
\newcommand{{\rhoval}}{{{texnum(branch['rho'],10)}}}
\newcommand{{\Lambdaval}}{{{texnum(branch['Lambda'],10)}}}
\newcommand{{\Jqval}}{{{texnum(branch['Jq'],10)}}}
"""
(outdir/'macros.tex').write_text(macros)
print(json.dumps(selected, indent=2))
