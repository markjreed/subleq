#!/usr/bin/env python
import sys
memory = [int(val) for val in open(sys.argv[1]).read().replace("\n"," ")[:-1].split(' ')]
ip = 0
while ip >= 0 and ip < len(memory):
  a, b, c = memory[ip:ip+3]
  ip += 3
  if a < 0:
    memory[b] = ord(sys.stdin.read(1))
  elif b < 0:
    sys.stdout.write(chr(memory[a]))
  else:
    memory[b] -= memory[a]
    if memory[b] <= 0:
      ip = c;
