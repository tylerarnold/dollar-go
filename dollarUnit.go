package main
import "fmt"
import "dollar"

func InsertPointUnitTest(){
	p0 := dollar.Point{0,0}
	p1 := dollar.Point{3,4}
	p2 := dollar.Point{6,8}
	p3 := dollar.Point{9,12}
	p4 := dollar.Point{4,5}
	points := []dollar.Point{p0,p1,p2}

	points = dollar.InsertPoint(3,points,p3)
	for i:= 0; i < len(points); i++ {
		//fmt.Printf("point(%d)is {%g,%g}\n",i,points[i].x,points[i].y)
	}
	points = dollar.InsertPoint(1,points,p4)
	for i:= 0; i < len(points); i++ {
		fmt.Printf("point(%d)is {%g,%g}\n",i,points[i].X,points[i].Y)
	}
}

func ResampleUnitTest(){
	p1 := dollar.Point{0,0}
	p2 := dollar.Point{4,0}
	p3 := dollar.Point{8,0}
	p4 := dollar.Point{12,0}
	points := []dollar.Point{p1,p2,p3,p4}

	//should stay the same
	rp:= dollar.Resample(points,4)
	for i:= 0; i < len(rp); i++ {
		if (rp[i].X != points[i].X || rp[i].X != points[i].X){
			fmt.Printf("ResampleUnitTest: failure\n")
			return
		}
	} 
	//decimate
	rp = dollar.Resample(points,2)
	if (len(rp) != 2){
		fmt.Printf("ResampleUnitTest: failure\n")
		return
	}
	for i:= 0; i < len(rp); i++ {
		fmt.Printf("point(%d)is {%g,%g}\n",i,rp[i].X,rp[i].Y)
	}
	fmt.Printf("ResampleUnitTest: OK\n")

}

func DistanceUnitTest(){
	p1 := dollar.Point{0,0}
	p2 := dollar.Point{3,4}
	d := dollar.Distance(p1,p2)
	//expect 5
	if (d != 5){
		fmt.Printf("DistanceUnitTest: Error, expected 5, got %g\n", d)
	} else {
		fmt.Printf("DistanceUnitTest:OK\n")
	}

	//fmt.Printf("distance is %g.",d)
}

func PathLengthUnitTest(){
	p1 := dollar.Point{0,0}
	p2 := dollar.Point{3,4}
	p3 := dollar.Point{6,8}
	points := []dollar.Point{p1,p2,p3}
	
	d1 := dollar.PathLength(points[:2])
	//expect 5
	if (d1 != 5){
		fmt.Printf("PathLengthUnitTest: Error, expected 5, got %g\n", d1)
	} else {
		fmt.Printf("PathLengthUnitTest:OK\n")
	}
	//expect 10
	d2 := dollar.PathLength(points)
	if (d2 != 10){
		fmt.Printf("PathLengthUnitTest: Error, expected 10, got %g\n", d2)
	} else {
		fmt.Printf("PathLengthUnitTest:OK\n")
	}
}

func VectorizeUnitTest(){
	p1 := dollar.Point{0,0}
	p2 := dollar.Point{3,4}
	p3 := dollar.Point{6,8}
	points := []dollar.Point{p1,p2,p3}
	v := dollar.Vectorize(points)
	for i:= 0; i < len(v); i++ {
		fmt.Printf("v(%d) = %g\n",i,v[i])
	}
	
}

func UnitTests (){
	InsertPointUnitTest()
	ResampleUnitTest()
	DistanceUnitTest()
	PathLengthUnitTest()
	VectorizeUnitTest()
}



//pre-defined gestures
//var trianglePoints []Point = []Point{Point{0,0},Point{125,250},Point{250,0},Point{0,0}}
//var   squarePoints []Point = []Point{Point{0,0},Point{125,0},Point{125,125},Point{125,0},Point{0,0}}


