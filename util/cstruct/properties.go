package cstruct

import (
	"reflect"
	"sync"
)

// StructProperties

type StructProperties struct {
	Prop      []*Properties
	fixedSize int
}

var (
	propertiesMu  sync.RWMutex
	propertiesMap = make(map[reflect.Type]*StructProperties)
)

func GetProperties(t reflect.Type) *StructProperties {
	if t.Kind() != reflect.Struct {
		panic("cstruct: type must have kind cstruct")
	}

	propertiesMu.RLock()
	sprop, ok := propertiesMap[t]
	propertiesMu.RUnlock()
	if ok {
		return sprop
	}

	propertiesMu.Lock()
	defer propertiesMu.Unlock()
	sprop = getPropertiesLocked(t)
	return sprop
}

func getPropertiesLocked(t reflect.Type) *StructProperties {
	if prop, ok := propertiesMap[t]; ok {
		return prop
	}

	prop := new(StructProperties)
	propertiesMap[t] = prop
	prop.Prop = make([]*Properties, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		p := new(Properties)
		name := f.Name
		p.init(f.Type, name, "", &f, &prop.fixedSize)
		prop.Prop[i] = p
	}

	return prop
}

// Properties

type encoder func(p *Buffer, prop *Properties, base structPointer) error
type decoder func(p *Buffer, prop *Properties, base structPointer) error
type sizer func(p *Buffer, prop *Properties, base structPointer) int

type Properties struct {
	Name  string
	field field
	enc   encoder
	dec   decoder
	siz   sizer
	t     reflect.Type
	stype reflect.Type
	sprop *StructProperties
}

func (p *Properties) init(typ reflect.Type, name, tag string, f *reflect.StructField, fixedSize *int) {
	p.Name = name
	if f != nil {
		p.field = toField(f)
	}
	p.setEncAndDec(typ, f, fixedSize)
}

