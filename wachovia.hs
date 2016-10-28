module Main where

import Control.Monad
import Data.Array.ST
import Data.Array

main = interact f
  where
    f str =
      let pair_input (w:v:nums) = (w, v) : (pair_input nums)
          (_N:nums) = map read (words str)
          pairs = pair_input nums
          test 0 _ = []
          test _N ((_K, _M):pairs) =
            (wachovia _K (take _M pairs)) : (test (_N - 1) (drop _M pairs))
          results = test _N pairs
      in unlines $
         map (\r -> "Hey stupid robber, you can get " ++ (show r) ++ ".") results

-- this is a better implementation than knapsack.hs
wachovia _K bags =
  maximum $
  elems $
  runSTArray $
  do sts <- newArray (0, _K) 0
     forM_ bags $
       \(w, v) -> do
         forM_ [0 .. (_K - w)] $
           \_K -> do
             _V1 <- readArray sts _K
             _V2 <- readArray sts (_K + w)
             writeArray sts _K (max _V1 (_V2 + v))
     return sts
