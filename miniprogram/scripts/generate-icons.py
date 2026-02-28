#!/usr/bin/env python3
"""
Generate simple PNG icon assets for the DANA Mini Program.
Uses only Python standard library (struct, zlib) - no external dependencies.
"""
import struct
import zlib
import os
import math

OUTPUT_DIR = os.path.join(os.path.dirname(__file__), '..', 'assets', 'images')


def create_png(width, height, pixels):
    """Create a PNG file from RGBA pixel data."""
    def chunk(chunk_type, data):
        c = chunk_type + data
        crc = struct.pack('>I', zlib.crc32(c) & 0xFFFFFFFF)
        return struct.pack('>I', len(data)) + c + crc

    header = b'\x89PNG\r\n\x1a\n'
    ihdr = chunk(b'IHDR', struct.pack('>IIBBBBB', width, height, 8, 6, 0, 0, 0))

    raw = b''
    for y in range(height):
        raw += b'\x00'
        for x in range(width):
            idx = (y * width + x) * 4
            raw += bytes(pixels[idx:idx+4])

    idat = chunk(b'IDAT', zlib.compress(raw))
    iend = chunk(b'IEND', b'')

    return header + ihdr + idat + iend


def new_canvas(w, h):
    return [0] * (w * h * 4)


def set_pixel(pixels, w, x, y, r, g, b, a=255):
    if 0 <= x < w and 0 <= y < w:
        idx = (y * w + x) * 4
        # Alpha blending
        if a < 255 and pixels[idx+3] > 0:
            old_a = pixels[idx+3] / 255.0
            new_a = a / 255.0
            out_a = new_a + old_a * (1 - new_a)
            if out_a > 0:
                pixels[idx] = int((r * new_a + pixels[idx] * old_a * (1 - new_a)) / out_a)
                pixels[idx+1] = int((g * new_a + pixels[idx+1] * old_a * (1 - new_a)) / out_a)
                pixels[idx+2] = int((b * new_a + pixels[idx+2] * old_a * (1 - new_a)) / out_a)
                pixels[idx+3] = int(out_a * 255)
        else:
            pixels[idx] = r
            pixels[idx+1] = g
            pixels[idx+2] = b
            pixels[idx+3] = a


def draw_circle(pixels, w, cx, cy, radius, r, g, b, a=255, fill=False, thickness=3):
    for y in range(w):
        for x in range(w):
            dist = math.sqrt((x - cx)**2 + (y - cy)**2)
            if fill:
                if dist <= radius:
                    set_pixel(pixels, w, x, y, r, g, b, a)
            else:
                if abs(dist - radius) < thickness / 2:
                    alpha = max(0, min(255, int(a * (1 - abs(dist - radius) / (thickness / 2)))))
                    set_pixel(pixels, w, x, y, r, g, b, alpha)


def draw_rect(pixels, w, x1, y1, x2, y2, r, g, b, a=255):
    for y in range(max(0, y1), min(w, y2)):
        for x in range(max(0, x1), min(w, x2)):
            set_pixel(pixels, w, x, y, r, g, b, a)


