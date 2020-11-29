import React from 'react';

import 'styles/RecordsView.css';
import { inColor, outColor } from 'theme/in_out';

type Record = {in: string, out: string}

function RecordCell({record}: {record: Record}) {
    return (
        <div className="cell">
            <div style={{backgroundColor: inColor}}>{record.in}</div>
            <div style={{backgroundColor: outColor}}>{record.out}</div>
        </div>
    );
}

// add some column grouping imagery (rounded container around cells)
function RecordRow({numEntries}: {numEntries: number}) {
    const records = Array<Record>(numEntries).fill({in: 'ininin', out: 'outout'});
    return (
        <div className="cell-row">
            <div>{new Date().toLocaleDateString()}</div>
            {records.map(record => <RecordCell record={record}/>)}
        </div> 
    );
}

function RecordsView() {
    const rows = Array<number>(30).fill(2);

    return (
        <div className="records-view">
            {rows.map(num  => <RecordRow numEntries={num}/>)}
        </div>
    );
}

export default RecordsView;
