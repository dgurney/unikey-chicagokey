package main

import (
	"fmt"
	"time"

	"github.com/dgurney/unikey/generator"
)

/*
   Copyright (C) 2021 Daniel Gurney
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// generationBenchmark generates the specified amount of keys and benchmarks it
func generationBenchmark(amount int) {
	chicago := generator.ChicagoCredentials{Build: "73f"}
	started := time.Now()
	for i := 0; i < amount; i++ {
		chicago.Generate()
	}

	var ended time.Duration
	switch {
	case time.Since(started).Round(time.Second) > 1:
		ended = time.Since(started).Round(time.Millisecond)
	default:
		ended = time.Since(started).Round(time.Microsecond)
	}
	fmt.Printf("Took %s to generate %d keys.\n", ended, amount)
}
