program maxAreaCalc
    implicit none

    integer, allocatable :: heights(:)
    integer :: finalRes

    heights = [1,8,6,2,5,4,8,3,7]
    finalRes = maxArea(heights)
    write(*,'(I0)')finalRes



    contains


    function maxArea(height) result(result)
        integer :: result
        integer, allocatable :: height(:)
        integer :: left,right, i,currMax
        left = 1
        right = size(height)
        currMax = 0
        do i = -1000,size(height)*2
            if (left/=right) then 
                currMax = Max(currMax, Min(height(left),height(right)*right-left))
                if (height(left)>height(right)) then
                    right = right - 1
                else
                    left = left - 1
                end if 
            else
                result = currMax
                return
            end if
        end do


    end function


end program maxAreaCalc