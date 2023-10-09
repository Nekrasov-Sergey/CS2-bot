package ds

type User struct {
	VkID            int    `db:"vk_id"`
	State           string `db:"state"`
	LongRange       bool   `db:"long_range"`
	StealthAccuracy bool   `db:"stealth_accuracy"`
	ShootOnMove     bool   `db:"shoot_on_move"`
	Power           bool   `db:"power"`
	ShootInHead     bool   `db:"shoot_in_head"`
	AccuracyDamage  bool   `db:"accuracy_damage"`
	ShootBursts     bool   `db:"shoot_bursts"`
	SmallCost       bool   `db:"small_cost"`
	Reward          bool   `db:"reward"`
	CloseRange      bool   `db:"close_range"`
	AmmoReserve     bool   `db:"ammo_reserve"`
	Light           bool   `db:"light"`
	MachineGun      bool   `db:"machine_gun"`
	FastRecharge    bool   `db:"fast_recharge"`
	LowRecoil       bool   `db:"low_recoil"`
	Beautiful       bool   `db:"beautiful"`
	MediumRange     bool   `db:"medium_range"`
	RateOfFire      bool   `db:"rate_of_fire"`
	Aggressive      bool   `db:"aggressive"`
	AccuracySlow    bool   `db:"accuracy_slow"`
}