var   squarePoints []dollar.Point = []dollar.Point{dollar.Point{78,149},dollar.Point{78,153},dollar.Point{78,157},dollar.Point{78,160},dollar.Point{79,162},dollar.Point{79,164},dollar.Point{79,167},dollar.Point{79,169},dollar.Point{79,173},dollar.Point{79,178},dollar.Point{79,183},dollar.Point{80,189},dollar.Point{80,193},dollar.Point{80,198},dollar.Point{80,202},dollar.Point{81,208},dollar.Point{81,210},dollar.Point{81,216},dollar.Point{82,222},dollar.Point{82,224},dollar.Point{82,227},dollar.Point{83,229},dollar.Point{83,231},dollar.Point{85,230},dollar.Point{88,232},dollar.Point{90,233},dollar.Point{92,232},dollar.Point{94,233},dollar.Point{99,232},dollar.Point{102,233},dollar.Point{106,233},dollar.Point{109,234},dollar.Point{117,235},dollar.Point{123,236},dollar.Point{126,236},dollar.Point{135,237},dollar.Point{142,238},dollar.Point{145,238},dollar.Point{152,238},dollar.Point{154,239},dollar.Point{165,238},dollar.Point{174,237},dollar.Point{179,236},dollar.Point{186,235},dollar.Point{191,235},dollar.Point{195,233},dollar.Point{197,233},dollar.Point{200,233},dollar.Point{201,235},dollar.Point{201,233},dollar.Point{199,231},dollar.Point{198,226},dollar.Point{198,220},dollar.Point{196,207},dollar.Point{195,195},dollar.Point{195,181},dollar.Point{195,173},dollar.Point{195,163},dollar.Point{194,155},dollar.Point{192,145},dollar.Point{192,143},dollar.Point{192,138},dollar.Point{191,135},dollar.Point{191,133},dollar.Point{191,130},dollar.Point{190,128},dollar.Point{188,129},dollar.Point{186,129},dollar.Point{181,132},dollar.Point{173,131},dollar.Point{162,131},dollar.Point{151,132},dollar.Point{149,132},dollar.Point{138,132},dollar.Point{136,132},dollar.Point{122,131},dollar.Point{120,131},dollar.Point{109,130},dollar.Point{107,130},dollar.Point{90,132},dollar.Point{81,133},dollar.Point{76,133}}
var trianglePoints []dollar.Point = []dollar.Point{dollar.Point{137,139}, dollar.Point{135,141}, dollar.Point{133,144}, dollar.Point{132,146}, dollar.Point{130,149}, dollar.Point{128,151}, dollar.Point{126,155}, dollar.Point{123,160}, dollar.Point{120,166}, dollar.Point{116,171}, dollar.Point{112,177}, dollar.Point{107,183}, dollar.Point{102,188}, dollar.Point{100,191}, dollar.Point{95,195}, dollar.Point{90,199}, dollar.Point{86,203}, dollar.Point{82,206}, dollar.Point{80,209}, dollar.Point{75,213}, dollar.Point{73,213}, dollar.Point{70,216}, dollar.Point{67,219}, dollar.Point{64,221}, dollar.Point{61,223}, dollar.Point{60,225}, dollar.Point{62,226}, dollar.Point{65,225}, dollar.Point{67,226}, dollar.Point{74,226}, dollar.Point{77,227}, dollar.Point{85,229}, dollar.Point{91,230}, dollar.Point{99,231}, dollar.Point{108,232}, dollar.Point{116,233}, dollar.Point{125,233}, dollar.Point{134,234}, dollar.Point{145,233}, dollar.Point{153,232}, dollar.Point{160,233}, dollar.Point{170,234}, dollar.Point{177,235}, dollar.Point{179,236}, dollar.Point{186,237}, dollar.Point{193,238}, dollar.Point{198,239}, dollar.Point{200,237}, dollar.Point{202,239}, dollar.Point{204,238}, dollar.Point{206,234}, dollar.Point{205,230}, dollar.Point{202,222}, dollar.Point{197,216}, dollar.Point{192,207}, dollar.Point{186,198}, dollar.Point{179,189}, dollar.Point{174,183}, dollar.Point{170,178}, dollar.Point{164,171}, dollar.Point{161,168}, dollar.Point{154,160}, dollar.Point{148,155}, dollar.Point{143,150}, dollar.Point{138,148}, dollar.Point{136,148}}
var xPoints []dollar.Point = []dollar.Point{dollar.Point{87,142},dollar.Point{89,145},dollar.Point{91,148},dollar.Point{93,151},dollar.Point{96,155},dollar.Point{98,157},dollar.Point{100,160},dollar.Point{102,162},dollar.Point{106,167},dollar.Point{108,169},dollar.Point{110,171},dollar.Point{115,177},dollar.Point{119,183},dollar.Point{123,189},dollar.Point{127,193},dollar.Point{129,196},dollar.Point{133,200},dollar.Point{137,206},dollar.Point{140,209},dollar.Point{143,212},dollar.Point{146,215},dollar.Point{151,220},dollar.Point{153,222},dollar.Point{155,223},dollar.Point{157,225},dollar.Point{158,223},dollar.Point{157,218},dollar.Point{155,211},dollar.Point{154,208},dollar.Point{152,200},dollar.Point{150,189},dollar.Point{148,179},dollar.Point{147,170},dollar.Point{147,158},dollar.Point{147,148},dollar.Point{147,141},dollar.Point{147,136},dollar.Point{144,135},dollar.Point{142,137},dollar.Point{140,139},dollar.Point{135,145},dollar.Point{131,152},dollar.Point{124,163},dollar.Point{116,177},dollar.Point{108,191},dollar.Point{100,206},dollar.Point{94,217},dollar.Point{91,222},dollar.Point{89,225},dollar.Point{87,226},dollar.Point{87,224}}

