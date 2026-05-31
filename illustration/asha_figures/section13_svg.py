from __future__ import annotations

from pathlib import Path
import html
from math import cos, sin, pi

from .style import PALETTE, STYLE, FigureStyle
from .readme_low_energy_action import FigureContract


def esc(s: object) -> str:
    return html.escape(str(s), quote=True)


def svg_open(title: str, desc: str, style: FigureStyle = STYLE) -> list[str]:
    p = PALETTE
    w, h = style.width, style.height
    return [f'''<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}" role="img" aria-labelledby="title desc">
<title id="title">{esc(title)}</title>
<desc id="desc">{esc(desc)}</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="74%"><stop offset="0%" stop-color="#10101B"/><stop offset="62%" stop-color="{p.abyss}"/><stop offset="100%" stop-color="#010103"/></radialGradient>
  <radialGradient id="coreGlow" cx="50%" cy="50%" r="58%"><stop offset="0%" stop-color="{p.gold}" stop-opacity="0.42"/><stop offset="46%" stop-color="{p.cyan}" stop-opacity="0.10"/><stop offset="100%" stop-color="{p.cyan}" stop-opacity="0"/></radialGradient>
  <linearGradient id="obsidian" x1="0" x2="1"><stop offset="0%" stop-color="#050508"/><stop offset="50%" stop-color="#15151E"/><stop offset="100%" stop-color="#050508"/></linearGradient>
  <linearGradient id="equationGlass" x1="0" x2="1"><stop offset="0%" stop-color="#0B1118" stop-opacity="0.92"/><stop offset="52%" stop-color="#13131B" stop-opacity="0.86"/><stop offset="100%" stop-color="#08080D" stop-opacity="0.94"/></linearGradient>
  <filter id="softGlow" x="-180%" y="-180%" width="460%" height="460%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-240%" y="-240%" width="580%" height="580%"><feGaussianBlur stdDeviation="16" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="10" markerHeight="10" refX="7" refY="3" orient="auto"><path d="M0,0 L0,6 L8,3 z" fill="{p.platinum}" fill-opacity="0.78"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 58px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 34px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 25px; fill: {p.platinum}; }}
    .mathTiny {{ font-family: {style.title_font}; font-size: 21px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 2.5px; fill: {p.muted}; }}
    .tiny {{ font-family: {style.label_font}; font-size: 14px; letter-spacing: 1.35px; fill: #8B93A5; }}
    .caption {{ font-family: {style.label_font}; font-size: 19px; fill: #AAB3C2; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{w-96}" height="{h-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README · LOW-ENERGY ACTION · SOURCE-TYPED SKELETON</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
<text class="mathSmall" x="{w/2:.0f}" y="278" text-anchor="middle">six-sector action architecture · ASHA Higgs/Yukawa inserts · inherited firewalls</text>
''']


def sector_card(parts: list[str], cx: float, cy: float, title: str, subtitle: str, color: str, index: int, angle: float, p) -> None:
    width, height = 430, 168
    x, y = cx - width/2, cy - height/2
    parts.append(f'<rect x="{x}" y="{y}" width="{width}" height="{height}" rx="34" fill="#050508" fill-opacity="0.74" stroke="{color}" stroke-opacity="0.46" filter="url(#softGlow)"/>')
    parts.append(f'<circle cx="{x+45}" cy="{y+45}" r="20" fill="{color}" fill-opacity="0.25" stroke="{color}" stroke-opacity="0.75"/>')
    parts.append(f'<text class="tiny" x="{x+45}" y="{y+51}" text-anchor="middle" fill="{color}">{index}</text>')
    parts.append(f'<text class="mathSmall" x="{x+82}" y="{y+54}" fill="{color}">{esc(title)}</text>')
    # two lines manually split by semicolon if present
    if ";" in subtitle:
        a, b = subtitle.split(";", 1)
        parts.append(f'<text class="tiny" x="{x+42}" y="{y+108}">{esc(a.strip())}</text>')
        parts.append(f'<text class="tiny" x="{x+42}" y="{y+136}">{esc(b.strip())}</text>')
    else:
        parts.append(f'<text class="tiny" x="{x+42}" y="{y+118}">{esc(subtitle)}</text>')
    # connector from card to central core
    rx = cx - 900
    ry = cy - 860
    c1x = 900 + 0.48*rx
    c1y = 860 + 0.48*ry
    parts.append(f'<path d="M {900:.1f} {860:.1f} C {c1x:.1f} {c1y:.1f} {cx:.1f} {cy:.1f} {cx - 0.43*width*cos(angle):.1f} {cy - 0.43*height*sin(angle):.1f}" stroke="{color}" stroke-opacity="0.26" stroke-width="2.5" fill="none" marker-end="url(#arrow)"/>')


def draw_core(parts: list[str], p) -> None:
    cx, cy = 900, 860
    parts.append(f'<ellipse cx="{cx}" cy="{cy}" rx="560" ry="400" fill="url(#coreGlow)" opacity="0.95"/>')
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="190" fill="#050508" fill-opacity="0.80" stroke="{p.gold}" stroke-opacity="0.55" stroke-width="2" filter="url(#bigGlow)"/>')
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="132" fill="none" stroke="{p.cyan}" stroke-opacity="0.20" stroke-width="1.5"/>')
    parts.append(f'<text class="math" x="{cx}" y="{cy-34}" text-anchor="middle" fill="{p.gold}">S_ASHA^low</text>')
    parts.append(f'<text class="mathTiny" x="{cx}" y="{cy+22}" text-anchor="middle">= S_grav + S_gauge + S_Higgs</text>')
    parts.append(f'<text class="mathTiny" x="{cx}" y="{cy+62}" text-anchor="middle">+ S_fermion + S_Yukawa + S_nu</text>')
    parts.append(f'<text class="tiny" x="{cx}" y="{cy+116}" text-anchor="middle">README skeleton: assembled low-energy effective action</text>')


