module Main where

main = interact f
  where
    f str =
      let ans:_ = map (willitst . read) (words str)
      in ans

willitst n
  | n <= 1 = "TAK"
willitst n =
  let willitst' n (pow:pows)
        | pow == n = "TAK"
      willitst' n (pow:pows)
        | pow > n = "NIE"
      willitst' n (_:pows) = willitst' n pows
  in willitst' n pow2s

pow2s = foldr (\_ recu pre -> (pre * 2) : (recu $ pre * 2)) (\_ -> []) [1 ..] 1
