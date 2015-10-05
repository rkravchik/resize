/*
Copyright (c) 2014, Charlie Vieth <charlie.vieth@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package resize

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type ycc struct {
	// Pix holds the image's pixels, in Y, Cb, Cr order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
	// SubsampleRatio is the subsample ratio of the original YCbCr image.
	SubsampleRatio image.YCbCrSubsampleRatio
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *ycc) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *ycc) Bounds() image.Rectangle {
	return p.Rect
}

func (p *ycc) ColorModel() color.Model {
	return color.YCbCrModel
}

func (p *ycc) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color.YCbCr{}
	}
	i := p.PixOffset(x, y)
	return color.YCbCr{
		p.Pix[i+0],
		p.Pix[i+1],
		p.Pix[i+2],
	}
}

func (p *ycc) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *ycc) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	if r.Empty() {
		return &ycc{SubsampleRatio: p.SubsampleRatio}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &ycc{
		Pix:            p.Pix[i:],
		Stride:         p.Stride,
		Rect:           r,
		SubsampleRatio: p.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func newYCC(r image.Rectangle, s image.YCbCrSubsampleRatio) *ycc {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 3*w*h)
	return &ycc{Pix: buf, Stride: 3 * w, Rect: r, SubsampleRatio: s}
}
