module Main where

main = interact f
  where
    f str =
      let t:nums = map read (words str)
          sums = take t (map divsum nums)
      in unlines (map show sums)

divsum :: Int -> Integer
divsum num =
  let factors = strip_primes primes num
      f prod (_prime, pow) = fac * prod
        where
          fac = truncate ((prime ^ (pow + 1) - 1) / (prime - 1))
          prime = fromIntegral _prime
  in (foldl f 1 factors) - (fromIntegral num)

strip_primes _ 1 = []
strip_primes (prime:primes) num =
  if (fromIntegral prime) > ((sqrt . fromIntegral) num)
    then [(num, 1)]
    else let (rest, pow) = strip_prime prime num
         in (prime, pow) : (strip_primes primes rest)
  where
    strip_prime prime num =
      if num `mod` prime /= 0
        then (num, 0)
        else let (rest, pow) = strip_prime prime (num `div` prime)
             in (rest, pow + 1)

primes = filter (prime_cmp primes) [2 ..]
  where
    prime_cmp _ 2 = True
    prime_cmp (factor:factors) num =
      if (fromIntegral factor) > ((sqrt . fromIntegral) num)
        then True
        else (num `mod` factor /= 0) && prime_cmp factors num
