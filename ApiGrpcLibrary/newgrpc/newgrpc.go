package main

import (
	"fmt"
	"log"
	"os"

	pb "newgrpc/proto"
)

func createLogsDirectory() error {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return err
	}
	return nil
}

func openLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func setLogger(logFile *os.File) {
	log.SetOutput(logFile)
}

func logLine() {
	log.Println("-----------------------------")
}

func main() {
	err := createLogsDirectory()
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := openLogFile()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	setLogger(logFile)

	logLine()
	log.Println("Start newgrpc")

	activity := &pb.Activity{
		Id:          1,
		Description: "Some activity",
	}

	// timestamp := &pb.Timestamp{
	// 	Seconds: time.Now().Unix(),
	// 	Nanos:   int32(time.Now().Nanosecond()),
	// }

	// activity.Time = timestamp

	fmt.Println("Id:", activity.Id)
	// fmt.Println("Time:", activity.Time)
	fmt.Println("Description:", activity.Description)
}
