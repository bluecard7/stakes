import React from 'react';
import Greeting from 'components/Greeting';
import Records from 'components/Records';

import 'styles/AppContainer.css';

function AppContainer() {  
    return (
      <div className="stakes-container">
        <Greeting/>
        <Records/>
      </div>
    );
  }
  
export default AppContainer;