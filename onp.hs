module Main where

main = interact f
  where
    f str =
      let t':exprs = lines str
          t = read t'
          rpns = map onp $ take t exprs
      in unlines rpns

onp expr = foldr iter (\st -> st) expr []
  where
    pop cond f [] = f []
    pop cond f (s:ops)
      | cond s = f (s : ops)
    pop cond f (s:ops) = s : (pop cond f ops)
    op_prio '^' = 1
    op_prio '*' = 2
    op_prio '/' = 2
    op_prio '+' = 3
    op_prio '-' = 3
    op_prio _ = 250
    iter '(' ff ops = ff ('(' : ops)
    iter ')' ff ops = pop (\s -> s == '(') (\(_:ops) -> ff ops) ops
    iter '^' ff ops = ff ('^' : ops)
    iter s ff ops
      | s == '*' || s == '/' =
        pop (\s -> (op_prio s) > (op_prio '*')) (\ops -> ff (s : ops)) ops
    iter s ff ops
      | s == '+' || s == '-' =
        pop (\s -> (op_prio s) > (op_prio '+')) (\ops -> ff (s : ops)) ops
    iter var ff ops = var : (ff ops)
