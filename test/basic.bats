BIN=./goclitem

@test "Can run the application" {
    run ${BIN}
    echo $output
    [ $status -eq 0 ]
}

