package checkin

import (
	"testing"
	"time"

	"github.com/berryhe/ding"
)

/*
   "appID":"839620577",
    "AppKey":"dingks0qa7897fttexqv",
    "AppSecret":"tyyBVO1oZ4qRfV1oQQYKTlgj9Ta582ucxfsTHhbuxKrinsPJ7OIVcbSqmjuzQNrR"
*/

func TestGetCheckin(t *testing.T) {
	dConfig := ding.Config{
		AgentID:      12345678,
		AppKey:       "dingAppkey",
		AppSecretKey: "dingAppSecret",
	}
	dCtx := ding.NewDCtx(dConfig)
	workDataFrom := time.Now().Format("2006-01-02") + " 09:00:00"
	workDataTo := time.Now().Format("2006-01-02") + " 17:30:00"
	dingIDs := []string{
		"122461524923443",
		"0354036337-121231",
		"05203467542",
		"0322200949657744",
		"1505450058232",
		"06463224549084",
		"1459051028781107",
		"0354036337-211403",
		"06683852-283968",
	}
	//
	//arc := entity.AttendanceCheckinListRequest{
	//	WorkDateFrom: workDataFrom,
	//	WorkDateTo:   workDataTo,
	//	UserIDList:   dingIDs,
	//	Limit:        50,
	//	Offset:       0,
	//	IsI18N:       false,
	//}
	resp, err := LoopAttendanceCheckinList(dCtx, dingIDs, workDataFrom, workDataTo, false)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", resp)
}