def equation_panel(parts: list[str], x: float, y: float, title: str, lines: list[str], color: str, p, height: float = 245) -> None:
    parts.append(f'<rect x="{x}" y="{y}" width="700" height="{height}" rx="38" fill="url(#equationGlass)" stroke="{color}" stroke-opacity="0.43" filter="url(#softGlow)"/>')
    parts.append(f'<text class="label" x="{x+42}" y="{y+55}" fill="{color}">{esc(title)}</text>')
    yy = y + 103
    for line in lines:
        parts.append(f'<text class="mathTiny" x="{x+42}" y="{yy}" fill="{p.platinum}">{esc(line)}</text>')
        yy += 38


def firewall_strip(parts: list[str], p) -> None:
    parts.append(f'<rect x="130" y="2010" width="1540" height="188" rx="42" fill="url(#obsidian)" stroke="{p.obsidian}" stroke-opacity="0.96" stroke-width="11"/>')
    parts.append(f'<text class="label" x="190" y="2068" fill="{p.warm}">FIREWALLS PRESERVED</text>')
    wounds = [
        "M_P^2 and Lambda remain inherited source-typed wounds",
        "Y_f^ASHA is shape-law/bridge/seal data, not a full native flavor theorem",
        "PMNS CP and Majorana selector remain unresolved",
    ]
    for i, wound in enumerate(wounds):
        y = 2110 + i*34
        parts.append(f'<circle cx="210" cy="{y-6}" r="5" fill="{p.warm}" fill-opacity="0.85"/>')
        parts.append(f'<text class="caption" x="232" y="{y}">{esc(wound)}</text>')


def render_low_energy_action(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p = PALETTE
    q = c.quantities
    parts = svg_open(c.title, "Six-sector current low-energy ASHA action with exact Higgs/Yukawa forms and inherited source-type firewalls.", style)

    # orbital source architecture
    draw_core(parts, p)
    sectors = [
        ("S_grav", "conditional metric lane; MP and Lambda not native", p.gold),
        ("S_gauge", "standard gauge kinetic lane; finite/internal data installed", p.cyan),
        ("S_Higgs^ASHA", "locked ASHA quartic/proxy lane; coefficient carried below", p.gold),
        ("S_fermion", "standard fermion kinetic lane; low-energy metric background", p.platinum),
        ("S_Yukawa^ASHA", "ASHA shape laws; sealed flavor/orientation data", p.frost),
        ("S_nu^seesaw", "rank-2 normal seesaw bridge; PMNS phases unresolved", p.warm),
    ]
    orbit_cx, orbit_cy = 900, 860
    rx, ry = 590, 430
    for i, (name, subtitle, color) in enumerate(sectors, 1):
        # start at top, clockwise; leave room for equation panels below
        angle = -pi/2 + (i-1) * 2*pi/6
        cx = orbit_cx + rx * cos(angle)
        cy = orbit_cy + ry * sin(angle)
        sector_card(parts, cx, cy, name, subtitle, color, i, angle, p)

    # exact interior equations from README
    equation_panel(
        parts,
        130,
        1390,
        "ASHA HIGGS INSERT",
        [
            "S_Higgs^ASHA = ∫√(-g)[ kinetic − λ_A(|φ|²−v²/2)² ] d⁴x",
            "λ_A = ⅜(1+L)(⅓−S)",
            "center = |φ|² − v²/2",
        ],
        p.gold,
        p,
        255,
    )
    equation_panel(
        parts,
        970,
        1390,
        "ASHA YUKAWA INSERT",
        [
            "S_Yukawa^ASHA = −∫√(-g) Σ_f[ ψ̄_L Y_f^ASHA φ_f ψ_R + h.c. ] d⁴x",
            "Y_f^ASHA = shape-law ledger + bridge/seal orientation data",
            "scope = low-energy skeleton, not complete native flavor theorem",
        ],
        p.frost,
        p,
        255,
    )

    # low-energy sum ribbon
    parts.append(f'<rect x="215" y="1718" width="1370" height="190" rx="42" fill="#050508" fill-opacity="0.62" stroke="{p.platinum}" stroke-opacity="0.30"/>')
    parts.append(f'<text class="label" x="900" y="1778" text-anchor="middle">EXACT README SUM</text>')
    parts.append(f'<text class="mathSmall" x="900" y="1830" text-anchor="middle">S_ASHA^low = S_grav + S_gauge + S_Higgs^ASHA</text>')
    parts.append(f'<text class="mathSmall" x="900" y="1870" text-anchor="middle">+ S_fermion + S_Yukawa^ASHA + S_nu^seesaw</text>')
    parts.append(f'<text class="tiny" x="900" y="1908" text-anchor="middle">six terms · two ASHA-specific interior equations · remaining wounds quarantined below</text>')

    firewall_strip(parts, p)

    parts.append(f'<text class="tiny" x="1580" y="2255" text-anchor="end">EXACT CHECK: 6 SECTORS · HIGGS COEFFICIENT (3/8)(1+L)(1/3-S) · YUKAWA SUM_F TERM</text>')
    parts.append(f'<text class="tiny" x="1580" y="2288" text-anchor="end">BOUNDARY: LOW-ENERGY SKELETON, NOT A SINGLE FULL NATIVE DERIVATION</text>')
    parts.append('</svg>')
    Path(out_path).write_text("".join(parts), encoding="utf-8")
