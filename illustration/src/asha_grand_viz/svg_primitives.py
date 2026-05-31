from __future__ import annotations

import html
from math import cos, sin, pi
from typing import Iterable, List, Tuple

Point = Tuple[float, float]

class SVG:
    def __init__(self, width: int, height: int):
        self.width = width
        self.height = height
        self.items: List[str] = []
        self.defs: List[str] = []

    def add_defs(self, s: str):
        self.defs.append(s)

    def add(self, s: str):
        self.items.append(s)

    def text(self, x, y, text, size=28, fill="#EDE7D1", family="Inter, DejaVu Sans, Arial", anchor="middle", weight="400", opacity=1.0, spacing=None):
        safe = html.escape(str(text))
        extra = f' letter-spacing="{spacing}"' if spacing else ""
        self.add(f'<text x="{x:.2f}" y="{y:.2f}" font-family="{family}" font-size="{size}" font-weight="{weight}" text-anchor="{anchor}" fill="{fill}" opacity="{opacity}"{extra}>{safe}</text>')

    def line(self, x1, y1, x2, y2, stroke="#FFF", width=1, opacity=1, dash=None, marker=None, cls=None, filt=None):
        dash_s = f' stroke-dasharray="{dash}"' if dash else ""
        marker_s = f' marker-end="url(#{marker})"' if marker else ""
        class_s = f' class="{cls}"' if cls else ""
        filter_s = f' filter="url(#{filt})"' if filt else ""
        self.add(f'<line x1="{x1:.2f}" y1="{y1:.2f}" x2="{x2:.2f}" y2="{y2:.2f}" stroke="{stroke}" stroke-width="{width}" opacity="{opacity}" fill="none"{dash_s}{marker_s}{class_s}{filter_s}/>')

    def path(self, d, stroke="#FFF", width=1, fill="none", opacity=1, dash=None, marker=None, filt=None):
        dash_s = f' stroke-dasharray="{dash}"' if dash else ""
        marker_s = f' marker-end="url(#{marker})"' if marker else ""
        filter_s = f' filter="url(#{filt})"' if filt else ""
        self.add(f'<path d="{d}" stroke="{stroke}" stroke-width="{width}" fill="{fill}" opacity="{opacity}" stroke-linecap="round" stroke-linejoin="round"{dash_s}{marker_s}{filter_s}/>')

    def circle(self, x, y, r, fill="#FFF", stroke="none", width=1, opacity=1, filt=None):
        filter_s = f' filter="url(#{filt})"' if filt else ""
        stroke_s = f' stroke="{stroke}" stroke-width="{width}"' if stroke != "none" else ""
        self.add(f'<circle cx="{x:.2f}" cy="{y:.2f}" r="{r:.2f}" fill="{fill}"{stroke_s} opacity="{opacity}"{filter_s}/>')

    def ellipse(self, x, y, rx, ry, fill="none", stroke="#FFF", width=1, opacity=1, rotate=0, dash=None, filt=None):
        dash_s = f' stroke-dasharray="{dash}"' if dash else ""
        filter_s = f' filter="url(#{filt})"' if filt else ""
        self.add(f'<ellipse cx="{x:.2f}" cy="{y:.2f}" rx="{rx:.2f}" ry="{ry:.2f}" fill="{fill}" stroke="{stroke}" stroke-width="{width}" opacity="{opacity}" transform="rotate({rotate:.2f} {x:.2f} {y:.2f})"{dash_s}{filter_s}/>')

    def rect(self, x, y, w, h, fill="#000", stroke="none", width=1, rx=0, opacity=1, filt=None):
        filter_s = f' filter="url(#{filt})"' if filt else ""
        stroke_s = f' stroke="{stroke}" stroke-width="{width}"' if stroke != "none" else ""
        self.add(f'<rect x="{x:.2f}" y="{y:.2f}" width="{w:.2f}" height="{h:.2f}" rx="{rx:.2f}" fill="{fill}"{stroke_s} opacity="{opacity}"{filter_s}/>')

    def polygon(self, pts: Iterable[Point], fill="none", stroke="#FFF", width=1, opacity=1, dash=None, filt=None):
        pts_s = " ".join(f"{x:.2f},{y:.2f}" for x, y in pts)
        dash_s = f' stroke-dasharray="{dash}"' if dash else ""
        filter_s = f' filter="url(#{filt})"' if filt else ""
        self.add(f'<polygon points="{pts_s}" fill="{fill}" stroke="{stroke}" stroke-width="{width}" opacity="{opacity}" stroke-linejoin="round"{dash_s}{filter_s}/>')

    def polyline(self, pts: Iterable[Point], stroke="#FFF", width=1, opacity=1, fill="none", dash=None, marker=None, filt=None):
        pts_s = " ".join(f"{x:.2f},{y:.2f}" for x, y in pts)
        dash_s = f' stroke-dasharray="{dash}"' if dash else ""
        marker_s = f' marker-end="url(#{marker})"' if marker else ""
        filter_s = f' filter="url(#{filt})"' if filt else ""
        self.add(f'<polyline points="{pts_s}" stroke="{stroke}" stroke-width="{width}" opacity="{opacity}" fill="{fill}" stroke-linecap="round" stroke-linejoin="round"{dash_s}{marker_s}{filter_s}/>')

    def group(self, content: str, opacity=1.0, filt=None, transform=None):
        f = f' filter="url(#{filt})"' if filt else ""
        t = f' transform="{transform}"' if transform else ""
        self.add(f'<g opacity="{opacity}"{f}{t}>{content}</g>')

    def render(self) -> str:
        defs = "\n".join(self.defs)
        body = "\n".join(self.items)
        return f'''<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="{self.width}" height="{self.height}" viewBox="0 0 {self.width} {self.height}">
<defs>
{defs}
</defs>
{body}
</svg>
'''


def polar(cx: float, cy: float, r: float, angle: float) -> Point:
    return (cx + r * cos(angle), cy + r * sin(angle))


def regular_polygon(cx: float, cy: float, r: float, n: int, start: float = -pi/2) -> List[Point]:
    return [polar(cx, cy, r, start + 2*pi*i/n) for i in range(n)]


def bezier_path(points: List[Point]) -> str:
    if not points:
        return ""
    d = f"M {points[0][0]:.2f} {points[0][1]:.2f}"
    for p in points[1:]:
        d += f" L {p[0]:.2f} {p[1]:.2f}"
    return d