def draw_line(pixels, w, x1, y1, x2, y2, r, g, b, a=255, thickness=3):
    dx = x2 - x1
    dy = y2 - y1
    length = max(1, int(math.sqrt(dx*dx + dy*dy)))
    for i in range(length + 1):
        t = i / length
        cx = x1 + dx * t
        cy = y1 + dy * t
        for ox in range(-thickness//2, thickness//2 + 1):
            for oy in range(-thickness//2, thickness//2 + 1):
                if ox*ox + oy*oy <= (thickness/2)**2:
                    set_pixel(pixels, w, int(cx+ox), int(cy+oy), r, g, b, a)


def hex_to_rgb(hex_color):
    hex_color = hex_color.lstrip('#')
    return tuple(int(hex_color[i:i+2], 16) for i in (0, 2, 4))


# ---- Icon Generators ----

def gen_about_icon(size, color):
    """Info circle icon (i)"""
    pixels = new_canvas(size, size)
    r, g, b = hex_to_rgb(color)
    cx, cy = size // 2, size // 2
    radius = int(size * 0.38)
    draw_circle(pixels, size, cx, cy, radius, r, g, b, thickness=3)
    # Draw 'i' - dot
    draw_circle(pixels, size, cx, cy - radius//3, 2, r, g, b, fill=True)
    # Draw 'i' - line
    draw_rect(pixels, size, cx-2, cy - radius//6, cx+2, cy + radius//2, r, g, b)
    return create_png(size, size, pixels)


def gen_listing_icon(size, color):
    """House icon"""
    pixels = new_canvas(size, size)
    r, g, b = hex_to_rgb(color)
    cx = size // 2

    # Roof (triangle)
    roof_top = int(size * 0.15)
    roof_bottom = int(size * 0.48)
    roof_left = int(size * 0.12)
    roof_right = int(size * 0.88)
    for y in range(roof_top, roof_bottom):
        t = (y - roof_top) / max(1, (roof_bottom - roof_top))
        half_w = int(t * (roof_right - cx))
        for x in range(cx - half_w, cx + half_w):
            d_edge = min(abs(x - (cx - half_w)), abs(x - (cx + half_w)))
            d_top = y - roof_top
            if d_edge < 3 or d_top < 3 or abs(y - roof_bottom) < 1:
                set_pixel(pixels, size, x, y, r, g, b)

    # Body
    body_top = roof_bottom - 2
    body_bottom = int(size * 0.85)
    body_left = int(size * 0.2)
    body_right = int(size * 0.8)
    for y in range(body_top, body_bottom):
        for x in range(body_left, body_right):
            if x < body_left + 3 or x >= body_right - 3 or y >= body_bottom - 3:
                set_pixel(pixels, size, x, y, r, g, b)

    # Door
    door_w = int(size * 0.16)
    door_h = int(size * 0.22)
    door_left = cx - door_w // 2
    draw_rect(pixels, size, door_left, body_bottom - door_h, door_left + door_w, body_bottom, r, g, b)

    return create_png(size, size, pixels)


def gen_search_icon(size, color):
    """Magnifying glass"""
    pixels = new_canvas(size, size)
    r, g, b = hex_to_rgb(color)
    cx, cy = int(size * 0.4), int(size * 0.4)
    radius = int(size * 0.25)
    draw_circle(pixels, size, cx, cy, radius, r, g, b, thickness=4)
    # Handle
    hx = cx + int(radius * 0.7)
    hy = cy + int(radius * 0.7)
    draw_line(pixels, size, hx, hy, int(size*0.82), int(size*0.82), r, g, b, thickness=4)
    return create_png(size, size, pixels)


def gen_grid_icon(size, color):
    """4 squares grid"""
    pixels = new_canvas(size, size)
    r, g, b = hex_to_rgb(color)
    margin = int(size * 0.15)
    gap = int(size * 0.1)
    sq = (size - 2 * margin - gap) // 2
    draw_rect(pixels, size, margin, margin, margin+sq, margin+sq, r, g, b)
    draw_rect(pixels, size, margin+sq+gap, margin, margin+2*sq+gap, margin+sq, r, g, b)
    draw_rect(pixels, size, margin, margin+sq+gap, margin+sq, margin+2*sq+gap, r, g, b)
    draw_rect(pixels, size, margin+sq+gap, margin+sq+gap, margin+2*sq+gap, margin+2*sq+gap, r, g, b)
    return create_png(size, size, pixels)


def gen_list_icon(size, color):
    """Horizontal lines"""
    pixels = new_canvas(size, size)
    r, g, b = hex_to_rgb(color)
    margin = int(size * 0.15)
    bar_h = int(size * 0.1)
    gap = int(size * 0.08)
    w = size - 2 * margin
    for i in range(4):
        y = margin + i * (bar_h + gap)
        draw_rect(pixels, size, margin, y, margin + w, y + bar_h, r, g, b)
    return create_png(size, size, pixels)


def gen_empty_icon(size, color):
    """Empty box with question mark"""
    pixels = new_canvas(size, size)
    r, g, b = hex_to_rgb(color)
    cx = size // 2

    # Box
    bx1, by1 = int(size*0.2), int(size*0.35)
    bx2, by2 = int(size*0.8), int(size*0.8)
    for y in range(by1, by2):
        for x in range(bx1, bx2):
            if x < bx1+3 or x >= bx2-3 or y < by1+3 or y >= by2-3:
                set_pixel(pixels, size, x, y, r, g, b)

    # Question mark
    qcx, qcy = cx, int(size * 0.6)
    draw_circle(pixels, size, qcx, qcy - 8, 8, r, g, b, thickness=3)
    draw_rect(pixels, size, qcx-2, qcy, qcx+2, qcy+6, r, g, b)
    draw_circle(pixels, size, qcx, qcy+10, 2, r, g, b, fill=True)

    return create_png(size, size, pixels)


def gen_logo(size):
    """Blue circle with white house"""
    pixels = new_canvas(size, size)
    cx, cy = size // 2, size // 2
    radius = int(size * 0.45)

    # Blue circle background
    draw_circle(pixels, size, cx, cy, radius, 22, 119, 255, fill=True)

    # White house
    # Roof
    roof_top = int(size * 0.22)
    roof_bottom = int(size * 0.48)
    for y in range(roof_top, roof_bottom):
        t = (y - roof_top) / max(1, (roof_bottom - roof_top))
        half_w = int(t * size * 0.28)
        draw_rect(pixels, size, cx - half_w, y, cx + half_w, y + 1, 255, 255, 255)

    # Body
    body_left = int(size * 0.3)
    body_right = int(size * 0.7)
    body_top = roof_bottom - 2
    body_bottom = int(size * 0.72)
    draw_rect(pixels, size, body_left, body_top, body_right, body_bottom, 255, 255, 255)

    # Door (blue cutout)
    door_w = int(size * 0.12)
    door_h = int(size * 0.18)
    draw_rect(pixels, size, cx - door_w//2, body_bottom - door_h, cx + door_w//2, body_bottom, 22, 119, 255)

    return create_png(size, size, pixels)


# ---- Main ----

def main():
    os.makedirs(OUTPUT_DIR, exist_ok=True)

    ICON_SIZE = 81
    INACTIVE = '#999999'
    ACTIVE = '#1677FF'

    icons = {
        'icon-about.png': gen_about_icon(ICON_SIZE, INACTIVE),
        'icon-about-active.png': gen_about_icon(ICON_SIZE, ACTIVE),
        'icon-listing.png': gen_listing_icon(ICON_SIZE, INACTIVE),
        'icon-listing-active.png': gen_listing_icon(ICON_SIZE, ACTIVE),
        'icon-search.png': gen_search_icon(ICON_SIZE, '#999999'),
        'icon-grid.png': gen_grid_icon(ICON_SIZE, '#666666'),
        'icon-list.png': gen_list_icon(ICON_SIZE, '#666666'),
        'icon-empty.png': gen_empty_icon(160, '#CCCCCC'),
        'logo.png': gen_logo(200),
    }

    for filename, data in icons.items():
        filepath = os.path.join(OUTPUT_DIR, filename)
        with open(filepath, 'wb') as f:
            f.write(data)
        print(f'Created: {filename} ({len(data)} bytes)')

    print('\nAll icons generated successfully!')


if __name__ == '__main__':
    main()
