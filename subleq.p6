#!/usr/bin/env perl6
sub MAIN($file) {
  my @memory = slurp($file).words;
  my $ip = 0;
  while $ip >= 0 && $ip < @memory {
    my ($a, $b, $c) = @memory[$ip,$ip+1,$ip+2];
    $ip += 3;
    if $a < 0 {
      @memory[$b] = $*IN.getc.ord;
    } elsif $b < 0 {
      print @memory[$a].chr;
    } else {
      if (@memory[$b] -= @memory[$a]) <= 0 {
        $ip = $c;
      } 
    }
  }
}
