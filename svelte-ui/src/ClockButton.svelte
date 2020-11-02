<script>
    import dayjs from 'dayjs';
    
    currTime = dayjs();
    console.log(currTime)

    // const minuteInMs = 60000;
    function timeInHoursAndMinutes() {
        return currTime.format('HH:mm')
    }
</script>

<style>
    button {

    }
</style>

<button> HEllo
    <!-- {timeInHoursAndMinutes()} -->
</button>
