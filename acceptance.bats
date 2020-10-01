
@test "Full abbreviation" {
    run ./abbreviate original strategy-limited
    [ "${lines[0]}" == "stg-ltd" ]
}

@test "No abbreviation" {
    run ./abbreviate original --max 99 strategy-limited
    [ "${lines[0]}" == "strategy-limited" ]
}

@test "Some abbreviation" {
    run ./abbreviate original --max 15 strategy-limited
    [ "${lines[0]}" == "strategy-ltd" ]
}

@test "More abbreviation" {
    run ./abbreviate original --max 6 strategy-limited
    [ "${lines[0]}" == "stg-ltd" ]
}

@test "Snake case" {
    run ./abbreviate snake --max 15 strategy-limited
    [ "${lines[0]}" == "strategy_ltd" ]
}

@test "Kebab case" {
    run ./abbreviate kebab --max 15 strategy_limited
    [ "${lines[0]}" == "strategy-ltd" ]
}

@test "Pascal case" {
    run ./abbreviate pascal --max 13 strategy-limited
    [ "${lines[0]}" == "StrategyLtd" ]
}

@test "Camel case" {
    run ./abbreviate camel --max 13 strategy-limited
    [ "${lines[0]}" == "strategyLtd" ]
}

@test "Camel case convertion" {
    run ./abbreviate camel --max 99 Strategy-limited
    [ "${lines[0]}" == "strategyLimited" ]
}

@test "Custom data" {
    echo short=longer > custom.txt
    run ./abbreviate original --custom ./custom.txt longer
    [ "${lines[0]}" == "short" ]
}

@test "Errors on bad custom data path" {
    run ./abbreviate original --custom ./unknown.txt longer
    [ "$status" -eq 1 ]
    [ "${lines[0]}" = "Unable to open custom abbreviations file: open ./unknown.txt: no such file or directory" ]
    rm -f custom.txt
}

@test "Prints abbreviations" {
    echo short=longer > custom.txt
    echo an=another >> custom.txt
    run ./abbreviate print --custom ./custom.txt
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "an=another" ]
    [ "${lines[1]}" = "short=longer" ]
    rm -f custom.txt
}

@test "Separated case" {
    run ./abbreviate separated --max 15 --separator + strategy-limited
    [ "${lines[0]}" == "strategy+ltd" ]
}

@test "Separated case, without separator flag" {
    run ./abbreviate separated strategy-limited
    [ "${lines[0]}" == "stgltd" ]
}

@test "Errors on unknown strategy" {
    run ./abbreviate original --strategy oops strategy-limited
    [ "$status" -eq 1 ]
    [ "${lines[0]}" = "Error: unknown abbreviation strategy 'oops'. Only allowed lookup" ]
}

@test "Select default strategy" {
    run ./abbreviate original --strategy lookup strategy-limited
    [ "${lines[0]}" == "stg-ltd" ]
}