//map of all strokes
var m map[string] dollar.Unistroke

func main() {

	UnitTests()

	//TODO: put this in method/struct init:
	m = make(map[string] dollar.Unistroke)
	//add some basic gestures	
	dollar.AddGesture("Triangle",trianglePoints,m)
	dollar.AddGesture("Square",squarePoints,m)
	dollar.AddGesture("X",xPoints,m)
	//protractor algorithm tests
	//should get 100% back if we throw the triangle back at it

	result := dollar.Recognize(trianglePoints,true,m)
	fmt.Printf("Trying Triangle...\n\n")
	if (result.Name  == "Triangle"){
		fmt.Printf("Recognize: PASS\n")
	} else {
		fmt.Printf("Recognize: FAIL\n")
		fmt.Printf("results = %s %g\n",result.Name ,result.Score)
	}
	result = dollar.Recognize(trianglePoints,false,m)
	fmt.Printf("Trying Triangle...\n\n")
	if (result.Name  == "Triangle"){
		fmt.Printf("Recognize: PASS\n")
	} else {
		fmt.Printf("Recognize: FAIL\n")
		fmt.Printf("results = %s %g\n",result.Name ,result.Score)
	}
	fmt.Printf("Trying square...\n\n")
	result = dollar.Recognize(squarePoints,true,m)
	if (result.Name  == "Square"){
		fmt.Printf("Recognize: PASS\n")
	} else {
		fmt.Printf("Recognize: FAIL\n")
		fmt.Printf("results = %s %g\n",result.Name ,result.Score)
	}
	fmt.Printf("Trying square...\n\n")
	result = dollar.Recognize(squarePoints,false,m)
	if (result.Name  == "Square"){
		fmt.Printf("Recognize: PASS\n")
	} else {
		fmt.Printf("Recognize: FAIL\n")
		fmt.Printf("results = %s %g\n",result.Name ,result.Score)
	}
	fmt.Printf("Trying X...\n\n")
	result = dollar.Recognize(xPoints,false,m)
	if (result.Name  == "X"){
		fmt.Printf("Recognize: PASS\n")
	} else {
		fmt.Printf("Recognize: FAIL\n")
		fmt.Printf("results = %s %g\n",result.Name ,result.Score)
	}

}