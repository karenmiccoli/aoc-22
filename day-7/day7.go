package day_7

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Files map[string]*Directory

type Directory struct {
	Name     string
	Size     int
	Children Files
}

func FindFileSizeUnder100000(filename string, part int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return findFileSize(lines, part)
}

func findFileSize(lines []string, part int) int {
	//Make root dir for "/" to help calculate at the end
	currentDir := &Directory{Name: "/", Children: Files{}}
	rootDir := currentDir

	for _, line := range lines[1:] {
		switch {
		case strings.HasPrefix(line, "$ cd"):
			//It's a command
			destination := line[5:]
			currentDir = currentDir.Children[destination]
		case strings.HasPrefix(line, "dir "):
			//It's a directory
			dirName := line[4:]
			currentDir.Children[dirName] = &Directory{Name: dirName, Children: Files{"..": currentDir}}

		case regexp.MustCompile(`\d`).MatchString(line):
			fileData := strings.Split(line, " ")
			fileName := fileData[1]
			fileSize, err := strconv.Atoi(fileData[0])
			if err != nil {
				log.Fatalf(err.Error(), "error converting string to int")
			}
			currentDir.Children[fileName] = &Directory{Name: fileName, Size: fileSize}
		}
	}

	totalSize := 0
	addSizesToDirs(rootDir)

	freeSpace := 70000000 - rootDir.Size
	spaceNeeded := 30000000 - freeSpace
	minSize := 1<<32 - 1

	calculateSize(rootDir, &totalSize, spaceNeeded, &minSize)
	if part == 1 {
		return totalSize
	}

	return minSize
}

func addSizesToDirs(dir *Directory) int {
	if len(dir.Children) == 0 {
		return dir.Size
	}

	fileSize := 0
	for i, file := range dir.Children {
		if i == ".." {
			continue
		}

		fileSize += addSizesToDirs(file)
	}

	dir.Size = fileSize

	return fileSize
}

func calculateSize(dir *Directory, total *int, spaceNeeded int, minSize *int) {
	if len(dir.Children) > 0 {
		if dir.Size <= 100000 {
			*total += dir.Size
		}

		if dir.Size >= spaceNeeded && dir.Size < *minSize {
			*minSize = dir.Size
		}
	}

	for i, file := range dir.Children {
		if i == ".." {
			continue
		}

		calculateSize(file, total, spaceNeeded, minSize)
	}
}
