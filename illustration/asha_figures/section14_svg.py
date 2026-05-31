from __future__ import annotations

from pathlib import Path
import html
from math import cos, sin, pi

from .style import PALETTE, STYLE, FigureStyle
from .readme_theorem_stack import FigureContract


def esc(s: object) -> str:
    return html.escape(str(s), quote=True)


def svg_open(title: str, desc: str, style: FigureStyle = STYLE) -> list[str]:
    p = PALETTE
    w, h = style.width, style.height
    return [f'''<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}" role="img" aria-labelledby="title desc">
<title id="title">{esc(title)}</title>
<desc id="desc">{esc(desc)}</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="41%" r="76%"><stop offset="0%" stop-color="#11111D"/><stop offset="62%" stop-color="{p.abyss}"/><stop offset="100%" stop-color="#010103"/></radialGradient>
  <radialGradient id="lawGlow" cx="50%" cy="50%" r="60%"><stop offset="0%" stop-color="{p.platinum}" stop-opacity="0.38"/><stop offset="42%" stop-color="{p.cyan}" stop-opacity="0.10"/><stop offset="100%" stop-color="{p.cyan}" stop-opacity="0"/></radialGradient>
  <linearGradient id="rail" x1="0" x2="1"><stop offset="0%" stop-color="{p.cyan}" stop-opacity="0.1"/><stop offset="45%" stop-color="{p.platinum}" stop-opacity="0.46"/><stop offset="100%" stop-color="{p.gold}" stop-opacity="0.24"/></linearGradient>
  <linearGradient id="glassCard" x1="0" x2="1"><stop offset="0%" stop-color="#071018" stop-opacity="0.88"/><stop offset="50%" stop-color="#11131B" stop-opacity="0.82"/><stop offset="100%" stop-color="#07070C" stop-opacity="0.92"/></linearGradient>
  <linearGradient id="obsidian" x1="0" x2="1"><stop offset="0%" stop-color="#050508"/><stop offset="48%" stop-color="#171721"/><stop offset="100%" stop-color="#050508"/></linearGradient>
  <filter id="softGlow" x="-180%" y="-180%" width="460%" height="460%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-250%" y="-250%" width="600%" height="600%"><feGaussianBlur stdDeviation="18" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="12" markerHeight="12" refX="9" refY="4" orient="auto"><path d="M0,0 L0,8 L10,4 z" fill="{p.platinum}" fill-opacity="0.72"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 62px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 34px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 25px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 2.6px; fill: {p.muted}; }}
    .tiny {{ font-family: {style.label_font}; font-size: 14px; letter-spacing: 1.25px; fill: #8B93A5; }}
    .caption {{ font-family: {style.label_font}; font-size: 18px; fill: #AAB3C2; }}
    .nodeText {{ font-family: {style.label_font}; font-size: 17px; letter-spacing: 0.35px; fill: #DCE6EE; }}
    .nodeTextSmall {{ font-family: {style.label_font}; font-size: 14px; letter-spacing: 0.55px; fill: #B8C5D0; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{w-96}" height="{h-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README · THEOREM RAIL · ACYCLIC SOURCE-TYPED STACK</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
<text class="mathSmall" x="{w/2:.0f}" y="278" text-anchor="middle">12 theorem/conditional-theorem nodes · 11 forward arrows · no physical filling promotion</text>
''']


def split_label(label: str) -> tuple[str, str]:
    # Manual splits for stable SVG layout.
    splits = {
        "Lorentzianized phase-space octave": ("Lorentzianized", "phase-space octave"),
        "Flat metric projection": ("Flat metric", "projection"),
        "Contact-seven source structure": ("Contact-seven", "source structure"),
        "Contact phase-triple projector algebra": ("Contact phase-triple", "projector algebra"),
        "Product-depth formal spectral extension": ("Product-depth", "formal spectral extension"),
        "Minimal commuting contact-depth matter extension": ("Minimal commuting", "contact-depth matter extension"),
        "Yukawa broadcast breaking by e^{-4πN_Q}": ("Yukawa broadcast breaking", "by e^{-4πN_Q}"),
        "Relative flavor orientation theorem": ("Relative flavor", "orientation theorem"),
        "Majorana Takagi and seesaw structure theorems": ("Majorana Takagi +", "seesaw structure theorems"),
        "Low-energy metric dynamics under standard assumptions": ("Low-energy metric dynamics", "under standard assumptions"),
        "Dimensional obstruction theorem": ("Dimensional", "obstruction theorem"),
        "Vacuum-zero independence theorem": ("Vacuum-zero", "independence theorem"),
    }
    return splits.get(label, (label, ""))


def draw_node(parts: list[str], x: float, y: float, idx: int, label: str, kind: str, p) -> None:
    w, h = 500, 112
    rx = 30
    color = p.gold if kind == "conditional" else p.cyan
    stroke = p.gold if kind == "conditional" else p.platinum
    parts.append(f'<rect x="{x-w/2}" y="{y-h/2}" width="{w}" height="{h}" rx="{rx}" fill="url(#glassCard)" stroke="{stroke}" stroke-opacity="0.40" filter="url(#softGlow)"/>')
    parts.append(f'<circle cx="{x-w/2+54}" cy="{y}" r="24" fill="{color}" fill-opacity="0.18" stroke="{color}" stroke-opacity="0.82"/>')
    parts.append(f'<text class="tiny" x="{x-w/2+54}" y="{y+5}" text-anchor="middle" fill="{color}">{idx:02d}</text>')
    a, b = split_label(label)
    parts.append(f'<text class="nodeText" x="{x-w/2+96}" y="{y-11}" fill="{p.platinum}">{esc(a)}</text>')
    if b:
        parts.append(f'<text class="nodeTextSmall" x="{x-w/2+96}" y="{y+22}">{esc(b)}</text>')
    if kind == "conditional":
        parts.append(f'<text class="tiny" x="{x+w/2-28}" y="{y+39}" text-anchor="end" fill="{p.gold}">CONDITIONAL STANDARD</text>')


