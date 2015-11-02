// This file was generated by nomdl/codegen.

package common

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __commonPackageInFile_incident_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Incident",
			[]types.Field{
				types.Field{"ID", types.MakePrimitiveTypeRef(types.Int64Kind), false},
				types.Field{"Category", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Description", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"DayOfWeek", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Date", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Time", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"PdDistrict", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Resolution", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Address", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Geoposition", types.MakeTypeRef(ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"), 0), false},
				types.Field{"PdID", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"),
	})
	__commonPackageInFile_incident_CachedRef = types.RegisterPackage(&p)
}

// Incident

type Incident struct {
	m   types.Map
	ref *ref.Ref
}

func NewIncident() Incident {
	return Incident{types.NewMap(
		types.NewString("ID"), types.Int64(0),
		types.NewString("Category"), types.NewString(""),
		types.NewString("Description"), types.NewString(""),
		types.NewString("DayOfWeek"), types.NewString(""),
		types.NewString("Date"), types.NewString(""),
		types.NewString("Time"), types.NewString(""),
		types.NewString("PdDistrict"), types.NewString(""),
		types.NewString("Resolution"), types.NewString(""),
		types.NewString("Address"), types.NewString(""),
		types.NewString("Geoposition"), NewGeoposition(),
		types.NewString("PdID"), types.NewString(""),
	), &ref.Ref{}}
}

type IncidentDef struct {
	ID          int64
	Category    string
	Description string
	DayOfWeek   string
	Date        string
	Time        string
	PdDistrict  string
	Resolution  string
	Address     string
	Geoposition GeopositionDef
	PdID        string
}

func (def IncidentDef) New() Incident {
	return Incident{
		types.NewMap(
			types.NewString("ID"), types.Int64(def.ID),
			types.NewString("Category"), types.NewString(def.Category),
			types.NewString("Description"), types.NewString(def.Description),
			types.NewString("DayOfWeek"), types.NewString(def.DayOfWeek),
			types.NewString("Date"), types.NewString(def.Date),
			types.NewString("Time"), types.NewString(def.Time),
			types.NewString("PdDistrict"), types.NewString(def.PdDistrict),
			types.NewString("Resolution"), types.NewString(def.Resolution),
			types.NewString("Address"), types.NewString(def.Address),
			types.NewString("Geoposition"), def.Geoposition.New(),
			types.NewString("PdID"), types.NewString(def.PdID),
		), &ref.Ref{}}
}

func (s Incident) Def() (d IncidentDef) {
	d.ID = int64(s.m.Get(types.NewString("ID")).(types.Int64))
	d.Category = s.m.Get(types.NewString("Category")).(types.String).String()
	d.Description = s.m.Get(types.NewString("Description")).(types.String).String()
	d.DayOfWeek = s.m.Get(types.NewString("DayOfWeek")).(types.String).String()
	d.Date = s.m.Get(types.NewString("Date")).(types.String).String()
	d.Time = s.m.Get(types.NewString("Time")).(types.String).String()
	d.PdDistrict = s.m.Get(types.NewString("PdDistrict")).(types.String).String()
	d.Resolution = s.m.Get(types.NewString("Resolution")).(types.String).String()
	d.Address = s.m.Get(types.NewString("Address")).(types.String).String()
	d.Geoposition = s.m.Get(types.NewString("Geoposition")).(Geoposition).Def()
	d.PdID = s.m.Get(types.NewString("PdID")).(types.String).String()
	return
}

var __typeRefForIncident types.TypeRef

func (m Incident) TypeRef() types.TypeRef {
	return __typeRefForIncident
}

func init() {
	__typeRefForIncident = types.MakeTypeRef(__commonPackageInFile_incident_CachedRef, 0)
	types.RegisterFromValFunction(__typeRefForIncident, func(v types.Value) types.Value {
		return IncidentFromVal(v)
	})
}

func IncidentFromVal(val types.Value) Incident {
	// TODO: Do we still need FromVal?
	if val, ok := val.(Incident); ok {
		return val
	}
	// TODO: Validate here
	return Incident{val.(types.Map), &ref.Ref{}}
}

func (s Incident) InternalImplementation() types.Map {
	return s.m
}

func (s Incident) Equals(other types.Value) bool {
	return other != nil && __typeRefForIncident.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s Incident) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Incident) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.TypeRef().Chunks()...)
	chunks = append(chunks, s.m.Chunks()...)
	return
}

