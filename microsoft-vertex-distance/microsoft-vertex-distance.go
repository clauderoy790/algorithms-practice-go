package microsoft_vertex_distance

import (
	"fmt"
	"log"
	"math"
	"sort"
)

var input = []Point{ {1,2}, {3,-1},{2,1},{2,3},{0,0} }
var vertex = Point{2,2}
var index = 2


func Resolve() {
	distances := calculateDistances()
	distMap := makeDistanceMap(input,distances)
	orderedPoints := sortByDistance(distMap)
	printMap(distMap,orderedPoints)

	if index < 0 || index >= len(distances) {
		log.Fatalf("error, invalid index: %v",index)
	} else {
		fmt.Println(fmt.Sprintf("the %v closest is %v from point %v",index,orderedPoints[index-1],vertex))
	}
}

func printMap(distMap map[Point]float64, keys []Point) {
	for _,k := range keys {
		fmt.Println(fmt.Sprintf("Key: %v, Value: %v",k,distMap[k]))
	}
}

func calculateDistances() []float64 {
	var distances []float64
	for _,point := range input {
		dist := distance(vertex,point)
		distances = append(distances,dist)
	}
	return distances
}

func makeDistanceMap(points []Point, distances []float64) map[Point]float64 {
	distMap := make(map[Point]float64)

	if len(points) != len(distances) {
		panic("lengths need to be the same")
	}


	for i,point := range points {
		distMap[point] = distances[i]
	}

	return distMap
}

func sortByDistance(distances map[Point]float64) []Point {
	var pairs []Pair
	//Split map into Pair list
	for k,v := range distances {
		pairs = append(pairs,Pair{k,v})
	}

	sort.Sort(byDistance(pairs))

	var orderedPoints []Point

	//add sorted pairs back into map
	for _,pair := range pairs {
		orderedPoints = append(orderedPoints,pair.point)
	}

	return orderedPoints
}

func distance(p1,p2 Point) float64 {
	xPow := math.Pow(float64(p1.X - p2.X),2)
	yPow := math.Pow(float64(p1.Y - p2.Y),2)
	return math.Sqrt(xPow + yPow)
}

type Point struct {
	X,Y int
}

type Pair struct {
	point Point
	distance float64
}

type byDistance []Pair

func (pair byDistance) Len() int {
	return len(pair)
}

func (pair byDistance) Less(i,j int) bool {
	return pair[i].distance <= pair[j].distance
}

func (pair byDistance) Swap(i, j int) {
	pair[i], pair[j] = pair[j], pair[i]
}

