#!/usr/bin/env bash

function init() {
    echo "=========== # start set wallet 1 ============="
    echo "=========== # save seed to wallet ============="
    result=$(./chain33-cli seed save -p fuzwmei1314 -s "tortoise main civil member grace happy century convince father cage beach hip maid merry rib" | jq ".isok")
    if [ "${result}" = "false" ]; then
        echo "save seed to wallet error seed, result: ${result}"
        exit 1
    fi

    sleep 2

    echo "=========== # unlock wallet ============="
    result=$(./chain33-cli wallet unlock -p fuzwmei1314 -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    sleep 2

    echo "=========== # import private key transfer ============="
    result=$(./chain33-cli account import_key -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944 -l transfer | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi
    result=$(./chain33-cli account import_key -k 2AFF1981291355322C7A6308D46A9C9BA311AA21D94F36B43FC6A6021A1334CF -l transfer1 | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi
    result=$(./chain33-cli account import_key -k 2116459C0EC8ED01AA0EEAE35CAC5C96F94473F7816F114873291217303F6989 -l transfer2 | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    sleep 2

    result=$(./chain33-cli account import_key -k 56942AD84CCF4788ED6DACBC005A1D0C4F91B63BCF0C99A02BE03C8DEAE71138 -l toaddr | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== # import private key mining ============="
    result=$(./chain33-cli account import_key -k 4257D8692EF7FE13C68B65D6A52F03933DB2FA5CE8FAF210B5B8B80C721CED01 -l mining | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    sleep 2
    echo "=========== # set auto mining ============="
    result=$(./chain33-cli wallet auto_mine -f 1 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    echo "=========== # end set wallet 1 ============="

}

set -x
init
set +x
