program progAddTwoNums
    implicit none

    type :: Listnode
        integer :: Value
        type(Listnode), pointer :: Next
    end type

    type(Listnode), target :: l10, l11, l12, l20, l21, l22
    type(Listnode), pointer :: head1, head2, resultPtr, node

    nullify(l10%Next, l11%Next, l12%Next, l20%Next, l21%Next, l22%Next)

    l10%Value = 2
    l10%Next => l11
    l11%Value = 4
    l11%Next => l12
    l12%Value = 3
    l12%Next => null()

    l20%Value = 5
    l20%Next => l21
    l21%Value = 6
    l21%Next => l22
    l22%Value = 4
    l22%Next => null()

    head1 => l10
    head2 => l20

    resultPtr => addTwoNumbers(head1,head2)

    node => resultPtr
    do while(associated(node))
        write(*,'(I0)', advance='no') node%Value
        node => node%Next    
    end do
    write(*,*)
    
    


    contains 
    function addTwoNumbers(l1,l2) result(result)
        implicit none
        type(Listnode), pointer :: l1,l2, result, current, p1, p2, newNode
        type(Listnode), target :: dummy
        integer :: carry, sum

        carry = 0

        nullify(dummy%Next)
        current => dummy
        p1 => l1
        p2 => l2


        do while (associated(p1) .or. associated(p2) .or. carry/=0)
            sum = carry
            if (associated(p1)) then
                sum = sum + p1%Value
                p1 => p1%Next
            end if
            if (associated(p2)) then
                sum = sum + p2%Value
                p2 => p2%Next
            end if

            allocate(newNode)
            newNode%Value = modulo(sum,10)
            nullify(newNode%Next)
            current%Next => newNode
            current => current%Next
            carry = sum/10

        end do


        result => dummy%Next
    end function
    
end program progAddTwoNums
