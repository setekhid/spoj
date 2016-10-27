main = interact f
  where
    f = unlines . takeWhile (/= "42") . words
