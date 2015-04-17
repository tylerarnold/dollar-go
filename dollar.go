package dollar
import "fmt"
import "math"

//
// port of $1 Unistroke Recogizer to the go program language
//  Tyler Arnold (tylernol@mac.com}, 2014
//

//constants
const NumUnistrokes = 16
const NumPoints = 64
const SquareSize = 250.0
var Origin = Point{0,0}
var Diagonal = math.Sqrt(SquareSize * SquareSize + SquareSize * SquareSize)
var HalfDiagonal = 0.5 * Diagonal
var AngleRange = Deg2Rad(45.0)
var AnglePrecision = Deg2Rad(2.0)
//golden ratio
var Phi = 0.5 * (-1.0 + math.Sqrt(5.0)) // Golden Ratio
type Point struct {
    X float64
    Y float64
}

type Rectangle struct {
	X float64
    Y float64
    Width float64
    Height float64
}

type Unistroke struct {
	Name string
	Points []Point
	Radians float64
	Vector []float64
} 

type Result struct {
	Name string
	Score float64
}

//constructor (TODO, method it)
func NewUnistroke(name string, points []Point) Unistroke{
	points = Resample(points,NumPoints)
	radians := IndicativeAngle(points)
	points = RotateBy(points, -radians)
	points = ScaleTo(points,SquareSize)
	points = TranslateTo(points,Origin)
	vector := Vectorize(points) 
	u := Unistroke{name,points,radians,vector}
	return u
}

//global variables



//API functions (turn these into methods)
func AddGesture(name string,points []Point,m map[string] Unistroke){
	unistroke := NewUnistroke(name,points) 
	m[name] = unistroke
}

//this.Recognize = function(points, useProtractor)
	
func Recognize(points []Point,useProtractor bool,m map[string] Unistroke) Result {
		points = Resample(points, NumPoints)
		radians := IndicativeAngle(points)
		points = RotateBy(points, -radians)
		points = ScaleTo(points, SquareSize)
		points = TranslateTo(points, Origin)
		vector := Vectorize(points) // for Protractor

		b := math.Inf(1)
	
		var bestUnistroke Unistroke;

		//TODO: closure this 
		var d float64
		for key, curUnistroke := range m {
			if (useProtractor) {
				d = OptimalCosineDistance(curUnistroke.Vector, vector)
				fmt.Printf("Recognize(pro)? unistroke %s distance %g\n",key,d)
			} else {
				// Golden Section Search (original $1) 
				d = DistanceAtBestAngle(points, curUnistroke, -AngleRange, +AngleRange, AnglePrecision)
				fmt.Printf("Recognize(gold)? unistroke %s distance %g\n",key,d)
			}
			if (d < b) {
				b = d // best (least) distance
				fmt.Printf("like %s more\n",curUnistroke.Name)
				bestUnistroke = curUnistroke 
			} else {
				fmt.Printf("like %s less\n",curUnistroke.Name)
			}
		}
		score := 0.0
		if (bestUnistroke.Name == ""){
			return Result{"No match.", score}
		} else {
			if (useProtractor){
				score = (1.0/b)
			} else {
				score = 1.0 - b / HalfDiagonal 
			}
			return Result{bestUnistroke.Name, score}
		}	
}


//private helper functions

//insert point into point at index i
//TODO: explore how to make this more efficient (change to linked list rather than slice)
//TODO make this a method for Point..?
func InsertPoint(i int, points []Point,point Point) []Point{
	if (i > len(points)+1){
		fmt.Printf("InsertPoint: bad index %d size is %d",i,len(points)+1)
	}
	t := make([]Point, len(points)+1) 		
	for j:= 0; j < i; j++ {
		t[j] = points[j]
	}
	t[i] = point
	for k:= i+1; k < len(t); k++ {
		t[k] = points[k-1]
	}
	return t
}



func Resample(points []Point, n int)[]Point{
	I := PathLength(points) / float64(n - 1) // interval length
	//fmt.Printf("I = %g, n = %d\n",I,n)
	D := 0.0
	newpoints := make([]Point, 1,len(points))
	newpoints[0] = points[0]
	for i:= 1; i < len(points); i++ {
		d := Distance(points[i - 1], points[i])
		//fmt.Printf("d = %g\n",d)
		if ((D + d) >= I){
			qx := points[i - 1].X + ((I - D) / d) * (points[i].X - points[i - 1].X)
			qy := points[i - 1].Y + ((I - D) / d) * (points[i].Y - points[i - 1].Y)
			q := Point{qx,qy}
			// append new point 'q'
			newpoints = append(newpoints,q)
			//insert q at index i, remove 0
			points = InsertPoint(i, points,q)
			D = 0.0
			//fmt.Printf("append/splice {%g, %g}\n",q.X,q.y)
		} else {
			D += d
			//fmt.Printf("D = %g\n",D)
		}
	}
	// somtimes we fall a rounding-error short of adding the last point, so add it if so
	if (len(newpoints) == n - 1) {
		q:= Point{points[len(newpoints) - 1].X, points[len(newpoints) - 1].Y}
		newpoints = append(newpoints,q)
		//newpoints[len(newpoints)] = Point{points[len(newpoints) - 1].X, points[len(newpoints) - 1].Y}
	}
	fmt.Printf("Resampled points size is %d\n",len(newpoints))
	return newpoints
}


func IndicativeAngle(points []Point) float64{
	c := Centroid(points)
	return math.Atan2(c.Y - points[0].Y, c.X - points[0].X)
}

