package entity

import "github.com/scantrust/dxf-golang/format"

type Vec = []float64

// Solid represents a triangle or square
// Minimum 3, Max 4 points.  If only 3 points, the 4th point == 3rd
// https://help.autodesk.com/view/OARX/2018/ENU/?guid=GUID-E0C5F04E-D0C5-48F5-AC09-32733E8848F2
//
// This is an incomplete implementation, enough for making our QR codes
type Solid struct {
	*entity
	Vertices  []Vec
	Thickness float64
	Extrusion Vec
}

// IsEntity is for Entity interface.
func (v *Solid) IsEntity() bool {
	return true
}

func NewSolid(vertices [][]float64) *Solid {
	v := &Solid{
		entity:    NewEntity(SOLID),
		Vertices:  vertices,
		Thickness: 0,
		Extrusion: []float64{},
	}
	return v
}

// Rectangle Creates a SOLID (AcDbTrace) at x,y as the LOWER LEFT CORNER
// Pass a negative height grow the rect down towards the X axis
func NewRect2D(x, y, width float64, height float64) *Solid {
	return NewSolid(
		[][]float64{
			{x, y, 0},
			{x + width, y, 0},
			{x, y + height, 0},
			{x + width, y + height, 0},
		},
	)
}

func (v *Solid) Format(f format.Formatter) {
	v.entity.Format(f)
	f.WriteString(100, "AcDbTrace")
	writeVector(f, 10, v.Vertices[0])
	writeVector(f, 11, v.Vertices[1])
	writeVector(f, 12, v.Vertices[2])
	writeVector(f, 13, v.Vertices[3])
	if v.Thickness != 0 {
		f.WriteFloat(39, v.Thickness) // thickness
	}
	if len(v.Extrusion) != 0 {
		writeVector(f, 210, v.Extrusion)
	}
}

func (v *Solid) String() string {
	f := format.NewASCII()
	return v.FormatString(f)
}

// FormatString outputs data using given formatter.
func (v *Solid) FormatString(f format.Formatter) string {
	v.Format(f)
	return f.Output()
}

// BBox fake implementation - not used anywhere, just return
// a unit square
func (v *Solid) BBox() ([]float64, []float64) {
	return []float64{0, 0, 0}, []float64{1, 1, 1}
}

func writeVector(f format.Formatter, startId int, vert []float64) {
	f.WriteFloat(startId, vert[0])
	f.WriteFloat(startId+10, vert[1])
	f.WriteFloat(startId+20, vert[2])
}
