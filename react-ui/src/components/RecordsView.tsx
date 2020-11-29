import React from 'react';
import { DateRangePicker } from 'react-date-range';

import 'react-date-range/dist/styles.css';
import 'react-date-range/dist/theme/default.css';

type Record = {in: string, out: string}

// would be cool to pull events and display some icon for certain holidays
function RecordCell({record}: {record: Record}) {
    return (
        <div style={{width: 100, margin: 10, color: '#fff', textAlign: 'center'}}>
            <div style={{
                backgroundColor: '#606060',
            }}>{new Date().toLocaleDateString()}</div>
            <div style={{
                backgroundColor: '#228B22',
            }}>{record.in}</div>
            <div style={{
                backgroundColor: '#800000',
            }}>{record.out}</div>
        </div>
    );
}

// add some column grouping imagery (rounded container around cells)
function RecordColumn({numEntries}: {numEntries: number}) {
    const records = Array<Record>(numEntries).fill({in: 'ininin', out: 'outout'});
    return (
        <div style={{
            display: 'flex',
            flexFlow: 'column',
        }}>
            {
                records.map(record => <RecordCell record={record}/>)
            }
        </div> 
    );
}

function RecordsView() {
    const columns = Array<number>(5).fill(2);

    return (
        <div style={{
            display: 'flex',
            flexFlow: 'row',
        }}>
            {
                columns.map(num  => <RecordColumn numEntries={num}/>)
            }
        </div>
    );
}

export default RecordsView;
