package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	var part1Valid, part2Valid int
	var currentPassport passport
	for _, line := range lines {
		if line == "" {
			if currentPassport.isValidPart1() {
				part1Valid++
			}
			if currentPassport.isValidPart2() {
				part2Valid++
			}
			currentPassport = passport{}
			continue
		}

		segments := strings.Split(line, " ")

		for _, seg := range segments {
			if strings.HasPrefix(seg, "byr") {
				currentPassport.byr = seg[4:]
				continue
			}

			if strings.HasPrefix(seg, "iyr") {
				currentPassport.iyr = seg[4:]
				continue
			}

			if strings.HasPrefix(seg, "eyr") {
				currentPassport.eyr = seg[4:]
				continue
			}

			if strings.HasPrefix(seg, "hgt") {
				currentPassport.hgt = seg[4:]
				continue
			}

			if strings.HasPrefix(seg, "hcl") {
				currentPassport.hcl = seg[4:]
				continue
			}

			if strings.HasPrefix(seg, "ecl") {
				currentPassport.ecl = seg[4:]
				continue
			}

			if strings.HasPrefix(seg, "pid") {
				currentPassport.pid = seg[4:]
				continue
			}
		}
	}

	fmt.Println("Part1:", part1Valid)
	fmt.Println("Part2:", part2Valid)
}

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (p *passport) isValidPart1() bool {
	if p.byr == "" ||
		p.iyr == "" ||
		p.eyr == "" ||
		p.hgt == "" ||
		p.hcl == "" ||
		p.ecl == "" ||
		p.pid == "" {
		return false
	}

	return true
}

func (p *passport) isValidPart2() bool {
	if p.byr == "" {
		return false
	}

	d, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}

	if d < 1920 || d > 2002 {
		return false
	}

	if p.iyr == "" {
		return false
	}

	d, err = strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}

	if d < 2010 || d > 2020 {
		return false
	}

	if p.eyr == "" {
		return false
	}

	d, err = strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}

	if d < 2020 || d > 2030 {
		return false
	}

	if p.pid == "" {
		return false
	}

	_, err = strconv.Atoi(p.pid)
	if err != nil {
		return false
	}

	if len(p.pid) != 9 {
		return false
	}

	if p.hgt == "" {
		return false
	}

	if strings.HasSuffix(p.hgt, "cm") {
		d, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm"))
		if err != nil {
			return false
		}

		if d < 150 || d > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		d, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in"))
		if err != nil {
			return false
		}

		if d < 59 || d > 76 {
			return false
		}
	} else {
		return false
	}

	match, err := regexp.MatchString("#[a-f0-9]{6}", p.hcl)
	if err != nil || !match {
		return false
	}

	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}

	return true
}
