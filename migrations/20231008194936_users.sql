-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users (
   "vk_id" integer not null primary key,
   "state" varchar(255) not null,
    long_range BOOLEAN default false,
    stealth_accuracy BOOLEAN default false,
    shoot_on_move BOOLEAN default false,
    power BOOLEAN default false,
    shoot_in_head BOOLEAN default false,
    accuracy_damage BOOLEAN default false,
    shoot_bursts BOOLEAN default false,
    small_cost BOOLEAN default false,
    reward BOOLEAN default false,
    close_range BOOLEAN default false,
    ammo_reserve BOOLEAN default false,
    light BOOLEAN default false,
    machine_gun BOOLEAN default false,
    fast_recharge BOOLEAN default false,
    low_recoil BOOLEAN default false,
    beautiful BOOLEAN default false,
    medium_range BOOLEAN default false,
    rate_of_fire BOOLEAN default false,
    aggressive BOOLEAN default false,
    accuracy_slow BOOLEAN default false
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
