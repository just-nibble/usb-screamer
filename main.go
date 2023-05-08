package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	last_alert := time.Now()
	cmd := exec.Command("udevadm", "monitor", "--udev")
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("%s install fail\n", err)
	} else {
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				this_alert := time.Now()
				// fmt.Println(scanner.Text())
				if strings.Contains(scanner.Text(), "bind") {
					if this_alert.Sub(last_alert) > 10*time.Second {
						last_alert = this_alert
						// fmt.Println("yamate")
						_, err := exec.Command("paplay", string("usb_audio.ogg")).Output()
						if err != nil {
							fmt.Println(fmt.Sprint(err))
						}

					}

				}

			}
		}()

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}
}
