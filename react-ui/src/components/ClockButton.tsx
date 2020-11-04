import React, { useEffect, useState } from 'react';
import dayjs from 'dayjs';

import 'styles/ClockButton.css'

/**
 * Possible 'features':
 *  Allow user to change default colors for clock in and out
 */
function ClockButton() {
    const [clockIn, setClockIn] = useState(true);
    const [time, setTime] = useState(dayjs());
    // useEffect(() => {
    //     const interval = setInterval(() => {

    //     }, );
    //     return () => clearInterval(interval);
    // });

    const forestGreen = '#228B22';
    const maroon = '#800000';

    const toggleInOut = () => setClockIn(!clockIn);
    const buttonColor = { backgroundColor: `${clockIn ? forestGreen : maroon}` }

    return (
        <button
            className="rounded-corners sizable" 
            onClick={toggleInOut}
            style={buttonColor}
        >
            {time.format('HH:mm')}
        </button>
    )
}

export default ClockButton;
