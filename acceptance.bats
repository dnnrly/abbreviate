
@test "Full abbreviation" {
    result="$(./abbreviate original --newline=false strategy-limited)"
    [ $result == "stg-ltd" ]
}

@test "No abbreviation" {
    result="$(./abbreviate original --max 99 --newline=false strategy-limited)"
    [ $result == "strategy-limited" ]
}

@test "Some abbreviation" {
    result="$(./abbreviate original --max 15 --newline=false strategy-limited)"
    [ $result == "strategy-ltd" ]
}

@test "More abbreviation" {
    result="$(./abbreviate original --max 6 --newline=false strategy-limited)"
    [ $result == "stg-ltd" ]
}

@test "Snake case" {
    result="$(./abbreviate snake --max 15 --newline=false strategy-limited)"
    [ $result == "strategy_ltd" ]
}

@test "Snake case with separator" {
    result="$(./abbreviate snake --max 15 --newline=false --separator + strategy-limited)"
    [ $result == "strategy+ltd" ]
}

@test "Pascal case" {
    result="$(./abbreviate pascal --max 13 --newline=false strategy-limited)"
    [ $result == "StrategyLtd" ]
}

@test "Camel case" {
    result="$(./abbreviate camel --max 13 --newline=false strategy-limited)"
    [ $result == "strategyLtd" ]
}

@test "Camel case convertion" {
    result="$(./abbreviate camel --max 99 --newline=false Strategy-limited)"
    [ $result == "strategyLimited" ]
}
