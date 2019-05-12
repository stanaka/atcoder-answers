package main

import (
	"fmt"
	"io"
	"os"
)

func primeFactors(n uint64) (pfs map[uint64]uint64) {
	pfs = make(map[uint64]uint64)
	for n%2 == 0 {
		pfs[2]++
		n = n / 2
	}

	var i uint64
	for i = 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i]++
			n = n / i
		}
	}

	if n > 2 {
		pfs[n]++
	}
	return
}

func pow(x uint64, y uint64) uint64 {
	//_, _ = fmt.Printf("pow: %v, %v\n", x,y)
	z := x
	if y == 0 {
		return 1
	}
	var i uint64
	for i =1;i<y;i++ {
		z = z * x
	}
	return z
}

func incIdx(idx *[]uint64, max *[]uint64) bool {
	for i:=0;i<len(*idx);i++ {
		if (*idx)[i] < (*max)[i] {
			(*idx)[i]++
			return true
		} else {
			(*idx)[i] = 0
		}
	}
	return false
}

func answer(reader io.Reader, writer io.Writer) {
	var n uint64
	_, _ = fmt.Fscan(reader, &n)
	var sum uint64
	pfs := primeFactors(n)
	//_, _ = fmt.Printf("%v, %v\n", pfs, len(pfs))
	numPf := len(pfs)
	idx := make([]uint64, len(pfs))
	max := make([]uint64, len(pfs))
	listPf := make([]uint64, len(pfs))
	c := 0
	for k, v := range pfs {
		listPf[c] = k
		max[c] = v
		c++
	}

	for {
		var pf uint64
		pf = 1
		for i:=0;i<numPf;i++{
			pf = pf * pow(listPf[i],idx[i])
		}
		//_, _ = fmt.Printf("pf: %v\n", pf)
		if pf > 2 {
			if n / pf < pf && n/(pf-1) == n%(pf-1) {
				sum += pf - 1
				//_, _ = fmt.Printf("m: %v, %v, %v\n", pf-1, n/(pf-1), n%(pf-1))
			}
		}
		if !incIdx(&idx, &max) {
			break
		}
	}
	//_, _ = fmt.Printf("%v\n", sum)
	_, _ = fmt.Fprintf(writer, "%v", sum)
}

func main() {
	answer(os.Stdin, os.Stdout)
}
