import React from 'react';
import Greeting from 'components/Greeting';
import ClockButton from 'components/ClockButton';
import RecordsView from 'components/RecordsView';

import 'styles/AppContainer.css';

function AppContainer() {  
    return (
      <div className="stakes-container">
        <Greeting/>
        <ClockButton/>
        <RecordsView/>
      </div>
    );
  }
  
export default AppContainer;