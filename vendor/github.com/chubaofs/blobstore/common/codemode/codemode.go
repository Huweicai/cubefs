package codemode

import "fmt"

// CodeMode EC encode and decode mode
type CodeMode uint8

// pre-defined mode
const (
	MinCodeMode CodeMode = iota
	EC15P12
	EC6P6
	EC16P20L2
	EC6P10L2
	EC6P3L3
	EC6P6Align0
	EC6P6Align512
	EC4P4L2
	EC12P4
	EC16P4
	EC3P3
	MaxCodeMode
)

// Note: Don't modify it unless you know very well how codemode works.
const (
	// align size per shard
	alignSize0B   = 0    // 0B
	alignSize512B = 512  // 512B
	alignSize2KB  = 2048 // 2KB
)

// The tactic is fixed pairing with one codemode.
// Add a new codemode if you want other features.
var constCodeModeTactic = map[CodeMode]Tactic{
	EC15P12:   {N: 15, M: 12, L: 0, AZCount: 3, PutQuorum: 24, GetQuorum: 0, MinShardSize: alignSize2KB},
	EC6P6:     {N: 06, M: 06, L: 0, AZCount: 3, PutQuorum: 11, GetQuorum: 0, MinShardSize: alignSize2KB},
	EC16P20L2: {N: 16, M: 20, L: 2, AZCount: 2, PutQuorum: 34, GetQuorum: 0, MinShardSize: alignSize2KB},
	EC6P10L2:  {N: 06, M: 10, L: 2, AZCount: 2, PutQuorum: 14, GetQuorum: 0, MinShardSize: alignSize2KB},

	// single az
	EC12P4: {N: 12, M: 04, L: 0, AZCount: 1, PutQuorum: 15, GetQuorum: 0, MinShardSize: alignSize2KB},
	EC16P4: {N: 16, M: 04, L: 0, AZCount: 1, PutQuorum: 19, GetQuorum: 0, MinShardSize: alignSize2KB},
	EC3P3:  {N: 3, M: 03, L: 0, AZCount: 1, PutQuorum: 5, GetQuorum: 0, MinShardSize: alignSize2KB},
	// for env test
	EC6P3L3:       {N: 06, M: 03, L: 3, AZCount: 3, PutQuorum: 9, GetQuorum: 0, MinShardSize: alignSize2KB},
	EC6P6Align0:   {N: 06, M: 06, L: 0, AZCount: 3, PutQuorum: 11, GetQuorum: 0, MinShardSize: alignSize0B},
	EC6P6Align512: {N: 06, M: 06, L: 0, AZCount: 3, PutQuorum: 11, GetQuorum: 0, MinShardSize: alignSize512B},
	EC4P4L2:       {N: 04, M: 04, L: 2, AZCount: 2, PutQuorum: 6, GetQuorum: 0, MinShardSize: alignSize2KB},
}

//vol layout ep:EC6P10L2
//|----N------|--------M----------------|--L--|
//|0,1,2,3,4,5|6,7,8,9,10,11,12,13,14,15|16,17|

//global stripe:[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15], n=6 m=10
//two local stripes:
//local stripe1:[0,1,2,  6, 7, 8, 9,10, 16] n=8 m=1
//local stripe2:[3,4,5, 11,12,13,14,15, 17] n=8 m=1

// Tactic constant strategy of one CodeMode
type Tactic struct {
	N int
	M int
	// local parity count
	L int
	// the count of AZ, access use this for split data shards and parity shards
	AZCount int

	// PutQuorum write quorum,
	// MUST make sure that ec data is recoverable if one AZ was down
	// We SHOULD ignore the local shards
	// (N + M) / AZCount + N <= PutQuorum <= M + N
	PutQuorum int

	// get quorum config
	GetQuorum int

	// MinShardSize min size per shard, fill data into shards 0-N continuously,
	// align with zero bytes if data size less than MinShardSize*N
	//
	// length of data less than MinShardSize*N, size of per shard = MinShardSize
	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//  |  data  |                 align zero bytes                     |
	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//  |    0    |    1    |    2    |   ....                |    N    |
	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//
	// length of data more than MinShardSize*N, size of per shard = ceil(len(data)/N)
	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//  |                           data                        |padding|
	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//  |    0    |    1    |    2    |   ....                |    N    |
	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	MinShardSize int
}

