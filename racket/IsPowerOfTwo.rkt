#lang racket
(define/contract (is-power-of-two n)
  (-> exact-integer? boolean?)
  (cond
    
    [(<= n 0) #f]


    [( = n 1) #t]

    
    [(odd? n) #f]

    
    [else (is-power-of-two (/ n 2))]
    )
  )