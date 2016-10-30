module Main where

main = return ()

lis2' pseq =
  let longest _ [] [] = 0
      longest (x, y) ((x', y'):pseq) (len:lens) =
        if x < x' && y < y'
          then max len (longest (x, y) pseq lens)
          else longest (x, y) pseq lens
      lens [] = []
      lens (pair:pseq) = ((+) 1 $ longest pair pseq lens') : lens'
        where
          lens' = lens pseq
  in maximum $ lens pseq

lis2 pseq = 
