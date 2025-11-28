program twoSum
    implicit none
    integer, allocatable :: numbers(:)
    integer :: result(2)
    integer :: target, arrSize
    target = 9

    numbers = [2,7,11,15,7]
    result = twoSumFunc(numbers, target)
    write(*,'(I0)') result
    contains

    function twoSumFunc(nums, target) result(theSum)
        integer:: theSum(2)
        integer, allocatable, intent(in) :: nums(:)
        integer, intent(in) :: target

        integer :: i,j
        do i=1,size(nums)
            do j=i+1,size(nums)
                if(nums(i)+nums(j)==target) then
                    theSum(1) = i
                    theSum(2) = j
                    exit
                end if
            end do
        end do

    end function twoSumFunc

end program twoSum