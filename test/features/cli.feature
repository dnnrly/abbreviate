Feature: Simple CLI commands

    @Acceptance
    Scenario: Prints help correctly
        When the app runs with parameters "-h"
        Then the app exits without error
        And the app output contains "Usage:"

    @Acceptance
    Scenario: Full abbreviation
        When the app runs with parameters "original strategy-limited"
        Then the app exits without error
        And the app output contains exactly "stg-ltd"

    @Acceptance
    Scenario: No abbreviation
        When the app runs with parameters "original --max 99 strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategy-limited"

    @Acceptance
    Scenario: Some abbreviation
        When the app runs with parameters "original --max 15 strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategy-ltd"

    @Acceptance
    Scenario: More abbreviation
        When the app runs with parameters "original --max 6 strategy-limited"
        Then the app exits without error
        And the app output contains exactly "stg-ltd"

    @Acceptance
    Scenario: Snake case
        When the app runs with parameters "snake --max 15 strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategy_ltd"

    @Acceptance
    Scenario: Kebab case
        When the app runs with parameters "kebab --max 15 strategy_limited"
        Then the app exits without error
        And the app output contains exactly "strategy-ltd"
    
    @Acceptance
    Scenario: Title case
        When the app runs with parameters "title --max 14 strategy_limited"
        Then the app exits without error
        And the app output contains exactly "Strategy Ltd"

    @Acceptance
    Scenario: Pascal case
        When the app runs with parameters "pascal --max 13 strategy-limited"
        Then the app exits without error
        And the app output contains exactly "StrategyLtd"

    @Acceptance
    Scenario: Camel case
        When the app runs with parameters "camel --max 13 strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategyLtd"

    @Acceptance
    Scenario: Camel case convertion
        When the app runs with parameters "camel --max 99 Strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategyLimited"

    @Acceptance
    Scenario: Custom data
        When the app runs with parameters "original --custom fixtures/custom.txt longer"
        Then the app exits without error
        And the app output contains exactly "short"

    @Acceptance
    Scenario: Errors on bad custom data path
        When the app runs with parameters "original --custom ./unknown.txt longer"
        Then the app exits with an error
        And the app output contains "Unable to open custom abbreviations file: open ./unknown.txt: no such file or directory"

    @Acceptance
    Scenario: Prints abbreviations
        When the app runs with parameters "print --custom fixtures/custom.txt"
        Then the app exits without error
        And the app output contains "an=another"
        And the app output contains "short=longer"

    @Acceptance
    Scenario: Separated case
        When the app runs with parameters "separated --max 15 --separator + strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategy+ltd"

    @Acceptance
    Scenario: Separated case, without separator flag
        When the app runs with parameters "separated strategy-limited"
        Then the app exits without error
        And the app output contains exactly "stgltd"

    @Acceptance
    Scenario: Errors on unknown strategy
        When the app runs with parameters "original --strategy oops strategy-limited"
        Then the app exits with an error
        And the app output contains "Error: unknown abbreviation strategy 'oops'. Only allowed lookup, no-abbreviation"

    @Acceptance
    Scenario: Use lookup strategy
        When the app runs with parameters "original --strategy lookup strategy-limited"
        Then the app exits without error
        And the app output contains exactly "stg-ltd"

    @Acceptance
    Scenario: Use no-abbreviation strategy
        When the app runs with parameters "original --strategy no-abbreviation strategy-limited"
        Then the app exits without error
        And the app output contains exactly "strategy-limited"

    @Acceptance
    Scenario: Lookup includes prefixes
        When the app runs with parameters "original --strategy lookup prestrategy-limited"
        Then the app exits without error
        And the app output contains exactly "prstg-ltd"

    @Acceptance
    Scenario: Lookup includes suffixes
        When the app runs with parameters "original --strategy lookup strategy-limitedment"
        Then the app exits without error
        And the app output contains exactly "stg-ltdmnt"

    @Acceptance
    Scenario: Lookup includes prefixes and suffixes
        When the app runs with parameters "original --strategy lookup prestrategy-limitedment"
        Then the app exits without error
        And the app output contains exactly "prstg-ltdmnt"