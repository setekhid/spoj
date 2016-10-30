module Main where

import Data.Array

main = interact f
  where
    f str =
      let t:nums = map read $ words str
          ans = map (show . fctrl2) $ take t nums
      in unlines ans

fctrl2 = (fctrl100 !)
  where
    fctrl100 =
      array (1, 100) $
      foldr
        (\n recu pre ->
            let fctrl = pre * (toInteger n)
            in (n, fctrl) : (recu fctrl))
        (\_ -> [])
        [1 .. 100]
        1
