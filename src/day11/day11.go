	package main

	import (
		"util"
		"strings"
		"fmt"
	)

	func getDistance(n int, ne int, se int) int{
		if n > 0{
			//n ne se
			if ne >= 0  && se >=0{
				return n+util.Max(ne,se)
			//n sw nw
			}else if ne < 0 && se < 0{
				return util.Max(n,-1*se)+util.Max(-1*ne,-1*se)
			//n ne nw
			}else if ne > 0 && se  <0{
				return n+util.Max(ne,-1*se)
			//n sw se
			}else{
				return util.Max(n,util.Abs(ne))+util.Max(util.Abs(se),util.Abs(ne))
			}

		}else {
			//s ne se
			if ne >= 0  && se >=0{
				return se+util.Max(util.Abs(n),ne)
				//s sw nw
			}else if ne < 0 && se < 0{
				return util.Max(util.Abs(n), util.Abs(se))+util.Max(util.Abs(ne), util.Abs(se))
				//s ne nw
			}else if ne > 0 && se  <0{
				return util.Max(util.Abs(n), util.Abs(ne))+util.Max(util.Abs(n), util.Abs(se))
				//s sw se
			}else{
				return util.Abs(n)+util.Max(util.Abs(ne), util.Abs(se))
			}
		}


	}

	func main() {
		lines := util.ReadFileLines("inputs/day11.txt")
		for _,line := range lines {
			max := 0
			commands := strings.Split(line, ",")
			n, ne, se := 0, 0, 0
			for _, command := range commands {
				switch command {
				case "n":
					n++
				case "ne":
					ne++
				case "se":
					se++
				case "s":
					n--
				case "sw":
					ne--
				case "nw":
					se--
				}
				max = util.Max(getDistance(n, ne, se),max)
			}

			fmt.Println("Part 1:", getDistance(n, ne, se))
			fmt.Println("Part 2:", max)
		}


	}