func init() {
	// assert all codemode
	for _, pair := range []struct {
		Mode CodeMode
		Size int
	}{
		{Mode: EC15P12, Size: alignSize2KB},
		{Mode: EC6P6, Size: alignSize2KB},
		{Mode: EC16P20L2, Size: alignSize2KB},
		{Mode: EC6P10L2, Size: alignSize2KB},

		{Mode: EC6P3L3, Size: alignSize2KB},
		{Mode: EC6P6Align0, Size: alignSize0B},
		{Mode: EC6P6Align512, Size: alignSize512B},
		//{Mode: EC4P4L2, Size: alignSize2KB},
	} {
		tactic := GetTactic(pair.Mode)
		if !tactic.IsValid() {
			panic(fmt.Sprintf("Invalid codemode:%d Tactic:%+v", pair.Mode, tactic))
		}

		min := tactic.N + (tactic.N+tactic.M)/tactic.AZCount
		max := tactic.N + tactic.M
		if tactic.PutQuorum < min || tactic.PutQuorum > max {
			panic(fmt.Sprintf("Invalid codemode:%d PutQuorum:%d([%d,%d])", pair.Mode,
				tactic.PutQuorum, min, max))
		}

		if tactic.MinShardSize != pair.Size {
			panic(fmt.Sprintf("Invalid codemode:%d MinShardSize:%d(%d)", pair.Mode,
				tactic.MinShardSize, pair.Size))
		}
	}
}

// Tactic returns its constant tactic
func (c CodeMode) Tactic() Tactic {
	return GetTactic(c)
}

// IsValid returns CodeMode is valid or not
func IsValid(mode CodeMode) bool {
	return mode > MinCodeMode && mode < MaxCodeMode
}

// GetShardNum L+M+N of this CodeMode
func GetShardNum(mode CodeMode) int {
	tactic := GetTactic(mode)
	return tactic.L + tactic.M + tactic.N
}

// GetTactic get tactic of this CodeMode
// panic if the mode is invalid
func GetTactic(mode CodeMode) Tactic {
	if tactic, ok := constCodeModeTactic[mode]; ok {
		return tactic
	}

	panic(fmt.Sprintf("Invalid codemode:%d defined:(%d, %d)", mode, MinCodeMode, MaxCodeMode))
}

// IsValid ec tactic valid or not
func (c *Tactic) IsValid() bool {
	return c.N > 0 && c.M > 0 && c.L >= 0 && c.AZCount > 0 &&
		c.PutQuorum > 0 && c.GetQuorum >= 0 && c.MinShardSize >= 0 &&
		c.N%c.AZCount == 0 && c.M%c.AZCount == 0 && c.L%c.AZCount == 0
}

// GetECLayoutByAZ ec layout by AZ
func (c *Tactic) GetECLayoutByAZ() (azStripes [][]int) {
	azStripes = make([][]int, c.AZCount)
	n, m, l := c.N/c.AZCount, c.M/c.AZCount, c.L/c.AZCount
	for idx := range azStripes {
		stripe := make([]int, 0, n+m+l)
		for i := 0; i < n; i++ {
			stripe = append(stripe, idx*n+i)
		}
		for i := 0; i < m; i++ {
			stripe = append(stripe, c.N+idx*m+i)
		}
		for i := 0; i < l; i++ {
			stripe = append(stripe, c.N+c.M+idx*l+i)
		}
		azStripes[idx] = stripe
	}
	return azStripes
}

// GlobalStripe returns initial stripe
func (c *Tactic) GlobalStripe() (indexes []int, n, m int) {
	indexes = make([]int, c.N+c.M)
	for i := 0; i < c.N+c.M; i++ {
		indexes[i] = i
	}
	return indexes, c.N, c.M
}

// AllLocalStripe returns all local stripes
func (c *Tactic) AllLocalStripe() (stripes [][]int, n, m int) {
	if c.L == 0 {
		return
	}

	n, m, l := c.N/c.AZCount, c.M/c.AZCount, c.L/c.AZCount
	return c.GetECLayoutByAZ(), n + m, l
}

// LocalStripe get local stripe by index
func (c *Tactic) LocalStripe(index int) (localStripe []int, n, m int) {
	if c.L == 0 {
		return nil, 0, 0
	}

	n, m, l := c.N/c.AZCount, c.M/c.AZCount, c.L/c.AZCount
	var azIdx int
	if index < c.N {
		azIdx = index / n
	} else if index < c.N+c.M {
		azIdx = (index - c.N) / m
	} else if index < c.N+c.M+c.L {
		azIdx = (index - c.N - c.M) / l
	} else {
		return nil, 0, 0
	}

	return c.LocalStripeInAZ(azIdx)
}

// LocalStripeInAZ get local stripe in az index
func (c *Tactic) LocalStripeInAZ(azIndex int) (localStripe []int, n, m int) {
	if c.L == 0 {
		return nil, 0, 0
	}

	n, m, l := c.N/c.AZCount, c.M/c.AZCount, c.L/c.AZCount
	azStripes := c.GetECLayoutByAZ()
	if azIndex < 0 || azIndex >= len(azStripes) {
		return nil, 0, 0
	}
	return azStripes[azIndex][:], n + m, l
}
