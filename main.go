package main

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

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgurney/unikey/generator"
)

const version = "0.4.0"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	bench := flag.Int("bench", 0, "Benchmark generation of N credentials")
	build := flag.String("b", "", "Generate credentials for build (can be 73f, 73g, 81, 99 (works up to 116), 122 (works up to 189), 216 (works up to 302), ie4july (4.70.1169), ie4sept (4.71.0225))")
	repeat := flag.Int("r", 1, "Repeat N times")
	t := flag.Bool("t", false, "Show elapsed time")
	ver := flag.Bool("ver", false, "Show version information and exit.")
	flag.Parse()

	if *ver {
		fmt.Printf("unikey-chicagokey v%s by Daniel Gurney\n", version)
		return
	}

	if *bench > 0 {
		generationBenchmark(*bench)
		return
	}

	if *repeat < 1 {
		*repeat = 1
	}

	if *build == "" || (*build != "73g" && *build != "73f" && *build != "81" && *build != "99" && *build != "122" && *build != "216" && *build != "ie4july" && *build != "ie4sept") {
		fmt.Println("You must specify a valid build! Usage:")
		flag.PrintDefaults()
		return
	}

	var started time.Time
	if *t {
		started = time.Now()
	}

	key := generator.ChicagoCredentials{Build: *build}
	for i := 0; i < *repeat; i++ {
		k, err := generator.Generate(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(k.String())
	}

	if *t {
		var ended time.Duration
		switch {
		case time.Since(started).Round(time.Second) > 1:
			ended = time.Since(started).Round(time.Millisecond)
		default:
			ended = time.Since(started).Round(time.Microsecond)
		}
		if ended < 1 {
			// Oh Windows...
			fmt.Println("Could not display elapsed time correctly :(")
			return
		}
		switch {
		case *repeat > 1:
			fmt.Printf("Took %s to generate %d keys.\n", ended, *repeat)
			return
		case *repeat == 1:
			fmt.Printf("Took %s to generate %d key.\n", ended, *repeat)
			return
		}
	}
}
