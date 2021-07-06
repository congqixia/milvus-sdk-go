// Code generated by go generate; DO NOT EDIT
// This file is generated by go genrated at 2021-07-05 22:52:01.455014639 &#43;0800 CST m=&#43;0.001674288

//package entity defines entities used in sdk
package entity 

import "github.com/milvus-io/milvus-sdk-go/internal/proto/schema"

// columnDouble generated columns type for Double
type columnDouble struct {
	name   string
	values []float64
}

func (c *columnDouble) Name() string {
	return c.name
}

func (c *columnDouble) Type() FieldType {
	return FieldTypeDouble
}

func (c *columnDouble) FieldData() *schema.FieldData {
	fd := &schema.FieldData{
		FieldName: c.name,
	}
	fd.Field = &schema.FieldData_Scalars{
		Scalars: &schema.ScalarField{
			Data: &schema.ScalarField_IntData{
				IntData: &schema.IntArray{
					Data: []int32{},
				},
			},
		},
	}
	return fd
}

func NewColumnDouble(name string, values []float64) Column {
	return &columnDouble {
		name: name,
		values: values,
	}
}
