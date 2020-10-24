package main

import (
	"fmt"
	"strings"
)

func (rp *RaidPlayer) toLua() string {
	return fmt.Sprintf("{'%s', '%s', %v}", rp.Name, rp.Spec, rp.Alt)
}

func (p *Party) toLua() string {
	strs := make([]string, len(p.Players))
	for i, player := range p.Players {
		strs[i] = player.toLua()
	}
	return joinLua(strs)
}

func (rp *RaidParty) toLua() string {
	strs := make([]string, len(rp.Parties))
	for i, party := range rp.Parties {
		strs[i] = party.toLua()
	}
	return fmt.Sprintf("{%d, '%s', '%s', %s}", rp.Id, rp.Timing, rp.RaidLeaderName, joinLua(strs))
}

func (lp *LootPriority) toLua() string {
	return fmt.Sprintf("{%d, '%s'}", lp.ItemID, lp.Priority)
}

func (mw *Miniwarlock) toLua() string {
	return fmt.Sprintf("{'%s', '%s', %v}", mw.OriginalName, mw.WarlockName, mw.Ready)
}

func (z *WarlocksZone) toLua() string {
	strs := make([]string, len(z.Warlocks))
	for i, warlock := range z.Warlocks {
		strs[i] = warlock.toLua()
	}
	return fmt.Sprintf("{'%s', %s}", z.Location, joinLua(strs))
}

func (n *News) toLua() string {
	return fmt.Sprintf("{'%s', '%s', %d}", n.Author, strings.ReplaceAll(n.Text, "'", ""), n.Time)
}

func raidPartiesToLua(raids *[]RaidParty) string {
	strs := make([]string, len(*raids))
	for i, raid := range *raids {
		strs[i] = raid.toLua()
	}
	return joinLua(strs)
}

func lootPrioritiesToLua(priorities *[]LootPriority) string {
	strs := make([]string, len(*priorities))
	for i, priority := range *priorities {
		strs[i] = priority.toLua()
	}
	return joinLua(strs)
}

func warlocksToLua(zones *[]WarlocksZone) string {
	strs := make([]string, len(*zones))
	for i, zone := range *zones {
		strs[i] = zone.toLua()
	}
	return joinLua(strs)
}

func newsToLua(news *[]News) string {
	strs := make([]string, len(*news))
	for i, newsItem := range *news {
		strs[i] = newsItem.toLua()
	}
	return joinLua(strs)
}

func joinLua(values []string) string {
	return "{" + strings.Join(values, ", ") + "}"
}

func mapToLua(value map[string]string) string {
	var args []string
	for k, v := range value {
		args = append(args, k + " = " + v)
	}
	return "{" + strings.Join(args, ", ") + "}"
}
