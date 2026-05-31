from __future__ import annotations

from pathlib import Path
from math import log10, sin, cos, pi
import html

from .style import PALETTE, STYLE, FigureStyle
from .readme_next5 import FigureContract, L, S_SPLIT


def esc(s: object) -> str:
    return html.escape(str(s), quote=True)


def svg_open(title: str, desc: str, style: FigureStyle = STYLE) -> list[str]:
    p=PALETTE; w=style.width; h=style.height
    return [f'''<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}" role="img" aria-labelledby="title desc">
<title id="title">{esc(title)}</title>
<desc id="desc">{esc(desc)}</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="72%"><stop offset="0%" stop-color="#10101B"/><stop offset="62%" stop-color="{p.abyss}"/><stop offset="100%" stop-color="#010103"/></radialGradient>
  <linearGradient id="goldLine" x1="0" x2="1"><stop offset="0%" stop-color="#7D6434" stop-opacity="0.25"/><stop offset="50%" stop-color="{p.gold}" stop-opacity="0.95"/><stop offset="100%" stop-color="#FFF2B8" stop-opacity="0.45"/></linearGradient>
  <linearGradient id="cyanLine" x1="0" x2="1"><stop offset="0%" stop-color="#2A6C75" stop-opacity="0.25"/><stop offset="55%" stop-color="{p.cyan}" stop-opacity="0.95"/><stop offset="100%" stop-color="#D7F7FF" stop-opacity="0.45"/></linearGradient>
  <linearGradient id="obsidian" x1="0" y1="0" x2="1" y2="1"><stop offset="0%" stop-color="#14141C"/><stop offset="70%" stop-color="#050508"/><stop offset="100%" stop-color="#000000"/></linearGradient>
  <filter id="softGlow" x="-200%" y="-200%" width="500%" height="500%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-250%" y="-250%" width="600%" height="600%"><feGaussianBlur stdDeviation="16" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="12" markerHeight="12" refX="10" refY="6" orient="auto"><path d="M 0 0 L 12 6 L 0 12 z" fill="{p.platinum}" fill-opacity="0.72"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 62px; letter-spacing: 0.3px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 19px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .label {{ font-family: {style.label_font}; font-size: 20px; letter-spacing: 3px; fill: {p.muted}; text-transform: uppercase; }}
    .caption {{ font-family: {style.label_font}; font-size: 22px; fill: #AAB3C2; }}
    .tiny {{ font-family: {style.label_font}; font-size: 15px; letter-spacing: 1.6px; fill: #8B93A5; }}
    .math {{ font-family: {style.title_font}; font-size: 36px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 26px; fill: {p.platinum}; }}
    .gold {{ fill: {p.gold}; }} .cyan {{ fill: {p.cyan}; }} .warm {{ fill: {p.warm}; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{w-96}" height="{h-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README VISUAL ATLAS · SOURCE-TYPED FIGURE</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
''']


def svg_close(footer_left: str, footer_right: str, style: FigureStyle = STYLE) -> str:
    w=style.width
    return f'''<g transform="translate(165 2128)">
  <text class="label" x="0" y="0">BOUNDARY</text>
  <text class="caption" x="0" y="48">{esc(footer_left)}</text>
</g>
<text class="tiny" x="{w-165}" y="2240" text-anchor="end">{esc(footer_right)}</text>
</svg>'''


