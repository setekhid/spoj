module Main where

main = interact f
  where
    f str =
      let (_N:nums) = map read (words str)
          _A = take _N nums
          _L = elis _A
      in show _L

elis' [] = []
elis' (a:_A) =
  let longest _ [] [] = 0
      longest a (a':_A') (s':sums') =
        if a' > a
          then max s' (longest a _A' sums')
          else longest a _A' sums'
      sums = elis' _A
      curr = (longest a _A sums) + 1
  in curr : sums

elis _A = maximum (elis' _A)
