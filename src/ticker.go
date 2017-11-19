package main

import (
	"fmt"
	"os"
	"./specparser"
	"time"
)

func main() {
	var lookAheadMins int = 10
	clock := new(specparser.ClockInterface)
	command := "1-15,42-46,55,57,59 * * * * /scripts/runBackup.sh"
	run(command, clock, lookAheadMins)
}

func run(spec string, clock *specparser.ClockInterface, lookAheadMins int) {
	for {
		var err error

		gotError := func(err error) {
			fmt.Println(err)
			fmt.Println("Quitting...")
			os.Exit(255)
		}

		startTime := clock.Now().Add(time.Duration(time.Second * 5)) // adding the offset inside the loop will cause skipped time slot issues
		// test for races

		fmt.Println("offset:", startTime.Second(), "seconds past minute")

		var taskSpec specparser.TaskSpec
		var taskList specparser.TaskList

		if taskSpec, err = specparser.NewTaskSpec(spec); err != nil {
			gotError(err)
			return
		}

		if taskList, err = specparser.NewTaskList(taskSpec, startTime, lookAheadMins); err != nil {
			gotError(err)
			return
		}

		if len(taskList.Schedule) < 1 {
			fmt.Println("No work...", startTime.Format("15:04:05"), "-", startTime.Add(time.Minute*time.Duration(10)).Format("15:04:05"))
		} else {
			fmt.Println("Jobs:", len(taskList.Schedule), "\n")
			doWork(taskList, 0, clock)
		}

		remainingTime := clock.Until(startTime.Add(time.Minute * time.Duration(10)))

		if remainingTime > 0 {
			fmt.Println(" ...sleeping...", remainingTime, clock.Now().Add(remainingTime).Format("15:04:05"))
			clock.Wait(remainingTime)
		}
	}
}

func doWork(taskList specparser.TaskList, listIndex int, clock *specparser.ClockInterface) {
	fmt.Printf("%s Job %d/%d - ", clock.Now().Format("15:04:05"), listIndex+1, len(taskList.Schedule))
	fmt.Printf("schedule for %s (%s)\n", taskList.Schedule[listIndex].Format("15:04:05"), clock.Until(taskList.Schedule[listIndex]))

	timeUntil := clock.Until(taskList.Schedule[listIndex])

	if timeUntil > (1 * time.Duration(1)) {
		clock.Wait(timeUntil)
	}

	fmt.Printf("%s Job %d/%d - dispatched command @ %s\n", clock.Now().Format("15:04:05"), listIndex+1, len(taskList.Schedule), clock.Now().Format("15:04:05"))
	execCommand(taskList.Work[taskList.Schedule[listIndex]].Command)

	if listIndex < len(taskList.Schedule)-1 {
		doWork(taskList, listIndex+1, clock) // TODO: make non blocking call and move above dispatch
		return
	} else {
		fmt.Println("Done ")
	}

	return
}

func execCommand(command string) {
	fmt.Printf("$ %s\n", command)
}