def draw_arrow(parts: list[str], x1: float, y1: float, x2: float, y2: float, p) -> None:
    # route from lower edge of prior card to upper edge of next card
    parts.append(f'<path d="M {x1:.1f},{y1+58:.1f} C {x1:.1f},{(y1+y2)/2:.1f} {x2:.1f},{(y1+y2)/2:.1f} {x2:.1f},{y2-58:.1f}" stroke="url(#rail)" stroke-opacity="0.54" stroke-width="2.3" fill="none" marker-end="url(#arrow)"/>')


def draw_firewall(parts: list[str], p) -> None:
    parts.append(f'<rect x="130" y="2012" width="1540" height="190" rx="42" fill="url(#obsidian)" stroke="{p.obsidian}" stroke-opacity="0.96" stroke-width="11"/>')
    parts.append(f'<text class="label" x="190" y="2070" fill="{p.warm}">FIREWALL BELOW THEOREM STACK</text>')
    lines = [
        "theorem stack ≠ locked physical filling list",
        "no promotion to v/M_P, Higgs mass, CKM numerics, particle masses, or PMNS phases",
        "arrows are typed dependencies, not downstream-native value derivations",
    ]
    for i, line in enumerate(lines):
        yy = 2112 + i*34
        parts.append(f'<circle cx="210" cy="{yy-6}" r="5" fill="{p.warm}" fill-opacity="0.86"/>')
        parts.append(f'<text class="caption" x="232" y="{yy}">{esc(line)}</text>')


def render_theorem_stack(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p = PALETTE
    q = c.quantities
    parts = svg_open(c.title, "ASHA README theorem-level stack rendered as an exact twelve-node acyclic source-typed theorem rail.", style)

    # central background aura and two-column rail layout
    parts.append('<ellipse cx="900" cy="1115" rx="705" ry="835" fill="url(#lawGlow)" opacity="0.86"/>')
    parts.append(f'<line x1="900" y1="338" x2="900" y2="1910" stroke="{p.platinum}" stroke-opacity="0.075" stroke-width="1.5"/>')
    parts.append(f'<text class="label" x="170" y="372" fill="{p.cyan}">NATIVE / THEOREM LAW SPACE</text>')
    parts.append(f'<text class="label" x="1300" y="372" fill="{p.gold}">CONDITIONAL STANDARD LANE</text>')

    # 12 nodes in a precise snake/rail: 6 left, 6 right, always forward top-to-bottom.
    nodes = q["theorem_nodes"]
    coords: list[tuple[float, float]] = []
    ys = [440, 610, 780, 950, 1120, 1290, 1460, 1630, 1800]
    # make first eight alternating columns, last four centered/tight to avoid bottom firewall
    base = [
        (585, 440), (1215, 560), (585, 680), (1215, 800),
        (585, 920), (1215, 1040), (585, 1160), (1215, 1280),
        (585, 1400), (1215, 1520), (585, 1640), (1215, 1760),
    ]
    coords = base
    # arrows first so nodes sit above them
    for i in range(len(coords)-1):
        draw_arrow(parts, coords[i][0], coords[i][1], coords[i+1][0], coords[i+1][1], p)
    # thin acyclic envelope
    parts.append(f'<path d="M 585 440 C 1210 600 585 720 1215 800 C 585 940 1215 1040 585 1160 C 1215 1280 585 1400 1215 1520 C 585 1640 1215 1760 1215 1760" fill="none" stroke="{p.gold}" stroke-opacity="0.09" stroke-width="18" filter="url(#bigGlow)"/>')
    for i, (label, (x, y)) in enumerate(zip(nodes, coords), 1):
        kind = "conditional" if label == "Low-energy metric dynamics under standard assumptions" else "native"
        draw_node(parts, x, y, i, label, kind, p)

    # exactness badge
    parts.append(f'<rect x="560" y="1888" width="680" height="80" rx="28" fill="#050508" fill-opacity="0.70" stroke="{p.gold}" stroke-opacity="0.34"/>')
    parts.append(f'<text class="mathSmall" x="900" y="1938" text-anchor="middle" fill="{p.gold}">EXACT: 12 NODES · 11 FORWARD ARROWS · DAG</text>')

    draw_firewall(parts, p)
    parts.append(f'<text class="tiny" x="1580" y="2255" text-anchor="end">README ANCHOR: THEOREM-LEVEL STACK · EXACT BULLET LIST PRESERVED</text>')
    parts.append(f'<text class="tiny" x="1580" y="2288" text-anchor="end">BOUNDARY: THEOREM WHERE THEOREM EXISTS; PHYSICAL FILLING STAYS SEPARATE</text>')
    parts.append('</svg>')
    Path(out_path).write_text("".join(parts), encoding="utf-8")
