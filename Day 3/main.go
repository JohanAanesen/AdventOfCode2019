package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	wire1 := []string{"R1008","U336","R184","D967","R451","D742","L235","U219","R57","D439","R869","U207","L574","U670","L808","D675","L203","D370","L279","U448","L890","U297","R279","D613","L411","D530","L372","D88","R986","U444","R319","D95","L385","D674","R887","U855","R794","U783","R633","U167","L587","D545","L726","D196","R681","U609","R677","U881","R153","D724","L63","U246","R343","U315","R580","U872","L516","U95","R463","D809","R9","U739","R540","U670","L434","D699","L158","U47","L383","D483","L341","U61","R933","D269","R816","D589","R488","D169","R972","D534","L995","D277","L887","D657","R628","D322","R753","U813","L284","D237","R676","D880","L50","D965","L401","D619","R858","U313","L156","U535","R664","U447","L251","U168","L352","U881","L734","U270","L177","D903","L114","U998","L102","D149","R848","D586","L98","D157","R942","U496","R857","U362","R398","U86","R469","U358","L721","D631","R176","D365","L112","U472","L557","D153","R97","D639","L457","U566","R570","U106","R504","D292","L94","U499","R358","U653","L704","D627","R544","D24","L407","U513","R28","U643","L510","U579","R825","D376","L867","U999","R134","D734","R654","D989","L920","U872","R64","U626","R751","D425","R620","U274","L471","D83","R979","U577","L43","D320","R673","D187","R300","U134","L451","D717","R857","U576","R570","U988","R745","U840","R799","U809","R573","U354","L208","D976","L417","U473","L555","U563","L955","U823","R712","D869","L145","D735","L780","D74","R421","U42","L158","U689","R718","D455","L670","U128","L744","U401","R149","U102","L122","U832","R872","D40","R45","D325","L553","U980","L565","D497","L435","U647","L209","D822","L593","D28","R936","U95","R349","U511","L243","U895","R421","U336","L986","U264","R376","D183","R480","D947","R416","D706","R118","D799","R424","D615","R384","U185","L338","U14","R576","D901","L734","D417","L62","D254","R784","D973","R987","D848","R32","D72","L535","D633","L668","D664","R308","D474","L418","D39","L473","U388","L518","D544","R118","D948","L844","D956","R605","U14","L948","D78","L689","U443","L996","U932","R81","D879","R556","D633","R131"}
	wire2 := []string{"L993","U227","L414","U228","L304","U53","R695","U765","R162","D264","L530","U870","R771","D395","R27","D200","L235","D834","L559","D128","R284","U912","L959","U358","R433","U404","L539","U799","R271","D734","L104","U261","R812","D15","L474","U887","R606","U366","L694","U156","R385","D667","L329","D434","R745","U776","L319","U756","R208","D457","R705","U999","R284","U98","L657","U214","R639","D937","R675","U444","L891","D587","L914","D4","R294","D896","R534","D584","L887","U878","L807","U202","R505","U234","L284","D5","R667","U261","L127","D482","R777","D223","L707","D468","L606","U345","L509","D967","R437","U995","R28","D376","L2","D959","L814","U702","L38","D154","L79","D439","L259","U143","L376","U700","R894","U165","L300","D58","R631","D47","R684","U935","R262","D857","R797","D599","L705","U792","R439","D444","R398","D887","L81","D40","R671","U332","L820","U252","R691","U412","L794","D661","R810","U157","R858","U566","R892","D543","R10","D725","L653","D812","R733","D804","R816","U862","R994","U221","L33","U271","R766","D591","R575","D970","R152","D693","L916","U404","L658","D847","L605","D433","L583","U587","L129","D103","R407","U780","L901","D676","L846","D687","L9","D47","R295","D597","L808","U134","L186","D676","L62","U305","L73","D369","L468","U30","L472","U280","L413","U961","L98","D966","R308","D178","L21","D789","L871","D671","R665","U927","L906","U633","L135","D894","R110","D205","R324","D665","R143","D450","L978","U385","R442","D853","L518","U542","R211","U857","R119","D872","L246","U380","L874","U463","R153","U982","R832","D784","L652","U545","R71","U386","R427","D827","R986","U870","R959","U232","R509","U675","R196","U389","R944","U149","R152","U571","R527","U495","L441","U511","L899","D996","L707","D455","L358","D423","L14","D427","R144","D703","L243","U157","R876","D538","R26","D577","L385","U622","L149","D852","L225","U475","L811","D520","L226","U523","L338","D79","R565","U766","L609","U496","L189","D446","R63","U396","L629","U312","L841","D639","R466","U808","L60","D589","L146","U114","R165","U96","L809","D704","L61"}


	type coordinate struct {
		x int
		y int
	}

	w1Path := make([]coordinate, 1)
	w2Path := make([]coordinate, 1)

	w1Path[0] = coordinate{x:0, y:0}
	w2Path[0] = coordinate{x:0, y:0}

	prevCor1 := coordinate{x:0, y:0}
	prevCor2 := coordinate{x:0, y:0}

	for _, w1 := range wire1 {
		if strings.HasPrefix(w1, "R"){
			//right
			nrString := strings.TrimPrefix(w1, "R")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w1Path = append(w1Path, coordinate{x:prevCor1.x+i, y:prevCor1.y})
			}

		} else if strings.HasPrefix(w1, "U") {
			//up
			nrString := strings.TrimPrefix(w1, "U")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w1Path = append(w1Path, coordinate{x:prevCor1.x, y:prevCor1.y+i})
			}
		} else if strings.HasPrefix(w1, "L") {
			//left
			nrString := strings.TrimPrefix(w1, "L")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w1Path = append(w1Path, coordinate{x:prevCor1.x-i, y:prevCor1.y})
			}
		} else if strings.HasPrefix(w1, "D") {
			//down
			nrString := strings.TrimPrefix(w1, "D")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w1Path = append(w1Path, coordinate{x:prevCor1.x, y:prevCor1.y-i})
			}
		}

		prevCor1 = w1Path[len(w1Path)-1]
	}

	for _, w2 := range wire2 {
		if strings.HasPrefix(w2, "R"){
			//right
			nrString := strings.TrimPrefix(w2, "R")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w2Path = append(w2Path, coordinate{x:prevCor2.x+i, y:prevCor2.y})
			}

		} else if strings.HasPrefix(w2, "U") {
			//up
			nrString := strings.TrimPrefix(w2, "U")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w2Path = append(w2Path, coordinate{x:prevCor2.x, y:prevCor2.y+i})
			}
		} else if strings.HasPrefix(w2, "L") {
			//left
			nrString := strings.TrimPrefix(w2, "L")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w2Path = append(w2Path, coordinate{x:prevCor2.x-i, y:prevCor2.y})
			}
		} else if strings.HasPrefix(w2, "D") {
			//down
			nrString := strings.TrimPrefix(w2, "D")

			nr, _ := strconv.Atoi(nrString)

			for i := 1; i <= nr; i++ {
				w2Path = append(w2Path, coordinate{x:prevCor2.x, y:prevCor2.y-i})
			}
		}

		prevCor2 = w2Path[len(w2Path)-1]
	}

	len := 99999999
	steps := 99999999

	for i, w1 := range w1Path{
		for j, w2 := range w2Path{
			if i > 2 && j > 2 {
				if w1.x == w2.x && w1.y == w2.y {
					temp := 0

					if w1.x < 0 {
						temp += w1.x * -1
					} else {
						temp += w1.x
					}

					if w1.y < 0 {
						temp += w1.y * -1
					} else {
						temp += w1.y
					}

					fmt.Println(w1)
					fmt.Println(temp)

					if temp < len {
						len = temp
					}

					if i+j < steps {
						steps = i+j
					}
				}
			}
		}
	}

	//Part1: len, part2: steps
	fmt.Println(len, steps)


	os.Exit(0) //quit



}
