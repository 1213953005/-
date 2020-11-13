// This file defines program parameters and routines for initializing
// them from the command line.

package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
)

// Parameters is a collection of all program parameters.
type Parameters struct {
	TTName  string     // Name of the input truth-table file
	MinQ    float64    // Minimum quadratic coefficient
	MaxQ    float64    // Maximum quadratic coefficient
	MinL    float64    // Minimum linear coefficient
	MaxL    float64    // Maximum linear coefficient
	TT      TruthTable // The truth-table proper
	NCols   int        // Number of columns in the truth table
	AllCols *mat.Dense // Matrix with all 2^n columns for n rows
}

// ParseCommandLine parses parameters from the command line.
func ParseCommandLine(p *Parameters) {
	// Parse the command line.
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [<options>] [<input.tt>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Float64Var(&p.MinQ, "qmin", -1.0, "Minimum quadratic coefficient")
	flag.Float64Var(&p.MaxQ, "qmax", 1.0, "Maximum quadratic coefficient")
	flag.Float64Var(&p.MinL, "lmin", -1.0, "Minimum linear coefficient")
	flag.Float64Var(&p.MaxL, "lmax", 1.0, "Maximum linear coefficient")
	flag.Parse()
	if flag.NArg() >= 1 {
		p.TTName = flag.Arg(0)
	}

	// Validate the arguments.
	if p.MinQ >= p.MaxQ {
		notify.Fatal("--qmin must specify a value that is less than --qmax")
	}
	if p.MinL >= p.MaxL {
		notify.Fatal("--lmin must specify a value that is less than --lmax")
	}
}