def render_matter_sockets(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p=PALETTE; w=style.width
    parts=svg_open(c.title, "Finite matter algebra blocks tensor-broadcast through the contact-depth triple.", style)
    parts.append(f'<text class="math" x="{w/2:.0f}" y="292" text-anchor="middle">Adepth = AF ⊗ Qcontact³</text>')
    # finite blocks
    blocks=[("C",1,370,650,130,130), ("H",4,370,1030,220,220), ("M₃(C)",9,370,1515,310,310)]
    for name,cells,cx,cy,bw,bh in blocks:
        parts.append(f'<rect x="{cx-bw/2}" y="{cy-bh/2}" width="{bw}" height="{bh}" rx="30" fill="#071018" fill-opacity="0.72" stroke="{p.cyan}" stroke-opacity="0.42" stroke-width="1.6"/>')
        n=int(cells**0.5)
        if cells==1:
            parts.append(f'<circle cx="{cx}" cy="{cy}" r="27" fill="{p.gold}" filter="url(#softGlow)"/>')
        else:
            gap=44 if cells==4 else 46
            startx=cx-gap*(n-1)/2; starty=cy-gap*(n-1)/2
            for i in range(n):
                for j in range(n):
                    parts.append(f'<circle cx="{startx+i*gap}" cy="{starty+j*gap}" r="14" fill="{p.cyan}" fill-opacity="0.82" filter="url(#softGlow)"/>')
        parts.append(f'<text class="math" x="{cx}" y="{cy+bh/2+68}" text-anchor="middle">{name}</text>')
        parts.append(f'<text class="tiny" x="{cx}" y="{cy+bh/2+102}" text-anchor="middle">{cells} SOCKET CELL{"S" if cells>1 else ""}</text>')
    # arrows to depth layers
    parts.append(f'<path d="M 620 1110 C 760 1110 770 1110 890 1110" fill="none" stroke="{p.platinum}" stroke-opacity="0.48" stroke-width="2.4" marker-end="url(#arrow)"/>')
    parts.append(f'<text class="tiny" x="750" y="1060" text-anchor="middle">TENSOR BROADCAST</text>')
    # depth layers columns
    qs=c.quantities
    for idx,(label,nq,wq) in enumerate(zip(["Π₁","Π₂","Π₃"], qs["N_Q"], qs["W_Q"])):
        x=1025+idx*245
        height=820-idx*95
        y=1180-height/2
        parts.append(f'<rect x="{x-80}" y="{y}" width="160" height="{height}" rx="80" fill="#0B1519" fill-opacity="0.62" stroke="{p.gold if idx==0 else p.cyan}" stroke-opacity="0.45" stroke-width="1.4"/>')
        parts.append(f'<circle cx="{x}" cy="{y+95}" r="40" fill="{p.gold}" fill-opacity="0.18" stroke="{p.gold}" stroke-opacity="0.78" filter="url(#softGlow)"/>')
        parts.append(f'<text class="math" x="{x}" y="{y+106}" text-anchor="middle" fill="{p.gold}">{label}</text>')
        parts.append(f'<text class="mathSmall" x="{x}" y="{y+190}" text-anchor="middle">N={nq:.6g}</text>')
        parts.append(f'<text class="tiny" x="{x}" y="{y+228}" text-anchor="middle">W={wq:.8f}</text>')
        for row in range(3):
            yy=y+300+row*130
            parts.append(f'<line x1="{x-54}" y1="{yy}" x2="{x+54}" y2="{yy}" stroke="{p.glass}" stroke-opacity="0.18"/>')
            parts.append(f'<circle cx="{x}" cy="{yy}" r="17" fill="{p.cyan}" fill-opacity="0.82" filter="url(#softGlow)"/>')
    # route labels
    parts.append(f'<path d="M 1010 1880 C 1120 1970 1330 1970 1478 1880" fill="none" stroke="{p.gold}" stroke-opacity="0.36" stroke-width="2"/>')
    parts.append(f'<text class="mathSmall" x="1245" y="2025" text-anchor="middle" fill="{p.gold}">DY depth = DY finite ⊗ exp(-4π NQ)</text>')
    parts.append(f'<text class="tiny" x="1245" y="2065" text-anchor="middle">TOP LANE σ=0 · CONTACT LANES EXAMPLE σ≈1</text>')
    parts.append(svg_close("Socket algebra is preserved; contact depth labels are a separate broadcast factor.", "EXACT CHECK: BLOCK CELLS 1,4,9 · DEPTH LAYERS 1/3,1/2,2/3", style))
    Path(out_path).write_text("".join(parts), encoding="utf-8")


def render_source_alphabet(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p=PALETTE; w=style.width
    q=c.quantities
    parts=svg_open(c.title, "A source-typing firewall for constants, finite numbers, and coefficient decompositions.", style)
    parts.append(f'<text class="math" x="{w/2:.0f}" y="305" text-anchor="middle">L = 1/(8π) · S = {S_SPLIT:.16f}</text>')
    cx,cy=900,950
    nums=q["finite_source_numbers"]
    radius=450
    for i,n in enumerate(nums):
        ang=-pi/2 + 2*pi*i/len(nums)
        x=cx+radius*cos(ang); y=cy+radius*sin(ang)
        parts.append(f'<line x1="{cx}" y1="{cy}" x2="{x:.1f}" y2="{y:.1f}" stroke="{p.gold}" stroke-opacity="0.12"/>')
        parts.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="58" fill="#071018" fill-opacity="0.82" stroke="{p.cyan}" stroke-opacity="0.45"/>')
        parts.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="24" fill="{p.cyan}" fill-opacity="0.74" filter="url(#softGlow)"/>')
        parts.append(f'<text class="math" x="{x:.1f}" y="{y+105:.1f}" text-anchor="middle">{n}</text>')
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="132" fill="#0B0B10" stroke="{p.gold}" stroke-opacity="0.58"/>')
    parts.append(f'<text class="math" x="{cx}" y="{cy-8}" text-anchor="middle" fill="{p.gold}">SOURCE</text><text class="math" x="{cx}" y="{cy+46}" text-anchor="middle" fill="{p.gold}">ALPHABET</text>')
    # coefficient firewall boxes
    y0=1540
    parts.append(f'<text class="label" x="300" y="{y0-80}">COEFFICIENT FIREWALL</text>')
    for i,(num,expr) in enumerate(q["typed_coefficients"].items()):
        x=330+i*395
        parts.append(f'<rect x="{x}" y="{y0}" width="310" height="170" rx="28" fill="url(#obsidian)" stroke="{p.gold}" stroke-opacity="0.44"/>')
        parts.append(f'<text class="math" x="{x+155}" y="{y0+70}" text-anchor="middle" fill="{p.gold}">{num}</text>')
        parts.append(f'<text class="mathSmall" x="{x+155}" y="{y0+126}" text-anchor="middle">{esc(expr)}</text>')
    parts.append(f'<path d="M 260 1815 L 1540 1815" stroke="#050508" stroke-width="20" stroke-linecap="round"/>')
    parts.append(f'<text class="tiny" x="900" y="1870" text-anchor="middle">NO COEFFICIENT MAY MOVE JUST TO IMPROVE A FIT</text>')
    parts.append(svg_close("This board locks the formula alphabet and marks untyped coefficient motion as forbidden.", "EXACT CHECK: L=1/(8π) · SOURCE NUMBERS 3,4,7,27,56,70,72", style))
    Path(out_path).write_text("".join(parts), encoding="utf-8")


def render_scale_bridge(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p=PALETTE; w=style.width
    q=c.quantities
    parts=svg_open(c.title, "Logarithmic drop from the Planck stiffness seal to the electroweak scale.", style)
    parts.append(f'<text class="math" x="{w/2:.0f}" y="305" text-anchor="middle">v = MP exp[-12π + √3/2 + 2S + 148S²]</text>')
    # log shaft
    x=910; y_top=470; y_bot=1830
    parts.append(f'<line x1="{x}" y1="{y_top}" x2="{x}" y2="{y_bot}" stroke="url(#goldLine)" stroke-width="5" filter="url(#softGlow)"/>')
    for k,label in enumerate(["1e18", "1e14", "1e10", "1e6", "1e2"]):
        y=y_top+k*(y_bot-y_top)/4
        parts.append(f'<line x1="{x-105}" y1="{y}" x2="{x+105}" y2="{y}" stroke="{p.glass}" stroke-opacity="0.22"/>')
        parts.append(f'<text class="mathSmall" x="{x-150}" y="{y+8}" text-anchor="end">{label} GeV</text>')
    parts.append(f'<circle cx="{x}" cy="{y_top}" r="72" fill="{p.platinum}" fill-opacity="0.10" stroke="{p.platinum}" stroke-opacity="0.65" filter="url(#bigGlow)"/>')
    parts.append(f'<text class="math" x="{x}" y="{y_top-105}" text-anchor="middle">MP</text>')
    parts.append(f'<circle cx="{x}" cy="{y_bot}" r="68" fill="{p.gold}" fill-opacity="0.28" stroke="{p.gold}" stroke-opacity="0.88" filter="url(#bigGlow)"/>')
    parts.append(f'<text class="math" x="{x}" y="{y_bot+120}" text-anchor="middle" fill="{p.gold}">v = {q["v_GeV"]:.6f} GeV</text>')
    # term capsules
    terms=[("-12π", q["terms"]["three_action_quantum"], 250, 650), ("√3/2", q["terms"]["triadic_amplitude"], 1280, 780), ("2S", q["terms"]["wall_response"], 265, 1280), ("148S²", q["terms"]["augmented_chamber"], 1275, 1390)]
    for lab,val,tx,ty in terms:
        parts.append(f'<rect x="{tx-145}" y="{ty-72}" width="290" height="144" rx="34" fill="#071018" fill-opacity="0.78" stroke="{p.cyan if val>0 else p.warm}" stroke-opacity="0.38"/>')
        parts.append(f'<text class="math" x="{tx}" y="{ty-14}" text-anchor="middle">{esc(lab)}</text>')
        parts.append(f'<text class="tiny" x="{tx}" y="{ty+34}" text-anchor="middle">{val:+.9f}</text>')
        parts.append(f'<path d="M {tx} {ty+72} C {(tx+x)/2:.1f} {ty+140} {(tx+x)/2:.1f} {(ty+y_bot)/2:.1f} {x} {(ty+y_bot)/2:.1f}" fill="none" stroke="{p.gold}" stroke-opacity="0.13"/>')
    parts.append(f'<text class="mathSmall" x="900" y="2015" text-anchor="middle">v/MP = {q["ratio_v_over_Mp"]:.12e}</text>')
    parts.append(svg_close("MP is a physical filling seal; the exponent is the source-typed scale bridge.", "EXACT CHECK: LOG DROP USES EXPONENT %.12f" % q["exponent"], style))
    Path(out_path).write_text("".join(parts), encoding="utf-8")


def render_higgs(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p=PALETTE; w=style.width
    q=c.quantities
    parts=svg_open(c.title, "Formula chain from L,S to the Higgs quartic and tree mass output.", style)
    parts.append(f'<text class="math" x="{w/2:.0f}" y="305" text-anchor="middle">lambda ASHA = 3/8(1+L)(1/3−S)</text>')
    # Formula chain nodes
    xs=[300,620,930,1240,1530]; labels=["L", "S", "lambda", "v", "mH"]
    vals=[f"{L:.9f}", f"{S_SPLIT:.9f}", f"{q['lambda_ASHA']:.9f}", f"{q['v_GeV_from_scale_bridge']:.6f} GeV", f"{q['m_H_GeV']:.6f} GeV"]
    for i,(x,lab,val) in enumerate(zip(xs,labels,vals)):
        col=p.gold if lab in ["lambda","mH"] else p.cyan
        parts.append(f'<circle cx="{x}" cy="610" r="72" fill="#071018" stroke="{col}" stroke-opacity="0.55" filter="url(#softGlow)"/>')
        parts.append(f'<text class="math" x="{x}" y="600" text-anchor="middle" fill="{col}">{lab}</text>')
        parts.append(f'<text class="tiny" x="{x}" y="652" text-anchor="middle">{esc(val)}</text>')
        if i < len(xs)-1:
            parts.append(f'<path d="M {x+78} 610 L {xs[i+1]-82} 610" stroke="{p.platinum}" stroke-opacity="0.40" stroke-width="2" marker-end="url(#arrow)"/>')
    # Mexican hat wire curves: use path rings and potential curve
    cx=900; cy=1320
    parts.append(f'<ellipse cx="{cx}" cy="{cy+120}" rx="545" ry="170" fill="{p.gold}" fill-opacity="0.05" stroke="{p.gold}" stroke-opacity="0.18"/>')
    parts.append(f'<ellipse cx="{cx}" cy="{cy+10}" rx="315" ry="90" fill="none" stroke="{p.cyan}" stroke-opacity="0.38"/>')
    parts.append(f'<ellipse cx="{cx}" cy="{cy-60}" rx="130" ry="38" fill="#050508" stroke="{p.warm}" stroke-opacity="0.45"/>')
    # curve path y = scaled (r^2-r0^2)^2 cross-section
    pts=[]
    for i in range(180):
        t=-1.55 + 3.10*i/179
        y=(t*t-0.72)**2
        sx=cx+t*330
        sy=cy+190-y*260
        pts.append((sx,sy))
    d="M "+" L ".join(f"{x:.1f},{y:.1f}" for x,y in pts)
    parts.append(f'<path d="{d}" fill="none" stroke="{p.gold}" stroke-opacity="0.95" stroke-width="4" filter="url(#softGlow)"/>')
    parts.append(f'<text class="math" x="{cx}" y="1780" text-anchor="middle" fill="{p.gold}">mH = v √(2 lambda)</text>')
    parts.append(f'<text class="math" x="{cx}" y="1860" text-anchor="middle" fill="{p.gold}">{q["m_H_GeV"]:.6f} GeV</text>')
    parts.append(svg_close("Higgs quartic and mass are locked physical-filling formulas; this does not derive the scalar sector natively.", "EXACT CHECK: λ=%.12f  · mH=%.6f GeV" % (q["lambda_ASHA"], q["m_H_GeV"]), style))
    Path(out_path).write_text("".join(parts), encoding="utf-8")


def render_charged_leptons(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p=PALETTE; w=style.width
    q=c.quantities; masses=q["solved_masses_GeV"]
    parts=svg_open(c.title, "Tau anchor plus Koide and logarithmic spacing shape laws for charged leptons.", style)
    parts.append(f'<text class="math" x="{w/2:.0f}" y="305" text-anchor="middle">A tau → m tau, then Ke and De fix the shape</text>')
    # action stack
    terms=[("4π/3",4*pi/3), ("3/10",0.3), ("7/72",7/72), ("−S",-S_SPLIT), ("½(72+27)S²",0.5*(72+27)*S_SPLIT*S_SPLIT)]
    x0=230; y0=560
    acc=0
    for i,(lab,val) in enumerate(terms):
        h=abs(val)*220
        h=max(34,min(300,h))
        x=x0+i*265
        col=p.gold if val>0 else p.warm
        parts.append(f'<rect x="{x}" y="{y0+310-h}" width="155" height="{h}" rx="20" fill="{col}" fill-opacity="0.24" stroke="{col}" stroke-opacity="0.60"/>')
        parts.append(f'<text class="mathSmall" x="{x+78}" y="{y0+365}" text-anchor="middle">{esc(lab)}</text>')
        parts.append(f'<text class="tiny" x="{x+78}" y="{y0+405}" text-anchor="middle">{val:+.9f}</text>')
        acc += val
    parts.append(f'<text class="math" x="900" y="1045" text-anchor="middle" fill="{p.gold}">A tau = {q["A_tau"]:.9f}</text>')
    parts.append(f'<path d="M 900 1085 C 900 1160 900 1190 900 1265" stroke="{p.platinum}" stroke-opacity="0.44" stroke-width="2" marker-end="url(#arrow)"/>')
    # mass ladder log x positions
    y=1500; left=335; right=1470
    log_vals={k:log10(v) for k,v in masses.items()}
    mn=min(log_vals.values()); mx=max(log_vals.values())
    parts.append(f'<line x1="{left}" y1="{y}" x2="{right}" y2="{y}" stroke="url(#cyanLine)" stroke-width="3"/>')
    for k,v in masses.items():
        lx=log10(v); x=left+(lx-mn)/(mx-mn)*(right-left)
        col=p.gold if k=="m_tau" else p.cyan
        r=54 if k=="m_tau" else 42
        parts.append(f'<circle cx="{x:.1f}" cy="{y}" r="{r}" fill="#071018" stroke="{col}" stroke-opacity="0.72" filter="url(#softGlow)"/>')
        parts.append(f'<text class="math" x="{x:.1f}" y="{y-92}" text-anchor="middle" fill="{col}">{esc(k.replace("m_","m"))}</text>')
        parts.append(f'<text class="tiny" x="{x:.1f}" y="{y+105}" text-anchor="middle">{v:.12f} GeV</text>')
    # shape capsules
    parts.append(f'<rect x="310" y="1745" width="505" height="170" rx="32" fill="url(#obsidian)" stroke="{p.gold}" stroke-opacity="0.40"/>')
    parts.append(f'<text class="math" x="562" y="1812" text-anchor="middle" fill="{p.gold}">Ke = {q["K_e"]:.12f}</text>')
    parts.append(f'<text class="tiny" x="562" y="1862" text-anchor="middle">2/3 − 4(1−2L)S²</text>')
    parts.append(f'<rect x="985" y="1745" width="505" height="170" rx="32" fill="url(#obsidian)" stroke="{p.gold}" stroke-opacity="0.40"/>')
    parts.append(f'<text class="math" x="1238" y="1812" text-anchor="middle" fill="{p.gold}">De = {q["D_e"]:.12f}</text>')
    parts.append(f'<text class="tiny" x="1238" y="1862" text-anchor="middle">√(2π)+2S−4(1−L)S²</text>')
    parts.append(svg_close("Tau anchors the scale; Ke and De encode the shape. This remains physical filling, not a native flavor theorem.", "EXACT CHECK: me < mmu < mtau · LOG MASS LADDER", style))
    Path(out_path).write_text("".join(parts), encoding="utf-8")


def render_contract_svg(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    if c.figure_id.endswith("matter_sockets_product_depth"):
        render_matter_sockets(c, out_path, style)
    elif c.figure_id.endswith("locked_constants_source_alphabet"):
        render_source_alphabet(c, out_path, style)
    elif c.figure_id.endswith("planck_to_electroweak_scale_bridge"):
        render_scale_bridge(c, out_path, style)
    elif c.figure_id.endswith("higgs_sector_quartic_mass_chain"):
        render_higgs(c, out_path, style)
    elif c.figure_id.endswith("charged_lepton_anchor_shape_laws"):
        render_charged_leptons(c, out_path, style)
    else:
        raise ValueError(f"unknown contract: {c.figure_id}")