//rotate points around centroid
func RotateBy(points []Point, radians float64) []Point{
	c := Centroid (points)
	cos := math.Cos(radians)
	sin := math.Sin(radians)
	qx  := 0.0
	qy  := 0.0
	newpoints := make([]Point, len(points))
	
	for i:= 0; i < len(points); i++ {
		qx = (points[i].X - c.X) * cos - (points[i].Y - c.Y) * sin + c.X
		qy = (points[i].X - c.X) * sin + (points[i].Y - c.Y) * cos + c.Y
		newpoints[i] = Point{qx,qy}
	}
	return newpoints
}
// non-uniform scale; assumes 2D gestures (i.e., no lines)
func ScaleTo(points []Point, size float64) []Point{
	B := BoundingBox(points)
	newpoints := make([]Point, len(points))
	for i:= 0; i < len(points); i++ {
		qx := points[i].X * (size / B.Width)
		qy := points[i].Y * (size / B.Height)
		newpoints[i] = Point{qx,qy}
	}
	return newpoints
}

func TranslateTo(points []Point, pt Point) []Point{
	c := Centroid(points)
	newpoints := make([]Point, len(points))
	qx  := 0.0
	qy  := 0.0
	for i:= 0; i < len(points); i++ {
		qx = points[i].X + pt.X - c.X;
		qy = points[i].Y + pt.Y - c.Y;
		newpoints[i] = Point{qx,qy}
	}
	return newpoints
}


func Vectorize(points []Point) []float64{
	sum := 0.0
	vector := make([]float64,0,len(points))
	for i:= 0; i < len(points); i++ {
		vector = append(vector,points[i].X)
		vector = append(vector,points[i].Y)
		sum += points[i].X * points[i].X + points[i].Y * points[i].Y
	}
	magnitude := math.Sqrt(sum)
	for i := 0; i < len(vector); i++ {
		vector[i] /= magnitude
	}
	return vector
}



func OptimalCosineDistance(v1 []float64, v2 []float64) float64{
	var a = 0.0;
	var b = 0.0;
	//verify that the lengths are equal
	if (len(v1) != len(v2)){
		fmt.Printf("OptimalCosineDistance: ERROR, vector size mismatch %d %d\n",len(v1),len(v2))
		return 0.0
	}
	for i:= 0; i < len(v1); i += 2 {
		a += v1[i] * v2[i] + v1[i + 1] * v2[i + 1]
        b += v1[i] * v2[i + 1] - v1[i + 1] * v2[i]
	}
	angle := math.Atan(b / a)
	return math.Acos(a*math.Cos(angle) + b*math.Sin(angle))
}

func DistanceAtBestAngle(points []Point, T Unistroke, a float64,b float64,threshold float64) float64{
	x1 := Phi * a + (1.0 - Phi) * b
	f1 := DistanceAtAngle(points, T, x1)
	x2 := (1.0 - Phi) * a + Phi * b
	f2 := DistanceAtAngle(points, T, x2)
	for math.Abs(b - a) > threshold {
		if (f1 < f2) {
			b = x2
			x2 = x1
			f2 = f1
			x1 = Phi * a + (1.0 - Phi) * b
			f1 = DistanceAtAngle(points, T, x1)
		} else {
			a  = x1
			x1 = x2
			f1 = f2
			x2 = (1.0 - Phi) * a + Phi * b
			f2 = DistanceAtAngle(points, T, x2)
		}	

	}
	return math.Min(f1,f2)
}

func DistanceAtAngle(points []Point, T Unistroke, radians float64) float64{
	newpoints := RotateBy(points, radians)
	return PathDistance(newpoints, T.Points)
}

func Centroid(points []Point) Point{
	x := 0.0
	y := 0.0
	for i:= 0; i < len(points); i++ {
		x += points[i].X
		y += points[i].Y	
	}
	x /= float64(len(points))
	y /= float64(len(points))
	p := Point{x,y}
	return p
}

func BoundingBox(points []Point) Rectangle{
	minX := math.Inf(1)
	maxX := math.Inf(-1)
	minY := math.Inf(1)
	maxY := math.Inf(-1)
	for i:= 0; i < len(points); i++ {
		minX = math.Min(minX, points[i].X);
		minY = math.Min(minY, points[i].Y);
		maxX = math.Max(maxX, points[i].X);
		maxY = math.Max(maxY, points[i].Y);	
	}
	r := Rectangle{minX,minY,maxX - minX,maxY - minY}
	return r
	//var minX = +Infinity, maxX = -Infinity, minY = +Infinity, maxY = -Infinity;
}

func PathDistance(Pts1, Pts2 []Point) float64{
	d := 0.0
	len1 := len(Pts1)
	len2 := len(Pts2)
	if (len1 == 0 || len2 == 0 || len1 != len2){
		fmt.Printf("bad parameters\n")
		return 0.0
	}
	for i:= 0; i < len(Pts1); i++ {
		d+= Distance(Pts1[i],Pts2[i])
	}
	return d / float64(len(Pts1))
}

func PathLength(Points []Point) float64{
	d := 0.0
	for i:= 1; i < len(Points); i++ {
		d += Distance(Points[i - 1], Points[i]); 
	}
	return d
}



func Distance(p1, p2 Point) float64{
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}



func Deg2Rad (d float64) float64 {
	return (d * math.Pi / 180.0)
}
//function Deg2Rad(d) { return (d * Math.PI / 180.0); }


