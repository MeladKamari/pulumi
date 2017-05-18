// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/pulumi/lumi/pkg/compiler/symbols"
	"github.com/pulumi/lumi/pkg/tokens"
)

// All of the primitive types.
var (
	Object  = symbols.NewPrimitiveType("object")  // the base of all types.
	Bool    = symbols.NewPrimitiveType("bool")    // a bool (true or false) primitive.
	Number  = symbols.NewPrimitiveType("number")  // a 64-bit IEEE754 floating point primitive.
	String  = symbols.NewPrimitiveType("string")  // a UTF8 encoded string.
	Null    = symbols.NewPrimitiveType("null")    // the special null literal type.
	Dynamic = symbols.NewPrimitiveType("dynamic") // a type that opts into automatic dynamic conversions.
)

// Primitives contains a map of all primitive types, keyed by their token/name.
var Primitives = map[tokens.TypeName]symbols.Type{
	Object.Nm:  Object,
	Bool.Nm:    Bool,
	Number.Nm:  Number,
	String.Nm:  String,
	Null.Nm:    Null,
	Dynamic.Nm: Dynamic,
}

// Special types that aren't intended for public use.
var (
	Error = symbols.NewPrimitiveType("<error>") // a type for internal compiler errors; not for direct use.
)
