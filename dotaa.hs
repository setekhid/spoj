module Main where

pass_Hs m _D [] = m <= 0 || _D <= 0
pass_Hs m _D (_H:_Hs) =
  if m <= 0
    then True
    else pass_Hs (m - ((_H - 1) `div` _D)) _D _Hs

main = interact f
  where
    f str =
      let t:nums = map read (words str)
          test 0 nums = []
          test t (n:m:_D:nums) = stat : (test (t - 1) (drop n nums))
            where
              stat = pass_Hs m _D (take n nums)
          to_YN True = "YES"
          to_YN False = "NO"
          yns = map to_YN (test t nums)
      in unlines yns
