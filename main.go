package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/wule61/go-mysql/mysql"
	"github.com/wule61/go-mysql/replication"
)

func main() {
	// Create a binlog syncer with a unique server id, the server id must be different from other MySQL's.
	// flavor is mysql or mariadb
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "wule",
		Password: "Wule!#%&",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	// Start sync with specified binlog file and position
	streamer, _ := syncer.StartSync(mysql.Position{"mysql-bin.000187", 0})

	// or you can start a gtid replication like
	// streamer, _ := syncer.StartSyncGTID(gtidSet)
	// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	// the mariadb GTID set likes this "0-1-100"

	for {

		ev, _ := streamer.GetEvent(context.Background())
		// Dump event
		buf := new(bytes.Buffer)
		ev.Dump(buf)
		if buf.Len() != 0 {
			fmt.Println(buf.String())
		}
		//println(ev.Header.EventType.String())
		//ev.Header.Dump(os.Stdout)
		//print(ev.Event.Decode(ev.RawData))
	}

	// or we can use a timeout context
	//for {
	//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//	ev, err := s.GetEvent(ctx)
	//	cancel()
	//
	//	if err == context.DeadlineExceeded {
	//		// meet timeout
	//		continue
	//	}
	//
	//	ev.Dump(os.Stdout)
	//}
}
