#!/usr/bin/env bats
# vim: ft=bats:list:ts=8:sts=4:sw=4:et:ai:si:

set -u

setup_file() {
    load "../lib/setup_file.sh"
    ./instance-data load
    ./instance-crowdsec start
    API_KEY=$(cscli bouncers add testbouncer -o raw)
    export API_KEY
    CROWDSEC_API_URL="http://localhost:8080"
    export CROWDSEC_API_URL
}

teardown_file() {
    load "../lib/teardown_file.sh"
}

setup() {
    load "../lib/setup.sh"
    if is_db_mysql; then sleep 0.3; fi
}

api() {
    URI="$1"
    curl -s -H "X-Api-Key: ${API_KEY}" "${CROWDSEC_API_URL}${URI}"
}

#----------

@test "cli - first decisions list: must be empty" {
    # delete community pull
    run -0 cscli decisions delete --all
    run -0 --separate-stderr cscli decisions list -o json
    assert_output 'null'
}

@test "API - first decisions list: must be empty" {
    run -0 --separate-stderr api '/v1/decisions'
    assert_output 'null'
}

@test "adding decision for 1.2.3.4" {
    run -0 --separate-stderr cscli decisions add -i '1.2.3.4'
    assert_stderr --partial 'Decision successfully added'
}

@test "CLI - all decisions" {
    run -0 --separate-stderr cscli decisions list -o json
    run -0 jq -r '.[0].decisions[0].value' <(output)
    assert_output '1.2.3.4'
}

@test "API - all decisions" {
    run -0 --separate-stderr api '/v1/decisions'
    run -0 jq -c '[ . | length, .[0].value ]' <(output)
    assert_output '[1,"1.2.3.4"]'
}

# check ip match

@test "CLI - decision for 1.2.3.4" {
    run -0 --separate-stderr cscli decisions list -i '1.2.3.4' -o json
    run -0 jq -r '.[0].decisions[0].value' <(output)
    assert_output '1.2.3.4'
}

@test "API - decision for 1.2.3.4" {
    run -0 --separate-stderr api '/v1/decisions?ip=1.2.3.4'
    run -0 jq -r '.[0].value' <(output)
    assert_output '1.2.3.4'
}

@test "CLI - decision for 1.2.3.5" {
    run -0 --separate-stderr cscli decisions list -i '1.2.3.5' -o json
    assert_output 'null'
}

@test "API - decision for 1.2.3.5" {
    run -0 --separate-stderr api '/v1/decisions?ip=1.2.3.5'
    assert_output 'null'
}

## check outer range match

@test "CLI - decision for 1.2.3.0/24" {
    run -0 --separate-stderr cscli decisions list -r '1.2.3.0/24' -o json
    assert_output 'null'
}

@test "API - decision for 1.2.3.0/24" {
    run -0 --separate-stderr api '/v1/decisions?range=1.2.3.0/24'
    assert_output 'null'
}

@test "CLI - decisions where IP in 1.2.3.0/24" {
    run -0 --separate-stderr cscli decisions list -r '1.2.3.0/24' --contained -o json
    run -0 jq -r '.[0].decisions[0].value' <(output)
    assert_output '1.2.3.4'
}

@test "API - decisions where IP in 1.2.3.0/24" {
    run -0 --separate-stderr api '/v1/decisions?range=1.2.3.0/24&contains=false'
    run -0 jq -r '.[0].value' <(output)
    assert_output '1.2.3.4'
}
