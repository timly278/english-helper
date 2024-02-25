package session

import (
	"fmt"
	"github/timly278/english-helper/clock"
	"github/timly278/english-helper/service"
	"os"
	"os/exec"
)

func DoDescribeImage(c *Config) {
	// Read all file names in the directory
	err := OpenFile(c.DescribeImage.Path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n************************************************************\n")
	fmt.Println("\tWelcome to Describe Image session")
	fmt.Printf("*************************************************************\n")

	fmt.Println("the image has been opened successfully!")
	AskForReady()
	clock.StartTimer(c.DescribeImage.Duration)
	fmt.Println("Finish!")
	clock.Sound(c.SoundPath)
}

func OpenFile(dirpath string) error {
	files, err := os.ReadDir(dirpath)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	// Print the names of all files in the directory
	fmt.Println("Files in directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	randIndex := service.RandomInt(0, len(files)-1)
	filepath := dirpath + "/" + files[randIndex].Name()
	cmd := exec.Command("open", filepath)
	_, err = cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
