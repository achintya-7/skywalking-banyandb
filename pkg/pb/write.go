// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pb

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/apache/skywalking-banyandb/api/proto/banyandb/v1"
	"github.com/apache/skywalking-banyandb/pkg/convert"
)

type writeEntityBuilder struct {
	we *v1.WriteEntity
}

func NewWriteEntityBuilder() *writeEntityBuilder {
	return &writeEntityBuilder{we: &v1.WriteEntity{}}
}

func (web *writeEntityBuilder) Metadata(group, name string) *writeEntityBuilder {
	web.we.Metadata = &v1.Metadata{
		Group: group,
		Name:  name,
	}
	return web
}

func (web *writeEntityBuilder) EntityValue(ev *v1.EntityValue) *writeEntityBuilder {
	web.we.Entity = ev
	return web
}

func (web *writeEntityBuilder) Build() *v1.WriteEntity {
	return web.we
}

type entityValueBuilder struct {
	ev *v1.EntityValue
}

func NewEntityValueBuilder() *entityValueBuilder {
	return &entityValueBuilder{ev: &v1.EntityValue{}}
}

func (evb *entityValueBuilder) EntityId(entityId string) *entityValueBuilder {
	evb.ev.EntityId = entityId
	return evb
}

func (evb *entityValueBuilder) Timestamp(time time.Time) *entityValueBuilder {
	evb.ev.Timestamp = timestamppb.New(time)
	return evb
}

func (evb *entityValueBuilder) DataBinary(data []byte) *entityValueBuilder {
	evb.ev.DataBinary = data
	return evb
}

func (evb *entityValueBuilder) Fields(items ...interface{}) *entityValueBuilder {
	evb.ev.Fields = make([]*v1.Field, len(items))
	for idx, item := range items {
		evb.ev.Fields[idx] = buildField(item)
	}
	return evb
}

func buildField(value interface{}) *v1.Field {
	if value == nil {
		return &v1.Field{ValueType: &v1.Field_Null{}}
	}
	switch v := value.(type) {
	case string:
		return &v1.Field{
			ValueType: &v1.Field_Str{
				Str: &v1.Str{
					Value: v,
				},
			},
		}
	case []string:
		return &v1.Field{
			ValueType: &v1.Field_StrArray{
				StrArray: &v1.StrArray{
					Value: v,
				},
			},
		}
	case int:
		return &v1.Field{
			ValueType: &v1.Field_Int{
				Int: &v1.Int{
					Value: int64(v),
				},
			},
		}
	case []int:
		return &v1.Field{
			ValueType: &v1.Field_IntArray{
				IntArray: &v1.IntArray{
					Value: convert.IntToInt64(v...),
				},
			},
		}
	case int64:
		return &v1.Field{
			ValueType: &v1.Field_Int{
				Int: &v1.Int{
					Value: v,
				},
			},
		}
	case []int64:
		return &v1.Field{
			ValueType: &v1.Field_IntArray{
				IntArray: &v1.IntArray{
					Value: v,
				},
			},
		}
	default:
		panic("type not supported")
	}
}

func (evb *entityValueBuilder) Build() *v1.EntityValue {
	return evb.ev
}