func (p *Properties) setEncAndDec(typ reflect.Type, f *reflect.StructField, fixedSize *int) {
	p.enc = nil
	p.dec = nil
	p.siz = nil
	p.t = typ

	switch typ.Kind() {
	case reflect.Bool: // bool
		p.enc = (*Buffer).enc_bool
		p.dec = (*Buffer).dec_bool
		*fixedSize += 1
	case reflect.Int8, reflect.Uint8: // int8 uint8
		p.enc = (*Buffer).enc_uint8
		p.dec = (*Buffer).dec_uint8
		*fixedSize += 1
	case reflect.Int16, reflect.Uint16: // int16 uint16
		p.enc = (*Buffer).enc_uint16
		p.dec = (*Buffer).dec_uint16
		*fixedSize += 2
	case reflect.Int32, reflect.Uint32, reflect.Float32: // int32 uint32 float32
		p.enc = (*Buffer).enc_uint32
		p.dec = (*Buffer).dec_uint32
		*fixedSize += 4
	case reflect.Int64, reflect.Uint64, reflect.Float64: // int64 uint64 float64
		p.enc = (*Buffer).enc_uint64
		p.dec = (*Buffer).dec_uint64
		*fixedSize += 8
	case reflect.String: // string
		p.enc = (*Buffer).enc_string
		p.dec = (*Buffer).dec_string
		p.siz = (*Buffer).size_string
		*fixedSize += 2
	case reflect.Ptr: // cstruct ptr
		if t2 := typ.Elem(); t2.Kind() == reflect.Struct {
			p.stype = t2
			p.sprop = getPropertiesLocked(p.stype)
			p.enc = (*Buffer).enc_substruct_ptr
			p.dec = (*Buffer).dec_substruct_ptr
			p.siz = (*Buffer).size_substruct_ptr
		} else {
			panic("cstruct: unknow type. field name =" + f.Name)
		}
	case reflect.Struct: // cstruct
		p.stype = typ
		p.sprop = getPropertiesLocked(p.stype)
		p.enc = (*Buffer).enc_substruct
		p.dec = (*Buffer).dec_substruct
		p.siz = (*Buffer).size_substruct
	case reflect.Slice:
		*fixedSize += 2
		switch t2 := typ.Elem(); t2.Kind() {
		case reflect.Uint8, reflect.Int8: // []byte []uint8 []int8
			p.enc = (*Buffer).enc_slice_byte
			p.dec = (*Buffer).dec_slice_byte
			p.siz = (*Buffer).size_slice_byte
		case reflect.Bool: // []bool
			p.enc = (*Buffer).enc_slice_bool
			p.dec = (*Buffer).dec_slice_bool
			p.siz = (*Buffer).size_slice_bool
		case reflect.Uint16, reflect.Int16: // []uint16 []int16
			p.enc = (*Buffer).enc_slice_uint16
			p.dec = (*Buffer).dec_slice_uint16
			p.siz = (*Buffer).size_slice_uint16
		case reflect.Uint32, reflect.Int32, reflect.Float32: // []uint32 []int32 []float32
			p.enc = (*Buffer).enc_slice_uint32
			p.dec = (*Buffer).dec_slice_uint32
			p.siz = (*Buffer).size_slice_uint32
		case reflect.Uint64, reflect.Int64, reflect.Float64: // []uint64 []int64 []float64
			p.enc = (*Buffer).enc_slice_uint64
			p.dec = (*Buffer).dec_slice_uint64
			p.siz = (*Buffer).size_slice_uint64
		case reflect.String: // []string
			p.enc = (*Buffer).enc_slice_string
			p.dec = (*Buffer).dec_slice_string
			p.siz = (*Buffer).size_slice_string
		case reflect.Struct: // [] cstruct
			p.stype = t2
			p.sprop = getPropertiesLocked(p.stype)
			p.enc = (*Buffer).enc_slice_substruct
			p.dec = (*Buffer).dec_slice_substruct
			p.siz = (*Buffer).size_slice_substruct
		case reflect.Ptr: // []*cstruct
			switch t3 := t2.Elem(); t3.Kind() {
			case reflect.Struct:
				p.stype = t3
				p.sprop = getPropertiesLocked(p.stype)
				if OptionSliceIgnoreNil == false {
					p.enc = (*Buffer).enc_slice_substruct_ptr
					p.dec = (*Buffer).dec_slice_substruct_ptr
					p.siz = (*Buffer).size_slice_substruct_ptr
				} else {
					p.enc = (*Buffer).enc_slice_substruct_ptr_ignore_nil
					p.dec = (*Buffer).dec_slice_substruct_ptr_ignore_nil
					p.siz = (*Buffer).size_slice_substruct_ptr_ignore_nil
				}
			default:
				panic("cstruct: unknow type. field name = " + f.Name)
			}
		case reflect.Slice:
			switch t2.Elem().Kind() {
			case reflect.Uint8:
				p.enc = (*Buffer).enc_slice_slice_byte
				p.dec = (*Buffer).dec_slice_slice_byte
				p.siz = (*Buffer).size_slice_slice_byte
			default:
				panic("cstruct: unknow type. field name = " + f.Name)
			}
		default:
			panic("cstruct: unknow type. field name = " + f.Name)
		}
	case reflect.Array:
		switch t2 := typ.Elem(); t2.Kind() {
		case reflect.Uint8, reflect.Int8, reflect.Bool: // [n]byte [n]uint8 [n]int8 [n]bool
			p.enc = (*Buffer).enc_array_byte
			p.dec = (*Buffer).dec_array_byte
			p.siz = (*Buffer).size_array_byte
		case reflect.Uint16, reflect.Int16: // [n]uint16 [n]int16
			p.enc = (*Buffer).enc_array_uint16
			p.dec = (*Buffer).dec_array_uint16
			p.siz = (*Buffer).size_array_uint16
		case reflect.Uint32, reflect.Int32, reflect.Float32: // [n]uint32 [n]int32 [n]float32
			p.enc = (*Buffer).enc_array_uint32
			p.dec = (*Buffer).dec_array_uint32
			p.siz = (*Buffer).size_array_uint32
		case reflect.Uint64, reflect.Int64, reflect.Float64: // [n]uint64 [n]int64 [n]float64
			p.enc = (*Buffer).enc_array_uint64
			p.dec = (*Buffer).dec_array_uint64
			p.siz = (*Buffer).size_array_uint64
		case reflect.Struct: // [n]cstruct
			p.stype = t2
			p.sprop = getPropertiesLocked(p.stype)
			p.enc = (*Buffer).enc_array_substruct
			p.dec = (*Buffer).dec_array_substruct
			p.siz = (*Buffer).size_array_substruct
		default:
			panic("cstruct: unknow type. field name = " + f.Name)
		}
	default:
		panic("cstruct: unknow type. field name = " + f.Name)
	}
}
