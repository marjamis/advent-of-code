#!/bin/bash -e

COMMAND_LINE_OPTIONS_HELP="Advent of Code helper"

ADVENT_SESSION=$(cat ~/.advent-of-code_cookie)
TEST_DIR="./test"
CODE_DIR="./internal/pkg"

function help {
    echo "Usage: helper -h for help";
    echo "$COMMAND_LINE_OPTIONS_HELP"
}

function error {
    printf "Error: $1\n"

    if [ ! "$3" == "no" ] ; then
        printf "Printing help...\n\n"
        help
    fi

    exit $2
}

function action_new_day {
    year=$(gum input --prompt "What year? " --value $(date "+%Y"))
    day=$(gum input --prompt "What day? " --value $(date "+%e"))

    if [ "$day" == "1" ] ; then
        mkdir $TEST_DIR/advent$year/
        mkdir $CODE_DIR/advent$year/
    fi

    curl --cookie "session=$ADVENT_SESSION" -o $TEST_DIR/advent$year/day$day.txt https://adventofcode.com/$year/day/$day/input

    gh issue create -R marjamis/advent-of-code --title "Year $year - Day $day" --label "enhancement" --assignee @me --body "Complete the days challenge."

    echo "package advent$year" > $CODE_DIR/advent$year/day$day.go
    echo "package advent$year" > $CODE_DIR/advent$year/day$(echo $day)_test.go
}

function select_action {
    LIST="New Day
"

    ACTION=$(echo -e "$LIST" | gum choose)

    case "$ACTION" in
        "New Day") action_new_day ;;
        \?) error "E_OPTERROR_UNKNOWNOPTION" 2 ;;
        *) error "E_OPTERROR_NOOPTION" 3 ;;
    esac
}

# Start of script execution

# Check if gum is installed and if not exit with details on how to install
if [ ! $(command -v gum) ] ; then
    error "Gum is not installed. Please refer to the documentation: https://github.com/charmbracelet/gum" 1 "no"
fi

while getopts ":h" flag ; do
    case "$flag" in
        h) help ;;
        \?) error "E_OPTERROR_UNKNOWNOPTION" 2 ;;
        *) error "E_OPTERROR_NOOPTION" 3 ;;
    esac
done

# Checks that a flag was provided and runs the script proper
if [ "$#" == 0 ]; then
   select_action
fi
