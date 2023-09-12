package job

import (
	"fmt"

	"github.com/baowk/dilu-core/core"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var c = cron.New(cron.WithSeconds())

func Add(spec string, cmd func()) {
	_, err := c.AddFunc(spec, cmd)
	if err != nil {
		core.Log.Error("job start err", zap.Error(err))
	}
}

func Start() {
	c.Start()
}

func Stop() {
	c.Stop()
}

func Remove(id int) {
	c.Remove(cron.EntryID(id))
}

func Jobs() []cron.Entry {
	return c.Entries()
}

func DemoTask() {
	fmt.Println("Haha")
}
