// +build !go1.5

package resize

import "image"

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (p *ycc) YCbCr() *image.YCbCr {
	ycbcr := image.NewYCbCr(p.Rect, p.SubsampleRatio)
	var off int

	switch ycbcr.SubsampleRatio {
	case image.YCbCrSubsampleRatio422:
		for y := ycbcr.Rect.Min.Y; y < ycbcr.Rect.Max.Y; y++ {
			yy := (y - ycbcr.Rect.Min.Y) * ycbcr.YStride
			cy := (y - ycbcr.Rect.Min.Y) * ycbcr.CStride
			for x := ycbcr.Rect.Min.X; x < ycbcr.Rect.Max.X; x++ {
				xx := (x - ycbcr.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx/2
				ycbcr.Y[yi] = p.Pix[off+0]
				ycbcr.Cb[ci] = p.Pix[off+1]
				ycbcr.Cr[ci] = p.Pix[off+2]
				off += 3
			}
		}
	case image.YCbCrSubsampleRatio420:
		for y := ycbcr.Rect.Min.Y; y < ycbcr.Rect.Max.Y; y++ {
			yy := (y - ycbcr.Rect.Min.Y) * ycbcr.YStride
			cy := (y/2 - ycbcr.Rect.Min.Y/2) * ycbcr.CStride
			for x := ycbcr.Rect.Min.X; x < ycbcr.Rect.Max.X; x++ {
				xx := (x - ycbcr.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx/2
				ycbcr.Y[yi] = p.Pix[off+0]
				ycbcr.Cb[ci] = p.Pix[off+1]
				ycbcr.Cr[ci] = p.Pix[off+2]
				off += 3
			}
		}
	case image.YCbCrSubsampleRatio440:
		for y := ycbcr.Rect.Min.Y; y < ycbcr.Rect.Max.Y; y++ {
			yy := (y - ycbcr.Rect.Min.Y) * ycbcr.YStride
			cy := (y/2 - ycbcr.Rect.Min.Y/2) * ycbcr.CStride
			for x := ycbcr.Rect.Min.X; x < ycbcr.Rect.Max.X; x++ {
				xx := (x - ycbcr.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx
				ycbcr.Y[yi] = p.Pix[off+0]
				ycbcr.Cb[ci] = p.Pix[off+1]
				ycbcr.Cr[ci] = p.Pix[off+2]
				off += 3
			}
		}
	default:
		// Default to 4:4:4 subsampling.
		for y := ycbcr.Rect.Min.Y; y < ycbcr.Rect.Max.Y; y++ {
			yy := (y - ycbcr.Rect.Min.Y) * ycbcr.YStride
			cy := (y - ycbcr.Rect.Min.Y) * ycbcr.CStride
			for x := ycbcr.Rect.Min.X; x < ycbcr.Rect.Max.X; x++ {
				xx := (x - ycbcr.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx
				ycbcr.Y[yi] = p.Pix[off+0]
				ycbcr.Cb[ci] = p.Pix[off+1]
				ycbcr.Cr[ci] = p.Pix[off+2]
				off += 3
			}
		}
	}
	return ycbcr
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func imageYCbCrToYCC(in *image.YCbCr) *ycc {
	w, h := in.Rect.Dx(), in.Rect.Dy()
	r := image.Rect(0, 0, w, h)
	buf := make([]uint8, 3*w*h)
	p := ycc{Pix: buf, Stride: 3 * w, Rect: r, SubsampleRatio: in.SubsampleRatio}
	var off int

	switch in.SubsampleRatio {
	case image.YCbCrSubsampleRatio422:
		for y := in.Rect.Min.Y; y < in.Rect.Max.Y; y++ {
			yy := (y - in.Rect.Min.Y) * in.YStride
			cy := (y - in.Rect.Min.Y) * in.CStride
			for x := in.Rect.Min.X; x < in.Rect.Max.X; x++ {
				xx := (x - in.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx/2
				p.Pix[off+0] = in.Y[yi]
				p.Pix[off+1] = in.Cb[ci]
				p.Pix[off+2] = in.Cr[ci]
				off += 3
			}
		}
	case image.YCbCrSubsampleRatio420:
		for y := in.Rect.Min.Y; y < in.Rect.Max.Y; y++ {
			yy := (y - in.Rect.Min.Y) * in.YStride
			cy := (y/2 - in.Rect.Min.Y/2) * in.CStride
			for x := in.Rect.Min.X; x < in.Rect.Max.X; x++ {
				xx := (x - in.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx/2
				p.Pix[off+0] = in.Y[yi]
				p.Pix[off+1] = in.Cb[ci]
				p.Pix[off+2] = in.Cr[ci]
				off += 3
			}
		}
	case image.YCbCrSubsampleRatio440:
		for y := in.Rect.Min.Y; y < in.Rect.Max.Y; y++ {
			yy := (y - in.Rect.Min.Y) * in.YStride
			cy := (y/2 - in.Rect.Min.Y/2) * in.CStride
			for x := in.Rect.Min.X; x < in.Rect.Max.X; x++ {
				xx := (x - in.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx
				p.Pix[off+0] = in.Y[yi]
				p.Pix[off+1] = in.Cb[ci]
				p.Pix[off+2] = in.Cr[ci]
				off += 3
			}
		}
	default:
		// Default to 4:4:4 subsampling.
		for y := in.Rect.Min.Y; y < in.Rect.Max.Y; y++ {
			yy := (y - in.Rect.Min.Y) * in.YStride
			cy := (y - in.Rect.Min.Y) * in.CStride
			for x := in.Rect.Min.X; x < in.Rect.Max.X; x++ {
				xx := (x - in.Rect.Min.X)
				yi := yy + xx
				ci := cy + xx
				p.Pix[off+0] = in.Y[yi]
				p.Pix[off+1] = in.Cb[ci]
				p.Pix[off+2] = in.Cr[ci]
				off += 3
			}
		}
	}
	return &p
}