func (s Incident) ID() int64 {
	return int64(s.m.Get(types.NewString("ID")).(types.Int64))
}

func (s Incident) SetID(val int64) Incident {
	return Incident{s.m.Set(types.NewString("ID"), types.Int64(val)), &ref.Ref{}}
}

func (s Incident) Category() string {
	return s.m.Get(types.NewString("Category")).(types.String).String()
}

func (s Incident) SetCategory(val string) Incident {
	return Incident{s.m.Set(types.NewString("Category"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Description() string {
	return s.m.Get(types.NewString("Description")).(types.String).String()
}

func (s Incident) SetDescription(val string) Incident {
	return Incident{s.m.Set(types.NewString("Description"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) DayOfWeek() string {
	return s.m.Get(types.NewString("DayOfWeek")).(types.String).String()
}

func (s Incident) SetDayOfWeek(val string) Incident {
	return Incident{s.m.Set(types.NewString("DayOfWeek"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Date() string {
	return s.m.Get(types.NewString("Date")).(types.String).String()
}

func (s Incident) SetDate(val string) Incident {
	return Incident{s.m.Set(types.NewString("Date"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Time() string {
	return s.m.Get(types.NewString("Time")).(types.String).String()
}

func (s Incident) SetTime(val string) Incident {
	return Incident{s.m.Set(types.NewString("Time"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) PdDistrict() string {
	return s.m.Get(types.NewString("PdDistrict")).(types.String).String()
}

func (s Incident) SetPdDistrict(val string) Incident {
	return Incident{s.m.Set(types.NewString("PdDistrict"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Resolution() string {
	return s.m.Get(types.NewString("Resolution")).(types.String).String()
}

func (s Incident) SetResolution(val string) Incident {
	return Incident{s.m.Set(types.NewString("Resolution"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Address() string {
	return s.m.Get(types.NewString("Address")).(types.String).String()
}

func (s Incident) SetAddress(val string) Incident {
	return Incident{s.m.Set(types.NewString("Address"), types.NewString(val)), &ref.Ref{}}
}

func (s Incident) Geoposition() Geoposition {
	return s.m.Get(types.NewString("Geoposition")).(Geoposition)
}

func (s Incident) SetGeoposition(val Geoposition) Incident {
	return Incident{s.m.Set(types.NewString("Geoposition"), val), &ref.Ref{}}
}

func (s Incident) PdID() string {
	return s.m.Get(types.NewString("PdID")).(types.String).String()
}

func (s Incident) SetPdID(val string) Incident {
	return Incident{s.m.Set(types.NewString("PdID"), types.NewString(val)), &ref.Ref{}}
}

// ListOfRefOfIncident

type ListOfRefOfIncident struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfRefOfIncident() ListOfRefOfIncident {
	return ListOfRefOfIncident{types.NewList(), &ref.Ref{}}
}

type ListOfRefOfIncidentDef []ref.Ref

func (def ListOfRefOfIncidentDef) New() ListOfRefOfIncident {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = NewRefOfIncident(d)
	}
	return ListOfRefOfIncident{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfRefOfIncident) Def() ListOfRefOfIncidentDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(RefOfIncident).TargetRef()
	}
	return d
}

func ListOfRefOfIncidentFromVal(val types.Value) ListOfRefOfIncident {
	// TODO: Do we still need FromVal?
	if val, ok := val.(ListOfRefOfIncident); ok {
		return val
	}
	// TODO: Validate here
	return ListOfRefOfIncident{val.(types.List), &ref.Ref{}}
}

func (l ListOfRefOfIncident) InternalImplementation() types.List {
	return l.l
}

func (l ListOfRefOfIncident) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfRefOfIncident.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfRefOfIncident) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfRefOfIncident) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfRefOfIncident.
var __typeRefForListOfRefOfIncident types.TypeRef

func (m ListOfRefOfIncident) TypeRef() types.TypeRef {
	return __typeRefForListOfRefOfIncident
}

func init() {
	__typeRefForListOfRefOfIncident = types.MakeCompoundTypeRef(types.ListKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(__commonPackageInFile_incident_CachedRef, 0)))
	types.RegisterFromValFunction(__typeRefForListOfRefOfIncident, func(v types.Value) types.Value {
		return ListOfRefOfIncidentFromVal(v)
	})
}

func (l ListOfRefOfIncident) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfIncident) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfIncident) Get(i uint64) RefOfIncident {
	return l.l.Get(i).(RefOfIncident)
}

func (l ListOfRefOfIncident) Slice(idx uint64, end uint64) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfIncident) Set(i uint64, val RefOfIncident) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfRefOfIncident) Append(v ...RefOfIncident) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfIncident) Insert(idx uint64, v ...RefOfIncident) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfIncident) Remove(idx uint64, end uint64) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfIncident) RemoveAt(idx uint64) ListOfRefOfIncident {
	return ListOfRefOfIncident{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfRefOfIncident) fromElemSlice(p []RefOfIncident) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfRefOfIncidentIterCallback func(v RefOfIncident, i uint64) (stop bool)

func (l ListOfRefOfIncident) Iter(cb ListOfRefOfIncidentIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfIncident), i)
	})
}

type ListOfRefOfIncidentIterAllCallback func(v RefOfIncident, i uint64)

func (l ListOfRefOfIncident) IterAll(cb ListOfRefOfIncidentIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(RefOfIncident), i)
	})
}

type ListOfRefOfIncidentFilterCallback func(v RefOfIncident, i uint64) (keep bool)

func (l ListOfRefOfIncident) Filter(cb ListOfRefOfIncidentFilterCallback) ListOfRefOfIncident {
	nl := NewListOfRefOfIncident()
	l.IterAll(func(v RefOfIncident, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// RefOfIncident

type RefOfIncident struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfIncident(target ref.Ref) RefOfIncident {
	return RefOfIncident{target, &ref.Ref{}}
}

func (r RefOfIncident) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfIncident) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfIncident) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfIncident.Equals(other.TypeRef()) && r.Ref() == other.Ref()
}

func (r RefOfIncident) Chunks() []ref.Ref {
	return r.TypeRef().Chunks()
}

func RefOfIncidentFromVal(val types.Value) RefOfIncident {
	// TODO: Do we still need FromVal?
	if val, ok := val.(RefOfIncident); ok {
		return val
	}
	return NewRefOfIncident(val.(types.Ref).TargetRef())
}

// A Noms Value that describes RefOfIncident.
var __typeRefForRefOfIncident types.TypeRef

func (m RefOfIncident) TypeRef() types.TypeRef {
	return __typeRefForRefOfIncident
}

func init() {
	__typeRefForRefOfIncident = types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(__commonPackageInFile_incident_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForRefOfIncident, func(v types.Value) types.Value {
		return RefOfIncidentFromVal(v)
	})
}

func (r RefOfIncident) TargetValue(cs chunks.ChunkSource) Incident {
	return types.ReadValue(r.target, cs).(Incident)
}

func (r RefOfIncident) SetTargetValue(val Incident, cs chunks.ChunkSink) RefOfIncident {
	return NewRefOfIncident(types.WriteValue(val, cs))
}