package main

import (
	"myHome/gin/configs"
	"myHome/gin/models/score"
	"myHome/gin/models/sys"
	"myHome/gin/services/stores"
	"myHome/gin/utils"
)

func main() {

	configs.InitSettings()

	// genTenantData()
	// genMemberData()
	// genMemberScore()

	defer configs.ReleaseSettings()
}

func genMemberScore() {
	ms := &score.MemberScore{
		MemberId:   "729c459f-36aa-4167-bb7e-dc260bbc0b88",
		MemberName: "王子安",
		Score:      0,
	}
	stores.NewModel(ms)
}

func genMemberData() {
	f := &sys.Member{
		MemberId:       utils.GetUUID(),
		TenantId:       "9cfe627d-1513-435b-9c64-6c2acbf5d958",
		MemberName:     "文涛",
		MemberType:     "dad",
		MemberTypeDesc: "爸爸",
	}
	stores.NewModel(f)
	m := &sys.Member{
		MemberId:       utils.GetUUID(),
		TenantId:       "9cfe627d-1513-435b-9c64-6c2acbf5d958",
		MemberName:     "睿睿",
		MemberType:     "mom",
		MemberTypeDesc: "妈妈",
	}
	stores.NewModel(m)

	c := &sys.Member{
		MemberId:       utils.GetUUID(),
		TenantId:       "9cfe627d-1513-435b-9c64-6c2acbf5d958",
		MemberName:     "王子安",
		MemberType:     "children",
		MemberTypeDesc: "孩子",
	}
	stores.NewModel(c)

}

func genTenantData() {

	tenant := &sys.Tenant{
		TenantId:   utils.GetUUID(),
		TenantName: "My Home",
		TenantType: "family",
	}
	stores.NewModel(tenant)
}
