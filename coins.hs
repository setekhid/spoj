module Main where

import Data.IntMap (IntMap)
import qualified Data.IntMap as IntMap

main = interact f
  where
    f str =
      let calc nums =
            foldr
              (\n recu mem ->
                  let iter = coins' n mem
                      ans = IntMap.findWithDefault (toInteger n) n iter
                  in ans : (recu iter))
              (\_ -> [])
              nums
              IntMap.empty
          nums = map read (words str)
          ans = (map show) $ calc nums
      in unlines ans

coins'' n
  | n >= 0 && n < 12 = n
coins'' n
  | n >= 12 && n < 23 =
    let sep n = ((n `div` 2), (n `div` 3), (n `div` 4))
        (a, b, c) = sep n
    in max (a + b + c) n
coins'' n
  | n >= 23 =
    let sep n = ((n `div` 2), (n `div` 3), (n `div` 4))
        (a, b, c) = sep n
    in (coins'' a) + (coins'' b) + (coins'' c)

coins' n mem
  | IntMap.member n mem = mem
coins' n mem
  | n >= 0 && n < 23 = IntMap.insert n (toInteger $ coins'' n) mem
coins' n mem
  | n >= 23 =
    let sep n = ((n `div` 2), (n `div` 3), (n `div` 4))
        (a, b, c) = sep n
        iter = (coins' c) $ (coins' b) $ (coins' a) mem
        va = IntMap.findWithDefault (toInteger a) a iter
        vb = IntMap.findWithDefault (toInteger b) b iter
        vc = IntMap.findWithDefault (toInteger c) c iter
    in IntMap.insert n (va + vb + vc) iter

coins n =
  let mem = coins' n IntMap.empty
  in IntMap.findWithDefault (toInteger n) n mem
