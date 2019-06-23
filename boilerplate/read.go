package main

import "strconv"

func nextUint64() uint64 {
	if !sc.Scan() {
		panic(nil)
	}
	i, e := strconv.ParseUint(sc.Text(), 0, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}
