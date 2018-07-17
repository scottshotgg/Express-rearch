%{

//TODO Put your favorite license here
		
// yacc source generated by ebnf2y[1]
// at 2018-07-12 17:08:40.031161687 -0500 CDT m=+0.000623483
//
//  $ ebnf2y -pkg gen -start Sched sched.ebnf
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
// 
//   [1]: http://github.com/cznic/ebnf2y

package gen //TODO real package name

//TODO required only be the demo _dump function
import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cznic/strutil"
)

%}

%union {
	item interface{} //TODO insert real field(s)
}

%token	COMMA
%token	DAY
%token	DECIMAL_DIGIT
%token	EVERY
%token	HRS
%token	MINS
%token	MONTH
%token	NEWLINE
%token	ORDINAL
%token	SECS
%token	TIME

%type	<item> 	/*TODO real type(s), if/where applicable */
	COMMA
	DAY
	DECIMAL_DIGIT
	EVERY
	HRS
	MINS
	MONTH
	NEWLINE
	ORDINAL
	SECS
	TIME

%token FROM
%token OF
%token TO

%type	<item> 	/*TODO real type(s), if/where applicable */
	Days
	Days1
	Months
	Months1
	RecurByDay
	RecurByDay1
	RecurByDay2
	RecurByDay21
	RecurByDay3
	RecurByTime
	RecurByTime1
	RecurByTime2
	Sched
	Sched1
	SchedLine
	Start

/*TODO %left, %right, ... declarations */

%start Start

%%

Days:
	DAY Days1
	{
		$$ = []Days{$1, $2} //TODO 1
	}

Days1:
	/* EMPTY */
	{
		$$ = []Days1(nil) //TODO 2
	}
|	Days1 COMMA DAY
	{
		$$ = append($1.([]Days1), $2, $3) //TODO 3
	}

Months:
	MONTH Months1
	{
		$$ = []Months{$1, $2} //TODO 4
	}

Months1:
	/* EMPTY */
	{
		$$ = []Months1(nil) //TODO 5
	}
|	Months1 COMMA MONTH
	{
		$$ = append($1.([]Months1), $2, $3) //TODO 6
	}

RecurByDay:
	RecurByDay1 Days RecurByDay2 RecurByDay3
	{
		$$ = []RecurByDay{$1, $2, $3, $4} //TODO 7
	}

RecurByDay1:
	EVERY
	{
		$$ = $1 //TODO 8
	}
|	ORDINAL
	{
		$$ = $1 //TODO 9
	}

RecurByDay2:
	/* EMPTY */
	{
		$$ = nil //TODO 10
	}
|	OF RecurByDay21
	{
		$$ = []RecurByDay2{"of", $2} //TODO 11
	}

RecurByDay21:
	Months
	{
		$$ = $1 //TODO 12
	}

RecurByDay3:
	TIME
	{
		$$ = $1 //TODO 13
	}

RecurByTime:
	EVERY DECIMAL_DIGIT RecurByTime1 RecurByTime2
	{
		$$ = []RecurByTime{$1, $2, $3, $4} //TODO 14
	}

RecurByTime1:
	HRS
	{
		$$ = $1 //TODO 15
	}
|	MINS
	{
		$$ = $1 //TODO 16
	}
|	SECS
	{
		$$ = $1 //TODO 17
	}

RecurByTime2:
	/* EMPTY */
	{
		$$ = nil //TODO 18
	}
|	FROM TIME TO TIME
	{
		$$ = []RecurByTime2{"from", $2, "to", $4} //TODO 19
	}

Sched:
	SchedLine Sched1
	{
		$$ = []Sched{$1, $2} //TODO 20
	}

Sched1:
	/* EMPTY */
	{
	}
|	Sched1 SchedLine
	{
	}

SchedLine:
	NEWLINE
	{
	}
|	RecurByTime
	{
		if v, ok := Schedlex.(*SchedLex); ok {
			v.OnSched($1)
		}	
	}
|	RecurByDay
	{
		if v, ok := Schedlex.(*SchedLex); ok {
			v.OnSched($1) 
		}	
	}

Start:
	Sched
	{
	}

%%

type (
	Days []string
	Days1 []string
	Months []string
	Months1 []string
	RecurByDay struct{
		Ord int
		Days []string
		Months []string
		Clock string
	}

	RecurByDay1 string
 	RecurByDay2 interface{}
	RecurByDay21 []string
	RecurByDay3 string
	RecurByTime struct{
		N          int
		ClockLbl   string
		ClockRange RecurByTime2
	}

 	RecurByTime1 interface{}
	RecurByTime2 struct {
		From string
		To string		
	}
 	Sched interface{}
 	Sched1 interface{}
 	SchedLine interface{}
 	Start interface{}
)