import React, { useEffect, useState } from 'react';
import dayjs from 'dayjs';

import 'styles/ClockButton.css'
import { inColor, toggleColor } from 'theme/in_out';

function ClockButton() {
    // needs to be set based on what data was passed
    const [btnColor, setBtnColor] = useState(inColor);
    const [time, setTime] = useState(dayjs());
    // useEffect(() => {
    //     const interval = setInterval(() => {

    //     }, );
    //     return () => clearInterval(interval);
    // });
    return (
        <button
            className="rounded-corners sizable" 
            onClick={() => setBtnColor(toggleColor(btnColor))}
            style={{ backgroundColor: btnColor }}
        >
            {time.format('HH:mm')}
        </button>
    )
}

export default ClockButton;
