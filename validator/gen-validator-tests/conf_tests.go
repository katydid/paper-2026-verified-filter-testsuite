//  Copyright 2026 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	. "github.com/katydid/validator-go/validator/combinator"
	"google.golang.org/protobuf/proto"
)

var ConfIsIn2026 = G{"main": InPath("Due", InPath("Y", Value(Eq(StringVar(), StringConst("2026")))))}

var ConfITP2026 = &Conf{
	N: proto.String("ITP"),
	Due: &Date{
		Y: proto.String("2026"),
	},
}

var ConfICFP2017 = &Conf{
	N: proto.String("ICFP"),
	Due: &Date{
		Y: proto.String("2017"),
	},
}

var ConfIsIn2026OrLate2025 = G{"main": InPath("Due",
	AnyOf(
		InPath("Y", Value(Eq(StringVar(), StringConst("2026")))),
		AllOf(
			InPath("Y", Value(Eq(StringVar(), StringConst("2025")))),
			InPath("M", Value(GE(StringVar(), StringConst("10")))),
		),
	),
)}

var ConfITPLate2025 = &Conf{
	N: proto.String("ITP"),
	Due: &Date{
		Y: proto.String("2025"),
		M: proto.String("11"),
	},
}

var ConfITPEarly2025 = &Conf{
	N: proto.String("ITP"),
	Due: &Date{
		Y: proto.String("2025"),
		M: proto.String("04"),
	},
}

var ConfIsIn2026OrLate2025AndEU = G{"main": InAnyOrder(
	In("Due",
		AnyOf(
			InPath("Y", Value(Eq(StringVar(), StringConst("2026")))),
			AllOf(
				InPath("Y", Value(Eq(StringVar(), StringConst("2025")))),
				InPath("M", Value(GE(StringVar(), StringConst("10")))),
			),
		),
	),
	In("L",
		InPath("Cont", Value(Eq(StringVar(), StringConst("EU")))),
	),
	Any(),
)}

var ConfITPEU = &Conf{
	N: proto.String("ITP"),
	Due: &Date{
		Y: proto.String("2025"),
		M: proto.String("11"),
	},
	L: &Location{
		Cont: proto.String("EU"),
	},
}

var ConfITPNotEU = &Conf{
	N: proto.String("ITP"),
	Due: &Date{
		Y: proto.String("2026"),
	},
	L: &Location{
		Cont: proto.String("AF"),
	},
}

func init() {
	ValidateProtoEtc("ConfIsIn2026ITP2026", ConfIsIn2026, ConfITP2026, true)
	ValidateProtoEtc("ConfIsIn2026ICFP2017", ConfIsIn2026, ConfICFP2017, false)
	ValidateProtoEtc("ConfIsIn2026OrLate2025ConfITPLate2025", ConfIsIn2026OrLate2025, ConfITPLate2025, true)
	ValidateProtoEtc("ConfIsIn2026OrLate2025ConfITPEarly2025", ConfIsIn2026OrLate2025, ConfITPEarly2025, false)
	ValidateProtoEtc("ConfIsIn2026OrLate2025AndEUConfITPEU", ConfIsIn2026OrLate2025AndEU, ConfITPEU, true)
	ValidateProtoEtc("ConfIsIn2026OrLate2025AndEUConfITPNotEU", ConfIsIn2026OrLate2025AndEU, ConfITPNotEU, false)
}
