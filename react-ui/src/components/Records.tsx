import React from 'react';
import { addDays } from 'date-fns';
import { useState } from 'react';
import ClockButton from 'components/ClockButton';
import RecordsView from 'components/RecordsView';
import { DateRangePicker } from 'react-date-range';

import 'react-date-range/dist/styles.css';
import 'react-date-range/dist/theme/default.css';

import 'styles/Records.css';

/**
 * Handles retrieving time records from stakes backend and setting 
 * ClockButton and RecordsView accrodingly.
 */
function Records() {

    const [state, setState] = useState([
        {
          startDate: new Date(),
          endDate: addDays(new Date(), 7),
          key: 'selection'
        }
    ]);

    return (
        <div className="records">
            <ClockButton/>
            <div style={{display: 'flex', flexFlow: 'row'}}>
                <DateRangePicker
                    onChange={(item: any) => setState([item.selection])}
                    minDate={addDays(new Date(), -365)}
                    maxDate={new Date()}
                    ranges={state}
                    direction="vertical"
                    scroll={{ enabled: true }}
                />
                <RecordsView/>
            </div>
        </div>
    )
}

export default Records;
