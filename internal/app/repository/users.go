package repository

import (
	"car_bot/internal/app/ds"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"sort"
)

func GetResults(db *sqlx.DB, VkID int) (string, error) {
	var user ds.User

	err := db.QueryRow("SELECT * FROM users WHERE vk_id = $1", VkID).Scan(&user.VkID, &user.State, &user.LongRange, &user.StealthAccuracy, &user.ShootOnMove, &user.Power, &user.ShootInHead, &user.AccuracyDamage, &user.ShootBursts, &user.SmallCost, &user.Reward, &user.CloseRange, &user.AmmoReserve, &user.Light, &user.MachineGun, &user.FastRecharge, &user.LowRecoil, &user.Beautiful, &user.MediumRange, &user.RateOfFire, &user.Aggressive, &user.AccuracySlow)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Row with id unknown")
		} else {
			log.Println("Couldn't find the line with the user")
		}
		log.Error(err)
		return "", err
	}

	weapons := make(map[string]int, 20)

	log.Println(user)

	for _, weapon := range ds.Weapons {
		if user.LongRange == weapon.LongRange {
			weapons[weapon.Name]++
		}
		if user.StealthAccuracy == weapon.StealthAccuracy {
			weapons[weapon.Name]++
		}
		if user.ShootOnMove == weapon.ShootOnMove {
			weapons[weapon.Name]++
		}
		if user.Power == weapon.Power {
			weapons[weapon.Name]++
		}
		if user.ShootInHead == weapon.ShootInHead {
			weapons[weapon.Name]++
		}
		if user.AccuracyDamage == weapon.AccuracyDamage {
			weapons[weapon.Name]++
		}
		if user.ShootBursts == weapon.ShootBursts {
			weapons[weapon.Name]++
		}
		if user.SmallCost == weapon.SmallCost {
			weapons[weapon.Name]++
		}
		if user.Reward == weapon.Reward {
			weapons[weapon.Name]++
		}
		if user.CloseRange == weapon.CloseRange {
			weapons[weapon.Name]++
		}
		if user.AmmoReserve == weapon.AmmoReserve {
			weapons[weapon.Name]++
		}
		if user.Light == weapon.Light {
			weapons[weapon.Name]++
		}
		if user.MachineGun == weapon.MachineGun {
			weapons[weapon.Name]++
		}
		if user.FastRecharge == weapon.FastRecharge {
			weapons[weapon.Name]++
		}
		if user.LowRecoil == weapon.LowRecoil {
			weapons[weapon.Name]++
		}
		if user.Beautiful == weapon.Beautiful {
			weapons[weapon.Name]++
		}
		if user.MediumRange == weapon.MediumRange {
			weapons[weapon.Name]++
		}
		if user.RateOfFire == weapon.RateOfFire {
			weapons[weapon.Name]++
		}
		if user.Aggressive == weapon.Aggressive {
			weapons[weapon.Name]++
		}
		if user.AccuracySlow == weapon.AccuracySlow {
			weapons[weapon.Name]++
		}
	}

	log.Println(weapons)
	res := make([]string, 0, len(weapons))

	for k := range weapons {
		res = append(res, k)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return weapons[res[i]] > weapons[res[j]]
	})

	var output string
	output += "Оружие\tСовместимость\n"

	for i := 0; i < 3; i++ {
		output += fmt.Sprintf("%s\t%d%%\n", res[i], weapons[res[i]]*100/20)
	}

	//for _, v := range res {
	//	output += fmt.Sprintf("%s\t%d%%\n", v, weapons[v]*100/20)
	//}
	return output, nil
}
