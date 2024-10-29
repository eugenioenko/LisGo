var demo = String.raw`
;   _        _          _____           _
;  | |      (_)        / ____|         | |
;  | |       _   ___  | |  __    ___   | |
;  | |      | | / __| | | |_ |  / _ \  |_|
;  | |____  | | \__ \ | |__| | | (_) |  _
;  |______| |_| |___/  \_____|  \___/  (_)
;
; Welcome to LisGo a functional programming language inspired by lisp!
; LisGo interpreter is crafted in GoLang and compiled to Wasm
; More info at https://github.com/eugenioenko/LisGo
;
; Hit "RUN" to execute the program

(print (+ 1 2 3))
(print (- 2))
(print (* 2 (/ 2 (- 8 4))))
(print (print "yes"))

(cond
  (0 (print "0"))
  (0 (print "1"))
  (2 (print "2"))
)

(if 1
  (print "yes")
  (print "no")
)

(print "hello world")

(:= index 0)
(print index)
(while (!= index 10)
  (print
    (:= index (+ index 1))
  )
)


(func function (alpha beta)
  (print alpha)
  (return-from "function" "the_return_value")
  (print beta)
)

(function "hello" "world")
(print (function "first" "second"))


(func addition (a b)
  (return (+ a b))
)

(print (addition 100 1))

`;
