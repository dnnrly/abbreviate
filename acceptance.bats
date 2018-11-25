
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
