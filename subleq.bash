#!/usr/local/bin/bash
read -a memory < <(tr $'\n' ' '<"$1")
let ip=0
while (( ip >= 0 && ip < ${#memory[@]} )); do
  let a=${memory[ip]} b=${memory[ip+1]} c=${memory[ip+2]} ip=ip+3
  if (( a == -1 )); then
    read -n 1 memory[b]
  elif (( b == -1 )); then
    printf "\x$(printf %x ${memory[a]})"
  else
    let memory[b]-=memory[a]
    if (( memory[b] <= 0 )); then
      let ip=c
    fi
  fi
done
