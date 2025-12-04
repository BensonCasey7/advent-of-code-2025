package day2

import (
	"strconv"
	"strings"
)

func PartTwo() int {
	input := "8284583-8497825,7171599589-7171806875,726-1031,109709-251143,1039-2064,650391-673817,674522-857785,53851-79525,8874170-8908147,4197684-4326484,22095-51217,92761-107689,23127451-23279882,4145708930-4145757240,375283-509798,585093-612147,7921-11457,899998-1044449,3-19,35-64,244-657,5514-7852,9292905274-9292965269,287261640-287314275,70-129,86249864-86269107,5441357-5687039,2493-5147,93835572-94041507,277109-336732,74668271-74836119,616692-643777,521461-548256,3131219357-3131417388"
	ranges := strings.Split(input, ",")

	output := 0

	for _, rangeStr := range ranges {
		endpoints := strings.Split(rangeStr, "-")
		start, _ := strconv.Atoi(endpoints[0])
		end, _ := strconv.Atoi(endpoints[1])

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			strLength := len(str)

			for seqLength := strLength / 2; seqLength >= 1; seqLength-- {
				if strLength%seqLength == 0 {
					firstSequence := str[:seqLength]
					matchFound := true
					for seqStart := 0; seqStart < strLength; seqStart += seqLength {
						seqEnd := seqStart + seqLength
						if firstSequence != str[seqStart:seqEnd] {
							matchFound = false
							break
						}
					}

					if matchFound {
						output += i
						break
					}
				}
			}
		}
	}

	return output
}
