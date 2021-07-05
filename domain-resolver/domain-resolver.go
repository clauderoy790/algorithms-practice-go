package domain_resolver

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input = []string{"900,google.com",
	"60,mail.yahoo.com",
	"10,mobile.sports.yahoo.com",
	"40,sports.yahoo.com",
	"300,yahoo.com",
	"10,stackoverflow.com",
	"20,overflow.com",
	"5,com.com",
	"2,en.wikipedia.org",
	"1,m.wikipedia.org",
	"1,mobile.sports",
	"1,google.co.uk"}

func Resolve() {
	fmt.Println("Resolving domains...")
	domainInfos := getDomainInfos(input)
	domainMap := buildDomainMap(domainInfos)
	pairList := sortResults(domainMap)
	printResults(pairList)
}

func printResults(pairs PairList) {
	for _,pair := range pairs {
		fmt.Println(fmt.Sprintf("domain: %v, has %v clicks.", pair.Key, pair.Value))
	}
}

func sortResults(domainMap map[string]int) PairList {
	pairs := make(PairList, len(domainMap))

	i := 0
	for k, v := range domainMap {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Sort(pairs)

	return pairs
}

func buildDomainMap(domainInfos []DomainInfo) map[string]int {
	domMap := make(map[string]int)

	for _, domInfo := range domainInfos {
		for _, domain := range domInfo.domains {
			if _, ok := domMap[domain]; !ok {
				domMap[domain] = 0
			}
			domMap[domain] += domInfo.clickCount
		}
	}

	return domMap
}

func getDomainInfos(domStrings []string) []DomainInfo {
	var domInfos []DomainInfo

	for _, str := range domStrings {
		info := getDomainInfo(str)
		domInfos = append(domInfos, info)
	}

	return domInfos
}

func getDomainInfo(input string) DomainInfo {
	domInfo := DomainInfo{}

	split := strings.Split(input,",")
	count, err := strconv.Atoi(split[0])

	if err != nil {
		panic(err)
	}
	domInfo.clickCount = count
	domInfo.domains = getDomains(split[1])


	return domInfo
}

func getDomains(fullDom string) []string {
	var doms []string

	split := strings.Split(fullDom,".")

	for i := 0; i<len(split);i++ {
		var domain string
		//Iterate to build domain
		for j := i;j < len(split);j++ {
			domain += split[j];

			//Add dot if we're not at the end
			if j < len(split)-1 {
				domain += "."
			}
		}
		doms = append(doms,domain)
	}

	return doms
}

type DomainInfo struct {
	clickCount int
	domains    []string
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair
func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }
