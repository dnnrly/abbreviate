
@test "Full abbreviation" {
    result="$(./abbreviate --newline=false strategy-limited)"
    [ $result == "stg-ltd" ]
}

@test "No abbreviation" {
    result="$(./abbreviate --max 99 --newline=false strategy-limited)"
    [ $result == "strategy-limited" ]
}

@test "Some abbreviation" {
    result="$(./abbreviate --max 15 --newline=false strategy-limited)"
    [ $result == "strategy-ltd" ]
}

@test "More abbreviation" {
    result="$(./abbreviate --max 6 --newline=false strategy-limited)"
    [ $result == "stg-ltd" ]
}

@test "Snake case" {
    result="$(./abbreviate --max 15 --newline=false --style snake strategy-limited)"
    [ $result == "strategy_ltd" ]
}

@test "Snake case with seperator" {
    result="$(./abbreviate --max 15 --newline=false --style snake --seperator + strategy-limited)"
    [ $result == "strategy+ltd" ]
}

@test "Pascal case" {
    result="$(./abbreviate --max 15 --newline=false --style pascal strategy-limited)"
    [ $result == "StrategyLtd" ]
}

@test "Camel case" {
    result="$(./abbreviate --max 15 --newline=false --style pascal strategy-limited)"
    [ $result == "strategyLtd" ]
}

@test "Camel case convertion" {
    result="$(./abbreviate --max 99 --newline=false --style pascal strategy-limited)"
    [ $result == "strategyLimited" ]
}
