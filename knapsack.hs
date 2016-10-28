module Main where

import Data.Map (Map)
import qualified Data.Map as Map

import System.Random

-- main for spoj test
main0 = interact f
  where
    f str =
      let _S:_N:nums = map read (words str)
          pair_items 0 _ = []
          pair_items _N (s:v:nums) = (s, v) : (pair_items (_N - 1) nums)
          items = pair_items _N nums
          value = knapsack1 _S items
      in show value

-- compare results of knapsack0 and knapsack1
main = run (take 70 (repeat tester))
  where
    run [] = return ()
    run (a:l) = a >> (run l)

tester = do
  _S:_N:nums <- case_gen
  let items = pair_items _N nums
  putStrLn
    ((show _S) ++
     "\t" ++
     (show _N) ++ "\t" ++ (show ((knapsack0 _S items) == (knapsack1 _S items))))
  where
    pair_items 0 _ = []
    pair_items _N (s:v:nums) = (s, v) : (pair_items (_N - 1) nums)
    rand_ls (a, b) n = sequence (take n (repeat (randomRIO (a, b))))
    case_gen = do
      (_S:_N:[]) <- rand_ls (1, 70) 2
      nums <- rand_ls (1, 70) (_N * 2)
      return (_S : _N : nums)

-- haskell will not memoize function result automatically, so this solution is
-- O(2^N)
knapsack0 _ [] = 0
knapsack0 0 _ = 0
knapsack0 _S ((s, v):items) =
  if s > _S
    then knapsack0 _S items
    else max (knapsack0 _S items) (v + (knapsack0 (_S - s) items))

-- this solution iterating the binary tree branched by every chosen of items.
-- each level has maximal S width, and the tree has N depth, so time cost is
-- less than O(S*N) and space cost is only O(S)
--
-- this implementation compares with knapsack.go, the only difference is the
-- status storage, golang implementation using an array, here it's a map and
-- created at every iterating level
--
-- a better implementation in wachovia.hs, using STArray replacing Map
knapsack1_ _SVs [] = foldl (\_V1 (_, _V2) -> max _V1 _V2) 0 _SVs
knapsack1_ _SVs (item:items) = knapsack1_ (f item _SVs) items
  where
    f (s, v) _SVs =
      let flatmap f l = foldr (\a b -> (f a) ++ b) [] l
          iterated =
            flatmap
              (\(_S, _V) ->
                  (_S, _V) :
                  (if s > _S
                     then []
                     else [(_S - s, _V + v)]))
              _SVs
          reduced = Map.fromListWith max iterated
      in Map.toList reduced

knapsack1 _S items = knapsack1_ [(_S, 0)] items
