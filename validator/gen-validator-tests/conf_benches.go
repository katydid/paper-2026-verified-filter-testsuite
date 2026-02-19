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
	"fmt"
	"log"
	"math/rand"

	"google.golang.org/protobuf/proto"
)

func RandomValidDate(r *rand.Rand) *Date {
	year := fmt.Sprintf("%d", 2000+r.Intn(100))
	month := fmt.Sprintf("%02d\n", 1+r.Intn(12))
	day := fmt.Sprintf("%02d\n", 1+r.Intn(30))
	return &Date{Year: proto.String(year), Month: proto.String(month), Day: proto.String(day)}
}

func RandomValidLocation(r *rand.Rand) *Location {
	l := RandomLocation(r).(*Location)
	conts := []string{"AF", "AN", "AS", "EU", "NA", "SA", "OC"}
	l.Cont = proto.String(conts[r.Intn(len(conts))])
	return l
}

func RandomValidConf(r *rand.Rand) ProtoMessage {
	c := RandomConf(r).(*Conf)
	if c.Due != nil {
		c.Due = RandomValidDate(r)
	}
	if c.Notify != nil {
		c.Notify = RandomValidDate(r)
	}
	if c.Loc != nil {
		c.Loc = RandomValidLocation(r)
	}
	return c
}

func RandomConf2026(r *rand.Rand) ProtoMessage {
	c := RandomValidConf(r).(*Conf)
	for c.GetDue() == nil {
		log.Printf("random looped: RandomValidConf")
		c = RandomValidConf(r).(*Conf)
		c.Due = RandomValidDate(r)
	}
	c.Due.Year = proto.String("2026")
	return c
}

func RandomConfNot2026(r *rand.Rand) ProtoMessage {
	c := RandomValidConf(r).(*Conf)
	for c.GetDue().GetYear() == "2026" {
		log.Printf("random looped: RandomValidConf")
		c = RandomValidConf(r).(*Conf)
		c.Due = RandomValidDate(r)
	}
	log.Printf("random returned: RandomValidConf")
	return c
}

func RandomConfIsIn2026OrLate2025AndEU(r *rand.Rand) ProtoMessage {
	c := RandomValidConf(r).(*Conf)
	for c.GetDue() == nil || c.GetLoc() == nil {
		log.Printf("random looped: RandomValidConf")
		c = RandomValidConf(r).(*Conf)
		c.Due = RandomValidDate(r)
		c.Loc = RandomValidLocation(r)
	}
	if r.Intn(2) == 0 {
		c.Due.Year = proto.String("2026")
	} else {
		c.Due.Year = proto.String("2025")
		switch r.Intn(2) {
		case 0:
			c.Due.Month = proto.String("11")
		case 1:
			c.Due.Month = proto.String("12")
		}
	}
	c.Loc.Cont = proto.String("EU")
	log.Printf("random returned: RandomValidConf")
	return c
}

func RandomConfComputerScience(r *rand.Rand) ProtoMessage {
	c := RandomValidConf(r).(*Conf)
	c.Category = proto.String("Computer Science")
	return c
}

func init() {
	BenchValidateProtoJsonReflect("ConfIs2026", ConfIsIn2026, RandomConf2026, RandomConfNot2026)
	BenchValidateProtoJsonReflect("ConfIsIn2026OrLate2025AndEU", ConfIsIn2026OrLate2025AndEU, RandomConfIsIn2026OrLate2025AndEU, RandomValidConf)
	BenchValidateProtoJsonReflect("ConfIsComputerScience", ConfIsComputerScience, RandomConfComputerScience, RandomValidConf)
}
