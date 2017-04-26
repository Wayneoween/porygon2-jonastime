package bot

import (
	"fmt"
	"time"

	"github.com/0x263b/porygon2"
)

func jonasTime(command *bot.Cmd, matches []string) (msg string, err error) {
	// layout shows by example how the reference time should be represented.
	const jonasLayout = "01/02/2006 03:04:05 PM"
	const correctLayout = "02.01.2006 15:04:05"

	t := time.Now()
	jonasLocation, _ := time.LoadLocation("US/Mountain")
	timeAtJones := t.In(jonasLocation).Format(jonasLayout)
	realTime := t.Format(correctLayout)
	return fmt.Sprintf("Jns: %s // De: %s", timeAtJones, realTime), nil
}

func init() {
	bot.RegisterCommand("^t$",
		jonasTime)

	bot.RegisterCommand("^time$",
		jonasTime)
}
